// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-operator/pkg/scyllaclient/internal/scylla_v2/models"
)

// FindConfigExperimentalReader is a Reader for the FindConfigExperimental structure.
type FindConfigExperimentalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigExperimentalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigExperimentalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigExperimentalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigExperimentalOK creates a FindConfigExperimentalOK with default headers values
func NewFindConfigExperimentalOK() *FindConfigExperimentalOK {
	return &FindConfigExperimentalOK{}
}

/*FindConfigExperimentalOK handles this case with default header values.

Config value
*/
type FindConfigExperimentalOK struct {
	Payload bool
}

func (o *FindConfigExperimentalOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigExperimentalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigExperimentalDefault creates a FindConfigExperimentalDefault with default headers values
func NewFindConfigExperimentalDefault(code int) *FindConfigExperimentalDefault {
	return &FindConfigExperimentalDefault{
		_statusCode: code,
	}
}

/*FindConfigExperimentalDefault handles this case with default header values.

unexpected error
*/
type FindConfigExperimentalDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config experimental default response
func (o *FindConfigExperimentalDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigExperimentalDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigExperimentalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigExperimentalDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}