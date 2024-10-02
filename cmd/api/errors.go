//Filename: cmd/api/errors.go

package main

import (
	"fmt"
	"net/http"
)

func (a *applicationDependencies)logError(r *http.Request, err error) {

	method := r.Method
	uri := r.URL.RequestURI()
	a.logger.Error(err.Error(), "method", method, "uri", uri)

}

// send an error response in JSON
func (a *applicationDependencies)errorResponseJSON(w http.ResponseWriter, r *http.Request, status int, message any) {

errorData := envelope{"error": message}
err := a.writeJSON(w, status, errorData, nil)
if err != nil {
	a.logError(r, err)
	w.WriteHeader(500)
}
}


// send an error message if our server messes up
func (a *applicationDependencies)serveErrorResponse(w http.ResponseWriter, r *http.Request, err error) {

// First thing is to log error message
	a.logError(r, err)
//prepare a response to send to the client
message := "the server encountered a problem and could not process your request"

a.errorResponseJSON(w, r, http.StatusInternalServerError, message)
}

func (a *applicationDependencies)notFoundResponse(w http.ResponseWriter, r *http.Request) {

// we only log server errors, not client erros
// prepare a response to send to the client

message := "the requested resource could not be found"
a.errorResponseJSON(w, r, http.StatusNotFound, message)
}

func (a *applicationDependencies)methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {

// we only log server errors, not client erros
// prepare a formatted response to send to the client

message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)

a.errorResponseJSON(w, r, http.StatusMethodNotAllowed, message)


}

