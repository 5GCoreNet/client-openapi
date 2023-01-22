/*
3gpp-mbs-us

API for MBS User Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package MBSUserService

import (
	"encoding/json"
)

// MBSUserService Represents the parameters of an MBS User Service.
type MBSUserService struct {
	ExtServiceIds []string `json:"extServiceIds"`
	ServType MbsServiceType `json:"servType"`
	// String providing an URI formatted according to RFC 3986.
	ServClass string `json:"servClass"`
	ServAnnModes []ServiceAnnouncementMode `json:"servAnnModes"`
	ServNameDescs []ServiceNameDescription `json:"servNameDescs,omitempty"`
	MainServLang *string `json:"mainServLang,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewMBSUserService instantiates a new MBSUserService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMBSUserService(extServiceIds []string, servType MbsServiceType, servClass string, servAnnModes []ServiceAnnouncementMode) *MBSUserService {
	this := MBSUserService{}
	this.ExtServiceIds = extServiceIds
	this.ServType = servType
	this.ServClass = servClass
	this.ServAnnModes = servAnnModes
	return &this
}

// NewMBSUserServiceWithDefaults instantiates a new MBSUserService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMBSUserServiceWithDefaults() *MBSUserService {
	this := MBSUserService{}
	return &this
}

// GetExtServiceIds returns the ExtServiceIds field value
func (o *MBSUserService) GetExtServiceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExtServiceIds
}

// GetExtServiceIdsOk returns a tuple with the ExtServiceIds field value
// and a boolean to check if the value has been set.
func (o *MBSUserService) GetExtServiceIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExtServiceIds, true
}

// SetExtServiceIds sets field value
func (o *MBSUserService) SetExtServiceIds(v []string) {
	o.ExtServiceIds = v
}

// GetServType returns the ServType field value
func (o *MBSUserService) GetServType() MbsServiceType {
	if o == nil {
		var ret MbsServiceType
		return ret
	}

	return o.ServType
}

// GetServTypeOk returns a tuple with the ServType field value
// and a boolean to check if the value has been set.
func (o *MBSUserService) GetServTypeOk() (*MbsServiceType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServType, true
}

// SetServType sets field value
func (o *MBSUserService) SetServType(v MbsServiceType) {
	o.ServType = v
}

// GetServClass returns the ServClass field value
func (o *MBSUserService) GetServClass() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServClass
}

// GetServClassOk returns a tuple with the ServClass field value
// and a boolean to check if the value has been set.
func (o *MBSUserService) GetServClassOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServClass, true
}

// SetServClass sets field value
func (o *MBSUserService) SetServClass(v string) {
	o.ServClass = v
}

// GetServAnnModes returns the ServAnnModes field value
func (o *MBSUserService) GetServAnnModes() []ServiceAnnouncementMode {
	if o == nil {
		var ret []ServiceAnnouncementMode
		return ret
	}

	return o.ServAnnModes
}

// GetServAnnModesOk returns a tuple with the ServAnnModes field value
// and a boolean to check if the value has been set.
func (o *MBSUserService) GetServAnnModesOk() ([]ServiceAnnouncementMode, bool) {
	if o == nil {
    return nil, false
	}
	return o.ServAnnModes, true
}

// SetServAnnModes sets field value
func (o *MBSUserService) SetServAnnModes(v []ServiceAnnouncementMode) {
	o.ServAnnModes = v
}

// GetServNameDescs returns the ServNameDescs field value if set, zero value otherwise.
func (o *MBSUserService) GetServNameDescs() []ServiceNameDescription {
	if o == nil || isNil(o.ServNameDescs) {
		var ret []ServiceNameDescription
		return ret
	}
	return o.ServNameDescs
}

// GetServNameDescsOk returns a tuple with the ServNameDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserService) GetServNameDescsOk() ([]ServiceNameDescription, bool) {
	if o == nil || isNil(o.ServNameDescs) {
    return nil, false
	}
	return o.ServNameDescs, true
}

// HasServNameDescs returns a boolean if a field has been set.
func (o *MBSUserService) HasServNameDescs() bool {
	if o != nil && !isNil(o.ServNameDescs) {
		return true
	}

	return false
}

// SetServNameDescs gets a reference to the given []ServiceNameDescription and assigns it to the ServNameDescs field.
func (o *MBSUserService) SetServNameDescs(v []ServiceNameDescription) {
	o.ServNameDescs = v
}

// GetMainServLang returns the MainServLang field value if set, zero value otherwise.
func (o *MBSUserService) GetMainServLang() string {
	if o == nil || isNil(o.MainServLang) {
		var ret string
		return ret
	}
	return *o.MainServLang
}

// GetMainServLangOk returns a tuple with the MainServLang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserService) GetMainServLangOk() (*string, bool) {
	if o == nil || isNil(o.MainServLang) {
    return nil, false
	}
	return o.MainServLang, true
}

// HasMainServLang returns a boolean if a field has been set.
func (o *MBSUserService) HasMainServLang() bool {
	if o != nil && !isNil(o.MainServLang) {
		return true
	}

	return false
}

// SetMainServLang gets a reference to the given string and assigns it to the MainServLang field.
func (o *MBSUserService) SetMainServLang(v string) {
	o.MainServLang = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *MBSUserService) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MBSUserService) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *MBSUserService) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *MBSUserService) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o MBSUserService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["extServiceIds"] = o.ExtServiceIds
	}
	if true {
		toSerialize["servType"] = o.ServType
	}
	if true {
		toSerialize["servClass"] = o.ServClass
	}
	if true {
		toSerialize["servAnnModes"] = o.ServAnnModes
	}
	if !isNil(o.ServNameDescs) {
		toSerialize["servNameDescs"] = o.ServNameDescs
	}
	if !isNil(o.MainServLang) {
		toSerialize["mainServLang"] = o.MainServLang
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableMBSUserService struct {
	value *MBSUserService
	isSet bool
}

func (v NullableMBSUserService) Get() *MBSUserService {
	return v.value
}

func (v *NullableMBSUserService) Set(val *MBSUserService) {
	v.value = val
	v.isSet = true
}

func (v NullableMBSUserService) IsSet() bool {
	return v.isSet
}

func (v *NullableMBSUserService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMBSUserService(val *MBSUserService) *NullableMBSUserService {
	return &NullableMBSUserService{value: val, isSet: true}
}

func (v NullableMBSUserService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMBSUserService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


