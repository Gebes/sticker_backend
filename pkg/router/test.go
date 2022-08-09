package router

import (
	"fmt"
	. "github.com/Gebes/there/v2"
	"io/ioutil"
	"net/http"
)

func TestIpServerGet(request HttpRequest) HttpResponse {
	resp, err := http.Get("https://api.ipify.org/")
	if err != nil {
		return Error(StatusInternalServerError, fmt.Errorf("make http request: %v", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Error(StatusInternalServerError, fmt.Errorf("read body: %v", err))
	}
	return Bytes(StatusOK, body)
}
