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

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageProxyMetricsCasWriteConditionNotMetGetParams creates a new StorageProxyMetricsCasWriteConditionNotMetGetParams object
// with the default values initialized.
func NewStorageProxyMetricsCasWriteConditionNotMetGetParams() *StorageProxyMetricsCasWriteConditionNotMetGetParams {

	return &StorageProxyMetricsCasWriteConditionNotMetGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageProxyMetricsCasWriteConditionNotMetGetParamsWithTimeout creates a new StorageProxyMetricsCasWriteConditionNotMetGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageProxyMetricsCasWriteConditionNotMetGetParamsWithTimeout(timeout time.Duration) *StorageProxyMetricsCasWriteConditionNotMetGetParams {

	return &StorageProxyMetricsCasWriteConditionNotMetGetParams{

		timeout: timeout,
	}
}

// NewStorageProxyMetricsCasWriteConditionNotMetGetParamsWithContext creates a new StorageProxyMetricsCasWriteConditionNotMetGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageProxyMetricsCasWriteConditionNotMetGetParamsWithContext(ctx context.Context) *StorageProxyMetricsCasWriteConditionNotMetGetParams {

	return &StorageProxyMetricsCasWriteConditionNotMetGetParams{

		Context: ctx,
	}
}

// NewStorageProxyMetricsCasWriteConditionNotMetGetParamsWithHTTPClient creates a new StorageProxyMetricsCasWriteConditionNotMetGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageProxyMetricsCasWriteConditionNotMetGetParamsWithHTTPClient(client *http.Client) *StorageProxyMetricsCasWriteConditionNotMetGetParams {

	return &StorageProxyMetricsCasWriteConditionNotMetGetParams{
		HTTPClient: client,
	}
}

/*
StorageProxyMetricsCasWriteConditionNotMetGetParams contains all the parameters to send to the API endpoint
for the storage proxy metrics cas write condition not met get operation typically these are written to a http.Request
*/
type StorageProxyMetricsCasWriteConditionNotMetGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage proxy metrics cas write condition not met get params
func (o *StorageProxyMetricsCasWriteConditionNotMetGetParams) WithTimeout(timeout time.Duration) *StorageProxyMetricsCasWriteConditionNotMetGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage proxy metrics cas write condition not met get params
func (o *StorageProxyMetricsCasWriteConditionNotMetGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage proxy metrics cas write condition not met get params
func (o *StorageProxyMetricsCasWriteConditionNotMetGetParams) WithContext(ctx context.Context) *StorageProxyMetricsCasWriteConditionNotMetGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage proxy metrics cas write condition not met get params
func (o *StorageProxyMetricsCasWriteConditionNotMetGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage proxy metrics cas write condition not met get params
func (o *StorageProxyMetricsCasWriteConditionNotMetGetParams) WithHTTPClient(client *http.Client) *StorageProxyMetricsCasWriteConditionNotMetGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage proxy metrics cas write condition not met get params
func (o *StorageProxyMetricsCasWriteConditionNotMetGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageProxyMetricsCasWriteConditionNotMetGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
