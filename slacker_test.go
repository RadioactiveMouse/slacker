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
	members, _ := GetUsers()
	ex.Pect(t, len(members), 2)
}

func TestListChannel(t *testing.T) {
	LoadToken()
	c, _ := ChannelsList()
	ex.Pect(t, len(c), 3)
}

func TestIMHistory(t *testing.T) {
	LoadToken()
	m, _ := IMHistory("jeff", 10)
	ex.Pect(t, len(m), 0)
}
