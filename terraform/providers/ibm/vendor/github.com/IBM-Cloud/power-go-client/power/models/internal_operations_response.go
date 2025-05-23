// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InternalOperationsResponse internal operations response
//
// swagger:model InternalOperationsResponse
type InternalOperationsResponse struct {

	// crn
	Crn CRN `json:"crn,omitempty"`
}

// Validate validates this internal operations response
func (m *InternalOperationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternalOperationsResponse) validateCrn(formats strfmt.Registry) error {
	if swag.IsZero(m.Crn) { // not required
		return nil
	}

	if err := m.Crn.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("crn")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("crn")
		}
		return err
	}

	return nil
}

// ContextValidate validate this internal operations response based on the context it is used
func (m *InternalOperationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCrn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternalOperationsResponse) contextValidateCrn(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Crn) { // not required
		return nil
	}

	if err := m.Crn.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("crn")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("crn")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InternalOperationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InternalOperationsResponse) UnmarshalBinary(b []byte) error {
	var res InternalOperationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
