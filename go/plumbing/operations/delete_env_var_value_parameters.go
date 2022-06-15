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

// NewDeleteEnvVarValueParams creates a new DeleteEnvVarValueParams object
// with the default values initialized.
func NewDeleteEnvVarValueParams() *DeleteEnvVarValueParams {
	var ()
	return &DeleteEnvVarValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteEnvVarValueParamsWithTimeout creates a new DeleteEnvVarValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteEnvVarValueParamsWithTimeout(timeout time.Duration) *DeleteEnvVarValueParams {
	var ()
	return &DeleteEnvVarValueParams{

		timeout: timeout,
	}
}

// NewDeleteEnvVarValueParamsWithContext creates a new DeleteEnvVarValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteEnvVarValueParamsWithContext(ctx context.Context) *DeleteEnvVarValueParams {
	var ()
	return &DeleteEnvVarValueParams{

		Context: ctx,
	}
}

// NewDeleteEnvVarValueParamsWithHTTPClient creates a new DeleteEnvVarValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteEnvVarValueParamsWithHTTPClient(client *http.Client) *DeleteEnvVarValueParams {
	var ()
	return &DeleteEnvVarValueParams{
		HTTPClient: client,
	}
}

/*DeleteEnvVarValueParams contains all the parameters to send to the API endpoint
for the delete env var value operation typically these are written to a http.Request
*/
type DeleteEnvVarValueParams struct {

	/*AccountID
	  Scope response to account_id

	*/
	AccountID string
	/*ID
	  The environment variable value's ID

	*/
	ID string
	/*Key
	  The environment variable key name (case-sensitive)

	*/
	Key string
	/*SiteID
	  If provided, delete the value from an environment variable on this site

	*/
	SiteID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete env var value params
func (o *DeleteEnvVarValueParams) WithTimeout(timeout time.Duration) *DeleteEnvVarValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete env var value params
func (o *DeleteEnvVarValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete env var value params
func (o *DeleteEnvVarValueParams) WithContext(ctx context.Context) *DeleteEnvVarValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete env var value params
func (o *DeleteEnvVarValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete env var value params
func (o *DeleteEnvVarValueParams) WithHTTPClient(client *http.Client) *DeleteEnvVarValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete env var value params
func (o *DeleteEnvVarValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the delete env var value params
func (o *DeleteEnvVarValueParams) WithAccountID(accountID string) *DeleteEnvVarValueParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete env var value params
func (o *DeleteEnvVarValueParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithID adds the id to the delete env var value params
func (o *DeleteEnvVarValueParams) WithID(id string) *DeleteEnvVarValueParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete env var value params
func (o *DeleteEnvVarValueParams) SetID(id string) {
	o.ID = id
}

// WithKey adds the key to the delete env var value params
func (o *DeleteEnvVarValueParams) WithKey(key string) *DeleteEnvVarValueParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the delete env var value params
func (o *DeleteEnvVarValueParams) SetKey(key string) {
	o.Key = key
}

// WithSiteID adds the siteID to the delete env var value params
func (o *DeleteEnvVarValueParams) WithSiteID(siteID *string) *DeleteEnvVarValueParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the delete env var value params
func (o *DeleteEnvVarValueParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteEnvVarValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
