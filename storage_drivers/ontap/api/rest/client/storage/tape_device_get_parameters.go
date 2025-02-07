// Code generated by go-swagger; DO NOT EDIT.

package storage

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
	"github.com/go-openapi/swag"
)

// NewTapeDeviceGetParams creates a new TapeDeviceGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTapeDeviceGetParams() *TapeDeviceGetParams {
	return &TapeDeviceGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTapeDeviceGetParamsWithTimeout creates a new TapeDeviceGetParams object
// with the ability to set a timeout on a request.
func NewTapeDeviceGetParamsWithTimeout(timeout time.Duration) *TapeDeviceGetParams {
	return &TapeDeviceGetParams{
		timeout: timeout,
	}
}

// NewTapeDeviceGetParamsWithContext creates a new TapeDeviceGetParams object
// with the ability to set a context for a request.
func NewTapeDeviceGetParamsWithContext(ctx context.Context) *TapeDeviceGetParams {
	return &TapeDeviceGetParams{
		Context: ctx,
	}
}

// NewTapeDeviceGetParamsWithHTTPClient creates a new TapeDeviceGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTapeDeviceGetParamsWithHTTPClient(client *http.Client) *TapeDeviceGetParams {
	return &TapeDeviceGetParams{
		HTTPClient: client,
	}
}

/* TapeDeviceGetParams contains all the parameters to send to the API endpoint
   for the tape device get operation.

   Typically these are written to a http.Request.
*/
type TapeDeviceGetParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceIDPathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tape device get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TapeDeviceGetParams) WithDefaults() *TapeDeviceGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tape device get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TapeDeviceGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tape device get params
func (o *TapeDeviceGetParams) WithTimeout(timeout time.Duration) *TapeDeviceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tape device get params
func (o *TapeDeviceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tape device get params
func (o *TapeDeviceGetParams) WithContext(ctx context.Context) *TapeDeviceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tape device get params
func (o *TapeDeviceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tape device get params
func (o *TapeDeviceGetParams) WithHTTPClient(client *http.Client) *TapeDeviceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tape device get params
func (o *TapeDeviceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceIDPathParameter adds the deviceID to the tape device get params
func (o *TapeDeviceGetParams) WithDeviceIDPathParameter(deviceID string) *TapeDeviceGetParams {
	o.SetDeviceIDPathParameter(deviceID)
	return o
}

// SetDeviceIDPathParameter adds the deviceId to the tape device get params
func (o *TapeDeviceGetParams) SetDeviceIDPathParameter(deviceID string) {
	o.DeviceIDPathParameter = deviceID
}

// WithFieldsQueryParameter adds the fields to the tape device get params
func (o *TapeDeviceGetParams) WithFieldsQueryParameter(fields []string) *TapeDeviceGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the tape device get params
func (o *TapeDeviceGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNodeUUIDPathParameter adds the nodeUUID to the tape device get params
func (o *TapeDeviceGetParams) WithNodeUUIDPathParameter(nodeUUID string) *TapeDeviceGetParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the tape device get params
func (o *TapeDeviceGetParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *TapeDeviceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceIDPathParameter); err != nil {
		return err
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTapeDeviceGet binds the parameter fields
func (o *TapeDeviceGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
