package app

import (
	"banking-app/dto"
	"banking-app/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// needs to have contact with the service
type AccountHandler struct {
	// interface to get things starting, leaving handlers to
	// just be a router
	service service.AccountService
}

func (a AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	// receives request from user side
	// customer id comes in the url so user needs to pass
	// account_type and amount
	// will need to write VALIDATION for this
	// validating incoming request lies with the service layer
	// handler just is in charge of sending and responding
	// javscript/the browser has to send the body correctly
	// to be able to DECODE it and have it match to the json
	// names
	// pull customer_if from the url
	vars := mux.Vars(r)
	customer_id := vars["id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// if it fails then someone is sending a request
		// not of NewAccountRequest type
		// bad request
		// send response back to user
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// now fully populated
	request.CustomerId = customer_id
	account, appError := a.service.NewAccount(request)
	if appError != nil {
		// don't panic cause that shuts down program
		// just tell them something went wrong
		writeResponse(w, appError.Code, appError.Message)
		return
	}

	// encode
	writeResponse(w, http.StatusCreated, account)
}

func (a AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// // get the account_id and customer_id from the URL
	// vars := mux.Vars(r)
	// accountId := vars["account_id"]
	// customerId := vars["id"]

	// // decoding incoming request
	// var request dto.TransactionRequest
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	writeResponse(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// // build the request object
	// request.AccountId = accountId
	// request.CustomerId = customerId

	// // make transaction
	// account, appError := a.service.MakeTransaction(request)

	// if appError != nil {
	// 	writeResponse(w, appError.Code, appError.AsMessage())
	// } else {
	// 	writeResponse(w, http.StatusOK, account)
	// }
	// get the account_id and customer_id from the URL
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {

		//build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// make transaction
		account, appError := a.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}

}
