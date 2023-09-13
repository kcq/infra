// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotLoadParams Defines the configuration used for handling snapshot resume. Exactly one of the two `mem_*` fields must be present in the body of the request.
//
// swagger:model SnapshotLoadParams
type SnapshotLoadParams struct {

	// Enable support for incremental (diff) snapshots by tracking dirty guest pages.
	EnableDiffSnapshots bool `json:"enable_diff_snapshots,omitempty"`

	// Configuration for the backend that handles memory load. If this field is specified, `mem_file_path` is forbidden. Either `mem_backend` or `mem_file_path` must be present at a time.
	MemBackend *MemoryBackend `json:"mem_backend,omitempty"`

	// Path to the file that contains the guest memory to be loaded. This parameter has been deprecated and is only allowed if `mem_backend` is not present.
	MemFilePath string `json:"mem_file_path,omitempty"`

	// When set to true, the vm is also resumed if the snapshot load is successful.
	ResumeVM bool `json:"resume_vm,omitempty"`

	// Path to the file that contains the microVM state to be loaded.
	// Required: true
	SnapshotPath *string `json:"snapshot_path"`
}

// Validate validates this snapshot load params
func (m *SnapshotLoadParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemBackend(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotLoadParams) validateMemBackend(formats strfmt.Registry) error {
	if swag.IsZero(m.MemBackend) { // not required
		return nil
	}

	if m.MemBackend != nil {
		if err := m.MemBackend.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mem_backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mem_backend")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotLoadParams) validateSnapshotPath(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_path", "body", m.SnapshotPath); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this snapshot load params based on the context it is used
func (m *SnapshotLoadParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMemBackend(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotLoadParams) contextValidateMemBackend(ctx context.Context, formats strfmt.Registry) error {

	if m.MemBackend != nil {

		if swag.IsZero(m.MemBackend) { // not required
			return nil
		}

		if err := m.MemBackend.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mem_backend")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mem_backend")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotLoadParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotLoadParams) UnmarshalBinary(b []byte) error {
	var res SnapshotLoadParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
