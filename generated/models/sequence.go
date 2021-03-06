// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Sequence Resource Sequence
//
// A sequence or ordering of resources within a Collection or Object.
// swagger:model Sequence
type Sequence struct {

	// URI for the JSON-LD context definitions.
	AtContext string `json:"@context,omitempty"`

	// The type of Sequence.
	// Required: true
	AtType *string `json:"@type"`

	// Label for the sequence or ordering.
	// Required: true
	Label *string `json:"label"`

	// members
	// Required: true
	Members *SequenceMembersTuple0 `json:"members"`

	// Identifier for the first member of the sequence.
	// Required: true
	StartMember *string `json:"startMember"`
}

// Validate validates this sequence
func (m *Sequence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAtType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartMember(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sequenceTypeAtTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http://sdr.sul.stanford.edu/models/sdr3-#/definitions/Sequenceld","http://sdr.sul.stanford.edu/models/sdr3-primary-sequence.jsonld"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sequenceTypeAtTypePropEnum = append(sequenceTypeAtTypePropEnum, v)
	}
}

const (
	// SequenceAtTypeHTTPSdrSulStanfordEduModelsSdr3DefinitionsSequenceld captures enum value "http://sdr.sul.stanford.edu/models/sdr3-#/definitions/Sequenceld"
	SequenceAtTypeHTTPSdrSulStanfordEduModelsSdr3DefinitionsSequenceld string = "http://sdr.sul.stanford.edu/models/sdr3-#/definitions/Sequenceld"
	// SequenceAtTypeHTTPSdrSulStanfordEduModelsSdr3PrimarySequenceJSONLD captures enum value "http://sdr.sul.stanford.edu/models/sdr3-primary-sequence.jsonld"
	SequenceAtTypeHTTPSdrSulStanfordEduModelsSdr3PrimarySequenceJSONLD string = "http://sdr.sul.stanford.edu/models/sdr3-primary-sequence.jsonld"
)

// prop value enum
func (m *Sequence) validateAtTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sequenceTypeAtTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Sequence) validateAtType(formats strfmt.Registry) error {

	if err := validate.Required("@type", "body", m.AtType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAtTypeEnum("@type", "body", *m.AtType); err != nil {
		return err
	}

	return nil
}

func (m *Sequence) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *Sequence) validateMembers(formats strfmt.Registry) error {

	if err := validate.Required("members", "body", m.Members); err != nil {
		return err
	}

	if m.Members != nil {

		if err := m.Members.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("members")
			}
			return err
		}
	}

	return nil
}

func (m *Sequence) validateStartMember(formats strfmt.Registry) error {

	if err := validate.Required("startMember", "body", m.StartMember); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Sequence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sequence) UnmarshalBinary(b []byte) error {
	var res Sequence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SequenceMembersTuple0 SequenceMembersTuple0 a representation of an anonymous Tuple type
// swagger:model SequenceMembersTuple0
type SequenceMembersTuple0 struct {

	// p0
	// Required: true
	P0 *string `json:"-"` // custom serializer

}

// UnmarshalJSON unmarshals this tuple type from a JSON array
func (m *SequenceMembersTuple0) UnmarshalJSON(raw []byte) error {
	// stage 1, get the array but just the array
	var stage1 []json.RawMessage
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&stage1); err != nil {
		return err
	}

	// stage 2
	if len(stage1) > 0 {
		buf = bytes.NewBuffer(stage1[0])
		dec := json.NewDecoder(buf)
		dec.UseNumber()
		if err := dec.Decode(m.P0); err != nil {
			return err
		}
	}

	return nil
}

// MarshalJSON marshals this tuple type into a JSON array
func (m SequenceMembersTuple0) MarshalJSON() ([]byte, error) {
	data := []interface{}{
		m.P0,
	}

	return json.Marshal(data)
}

// Validate validates this sequence members tuple0
func (m *SequenceMembersTuple0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateP0(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SequenceMembersTuple0) validateP0(formats strfmt.Registry) error {

	if err := validate.Required("P0", "body", m.P0); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SequenceMembersTuple0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SequenceMembersTuple0) UnmarshalBinary(b []byte) error {
	var res SequenceMembersTuple0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
