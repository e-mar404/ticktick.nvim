package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	TickTickOAuthURL = "https://ticktick.com/oauth"
)

var (
	ErrNoCodeProvided   = errors.New("No code given by redirect")
	ErrStateMismatch    = errors.New("Either state is missing or there is a missmatch on state passed by redirect")
	ErrTokenNotRecieved = errors.New("Access token not received from TickTick")
)

type Auth struct{}

type AuthArgs struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TickTickOAuthRes struct {
	Code string
	Err  error
}

// TODO: at some point this will also need to get the "expires_in" so it can be saved and we can query the saved access token to see if it needs a refresh
type TickTickTokenRes struct {
	AccessToken      string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (a *Auth) Login(args *AuthArgs, reply *bool) error {
	// TODO: save auth creds here?

	if err := openOAuthPage(args.ClientID); err != nil {
		*reply = false
		log.Printf("unable to open OAuth Page: %v\n", err)
	}

	callback := make(chan TickTickOAuthRes)
	go startCallbackServer(callback)
	oauthRes := <-callback

	if oauthRes.Err != nil {
		*reply = false
		fmt.Printf("[Auth.Login] set result on reply to %v\n", *reply)
		return oauthRes.Err
	}

	tokenRes := fetchAccessToken(args.ClientID, args.ClientSecret, oauthRes.Code)
	log.Printf("got access token: %s\n", tokenRes.AccessToken)

	if tokenRes.Error != "" {
		*reply = false
		fmt.Printf("[Auth.Login] set result on reply to %v\n", *reply)
		log.Printf("Err: %s\nDesc: %s\n", tokenRes.Error, tokenRes.ErrorDescription)
		return fmt.Errorf("%s", tokenRes.Error)
	}

	// TODO: once access token is successfully gotten save it to disk along side the expires_in field
	*reply = true
	log.Printf("[Auth.Login] set result on reply to %v\n", *reply)

	return nil
}

func fetchAccessToken(clientID, clientSecret, code string) TickTickTokenRes {
	urlValues := url.Values{}
	redirectURI := fmt.Sprintf("http://127.0.0.1%s/callback", CALLBACK_PORT)

	urlValues.Add("client_id", clientID)
	urlValues.Add("client_secret", clientSecret)
	urlValues.Add("code", code)
	urlValues.Add("grant_type", "authorization_code")
	urlValues.Add("scope", "tasks:write tasks:read")
	urlValues.Add("redirect_uri", redirectURI)

	u := TickTickOAuthURL + "/token?" + urlValues.Encode()

	req, err := http.NewRequest("POST", u, nil)
	if err != nil {
		return TickTickTokenRes{
			Error: err.Error(),
		}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return TickTickTokenRes{
			Error: err.Error(),
		}
	}

	if res.StatusCode != http.StatusOK {
		return TickTickTokenRes{
			Error: res.Status,
		}
	}

	defer res.Body.Close()

	tokenRes := TickTickTokenRes{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&tokenRes); err != nil {
		return TickTickTokenRes{
			Error: err.Error(),
		}
	}

	return tokenRes
}
