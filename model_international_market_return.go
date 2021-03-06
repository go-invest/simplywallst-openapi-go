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

// InternationalMarketReturn struct for InternationalMarketReturn
type InternationalMarketReturn struct {
	Return7d     *[]InternationalMarketReturnDate `json:"return_7d,omitempty"`
	Return30d    *[]InternationalMarketReturnDate `json:"return_30d,omitempty"`
	Return90d    *[]InternationalMarketReturnDate `json:"return_90d,omitempty"`
	Return1yrAbs *[]InternationalMarketReturnDate `json:"return_1yr_abs,omitempty"`
	Return3yrAbs *[]InternationalMarketReturnDate `json:"return_3yr_abs,omitempty"`
	Return5yrAbs *[]InternationalMarketReturnDate `json:"return_5yr_abs,omitempty"`
}

// NewInternationalMarketReturn instantiates a new InternationalMarketReturn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternationalMarketReturn() *InternationalMarketReturn {
	this := InternationalMarketReturn{}
	return &this
}

// NewInternationalMarketReturnWithDefaults instantiates a new InternationalMarketReturn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternationalMarketReturnWithDefaults() *InternationalMarketReturn {
	this := InternationalMarketReturn{}
	return &this
}

// GetReturn7d returns the Return7d field value if set, zero value otherwise.
func (o *InternationalMarketReturn) GetReturn7d() []InternationalMarketReturnDate {
	if o == nil || o.Return7d == nil {
		var ret []InternationalMarketReturnDate
		return ret
	}
	return *o.Return7d
}

// GetReturn7dOk returns a tuple with the Return7d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalMarketReturn) GetReturn7dOk() (*[]InternationalMarketReturnDate, bool) {
	if o == nil || o.Return7d == nil {
		return nil, false
	}
	return o.Return7d, true
}

// HasReturn7d returns a boolean if a field has been set.
func (o *InternationalMarketReturn) HasReturn7d() bool {
	if o != nil && o.Return7d != nil {
		return true
	}

	return false
}

// SetReturn7d gets a reference to the given []InternationalMarketReturnDate and assigns it to the Return7d field.
func (o *InternationalMarketReturn) SetReturn7d(v []InternationalMarketReturnDate) {
	o.Return7d = &v
}

// GetReturn30d returns the Return30d field value if set, zero value otherwise.
func (o *InternationalMarketReturn) GetReturn30d() []InternationalMarketReturnDate {
	if o == nil || o.Return30d == nil {
		var ret []InternationalMarketReturnDate
		return ret
	}
	return *o.Return30d
}

// GetReturn30dOk returns a tuple with the Return30d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalMarketReturn) GetReturn30dOk() (*[]InternationalMarketReturnDate, bool) {
	if o == nil || o.Return30d == nil {
		return nil, false
	}
	return o.Return30d, true
}

// HasReturn30d returns a boolean if a field has been set.
func (o *InternationalMarketReturn) HasReturn30d() bool {
	if o != nil && o.Return30d != nil {
		return true
	}

	return false
}

// SetReturn30d gets a reference to the given []InternationalMarketReturnDate and assigns it to the Return30d field.
func (o *InternationalMarketReturn) SetReturn30d(v []InternationalMarketReturnDate) {
	o.Return30d = &v
}

// GetReturn90d returns the Return90d field value if set, zero value otherwise.
func (o *InternationalMarketReturn) GetReturn90d() []InternationalMarketReturnDate {
	if o == nil || o.Return90d == nil {
		var ret []InternationalMarketReturnDate
		return ret
	}
	return *o.Return90d
}

// GetReturn90dOk returns a tuple with the Return90d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalMarketReturn) GetReturn90dOk() (*[]InternationalMarketReturnDate, bool) {
	if o == nil || o.Return90d == nil {
		return nil, false
	}
	return o.Return90d, true
}

// HasReturn90d returns a boolean if a field has been set.
func (o *InternationalMarketReturn) HasReturn90d() bool {
	if o != nil && o.Return90d != nil {
		return true
	}

	return false
}

// SetReturn90d gets a reference to the given []InternationalMarketReturnDate and assigns it to the Return90d field.
func (o *InternationalMarketReturn) SetReturn90d(v []InternationalMarketReturnDate) {
	o.Return90d = &v
}

// GetReturn1yrAbs returns the Return1yrAbs field value if set, zero value otherwise.
func (o *InternationalMarketReturn) GetReturn1yrAbs() []InternationalMarketReturnDate {
	if o == nil || o.Return1yrAbs == nil {
		var ret []InternationalMarketReturnDate
		return ret
	}
	return *o.Return1yrAbs
}

// GetReturn1yrAbsOk returns a tuple with the Return1yrAbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalMarketReturn) GetReturn1yrAbsOk() (*[]InternationalMarketReturnDate, bool) {
	if o == nil || o.Return1yrAbs == nil {
		return nil, false
	}
	return o.Return1yrAbs, true
}

// HasReturn1yrAbs returns a boolean if a field has been set.
func (o *InternationalMarketReturn) HasReturn1yrAbs() bool {
	if o != nil && o.Return1yrAbs != nil {
		return true
	}

	return false
}

// SetReturn1yrAbs gets a reference to the given []InternationalMarketReturnDate and assigns it to the Return1yrAbs field.
func (o *InternationalMarketReturn) SetReturn1yrAbs(v []InternationalMarketReturnDate) {
	o.Return1yrAbs = &v
}

// GetReturn3yrAbs returns the Return3yrAbs field value if set, zero value otherwise.
func (o *InternationalMarketReturn) GetReturn3yrAbs() []InternationalMarketReturnDate {
	if o == nil || o.Return3yrAbs == nil {
		var ret []InternationalMarketReturnDate
		return ret
	}
	return *o.Return3yrAbs
}

// GetReturn3yrAbsOk returns a tuple with the Return3yrAbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalMarketReturn) GetReturn3yrAbsOk() (*[]InternationalMarketReturnDate, bool) {
	if o == nil || o.Return3yrAbs == nil {
		return nil, false
	}
	return o.Return3yrAbs, true
}

// HasReturn3yrAbs returns a boolean if a field has been set.
func (o *InternationalMarketReturn) HasReturn3yrAbs() bool {
	if o != nil && o.Return3yrAbs != nil {
		return true
	}

	return false
}

// SetReturn3yrAbs gets a reference to the given []InternationalMarketReturnDate and assigns it to the Return3yrAbs field.
func (o *InternationalMarketReturn) SetReturn3yrAbs(v []InternationalMarketReturnDate) {
	o.Return3yrAbs = &v
}

// GetReturn5yrAbs returns the Return5yrAbs field value if set, zero value otherwise.
func (o *InternationalMarketReturn) GetReturn5yrAbs() []InternationalMarketReturnDate {
	if o == nil || o.Return5yrAbs == nil {
		var ret []InternationalMarketReturnDate
		return ret
	}
	return *o.Return5yrAbs
}

// GetReturn5yrAbsOk returns a tuple with the Return5yrAbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternationalMarketReturn) GetReturn5yrAbsOk() (*[]InternationalMarketReturnDate, bool) {
	if o == nil || o.Return5yrAbs == nil {
		return nil, false
	}
	return o.Return5yrAbs, true
}

// HasReturn5yrAbs returns a boolean if a field has been set.
func (o *InternationalMarketReturn) HasReturn5yrAbs() bool {
	if o != nil && o.Return5yrAbs != nil {
		return true
	}

	return false
}

// SetReturn5yrAbs gets a reference to the given []InternationalMarketReturnDate and assigns it to the Return5yrAbs field.
func (o *InternationalMarketReturn) SetReturn5yrAbs(v []InternationalMarketReturnDate) {
	o.Return5yrAbs = &v
}

func (o InternationalMarketReturn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Return7d != nil {
		toSerialize["return_7d"] = o.Return7d
	}
	if o.Return30d != nil {
		toSerialize["return_30d"] = o.Return30d
	}
	if o.Return90d != nil {
		toSerialize["return_90d"] = o.Return90d
	}
	if o.Return1yrAbs != nil {
		toSerialize["return_1yr_abs"] = o.Return1yrAbs
	}
	if o.Return3yrAbs != nil {
		toSerialize["return_3yr_abs"] = o.Return3yrAbs
	}
	if o.Return5yrAbs != nil {
		toSerialize["return_5yr_abs"] = o.Return5yrAbs
	}
	return json.Marshal(toSerialize)
}

type NullableInternationalMarketReturn struct {
	value *InternationalMarketReturn
	isSet bool
}

func (v NullableInternationalMarketReturn) Get() *InternationalMarketReturn {
	return v.value
}

func (v *NullableInternationalMarketReturn) Set(val *InternationalMarketReturn) {
	v.value = val
	v.isSet = true
}

func (v NullableInternationalMarketReturn) IsSet() bool {
	return v.isSet
}

func (v *NullableInternationalMarketReturn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternationalMarketReturn(val *InternationalMarketReturn) *NullableInternationalMarketReturn {
	return &NullableInternationalMarketReturn{value: val, isSet: true}
}

func (v NullableInternationalMarketReturn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternationalMarketReturn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
