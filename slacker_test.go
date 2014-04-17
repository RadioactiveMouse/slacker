package slacker

import (
	"testing"
	"fmt"

	"github.com/RadioactiveMouse/ex"
)

func TestGenerateRequest(t *testing.T) {
	_ = LoadToken()
	request, _ := generateRequest("users.list")
	ex.Pect(t, request.URL.Path, "/api/users.list")
	ex.Pect(t, request.URL.RawQuery, fmt.Sprintf("token=%s",token))
}
