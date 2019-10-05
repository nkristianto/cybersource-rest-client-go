// Code generated by go-swagger; DO NOT EDIT.

package download_d_t_d

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetDTDV2Reader is a Reader for the GetDTDV2 structure.
type GetDTDV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDTDV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDTDV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDTDV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDTDV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDTDV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDTDV2OK creates a GetDTDV2OK with default headers values
func NewGetDTDV2OK() *GetDTDV2OK {
	return &GetDTDV2OK{}
}

/*GetDTDV2OK handles this case with default header values.

Ok
*/
type GetDTDV2OK struct {
}

func (o *GetDTDV2OK) Error() string {
	return fmt.Sprintf("[GET /dtds/{reportDefinitionNameVersion}][%d] getDTDV2OK ", 200)
}

func (o *GetDTDV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDTDV2BadRequest creates a GetDTDV2BadRequest with default headers values
func NewGetDTDV2BadRequest() *GetDTDV2BadRequest {
	return &GetDTDV2BadRequest{}
}

/*GetDTDV2BadRequest handles this case with default header values.

Bad request. DTD file name may be invalid
*/
type GetDTDV2BadRequest struct {
}

func (o *GetDTDV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /dtds/{reportDefinitionNameVersion}][%d] getDTDV2BadRequest ", 400)
}

func (o *GetDTDV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDTDV2NotFound creates a GetDTDV2NotFound with default headers values
func NewGetDTDV2NotFound() *GetDTDV2NotFound {
	return &GetDTDV2NotFound{}
}

/*GetDTDV2NotFound handles this case with default header values.

DTD file not found
*/
type GetDTDV2NotFound struct {
}

func (o *GetDTDV2NotFound) Error() string {
	return fmt.Sprintf("[GET /dtds/{reportDefinitionNameVersion}][%d] getDTDV2NotFound ", 404)
}

func (o *GetDTDV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDTDV2InternalServerError creates a GetDTDV2InternalServerError with default headers values
func NewGetDTDV2InternalServerError() *GetDTDV2InternalServerError {
	return &GetDTDV2InternalServerError{}
}

/*GetDTDV2InternalServerError handles this case with default header values.

Internal Server Error
*/
type GetDTDV2InternalServerError struct {
}

func (o *GetDTDV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /dtds/{reportDefinitionNameVersion}][%d] getDTDV2InternalServerError ", 500)
}

func (o *GetDTDV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
