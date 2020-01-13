package dwelo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const url = "https://api.dwelo.com/v3"
const tokenRefresh = time.Hour * 12

// stuff for the session function dont mess with this
type token struct {
	value     string
	retreived time.Time

	lock *sync.RWMutex
}

var userToken token

// getToken retreives a token using locks or the same one if the locks arent needed,
// this will panic if a token is getting grabbed but an error occurs
func getToken() string {
	userToken.lock.RLock()
	if time.Since(userToken.retreived) > tokenRefresh {
		userToken.lock.RUnlock()

		log.Info("[token] Getting new token...")

		// refresh the token
		req := NewLoginRequest(os.Getenv("DWELO_EMAIL"), os.Getenv("DWELO_PW"))
		resp, err := req.login()
		if err != nil {
			log.Errorf("error logging in, going to panic: %v", err)
			panic(err)
		}

		// set the new token using write locks
		userToken.lock.Lock()
		userToken.retreived = time.Now()
		userToken.value = resp.Token
		userToken.lock.Unlock()
		return resp.Token
	}
	// get the token that already exists
	t := userToken.value
	userToken.lock.RUnlock()
	log.Info("[token] Using existing token...")

	return t
}

// login logs a user in and returns a true inside of the LoginResponse that gets returned
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

// Do is basically the command structure to send a command
func Do(w http.ResponseWriter, r *http.Request) {
	reqID := uuid.New()
	if r.Method != http.MethodPost {
		log.Warnf("got a bad method for request: %v", reqID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userToken.lock.RLock()
	if time.Since(userToken.retreived) > tokenRefresh {
		// refresh the token

	}
}
