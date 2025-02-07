// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewAntiRansomwareSuspectDeleteParams creates a new AntiRansomwareSuspectDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAntiRansomwareSuspectDeleteParams() *AntiRansomwareSuspectDeleteParams {
	return &AntiRansomwareSuspectDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAntiRansomwareSuspectDeleteParamsWithTimeout creates a new AntiRansomwareSuspectDeleteParams object
// with the ability to set a timeout on a request.
func NewAntiRansomwareSuspectDeleteParamsWithTimeout(timeout time.Duration) *AntiRansomwareSuspectDeleteParams {
	return &AntiRansomwareSuspectDeleteParams{
		timeout: timeout,
	}
}

// NewAntiRansomwareSuspectDeleteParamsWithContext creates a new AntiRansomwareSuspectDeleteParams object
// with the ability to set a context for a request.
func NewAntiRansomwareSuspectDeleteParamsWithContext(ctx context.Context) *AntiRansomwareSuspectDeleteParams {
	return &AntiRansomwareSuspectDeleteParams{
		Context: ctx,
	}
}

// NewAntiRansomwareSuspectDeleteParamsWithHTTPClient creates a new AntiRansomwareSuspectDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAntiRansomwareSuspectDeleteParamsWithHTTPClient(client *http.Client) *AntiRansomwareSuspectDeleteParams {
	return &AntiRansomwareSuspectDeleteParams{
		HTTPClient: client,
	}
}

/* AntiRansomwareSuspectDeleteParams contains all the parameters to send to the API endpoint
   for the anti ransomware suspect delete operation.

   Typically these are written to a http.Request.
*/
type AntiRansomwareSuspectDeleteParams struct {

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* VolumeUUID.

	   Identification of the Anti-ransomware suspect file for the deletion.

	   Format: uuid
	*/
	VolumeUUIDPathParameter strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the anti ransomware suspect delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AntiRansomwareSuspectDeleteParams) WithDefaults() *AntiRansomwareSuspectDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the anti ransomware suspect delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AntiRansomwareSuspectDeleteParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := AntiRansomwareSuspectDeleteParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) WithTimeout(timeout time.Duration) *AntiRansomwareSuspectDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) WithContext(ctx context.Context) *AntiRansomwareSuspectDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) WithHTTPClient(client *http.Client) *AntiRansomwareSuspectDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnRecordsQueryParameter adds the returnRecords to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) WithReturnRecordsQueryParameter(returnRecords *bool) *AntiRansomwareSuspectDeleteParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *AntiRansomwareSuspectDeleteParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithVolumeUUIDPathParameter adds the volumeUUID to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) WithVolumeUUIDPathParameter(volumeUUID strfmt.UUID) *AntiRansomwareSuspectDeleteParams {
	o.SetVolumeUUIDPathParameter(volumeUUID)
	return o
}

// SetVolumeUUIDPathParameter adds the volumeUuid to the anti ransomware suspect delete params
func (o *AntiRansomwareSuspectDeleteParams) SetVolumeUUIDPathParameter(volumeUUID strfmt.UUID) {
	o.VolumeUUIDPathParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AntiRansomwareSuspectDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUIDPathParameter.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
