// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetAllCertificatesParams creates a new GetAllCertificatesParams object
// with the default values initialized.
func NewGetAllCertificatesParams() *GetAllCertificatesParams {
	var ()
	return &GetAllCertificatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllCertificatesParamsWithTimeout creates a new GetAllCertificatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllCertificatesParamsWithTimeout(timeout time.Duration) *GetAllCertificatesParams {
	var ()
	return &GetAllCertificatesParams{

		timeout: timeout,
	}
}

// NewGetAllCertificatesParamsWithContext creates a new GetAllCertificatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllCertificatesParamsWithContext(ctx context.Context) *GetAllCertificatesParams {
	var ()
	return &GetAllCertificatesParams{

		Context: ctx,
	}
}

// NewGetAllCertificatesParamsWithHTTPClient creates a new GetAllCertificatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllCertificatesParamsWithHTTPClient(client *http.Client) *GetAllCertificatesParams {
	var ()
	return &GetAllCertificatesParams{
		HTTPClient: client,
	}
}

/*
GetAllCertificatesParams contains all the parameters to send to the API endpoint
for the get all certificates operation typically these are written to a http.Request
*/
type GetAllCertificatesParams struct {

	/*Domain*/
	Domain string
	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all certificates params
func (o *GetAllCertificatesParams) WithTimeout(timeout time.Duration) *GetAllCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all certificates params
func (o *GetAllCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all certificates params
func (o *GetAllCertificatesParams) WithContext(ctx context.Context) *GetAllCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all certificates params
func (o *GetAllCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all certificates params
func (o *GetAllCertificatesParams) WithHTTPClient(client *http.Client) *GetAllCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all certificates params
func (o *GetAllCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomain adds the domain to the get all certificates params
func (o *GetAllCertificatesParams) WithDomain(domain string) *GetAllCertificatesParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the get all certificates params
func (o *GetAllCertificatesParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithSiteID adds the siteID to the get all certificates params
func (o *GetAllCertificatesParams) WithSiteID(siteID string) *GetAllCertificatesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get all certificates params
func (o *GetAllCertificatesParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param domain
	qrDomain := o.Domain
	qDomain := qrDomain
	if qDomain != "" {
		if err := r.SetQueryParam("domain", qDomain); err != nil {
			return err
		}
	}

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
