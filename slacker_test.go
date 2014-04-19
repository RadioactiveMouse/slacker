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
	token = ""
	notok, _ := AuthTest()
	ex.Pect(t, notok, false)
	LoadToken()
	ok, _ := AuthTest()
	ex.Pect(t, ok, true)
}

func TestGetUsers(t *testing.T) {
	LoadToken()
	_, err := UsersList()
	if err != nil {
		t.Error(err)
	}
}

func TestChannelsList(t *testing.T) {
	LoadToken()
	_, err := ChannelsList()
	if err != nil {
		t.Error(err)
	}
}

func TestChannelHistory(t *testing.T) {
	LoadToken()
	_, err := ChannelHistory("C028YLM2M", 10)
	if err != nil {
		t.Error(err)
	}
}

func TestChannelMark(t *testing.T) {
	LoadToken()
	_, err := ChannelMark("C028YLM2M", "")
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
	_, err := IMList()
	if err != nil {
		t.Error(err)
	}
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

func TestFileInfo(t *testing.T) {
	LoadToken()
	_, err := FilesInfo("demo", 1)
	if err != nil {
		ex.Pect(t, fmt.Sprintf("%s", err), "file_not_found")
	}
}

func TestFilesList(t *testing.T) {
	LoadToken()
	_, err := FilesList()
	if err != nil {
		t.Error(err)
	}
}
