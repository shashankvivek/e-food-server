// Code generated by go-swagger; DO NOT EDIT.

package guest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"e-food/api/models"
)

// AddItemOKCode is the HTTP code returned for type AddItemOK
const AddItemOKCode int = 200

/*AddItemOK Success response when item is added successfully

swagger:response addItemOK
*/
type AddItemOK struct {

	/*
	  In: Body
	*/
	Payload *models.CartSuccessResponse `json:"body,omitempty"`
}

// NewAddItemOK creates AddItemOK with default headers values
func NewAddItemOK() *AddItemOK {

	return &AddItemOK{}
}

// WithPayload adds the payload to the add item o k response
func (o *AddItemOK) WithPayload(payload *models.CartSuccessResponse) *AddItemOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add item o k response
func (o *AddItemOK) SetPayload(payload *models.CartSuccessResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddItemOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddItemBadRequestCode is the HTTP code returned for type AddItemBadRequest
const AddItemBadRequestCode int = 400

/*AddItemBadRequest Bad Request

swagger:response addItemBadRequest
*/
type AddItemBadRequest struct {
}

// NewAddItemBadRequest creates AddItemBadRequest with default headers values
func NewAddItemBadRequest() *AddItemBadRequest {

	return &AddItemBadRequest{}
}

// WriteResponse to the client
func (o *AddItemBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// AddItemNotFoundCode is the HTTP code returned for type AddItemNotFound
const AddItemNotFoundCode int = 404

/*AddItemNotFound Item to be added Not Found

swagger:response addItemNotFound
*/
type AddItemNotFound struct {
}

// NewAddItemNotFound creates AddItemNotFound with default headers values
func NewAddItemNotFound() *AddItemNotFound {

	return &AddItemNotFound{}
}

// WriteResponse to the client
func (o *AddItemNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// AddItemInternalServerErrorCode is the HTTP code returned for type AddItemInternalServerError
const AddItemInternalServerErrorCode int = 500

/*AddItemInternalServerError Server error

swagger:response addItemInternalServerError
*/
type AddItemInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddItemInternalServerError creates AddItemInternalServerError with default headers values
func NewAddItemInternalServerError() *AddItemInternalServerError {

	return &AddItemInternalServerError{}
}

// WithPayload adds the payload to the add item internal server error response
func (o *AddItemInternalServerError) WithPayload(payload string) *AddItemInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add item internal server error response
func (o *AddItemInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddItemInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
