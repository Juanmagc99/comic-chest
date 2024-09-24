package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"juanmagc99.comic-chest/internal/data"
)

// Handler for GET /v1/gnovels
func (app *application) listGraphicNovelsHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para listar todas las graphic novels
	fmt.Fprintln(w, "Listing all graphic novels")
}

// Handler for POST /v1/gnovels
func (app *application) createGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para crear una nueva graphic novel
	fmt.Fprintln(w, "Creating a new graphic novel")
}

// Handler for GET /v1/gnovels/:id
func (app *application) getGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIntParam(r, "id")
	if err != nil {
		app.notFoundResponse(w, r)
	}

	gnovel := data.Gnovel{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Black Lagoon",
		Description: `The series follows the Lagoon Company, a four-member team of pirate mercenaries smuggling
		 goods in and around the seas of Southeast Asia with their PT boat, the Black Lagoon.The group 
		 takes on various jobs, usually involving criminal organizations, and resulting in violent gunfights.`,
		Genres:   []string{"Action", "Drama"},
		Author:   "Rei Hiroe",
		Status:   "ongoing",
		NChapers: 78,
		Year:     2002,
		GNType:   "Manga",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"gnovel": gnovel}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Handler for PUT /v1/gnovels/:id
func (app *application) updateGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para actualizar una graphic novel por su ID
	id := r.URL.Query().Get(":id")
	fmt.Fprintf(w, "Updating graphic novel with ID: %s\n", id)
}

// Handler for DELETE /v1/gnovels/:id
func (app *application) deleteGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para eliminar una graphic novel por su ID
	id := r.URL.Query().Get(":id")
	fmt.Fprintf(w, "Deleting graphic novel with ID: %s\n", id)
}

// Handler for GET /v1/gnovels/:id/chapter/:nchapter
func (app *application) getGraphicNovelChapterHandler(w http.ResponseWriter, r *http.Request) {
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
