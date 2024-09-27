package app_test

import (
	"io"
	"log/slog"
	"net/http"
	"testing"

	"simple-leave-tracker/internal/app"

	"github.com/stretchr/testify/assert"
)

func TestRegisterRoutes(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))

	for _, tc := range getTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			sut := app.New(&app.Config{
				Log:     logger,
				Modules: tc.modules,
			})

			mux := http.NewServeMux()

			err := sut.RegisterRoutes(mux)
			if tc.expError != "" {
				assert.EqualError(t, err, tc.expError)
			}
		})
	}
}

type TT struct {
	name     string
	modules  []app.RouteRegister
	expError string
}

func getTestCases() []TT {
	return []TT{
		{
			name: "invalid route method",
			modules: []app.RouteRegister{
				&mockModule{
					routes: app.Routes{
						"invalid_method": map[app.Path]http.HandlerFunc{
							"/abc": http.NotFound,
						},
					},
				},
			},
			expError: "The method [invalid_method] is not supported",
		},
		{
			name: "try to register same route more than once",
			modules: []app.RouteRegister{
				&mockModule{
					routes: app.Routes{
						"POST": map[app.Path]http.HandlerFunc{
							"/abc": http.NotFound,
						},
					},
				},
				&mockModule{
					routes: app.Routes{
						"POST": map[app.Path]http.HandlerFunc{
							"/abc": http.NotFound,
						},
					},
				},
			},
			expError: "The route [POST /abc] is already registered",
		},
	}
}

type mockModule struct {
	routes app.Routes
}

func (*mockModule) Name() string {
	return "test_module"
}

func (m *mockModule) Register() app.Routes {
	return m.routes
}
