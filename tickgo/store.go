package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/adrg/xdg"
)

type AuthStore struct {
	path string
}

type AuthState struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
}

// All data will go to xdg_data_dir, there is no configuration that should be saved from here
func (aStore *AuthStore) save(state AuthState) error {
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		log.Printf("unable to marshal obj: %v\n", err)
		return err
	}

	return os.WriteFile(aStore.path, data, 0644)
}

func NewAuthStore() *AuthStore {
	path, err := xdg.DataFile("tickgo/auth_store.json")
	if err != nil {
		log.Printf("auth store not fully loaded, path is blank: %v\n", err)
		return nil
	}

	return &AuthStore{
		path: path,
	}
}
