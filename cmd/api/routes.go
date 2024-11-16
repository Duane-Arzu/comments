<<<<<<< HEAD
// Filename: cmd/api/routes.go
=======
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

<<<<<<< HEAD
func (a *applicationDependencies) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(a.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/comments", a.requireActivatedUser(a.listCommentsHandler))

	router.HandlerFunc(http.MethodPost, "/v1/comments", a.requireActivatedUser(a.createCommentHandler))

	router.HandlerFunc(http.MethodGet, "/v1/comments/:id", a.requireActivatedUser(a.displayCommentHandler))

	router.HandlerFunc(http.MethodPatch, "/v1/comments/:id", a.requireActivatedUser(a.updateCommentHandler))

	router.HandlerFunc(http.MethodDelete, "/v1/comments/:id", a.requireActivatedUser(a.deleteCommentHandler))

	router.HandlerFunc(http.MethodPut, "/v1/users/activated", a.activateUserHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", a.createAuthenticationTokenHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", a.registerUserHandler)

	return a.recoverPanic(a.rateLimit(a.authenticate(router)))

=======
func (a *applicationDependences) routes() http.Handler {
	//setup a new router
	router := httprouter.New()

	//handle 405
	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedResponse)

	//method 404
	router.NotFound = http.HandlerFunc(a.notFoundResponse)

	//setup routes
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthChechHandler)
	router.HandlerFunc(http.MethodPost, "/v1/comments", a.createCommentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/comments/:id", a.displayCommentHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/comments/:id", a.updateCommentHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/comments/:id", a.deleteCommentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/comments", a.listCommentHandler)
	return a.recoverPanic(router)
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
}
