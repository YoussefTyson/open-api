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

// SetEnvVarValueReader is a Reader for the SetEnvVarValue structure.
type SetEnvVarValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEnvVarValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSetEnvVarValueCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetEnvVarValueDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetEnvVarValueCreated creates a SetEnvVarValueCreated with default headers values
func NewSetEnvVarValueCreated() *SetEnvVarValueCreated {
	return &SetEnvVarValueCreated{}
}

/*SetEnvVarValueCreated handles this case with default header values.

Created (success)
*/
type SetEnvVarValueCreated struct {
	Payload *models.EnvVar
}

func (o *SetEnvVarValueCreated) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{account_id}/env/{key}][%d] setEnvVarValueCreated  %+v", 201, o.Payload)
}

func (o *SetEnvVarValueCreated) GetPayload() *models.EnvVar {
	return o.Payload
}

func (o *SetEnvVarValueCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvVar)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetEnvVarValueDefault creates a SetEnvVarValueDefault with default headers values
func NewSetEnvVarValueDefault(code int) *SetEnvVarValueDefault {
	return &SetEnvVarValueDefault{
		_statusCode: code,
	}
}

/*SetEnvVarValueDefault handles this case with default header values.

error
*/
type SetEnvVarValueDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the set env var value default response
func (o *SetEnvVarValueDefault) Code() int {
	return o._statusCode
}

func (o *SetEnvVarValueDefault) Error() string {
	return fmt.Sprintf("[PATCH /accounts/{account_id}/env/{key}][%d] setEnvVarValue default  %+v", o._statusCode, o.Payload)
}

func (o *SetEnvVarValueDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetEnvVarValueDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
