package main

import (
	"fmt"
	"net/http"
)

// Handler for GET /v1/gnovels/:id/chapter/:nchapter
func (app *application) getGraphicNovelChapterHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para obtener un capítulo de una graphic novel por su número de capítulo
	id := r.URL.Query().Get(":id")
	nchapter := r.URL.Query().Get(":nchapter")
	fmt.Fprintf(w, "Getting chapter %s of graphic novel with ID: %s\n", nchapter, id)
}

// Handler for GET /v1/gnovels/:id/chapter/:nchapter
func (app *application) createGraphicNovelChapterHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para obtener un capítulo de una graphic novel por su número de capítulo
	id := r.URL.Query().Get(":id")
	nchapter := r.URL.Query().Get(":nchapter")
	fmt.Fprintf(w, "Getting chapter %s of graphic novel with ID: %s\n", nchapter, id)
}

// Handler for PUT /v1/gnovels/:id/chapter/:nchapter
func (app *application) updateGraphicNovelChapterHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para actualizar un capítulo de una graphic novel
	id := r.URL.Query().Get(":id")
	nchapter := r.URL.Query().Get(":nchapter")
	fmt.Fprintf(w, "Updating chapter %s of graphic novel with ID: %s\n", nchapter, id)
}

// Handler for DELETE /v1/gnovels/:id/chapter/:nchapter
func (app *application) deleteGraphicNovelChapterHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para eliminar un capítulo de una graphic novel
	id := r.URL.Query().Get(":id")
	nchapter := r.URL.Query().Get(":nchapter")
	fmt.Fprintf(w, "Deleting chapter %s of graphic novel with ID: %s\n", nchapter, id)
}
