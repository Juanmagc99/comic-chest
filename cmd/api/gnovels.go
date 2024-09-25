package main

import (
	"errors"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"juanmagc99.comic-chest/internal/data"
	"juanmagc99.comic-chest/internal/validator"
)

// Handler for GET /v1/gnovels
func (app *application) listGraphicNovelsHandler(w http.ResponseWriter, r *http.Request) {
	// Aquí iría la lógica para listar todas las graphic novels
	fmt.Fprintln(w, "Listing all graphic novels")
}

// Handler for POST /v1/gnovels
func (app *application) createGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		GNType      string   `json:"type"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Genres      []string `json:"genres"`
		Status      string   `json:"status"`
		Author      string   `json:"author"`
		Year        int32    `json:"year"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	gnovel := &data.Gnovel{
		GNType:      input.GNType,
		Title:       input.Title,
		Description: input.Description,
		Genres:      input.Genres,
		Status:      input.Status,
		Author:      input.Author,
		Year:        input.Year,
	}

	v := validator.New()

	if data.ValidateMovie(v, gnovel); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Gnovels.Insert(gnovel)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/gnovels/%d", gnovel.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"gnovel": gnovel}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Handler for GET /v1/gnovels/:id
func (app *application) getGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIntParam(r, "id")
	if err != nil {
		app.notFoundResponse(w, r)
	}

	gnovel, err := app.models.Gnovels.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	/*gnovel := data.Gnovel{
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
	}*/

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
