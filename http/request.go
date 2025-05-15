package http

import (
	"encoding/json"
	netHttp "net/http"
	"strings"
)

type Request struct {
	HttpRequest *netHttp.Request
	Route       string
}

func (r *Request) Body(output any) {
	json.NewDecoder(r.HttpRequest.Body).Decode(&output)
}

func (r *Request) Param(key string) string {
	params := make(map[string]string)
	routeSplit := strings.Split(r.Route, "/")
	pathSplit := strings.Split(r.HttpRequest.URL.Path, "/")

	for index, value := range routeSplit {
		if strings.HasPrefix(value, ":") {
			params[value] = pathSplit[index]
		}
	}

	return params[":"+key]
}
