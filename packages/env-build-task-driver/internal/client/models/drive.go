// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Drive drive
//
// swagger:model Drive
type Drive struct {

	// Represents the caching strategy for the block device.
	// Enum: [Unsafe Writeback]
	CacheType *string `json:"cache_type,omitempty"`

	// drive id
	// Required: true
	DriveID *string `json:"drive_id"`

	// Type of the IO engine used by the device. "Async" is supported on host kernels newer than 5.10.51.
	// Enum: [Sync Async]
	IoEngine *string `json:"io_engine,omitempty"`

	// is read only
	// Required: true
	IsReadOnly *bool `json:"is_read_only"`

	// is root device
	// Required: true
	IsRootDevice *bool `json:"is_root_device"`

	// Represents the unique id of the boot partition of this device. It is optional and it will be taken into account only if the is_root_device field is true.
	Partuuid string `json:"partuuid,omitempty"`

	// Host level path for the guest drive
	// Required: true
	PathOnHost *string `json:"path_on_host"`

	// rate limiter
	RateLimiter *RateLimiter `json:"rate_limiter,omitempty"`
}

// Validate validates this drive
func (m *Drive) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCacheType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDriveID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsReadOnly(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRootDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathOnHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRateLimiter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var driveTypeCacheTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unsafe","Writeback"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		driveTypeCacheTypePropEnum = append(driveTypeCacheTypePropEnum, v)
	}
}

const (

	// DriveCacheTypeUnsafe captures enum value "Unsafe"
	DriveCacheTypeUnsafe string = "Unsafe"

	// DriveCacheTypeWriteback captures enum value "Writeback"
	DriveCacheTypeWriteback string = "Writeback"
)

// prop value enum
func (m *Drive) validateCacheTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, driveTypeCacheTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Drive) validateCacheType(formats strfmt.Registry) error {
	if swag.IsZero(m.CacheType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCacheTypeEnum("cache_type", "body", *m.CacheType); err != nil {
		return err
	}

	return nil
}

func (m *Drive) validateDriveID(formats strfmt.Registry) error {

	if err := validate.Required("drive_id", "body", m.DriveID); err != nil {
		return err
	}

	return nil
}

var driveTypeIoEnginePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sync","Async"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		driveTypeIoEnginePropEnum = append(driveTypeIoEnginePropEnum, v)
	}
}

const (

	// DriveIoEngineSync captures enum value "Sync"
	DriveIoEngineSync string = "Sync"

	// DriveIoEngineAsync captures enum value "Async"
	DriveIoEngineAsync string = "Async"
)

// prop value enum
func (m *Drive) validateIoEngineEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, driveTypeIoEnginePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Drive) validateIoEngine(formats strfmt.Registry) error {
	if swag.IsZero(m.IoEngine) { // not required
		return nil
	}

	// value enum
	if err := m.validateIoEngineEnum("io_engine", "body", *m.IoEngine); err != nil {
		return err
	}

	return nil
}

func (m *Drive) validateIsReadOnly(formats strfmt.Registry) error {

	if err := validate.Required("is_read_only", "body", m.IsReadOnly); err != nil {
		return err
	}

	return nil
}

func (m *Drive) validateIsRootDevice(formats strfmt.Registry) error {

	if err := validate.Required("is_root_device", "body", m.IsRootDevice); err != nil {
		return err
	}

	return nil
}

func (m *Drive) validatePathOnHost(formats strfmt.Registry) error {

	if err := validate.Required("path_on_host", "body", m.PathOnHost); err != nil {
		return err
	}

	return nil
}

func (m *Drive) validateRateLimiter(formats strfmt.Registry) error {
	if swag.IsZero(m.RateLimiter) { // not required
		return nil
	}

	if m.RateLimiter != nil {
		if err := m.RateLimiter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate_limiter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rate_limiter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this drive based on the context it is used
func (m *Drive) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRateLimiter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Drive) contextValidateRateLimiter(ctx context.Context, formats strfmt.Registry) error {

	if m.RateLimiter != nil {

		if swag.IsZero(m.RateLimiter) { // not required
			return nil
		}

		if err := m.RateLimiter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate_limiter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rate_limiter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Drive) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Drive) UnmarshalBinary(b []byte) error {
	var res Drive
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
