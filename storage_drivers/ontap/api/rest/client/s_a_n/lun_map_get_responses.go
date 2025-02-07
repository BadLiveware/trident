// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LunMapGetReader is a Reader for the LunMapGet structure.
type LunMapGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunMapGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapGetOK creates a LunMapGetOK with default headers values
func NewLunMapGetOK() *LunMapGetOK {
	return &LunMapGetOK{}
}

/* LunMapGetOK describes a response with status code 200, with default header values.

OK
*/
type LunMapGetOK struct {
	Payload *models.LunMap
}

func (o *LunMapGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lunMapGetOK  %+v", 200, o.Payload)
}
func (o *LunMapGetOK) GetPayload() *models.LunMap {
	return o.Payload
}

func (o *LunMapGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunMapGetDefault creates a LunMapGetDefault with default headers values
func NewLunMapGetDefault(code int) *LunMapGetDefault {
	return &LunMapGetDefault{
		_statusCode: code,
	}
}

/* LunMapGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5374852 | The initiator group does not exist or is not accessible to the caller. |
| 5374875 | The LUN does not exist or is not accessible to the caller. |

*/
type LunMapGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun map get default response
func (o *LunMapGetDefault) Code() int {
	return o._statusCode
}

func (o *LunMapGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}][%d] lun_map_get default  %+v", o._statusCode, o.Payload)
}
func (o *LunMapGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
