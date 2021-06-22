/*
 * Subset of Influx API covered by Influx CLI
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// TemplateExport struct for TemplateExport
type TemplateExport struct {
	StackID   *string                   `json:"stackID,omitempty"`
	OrgIDs    *[]TemplateExportOrgIDs   `json:"orgIDs,omitempty"`
	Resources []TemplateExportResources `json:"resources"`
}

// NewTemplateExport instantiates a new TemplateExport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateExport(resources []TemplateExportResources) *TemplateExport {
	this := TemplateExport{}
	this.Resources = resources
	return &this
}

// NewTemplateExportWithDefaults instantiates a new TemplateExport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateExportWithDefaults() *TemplateExport {
	this := TemplateExport{}
	return &this
}

// GetStackID returns the StackID field value if set, zero value otherwise.
func (o *TemplateExport) GetStackID() string {
	if o == nil || o.StackID == nil {
		var ret string
		return ret
	}
	return *o.StackID
}

// GetStackIDOk returns a tuple with the StackID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateExport) GetStackIDOk() (*string, bool) {
	if o == nil || o.StackID == nil {
		return nil, false
	}
	return o.StackID, true
}

// HasStackID returns a boolean if a field has been set.
func (o *TemplateExport) HasStackID() bool {
	if o != nil && o.StackID != nil {
		return true
	}

	return false
}

// SetStackID gets a reference to the given string and assigns it to the StackID field.
func (o *TemplateExport) SetStackID(v string) {
	o.StackID = &v
}

// GetOrgIDs returns the OrgIDs field value if set, zero value otherwise.
func (o *TemplateExport) GetOrgIDs() []TemplateExportOrgIDs {
	if o == nil || o.OrgIDs == nil {
		var ret []TemplateExportOrgIDs
		return ret
	}
	return *o.OrgIDs
}

// GetOrgIDsOk returns a tuple with the OrgIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateExport) GetOrgIDsOk() (*[]TemplateExportOrgIDs, bool) {
	if o == nil || o.OrgIDs == nil {
		return nil, false
	}
	return o.OrgIDs, true
}

// HasOrgIDs returns a boolean if a field has been set.
func (o *TemplateExport) HasOrgIDs() bool {
	if o != nil && o.OrgIDs != nil {
		return true
	}

	return false
}

// SetOrgIDs gets a reference to the given []TemplateExportOrgIDs and assigns it to the OrgIDs field.
func (o *TemplateExport) SetOrgIDs(v []TemplateExportOrgIDs) {
	o.OrgIDs = &v
}

// GetResources returns the Resources field value
func (o *TemplateExport) GetResources() []TemplateExportResources {
	if o == nil {
		var ret []TemplateExportResources
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *TemplateExport) GetResourcesOk() (*[]TemplateExportResources, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resources, true
}

// SetResources sets field value
func (o *TemplateExport) SetResources(v []TemplateExportResources) {
	o.Resources = v
}

func (o TemplateExport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StackID != nil {
		toSerialize["stackID"] = o.StackID
	}
	if o.OrgIDs != nil {
		toSerialize["orgIDs"] = o.OrgIDs
	}
	if true {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateExport struct {
	value *TemplateExport
	isSet bool
}

func (v NullableTemplateExport) Get() *TemplateExport {
	return v.value
}

func (v *NullableTemplateExport) Set(val *TemplateExport) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateExport) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateExport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateExport(val *TemplateExport) *NullableTemplateExport {
	return &NullableTemplateExport{value: val, isSet: true}
}

func (v NullableTemplateExport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateExport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}