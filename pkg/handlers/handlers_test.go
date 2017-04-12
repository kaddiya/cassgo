package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kaddiya/cassgo/pkg/config"
	"github.com/kaddiya/cassgo/pkg/services"
)

var routerTests = []struct {
	URLPath    string
	HTTPMethod string
}{{
	URLPath:    "/foo/",
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

var getFooTests = []struct {
	URLPath    string
	HTTPMethod string
	StatusCode int
	Message    string
}{{
	URLPath:    "/foo/",
	HTTPMethod: "GET",
	StatusCode: 200,
	Message:    "Please enter a valid email id",
}}

func TestGetFoo(t *testing.T) {
	a := &config.AppContainer{AppName: "Testing", FooServiceImpl: &services.FooServiceImpl{}}
	r := InitRouter(a)
	server := httptest.NewServer(r)
	for _, fixture := range getFooTests {
		request, _ := http.NewRequest(fixture.HTTPMethod, fixture.URLPath, &NoOpReader{})
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(ListFoo(a))
		handler.ServeHTTP(w, request)
		if w.Code != fixture.StatusCode {
			t.Fail()
		}
	}
	server.Close()

}
