package main

import (
	"fmt"
	"net/http"

	"juanmagc99.comic-chest/internal/data"
	"juanmagc99.comic-chest/internal/validator"
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
	const maxUploadSize = 100 * 1024 * 1024 // 50MB

	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("the file is too heavy"))
		return
	}

	id, err := app.readIntParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	nchapter, err := app.readIntParam(r, "nchapter")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Recoger el archivo subido (campo "file")
	file, _, err := r.FormFile("file")
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("unable to obtain archive: %v", err))
		return
	}
	defer file.Close()

	chapter := &data.Chapter{
		GnovelID: id,
		Number:   int(nchapter),
	}

	v := validator.New()
	if data.ValidateChapter(v, chapter, app.models.Gnovels); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	filepath, err := app.uploadChapter(chapter.GnovelID, chapter.Number, file)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}

	chapter.FilePath = filepath

	err = app.models.Chapters.Insert(chapter)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/chapters/%d", chapter.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"chapter": chapter}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

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
