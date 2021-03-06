// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// EnvFromSource EnvFromSource represents the source of a set of ConfigMaps
//
// swagger:model EnvFromSource
type EnvFromSource struct {

	// The ConfigMap to select from
	ConfigMapRef *ConfigMapEnvSource `json:"configMapRef,omitempty"`

	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.
	Prefix string `json:"prefix,omitempty"`

	// The Secret to select from
	SecretRef *SecretEnvSource `json:"secretRef,omitempty"`
}
