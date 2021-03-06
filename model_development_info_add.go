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

// DevelopmentInfoAdd struct for DevelopmentInfoAdd
type DevelopmentInfoAdd struct {
	Id        *int64  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	Group     *string `json:"group,omitempty"`
}

// NewDevelopmentInfoAdd instantiates a new DevelopmentInfoAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevelopmentInfoAdd() *DevelopmentInfoAdd {
	this := DevelopmentInfoAdd{}
	return &this
}

// NewDevelopmentInfoAddWithDefaults instantiates a new DevelopmentInfoAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevelopmentInfoAddWithDefaults() *DevelopmentInfoAdd {
	this := DevelopmentInfoAdd{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DevelopmentInfoAdd) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentInfoAdd) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DevelopmentInfoAdd) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *DevelopmentInfoAdd) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DevelopmentInfoAdd) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentInfoAdd) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DevelopmentInfoAdd) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DevelopmentInfoAdd) SetName(v string) {
	o.Name = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *DevelopmentInfoAdd) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentInfoAdd) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *DevelopmentInfoAdd) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *DevelopmentInfoAdd) SetShortName(v string) {
	o.ShortName = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *DevelopmentInfoAdd) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevelopmentInfoAdd) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *DevelopmentInfoAdd) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *DevelopmentInfoAdd) SetGroup(v string) {
	o.Group = &v
}

func (o DevelopmentInfoAdd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	return json.Marshal(toSerialize)
}

type NullableDevelopmentInfoAdd struct {
	value *DevelopmentInfoAdd
	isSet bool
}

func (v NullableDevelopmentInfoAdd) Get() *DevelopmentInfoAdd {
	return v.value
}

func (v *NullableDevelopmentInfoAdd) Set(val *DevelopmentInfoAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableDevelopmentInfoAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableDevelopmentInfoAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevelopmentInfoAdd(val *DevelopmentInfoAdd) *NullableDevelopmentInfoAdd {
	return &NullableDevelopmentInfoAdd{value: val, isSet: true}
}

func (v NullableDevelopmentInfoAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevelopmentInfoAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
