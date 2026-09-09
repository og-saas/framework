package jwtx

import (
	"context"
	"net/http"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

var j = NewJWT().
	WithScene("api").
	WithSecret("123456").
	WithSso(true).
	WithTTL(500)

func TestGenToken(t *testing.T) {

	token, err := j.GenerateToken(context.Background(), 1, jwt.MapClaims{
		"name": "1",
		"age":  1,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZ2UiOjEsImV4cCI6MTc4ODkyMTU0MSwiaWF0IjoxNzg4OTIxMDQxLCJuYW1lIjoiMSIsInVzZXJfaWQiOjF9.RNwQS9SJ8-wDwj8LS1Y_xR5FRVDL6MXPiHa9DAKNdz4"

	_, data, err := j.ParseToken(&http.Request{
		Header: http.Header{"Authorization": []string{"Bearer " + token}},
	})
	t.Log(err)
	t.Log(cast.ToInt64(data["exp"]))
}
