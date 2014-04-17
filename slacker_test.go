package slacker

import (
	"testing"
	"fmt"

	"github.com/RadioactiveMouse/ex"
)

func TestGenerateRequest(t *testing.T) {
	LoadToken()
	request, _ := generateRequest("users.list")
	ex.Pect(t, request.URL.Path, "/api/users.list")
	ex.Pect(t, request.URL.RawQuery, fmt.Sprintf("token=%s",token))
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
