/*
 * Influx API Service (V1 compatible endpoints)
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// InfluxQLResponseResults struct for InfluxQLResponseResults
type InfluxQLResponseResults struct {
	StatementId *int32                    `json:"statement_id,omitempty"`
	Series      *[]InfluxQLResponseSeries `json:"series,omitempty"`
}

// NewInfluxQLResponseResults instantiates a new InfluxQLResponseResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfluxQLResponseResults() *InfluxQLResponseResults {
	this := InfluxQLResponseResults{}
	return &this
}

// NewInfluxQLResponseResultsWithDefaults instantiates a new InfluxQLResponseResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfluxQLResponseResultsWithDefaults() *InfluxQLResponseResults {
	this := InfluxQLResponseResults{}
	return &this
}

// GetStatementId returns the StatementId field value if set, zero value otherwise.
func (o *InfluxQLResponseResults) GetStatementId() int32 {
	if o == nil || o.StatementId == nil {
		var ret int32
		return ret
	}
	return *o.StatementId
}

// GetStatementIdOk returns a tuple with the StatementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxQLResponseResults) GetStatementIdOk() (*int32, bool) {
	if o == nil || o.StatementId == nil {
		return nil, false
	}
	return o.StatementId, true
}

// HasStatementId returns a boolean if a field has been set.
func (o *InfluxQLResponseResults) HasStatementId() bool {
	if o != nil && o.StatementId != nil {
		return true
	}

	return false
}

// SetStatementId gets a reference to the given int32 and assigns it to the StatementId field.
func (o *InfluxQLResponseResults) SetStatementId(v int32) {
	o.StatementId = &v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *InfluxQLResponseResults) GetSeries() []InfluxQLResponseSeries {
	if o == nil || o.Series == nil {
		var ret []InfluxQLResponseSeries
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfluxQLResponseResults) GetSeriesOk() (*[]InfluxQLResponseSeries, bool) {
	if o == nil || o.Series == nil {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *InfluxQLResponseResults) HasSeries() bool {
	if o != nil && o.Series != nil {
		return true
	}

	return false
}

// SetSeries gets a reference to the given []InfluxQLResponseSeries and assigns it to the Series field.
func (o *InfluxQLResponseResults) SetSeries(v []InfluxQLResponseSeries) {
	o.Series = &v
}

func (o InfluxQLResponseResults) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StatementId != nil {
		toSerialize["statement_id"] = o.StatementId
	}
	if o.Series != nil {
		toSerialize["series"] = o.Series
	}
	return json.Marshal(toSerialize)
}

type NullableInfluxQLResponseResults struct {
	value *InfluxQLResponseResults
	isSet bool
}

func (v NullableInfluxQLResponseResults) Get() *InfluxQLResponseResults {
	return v.value
}

func (v *NullableInfluxQLResponseResults) Set(val *InfluxQLResponseResults) {
	v.value = val
	v.isSet = true
}

func (v NullableInfluxQLResponseResults) IsSet() bool {
	return v.isSet
}

func (v *NullableInfluxQLResponseResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfluxQLResponseResults(val *InfluxQLResponseResults) *NullableInfluxQLResponseResults {
	return &NullableInfluxQLResponseResults{value: val, isSet: true}
}

func (v NullableInfluxQLResponseResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfluxQLResponseResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
