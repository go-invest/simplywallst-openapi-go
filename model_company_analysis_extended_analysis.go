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

// CompanyAnalysisExtendedAnalysis struct for CompanyAnalysisExtendedAnalysis
type CompanyAnalysisExtendedAnalysis struct {
	Dividend *CompanyAnalysisExtendedAnalysisDividend `json:"dividend,omitempty"`
	Past     *CompanyAnalysisExtendedAnalysisPast     `json:"past,omitempty"`
	Future   *CompanyAnalysisExtendedAnalysisFuture   `json:"future,omitempty"`
	Health   *CompanyAnalysisExtendedAnalysisHealth   `json:"health,omitempty"`
	Misc     *CompanyAnalysisExtendedAnalysisMisc     `json:"misc,omitempty"`
	Value    *CompanyAnalysisExtendedAnalysisValue    `json:"value,omitempty"`
}

// NewCompanyAnalysisExtendedAnalysis instantiates a new CompanyAnalysisExtendedAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyAnalysisExtendedAnalysis() *CompanyAnalysisExtendedAnalysis {
	this := CompanyAnalysisExtendedAnalysis{}
	return &this
}

// NewCompanyAnalysisExtendedAnalysisWithDefaults instantiates a new CompanyAnalysisExtendedAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyAnalysisExtendedAnalysisWithDefaults() *CompanyAnalysisExtendedAnalysis {
	this := CompanyAnalysisExtendedAnalysis{}
	return &this
}

// GetDividend returns the Dividend field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedAnalysis) GetDividend() CompanyAnalysisExtendedAnalysisDividend {
	if o == nil || o.Dividend == nil {
		var ret CompanyAnalysisExtendedAnalysisDividend
		return ret
	}
	return *o.Dividend
}

// GetDividendOk returns a tuple with the Dividend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedAnalysis) GetDividendOk() (*CompanyAnalysisExtendedAnalysisDividend, bool) {
	if o == nil || o.Dividend == nil {
		return nil, false
	}
	return o.Dividend, true
}

// HasDividend returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedAnalysis) HasDividend() bool {
	if o != nil && o.Dividend != nil {
		return true
	}

	return false
}

// SetDividend gets a reference to the given CompanyAnalysisExtendedAnalysisDividend and assigns it to the Dividend field.
func (o *CompanyAnalysisExtendedAnalysis) SetDividend(v CompanyAnalysisExtendedAnalysisDividend) {
	o.Dividend = &v
}

// GetPast returns the Past field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedAnalysis) GetPast() CompanyAnalysisExtendedAnalysisPast {
	if o == nil || o.Past == nil {
		var ret CompanyAnalysisExtendedAnalysisPast
		return ret
	}
	return *o.Past
}

// GetPastOk returns a tuple with the Past field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedAnalysis) GetPastOk() (*CompanyAnalysisExtendedAnalysisPast, bool) {
	if o == nil || o.Past == nil {
		return nil, false
	}
	return o.Past, true
}

// HasPast returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedAnalysis) HasPast() bool {
	if o != nil && o.Past != nil {
		return true
	}

	return false
}

// SetPast gets a reference to the given CompanyAnalysisExtendedAnalysisPast and assigns it to the Past field.
func (o *CompanyAnalysisExtendedAnalysis) SetPast(v CompanyAnalysisExtendedAnalysisPast) {
	o.Past = &v
}

// GetFuture returns the Future field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedAnalysis) GetFuture() CompanyAnalysisExtendedAnalysisFuture {
	if o == nil || o.Future == nil {
		var ret CompanyAnalysisExtendedAnalysisFuture
		return ret
	}
	return *o.Future
}

// GetFutureOk returns a tuple with the Future field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedAnalysis) GetFutureOk() (*CompanyAnalysisExtendedAnalysisFuture, bool) {
	if o == nil || o.Future == nil {
		return nil, false
	}
	return o.Future, true
}

// HasFuture returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedAnalysis) HasFuture() bool {
	if o != nil && o.Future != nil {
		return true
	}

	return false
}

// SetFuture gets a reference to the given CompanyAnalysisExtendedAnalysisFuture and assigns it to the Future field.
func (o *CompanyAnalysisExtendedAnalysis) SetFuture(v CompanyAnalysisExtendedAnalysisFuture) {
	o.Future = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedAnalysis) GetHealth() CompanyAnalysisExtendedAnalysisHealth {
	if o == nil || o.Health == nil {
		var ret CompanyAnalysisExtendedAnalysisHealth
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedAnalysis) GetHealthOk() (*CompanyAnalysisExtendedAnalysisHealth, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedAnalysis) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given CompanyAnalysisExtendedAnalysisHealth and assigns it to the Health field.
func (o *CompanyAnalysisExtendedAnalysis) SetHealth(v CompanyAnalysisExtendedAnalysisHealth) {
	o.Health = &v
}

// GetMisc returns the Misc field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedAnalysis) GetMisc() CompanyAnalysisExtendedAnalysisMisc {
	if o == nil || o.Misc == nil {
		var ret CompanyAnalysisExtendedAnalysisMisc
		return ret
	}
	return *o.Misc
}

// GetMiscOk returns a tuple with the Misc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedAnalysis) GetMiscOk() (*CompanyAnalysisExtendedAnalysisMisc, bool) {
	if o == nil || o.Misc == nil {
		return nil, false
	}
	return o.Misc, true
}

// HasMisc returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedAnalysis) HasMisc() bool {
	if o != nil && o.Misc != nil {
		return true
	}

	return false
}

// SetMisc gets a reference to the given CompanyAnalysisExtendedAnalysisMisc and assigns it to the Misc field.
func (o *CompanyAnalysisExtendedAnalysis) SetMisc(v CompanyAnalysisExtendedAnalysisMisc) {
	o.Misc = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedAnalysis) GetValue() CompanyAnalysisExtendedAnalysisValue {
	if o == nil || o.Value == nil {
		var ret CompanyAnalysisExtendedAnalysisValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedAnalysis) GetValueOk() (*CompanyAnalysisExtendedAnalysisValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedAnalysis) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given CompanyAnalysisExtendedAnalysisValue and assigns it to the Value field.
func (o *CompanyAnalysisExtendedAnalysis) SetValue(v CompanyAnalysisExtendedAnalysisValue) {
	o.Value = &v
}

func (o CompanyAnalysisExtendedAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dividend != nil {
		toSerialize["dividend"] = o.Dividend
	}
	if o.Past != nil {
		toSerialize["past"] = o.Past
	}
	if o.Future != nil {
		toSerialize["future"] = o.Future
	}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	if o.Misc != nil {
		toSerialize["misc"] = o.Misc
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyAnalysisExtendedAnalysis struct {
	value *CompanyAnalysisExtendedAnalysis
	isSet bool
}

func (v NullableCompanyAnalysisExtendedAnalysis) Get() *CompanyAnalysisExtendedAnalysis {
	return v.value
}

func (v *NullableCompanyAnalysisExtendedAnalysis) Set(val *CompanyAnalysisExtendedAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyAnalysisExtendedAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyAnalysisExtendedAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyAnalysisExtendedAnalysis(val *CompanyAnalysisExtendedAnalysis) *NullableCompanyAnalysisExtendedAnalysis {
	return &NullableCompanyAnalysisExtendedAnalysis{value: val, isSet: true}
}

func (v NullableCompanyAnalysisExtendedAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyAnalysisExtendedAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
