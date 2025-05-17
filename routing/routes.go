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

func Handler(responseWriter netHttp.ResponseWriter, request *netHttp.Request) {
	setJsonHeader(responseWriter)

	for _, route := range routes[request.Method] {
		if hasMatch(request.URL.Path, route.path) {
			handleMatch(route, request, responseWriter)
			return // break on first match
		}
	}

	writeResponse(responseWriter, http.ResponseNotFound())
}

func setJsonHeader(responseWriter netHttp.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
}

func handleMatch(route routeMap, request *netHttp.Request, responseWriter netHttp.ResponseWriter) {
	response := route.controller(&http.Request{HttpRequest: request, Route: route.path})
	writeResponse(responseWriter, response)
}

func writeResponse(responseWriter netHttp.ResponseWriter, response http.Response) {
	responseWriter.WriteHeader(response.Status)
	json.NewEncoder(responseWriter).Encode(response.Data)
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
