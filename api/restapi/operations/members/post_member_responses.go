// Code generated by go-swagger; DO NOT EDIT.

package members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "secret-santa/api/models"
)

// PostMemberOKCode is the HTTP code returned for type PostMemberOK
const PostMemberOKCode int = 200

/*PostMemberOK Success

swagger:response postMemberOK
*/
type PostMemberOK struct {

	/*
	  In: Body
	*/
	Payload *models.Member `json:"body,omitempty"`
}

// NewPostMemberOK creates PostMemberOK with default headers values
func NewPostMemberOK() *PostMemberOK {

	return &PostMemberOK{}
}

// WithPayload adds the payload to the post member o k response
func (o *PostMemberOK) WithPayload(payload *models.Member) *PostMemberOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post member o k response
func (o *PostMemberOK) SetPayload(payload *models.Member) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostMemberOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostMemberBadRequestCode is the HTTP code returned for type PostMemberBadRequest
const PostMemberBadRequestCode int = 400

/*PostMemberBadRequest Bad Request

swagger:response postMemberBadRequest
*/
type PostMemberBadRequest struct {
}

// NewPostMemberBadRequest creates PostMemberBadRequest with default headers values
func NewPostMemberBadRequest() *PostMemberBadRequest {

	return &PostMemberBadRequest{}
}

// WriteResponse to the client
func (o *PostMemberBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}