package routing

import (
	"encoding/json"
	netHttp "net/http"
	"strings"

	"github.com/S3ergio31/url-shortener/http"
)

type Controller func(r *http.Request) http.Response

type routeMap struct {
	path       string
	controller Controller
}

var routes = make(map[string][]routeMap)

func Handler(w netHttp.ResponseWriter, r *netHttp.Request) {
	routesByMethod := routes[r.Method]
	w.Header().Set("Content-Type", "application/json")
	found := false

	for _, route := range routesByMethod {
		if hasMatch(r.URL.Path, route.path) {
			found = true
			response := route.controller(&http.Request{HttpRequest: r, Route: route.path})
			w.WriteHeader(response.Status)
			json.NewEncoder(w).Encode(response.Data)
			break // break on first match
		}
	}

	if found {
		return
	}

	responseNotFound := http.ResponseNotFound()
	w.WriteHeader(responseNotFound.Status)
	json.NewEncoder(w).Encode(responseNotFound.Data)
}

func hasMatch(path string, route string) bool {
	pathSplit := strings.Split(path, "/")
	routeSplit := strings.Split(route, "/")

	if len(pathSplit) != len(routeSplit) {
		return false
	}

	matchSplit := make([]string, len(pathSplit))

	for key, routeSection := range routeSplit {
		matchSplit[key] = routeSection

		if strings.HasPrefix(routeSection, ":") {
			matchSplit[key] = pathSplit[key]
		}
	}

	joinedMatch := strings.Join(matchSplit, "/")
	joinedPath := strings.Join(pathSplit, "/")

	return joinedMatch == joinedPath
}

func add(route string, method string, controller Controller) {
	routesByMethod := routes[method]
	routes[method] = append(routesByMethod, routeMap{route, controller})
}

func Get(route string, controller Controller) {
	add(route, "GET", controller)
}

func Post(route string, controller Controller) {
	add(route, "POST", controller)
}

func Put(route string, controller Controller) {
	add(route, "PUT", controller)
}

func Delete(route string, controller Controller) {
	add(route, "DELETE", controller)
}
