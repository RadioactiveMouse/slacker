package slacker

import (
	"os"
	"net/http"
	"errors"
	"fmt"
	"net/url"
	"log"
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
	log.Println(response.Body)
	return nil, nil
}
