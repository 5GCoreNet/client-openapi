/*
NSSF NSSAI Availability

NSSF NSSAI Availability Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnssf_NSSAIAvailability

import (
	"encoding/json"
	"time"
)

// NssfEventSubscriptionCreateData This contains the information for event subscription
type NssfEventSubscriptionCreateData struct {
	// String providing an URI formatted according to RFC 3986.
	NfNssaiAvailabilityUri string `json:"nfNssaiAvailabilityUri"`
	TaiList []Tai `json:"taiList"`
	Event NssfEventType `json:"event"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
	AmfSetId *string `json:"amfSetId,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfId *string `json:"amfId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	AllAmfSetTaiInd *bool `json:"allAmfSetTaiInd,omitempty"`
}

// NewNssfEventSubscriptionCreateData instantiates a new NssfEventSubscriptionCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNssfEventSubscriptionCreateData(nfNssaiAvailabilityUri string, taiList []Tai, event NssfEventType) *NssfEventSubscriptionCreateData {
	this := NssfEventSubscriptionCreateData{}
	this.NfNssaiAvailabilityUri = nfNssaiAvailabilityUri
	this.TaiList = taiList
	this.Event = event
	var allAmfSetTaiInd bool = false
	this.AllAmfSetTaiInd = &allAmfSetTaiInd
	return &this
}

// NewNssfEventSubscriptionCreateDataWithDefaults instantiates a new NssfEventSubscriptionCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNssfEventSubscriptionCreateDataWithDefaults() *NssfEventSubscriptionCreateData {
	this := NssfEventSubscriptionCreateData{}
	var allAmfSetTaiInd bool = false
	this.AllAmfSetTaiInd = &allAmfSetTaiInd
	return &this
}

// GetNfNssaiAvailabilityUri returns the NfNssaiAvailabilityUri field value
func (o *NssfEventSubscriptionCreateData) GetNfNssaiAvailabilityUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfNssaiAvailabilityUri
}

// GetNfNssaiAvailabilityUriOk returns a tuple with the NfNssaiAvailabilityUri field value
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetNfNssaiAvailabilityUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfNssaiAvailabilityUri, true
}

// SetNfNssaiAvailabilityUri sets field value
func (o *NssfEventSubscriptionCreateData) SetNfNssaiAvailabilityUri(v string) {
	o.NfNssaiAvailabilityUri = v
}

// GetTaiList returns the TaiList field value
func (o *NssfEventSubscriptionCreateData) GetTaiList() []Tai {
	if o == nil {
		var ret []Tai
		return ret
	}

	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetTaiListOk() ([]Tai, bool) {
	if o == nil {
    return nil, false
	}
	return o.TaiList, true
}

// SetTaiList sets field value
func (o *NssfEventSubscriptionCreateData) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetEvent returns the Event field value
func (o *NssfEventSubscriptionCreateData) GetEvent() NssfEventType {
	if o == nil {
		var ret NssfEventType
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetEventOk() (*NssfEventType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *NssfEventSubscriptionCreateData) SetEvent(v NssfEventType) {
	o.Event = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreateData) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreateData) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *NssfEventSubscriptionCreateData) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetAmfSetId returns the AmfSetId field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreateData) GetAmfSetId() string {
	if o == nil || isNil(o.AmfSetId) {
		var ret string
		return ret
	}
	return *o.AmfSetId
}

// GetAmfSetIdOk returns a tuple with the AmfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetAmfSetIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfSetId) {
    return nil, false
	}
	return o.AmfSetId, true
}

// HasAmfSetId returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreateData) HasAmfSetId() bool {
	if o != nil && !isNil(o.AmfSetId) {
		return true
	}

	return false
}

// SetAmfSetId gets a reference to the given string and assigns it to the AmfSetId field.
func (o *NssfEventSubscriptionCreateData) SetAmfSetId(v string) {
	o.AmfSetId = &v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreateData) GetTaiRangeList() []TaiRange {
	if o == nil || isNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || isNil(o.TaiRangeList) {
    return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreateData) HasTaiRangeList() bool {
	if o != nil && !isNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NssfEventSubscriptionCreateData) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetAmfId returns the AmfId field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreateData) GetAmfId() string {
	if o == nil || isNil(o.AmfId) {
		var ret string
		return ret
	}
	return *o.AmfId
}

// GetAmfIdOk returns a tuple with the AmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetAmfIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfId) {
    return nil, false
	}
	return o.AmfId, true
}

// HasAmfId returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreateData) HasAmfId() bool {
	if o != nil && !isNil(o.AmfId) {
		return true
	}

	return false
}

// SetAmfId gets a reference to the given string and assigns it to the AmfId field.
func (o *NssfEventSubscriptionCreateData) SetAmfId(v string) {
	o.AmfId = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreateData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreateData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *NssfEventSubscriptionCreateData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAllAmfSetTaiInd returns the AllAmfSetTaiInd field value if set, zero value otherwise.
func (o *NssfEventSubscriptionCreateData) GetAllAmfSetTaiInd() bool {
	if o == nil || isNil(o.AllAmfSetTaiInd) {
		var ret bool
		return ret
	}
	return *o.AllAmfSetTaiInd
}

// GetAllAmfSetTaiIndOk returns a tuple with the AllAmfSetTaiInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NssfEventSubscriptionCreateData) GetAllAmfSetTaiIndOk() (*bool, bool) {
	if o == nil || isNil(o.AllAmfSetTaiInd) {
    return nil, false
	}
	return o.AllAmfSetTaiInd, true
}

// HasAllAmfSetTaiInd returns a boolean if a field has been set.
func (o *NssfEventSubscriptionCreateData) HasAllAmfSetTaiInd() bool {
	if o != nil && !isNil(o.AllAmfSetTaiInd) {
		return true
	}

	return false
}

// SetAllAmfSetTaiInd gets a reference to the given bool and assigns it to the AllAmfSetTaiInd field.
func (o *NssfEventSubscriptionCreateData) SetAllAmfSetTaiInd(v bool) {
	o.AllAmfSetTaiInd = &v
}

func (o NssfEventSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nfNssaiAvailabilityUri"] = o.NfNssaiAvailabilityUri
	}
	if true {
		toSerialize["taiList"] = o.TaiList
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !isNil(o.AmfSetId) {
		toSerialize["amfSetId"] = o.AmfSetId
	}
	if !isNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !isNil(o.AmfId) {
		toSerialize["amfId"] = o.AmfId
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.AllAmfSetTaiInd) {
		toSerialize["allAmfSetTaiInd"] = o.AllAmfSetTaiInd
	}
	return json.Marshal(toSerialize)
}

type NullableNssfEventSubscriptionCreateData struct {
	value *NssfEventSubscriptionCreateData
	isSet bool
}

func (v NullableNssfEventSubscriptionCreateData) Get() *NssfEventSubscriptionCreateData {
	return v.value
}

func (v *NullableNssfEventSubscriptionCreateData) Set(val *NssfEventSubscriptionCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableNssfEventSubscriptionCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableNssfEventSubscriptionCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNssfEventSubscriptionCreateData(val *NssfEventSubscriptionCreateData) *NullableNssfEventSubscriptionCreateData {
	return &NullableNssfEventSubscriptionCreateData{value: val, isSet: true}
}

func (v NullableNssfEventSubscriptionCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNssfEventSubscriptionCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


