package slacker

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/RadioactiveMouse/ex"
)

func TestGenerateRequestNoValues(t *testing.T) {
	LoadToken()
	request, _ := generateRequest("users.list", nil)
	ex.Pect(t, request.URL.Path, "/api/users.list")
	ex.Pect(t, request.URL.RawQuery, fmt.Sprintf("token=%s", token))
}

func TestGenerateRequestWithValues(t *testing.T) {
	LoadToken()
	vals := url.Values{}
	vals.Set("testing", "test")
	request, _ := generateRequest("im.history", vals)
	ex.Pect(t, request.URL.Path, "/api/im.history")
	ex.Pect(t, request.URL.RawQuery, fmt.Sprintf("testing=test&token=%s", token))
}

func TestAuthTest(t *testing.T) {
	LoadToken()
	ok, _ := AuthTest()
	ex.Pect(t, ok, true)
}

func TestGetUsers(t *testing.T) {
	LoadToken()
	members, _ := UsersList()
	ex.Pect(t, len(members), 2)
}

func TestChannelsList(t *testing.T) {
	LoadToken()
	c, _ := ChannelsList()
	ex.Pect(t, len(c), 3)
}

func TestChannelHistory(t *testing.T) {
	LoadToken()
	_, err := ChannelHistory("C028YLM2M", 10)
	if err != nil {
		t.Error(err)
	}
}

func TestIMHistory(t *testing.T) {
	LoadToken()
	_, err := IMHistory("D028YM1S7", 10)
	if err != nil {
		t.Error(err)
	}
}

func TestIMList(t *testing.T) {
	LoadToken()
	ims, err := IMList()
	if err != nil {
		t.Error(err)
	}
	ex.Pect(t, len(ims), 2)
}

func TestChatPostMessage(t *testing.T) {
	LoadToken()
	_, err := ChatPostMessage("C028YLM2M", "Hi from Slacker", "demobot")
	if err != nil {
		t.Error(err)
	}
}

func TestGroupList(t *testing.T) {
	LoadToken()
	_, err := GroupList()
	if err != nil {
		t.Error(err)
	}
}

func TestGroupHistory(t *testing.T) {
	LoadToken()
	_, err := GroupHistory("G028ZJX59", 10)
	if err != nil {
		t.Error(err)
	}
}

func TestStars(t *testing.T) {
	LoadToken()
	_, err := Stars("", 10)
	if err != nil {
		t.Error(err)
	}
}
