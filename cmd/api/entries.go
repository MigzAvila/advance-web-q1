// Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

// createEntryHandler for POST /v1/entries endpoints
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating a new entry...")

}

// createEntryHandler for GET /v1/entries endpoints
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	//Utilize Utility Methods From helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Displaying out the Entry ID
	fmt.Fprintf(w, "show the details for the Entry %d\n", id)

}
