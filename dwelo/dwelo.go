package dwelo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const url = "https://api.dwelo.com/v3"

func (req LoginRequest) login() (*LoginResponse, error) {
	log.Infof("starting new login for %v", req.id)

	// create the body of the request from LoginRequest
	bs, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(bs)

	newRequest, err := http.NewRequest(http.MethodPost, url+"/login/", reader)
	if err != nil {
		log.Errorf("error creating new request: %v", err)
		return nil, err
	}
	newRequest.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		log.Errorf("error sending request: %v", err)
		return nil, err
	}

	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("error reading response body: %v: %v", resp, err)
		return nil, err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		// the user is not logged in so ignore everything but logged in
		log.Infof("%v was not logged in from request: %v", req.Email, req.id)
		return &LoginResponse{loggedIn: false}, nil
	}

	// if the status is created that means that the user was logged in
	if resp.StatusCode != http.StatusCreated {
		log.Warnf("got an unexpected status code for %v (%v)", req.id, resp.StatusCode)
		return &LoginResponse{loggedIn: false}, nil
	}

	var lr LoginResponse
	err = json.Unmarshal(rBody, &lr)
	if err != nil {
		log.Errorf("Erroro unmarshalling body %v: %v ", string(rBody), err)
		return nil, err
	}
	log.Infof("got response for request: %v", err)
	lr.loggedIn = true

	return &lr, nil
}
