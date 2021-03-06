// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceAdministrative Administrative metadata for the resource.
// swagger:model resourceAdministrative
type ResourceAdministrative struct {

	// When the resource in TAQUITO was created.
	// Required: true
	Created *strfmt.DateTime `json:"created"`

	// If the resource has been deleted (but not purged).
	Deleted bool `json:"deleted,omitempty"`

	// Message describing why the object was deleted.
	GravestoneMessage string `json:"gravestoneMessage,omitempty"`

	// Pointer to the PURL/XML file that is a traditional representation of the metadata for the resource.
	// Required: true
	IsDescribedBy *string `json:"isDescribedBy"`

	// When the resource in TAQUITO was last updated.
	LastUpdated strfmt.DateTime `json:"lastUpdated,omitempty"`

	// Administrative or Internal project this resource is a part of.
	PartOfProject string `json:"partOfProject,omitempty"`

	// remediated by
	RemediatedBy ResourceAdministrativeRemediatedBy `json:"remediatedBy"`

	// If this resource should be sent to Preservation.
	// Required: true
	SdrPreserve *bool `json:"sdrPreserve"`
}

// Validate validates this resource administrative
func (m *ResourceAdministrative) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsDescribedBy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSdrPreserve(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceAdministrative) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResourceAdministrative) validateIsDescribedBy(formats strfmt.Registry) error {

	if err := validate.Required("isDescribedBy", "body", m.IsDescribedBy); err != nil {
		return err
	}

	return nil
}

func (m *ResourceAdministrative) validateSdrPreserve(formats strfmt.Registry) error {

	if err := validate.Required("sdrPreserve", "body", m.SdrPreserve); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceAdministrative) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceAdministrative) UnmarshalBinary(b []byte) error {
	var res ResourceAdministrative
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
