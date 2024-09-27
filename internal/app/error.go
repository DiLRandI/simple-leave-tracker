package app

import "fmt"

type ErrorRouteAlreadyRegistered struct {
	route  Path
	method Method
}

func (e *ErrorRouteAlreadyRegistered) Error() string {
	return fmt.Sprintf("The route [%s %s] is already registered", e.method, e.route)
}

type ErrorInvalidMethod struct {
	method Method
}

func (e *ErrorInvalidMethod) Error() string {
	return fmt.Sprintf("The method [%s] is not supported", e.method)
}
