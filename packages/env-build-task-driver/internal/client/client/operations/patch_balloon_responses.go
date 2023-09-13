// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/e2b-dev/api/packages/env-build-task-driver/internal/client/models"
)

// PatchBalloonReader is a Reader for the PatchBalloon structure.
type PatchBalloonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBalloonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchBalloonNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchBalloonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchBalloonDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchBalloonNoContent creates a PatchBalloonNoContent with default headers values
func NewPatchBalloonNoContent() *PatchBalloonNoContent {
	return &PatchBalloonNoContent{}
}

/*
PatchBalloonNoContent describes a response with status code 204, with default header values.

Balloon device updated
*/
type PatchBalloonNoContent struct {
}

// IsSuccess returns true when this patch balloon no content response has a 2xx status code
func (o *PatchBalloonNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch balloon no content response has a 3xx status code
func (o *PatchBalloonNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch balloon no content response has a 4xx status code
func (o *PatchBalloonNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch balloon no content response has a 5xx status code
func (o *PatchBalloonNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch balloon no content response a status code equal to that given
func (o *PatchBalloonNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the patch balloon no content response
func (o *PatchBalloonNoContent) Code() int {
	return 204
}

func (o *PatchBalloonNoContent) Error() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloonNoContent ", 204)
}

func (o *PatchBalloonNoContent) String() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloonNoContent ", 204)
}

func (o *PatchBalloonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchBalloonBadRequest creates a PatchBalloonBadRequest with default headers values
func NewPatchBalloonBadRequest() *PatchBalloonBadRequest {
	return &PatchBalloonBadRequest{}
}

/*
PatchBalloonBadRequest describes a response with status code 400, with default header values.

Balloon device cannot be updated due to bad input
*/
type PatchBalloonBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this patch balloon bad request response has a 2xx status code
func (o *PatchBalloonBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch balloon bad request response has a 3xx status code
func (o *PatchBalloonBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch balloon bad request response has a 4xx status code
func (o *PatchBalloonBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch balloon bad request response has a 5xx status code
func (o *PatchBalloonBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch balloon bad request response a status code equal to that given
func (o *PatchBalloonBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch balloon bad request response
func (o *PatchBalloonBadRequest) Code() int {
	return 400
}

func (o *PatchBalloonBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloonBadRequest  %+v", 400, o.Payload)
}

func (o *PatchBalloonBadRequest) String() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloonBadRequest  %+v", 400, o.Payload)
}

func (o *PatchBalloonBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchBalloonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBalloonDefault creates a PatchBalloonDefault with default headers values
func NewPatchBalloonDefault(code int) *PatchBalloonDefault {
	return &PatchBalloonDefault{
		_statusCode: code,
	}
}

/*
PatchBalloonDefault describes a response with status code -1, with default header values.

Internal server error
*/
type PatchBalloonDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this patch balloon default response has a 2xx status code
func (o *PatchBalloonDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch balloon default response has a 3xx status code
func (o *PatchBalloonDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch balloon default response has a 4xx status code
func (o *PatchBalloonDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch balloon default response has a 5xx status code
func (o *PatchBalloonDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch balloon default response a status code equal to that given
func (o *PatchBalloonDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch balloon default response
func (o *PatchBalloonDefault) Code() int {
	return o._statusCode
}

func (o *PatchBalloonDefault) Error() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloon default  %+v", o._statusCode, o.Payload)
}

func (o *PatchBalloonDefault) String() string {
	return fmt.Sprintf("[PATCH /balloon][%d] patchBalloon default  %+v", o._statusCode, o.Payload)
}

func (o *PatchBalloonDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchBalloonDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
