/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EstimateCoverage struct for EstimateCoverage
type EstimateCoverage struct {
	CompanyId *int64                   `json:"company_id,omitempty"`
	Brokers   *EstimateCoverageBrokers `json:"brokers,omitempty"`
}

// NewEstimateCoverage instantiates a new EstimateCoverage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimateCoverage() *EstimateCoverage {
	this := EstimateCoverage{}
	return &this
}

// NewEstimateCoverageWithDefaults instantiates a new EstimateCoverage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimateCoverageWithDefaults() *EstimateCoverage {
	this := EstimateCoverage{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *EstimateCoverage) GetCompanyId() int64 {
	if o == nil || o.CompanyId == nil {
		var ret int64
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateCoverage) GetCompanyIdOk() (*int64, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *EstimateCoverage) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int64 and assigns it to the CompanyId field.
func (o *EstimateCoverage) SetCompanyId(v int64) {
	o.CompanyId = &v
}

// GetBrokers returns the Brokers field value if set, zero value otherwise.
func (o *EstimateCoverage) GetBrokers() EstimateCoverageBrokers {
	if o == nil || o.Brokers == nil {
		var ret EstimateCoverageBrokers
		return ret
	}
	return *o.Brokers
}

// GetBrokersOk returns a tuple with the Brokers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimateCoverage) GetBrokersOk() (*EstimateCoverageBrokers, bool) {
	if o == nil || o.Brokers == nil {
		return nil, false
	}
	return o.Brokers, true
}

// HasBrokers returns a boolean if a field has been set.
func (o *EstimateCoverage) HasBrokers() bool {
	if o != nil && o.Brokers != nil {
		return true
	}

	return false
}

// SetBrokers gets a reference to the given EstimateCoverageBrokers and assigns it to the Brokers field.
func (o *EstimateCoverage) SetBrokers(v EstimateCoverageBrokers) {
	o.Brokers = &v
}

func (o EstimateCoverage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["company_id"] = o.CompanyId
	}
	if o.Brokers != nil {
		toSerialize["brokers"] = o.Brokers
	}
	return json.Marshal(toSerialize)
}

type NullableEstimateCoverage struct {
	value *EstimateCoverage
	isSet bool
}

func (v NullableEstimateCoverage) Get() *EstimateCoverage {
	return v.value
}

func (v *NullableEstimateCoverage) Set(val *EstimateCoverage) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimateCoverage) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimateCoverage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimateCoverage(val *EstimateCoverage) *NullableEstimateCoverage {
	return &NullableEstimateCoverage{value: val, isSet: true}
}

func (v NullableEstimateCoverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimateCoverage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
