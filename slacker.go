package slacker

import (
	"os"
	"net/http"
	"errors"
	"fmt"
	"net/url"
	"encoding/json"
	"io/ioutil"
)

const rootPath = "https://slack.com/api/"
// api token
var token = ""

// struct to encapsulate the member details
type Member struct {
	Id	string	`json:"id"`
	Name	string	`json:"name"`
	Deleted	bool	`json:"deleted"`
	Color	string	`json:"color"`
	IsAdmin	bool	`json:"is_admin"`
	IsOwner	bool	`json:"is_owner"`
	IsPrimaryOwner	bool	`json:"is_primary_owner"`
}

// struct to encapsulate channel information
type Channel struct {
	Name	string	`"json=name"`
}

// load a token from environment variables
func LoadToken() error {
	token = os.Getenv("SLACK_API_TOKEN")
	if len(token) == 0 {
		return errors.New("Could not find SLACK_API_TOKEN")
	}
	return nil
}

// generate request for query
func generateRequest(method string) (*http.Request, error) {
	route := fmt.Sprintf("%s%s", rootPath, method)
	u, err := url.Parse(route)
	if err != nil {
		return nil, err
	}
	query := url.Values{}
	query.Set("token", token)
	u.RawQuery = query.Encode()
	return http.NewRequest("GET", u.String(), nil)
}

// return all the users in the team
func GetUsers() ([]Member, error) {
	request, err := generateRequest("users.list")
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	type resp struct {
		Ok	bool
		Members	[]Member
	}
	r := resp{Members: make([]Member,0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Members, err
	}
	err = json.Unmarshal(b, &r)
	return r.Members, err
}

// check the credentials
func AuthTest() (bool, error) {
	type returned struct {
		Ok	bool	`json="ok"`
		
	}
	request, err := generateRequest("auth.test")
	if err != nil {
		return false, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return false, err
	}
	r := returned{}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(b, &r)
	return r.Ok, nil
}

// return a list the channels
func ChannelsList() ([]Channel, error) {
	type respo struct {
		Ok	bool	`"json:ok"`
		Channels	[]Channel	`"json:channels"`
	}
	request, err := generateRequest("channels.list")
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	r := respo{Channels: make([]Channel,0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Channels, err
	}
	err = json.Unmarshal(b, &r)
	return r.Channels, err
}
