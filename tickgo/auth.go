package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	TickTickAuthURL = "https://ticktick.com/oauth/authorize"
)

type Auth struct {}

type AuthArgs struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (a *Auth) Login(args *AuthArgs, reply *bool) error {
	// redirect user to ticktick auth page
	// needs client_id, scope, state, redirect_uri, and response_type
	urlValues := url.Values {}
	urlValues.Add("client_id", args.ClientID)
	urlValues.Add("scope", "tasks:write tasks:read")
	urlValues.Add("state", "state") // TODO: needs actual random state
	urlValues.Add("redirect_uri", "http://127.0.0.1:9090/callback") // TODO: needs to be an env var
	urlValues.Add("response_type", "code")

	openURL(TickTickAuthURL + "?" + urlValues.Encode())
	
	// wait for callback by blocking on a channel
	callback := make(chan bool)
	
	go callbackServer(callback)

	res := <- callback
	*reply = res

	fmt.Printf("[Auth.Login] set result on reply to %v\n", *reply)

	return nil
}

func callbackServer(callback chan bool) {
	mux := http.NewServeMux()

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {

		code := r.URL.Query().Get("code")
		if code == "" {
			log.Printf("no code found on the redirect_uri\n")
			callback <- false
			w.WriteHeader(http.StatusBadRequest)
		}

		state := r.URL.Query().Get("state") 
		if state == "" {
			log.Printf("no state found on the redirect_uri\n")
			callback <- false
			w.WriteHeader(http.StatusBadRequest)
		}

		log.Printf("code: %s, state: %s\n", code, state)
		callback <- true

		res := []byte(`<html>
	<p>you can close this tab</p>
</html>`)

		w.Write(res)
	})

	log.Printf("callback server started on :9090\n")
	if err := http.ListenAndServe(":9090", mux); err != nil {
		log.Printf("error on callback server: %v\n", err)
	}
}
