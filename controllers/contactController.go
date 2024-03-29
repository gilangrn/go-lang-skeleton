package controllers

import (
	"encoding/json"
	"net/http"

	"../models"
	u "../utils"
)

// CreateContact ...
var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserID = user
	resp := contact.Create()
	u.Respond(w, resp)
}

// GetContactsFor ...
var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("user").(uint)
	data := models.GetContacts(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)

	// params := mux.Vars(r)
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	//The passed path parameter is not an integer
	// 	u.Respond(w, u.Message(false, "There was an error in your request"))
	// 	return
	// }

	// data := models.GetContacts(uint(id))
	// resp := u.Message(true, "success")
	// resp["data"] = data
	// u.Respond(w, resp)
}
