package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/gnovels", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/gnovels", app.createGraphicNovelHandler)
	router.HandlerFunc(http.MethodGet, "/v1/gnovels/:id", app.getGraphicNovelHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/gnovels/:id", app.updateGraphicNovelHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/gnovels/:id", app.deleteGraphicNovelHandler)

	router.HandlerFunc(http.MethodGet, "/v1/gnovels/:id/chapter/:nchapter", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPut, "/v1/gnovels/:id/chapter/:nchapter", app.healthcheckHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/gnovels/:id/chapter/:nchapter", app.healthcheckHandler)

	// Return the httprouter instance.
	return app.recoverPanic(router)
}
