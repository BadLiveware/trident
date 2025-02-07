// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SvmSSHServerModifyReader is a Reader for the SvmSSHServerModify structure.
type SvmSSHServerModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmSSHServerModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmSSHServerModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmSSHServerModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmSSHServerModifyOK creates a SvmSSHServerModifyOK with default headers values
func NewSvmSSHServerModifyOK() *SvmSSHServerModifyOK {
	return &SvmSSHServerModifyOK{}
}

/* SvmSSHServerModifyOK describes a response with status code 200, with default header values.

OK
*/
type SvmSSHServerModifyOK struct {
}

func (o *SvmSSHServerModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /security/ssh/svms/{svm.uuid}][%d] svmSshServerModifyOK ", 200)
}

func (o *SvmSSHServerModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSvmSSHServerModifyDefault creates a SvmSSHServerModifyDefault with default headers values
func NewSvmSSHServerModifyDefault(code int) *SvmSSHServerModifyDefault {
	return &SvmSSHServerModifyDefault{
		_statusCode: code,
	}
}

/* SvmSSHServerModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 10682372 | There must be at least one key exchange algorithm associated with the SSH configuration. |
| 10682373 | There must be at least one cipher associated with the SSH configuration. |
| 10682375 | Failed to modify SSH key exchange algorithms. |
| 10682378 | Failed to modify SSH ciphers. |
| 10682399 | Key exchange algorithm not supported in FIPS-enabled mode. |
| 10682400 | Failed to modify SSH MAC algorithms. |
| 10682401 | MAC algorithm not supported in FIPS-enabled mode. |
| 10682403 | There must be at least one MAC algorithm with the SSH configuration. |
| 10682413 | Failed to modify maximum authentication retry attempts. |
| 10682418 | Cipher not supported in FIPS-enabled mode. |

*/
type SvmSSHServerModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the svm ssh server modify default response
func (o *SvmSSHServerModifyDefault) Code() int {
	return o._statusCode
}

func (o *SvmSSHServerModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /security/ssh/svms/{svm.uuid}][%d] svm_ssh_server_modify default  %+v", o._statusCode, o.Payload)
}
func (o *SvmSSHServerModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmSSHServerModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
