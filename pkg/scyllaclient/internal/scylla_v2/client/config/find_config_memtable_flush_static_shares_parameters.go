// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindConfigMemtableFlushStaticSharesParams creates a new FindConfigMemtableFlushStaticSharesParams object
// with the default values initialized.
func NewFindConfigMemtableFlushStaticSharesParams() *FindConfigMemtableFlushStaticSharesParams {

	return &FindConfigMemtableFlushStaticSharesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigMemtableFlushStaticSharesParamsWithTimeout creates a new FindConfigMemtableFlushStaticSharesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigMemtableFlushStaticSharesParamsWithTimeout(timeout time.Duration) *FindConfigMemtableFlushStaticSharesParams {

	return &FindConfigMemtableFlushStaticSharesParams{

		timeout: timeout,
	}
}

// NewFindConfigMemtableFlushStaticSharesParamsWithContext creates a new FindConfigMemtableFlushStaticSharesParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigMemtableFlushStaticSharesParamsWithContext(ctx context.Context) *FindConfigMemtableFlushStaticSharesParams {

	return &FindConfigMemtableFlushStaticSharesParams{

		Context: ctx,
	}
}

// NewFindConfigMemtableFlushStaticSharesParamsWithHTTPClient creates a new FindConfigMemtableFlushStaticSharesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigMemtableFlushStaticSharesParamsWithHTTPClient(client *http.Client) *FindConfigMemtableFlushStaticSharesParams {

	return &FindConfigMemtableFlushStaticSharesParams{
		HTTPClient: client,
	}
}

/*
FindConfigMemtableFlushStaticSharesParams contains all the parameters to send to the API endpoint
for the find config memtable flush static shares operation typically these are written to a http.Request
*/
type FindConfigMemtableFlushStaticSharesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config memtable flush static shares params
func (o *FindConfigMemtableFlushStaticSharesParams) WithTimeout(timeout time.Duration) *FindConfigMemtableFlushStaticSharesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config memtable flush static shares params
func (o *FindConfigMemtableFlushStaticSharesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config memtable flush static shares params
func (o *FindConfigMemtableFlushStaticSharesParams) WithContext(ctx context.Context) *FindConfigMemtableFlushStaticSharesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config memtable flush static shares params
func (o *FindConfigMemtableFlushStaticSharesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config memtable flush static shares params
func (o *FindConfigMemtableFlushStaticSharesParams) WithHTTPClient(client *http.Client) *FindConfigMemtableFlushStaticSharesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config memtable flush static shares params
func (o *FindConfigMemtableFlushStaticSharesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigMemtableFlushStaticSharesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
