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

// Prices struct for Prices
type Prices struct {
	Data *[]Price `json:"data,omitempty"`
}

// NewPrices instantiates a new Prices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrices() *Prices {
	this := Prices{}
	return &this
}

// NewPricesWithDefaults instantiates a new Prices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPricesWithDefaults() *Prices {
	this := Prices{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Prices) GetData() []Price {
	if o == nil || o.Data == nil {
		var ret []Price
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prices) GetDataOk() (*[]Price, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Prices) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Price and assigns it to the Data field.
func (o *Prices) SetData(v []Price) {
	o.Data = &v
}

func (o Prices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullablePrices struct {
	value *Prices
	isSet bool
}

func (v NullablePrices) Get() *Prices {
	return v.value
}

func (v *NullablePrices) Set(val *Prices) {
	v.value = val
	v.isSet = true
}

func (v NullablePrices) IsSet() bool {
	return v.isSet
}

func (v *NullablePrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrices(val *Prices) *NullablePrices {
	return &NullablePrices{value: val, isSet: true}
}

func (v NullablePrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
