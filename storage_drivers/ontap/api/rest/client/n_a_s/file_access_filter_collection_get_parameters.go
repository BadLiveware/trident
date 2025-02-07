// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewFileAccessFilterCollectionGetParams creates a new FileAccessFilterCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileAccessFilterCollectionGetParams() *FileAccessFilterCollectionGetParams {
	return &FileAccessFilterCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileAccessFilterCollectionGetParamsWithTimeout creates a new FileAccessFilterCollectionGetParams object
// with the ability to set a timeout on a request.
func NewFileAccessFilterCollectionGetParamsWithTimeout(timeout time.Duration) *FileAccessFilterCollectionGetParams {
	return &FileAccessFilterCollectionGetParams{
		timeout: timeout,
	}
}

// NewFileAccessFilterCollectionGetParamsWithContext creates a new FileAccessFilterCollectionGetParams object
// with the ability to set a context for a request.
func NewFileAccessFilterCollectionGetParamsWithContext(ctx context.Context) *FileAccessFilterCollectionGetParams {
	return &FileAccessFilterCollectionGetParams{
		Context: ctx,
	}
}

// NewFileAccessFilterCollectionGetParamsWithHTTPClient creates a new FileAccessFilterCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileAccessFilterCollectionGetParamsWithHTTPClient(client *http.Client) *FileAccessFilterCollectionGetParams {
	return &FileAccessFilterCollectionGetParams{
		HTTPClient: client,
	}
}

/* FileAccessFilterCollectionGetParams contains all the parameters to send to the API endpoint
   for the file access filter collection get operation.

   Typically these are written to a http.Request.
*/
type FileAccessFilterCollectionGetParams struct {

	/* ClientIP.

	   Filter by client_ip
	*/
	ClientIPQueryParameter *string

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* Path.

	   Filter by path
	*/
	PathQueryParameter *string

	/* Protocol.

	   Filter by protocol
	*/
	ProtocolQueryParameter *string

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

	/* TraceAllowedOps.

	   Filter by trace_allowed_ops
	*/
	TraceAllowedOpsQueryParameter *bool

	/* UnixUser.

	   Filter by unix_user
	*/
	UnixUserQueryParameter *string

	/* WindowsUser.

	   Filter by windows_user
	*/
	WindowsUserQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file access filter collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileAccessFilterCollectionGetParams) WithDefaults() *FileAccessFilterCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file access filter collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileAccessFilterCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := FileAccessFilterCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithTimeout(timeout time.Duration) *FileAccessFilterCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithContext(ctx context.Context) *FileAccessFilterCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithHTTPClient(client *http.Client) *FileAccessFilterCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientIPQueryParameter adds the clientIP to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithClientIPQueryParameter(clientIP *string) *FileAccessFilterCollectionGetParams {
	o.SetClientIPQueryParameter(clientIP)
	return o
}

// SetClientIPQueryParameter adds the clientIp to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetClientIPQueryParameter(clientIP *string) {
	o.ClientIPQueryParameter = clientIP
}

// WithEnabledQueryParameter adds the enabled to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithEnabledQueryParameter(enabled *bool) *FileAccessFilterCollectionGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithFieldsQueryParameter adds the fields to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithFieldsQueryParameter(fields []string) *FileAccessFilterCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexQueryParameter adds the index to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithIndexQueryParameter(index *int64) *FileAccessFilterCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecordsQueryParameter adds the maxRecords to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *FileAccessFilterCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *FileAccessFilterCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPathQueryParameter adds the path to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithPathQueryParameter(path *string) *FileAccessFilterCollectionGetParams {
	o.SetPathQueryParameter(path)
	return o
}

// SetPathQueryParameter adds the path to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetPathQueryParameter(path *string) {
	o.PathQueryParameter = path
}

// WithProtocolQueryParameter adds the protocol to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithProtocolQueryParameter(protocol *string) *FileAccessFilterCollectionGetParams {
	o.SetProtocolQueryParameter(protocol)
	return o
}

// SetProtocolQueryParameter adds the protocol to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetProtocolQueryParameter(protocol *string) {
	o.ProtocolQueryParameter = protocol
}

// WithReturnRecordsQueryParameter adds the returnRecords to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *FileAccessFilterCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *FileAccessFilterCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *FileAccessFilterCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *FileAccessFilterCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTraceAllowedOpsQueryParameter adds the traceAllowedOps to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithTraceAllowedOpsQueryParameter(traceAllowedOps *bool) *FileAccessFilterCollectionGetParams {
	o.SetTraceAllowedOpsQueryParameter(traceAllowedOps)
	return o
}

// SetTraceAllowedOpsQueryParameter adds the traceAllowedOps to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetTraceAllowedOpsQueryParameter(traceAllowedOps *bool) {
	o.TraceAllowedOpsQueryParameter = traceAllowedOps
}

// WithUnixUserQueryParameter adds the unixUser to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithUnixUserQueryParameter(unixUser *string) *FileAccessFilterCollectionGetParams {
	o.SetUnixUserQueryParameter(unixUser)
	return o
}

// SetUnixUserQueryParameter adds the unixUser to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetUnixUserQueryParameter(unixUser *string) {
	o.UnixUserQueryParameter = unixUser
}

// WithWindowsUserQueryParameter adds the windowsUser to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) WithWindowsUserQueryParameter(windowsUser *string) *FileAccessFilterCollectionGetParams {
	o.SetWindowsUserQueryParameter(windowsUser)
	return o
}

// SetWindowsUserQueryParameter adds the windowsUser to the file access filter collection get params
func (o *FileAccessFilterCollectionGetParams) SetWindowsUserQueryParameter(windowsUser *string) {
	o.WindowsUserQueryParameter = windowsUser
}

// WriteToRequest writes these params to a swagger request
func (o *FileAccessFilterCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientIPQueryParameter != nil {

		// query param client_ip
		var qrClientIP string

		if o.ClientIPQueryParameter != nil {
			qrClientIP = *o.ClientIPQueryParameter
		}
		qClientIP := qrClientIP
		if qClientIP != "" {

			if err := r.SetQueryParam("client_ip", qClientIP); err != nil {
				return err
			}
		}
	}

	if o.EnabledQueryParameter != nil {

		// query param enabled
		var qrEnabled bool

		if o.EnabledQueryParameter != nil {
			qrEnabled = *o.EnabledQueryParameter
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
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

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
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

	if o.PathQueryParameter != nil {

		// query param path
		var qrPath string

		if o.PathQueryParameter != nil {
			qrPath = *o.PathQueryParameter
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	if o.ProtocolQueryParameter != nil {

		// query param protocol
		var qrProtocol string

		if o.ProtocolQueryParameter != nil {
			qrProtocol = *o.ProtocolQueryParameter
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
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

	if o.TraceAllowedOpsQueryParameter != nil {

		// query param trace_allowed_ops
		var qrTraceAllowedOps bool

		if o.TraceAllowedOpsQueryParameter != nil {
			qrTraceAllowedOps = *o.TraceAllowedOpsQueryParameter
		}
		qTraceAllowedOps := swag.FormatBool(qrTraceAllowedOps)
		if qTraceAllowedOps != "" {

			if err := r.SetQueryParam("trace_allowed_ops", qTraceAllowedOps); err != nil {
				return err
			}
		}
	}

	if o.UnixUserQueryParameter != nil {

		// query param unix_user
		var qrUnixUser string

		if o.UnixUserQueryParameter != nil {
			qrUnixUser = *o.UnixUserQueryParameter
		}
		qUnixUser := qrUnixUser
		if qUnixUser != "" {

			if err := r.SetQueryParam("unix_user", qUnixUser); err != nil {
				return err
			}
		}
	}

	if o.WindowsUserQueryParameter != nil {

		// query param windows_user
		var qrWindowsUser string

		if o.WindowsUserQueryParameter != nil {
			qrWindowsUser = *o.WindowsUserQueryParameter
		}
		qWindowsUser := qrWindowsUser
		if qWindowsUser != "" {

			if err := r.SetQueryParam("windows_user", qWindowsUser); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFileAccessFilterCollectionGet binds the parameter fields
func (o *FileAccessFilterCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamFileAccessFilterCollectionGet binds the parameter order_by
func (o *FileAccessFilterCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
