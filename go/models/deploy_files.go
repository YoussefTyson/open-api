// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeployFiles deploy files
//
// swagger:model deployFiles
type DeployFiles struct {

	// async
	Async bool `json:"async,omitempty"`

	// await ready signal
	AwaitReadySignal bool `json:"await_ready_signal,omitempty"`

	// branch
	Branch string `json:"branch,omitempty"`

	// draft
	Draft bool `json:"draft,omitempty"`

	// files
	Files interface{} `json:"files,omitempty"`

	// framework
	Framework string `json:"framework,omitempty"`

	// function schedules
	FunctionSchedules []*FunctionSchedule `json:"function_schedules"`

	// functions
	Functions interface{} `json:"functions,omitempty"`

	// functions config
	FunctionsConfig map[string]FunctionConfig `json:"functions_config,omitempty"`
}

// Validate validates this deploy files
func (m *DeployFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFunctionSchedules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctionsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeployFiles) validateFunctionSchedules(formats strfmt.Registry) error {

	if swag.IsZero(m.FunctionSchedules) { // not required
		return nil
	}

	for i := 0; i < len(m.FunctionSchedules); i++ {
		if swag.IsZero(m.FunctionSchedules[i]) { // not required
			continue
		}

		if m.FunctionSchedules[i] != nil {
			if err := m.FunctionSchedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("function_schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DeployFiles) validateFunctionsConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FunctionsConfig) { // not required
		return nil
	}

	for k := range m.FunctionsConfig {

		if err := validate.Required("functions_config"+"."+k, "body", m.FunctionsConfig[k]); err != nil {
			return err
		}
		if val, ok := m.FunctionsConfig[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeployFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeployFiles) UnmarshalBinary(b []byte) error {
	var res DeployFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
