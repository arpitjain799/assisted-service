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

// KernelArgument pair of [operation, argument] specifying the argument and what operation should be applied on it.
//
// swagger:model kernel_argument
type KernelArgument struct {

	// The operation to apply on the kernel argument.
	// Enum: [append replace delete]
	Operation string `json:"operation,omitempty"`

	// Kernel argument can have the form <parameter> or <parameter>=<value>. The following examples should
	// be supported:
	// rd.net.timeout.carrier=60
	// isolcpus=1,2,10-20,100-2000:2/25
	// quiet
	// The parsing by the command line parser in linux kernel is much looser and this pattern follows it.
	//
	// Pattern: ^(?:(?:[^ \t\n\r"]+)|(?:"[^"]*"))+$
	Value string `json:"value,omitempty"`
}

// Validate validates this kernel argument
func (m *KernelArgument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kernelArgumentTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["append","replace","delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kernelArgumentTypeOperationPropEnum = append(kernelArgumentTypeOperationPropEnum, v)
	}
}

const (

	// KernelArgumentOperationAppend captures enum value "append"
	KernelArgumentOperationAppend string = "append"

	// KernelArgumentOperationReplace captures enum value "replace"
	KernelArgumentOperationReplace string = "replace"

	// KernelArgumentOperationDelete captures enum value "delete"
	KernelArgumentOperationDelete string = "delete"
)

// prop value enum
func (m *KernelArgument) validateOperationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kernelArgumentTypeOperationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KernelArgument) validateOperation(formats strfmt.Registry) error {
	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *KernelArgument) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.Pattern("value", "body", m.Value, `^(?:(?:[^ \t\n\r"]+)|(?:"[^"]*"))+$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this kernel argument based on context it is used
func (m *KernelArgument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KernelArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KernelArgument) UnmarshalBinary(b []byte) error {
	var res KernelArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}