package auth

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const url = "https://authserver.mojang.com"

func Authenticate(email string, password string) string {
	text := "{\"agent\":{\"name\":\"Minecraft\",\"version\": 1},\"username\": \"" + email + "\",\"password\": \"" + password + "\"}"
	resp := request(text, "/authenticate")
	return resp
}

func Validate(accessToken string) (success bool) {
	text := "{\"accessToken\": \"" + accessToken + "\"}"
	resp := request(text, "/validate")
	if len(resp) > 0 {
		return false
	} else {
		return true
	}
}

func SignOut(username string, password string) (success bool) {
	text := `{"username":"` + username + `","password":"` + password + `"}`
	resp := request(text, "/signout")
	if len(resp) > 0 {
		return false
	} else {
		return true
	}
}

func Invalidate(accessToken string, clientToken string) (success bool) {
	text := `"accessToken":"` + accessToken + `","clientToken":"` + clientToken + `"}`
	resp := request(text, "/invalidate")
	if len(resp) > 0 {
		return false
	} else {
		return true
	}
}

func Refresh(accessToken string, clientToken string) string {
	text := `{"accessToken":"` + accessToken + `","clientToken":"` + clientToken + `"}`
	resp := request(text, "/refresh")
	return resp
}

func request(text string, endpoint string) string {
	buff := bytes.NewBufferString(text)
	resp, err := http.Post(url+endpoint, "application/json", buff)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		return str
	} else {
		panic(err)
	}
}
