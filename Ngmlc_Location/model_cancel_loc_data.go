/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
)

// CancelLocData Contains the input parameters in CancelLocation service operation
type CancelLocData struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String identifying External Group Identifier that identifies a group made up of one or more  subscriptions associated to a group of IMSIs, as specified in clause 19.7.3 of 3GPP TS 23.003.  
	ExtGroupId *string `json:"extGroupId,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	IntGroupId *string `json:"intGroupId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HgmlcCallBackUri string `json:"hgmlcCallBackUri"`
	// LDR Reference.
	LdrReference string `json:"ldrReference"`
	// LMF identification.
	LmfIdentification *string `json:"lmfIdentification,omitempty"`
	// String identifying the AMF ID composed of AMF Region ID (8 bits), AMF Set ID (10 bits) and AMF  Pointer (6 bits) as specified in clause 2.10.1 of 3GPP TS 23.003. It is encoded as a string of  6 hexadecimal characters (i.e., 24 bits).  
	AmfId *string `json:"amfId,omitempty"`
}

// NewCancelLocData instantiates a new CancelLocData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelLocData(hgmlcCallBackUri string, ldrReference string) *CancelLocData {
	this := CancelLocData{}
	this.HgmlcCallBackUri = hgmlcCallBackUri
	this.LdrReference = ldrReference
	return &this
}

// NewCancelLocDataWithDefaults instantiates a new CancelLocData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelLocDataWithDefaults() *CancelLocData {
	this := CancelLocData{}
	return &this
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *CancelLocData) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *CancelLocData) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *CancelLocData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *CancelLocData) GetSupi() string {
	if o == nil || isNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetSupiOk() (*string, bool) {
	if o == nil || isNil(o.Supi) {
    return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *CancelLocData) HasSupi() bool {
	if o != nil && !isNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *CancelLocData) SetSupi(v string) {
	o.Supi = &v
}

// GetExtGroupId returns the ExtGroupId field value if set, zero value otherwise.
func (o *CancelLocData) GetExtGroupId() string {
	if o == nil || isNil(o.ExtGroupId) {
		var ret string
		return ret
	}
	return *o.ExtGroupId
}

// GetExtGroupIdOk returns a tuple with the ExtGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetExtGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ExtGroupId) {
    return nil, false
	}
	return o.ExtGroupId, true
}

// HasExtGroupId returns a boolean if a field has been set.
func (o *CancelLocData) HasExtGroupId() bool {
	if o != nil && !isNil(o.ExtGroupId) {
		return true
	}

	return false
}

// SetExtGroupId gets a reference to the given string and assigns it to the ExtGroupId field.
func (o *CancelLocData) SetExtGroupId(v string) {
	o.ExtGroupId = &v
}

// GetIntGroupId returns the IntGroupId field value if set, zero value otherwise.
func (o *CancelLocData) GetIntGroupId() string {
	if o == nil || isNil(o.IntGroupId) {
		var ret string
		return ret
	}
	return *o.IntGroupId
}

// GetIntGroupIdOk returns a tuple with the IntGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetIntGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.IntGroupId) {
    return nil, false
	}
	return o.IntGroupId, true
}

// HasIntGroupId returns a boolean if a field has been set.
func (o *CancelLocData) HasIntGroupId() bool {
	if o != nil && !isNil(o.IntGroupId) {
		return true
	}

	return false
}

// SetIntGroupId gets a reference to the given string and assigns it to the IntGroupId field.
func (o *CancelLocData) SetIntGroupId(v string) {
	o.IntGroupId = &v
}

// GetHgmlcCallBackUri returns the HgmlcCallBackUri field value
func (o *CancelLocData) GetHgmlcCallBackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HgmlcCallBackUri
}

// GetHgmlcCallBackUriOk returns a tuple with the HgmlcCallBackUri field value
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetHgmlcCallBackUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.HgmlcCallBackUri, true
}

// SetHgmlcCallBackUri sets field value
func (o *CancelLocData) SetHgmlcCallBackUri(v string) {
	o.HgmlcCallBackUri = v
}

// GetLdrReference returns the LdrReference field value
func (o *CancelLocData) GetLdrReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdrReference
}

// GetLdrReferenceOk returns a tuple with the LdrReference field value
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetLdrReferenceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.LdrReference, true
}

// SetLdrReference sets field value
func (o *CancelLocData) SetLdrReference(v string) {
	o.LdrReference = v
}

// GetLmfIdentification returns the LmfIdentification field value if set, zero value otherwise.
func (o *CancelLocData) GetLmfIdentification() string {
	if o == nil || isNil(o.LmfIdentification) {
		var ret string
		return ret
	}
	return *o.LmfIdentification
}

// GetLmfIdentificationOk returns a tuple with the LmfIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetLmfIdentificationOk() (*string, bool) {
	if o == nil || isNil(o.LmfIdentification) {
    return nil, false
	}
	return o.LmfIdentification, true
}

// HasLmfIdentification returns a boolean if a field has been set.
func (o *CancelLocData) HasLmfIdentification() bool {
	if o != nil && !isNil(o.LmfIdentification) {
		return true
	}

	return false
}

// SetLmfIdentification gets a reference to the given string and assigns it to the LmfIdentification field.
func (o *CancelLocData) SetLmfIdentification(v string) {
	o.LmfIdentification = &v
}

// GetAmfId returns the AmfId field value if set, zero value otherwise.
func (o *CancelLocData) GetAmfId() string {
	if o == nil || isNil(o.AmfId) {
		var ret string
		return ret
	}
	return *o.AmfId
}

// GetAmfIdOk returns a tuple with the AmfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelLocData) GetAmfIdOk() (*string, bool) {
	if o == nil || isNil(o.AmfId) {
    return nil, false
	}
	return o.AmfId, true
}

// HasAmfId returns a boolean if a field has been set.
func (o *CancelLocData) HasAmfId() bool {
	if o != nil && !isNil(o.AmfId) {
		return true
	}

	return false
}

// SetAmfId gets a reference to the given string and assigns it to the AmfId field.
func (o *CancelLocData) SetAmfId(v string) {
	o.AmfId = &v
}

func (o CancelLocData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.ExtGroupId) {
		toSerialize["extGroupId"] = o.ExtGroupId
	}
	if !isNil(o.IntGroupId) {
		toSerialize["intGroupId"] = o.IntGroupId
	}
	if true {
		toSerialize["hgmlcCallBackUri"] = o.HgmlcCallBackUri
	}
	if true {
		toSerialize["ldrReference"] = o.LdrReference
	}
	if !isNil(o.LmfIdentification) {
		toSerialize["lmfIdentification"] = o.LmfIdentification
	}
	if !isNil(o.AmfId) {
		toSerialize["amfId"] = o.AmfId
	}
	return json.Marshal(toSerialize)
}

type NullableCancelLocData struct {
	value *CancelLocData
	isSet bool
}

func (v NullableCancelLocData) Get() *CancelLocData {
	return v.value
}

func (v *NullableCancelLocData) Set(val *CancelLocData) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelLocData) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelLocData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelLocData(val *CancelLocData) *NullableCancelLocData {
	return &NullableCancelLocData{value: val, isSet: true}
}

func (v NullableCancelLocData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelLocData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


