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

type Profile struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	RealName  string `json:"real_name"`
	Email     string `json:"email"`
	Skype     string `json:"skype"`
	Phone     string `json:"phone"`
	Image24   string `json:"image_24"`
	Image32   string `json:"image_32"`
	Image48   string `json:"image_48"`
	Image72   string `json:"image_72"`
	Image192  string `json:"image_192"`
}

// struct to encapsulate channel information
type Channel struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Created    int64          `json:"created"`
	Creator    string         `json:"creator"`
	Archived   bool           `json:"is_archived"`
	IsMember   bool           `json:"is_member"`
	NumMembers int64          `json:"num_members"`
	General    bool           `json:"is_general"`
	Topic      TopicOrPurpose `json:"topic"`
	Purpose    TopicOrPurpose `json:"purpose"`
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

// struct to encapsulate Groups
type Group struct {
	Id       string         `json:"id"`
	Name     string         `json:"name"`
	Created  int64          `json:"created"`
	Creator  string         `json:"creator"`
	Archived bool           `json:"is_archived"`
	Members  []string       `json:"members"`
	Topic    TopicOrPurpose `json:"topic"`
	Purpose  TopicOrPurpose `json:"purpose"`
}

type TopicOrPurpose struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet string `json:"last_set"`
}

type Starred struct {
	Type    string  `json:"type"`
	Channel string  `json:"channel"`
	Message Message `json:"message"`
	File    string  `json:"file"`
	Comment string  `json:"comment"`
}

type File struct {
	Id                 string    `json:"id"`
	TimeStamp          int64     `json:"timestamp"`
	Name               string    `json:"name"`
	Title              string    `json:"title"`
	MimeType           string    `json:"mimetype"`
	FileType           string    `json:"filetype"`
	Pretty             string    `json:"pretty_type"`
	User               string    `json:"user"`
	Mode               string    `json:"mode"`
	Editable           bool      `json:"editable"`
	IsExternal         bool      `json:"is_external"`
	ExternalType       string    `json:"external_type"`
	Size               int64     `json:"size"`
	Url                string    `json:"url"`
	UrlDownload        string    `json:"url_download"`
	UrlPrivate         string    `json:"url_private"`
	UrlPrivateDownload string    `json:"url_private_download"`
	Thumbnail64        string    `json:"thumb_64"`
	Thumbnail80        string    `json:"thumb_80"`
	Thumbnail360       string    `json:"thumb_360"`
	Thumbnail360Gif    string    `json:"thumb_360_gif"`
	Thumbnail360W      string    `json:"thumb_360_w"`
	Thumbnail360H      string    `json:"thumb_360_h"`
	Permalink          string    `json:"permalink"`
	EditLink           string    `json:"edit_link"`
	Preview            string    `json:"preview"`
	PreviewHighlight   string    `json:"preview_highlight"`
	Lines              int       `json:"lines"`
	LinesMore          int       `json:"lines_more"`
	IsPublic           bool      `json:"is_public"`
	PublicUrlShared    bool      `json:"public_url_shared"`
	Channels           []Channel `json:"channels"`
	Groups             []Group   `json:"groups"`
	InitialComment     string    `json:"initial_comment"`
	NumStars           int       `json:num_stars"`
	IsStarred          bool      `json:"is_starred"`
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
		Error   string `json:"error"`
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
	return nil, errors.New(r.Error)
}

// check the credentials
func AuthTest() (bool, error) {
	type returned struct {
		Ok    bool   `json="ok"`
		Error string `json:"error"`
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
		Ok       bool      `json:"ok"`
		Channels []Channel `json:"channels"`
		Error    string    `json:"error"`
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
	return r.Channels, errors.New(r.Error)
}

// return a list of the im history
func ChannelHistory(channel string, count int) ([]Message, error) {
	type respo struct {
		Ok       bool      `json:"ok"`
		Messages []Message `json:"messages"`
		Error    string    `json:"error"`
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
	return r.Messages, errors.New(r.Error)
}

func ChannelMark(channel string, timestamp string) (bool, error) {
	type respo struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}
	// modify the params
	vals := url.Values{}
	vals.Set("channel", channel)
	vals.Set("ts", timestamp)
	request, err := generateRequest("channels.mark", vals)
	if err != nil {
		return false, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return false, err
	}
	r := respo{}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return false, err
	}
	if r.Ok {
		return true, nil
	}
	return false, errors.New(r.Error)
}

// return a list of the im history
func IMHistory(channel string, count int) ([]Message, error) {
	type respo struct {
		Ok       bool      `json:"ok"`
		Messages []Message `json:"messages"`
		Error    string    `json:"error"`
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
	return r.Messages, errors.New(r.Error)
}

// list the IMs
func IMList() ([]IM, error) {
	type respo struct {
		Ok    bool   `json:"ok"`
		IMs   []IM   `json:"ims"`
		Error string `json:"error"`
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
	return r.IMs, errors.New(r.Error)
}

func ChatPostMessage(channel string, text string, botName string) (string, error) {
	type respo struct {
		Ok        bool  `json:"ok"`
		TimeStamp string `json:"timestamp"`
		Error     string `json:"error"`
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
	if r.Ok {
		return r.TimeStamp, nil
	}
	return r.TimeStamp, errors.New(r.Error)
}

func GroupList() ([]Group, error) {
	type respo struct {
		Ok     bool    `json:"ok"`
		Groups []Group `json:"groups"`
		Error  string  `json:"error"`
	}
	request, err := generateRequest("groups.list", nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	r := respo{Groups: make([]Group, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Groups, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.Ok {
		return r.Groups, nil
	}
	return r.Groups, errors.New(r.Error)
}

func GroupHistory(channel string, count int) ([]Message, error) {
	type respo struct {
		Ok       bool      `json:"ok"`
		Messages []Message `json:"messages"`
		Error    string    `json:"error"`
	}
	// modify the params
	vals := url.Values{}
	vals.Set("channel", channel)
	vals.Set("latest", "")
	vals.Set("oldest", "")
	vals.Set("count", strconv.Itoa(count))
	request, err := generateRequest("groups.history", vals)
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
	return r.Messages, errors.New(r.Error)
}

func StarsList(user string, count int) ([]Starred, error) {
	type respo struct {
		Ok      bool      `json:"ok"`
		Starred []Starred `json:"items"`
		Error   string    `json:"error"`
	}
	// modify the params
	vals := url.Values{}
	vals.Set("user", user)
	vals.Set("latest", "")
	vals.Set("oldest", "")
	vals.Set("count", strconv.Itoa(count))
	request, err := generateRequest("stars.list", vals)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	r := respo{Starred: make([]Starred, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Starred, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}
	if r.Ok {
		return r.Starred, nil
	}
	return r.Starred, errors.New(r.Error)
}

func FilesInfo(file string, count int) (File, error) {
	type respo struct {
		Ok   bool `json:"ok"`
		File File `json:"file"`
		//Comments []Comment `json:"comments"`
		Error string `json:"error"`
	}
	// modify the params
	vals := url.Values{}
	vals.Set("file", file)
	vals.Set("count", strconv.Itoa(count))
	request, err := generateRequest("files.info", vals)
	if err != nil {
		return File{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return File{}, err
	}
	r := respo{} //Comments: make([]Comment, 0)}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.File, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return File{}, err
	}
	if r.Ok {
		return r.File, nil
	}
	return r.File, errors.New(r.Error)
}

func FilesList() ([]File, error) {
	type respo struct {
		Ok    bool   `json:"ok"`
		Files []File `json:"files"`
		Error string `json:"error"`
	}
	// modify the params
	r := respo{Files: make([]File, 0)}
	request, err := generateRequest("files.list", nil)
	if err != nil {
		return r.Files, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return r.Files, err
	}
	defer response.Body.Close()
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r.Files, err
	}
	err = json.Unmarshal(b, &r)
	if err != nil {
		return r.Files, err
	}
	if r.Ok {
		return r.Files, nil
	}
	return r.Files, errors.New(r.Error)
}
