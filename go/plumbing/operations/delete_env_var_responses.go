// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// DeleteEnvVarReader is a Reader for the DeleteEnvVar structure.
type DeleteEnvVarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEnvVarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEnvVarNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteEnvVarDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteEnvVarNoContent creates a DeleteEnvVarNoContent with default headers values
func NewDeleteEnvVarNoContent() *DeleteEnvVarNoContent {
	return &DeleteEnvVarNoContent{}
}

/*DeleteEnvVarNoContent handles this case with default header values.

No Content (success)
*/
type DeleteEnvVarNoContent struct {
}

func (o *DeleteEnvVarNoContent) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{account_id}/env/{key}][%d] deleteEnvVarNoContent ", 204)
}

func (o *DeleteEnvVarNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEnvVarDefault creates a DeleteEnvVarDefault with default headers values
func NewDeleteEnvVarDefault(code int) *DeleteEnvVarDefault {
	return &DeleteEnvVarDefault{
		_statusCode: code,
	}
}

/*DeleteEnvVarDefault handles this case with default header values.

error
*/
type DeleteEnvVarDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete env var default response
func (o *DeleteEnvVarDefault) Code() int {
	return o._statusCode
}

func (o *DeleteEnvVarDefault) Error() string {
	return fmt.Sprintf("[DELETE /accounts/{account_id}/env/{key}][%d] deleteEnvVar default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteEnvVarDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteEnvVarDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
