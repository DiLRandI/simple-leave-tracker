package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
)

type Config struct {
	Log     *slog.Logger
	Modules []RouteRegister
}

type (
	Path   string
	Method string
	Routes map[Method]map[Path]http.HandlerFunc
)

type RouteRegister interface {
	Name() string
	Register() Routes
}

type App struct {
	log     *slog.Logger
	modules []RouteRegister
}

func New(c *Config) *App {
	return &App{
		log:     c.Log,
		modules: c.Modules,
	}
}

func (a *App) RegisterRoutes(mux *http.ServeMux) error {
	routes := make(Routes)

	for _, m := range a.modules {
		for method, methods := range m.Register() {
			for path, handlerFn := range methods {
				_, routeOk := routes[method]
				if !routeOk {
					routes[method] = make(map[Path]http.HandlerFunc)
				}

				if !validateMethod(method) {
					return &ErrorInvalidMethod{
						method: method,
					}
				}

				_, methodOk := routes[method][path]
				if methodOk {
					return &ErrorRouteAlreadyRegistered{
						route:  path,
						method: method,
					}
				}

				routes[method][path] = handlerFn
				mux.Handle(fmt.Sprintf("%s %s", method, path), handlerFn)
				a.log.Debug("route registered", "method", method, "path", path)
			}
		}
	}

	return nil
}

func validateMethod(method Method) bool {
	validMethods := []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	}

	return slices.Contains(validMethods, string(method))
}
