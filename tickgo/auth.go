package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	TickTickOAuthURL = "https://ticktick.com/oauth"
)

var (
	ErrNoCodeProvided   = errors.New("No code given by redirect")
	ErrStateMismatch    = errors.New("Either state is missing or state passed by redirect is not the one set by the tickgo server")
	ErrTokenNotRecieved = errors.New("Access token not received from TickTick")
)

type Auth struct {
	store *AuthStore
}

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
	ExpiresIn        int    `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func init() {
	rpcServices["Auth"] = NewAuthService()
}

func (a *Auth) Login(args *AuthArgs, reply *bool) error {
	state := generateState()

	if err := openOAuthPage(args.ClientID, state); err != nil {
		*reply = false
		l.Errorf("unable to open OAuth Page: %v\n", err)
	}

	callback := make(chan TickTickOAuthRes)
	go startCallbackServer(callback, state)
	oauthRes := <-callback

	if oauthRes.Err != nil {
		*reply = false
		fmt.Printf("[Auth.Login] set result on reply to %v\n", *reply)
		return oauthRes.Err
	}

	tokenRes := fetchAccessToken(args.ClientID, args.ClientSecret, oauthRes.Code)
	l.Debug("successfully retrieved access token", "access_token", tokenRes.AccessToken)

	if tokenRes.Error != "" {
		*reply = false
		fmt.Printf("[Auth.Login] set result on reply to %v\n", *reply)
		l.Error("Received error from last OAuth step",
			"Err", tokenRes.Error,
			"Description", tokenRes.ErrorDescription,
		)
		return fmt.Errorf("%s", tokenRes.Error)
	}

	authState := AuthState{
		ClientID:     args.ClientID,
		ClientSecret: args.ClientSecret,
		AccessToken:  tokenRes.AccessToken,
		ExpiresIn:    tokenRes.ExpiresIn,
	}

	if err := a.store.save(authState); err != nil {
		l.Errorf("unable to save login credentials: %v\n", err)
		return err
	}

	*reply = true
	l.Info("[Auth.Login] responded", "reply", *reply)

	return nil
}

func NewAuthService() *Auth {
	store := NewAuthStore()
	if store == nil {
		return nil
	}

	return &Auth{
		store: store,
	}
}

func generateState() string {
	buf := make([]byte, 32)
	rand.Read(buf)
	return base64.URLEncoding.EncodeToString(buf)
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
