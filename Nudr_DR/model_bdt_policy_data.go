/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nudr_DR

import (
	"encoding/json"
)

// BdtPolicyData Represents applied BDT policy data.
type BdtPolicyData struct {
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InterGroupId *string `json:"interGroupId,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	ResUri *string `json:"resUri,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewBdtPolicyData instantiates a new BdtPolicyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBdtPolicyData(bdtRefId string) *BdtPolicyData {
	this := BdtPolicyData{}
	this.BdtRefId = bdtRefId
	return &this
}

// NewBdtPolicyDataWithDefaults instantiates a new BdtPolicyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBdtPolicyDataWithDefaults() *BdtPolicyData {
	this := BdtPolicyData{}
	return &this
}

// GetInterGroupId returns the InterGroupId field value if set, zero value otherwise.
func (o *BdtPolicyData) GetInterGroupId() string {
	if o == nil || isNil(o.InterGroupId) {
		var ret string
		return ret
	}
	return *o.InterGroupId
}

// GetInterGroupIdOk returns a tuple with the InterGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetInterGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.InterGroupId) {
    return nil, false
	}
	return o.InterGroupId, true
}

// HasInterGroupId returns a boolean if a field has been set.
func (o *BdtPolicyData) HasInterGroupId() bool {
	if o != nil && !isNil(o.InterGroupId) {
		return true
	}

	return false
}

// SetInterGroupId gets a reference to the given string and assigns it to the InterGroupId field.
func (o *BdtPolicyData) SetInterGroupId(v string) {
	o.InterGroupId = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *BdtPolicyData) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *BdtPolicyData) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *BdtPolicyData) SetSupi(v string) {
	o.Supi = &v
}

// GetBdtRefId returns the BdtRefId field value
func (o *BdtPolicyData) GetBdtRefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetBdtRefIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.BdtRefId, true
}

// SetBdtRefId sets field value
func (o *BdtPolicyData) SetBdtRefId(v string) {
	o.BdtRefId = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *BdtPolicyData) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *BdtPolicyData) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *BdtPolicyData) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *BdtPolicyData) GetSnssai() Snssai {
	if o == nil || isNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || isNil(o.Snssai) {
    return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *BdtPolicyData) HasSnssai() bool {
	if o != nil && !isNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *BdtPolicyData) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetResUri returns the ResUri field value if set, zero value otherwise.
func (o *BdtPolicyData) GetResUri() string {
	if o == nil || isNil(o.ResUri) {
		var ret string
		return ret
	}
	return *o.ResUri
}

// GetResUriOk returns a tuple with the ResUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetResUriOk() (*string, bool) {
	if o == nil || isNil(o.ResUri) {
    return nil, false
	}
	return o.ResUri, true
}

// HasResUri returns a boolean if a field has been set.
func (o *BdtPolicyData) HasResUri() bool {
	if o != nil && !isNil(o.ResUri) {
		return true
	}

	return false
}

// SetResUri gets a reference to the given string and assigns it to the ResUri field.
func (o *BdtPolicyData) SetResUri(v string) {
	o.ResUri = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *BdtPolicyData) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *BdtPolicyData) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *BdtPolicyData) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o BdtPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.InterGroupId) {
		toSerialize["interGroupId"] = o.InterGroupId
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if true {
		toSerialize["bdtRefId"] = o.BdtRefId
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !isNil(o.ResUri) {
		toSerialize["resUri"] = o.ResUri
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return json.Marshal(toSerialize)
}

type NullableBdtPolicyData struct {
	value *BdtPolicyData
	isSet bool
}

func (v NullableBdtPolicyData) Get() *BdtPolicyData {
	return v.value
}

func (v *NullableBdtPolicyData) Set(val *BdtPolicyData) {
	v.value = val
	v.isSet = true
}

func (v NullableBdtPolicyData) IsSet() bool {
	return v.isSet
}

func (v *NullableBdtPolicyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBdtPolicyData(val *BdtPolicyData) *NullableBdtPolicyData {
	return &NullableBdtPolicyData{value: val, isSet: true}
}

func (v NullableBdtPolicyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBdtPolicyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


