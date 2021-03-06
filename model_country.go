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

// Country struct for Country
type Country struct {
	Name   *string        `json:"name,omitempty"`
	Iso2   *string        `json:"iso2,omitempty"`
	Iso3   *string        `json:"iso3,omitempty"`
	Type   *string        `json:"type,omitempty"`
	Links  *CountryLinks  `json:"links,omitempty"`
	Fields *CountryFields `json:"fields,omitempty"`
}

// NewCountry instantiates a new Country object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountry() *Country {
	this := Country{}
	return &this
}

// NewCountryWithDefaults instantiates a new Country object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryWithDefaults() *Country {
	this := Country{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Country) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Country) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Country) SetName(v string) {
	o.Name = &v
}

// GetIso2 returns the Iso2 field value if set, zero value otherwise.
func (o *Country) GetIso2() string {
	if o == nil || o.Iso2 == nil {
		var ret string
		return ret
	}
	return *o.Iso2
}

// GetIso2Ok returns a tuple with the Iso2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetIso2Ok() (*string, bool) {
	if o == nil || o.Iso2 == nil {
		return nil, false
	}
	return o.Iso2, true
}

// HasIso2 returns a boolean if a field has been set.
func (o *Country) HasIso2() bool {
	if o != nil && o.Iso2 != nil {
		return true
	}

	return false
}

// SetIso2 gets a reference to the given string and assigns it to the Iso2 field.
func (o *Country) SetIso2(v string) {
	o.Iso2 = &v
}

// GetIso3 returns the Iso3 field value if set, zero value otherwise.
func (o *Country) GetIso3() string {
	if o == nil || o.Iso3 == nil {
		var ret string
		return ret
	}
	return *o.Iso3
}

// GetIso3Ok returns a tuple with the Iso3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetIso3Ok() (*string, bool) {
	if o == nil || o.Iso3 == nil {
		return nil, false
	}
	return o.Iso3, true
}

// HasIso3 returns a boolean if a field has been set.
func (o *Country) HasIso3() bool {
	if o != nil && o.Iso3 != nil {
		return true
	}

	return false
}

// SetIso3 gets a reference to the given string and assigns it to the Iso3 field.
func (o *Country) SetIso3(v string) {
	o.Iso3 = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Country) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Country) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Country) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Country) GetLinks() CountryLinks {
	if o == nil || o.Links == nil {
		var ret CountryLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetLinksOk() (*CountryLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Country) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given CountryLinks and assigns it to the Links field.
func (o *Country) SetLinks(v CountryLinks) {
	o.Links = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *Country) GetFields() CountryFields {
	if o == nil || o.Fields == nil {
		var ret CountryFields
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetFieldsOk() (*CountryFields, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *Country) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given CountryFields and assigns it to the Fields field.
func (o *Country) SetFields(v CountryFields) {
	o.Fields = &v
}

func (o Country) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Iso2 != nil {
		toSerialize["iso2"] = o.Iso2
	}
	if o.Iso3 != nil {
		toSerialize["iso3"] = o.Iso3
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	return json.Marshal(toSerialize)
}

type NullableCountry struct {
	value *Country
	isSet bool
}

func (v NullableCountry) Get() *Country {
	return v.value
}

func (v *NullableCountry) Set(val *Country) {
	v.value = val
	v.isSet = true
}

func (v NullableCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountry(val *Country) *NullableCountry {
	return &NullableCountry{value: val, isSet: true}
}

func (v NullableCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
