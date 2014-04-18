package slacker

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const rootPath = "https://slack.com/api/"

// api token
var token = ""

// struct to encapsulate the member details
type Member struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Deleted        bool   `json:"deleted"`
	Color          string `json:"color"`
	IsAdmin        bool   `json:"is_admin"`
	IsOwner        bool   `json:"is_owner"`
	IsPrimaryOwner bool   `json:"is_primary_owner"`
}

// struct to encapsulate channel information
type Channel struct {
	Name string `"json=name"`
}

// struct to encapsulate messages
type Message struct {
	Type      string `json:"type"`
	TimeStamp string `json:"ts"`
	User      string `json:"user"`
	Text      string `json:"text"`
	Starred   bool   `json:"is_starred"`
}

// struct to encapsulate IMs
type IM struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Created     int64  `json:"created"`
	UserDeleted bool   `json:"is_user_deleted"`
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
func generateRequest(method string, values url.Values) (*http.Request, error) {
	route := fmt.Sprintf("%s%s", rootPath, method)
	u, err := url.Parse(route)
	if err != nil {
		return nil, err
	}
	if len(values) > 0 {
		values.Set("token", token)
		u.RawQuery = values.Encode()
	} else {
		query := url.Values{}
		query.Set("token", token)
		u.RawQuery = query.Encode()
	}
	return http.NewRequest("GET", u.String(), nil)
}

// return all the users in the team
func UsersList() ([]Member, error) {
	request, err := generateRequest("users.list", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	type resp struct {
		Ok      bool
		Members []Member
	}
	r := resp{Members: make([]Member, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Members, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.Ok {
		return r.Members, nil
	}
	return nil, errors.New("Non ok value returned from API.")
}

// check the credentials
func AuthTest() (bool, error) {
	type returned struct {
		Ok bool `json="ok"`
	}
	request, err := generateRequest("auth.test", nil)
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
		Ok       bool      `"json:ok"`
		Channels []Channel `"json:channels"`
	}
	request, err := generateRequest("channels.list", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	r := respo{Channels: make([]Channel, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Channels, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.Ok {
		return r.Channels, nil
	}
	return r.Channels, errors.New("Non ok value receieved from API")
}

// return a list of the im history
func ChannelHistory(channel string, count int) ([]Message, error) {
	type respo struct {
		Ok       bool      `"json:ok"`
		Messages []Message `"json:messages"`
	}
	// modify the params
	vals := url.Values{}
	vals.Set("channel", channel)
	vals.Set("latest", "")
	vals.Set("oldest", "")
	vals.Set("count", strconv.Itoa(count))
	request, err := generateRequest("channels.history", vals)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	r := respo{Messages: make([]Message, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Messages, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.Ok {
		return r.Messages, nil
	}
	return r.Messages, errors.New("Non ok value receieved from API")
}

// return a list of the im history
func IMHistory(channel string, count int) ([]Message, error) {
	type respo struct {
		Ok       bool      `"json:ok"`
		Messages []Message `"json:messages"`
	}
	// modify the params
	vals := url.Values{}
	vals.Set("channel", channel)
	vals.Set("latest", "")
	vals.Set("oldest", "")
	vals.Set("count", strconv.Itoa(count))
	request, err := generateRequest("im.history", vals)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	r := respo{Messages: make([]Message, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Messages, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.Ok {
		return r.Messages, nil
	}
	return r.Messages, errors.New("Non ok value receieved from API")
}

// list the IMs
func IMList() ([]IM, error) {
	type respo struct {
		Ok  bool `"json:ok"`
		IMs []IM `"json:ims"`
	}
	request, err := generateRequest("im.list", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	r := respo{IMs: make([]IM, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.IMs, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.Ok {
		return r.IMs, nil
	}
	return r.IMs, errors.New("Non ok value receieved from API")
}

func ChatPostMessage(channel string, text string, botName string) (string, error) {
	type respo struct {
		Ok        int64  `"json:ok"`
		TimeStamp string `"json:timestamp"`
	}
	// modify the params
	r := respo{}
	vals := url.Values{}
	vals.Set("channel", channel)
	vals.Set("text", text)
	vals.Set("username", botName)
	request, err := generateRequest("chat.postMessage", vals)
	if err != nil {
		return r.TimeStamp, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return r.TimeStamp, err
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.TimeStamp, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return r.TimeStamp, err
	}
	if r.Ok == 1 {
		return r.TimeStamp, nil
	}
	return r.TimeStamp, errors.New("Non ok value receieved from API")
}
