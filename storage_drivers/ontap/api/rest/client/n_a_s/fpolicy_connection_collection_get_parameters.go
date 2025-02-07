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

// NewFpolicyConnectionCollectionGetParams creates a new FpolicyConnectionCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyConnectionCollectionGetParams() *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyConnectionCollectionGetParamsWithTimeout creates a new FpolicyConnectionCollectionGetParams object
// with the ability to set a timeout on a request.
func NewFpolicyConnectionCollectionGetParamsWithTimeout(timeout time.Duration) *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		timeout: timeout,
	}
}

// NewFpolicyConnectionCollectionGetParamsWithContext creates a new FpolicyConnectionCollectionGetParams object
// with the ability to set a context for a request.
func NewFpolicyConnectionCollectionGetParamsWithContext(ctx context.Context) *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		Context: ctx,
	}
}

// NewFpolicyConnectionCollectionGetParamsWithHTTPClient creates a new FpolicyConnectionCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyConnectionCollectionGetParamsWithHTTPClient(client *http.Client) *FpolicyConnectionCollectionGetParams {
	return &FpolicyConnectionCollectionGetParams{
		HTTPClient: client,
	}
}

/* FpolicyConnectionCollectionGetParams contains all the parameters to send to the API endpoint
   for the fpolicy connection collection get operation.

   Typically these are written to a http.Request.
*/
type FpolicyConnectionCollectionGetParams struct {

	/* DisconnectedReasonCode.

	   Filter by disconnected_reason.code
	*/
	DisconnectedReasonCodeQueryParameter *int64

	/* DisconnectedReasonMessage.

	   Filter by disconnected_reason.message
	*/
	DisconnectedReasonMessageQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* PassthroughRead.

	   Whether to view only passthrough-read connections
	*/
	PassthroughReadQueryParameter *bool

	/* PolicyName.

	   Filter by policy.name
	*/
	PolicyNameQueryParameter *string

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

	/* Server.

	   Filter by server
	*/
	ServerQueryParameter *string

	/* SessionUUID.

	   Filter by session_uuid
	*/
	SessionUUIDQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* UpdateTime.

	   Filter by update_time
	*/
	UpdateTimeQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy connection collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyConnectionCollectionGetParams) WithDefaults() *FpolicyConnectionCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy connection collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyConnectionCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := FpolicyConnectionCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithTimeout(timeout time.Duration) *FpolicyConnectionCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithContext(ctx context.Context) *FpolicyConnectionCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithHTTPClient(client *http.Client) *FpolicyConnectionCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisconnectedReasonCodeQueryParameter adds the disconnectedReasonCode to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithDisconnectedReasonCodeQueryParameter(disconnectedReasonCode *int64) *FpolicyConnectionCollectionGetParams {
	o.SetDisconnectedReasonCodeQueryParameter(disconnectedReasonCode)
	return o
}

// SetDisconnectedReasonCodeQueryParameter adds the disconnectedReasonCode to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetDisconnectedReasonCodeQueryParameter(disconnectedReasonCode *int64) {
	o.DisconnectedReasonCodeQueryParameter = disconnectedReasonCode
}

// WithDisconnectedReasonMessageQueryParameter adds the disconnectedReasonMessage to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithDisconnectedReasonMessageQueryParameter(disconnectedReasonMessage *string) *FpolicyConnectionCollectionGetParams {
	o.SetDisconnectedReasonMessageQueryParameter(disconnectedReasonMessage)
	return o
}

// SetDisconnectedReasonMessageQueryParameter adds the disconnectedReasonMessage to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetDisconnectedReasonMessageQueryParameter(disconnectedReasonMessage *string) {
	o.DisconnectedReasonMessageQueryParameter = disconnectedReasonMessage
}

// WithFieldsQueryParameter adds the fields to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithFieldsQueryParameter(fields []string) *FpolicyConnectionCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *FpolicyConnectionCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *FpolicyConnectionCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *FpolicyConnectionCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderByQueryParameter adds the orderBy to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *FpolicyConnectionCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPassthroughReadQueryParameter adds the passthroughRead to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithPassthroughReadQueryParameter(passthroughRead *bool) *FpolicyConnectionCollectionGetParams {
	o.SetPassthroughReadQueryParameter(passthroughRead)
	return o
}

// SetPassthroughReadQueryParameter adds the passthroughRead to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetPassthroughReadQueryParameter(passthroughRead *bool) {
	o.PassthroughReadQueryParameter = passthroughRead
}

// WithPolicyNameQueryParameter adds the policyName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithPolicyNameQueryParameter(policyName *string) *FpolicyConnectionCollectionGetParams {
	o.SetPolicyNameQueryParameter(policyName)
	return o
}

// SetPolicyNameQueryParameter adds the policyName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetPolicyNameQueryParameter(policyName *string) {
	o.PolicyNameQueryParameter = policyName
}

// WithReturnRecordsQueryParameter adds the returnRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *FpolicyConnectionCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *FpolicyConnectionCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithServerQueryParameter adds the server to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithServerQueryParameter(server *string) *FpolicyConnectionCollectionGetParams {
	o.SetServerQueryParameter(server)
	return o
}

// SetServerQueryParameter adds the server to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetServerQueryParameter(server *string) {
	o.ServerQueryParameter = server
}

// WithSessionUUIDQueryParameter adds the sessionUUID to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithSessionUUIDQueryParameter(sessionUUID *string) *FpolicyConnectionCollectionGetParams {
	o.SetSessionUUIDQueryParameter(sessionUUID)
	return o
}

// SetSessionUUIDQueryParameter adds the sessionUuid to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetSessionUUIDQueryParameter(sessionUUID *string) {
	o.SessionUUIDQueryParameter = sessionUUID
}

// WithStateQueryParameter adds the state to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithStateQueryParameter(state *string) *FpolicyConnectionCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithSVMNameQueryParameter adds the svmName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *FpolicyConnectionCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDPathParameter adds the svmUUID to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithSVMUUIDPathParameter(svmUUID string) *FpolicyConnectionCollectionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithTypeQueryParameter adds the typeVar to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithTypeQueryParameter(typeVar *string) *FpolicyConnectionCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithUpdateTimeQueryParameter adds the updateTime to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) WithUpdateTimeQueryParameter(updateTime *string) *FpolicyConnectionCollectionGetParams {
	o.SetUpdateTimeQueryParameter(updateTime)
	return o
}

// SetUpdateTimeQueryParameter adds the updateTime to the fpolicy connection collection get params
func (o *FpolicyConnectionCollectionGetParams) SetUpdateTimeQueryParameter(updateTime *string) {
	o.UpdateTimeQueryParameter = updateTime
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyConnectionCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisconnectedReasonCodeQueryParameter != nil {

		// query param disconnected_reason.code
		var qrDisconnectedReasonCode int64

		if o.DisconnectedReasonCodeQueryParameter != nil {
			qrDisconnectedReasonCode = *o.DisconnectedReasonCodeQueryParameter
		}
		qDisconnectedReasonCode := swag.FormatInt64(qrDisconnectedReasonCode)
		if qDisconnectedReasonCode != "" {

			if err := r.SetQueryParam("disconnected_reason.code", qDisconnectedReasonCode); err != nil {
				return err
			}
		}
	}

	if o.DisconnectedReasonMessageQueryParameter != nil {

		// query param disconnected_reason.message
		var qrDisconnectedReasonMessage string

		if o.DisconnectedReasonMessageQueryParameter != nil {
			qrDisconnectedReasonMessage = *o.DisconnectedReasonMessageQueryParameter
		}
		qDisconnectedReasonMessage := qrDisconnectedReasonMessage
		if qDisconnectedReasonMessage != "" {

			if err := r.SetQueryParam("disconnected_reason.message", qDisconnectedReasonMessage); err != nil {
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

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.PassthroughReadQueryParameter != nil {

		// query param passthrough_read
		var qrPassthroughRead bool

		if o.PassthroughReadQueryParameter != nil {
			qrPassthroughRead = *o.PassthroughReadQueryParameter
		}
		qPassthroughRead := swag.FormatBool(qrPassthroughRead)
		if qPassthroughRead != "" {

			if err := r.SetQueryParam("passthrough_read", qPassthroughRead); err != nil {
				return err
			}
		}
	}

	if o.PolicyNameQueryParameter != nil {

		// query param policy.name
		var qrPolicyName string

		if o.PolicyNameQueryParameter != nil {
			qrPolicyName = *o.PolicyNameQueryParameter
		}
		qPolicyName := qrPolicyName
		if qPolicyName != "" {

			if err := r.SetQueryParam("policy.name", qPolicyName); err != nil {
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

	if o.ServerQueryParameter != nil {

		// query param server
		var qrServer string

		if o.ServerQueryParameter != nil {
			qrServer = *o.ServerQueryParameter
		}
		qServer := qrServer
		if qServer != "" {

			if err := r.SetQueryParam("server", qServer); err != nil {
				return err
			}
		}
	}

	if o.SessionUUIDQueryParameter != nil {

		// query param session_uuid
		var qrSessionUUID string

		if o.SessionUUIDQueryParameter != nil {
			qrSessionUUID = *o.SessionUUIDQueryParameter
		}
		qSessionUUID := qrSessionUUID
		if qSessionUUID != "" {

			if err := r.SetQueryParam("session_uuid", qSessionUUID); err != nil {
				return err
			}
		}
	}

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.UpdateTimeQueryParameter != nil {

		// query param update_time
		var qrUpdateTime string

		if o.UpdateTimeQueryParameter != nil {
			qrUpdateTime = *o.UpdateTimeQueryParameter
		}
		qUpdateTime := qrUpdateTime
		if qUpdateTime != "" {

			if err := r.SetQueryParam("update_time", qUpdateTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFpolicyConnectionCollectionGet binds the parameter fields
func (o *FpolicyConnectionCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamFpolicyConnectionCollectionGet binds the parameter order_by
func (o *FpolicyConnectionCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
