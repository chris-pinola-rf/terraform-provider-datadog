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

// UsageLambdaResponse Reponse containing the umber of lambda functions and sum of the invocations of all lambda functions for each hour for a given organization.
type UsageLambdaResponse struct {
	// Get hourly usage for Lambda.
	Usage *[]UsageLambdaHour `json:"usage,omitempty"`
}

// NewUsageLambdaResponse instantiates a new UsageLambdaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLambdaResponse() *UsageLambdaResponse {
	this := UsageLambdaResponse{}
	return &this
}

// NewUsageLambdaResponseWithDefaults instantiates a new UsageLambdaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLambdaResponseWithDefaults() *UsageLambdaResponse {
	this := UsageLambdaResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageLambdaResponse) GetUsage() []UsageLambdaHour {
	if o == nil || o.Usage == nil {
		var ret []UsageLambdaHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLambdaResponse) GetUsageOk() (*[]UsageLambdaHour, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageLambdaResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageLambdaHour and assigns it to the Usage field.
func (o *UsageLambdaResponse) SetUsage(v []UsageLambdaHour) {
	o.Usage = &v
}

func (o UsageLambdaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageLambdaResponse struct {
	value *UsageLambdaResponse
	isSet bool
}

func (v NullableUsageLambdaResponse) Get() *UsageLambdaResponse {
	return v.value
}

func (v *NullableUsageLambdaResponse) Set(val *UsageLambdaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLambdaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLambdaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLambdaResponse(val *UsageLambdaResponse) *NullableUsageLambdaResponse {
	return &NullableUsageLambdaResponse{value: val, isSet: true}
}

func (v NullableUsageLambdaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLambdaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
