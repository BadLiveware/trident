// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// IpspacesCreateReader is a Reader for the IpspacesCreate structure.
type IpspacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpspacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpspacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpspacesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpspacesCreateCreated creates a IpspacesCreateCreated with default headers values
func NewIpspacesCreateCreated() *IpspacesCreateCreated {
	return &IpspacesCreateCreated{}
}

/* IpspacesCreateCreated describes a response with status code 201, with default header values.

Created
*/
type IpspacesCreateCreated struct {
}

func (o *IpspacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /network/ipspaces][%d] ipspacesCreateCreated ", 201)
}

func (o *IpspacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpspacesCreateDefault creates a IpspacesCreateDefault with default headers values
func NewIpspacesCreateDefault(code int) *IpspacesCreateDefault {
	return &IpspacesCreateDefault{
		_statusCode: code,
	}
}

/* IpspacesCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1966586 | The specified IPspace name is invalid because it is already used by a peered SVM. |
| 1967102 | A POST operation might have left the configuration in an inconsistent state. Check the configuration. |
ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 9240591 | The name is not valid. The name is already in use by a cluster node, Vserver, or it is the name of the local cluster. |

*/
type IpspacesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the ipspaces create default response
func (o *IpspacesCreateDefault) Code() int {
	return o._statusCode
}

func (o *IpspacesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /network/ipspaces][%d] ipspaces_create default  %+v", o._statusCode, o.Payload)
}
func (o *IpspacesCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IpspacesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
