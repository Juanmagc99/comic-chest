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

	if data.ValidateGnovel(v, gnovel); !v.Valid() {
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

	err = app.writeJSON(w, http.StatusOK, envelope{"gnovel": gnovel}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Handler for PATCH /v1/gnovels/:id
func (app *application) updateGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIntParam(r, "id")
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	gnovel, err := app.models.Gnovels.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
	}

	var input struct {
		GNType      *string  `json:"type"`
		Title       *string  `json:"title"`
		Description *string  `json:"description"`
		Genres      []string `json:"genres"`
		Status      *string  `json:"status"`
		Author      *string  `json:"author"`
		Year        *int32   `json:"year"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.GNType != nil {
		gnovel.GNType = *input.GNType
	}

	if input.Title != nil {
		gnovel.Title = *input.Title
	}

	if input.Description != nil {
		gnovel.Description = *input.Description
	}

	if input.Genres != nil {
		gnovel.Genres = input.Genres
	}

	if input.Status != nil {
		gnovel.Status = *input.Status
	}

	if input.Author != nil {
		gnovel.Author = *input.Author
	}

	if input.Year != nil {
		gnovel.Year = *input.Year
	}

	v := validator.New()
	if data.ValidateGnovel(v, gnovel); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
	}

	err = app.models.Gnovels.Update(gnovel)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"gnovel": gnovel}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Handler for DELETE /v1/gnovels/:id
func (app *application) deleteGraphicNovelHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIntParam(r, "id")
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Gnovels.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "gnovel succesfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
