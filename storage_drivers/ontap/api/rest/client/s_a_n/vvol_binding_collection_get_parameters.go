// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewVvolBindingCollectionGetParams creates a new VvolBindingCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVvolBindingCollectionGetParams() *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVvolBindingCollectionGetParamsWithTimeout creates a new VvolBindingCollectionGetParams object
// with the ability to set a timeout on a request.
func NewVvolBindingCollectionGetParamsWithTimeout(timeout time.Duration) *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		timeout: timeout,
	}
}

// NewVvolBindingCollectionGetParamsWithContext creates a new VvolBindingCollectionGetParams object
// with the ability to set a context for a request.
func NewVvolBindingCollectionGetParamsWithContext(ctx context.Context) *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		Context: ctx,
	}
}

// NewVvolBindingCollectionGetParamsWithHTTPClient creates a new VvolBindingCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVvolBindingCollectionGetParamsWithHTTPClient(client *http.Client) *VvolBindingCollectionGetParams {
	return &VvolBindingCollectionGetParams{
		HTTPClient: client,
	}
}

/* VvolBindingCollectionGetParams contains all the parameters to send to the API endpoint
   for the vvol binding collection get operation.

   Typically these are written to a http.Request.
*/
type VvolBindingCollectionGetParams struct {

	/* Count.

	   Filter by count
	*/
	CountQueryParameter *int64

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* ID.

	   Filter by id
	*/
	IDQueryParameter *int64

	/* IsOptimal.

	   Filter by is_optimal
	*/
	IsOptimalQueryParameter *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ProtocolEndpointName.

	   Filter by protocol_endpoint.name
	*/
	ProtocolEndpointNameQueryParameter *string

	/* ProtocolEndpointUUID.

	   Filter by protocol_endpoint.uuid
	*/
	ProtocolEndpointUUIDQueryParameter *string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* VvolName.

	   Filter by vvol.name
	*/
	VvolNameQueryParameter *string

	/* VvolUUID.

	   Filter by vvol.uuid
	*/
	VvolUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vvol binding collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VvolBindingCollectionGetParams) WithDefaults() *VvolBindingCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vvol binding collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VvolBindingCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := VvolBindingCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithTimeout(timeout time.Duration) *VvolBindingCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithContext(ctx context.Context) *VvolBindingCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithHTTPClient(client *http.Client) *VvolBindingCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCountQueryParameter adds the count to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithCountQueryParameter(count *int64) *VvolBindingCollectionGetParams {
	o.SetCountQueryParameter(count)
	return o
}

// SetCountQueryParameter adds the count to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetCountQueryParameter(count *int64) {
	o.CountQueryParameter = count
}

// WithFieldsQueryParameter adds the fields to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithFieldsQueryParameter(fields []string) *VvolBindingCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIDQueryParameter adds the id to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithIDQueryParameter(id *int64) *VvolBindingCollectionGetParams {
	o.SetIDQueryParameter(id)
	return o
}

// SetIDQueryParameter adds the id to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetIDQueryParameter(id *int64) {
	o.IDQueryParameter = id
}

// WithIsOptimalQueryParameter adds the isOptimal to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithIsOptimalQueryParameter(isOptimal *bool) *VvolBindingCollectionGetParams {
	o.SetIsOptimalQueryParameter(isOptimal)
	return o
}

// SetIsOptimalQueryParameter adds the isOptimal to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetIsOptimalQueryParameter(isOptimal *bool) {
	o.IsOptimalQueryParameter = isOptimal
}

// WithMaxRecordsQueryParameter adds the maxRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *VvolBindingCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *VvolBindingCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithProtocolEndpointNameQueryParameter adds the protocolEndpointName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithProtocolEndpointNameQueryParameter(protocolEndpointName *string) *VvolBindingCollectionGetParams {
	o.SetProtocolEndpointNameQueryParameter(protocolEndpointName)
	return o
}

// SetProtocolEndpointNameQueryParameter adds the protocolEndpointName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetProtocolEndpointNameQueryParameter(protocolEndpointName *string) {
	o.ProtocolEndpointNameQueryParameter = protocolEndpointName
}

// WithProtocolEndpointUUIDQueryParameter adds the protocolEndpointUUID to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithProtocolEndpointUUIDQueryParameter(protocolEndpointUUID *string) *VvolBindingCollectionGetParams {
	o.SetProtocolEndpointUUIDQueryParameter(protocolEndpointUUID)
	return o
}

// SetProtocolEndpointUUIDQueryParameter adds the protocolEndpointUuid to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetProtocolEndpointUUIDQueryParameter(protocolEndpointUUID *string) {
	o.ProtocolEndpointUUIDQueryParameter = protocolEndpointUUID
}

// WithReturnRecordsQueryParameter adds the returnRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *VvolBindingCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *VvolBindingCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *VvolBindingCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *VvolBindingCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithVvolNameQueryParameter adds the vvolName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithVvolNameQueryParameter(vvolName *string) *VvolBindingCollectionGetParams {
	o.SetVvolNameQueryParameter(vvolName)
	return o
}

// SetVvolNameQueryParameter adds the vvolName to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetVvolNameQueryParameter(vvolName *string) {
	o.VvolNameQueryParameter = vvolName
}

// WithVvolUUIDQueryParameter adds the vvolUUID to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) WithVvolUUIDQueryParameter(vvolUUID *string) *VvolBindingCollectionGetParams {
	o.SetVvolUUIDQueryParameter(vvolUUID)
	return o
}

// SetVvolUUIDQueryParameter adds the vvolUuid to the vvol binding collection get params
func (o *VvolBindingCollectionGetParams) SetVvolUUIDQueryParameter(vvolUUID *string) {
	o.VvolUUIDQueryParameter = vvolUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VvolBindingCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CountQueryParameter != nil {

		// query param count
		var qrCount int64

		if o.CountQueryParameter != nil {
			qrCount = *o.CountQueryParameter
		}
		qCount := swag.FormatInt64(qrCount)
		if qCount != "" {

			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IDQueryParameter != nil {

		// query param id
		var qrID int64

		if o.IDQueryParameter != nil {
			qrID = *o.IDQueryParameter
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IsOptimalQueryParameter != nil {

		// query param is_optimal
		var qrIsOptimal bool

		if o.IsOptimalQueryParameter != nil {
			qrIsOptimal = *o.IsOptimalQueryParameter
		}
		qIsOptimal := swag.FormatBool(qrIsOptimal)
		if qIsOptimal != "" {

			if err := r.SetQueryParam("is_optimal", qIsOptimal); err != nil {
				return err
			}
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ProtocolEndpointNameQueryParameter != nil {

		// query param protocol_endpoint.name
		var qrProtocolEndpointName string

		if o.ProtocolEndpointNameQueryParameter != nil {
			qrProtocolEndpointName = *o.ProtocolEndpointNameQueryParameter
		}
		qProtocolEndpointName := qrProtocolEndpointName
		if qProtocolEndpointName != "" {

			if err := r.SetQueryParam("protocol_endpoint.name", qProtocolEndpointName); err != nil {
				return err
			}
		}
	}

	if o.ProtocolEndpointUUIDQueryParameter != nil {

		// query param protocol_endpoint.uuid
		var qrProtocolEndpointUUID string

		if o.ProtocolEndpointUUIDQueryParameter != nil {
			qrProtocolEndpointUUID = *o.ProtocolEndpointUUIDQueryParameter
		}
		qProtocolEndpointUUID := qrProtocolEndpointUUID
		if qProtocolEndpointUUID != "" {

			if err := r.SetQueryParam("protocol_endpoint.uuid", qProtocolEndpointUUID); err != nil {
				return err
			}
		}
	}

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

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.VvolNameQueryParameter != nil {

		// query param vvol.name
		var qrVvolName string

		if o.VvolNameQueryParameter != nil {
			qrVvolName = *o.VvolNameQueryParameter
		}
		qVvolName := qrVvolName
		if qVvolName != "" {

			if err := r.SetQueryParam("vvol.name", qVvolName); err != nil {
				return err
			}
		}
	}

	if o.VvolUUIDQueryParameter != nil {

		// query param vvol.uuid
		var qrVvolUUID string

		if o.VvolUUIDQueryParameter != nil {
			qrVvolUUID = *o.VvolUUIDQueryParameter
		}
		qVvolUUID := qrVvolUUID
		if qVvolUUID != "" {

			if err := r.SetQueryParam("vvol.uuid", qVvolUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVvolBindingCollectionGet binds the parameter fields
func (o *VvolBindingCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVvolBindingCollectionGet binds the parameter order_by
func (o *VvolBindingCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
