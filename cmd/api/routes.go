package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	//router.NotFound = http.HandlerFunc(app.notFoundResponse)

	//router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	// Register the relevant methods, URL patterns and handler functions for our
	// endpoints using the HandlerFunc() method. Note that http.MethodGet and
	// http.MethodPost are constants which equate to the strings "GET" and "POST"
	// respectively.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/gnovels", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/gnovels", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/gnovels/:id", app.getGraphicNovelHandler)
	router.HandlerFunc(http.MethodPut, "/v1/gnovels/:id", app.healthcheckHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/gnovels/:id", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/gnovels/:id/chapter/:nchapter", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPut, "/v1/gnovels/:id/chapter/:nchapter", app.healthcheckHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/gnovels/:id/chapter/:nchapter", app.healthcheckHandler)

	// Return the httprouter instance.
	return app.recoverPanic(router)
}
