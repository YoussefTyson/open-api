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

// GetAllCertificatesReader is a Reader for the GetAllCertificates structure.
type GetAllCertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllCertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllCertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllCertificatesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetAllCertificatesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllCertificatesOK creates a GetAllCertificatesOK with default headers values
func NewGetAllCertificatesOK() *GetAllCertificatesOK {
	return &GetAllCertificatesOK{}
}

/*
GetAllCertificatesOK handles this case with default header values.

Array of SNI Certificates
*/
type GetAllCertificatesOK struct {
	Payload []*models.SniCertificate
}

func (o *GetAllCertificatesOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/ssl/certificates][%d] getAllCertificatesOK  %+v", 200, o.Payload)
}

func (o *GetAllCertificatesOK) GetPayload() []*models.SniCertificate {
	return o.Payload
}

func (o *GetAllCertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllCertificatesNotFound creates a GetAllCertificatesNotFound with default headers values
func NewGetAllCertificatesNotFound() *GetAllCertificatesNotFound {
	return &GetAllCertificatesNotFound{}
}

/*
GetAllCertificatesNotFound handles this case with default header values.

Not Found
*/
type GetAllCertificatesNotFound struct {
}

func (o *GetAllCertificatesNotFound) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/ssl/certificates][%d] getAllCertificatesNotFound ", 404)
}

func (o *GetAllCertificatesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCertificatesUnprocessableEntity creates a GetAllCertificatesUnprocessableEntity with default headers values
func NewGetAllCertificatesUnprocessableEntity() *GetAllCertificatesUnprocessableEntity {
	return &GetAllCertificatesUnprocessableEntity{}
}

/*
GetAllCertificatesUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type GetAllCertificatesUnprocessableEntity struct {
}

func (o *GetAllCertificatesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/ssl/certificates][%d] getAllCertificatesUnprocessableEntity ", 422)
}

func (o *GetAllCertificatesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
