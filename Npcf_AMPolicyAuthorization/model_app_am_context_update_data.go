/*
Npcf_AMPolicyAuthorization Service API

PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Npcf_AMPolicyAuthorization

import (
	"encoding/json"
)

// AppAmContextUpdateData Describes the modifications to an Individual Application AM resource.
type AppAmContextUpdateData struct {
	// String providing an URI formatted according to RFC 3986.
	TermNotifUri *string `json:"termNotifUri,omitempty"`
	EvSubsc NullableAmEventsSubscDataRm `json:"evSubsc,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	Expiry NullableInt32 `json:"expiry,omitempty"`
	// Indicates whether high throughput is desired for the indicated UE traffic.
	HighThruInd NullableBool `json:"highThruInd,omitempty"`
	// Identifies a list of Tracking Areas per serving network where service is allowed.
	CovReq []ServiceAreaCoverageInfo `json:"covReq,omitempty"`
	AsTimeDisParam NullableAsTimeDistributionParam `json:"asTimeDisParam,omitempty"`
}

// NewAppAmContextUpdateData instantiates a new AppAmContextUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAmContextUpdateData() *AppAmContextUpdateData {
	this := AppAmContextUpdateData{}
	return &this
}

// NewAppAmContextUpdateDataWithDefaults instantiates a new AppAmContextUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAmContextUpdateDataWithDefaults() *AppAmContextUpdateData {
	this := AppAmContextUpdateData{}
	return &this
}

// GetTermNotifUri returns the TermNotifUri field value if set, zero value otherwise.
func (o *AppAmContextUpdateData) GetTermNotifUri() string {
	if o == nil || isNil(o.TermNotifUri) {
		var ret string
		return ret
	}
	return *o.TermNotifUri
}

// GetTermNotifUriOk returns a tuple with the TermNotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppAmContextUpdateData) GetTermNotifUriOk() (*string, bool) {
	if o == nil || isNil(o.TermNotifUri) {
    return nil, false
	}
	return o.TermNotifUri, true
}

// HasTermNotifUri returns a boolean if a field has been set.
func (o *AppAmContextUpdateData) HasTermNotifUri() bool {
	if o != nil && !isNil(o.TermNotifUri) {
		return true
	}

	return false
}

// SetTermNotifUri gets a reference to the given string and assigns it to the TermNotifUri field.
func (o *AppAmContextUpdateData) SetTermNotifUri(v string) {
	o.TermNotifUri = &v
}

// GetEvSubsc returns the EvSubsc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAmContextUpdateData) GetEvSubsc() AmEventsSubscDataRm {
	if o == nil || isNil(o.EvSubsc.Get()) {
		var ret AmEventsSubscDataRm
		return ret
	}
	return *o.EvSubsc.Get()
}

// GetEvSubscOk returns a tuple with the EvSubsc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAmContextUpdateData) GetEvSubscOk() (*AmEventsSubscDataRm, bool) {
	if o == nil {
    return nil, false
	}
	return o.EvSubsc.Get(), o.EvSubsc.IsSet()
}

// HasEvSubsc returns a boolean if a field has been set.
func (o *AppAmContextUpdateData) HasEvSubsc() bool {
	if o != nil && o.EvSubsc.IsSet() {
		return true
	}

	return false
}

// SetEvSubsc gets a reference to the given NullableAmEventsSubscDataRm and assigns it to the EvSubsc field.
func (o *AppAmContextUpdateData) SetEvSubsc(v AmEventsSubscDataRm) {
	o.EvSubsc.Set(&v)
}
// SetEvSubscNil sets the value for EvSubsc to be an explicit nil
func (o *AppAmContextUpdateData) SetEvSubscNil() {
	o.EvSubsc.Set(nil)
}

// UnsetEvSubsc ensures that no value is present for EvSubsc, not even an explicit nil
func (o *AppAmContextUpdateData) UnsetEvSubsc() {
	o.EvSubsc.Unset()
}

// GetExpiry returns the Expiry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAmContextUpdateData) GetExpiry() int32 {
	if o == nil || isNil(o.Expiry.Get()) {
		var ret int32
		return ret
	}
	return *o.Expiry.Get()
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAmContextUpdateData) GetExpiryOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.Expiry.Get(), o.Expiry.IsSet()
}

// HasExpiry returns a boolean if a field has been set.
func (o *AppAmContextUpdateData) HasExpiry() bool {
	if o != nil && o.Expiry.IsSet() {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given NullableInt32 and assigns it to the Expiry field.
func (o *AppAmContextUpdateData) SetExpiry(v int32) {
	o.Expiry.Set(&v)
}
// SetExpiryNil sets the value for Expiry to be an explicit nil
func (o *AppAmContextUpdateData) SetExpiryNil() {
	o.Expiry.Set(nil)
}

// UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
func (o *AppAmContextUpdateData) UnsetExpiry() {
	o.Expiry.Unset()
}

// GetHighThruInd returns the HighThruInd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAmContextUpdateData) GetHighThruInd() bool {
	if o == nil || isNil(o.HighThruInd.Get()) {
		var ret bool
		return ret
	}
	return *o.HighThruInd.Get()
}

// GetHighThruIndOk returns a tuple with the HighThruInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAmContextUpdateData) GetHighThruIndOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.HighThruInd.Get(), o.HighThruInd.IsSet()
}

// HasHighThruInd returns a boolean if a field has been set.
func (o *AppAmContextUpdateData) HasHighThruInd() bool {
	if o != nil && o.HighThruInd.IsSet() {
		return true
	}

	return false
}

// SetHighThruInd gets a reference to the given NullableBool and assigns it to the HighThruInd field.
func (o *AppAmContextUpdateData) SetHighThruInd(v bool) {
	o.HighThruInd.Set(&v)
}
// SetHighThruIndNil sets the value for HighThruInd to be an explicit nil
func (o *AppAmContextUpdateData) SetHighThruIndNil() {
	o.HighThruInd.Set(nil)
}

// UnsetHighThruInd ensures that no value is present for HighThruInd, not even an explicit nil
func (o *AppAmContextUpdateData) UnsetHighThruInd() {
	o.HighThruInd.Unset()
}

// GetCovReq returns the CovReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAmContextUpdateData) GetCovReq() []ServiceAreaCoverageInfo {
	if o == nil {
		var ret []ServiceAreaCoverageInfo
		return ret
	}
	return o.CovReq
}

// GetCovReqOk returns a tuple with the CovReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAmContextUpdateData) GetCovReqOk() ([]ServiceAreaCoverageInfo, bool) {
	if o == nil || isNil(o.CovReq) {
    return nil, false
	}
	return o.CovReq, true
}

// HasCovReq returns a boolean if a field has been set.
func (o *AppAmContextUpdateData) HasCovReq() bool {
	if o != nil && isNil(o.CovReq) {
		return true
	}

	return false
}

// SetCovReq gets a reference to the given []ServiceAreaCoverageInfo and assigns it to the CovReq field.
func (o *AppAmContextUpdateData) SetCovReq(v []ServiceAreaCoverageInfo) {
	o.CovReq = v
}

// GetAsTimeDisParam returns the AsTimeDisParam field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppAmContextUpdateData) GetAsTimeDisParam() AsTimeDistributionParam {
	if o == nil || isNil(o.AsTimeDisParam.Get()) {
		var ret AsTimeDistributionParam
		return ret
	}
	return *o.AsTimeDisParam.Get()
}

// GetAsTimeDisParamOk returns a tuple with the AsTimeDisParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppAmContextUpdateData) GetAsTimeDisParamOk() (*AsTimeDistributionParam, bool) {
	if o == nil {
    return nil, false
	}
	return o.AsTimeDisParam.Get(), o.AsTimeDisParam.IsSet()
}

// HasAsTimeDisParam returns a boolean if a field has been set.
func (o *AppAmContextUpdateData) HasAsTimeDisParam() bool {
	if o != nil && o.AsTimeDisParam.IsSet() {
		return true
	}

	return false
}

// SetAsTimeDisParam gets a reference to the given NullableAsTimeDistributionParam and assigns it to the AsTimeDisParam field.
func (o *AppAmContextUpdateData) SetAsTimeDisParam(v AsTimeDistributionParam) {
	o.AsTimeDisParam.Set(&v)
}
// SetAsTimeDisParamNil sets the value for AsTimeDisParam to be an explicit nil
func (o *AppAmContextUpdateData) SetAsTimeDisParamNil() {
	o.AsTimeDisParam.Set(nil)
}

// UnsetAsTimeDisParam ensures that no value is present for AsTimeDisParam, not even an explicit nil
func (o *AppAmContextUpdateData) UnsetAsTimeDisParam() {
	o.AsTimeDisParam.Unset()
}

func (o AppAmContextUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TermNotifUri) {
		toSerialize["termNotifUri"] = o.TermNotifUri
	}
	if o.EvSubsc.IsSet() {
		toSerialize["evSubsc"] = o.EvSubsc.Get()
	}
	if o.Expiry.IsSet() {
		toSerialize["expiry"] = o.Expiry.Get()
	}
	if o.HighThruInd.IsSet() {
		toSerialize["highThruInd"] = o.HighThruInd.Get()
	}
	if o.CovReq != nil {
		toSerialize["covReq"] = o.CovReq
	}
	if o.AsTimeDisParam.IsSet() {
		toSerialize["asTimeDisParam"] = o.AsTimeDisParam.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAppAmContextUpdateData struct {
	value *AppAmContextUpdateData
	isSet bool
}

func (v NullableAppAmContextUpdateData) Get() *AppAmContextUpdateData {
	return v.value
}

func (v *NullableAppAmContextUpdateData) Set(val *AppAmContextUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAmContextUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAmContextUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAmContextUpdateData(val *AppAmContextUpdateData) *NullableAppAmContextUpdateData {
	return &NullableAppAmContextUpdateData{value: val, isSet: true}
}

func (v NullableAppAmContextUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAmContextUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


