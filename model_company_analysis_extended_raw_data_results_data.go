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

// CompanyAnalysisExtendedRawDataResultsData struct for CompanyAnalysisExtendedRawDataResultsData
type CompanyAnalysisExtendedRawDataResultsData struct {
	Quarter     *int64   `json:"quarter,omitempty"`
	Year        *int64   `json:"year,omitempty"`
	Latest      *bool    `json:"latest,omitempty"`
	Id          *string  `json:"id,omitempty"`
	Value       *float64 `json:"value,omitempty"`
	EndDate     *int64   `json:"end_date,omitempty"`
	Unit        *string  `json:"unit,omitempty"`
	UnitNumeric *float64 `json:"unit_numeric,omitempty"`
	Currency    *string  `json:"currency,omitempty"`
}

// NewCompanyAnalysisExtendedRawDataResultsData instantiates a new CompanyAnalysisExtendedRawDataResultsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyAnalysisExtendedRawDataResultsData() *CompanyAnalysisExtendedRawDataResultsData {
	this := CompanyAnalysisExtendedRawDataResultsData{}
	return &this
}

// NewCompanyAnalysisExtendedRawDataResultsDataWithDefaults instantiates a new CompanyAnalysisExtendedRawDataResultsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyAnalysisExtendedRawDataResultsDataWithDefaults() *CompanyAnalysisExtendedRawDataResultsData {
	this := CompanyAnalysisExtendedRawDataResultsData{}
	return &this
}

// GetQuarter returns the Quarter field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetQuarter() int64 {
	if o == nil || o.Quarter == nil {
		var ret int64
		return ret
	}
	return *o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetQuarterOk() (*int64, bool) {
	if o == nil || o.Quarter == nil {
		return nil, false
	}
	return o.Quarter, true
}

// HasQuarter returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasQuarter() bool {
	if o != nil && o.Quarter != nil {
		return true
	}

	return false
}

// SetQuarter gets a reference to the given int64 and assigns it to the Quarter field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetQuarter(v int64) {
	o.Quarter = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetYear() int64 {
	if o == nil || o.Year == nil {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetYearOk() (*int64, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetYear(v int64) {
	o.Year = &v
}

// GetLatest returns the Latest field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetLatest() bool {
	if o == nil || o.Latest == nil {
		var ret bool
		return ret
	}
	return *o.Latest
}

// GetLatestOk returns a tuple with the Latest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetLatestOk() (*bool, bool) {
	if o == nil || o.Latest == nil {
		return nil, false
	}
	return o.Latest, true
}

// HasLatest returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasLatest() bool {
	if o != nil && o.Latest != nil {
		return true
	}

	return false
}

// SetLatest gets a reference to the given bool and assigns it to the Latest field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetLatest(v bool) {
	o.Latest = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetValue(v float64) {
	o.Value = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetEndDate() int64 {
	if o == nil || o.EndDate == nil {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetEndDateOk() (*int64, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitNumeric returns the UnitNumeric field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetUnitNumeric() float64 {
	if o == nil || o.UnitNumeric == nil {
		var ret float64
		return ret
	}
	return *o.UnitNumeric
}

// GetUnitNumericOk returns a tuple with the UnitNumeric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetUnitNumericOk() (*float64, bool) {
	if o == nil || o.UnitNumeric == nil {
		return nil, false
	}
	return o.UnitNumeric, true
}

// HasUnitNumeric returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasUnitNumeric() bool {
	if o != nil && o.UnitNumeric != nil {
		return true
	}

	return false
}

// SetUnitNumeric gets a reference to the given float64 and assigns it to the UnitNumeric field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetUnitNumeric(v float64) {
	o.UnitNumeric = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *CompanyAnalysisExtendedRawDataResultsData) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *CompanyAnalysisExtendedRawDataResultsData) SetCurrency(v string) {
	o.Currency = &v
}

func (o CompanyAnalysisExtendedRawDataResultsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Quarter != nil {
		toSerialize["quarter"] = o.Quarter
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	if o.Latest != nil {
		toSerialize["latest"] = o.Latest
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.UnitNumeric != nil {
		toSerialize["unit_numeric"] = o.UnitNumeric
	}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	return json.Marshal(toSerialize)
}

type NullableCompanyAnalysisExtendedRawDataResultsData struct {
	value *CompanyAnalysisExtendedRawDataResultsData
	isSet bool
}

func (v NullableCompanyAnalysisExtendedRawDataResultsData) Get() *CompanyAnalysisExtendedRawDataResultsData {
	return v.value
}

func (v *NullableCompanyAnalysisExtendedRawDataResultsData) Set(val *CompanyAnalysisExtendedRawDataResultsData) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyAnalysisExtendedRawDataResultsData) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyAnalysisExtendedRawDataResultsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyAnalysisExtendedRawDataResultsData(val *CompanyAnalysisExtendedRawDataResultsData) *NullableCompanyAnalysisExtendedRawDataResultsData {
	return &NullableCompanyAnalysisExtendedRawDataResultsData{value: val, isSet: true}
}

func (v NullableCompanyAnalysisExtendedRawDataResultsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyAnalysisExtendedRawDataResultsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}