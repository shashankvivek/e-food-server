// Code generated by go-swagger; DO NOT EDIT.

package menu

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"e-food/models"
)

// CategoryListOKCode is the HTTP code returned for type CategoryListOK
const CategoryListOKCode int = 200

/*CategoryListOK Get Category to show menu

swagger:response categoryListOK
*/
type CategoryListOK struct {

	/*
	  In: Body
	*/
	Payload models.Categories `json:"body,omitempty"`
}

// NewCategoryListOK creates CategoryListOK with default headers values
func NewCategoryListOK() *CategoryListOK {

	return &CategoryListOK{}
}

// WithPayload adds the payload to the category list o k response
func (o *CategoryListOK) WithPayload(payload models.Categories) *CategoryListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the category list o k response
func (o *CategoryListOK) SetPayload(payload models.Categories) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CategoryListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Categories{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// CategoryListBadRequestCode is the HTTP code returned for type CategoryListBadRequest
const CategoryListBadRequestCode int = 400

/*CategoryListBadRequest Bad Request

swagger:response categoryListBadRequest
*/
type CategoryListBadRequest struct {
}

// NewCategoryListBadRequest creates CategoryListBadRequest with default headers values
func NewCategoryListBadRequest() *CategoryListBadRequest {

	return &CategoryListBadRequest{}
}

// WriteResponse to the client
func (o *CategoryListBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// CategoryListNotFoundCode is the HTTP code returned for type CategoryListNotFound
const CategoryListNotFoundCode int = 404

/*CategoryListNotFound Categories Not Found

swagger:response categoryListNotFound
*/
type CategoryListNotFound struct {
}

// NewCategoryListNotFound creates CategoryListNotFound with default headers values
func NewCategoryListNotFound() *CategoryListNotFound {

	return &CategoryListNotFound{}
}

// WriteResponse to the client
func (o *CategoryListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// CategoryListInternalServerErrorCode is the HTTP code returned for type CategoryListInternalServerError
const CategoryListInternalServerErrorCode int = 500

/*CategoryListInternalServerError Server Error

swagger:response categoryListInternalServerError
*/
type CategoryListInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCategoryListInternalServerError creates CategoryListInternalServerError with default headers values
func NewCategoryListInternalServerError() *CategoryListInternalServerError {

	return &CategoryListInternalServerError{}
}

// WithPayload adds the payload to the category list internal server error response
func (o *CategoryListInternalServerError) WithPayload(payload string) *CategoryListInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the category list internal server error response
func (o *CategoryListInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CategoryListInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
