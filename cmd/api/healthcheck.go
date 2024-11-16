package main

import (
	"net/http"
)

<<<<<<< HEAD
func (a *applicationDependencies) healthcheckHandler(w http.ResponseWriter,
	r *http.Request) {
	//panic("Apples & Oranges")
=======
func (a *applicationDependences) healthChechHandler(w http.ResponseWriter, r *http.Request) {

>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": a.config.environment,
			"version":     appVersion,
		},
	}
<<<<<<< HEAD
	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)

=======

	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
	}
}
