/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// AuthorizedDefaultQos Represents the Authorized Default QoS.
type AuthorizedDefaultQos struct {
	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Var5qi *int32 `json:"5qi,omitempty"`
	Arp *Arp `json:"arp,omitempty"`
	// This data type is defined in the same way as the '5QiPriorityLevel' data type, but with the OpenAPI 'nullable: true' property. 
	PriorityLevel NullableInt32 `json:"priorityLevel,omitempty"`
	// This data type is defined in the same way as the 'AverWindow' data type, but with the OpenAPI 'nullable: true' property. 
	AverWindow NullableInt32 `json:"averWindow,omitempty"`
	// This data type is defined in the same way as the 'MaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property. 
	MaxDataBurstVol NullableInt32 `json:"maxDataBurstVol,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxbrUl NullableString `json:"maxbrUl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxbrDl NullableString `json:"maxbrDl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	GbrUl NullableString `json:"gbrUl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	GbrDl NullableString `json:"gbrDl,omitempty"`
	// This data type is defined in the same way as the 'ExtMaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property. 
	ExtMaxDataBurstVol NullableInt32 `json:"extMaxDataBurstVol,omitempty"`
}

// NewAuthorizedDefaultQos instantiates a new AuthorizedDefaultQos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizedDefaultQos() *AuthorizedDefaultQos {
	this := AuthorizedDefaultQos{}
	var averWindow int32 = 2000
	this.AverWindow = *NewNullableInt32(&averWindow)
	return &this
}

// NewAuthorizedDefaultQosWithDefaults instantiates a new AuthorizedDefaultQos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizedDefaultQosWithDefaults() *AuthorizedDefaultQos {
	this := AuthorizedDefaultQos{}
	var averWindow int32 = 2000
	this.AverWindow = *NewNullableInt32(&averWindow)
	return &this
}

// GetVar5qi returns the Var5qi field value if set, zero value otherwise.
func (o *AuthorizedDefaultQos) GetVar5qi() int32 {
	if o == nil || isNil(o.Var5qi) {
		var ret int32
		return ret
	}
	return *o.Var5qi
}

// GetVar5qiOk returns a tuple with the Var5qi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedDefaultQos) GetVar5qiOk() (*int32, bool) {
	if o == nil || isNil(o.Var5qi) {
    return nil, false
	}
	return o.Var5qi, true
}

// HasVar5qi returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasVar5qi() bool {
	if o != nil && !isNil(o.Var5qi) {
		return true
	}

	return false
}

// SetVar5qi gets a reference to the given int32 and assigns it to the Var5qi field.
func (o *AuthorizedDefaultQos) SetVar5qi(v int32) {
	o.Var5qi = &v
}

// GetArp returns the Arp field value if set, zero value otherwise.
func (o *AuthorizedDefaultQos) GetArp() Arp {
	if o == nil || isNil(o.Arp) {
		var ret Arp
		return ret
	}
	return *o.Arp
}

// GetArpOk returns a tuple with the Arp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedDefaultQos) GetArpOk() (*Arp, bool) {
	if o == nil || isNil(o.Arp) {
    return nil, false
	}
	return o.Arp, true
}

// HasArp returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasArp() bool {
	if o != nil && !isNil(o.Arp) {
		return true
	}

	return false
}

// SetArp gets a reference to the given Arp and assigns it to the Arp field.
func (o *AuthorizedDefaultQos) SetArp(v Arp) {
	o.Arp = &v
}

// GetPriorityLevel returns the PriorityLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetPriorityLevel() int32 {
	if o == nil || isNil(o.PriorityLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.PriorityLevel.Get()
}

// GetPriorityLevelOk returns a tuple with the PriorityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetPriorityLevelOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.PriorityLevel.Get(), o.PriorityLevel.IsSet()
}

// HasPriorityLevel returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasPriorityLevel() bool {
	if o != nil && o.PriorityLevel.IsSet() {
		return true
	}

	return false
}

// SetPriorityLevel gets a reference to the given NullableInt32 and assigns it to the PriorityLevel field.
func (o *AuthorizedDefaultQos) SetPriorityLevel(v int32) {
	o.PriorityLevel.Set(&v)
}
// SetPriorityLevelNil sets the value for PriorityLevel to be an explicit nil
func (o *AuthorizedDefaultQos) SetPriorityLevelNil() {
	o.PriorityLevel.Set(nil)
}

// UnsetPriorityLevel ensures that no value is present for PriorityLevel, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetPriorityLevel() {
	o.PriorityLevel.Unset()
}

// GetAverWindow returns the AverWindow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetAverWindow() int32 {
	if o == nil || isNil(o.AverWindow.Get()) {
		var ret int32
		return ret
	}
	return *o.AverWindow.Get()
}

// GetAverWindowOk returns a tuple with the AverWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetAverWindowOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.AverWindow.Get(), o.AverWindow.IsSet()
}

// HasAverWindow returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasAverWindow() bool {
	if o != nil && o.AverWindow.IsSet() {
		return true
	}

	return false
}

// SetAverWindow gets a reference to the given NullableInt32 and assigns it to the AverWindow field.
func (o *AuthorizedDefaultQos) SetAverWindow(v int32) {
	o.AverWindow.Set(&v)
}
// SetAverWindowNil sets the value for AverWindow to be an explicit nil
func (o *AuthorizedDefaultQos) SetAverWindowNil() {
	o.AverWindow.Set(nil)
}

// UnsetAverWindow ensures that no value is present for AverWindow, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetAverWindow() {
	o.AverWindow.Unset()
}

// GetMaxDataBurstVol returns the MaxDataBurstVol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetMaxDataBurstVol() int32 {
	if o == nil || isNil(o.MaxDataBurstVol.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxDataBurstVol.Get()
}

// GetMaxDataBurstVolOk returns a tuple with the MaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetMaxDataBurstVolOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxDataBurstVol.Get(), o.MaxDataBurstVol.IsSet()
}

// HasMaxDataBurstVol returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasMaxDataBurstVol() bool {
	if o != nil && o.MaxDataBurstVol.IsSet() {
		return true
	}

	return false
}

// SetMaxDataBurstVol gets a reference to the given NullableInt32 and assigns it to the MaxDataBurstVol field.
func (o *AuthorizedDefaultQos) SetMaxDataBurstVol(v int32) {
	o.MaxDataBurstVol.Set(&v)
}
// SetMaxDataBurstVolNil sets the value for MaxDataBurstVol to be an explicit nil
func (o *AuthorizedDefaultQos) SetMaxDataBurstVolNil() {
	o.MaxDataBurstVol.Set(nil)
}

// UnsetMaxDataBurstVol ensures that no value is present for MaxDataBurstVol, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetMaxDataBurstVol() {
	o.MaxDataBurstVol.Unset()
}

// GetMaxbrUl returns the MaxbrUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetMaxbrUl() string {
	if o == nil || isNil(o.MaxbrUl.Get()) {
		var ret string
		return ret
	}
	return *o.MaxbrUl.Get()
}

// GetMaxbrUlOk returns a tuple with the MaxbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetMaxbrUlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxbrUl.Get(), o.MaxbrUl.IsSet()
}

// HasMaxbrUl returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasMaxbrUl() bool {
	if o != nil && o.MaxbrUl.IsSet() {
		return true
	}

	return false
}

// SetMaxbrUl gets a reference to the given NullableString and assigns it to the MaxbrUl field.
func (o *AuthorizedDefaultQos) SetMaxbrUl(v string) {
	o.MaxbrUl.Set(&v)
}
// SetMaxbrUlNil sets the value for MaxbrUl to be an explicit nil
func (o *AuthorizedDefaultQos) SetMaxbrUlNil() {
	o.MaxbrUl.Set(nil)
}

// UnsetMaxbrUl ensures that no value is present for MaxbrUl, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetMaxbrUl() {
	o.MaxbrUl.Unset()
}

// GetMaxbrDl returns the MaxbrDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetMaxbrDl() string {
	if o == nil || isNil(o.MaxbrDl.Get()) {
		var ret string
		return ret
	}
	return *o.MaxbrDl.Get()
}

// GetMaxbrDlOk returns a tuple with the MaxbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetMaxbrDlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxbrDl.Get(), o.MaxbrDl.IsSet()
}

// HasMaxbrDl returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasMaxbrDl() bool {
	if o != nil && o.MaxbrDl.IsSet() {
		return true
	}

	return false
}

// SetMaxbrDl gets a reference to the given NullableString and assigns it to the MaxbrDl field.
func (o *AuthorizedDefaultQos) SetMaxbrDl(v string) {
	o.MaxbrDl.Set(&v)
}
// SetMaxbrDlNil sets the value for MaxbrDl to be an explicit nil
func (o *AuthorizedDefaultQos) SetMaxbrDlNil() {
	o.MaxbrDl.Set(nil)
}

// UnsetMaxbrDl ensures that no value is present for MaxbrDl, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetMaxbrDl() {
	o.MaxbrDl.Unset()
}

// GetGbrUl returns the GbrUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetGbrUl() string {
	if o == nil || isNil(o.GbrUl.Get()) {
		var ret string
		return ret
	}
	return *o.GbrUl.Get()
}

// GetGbrUlOk returns a tuple with the GbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetGbrUlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.GbrUl.Get(), o.GbrUl.IsSet()
}

// HasGbrUl returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasGbrUl() bool {
	if o != nil && o.GbrUl.IsSet() {
		return true
	}

	return false
}

// SetGbrUl gets a reference to the given NullableString and assigns it to the GbrUl field.
func (o *AuthorizedDefaultQos) SetGbrUl(v string) {
	o.GbrUl.Set(&v)
}
// SetGbrUlNil sets the value for GbrUl to be an explicit nil
func (o *AuthorizedDefaultQos) SetGbrUlNil() {
	o.GbrUl.Set(nil)
}

// UnsetGbrUl ensures that no value is present for GbrUl, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetGbrUl() {
	o.GbrUl.Unset()
}

// GetGbrDl returns the GbrDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetGbrDl() string {
	if o == nil || isNil(o.GbrDl.Get()) {
		var ret string
		return ret
	}
	return *o.GbrDl.Get()
}

// GetGbrDlOk returns a tuple with the GbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetGbrDlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.GbrDl.Get(), o.GbrDl.IsSet()
}

// HasGbrDl returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasGbrDl() bool {
	if o != nil && o.GbrDl.IsSet() {
		return true
	}

	return false
}

// SetGbrDl gets a reference to the given NullableString and assigns it to the GbrDl field.
func (o *AuthorizedDefaultQos) SetGbrDl(v string) {
	o.GbrDl.Set(&v)
}
// SetGbrDlNil sets the value for GbrDl to be an explicit nil
func (o *AuthorizedDefaultQos) SetGbrDlNil() {
	o.GbrDl.Set(nil)
}

// UnsetGbrDl ensures that no value is present for GbrDl, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetGbrDl() {
	o.GbrDl.Unset()
}

// GetExtMaxDataBurstVol returns the ExtMaxDataBurstVol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthorizedDefaultQos) GetExtMaxDataBurstVol() int32 {
	if o == nil || isNil(o.ExtMaxDataBurstVol.Get()) {
		var ret int32
		return ret
	}
	return *o.ExtMaxDataBurstVol.Get()
}

// GetExtMaxDataBurstVolOk returns a tuple with the ExtMaxDataBurstVol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthorizedDefaultQos) GetExtMaxDataBurstVolOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.ExtMaxDataBurstVol.Get(), o.ExtMaxDataBurstVol.IsSet()
}

// HasExtMaxDataBurstVol returns a boolean if a field has been set.
func (o *AuthorizedDefaultQos) HasExtMaxDataBurstVol() bool {
	if o != nil && o.ExtMaxDataBurstVol.IsSet() {
		return true
	}

	return false
}

// SetExtMaxDataBurstVol gets a reference to the given NullableInt32 and assigns it to the ExtMaxDataBurstVol field.
func (o *AuthorizedDefaultQos) SetExtMaxDataBurstVol(v int32) {
	o.ExtMaxDataBurstVol.Set(&v)
}
// SetExtMaxDataBurstVolNil sets the value for ExtMaxDataBurstVol to be an explicit nil
func (o *AuthorizedDefaultQos) SetExtMaxDataBurstVolNil() {
	o.ExtMaxDataBurstVol.Set(nil)
}

// UnsetExtMaxDataBurstVol ensures that no value is present for ExtMaxDataBurstVol, not even an explicit nil
func (o *AuthorizedDefaultQos) UnsetExtMaxDataBurstVol() {
	o.ExtMaxDataBurstVol.Unset()
}

func (o AuthorizedDefaultQos) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Var5qi) {
		toSerialize["5qi"] = o.Var5qi
	}
	if !isNil(o.Arp) {
		toSerialize["arp"] = o.Arp
	}
	if o.PriorityLevel.IsSet() {
		toSerialize["priorityLevel"] = o.PriorityLevel.Get()
	}
	if o.AverWindow.IsSet() {
		toSerialize["averWindow"] = o.AverWindow.Get()
	}
	if o.MaxDataBurstVol.IsSet() {
		toSerialize["maxDataBurstVol"] = o.MaxDataBurstVol.Get()
	}
	if o.MaxbrUl.IsSet() {
		toSerialize["maxbrUl"] = o.MaxbrUl.Get()
	}
	if o.MaxbrDl.IsSet() {
		toSerialize["maxbrDl"] = o.MaxbrDl.Get()
	}
	if o.GbrUl.IsSet() {
		toSerialize["gbrUl"] = o.GbrUl.Get()
	}
	if o.GbrDl.IsSet() {
		toSerialize["gbrDl"] = o.GbrDl.Get()
	}
	if o.ExtMaxDataBurstVol.IsSet() {
		toSerialize["extMaxDataBurstVol"] = o.ExtMaxDataBurstVol.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizedDefaultQos struct {
	value *AuthorizedDefaultQos
	isSet bool
}

func (v NullableAuthorizedDefaultQos) Get() *AuthorizedDefaultQos {
	return v.value
}

func (v *NullableAuthorizedDefaultQos) Set(val *AuthorizedDefaultQos) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizedDefaultQos) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizedDefaultQos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizedDefaultQos(val *AuthorizedDefaultQos) *NullableAuthorizedDefaultQos {
	return &NullableAuthorizedDefaultQos{value: val, isSet: true}
}

func (v NullableAuthorizedDefaultQos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizedDefaultQos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


