package main

import (
	"encoding/json"
	"net/http"
)

// TODO: See if there is a way to call this just Project
type Project struct{}

// TODO: see if I can change the name to ProjectList or something similar to free up the 'Project' name
type ProjectList []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	// TODO register service to rpc server
	rpcServices["Project"] = NewProjectService()
}

func (p *Project) GetAll(args *any, reply *ProjectList) error {
	state := &AuthState{}
	if err := NewAuthStore().load(state); err != nil {
		logger.Error("Could not retrieve saved credentials", "error", err)
		return err
	}

	u := API_URL + "/project"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		logger.Error("Unable to create project request", "error", err)
		return err
	}

	req.Header.Add("Host", API_HOST)
	req.Header.Add("Authorization", "Bearer "+state.AccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("could not complete request", "error", err)
		return err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(reply); err != nil {
		logger.Error("could not decode response", "error", err)
	}

	// TODO: should save all of this into a ProjectResponse struct or something similar
	logger.Debug("sending reply", "projects", reply)

	return nil
}

func NewProjectService() *Project {
	// TODO: do any setup required
	return &Project{}
}
