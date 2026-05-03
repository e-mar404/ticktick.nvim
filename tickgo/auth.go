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

type TickTickTokenRes struct {
	AccessToken      string `json:"access_token"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// TODO: reply should be of a type that will allow for errors to be passed and displayed on neovim lua side
func (a *Auth) Login(args *AuthArgs, reply *bool) error {
	// TODO: save auth creds here?

	if err := openOAuthPage(args.ClientID); err != nil {
		*reply = false
		log.Printf("unable to open OAuth Page: %v\n", err)
	}

	// wait for code from OAuth page by blocking on a channel
	callback := make(chan TickTickOAuthRes)
	go startCallbackServer(callback)
	oauthRes := <-callback

	if oauthRes.Err != nil {
		*reply = false
		fmt.Printf("[Auth.Login] set result on reply to %v\n", *reply)
		return oauthRes.Err
	}

	// TODO: make POST call to TickTickOAuthURL + /token to get access token
	tokenRes := fetchAccessToken(args.ClientID, args.ClientID, oauthRes.Code)
	log.Printf("got access token: %s\n", tokenRes.AccessToken)

	// TODO: check for tokenRes.Err

	*reply = true
	fmt.Printf("[Auth.Login] set result on reply to %v\n", *reply)

	return nil
}

func fetchAccessToken(clientID, clientSecret, code string) TickTickTokenRes {
	urlValues := url.Values{}

	urlValues.Add("client_id", clientID)
	urlValues.Add("client_secret", clientSecret)
	urlValues.Add("code", code)
	urlValues.Add("grant_type", "authorization_code")
	urlValues.Add("scope", "tasks:write tasks:read")
	urlValues.Add("redirect_uri", "http://127.0.0.1:9090/callback")

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
