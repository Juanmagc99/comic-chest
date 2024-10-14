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

	router.HandlerFunc(http.MethodGet, "/v1/gnovels", app.requirePermission("gnovels:read", app.listGraphicNovelsHandler))

	router.HandlerFunc(http.MethodPost, "/v1/gnovels", app.requirePermission("gnovels:read", app.createGraphicNovelHandler))
	router.HandlerFunc(http.MethodGet, "/v1/gnovels/:id", app.requirePermission("gnovels:read", app.getGraphicNovelHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/gnovels/:id", app.requirePermission("gnovels:read", app.updateGraphicNovelHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/gnovels/:id", app.requirePermission("gnovels:read", app.deleteGraphicNovelHandler))

	router.HandlerFunc(http.MethodGet, "/v1/gnovels/:id/chapter/:number", app.requirePermission("gnovels:read", app.getChapterHandler))
	router.HandlerFunc(http.MethodPost, "/v1/gnovels/:id/chapter/:number", app.requirePermission("gnovels:read", app.createChapterHandler))
	router.HandlerFunc(http.MethodPut, "/v1/gnovels/:id/chapter/:number", app.requirePermission("gnovels:read", app.updateChapterHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/gnovels/:id/chapter/:number", app.requirePermission("gnovels:read", app.deleteChapterHandler))

	router.HandlerFunc(http.MethodGet, "/v1/serve/:id/chapter/:number", app.requirePermission("gnovels:read", app.serveChapterHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.createUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	// Return the httprouter instance.
	return app.recoverPanic(app.authenticate(router))
}
