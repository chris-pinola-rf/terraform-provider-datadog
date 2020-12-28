/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsPrivateLocationSecretsConfigDecryption Private key for the private location.
type SyntheticsPrivateLocationSecretsConfigDecryption struct {
	// Private key for the private location.
	Key *string `json:"key,omitempty"`
}

// NewSyntheticsPrivateLocationSecretsConfigDecryption instantiates a new SyntheticsPrivateLocationSecretsConfigDecryption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsPrivateLocationSecretsConfigDecryption() *SyntheticsPrivateLocationSecretsConfigDecryption {
	this := SyntheticsPrivateLocationSecretsConfigDecryption{}
	return &this
}

// NewSyntheticsPrivateLocationSecretsConfigDecryptionWithDefaults instantiates a new SyntheticsPrivateLocationSecretsConfigDecryption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsPrivateLocationSecretsConfigDecryptionWithDefaults() *SyntheticsPrivateLocationSecretsConfigDecryption {
	this := SyntheticsPrivateLocationSecretsConfigDecryption{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SyntheticsPrivateLocationSecretsConfigDecryption) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPrivateLocationSecretsConfigDecryption) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SyntheticsPrivateLocationSecretsConfigDecryption) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SyntheticsPrivateLocationSecretsConfigDecryption) SetKey(v string) {
	o.Key = &v
}

func (o SyntheticsPrivateLocationSecretsConfigDecryption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsPrivateLocationSecretsConfigDecryption struct {
	value *SyntheticsPrivateLocationSecretsConfigDecryption
	isSet bool
}

func (v NullableSyntheticsPrivateLocationSecretsConfigDecryption) Get() *SyntheticsPrivateLocationSecretsConfigDecryption {
	return v.value
}

func (v *NullableSyntheticsPrivateLocationSecretsConfigDecryption) Set(val *SyntheticsPrivateLocationSecretsConfigDecryption) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsPrivateLocationSecretsConfigDecryption) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsPrivateLocationSecretsConfigDecryption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsPrivateLocationSecretsConfigDecryption(val *SyntheticsPrivateLocationSecretsConfigDecryption) *NullableSyntheticsPrivateLocationSecretsConfigDecryption {
	return &NullableSyntheticsPrivateLocationSecretsConfigDecryption{value: val, isSet: true}
}

func (v NullableSyntheticsPrivateLocationSecretsConfigDecryption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsPrivateLocationSecretsConfigDecryption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}