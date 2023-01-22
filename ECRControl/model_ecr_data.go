/*
3gpp-ecr-control

API for enhanced converage restriction control.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ECRControl

import (
	"encoding/json"
)

// ECRData Represents the current visited PLMN (if any) and the current settings of enhanced coverage restriction.
type ECRData struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures"`
	VisitedPlmnId *PlmnId `json:"visitedPlmnId,omitempty"`
	EcrDataWbs []PlmnEcRestrictionDataWb `json:"ecrDataWbs,omitempty"`
	// Indicates a complete list (and possibly empty) of serving PLMNs where Enhanced Coverage shall be restricted.
	RestrictedPlmnIds []PlmnId `json:"restrictedPlmnIds,omitempty"`
	// Indicates a complete list (and possibly empty) of serving PLMNs where Enhanced Coverage shall be allowed.
	AllowedPlmnIds []PlmnId `json:"allowedPlmnIds,omitempty"`
}

// NewECRData instantiates a new ECRData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECRData(supportedFeatures string) *ECRData {
	this := ECRData{}
	this.SupportedFeatures = supportedFeatures
	return &this
}

// NewECRDataWithDefaults instantiates a new ECRData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECRDataWithDefaults() *ECRData {
	this := ECRData{}
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value
func (o *ECRData) GetSupportedFeatures() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value
// and a boolean to check if the value has been set.
func (o *ECRData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SupportedFeatures, true
}

// SetSupportedFeatures sets field value
func (o *ECRData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = v
}

// GetVisitedPlmnId returns the VisitedPlmnId field value if set, zero value otherwise.
func (o *ECRData) GetVisitedPlmnId() PlmnId {
	if o == nil || isNil(o.VisitedPlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.VisitedPlmnId
}

// GetVisitedPlmnIdOk returns a tuple with the VisitedPlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECRData) GetVisitedPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.VisitedPlmnId) {
    return nil, false
	}
	return o.VisitedPlmnId, true
}

// HasVisitedPlmnId returns a boolean if a field has been set.
func (o *ECRData) HasVisitedPlmnId() bool {
	if o != nil && !isNil(o.VisitedPlmnId) {
		return true
	}

	return false
}

// SetVisitedPlmnId gets a reference to the given PlmnId and assigns it to the VisitedPlmnId field.
func (o *ECRData) SetVisitedPlmnId(v PlmnId) {
	o.VisitedPlmnId = &v
}

// GetEcrDataWbs returns the EcrDataWbs field value if set, zero value otherwise.
func (o *ECRData) GetEcrDataWbs() []PlmnEcRestrictionDataWb {
	if o == nil || isNil(o.EcrDataWbs) {
		var ret []PlmnEcRestrictionDataWb
		return ret
	}
	return o.EcrDataWbs
}

// GetEcrDataWbsOk returns a tuple with the EcrDataWbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECRData) GetEcrDataWbsOk() ([]PlmnEcRestrictionDataWb, bool) {
	if o == nil || isNil(o.EcrDataWbs) {
    return nil, false
	}
	return o.EcrDataWbs, true
}

// HasEcrDataWbs returns a boolean if a field has been set.
func (o *ECRData) HasEcrDataWbs() bool {
	if o != nil && !isNil(o.EcrDataWbs) {
		return true
	}

	return false
}

// SetEcrDataWbs gets a reference to the given []PlmnEcRestrictionDataWb and assigns it to the EcrDataWbs field.
func (o *ECRData) SetEcrDataWbs(v []PlmnEcRestrictionDataWb) {
	o.EcrDataWbs = v
}

// GetRestrictedPlmnIds returns the RestrictedPlmnIds field value if set, zero value otherwise.
func (o *ECRData) GetRestrictedPlmnIds() []PlmnId {
	if o == nil || isNil(o.RestrictedPlmnIds) {
		var ret []PlmnId
		return ret
	}
	return o.RestrictedPlmnIds
}

// GetRestrictedPlmnIdsOk returns a tuple with the RestrictedPlmnIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECRData) GetRestrictedPlmnIdsOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.RestrictedPlmnIds) {
    return nil, false
	}
	return o.RestrictedPlmnIds, true
}

// HasRestrictedPlmnIds returns a boolean if a field has been set.
func (o *ECRData) HasRestrictedPlmnIds() bool {
	if o != nil && !isNil(o.RestrictedPlmnIds) {
		return true
	}

	return false
}

// SetRestrictedPlmnIds gets a reference to the given []PlmnId and assigns it to the RestrictedPlmnIds field.
func (o *ECRData) SetRestrictedPlmnIds(v []PlmnId) {
	o.RestrictedPlmnIds = v
}

// GetAllowedPlmnIds returns the AllowedPlmnIds field value if set, zero value otherwise.
func (o *ECRData) GetAllowedPlmnIds() []PlmnId {
	if o == nil || isNil(o.AllowedPlmnIds) {
		var ret []PlmnId
		return ret
	}
	return o.AllowedPlmnIds
}

// GetAllowedPlmnIdsOk returns a tuple with the AllowedPlmnIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECRData) GetAllowedPlmnIdsOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.AllowedPlmnIds) {
    return nil, false
	}
	return o.AllowedPlmnIds, true
}

// HasAllowedPlmnIds returns a boolean if a field has been set.
func (o *ECRData) HasAllowedPlmnIds() bool {
	if o != nil && !isNil(o.AllowedPlmnIds) {
		return true
	}

	return false
}

// SetAllowedPlmnIds gets a reference to the given []PlmnId and assigns it to the AllowedPlmnIds field.
func (o *ECRData) SetAllowedPlmnIds(v []PlmnId) {
	o.AllowedPlmnIds = v
}

func (o ECRData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.VisitedPlmnId) {
		toSerialize["visitedPlmnId"] = o.VisitedPlmnId
	}
	if !isNil(o.EcrDataWbs) {
		toSerialize["ecrDataWbs"] = o.EcrDataWbs
	}
	if !isNil(o.RestrictedPlmnIds) {
		toSerialize["restrictedPlmnIds"] = o.RestrictedPlmnIds
	}
	if !isNil(o.AllowedPlmnIds) {
		toSerialize["allowedPlmnIds"] = o.AllowedPlmnIds
	}
	return json.Marshal(toSerialize)
}

type NullableECRData struct {
	value *ECRData
	isSet bool
}

func (v NullableECRData) Get() *ECRData {
	return v.value
}

func (v *NullableECRData) Set(val *ECRData) {
	v.value = val
	v.isSet = true
}

func (v NullableECRData) IsSet() bool {
	return v.isSet
}

func (v *NullableECRData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECRData(val *ECRData) *NullableECRData {
	return &NullableECRData{value: val, isSet: true}
}

func (v NullableECRData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECRData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


