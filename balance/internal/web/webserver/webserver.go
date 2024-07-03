package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(webServerPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: webServerPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	key := method + ":" + path
	s.Handlers[key] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for key, handler := range s.Handlers {
		method := key[:strings.Index(key, ":")]
		path := key[strings.Index(key, ":")+1:]
		s.Router.MethodFunc(method, path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
