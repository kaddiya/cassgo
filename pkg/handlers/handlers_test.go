package handlers

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kaddiya/cassgo/pkg/config"
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
