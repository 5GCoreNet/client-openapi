/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnrf_NFDiscovery

import (
	"encoding/json"
)

// PreferredSearch Contains information on whether the returned NFProfiles match the preferred query parameters 
type PreferredSearch struct {
	PreferredTaiMatchInd *bool `json:"preferredTaiMatchInd,omitempty"`
	PreferredFullPlmnMatchInd *bool `json:"preferredFullPlmnMatchInd,omitempty"`
	PreferredApiVersionsMatchInd *bool `json:"preferredApiVersionsMatchInd,omitempty"`
	OtherApiVersionsInd *bool `json:"otherApiVersionsInd,omitempty"`
	PreferredLocalityMatchInd *bool `json:"preferredLocalityMatchInd,omitempty"`
	OtherLocalityInd *bool `json:"otherLocalityInd,omitempty"`
	PreferredVendorSpecificFeaturesInd *bool `json:"preferredVendorSpecificFeaturesInd,omitempty"`
	PreferredCollocatedNfTypeInd *bool `json:"preferredCollocatedNfTypeInd,omitempty"`
	PreferredPgwMatchInd *bool `json:"preferredPgwMatchInd,omitempty"`
	PreferredAnalyticsDelaysInd *bool `json:"preferredAnalyticsDelaysInd,omitempty"`
}

// NewPreferredSearch instantiates a new PreferredSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreferredSearch() *PreferredSearch {
	this := PreferredSearch{}
	var preferredTaiMatchInd bool = false
	this.PreferredTaiMatchInd = &preferredTaiMatchInd
	var preferredFullPlmnMatchInd bool = false
	this.PreferredFullPlmnMatchInd = &preferredFullPlmnMatchInd
	var preferredLocalityMatchInd bool = false
	this.PreferredLocalityMatchInd = &preferredLocalityMatchInd
	var otherLocalityInd bool = false
	this.OtherLocalityInd = &otherLocalityInd
	var preferredVendorSpecificFeaturesInd bool = false
	this.PreferredVendorSpecificFeaturesInd = &preferredVendorSpecificFeaturesInd
	var preferredCollocatedNfTypeInd bool = false
	this.PreferredCollocatedNfTypeInd = &preferredCollocatedNfTypeInd
	return &this
}

// NewPreferredSearchWithDefaults instantiates a new PreferredSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreferredSearchWithDefaults() *PreferredSearch {
	this := PreferredSearch{}
	var preferredTaiMatchInd bool = false
	this.PreferredTaiMatchInd = &preferredTaiMatchInd
	var preferredFullPlmnMatchInd bool = false
	this.PreferredFullPlmnMatchInd = &preferredFullPlmnMatchInd
	var preferredLocalityMatchInd bool = false
	this.PreferredLocalityMatchInd = &preferredLocalityMatchInd
	var otherLocalityInd bool = false
	this.OtherLocalityInd = &otherLocalityInd
	var preferredVendorSpecificFeaturesInd bool = false
	this.PreferredVendorSpecificFeaturesInd = &preferredVendorSpecificFeaturesInd
	var preferredCollocatedNfTypeInd bool = false
	this.PreferredCollocatedNfTypeInd = &preferredCollocatedNfTypeInd
	return &this
}

// GetPreferredTaiMatchInd returns the PreferredTaiMatchInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredTaiMatchInd() bool {
	if o == nil || isNil(o.PreferredTaiMatchInd) {
		var ret bool
		return ret
	}
	return *o.PreferredTaiMatchInd
}

// GetPreferredTaiMatchIndOk returns a tuple with the PreferredTaiMatchInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredTaiMatchIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredTaiMatchInd) {
    return nil, false
	}
	return o.PreferredTaiMatchInd, true
}

// HasPreferredTaiMatchInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredTaiMatchInd() bool {
	if o != nil && !isNil(o.PreferredTaiMatchInd) {
		return true
	}

	return false
}

// SetPreferredTaiMatchInd gets a reference to the given bool and assigns it to the PreferredTaiMatchInd field.
func (o *PreferredSearch) SetPreferredTaiMatchInd(v bool) {
	o.PreferredTaiMatchInd = &v
}

// GetPreferredFullPlmnMatchInd returns the PreferredFullPlmnMatchInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredFullPlmnMatchInd() bool {
	if o == nil || isNil(o.PreferredFullPlmnMatchInd) {
		var ret bool
		return ret
	}
	return *o.PreferredFullPlmnMatchInd
}

// GetPreferredFullPlmnMatchIndOk returns a tuple with the PreferredFullPlmnMatchInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredFullPlmnMatchIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredFullPlmnMatchInd) {
    return nil, false
	}
	return o.PreferredFullPlmnMatchInd, true
}

// HasPreferredFullPlmnMatchInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredFullPlmnMatchInd() bool {
	if o != nil && !isNil(o.PreferredFullPlmnMatchInd) {
		return true
	}

	return false
}

// SetPreferredFullPlmnMatchInd gets a reference to the given bool and assigns it to the PreferredFullPlmnMatchInd field.
func (o *PreferredSearch) SetPreferredFullPlmnMatchInd(v bool) {
	o.PreferredFullPlmnMatchInd = &v
}

// GetPreferredApiVersionsMatchInd returns the PreferredApiVersionsMatchInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredApiVersionsMatchInd() bool {
	if o == nil || isNil(o.PreferredApiVersionsMatchInd) {
		var ret bool
		return ret
	}
	return *o.PreferredApiVersionsMatchInd
}

// GetPreferredApiVersionsMatchIndOk returns a tuple with the PreferredApiVersionsMatchInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredApiVersionsMatchIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredApiVersionsMatchInd) {
    return nil, false
	}
	return o.PreferredApiVersionsMatchInd, true
}

// HasPreferredApiVersionsMatchInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredApiVersionsMatchInd() bool {
	if o != nil && !isNil(o.PreferredApiVersionsMatchInd) {
		return true
	}

	return false
}

// SetPreferredApiVersionsMatchInd gets a reference to the given bool and assigns it to the PreferredApiVersionsMatchInd field.
func (o *PreferredSearch) SetPreferredApiVersionsMatchInd(v bool) {
	o.PreferredApiVersionsMatchInd = &v
}

// GetOtherApiVersionsInd returns the OtherApiVersionsInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetOtherApiVersionsInd() bool {
	if o == nil || isNil(o.OtherApiVersionsInd) {
		var ret bool
		return ret
	}
	return *o.OtherApiVersionsInd
}

// GetOtherApiVersionsIndOk returns a tuple with the OtherApiVersionsInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetOtherApiVersionsIndOk() (*bool, bool) {
	if o == nil || isNil(o.OtherApiVersionsInd) {
    return nil, false
	}
	return o.OtherApiVersionsInd, true
}

// HasOtherApiVersionsInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasOtherApiVersionsInd() bool {
	if o != nil && !isNil(o.OtherApiVersionsInd) {
		return true
	}

	return false
}

// SetOtherApiVersionsInd gets a reference to the given bool and assigns it to the OtherApiVersionsInd field.
func (o *PreferredSearch) SetOtherApiVersionsInd(v bool) {
	o.OtherApiVersionsInd = &v
}

// GetPreferredLocalityMatchInd returns the PreferredLocalityMatchInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredLocalityMatchInd() bool {
	if o == nil || isNil(o.PreferredLocalityMatchInd) {
		var ret bool
		return ret
	}
	return *o.PreferredLocalityMatchInd
}

// GetPreferredLocalityMatchIndOk returns a tuple with the PreferredLocalityMatchInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredLocalityMatchIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredLocalityMatchInd) {
    return nil, false
	}
	return o.PreferredLocalityMatchInd, true
}

// HasPreferredLocalityMatchInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredLocalityMatchInd() bool {
	if o != nil && !isNil(o.PreferredLocalityMatchInd) {
		return true
	}

	return false
}

// SetPreferredLocalityMatchInd gets a reference to the given bool and assigns it to the PreferredLocalityMatchInd field.
func (o *PreferredSearch) SetPreferredLocalityMatchInd(v bool) {
	o.PreferredLocalityMatchInd = &v
}

// GetOtherLocalityInd returns the OtherLocalityInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetOtherLocalityInd() bool {
	if o == nil || isNil(o.OtherLocalityInd) {
		var ret bool
		return ret
	}
	return *o.OtherLocalityInd
}

// GetOtherLocalityIndOk returns a tuple with the OtherLocalityInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetOtherLocalityIndOk() (*bool, bool) {
	if o == nil || isNil(o.OtherLocalityInd) {
    return nil, false
	}
	return o.OtherLocalityInd, true
}

// HasOtherLocalityInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasOtherLocalityInd() bool {
	if o != nil && !isNil(o.OtherLocalityInd) {
		return true
	}

	return false
}

// SetOtherLocalityInd gets a reference to the given bool and assigns it to the OtherLocalityInd field.
func (o *PreferredSearch) SetOtherLocalityInd(v bool) {
	o.OtherLocalityInd = &v
}

// GetPreferredVendorSpecificFeaturesInd returns the PreferredVendorSpecificFeaturesInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredVendorSpecificFeaturesInd() bool {
	if o == nil || isNil(o.PreferredVendorSpecificFeaturesInd) {
		var ret bool
		return ret
	}
	return *o.PreferredVendorSpecificFeaturesInd
}

// GetPreferredVendorSpecificFeaturesIndOk returns a tuple with the PreferredVendorSpecificFeaturesInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredVendorSpecificFeaturesIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredVendorSpecificFeaturesInd) {
    return nil, false
	}
	return o.PreferredVendorSpecificFeaturesInd, true
}

// HasPreferredVendorSpecificFeaturesInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredVendorSpecificFeaturesInd() bool {
	if o != nil && !isNil(o.PreferredVendorSpecificFeaturesInd) {
		return true
	}

	return false
}

// SetPreferredVendorSpecificFeaturesInd gets a reference to the given bool and assigns it to the PreferredVendorSpecificFeaturesInd field.
func (o *PreferredSearch) SetPreferredVendorSpecificFeaturesInd(v bool) {
	o.PreferredVendorSpecificFeaturesInd = &v
}

// GetPreferredCollocatedNfTypeInd returns the PreferredCollocatedNfTypeInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredCollocatedNfTypeInd() bool {
	if o == nil || isNil(o.PreferredCollocatedNfTypeInd) {
		var ret bool
		return ret
	}
	return *o.PreferredCollocatedNfTypeInd
}

// GetPreferredCollocatedNfTypeIndOk returns a tuple with the PreferredCollocatedNfTypeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredCollocatedNfTypeIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredCollocatedNfTypeInd) {
    return nil, false
	}
	return o.PreferredCollocatedNfTypeInd, true
}

// HasPreferredCollocatedNfTypeInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredCollocatedNfTypeInd() bool {
	if o != nil && !isNil(o.PreferredCollocatedNfTypeInd) {
		return true
	}

	return false
}

// SetPreferredCollocatedNfTypeInd gets a reference to the given bool and assigns it to the PreferredCollocatedNfTypeInd field.
func (o *PreferredSearch) SetPreferredCollocatedNfTypeInd(v bool) {
	o.PreferredCollocatedNfTypeInd = &v
}

// GetPreferredPgwMatchInd returns the PreferredPgwMatchInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredPgwMatchInd() bool {
	if o == nil || isNil(o.PreferredPgwMatchInd) {
		var ret bool
		return ret
	}
	return *o.PreferredPgwMatchInd
}

// GetPreferredPgwMatchIndOk returns a tuple with the PreferredPgwMatchInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredPgwMatchIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredPgwMatchInd) {
    return nil, false
	}
	return o.PreferredPgwMatchInd, true
}

// HasPreferredPgwMatchInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredPgwMatchInd() bool {
	if o != nil && !isNil(o.PreferredPgwMatchInd) {
		return true
	}

	return false
}

// SetPreferredPgwMatchInd gets a reference to the given bool and assigns it to the PreferredPgwMatchInd field.
func (o *PreferredSearch) SetPreferredPgwMatchInd(v bool) {
	o.PreferredPgwMatchInd = &v
}

// GetPreferredAnalyticsDelaysInd returns the PreferredAnalyticsDelaysInd field value if set, zero value otherwise.
func (o *PreferredSearch) GetPreferredAnalyticsDelaysInd() bool {
	if o == nil || isNil(o.PreferredAnalyticsDelaysInd) {
		var ret bool
		return ret
	}
	return *o.PreferredAnalyticsDelaysInd
}

// GetPreferredAnalyticsDelaysIndOk returns a tuple with the PreferredAnalyticsDelaysInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreferredSearch) GetPreferredAnalyticsDelaysIndOk() (*bool, bool) {
	if o == nil || isNil(o.PreferredAnalyticsDelaysInd) {
    return nil, false
	}
	return o.PreferredAnalyticsDelaysInd, true
}

// HasPreferredAnalyticsDelaysInd returns a boolean if a field has been set.
func (o *PreferredSearch) HasPreferredAnalyticsDelaysInd() bool {
	if o != nil && !isNil(o.PreferredAnalyticsDelaysInd) {
		return true
	}

	return false
}

// SetPreferredAnalyticsDelaysInd gets a reference to the given bool and assigns it to the PreferredAnalyticsDelaysInd field.
func (o *PreferredSearch) SetPreferredAnalyticsDelaysInd(v bool) {
	o.PreferredAnalyticsDelaysInd = &v
}

func (o PreferredSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PreferredTaiMatchInd) {
		toSerialize["preferredTaiMatchInd"] = o.PreferredTaiMatchInd
	}
	if !isNil(o.PreferredFullPlmnMatchInd) {
		toSerialize["preferredFullPlmnMatchInd"] = o.PreferredFullPlmnMatchInd
	}
	if !isNil(o.PreferredApiVersionsMatchInd) {
		toSerialize["preferredApiVersionsMatchInd"] = o.PreferredApiVersionsMatchInd
	}
	if !isNil(o.OtherApiVersionsInd) {
		toSerialize["otherApiVersionsInd"] = o.OtherApiVersionsInd
	}
	if !isNil(o.PreferredLocalityMatchInd) {
		toSerialize["preferredLocalityMatchInd"] = o.PreferredLocalityMatchInd
	}
	if !isNil(o.OtherLocalityInd) {
		toSerialize["otherLocalityInd"] = o.OtherLocalityInd
	}
	if !isNil(o.PreferredVendorSpecificFeaturesInd) {
		toSerialize["preferredVendorSpecificFeaturesInd"] = o.PreferredVendorSpecificFeaturesInd
	}
	if !isNil(o.PreferredCollocatedNfTypeInd) {
		toSerialize["preferredCollocatedNfTypeInd"] = o.PreferredCollocatedNfTypeInd
	}
	if !isNil(o.PreferredPgwMatchInd) {
		toSerialize["preferredPgwMatchInd"] = o.PreferredPgwMatchInd
	}
	if !isNil(o.PreferredAnalyticsDelaysInd) {
		toSerialize["preferredAnalyticsDelaysInd"] = o.PreferredAnalyticsDelaysInd
	}
	return json.Marshal(toSerialize)
}

type NullablePreferredSearch struct {
	value *PreferredSearch
	isSet bool
}

func (v NullablePreferredSearch) Get() *PreferredSearch {
	return v.value
}

func (v *NullablePreferredSearch) Set(val *PreferredSearch) {
	v.value = val
	v.isSet = true
}

func (v NullablePreferredSearch) IsSet() bool {
	return v.isSet
}

func (v *NullablePreferredSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreferredSearch(val *PreferredSearch) *NullablePreferredSearch {
	return &NullablePreferredSearch{value: val, isSet: true}
}

func (v NullablePreferredSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreferredSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


