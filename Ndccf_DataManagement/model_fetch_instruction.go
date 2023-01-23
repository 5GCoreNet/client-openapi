/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"time"
)

// FetchInstruction The fetch instructions indicate whether the data or analytics are to be fetched by the consumer. 
type FetchInstruction struct {
	// String providing an URI formatted according to RFC 3986.
	FetchUri string `json:"fetchUri"`
	// The fetch correlation identifier(s) of the MFAF Data or Analytics.
	FetchCorrIds []string `json:"fetchCorrIds"`
	// string with format 'date-time' as defined in OpenAPI.
	Expiry *time.Time `json:"expiry,omitempty"`
}

// NewFetchInstruction instantiates a new FetchInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchInstruction(fetchUri string, fetchCorrIds []string) *FetchInstruction {
	this := FetchInstruction{}
	this.FetchUri = fetchUri
	this.FetchCorrIds = fetchCorrIds
	return &this
}

// NewFetchInstructionWithDefaults instantiates a new FetchInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchInstructionWithDefaults() *FetchInstruction {
	this := FetchInstruction{}
	return &this
}

// GetFetchUri returns the FetchUri field value
func (o *FetchInstruction) GetFetchUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FetchUri
}

// GetFetchUriOk returns a tuple with the FetchUri field value
// and a boolean to check if the value has been set.
func (o *FetchInstruction) GetFetchUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.FetchUri, true
}

// SetFetchUri sets field value
func (o *FetchInstruction) SetFetchUri(v string) {
	o.FetchUri = v
}

// GetFetchCorrIds returns the FetchCorrIds field value
func (o *FetchInstruction) GetFetchCorrIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FetchCorrIds
}

// GetFetchCorrIdsOk returns a tuple with the FetchCorrIds field value
// and a boolean to check if the value has been set.
func (o *FetchInstruction) GetFetchCorrIdsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FetchCorrIds, true
}

// SetFetchCorrIds sets field value
func (o *FetchInstruction) SetFetchCorrIds(v []string) {
	o.FetchCorrIds = v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *FetchInstruction) GetExpiry() time.Time {
	if o == nil || isNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchInstruction) GetExpiryOk() (*time.Time, bool) {
	if o == nil || isNil(o.Expiry) {
    return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *FetchInstruction) HasExpiry() bool {
	if o != nil && !isNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *FetchInstruction) SetExpiry(v time.Time) {
	o.Expiry = &v
}

func (o FetchInstruction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fetchUri"] = o.FetchUri
	}
	if true {
		toSerialize["fetchCorrIds"] = o.FetchCorrIds
	}
	if !isNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	return json.Marshal(toSerialize)
}

type NullableFetchInstruction struct {
	value *FetchInstruction
	isSet bool
}

func (v NullableFetchInstruction) Get() *FetchInstruction {
	return v.value
}

func (v *NullableFetchInstruction) Set(val *FetchInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchInstruction(val *FetchInstruction) *NullableFetchInstruction {
	return &NullableFetchInstruction{value: val, isSet: true}
}

func (v NullableFetchInstruction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


