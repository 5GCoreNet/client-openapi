/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// NFProfile NF profile stored in NRF, defined in TS 29.510
type NFProfile struct {
	// uuid of NF instance
	NFInstanceId *string `json:"nFInstanceId,omitempty"`
	NFType *NFType `json:"nFType,omitempty"`
	NFStatus *NFStatus `json:"nFStatus,omitempty"`
	Plmn *PlmnId `json:"plmn,omitempty"`
	SNssais *Snssai `json:"sNssais,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	InterPlmnFqdn *string `json:"interPlmnFqdn,omitempty"`
	NfServices []NFService `json:"nfServices,omitempty"`
}

// NewNFProfile instantiates a new NFProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFProfile() *NFProfile {
	this := NFProfile{}
	return &this
}

// NewNFProfileWithDefaults instantiates a new NFProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFProfileWithDefaults() *NFProfile {
	this := NFProfile{}
	return &this
}

// GetNFInstanceId returns the NFInstanceId field value if set, zero value otherwise.
func (o *NFProfile) GetNFInstanceId() string {
	if o == nil || isNil(o.NFInstanceId) {
		var ret string
		return ret
	}
	return *o.NFInstanceId
}

// GetNFInstanceIdOk returns a tuple with the NFInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNFInstanceIdOk() (*string, bool) {
	if o == nil || isNil(o.NFInstanceId) {
    return nil, false
	}
	return o.NFInstanceId, true
}

// HasNFInstanceId returns a boolean if a field has been set.
func (o *NFProfile) HasNFInstanceId() bool {
	if o != nil && !isNil(o.NFInstanceId) {
		return true
	}

	return false
}

// SetNFInstanceId gets a reference to the given string and assigns it to the NFInstanceId field.
func (o *NFProfile) SetNFInstanceId(v string) {
	o.NFInstanceId = &v
}

// GetNFType returns the NFType field value if set, zero value otherwise.
func (o *NFProfile) GetNFType() NFType {
	if o == nil || isNil(o.NFType) {
		var ret NFType
		return ret
	}
	return *o.NFType
}

// GetNFTypeOk returns a tuple with the NFType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNFTypeOk() (*NFType, bool) {
	if o == nil || isNil(o.NFType) {
    return nil, false
	}
	return o.NFType, true
}

// HasNFType returns a boolean if a field has been set.
func (o *NFProfile) HasNFType() bool {
	if o != nil && !isNil(o.NFType) {
		return true
	}

	return false
}

// SetNFType gets a reference to the given NFType and assigns it to the NFType field.
func (o *NFProfile) SetNFType(v NFType) {
	o.NFType = &v
}

// GetNFStatus returns the NFStatus field value if set, zero value otherwise.
func (o *NFProfile) GetNFStatus() NFStatus {
	if o == nil || isNil(o.NFStatus) {
		var ret NFStatus
		return ret
	}
	return *o.NFStatus
}

// GetNFStatusOk returns a tuple with the NFStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNFStatusOk() (*NFStatus, bool) {
	if o == nil || isNil(o.NFStatus) {
    return nil, false
	}
	return o.NFStatus, true
}

// HasNFStatus returns a boolean if a field has been set.
func (o *NFProfile) HasNFStatus() bool {
	if o != nil && !isNil(o.NFStatus) {
		return true
	}

	return false
}

// SetNFStatus gets a reference to the given NFStatus and assigns it to the NFStatus field.
func (o *NFProfile) SetNFStatus(v NFStatus) {
	o.NFStatus = &v
}

// GetPlmn returns the Plmn field value if set, zero value otherwise.
func (o *NFProfile) GetPlmn() PlmnId {
	if o == nil || isNil(o.Plmn) {
		var ret PlmnId
		return ret
	}
	return *o.Plmn
}

// GetPlmnOk returns a tuple with the Plmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetPlmnOk() (*PlmnId, bool) {
	if o == nil || isNil(o.Plmn) {
    return nil, false
	}
	return o.Plmn, true
}

// HasPlmn returns a boolean if a field has been set.
func (o *NFProfile) HasPlmn() bool {
	if o != nil && !isNil(o.Plmn) {
		return true
	}

	return false
}

// SetPlmn gets a reference to the given PlmnId and assigns it to the Plmn field.
func (o *NFProfile) SetPlmn(v PlmnId) {
	o.Plmn = &v
}

// GetSNssais returns the SNssais field value if set, zero value otherwise.
func (o *NFProfile) GetSNssais() Snssai {
	if o == nil || isNil(o.SNssais) {
		var ret Snssai
		return ret
	}
	return *o.SNssais
}

// GetSNssaisOk returns a tuple with the SNssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetSNssaisOk() (*Snssai, bool) {
	if o == nil || isNil(o.SNssais) {
    return nil, false
	}
	return o.SNssais, true
}

// HasSNssais returns a boolean if a field has been set.
func (o *NFProfile) HasSNssais() bool {
	if o != nil && !isNil(o.SNssais) {
		return true
	}

	return false
}

// SetSNssais gets a reference to the given Snssai and assigns it to the SNssais field.
func (o *NFProfile) SetSNssais(v Snssai) {
	o.SNssais = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *NFProfile) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *NFProfile) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *NFProfile) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetInterPlmnFqdn returns the InterPlmnFqdn field value if set, zero value otherwise.
func (o *NFProfile) GetInterPlmnFqdn() string {
	if o == nil || isNil(o.InterPlmnFqdn) {
		var ret string
		return ret
	}
	return *o.InterPlmnFqdn
}

// GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetInterPlmnFqdnOk() (*string, bool) {
	if o == nil || isNil(o.InterPlmnFqdn) {
    return nil, false
	}
	return o.InterPlmnFqdn, true
}

// HasInterPlmnFqdn returns a boolean if a field has been set.
func (o *NFProfile) HasInterPlmnFqdn() bool {
	if o != nil && !isNil(o.InterPlmnFqdn) {
		return true
	}

	return false
}

// SetInterPlmnFqdn gets a reference to the given string and assigns it to the InterPlmnFqdn field.
func (o *NFProfile) SetInterPlmnFqdn(v string) {
	o.InterPlmnFqdn = &v
}

// GetNfServices returns the NfServices field value if set, zero value otherwise.
func (o *NFProfile) GetNfServices() []NFService {
	if o == nil || isNil(o.NfServices) {
		var ret []NFService
		return ret
	}
	return o.NfServices
}

// GetNfServicesOk returns a tuple with the NfServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFProfile) GetNfServicesOk() ([]NFService, bool) {
	if o == nil || isNil(o.NfServices) {
    return nil, false
	}
	return o.NfServices, true
}

// HasNfServices returns a boolean if a field has been set.
func (o *NFProfile) HasNfServices() bool {
	if o != nil && !isNil(o.NfServices) {
		return true
	}

	return false
}

// SetNfServices gets a reference to the given []NFService and assigns it to the NfServices field.
func (o *NFProfile) SetNfServices(v []NFService) {
	o.NfServices = v
}

func (o NFProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NFInstanceId) {
		toSerialize["nFInstanceId"] = o.NFInstanceId
	}
	if !isNil(o.NFType) {
		toSerialize["nFType"] = o.NFType
	}
	if !isNil(o.NFStatus) {
		toSerialize["nFStatus"] = o.NFStatus
	}
	if !isNil(o.Plmn) {
		toSerialize["plmn"] = o.Plmn
	}
	if !isNil(o.SNssais) {
		toSerialize["sNssais"] = o.SNssais
	}
	if !isNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !isNil(o.InterPlmnFqdn) {
		toSerialize["interPlmnFqdn"] = o.InterPlmnFqdn
	}
	if !isNil(o.NfServices) {
		toSerialize["nfServices"] = o.NfServices
	}
	return json.Marshal(toSerialize)
}

type NullableNFProfile struct {
	value *NFProfile
	isSet bool
}

func (v NullableNFProfile) Get() *NFProfile {
	return v.value
}

func (v *NullableNFProfile) Set(val *NFProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableNFProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableNFProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFProfile(val *NFProfile) *NullableNFProfile {
	return &NullableNFProfile{value: val, isSet: true}
}

func (v NullableNFProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


