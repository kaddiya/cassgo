package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kaddiya/cassgo/pkg/config"
)

var routerTests = []struct {
	URLPath    string
	HTTPMethod string
}{{
	URLPath:    "/Todo/",
	HTTPMethod: "GET",
}}

type NoOpReader struct {
}

func (r *NoOpReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

// TestRoutes tests the HTTP routes
func TestRoutes(t *testing.T) {
	a := &config.AppContainer{AppName: "Testing"}
	r := InitRouter(a)

	for _, fixture := range routerTests {
		request, _ := http.NewRequest(fixture.HTTPMethod, fixture.URLPath, &NoOpReader{})
		var match mux.RouteMatch
		b := r.Match(request, &match)
		if !b {
			t.Logf("could not find the route %s", fixture.URLPath)
			t.Fail()
		}
	}

}

var getTodoTests = []struct {
	URLPath    string
	HTTPMethod string
	StatusCode int
	Message    string
}{{
	URLPath:    "/Todo/",
	HTTPMethod: "GET",
	StatusCode: 200,
	Message:    "Please enter a valid email id",
}}

func TestGetTodo(t *testing.T) {
	a := &config.AppContainer{AppName: "Testing"}
	r := InitRouter(a)
	server := httptest.NewServer(r)
	for _, fixture := range getTodoTests {
		request, _ := http.NewRequest(fixture.HTTPMethod, fixture.URLPath, &NoOpReader{})
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(ListTodo(a))
		handler.ServeHTTP(w, request)
		if w.Code != fixture.StatusCode {
			t.Fail()
		}
	}
	server.Close()

}
