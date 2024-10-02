package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"juanmagc99.comic-chest/internal/data"
	"juanmagc99.comic-chest/internal/validator"
)

// Handler for GET /v1/gnovels/:id/chapter/:number
func (app *application) getChapterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIntParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	number, err := app.readIntParam(r, "number")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	chapter, err := app.models.Chapters.Get(id, int(number))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	fmt.Println(chapter.FilePath)
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Itoa(chapter.Number)+".cbz")
	w.Header().Set("Content-Type", "application/zip")
	http.ServeFile(w, r, chapter.FilePath)
}

// Handler for POST /v1/gnovels/:id/chapter/:number
func (app *application) createChapterHandler(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 30 * 1024 * 1024 // 30MB

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

	number, err := app.readIntParam(r, "number")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Recoger el archivo subido (campo "file")
	file, fileheader, err := r.FormFile("file")
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("unable to obtain archive: %v", err))
		return
	}

	defer file.Close()

	filename := fileheader.Filename
	if !strings.HasSuffix(filename, ".cbz") {
		app.badRequestResponse(w, r, fmt.Errorf("only archives with .cbz are allowed"))
	}

	chapter := &data.Chapter{
		GnovelID: id,
		Number:   int(number),
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

// Handler for PUT /v1/gnovels/:id/chapter/:number
func (app *application) updateChapterHandler(w http.ResponseWriter, r *http.Request) {
	const maxUploadSize = 30 * 1024 * 1024 // 30MB

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

	number, err := app.readIntParam(r, "number")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Recoger el archivo subido (campo "file")
	file, fileheader, err := r.FormFile("file")
	if err != nil {
		app.badRequestResponse(w, r, fmt.Errorf("unable to obtain archive: %v", err))
		return
	}

	defer file.Close()

	filename := fileheader.Filename
	if !strings.HasSuffix(filename, ".cbz") {
		app.badRequestResponse(w, r, fmt.Errorf("only archives with .cbz are allowed"))
	}

	chapter := &data.Chapter{
		GnovelID: id,
		Number:   int(number),
	}

	v := validator.New()
	if data.ValidateChapter(v, chapter, app.models.Gnovels); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.deleteFile(id, int(number))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.Chapters.Delete(id, int(number))
	if err != nil {
		app.serverErrorResponse(w, r, err)
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

	err = app.writeJSON(w, http.StatusOK, envelope{"chapter": chapter}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// Handler for DELETE /v1/gnovels/:id/chapter/:number
func (app *application) deleteChapterHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIntParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	number, err := app.readIntParam(r, "number")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = app.deleteFile(id, int(number))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.Chapters.Delete(id, int(number))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "chapter deleted succesfully"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
