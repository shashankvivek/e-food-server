// Code generated by go-swagger; DO NOT EDIT.

package cart

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"e-food/models"
)

// GetItemsOKCode is the HTTP code returned for type GetItemsOK
const GetItemsOKCode int = 200

/*GetItemsOK All items in cart

swagger:response getItemsOK
*/
type GetItemsOK struct {

	/*
	  In: Body
	*/
	Payload models.CartPreview `json:"body,omitempty"`
}

// NewGetItemsOK creates GetItemsOK with default headers values
func NewGetItemsOK() *GetItemsOK {

	return &GetItemsOK{}
}

// WithPayload adds the payload to the get items o k response
func (o *GetItemsOK) WithPayload(payload models.CartPreview) *GetItemsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get items o k response
func (o *GetItemsOK) SetPayload(payload models.CartPreview) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.CartPreview{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetItemsBadRequestCode is the HTTP code returned for type GetItemsBadRequest
const GetItemsBadRequestCode int = 400

/*GetItemsBadRequest Bad Request

swagger:response getItemsBadRequest
*/
type GetItemsBadRequest struct {
}

// NewGetItemsBadRequest creates GetItemsBadRequest with default headers values
func NewGetItemsBadRequest() *GetItemsBadRequest {

	return &GetItemsBadRequest{}
}

// WriteResponse to the client
func (o *GetItemsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetItemsNotFoundCode is the HTTP code returned for type GetItemsNotFound
const GetItemsNotFoundCode int = 404

/*GetItemsNotFound Item to be added Not Found

swagger:response getItemsNotFound
*/
type GetItemsNotFound struct {
}

// NewGetItemsNotFound creates GetItemsNotFound with default headers values
func NewGetItemsNotFound() *GetItemsNotFound {

	return &GetItemsNotFound{}
}

// WriteResponse to the client
func (o *GetItemsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetItemsInternalServerErrorCode is the HTTP code returned for type GetItemsInternalServerError
const GetItemsInternalServerErrorCode int = 500

/*GetItemsInternalServerError Server error

swagger:response getItemsInternalServerError
*/
type GetItemsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetItemsInternalServerError creates GetItemsInternalServerError with default headers values
func NewGetItemsInternalServerError() *GetItemsInternalServerError {

	return &GetItemsInternalServerError{}
}

// WithPayload adds the payload to the get items internal server error response
func (o *GetItemsInternalServerError) WithPayload(payload string) *GetItemsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get items internal server error response
func (o *GetItemsInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
