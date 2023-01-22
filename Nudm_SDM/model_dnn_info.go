/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudm_SDM

import (
	"encoding/json"
)

// DnnInfo struct for DnnInfo
type DnnInfo struct {
	Dnn AccessAndMobilitySubscriptionDataSubscribedDnnListInner `json:"dnn"`
	DefaultDnnIndicator *bool `json:"defaultDnnIndicator,omitempty"`
	LboRoamingAllowed *bool `json:"lboRoamingAllowed,omitempty"`
	IwkEpsInd *bool `json:"iwkEpsInd,omitempty"`
	DnnBarred *bool `json:"dnnBarred,omitempty"`
	InvokeNefInd *bool `json:"invokeNefInd,omitempty"`
	SmfList []string `json:"smfList,omitempty"`
	SameSmfInd *bool `json:"sameSmfInd,omitempty"`
}

// NewDnnInfo instantiates a new DnnInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnInfo(dnn AccessAndMobilitySubscriptionDataSubscribedDnnListInner) *DnnInfo {
	this := DnnInfo{}
	this.Dnn = dnn
	return &this
}

// NewDnnInfoWithDefaults instantiates a new DnnInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnInfoWithDefaults() *DnnInfo {
	this := DnnInfo{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnInfo) GetDnn() AccessAndMobilitySubscriptionDataSubscribedDnnListInner {
	if o == nil {
		var ret AccessAndMobilitySubscriptionDataSubscribedDnnListInner
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetDnnOk() (*AccessAndMobilitySubscriptionDataSubscribedDnnListInner, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnInfo) SetDnn(v AccessAndMobilitySubscriptionDataSubscribedDnnListInner) {
	o.Dnn = v
}

// GetDefaultDnnIndicator returns the DefaultDnnIndicator field value if set, zero value otherwise.
func (o *DnnInfo) GetDefaultDnnIndicator() bool {
	if o == nil || isNil(o.DefaultDnnIndicator) {
		var ret bool
		return ret
	}
	return *o.DefaultDnnIndicator
}

// GetDefaultDnnIndicatorOk returns a tuple with the DefaultDnnIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetDefaultDnnIndicatorOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultDnnIndicator) {
    return nil, false
	}
	return o.DefaultDnnIndicator, true
}

// HasDefaultDnnIndicator returns a boolean if a field has been set.
func (o *DnnInfo) HasDefaultDnnIndicator() bool {
	if o != nil && !isNil(o.DefaultDnnIndicator) {
		return true
	}

	return false
}

// SetDefaultDnnIndicator gets a reference to the given bool and assigns it to the DefaultDnnIndicator field.
func (o *DnnInfo) SetDefaultDnnIndicator(v bool) {
	o.DefaultDnnIndicator = &v
}

// GetLboRoamingAllowed returns the LboRoamingAllowed field value if set, zero value otherwise.
func (o *DnnInfo) GetLboRoamingAllowed() bool {
	if o == nil || isNil(o.LboRoamingAllowed) {
		var ret bool
		return ret
	}
	return *o.LboRoamingAllowed
}

// GetLboRoamingAllowedOk returns a tuple with the LboRoamingAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetLboRoamingAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.LboRoamingAllowed) {
    return nil, false
	}
	return o.LboRoamingAllowed, true
}

// HasLboRoamingAllowed returns a boolean if a field has been set.
func (o *DnnInfo) HasLboRoamingAllowed() bool {
	if o != nil && !isNil(o.LboRoamingAllowed) {
		return true
	}

	return false
}

// SetLboRoamingAllowed gets a reference to the given bool and assigns it to the LboRoamingAllowed field.
func (o *DnnInfo) SetLboRoamingAllowed(v bool) {
	o.LboRoamingAllowed = &v
}

// GetIwkEpsInd returns the IwkEpsInd field value if set, zero value otherwise.
func (o *DnnInfo) GetIwkEpsInd() bool {
	if o == nil || isNil(o.IwkEpsInd) {
		var ret bool
		return ret
	}
	return *o.IwkEpsInd
}

// GetIwkEpsIndOk returns a tuple with the IwkEpsInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetIwkEpsIndOk() (*bool, bool) {
	if o == nil || isNil(o.IwkEpsInd) {
    return nil, false
	}
	return o.IwkEpsInd, true
}

// HasIwkEpsInd returns a boolean if a field has been set.
func (o *DnnInfo) HasIwkEpsInd() bool {
	if o != nil && !isNil(o.IwkEpsInd) {
		return true
	}

	return false
}

// SetIwkEpsInd gets a reference to the given bool and assigns it to the IwkEpsInd field.
func (o *DnnInfo) SetIwkEpsInd(v bool) {
	o.IwkEpsInd = &v
}

// GetDnnBarred returns the DnnBarred field value if set, zero value otherwise.
func (o *DnnInfo) GetDnnBarred() bool {
	if o == nil || isNil(o.DnnBarred) {
		var ret bool
		return ret
	}
	return *o.DnnBarred
}

// GetDnnBarredOk returns a tuple with the DnnBarred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetDnnBarredOk() (*bool, bool) {
	if o == nil || isNil(o.DnnBarred) {
    return nil, false
	}
	return o.DnnBarred, true
}

// HasDnnBarred returns a boolean if a field has been set.
func (o *DnnInfo) HasDnnBarred() bool {
	if o != nil && !isNil(o.DnnBarred) {
		return true
	}

	return false
}

// SetDnnBarred gets a reference to the given bool and assigns it to the DnnBarred field.
func (o *DnnInfo) SetDnnBarred(v bool) {
	o.DnnBarred = &v
}

// GetInvokeNefInd returns the InvokeNefInd field value if set, zero value otherwise.
func (o *DnnInfo) GetInvokeNefInd() bool {
	if o == nil || isNil(o.InvokeNefInd) {
		var ret bool
		return ret
	}
	return *o.InvokeNefInd
}

// GetInvokeNefIndOk returns a tuple with the InvokeNefInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetInvokeNefIndOk() (*bool, bool) {
	if o == nil || isNil(o.InvokeNefInd) {
    return nil, false
	}
	return o.InvokeNefInd, true
}

// HasInvokeNefInd returns a boolean if a field has been set.
func (o *DnnInfo) HasInvokeNefInd() bool {
	if o != nil && !isNil(o.InvokeNefInd) {
		return true
	}

	return false
}

// SetInvokeNefInd gets a reference to the given bool and assigns it to the InvokeNefInd field.
func (o *DnnInfo) SetInvokeNefInd(v bool) {
	o.InvokeNefInd = &v
}

// GetSmfList returns the SmfList field value if set, zero value otherwise.
func (o *DnnInfo) GetSmfList() []string {
	if o == nil || isNil(o.SmfList) {
		var ret []string
		return ret
	}
	return o.SmfList
}

// GetSmfListOk returns a tuple with the SmfList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetSmfListOk() ([]string, bool) {
	if o == nil || isNil(o.SmfList) {
    return nil, false
	}
	return o.SmfList, true
}

// HasSmfList returns a boolean if a field has been set.
func (o *DnnInfo) HasSmfList() bool {
	if o != nil && !isNil(o.SmfList) {
		return true
	}

	return false
}

// SetSmfList gets a reference to the given []string and assigns it to the SmfList field.
func (o *DnnInfo) SetSmfList(v []string) {
	o.SmfList = v
}

// GetSameSmfInd returns the SameSmfInd field value if set, zero value otherwise.
func (o *DnnInfo) GetSameSmfInd() bool {
	if o == nil || isNil(o.SameSmfInd) {
		var ret bool
		return ret
	}
	return *o.SameSmfInd
}

// GetSameSmfIndOk returns a tuple with the SameSmfInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnInfo) GetSameSmfIndOk() (*bool, bool) {
	if o == nil || isNil(o.SameSmfInd) {
    return nil, false
	}
	return o.SameSmfInd, true
}

// HasSameSmfInd returns a boolean if a field has been set.
func (o *DnnInfo) HasSameSmfInd() bool {
	if o != nil && !isNil(o.SameSmfInd) {
		return true
	}

	return false
}

// SetSameSmfInd gets a reference to the given bool and assigns it to the SameSmfInd field.
func (o *DnnInfo) SetSameSmfInd(v bool) {
	o.SameSmfInd = &v
}

func (o DnnInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.DefaultDnnIndicator) {
		toSerialize["defaultDnnIndicator"] = o.DefaultDnnIndicator
	}
	if !isNil(o.LboRoamingAllowed) {
		toSerialize["lboRoamingAllowed"] = o.LboRoamingAllowed
	}
	if !isNil(o.IwkEpsInd) {
		toSerialize["iwkEpsInd"] = o.IwkEpsInd
	}
	if !isNil(o.DnnBarred) {
		toSerialize["dnnBarred"] = o.DnnBarred
	}
	if !isNil(o.InvokeNefInd) {
		toSerialize["invokeNefInd"] = o.InvokeNefInd
	}
	if !isNil(o.SmfList) {
		toSerialize["smfList"] = o.SmfList
	}
	if !isNil(o.SameSmfInd) {
		toSerialize["sameSmfInd"] = o.SameSmfInd
	}
	return json.Marshal(toSerialize)
}

type NullableDnnInfo struct {
	value *DnnInfo
	isSet bool
}

func (v NullableDnnInfo) Get() *DnnInfo {
	return v.value
}

func (v *NullableDnnInfo) Set(val *DnnInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnInfo(val *DnnInfo) *NullableDnnInfo {
	return &NullableDnnInfo{value: val, isSet: true}
}

func (v NullableDnnInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


