/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
	"time"
)

// BootstrappingInfoResponse Response body of the HTTP POST operation for resource /bootstrapping-info-request
type BootstrappingInfoResponse struct {
	// ME Key Material (hex-encoded string)
	MeKeyMaterial string `json:"meKeyMaterial"`
	// UICC key material (hex-encoded string)
	UiccKeyMaterial *string `json:"uiccKeyMaterial,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	KeyExpiryTime *time.Time `json:"keyExpiryTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	BootstrappingInfoCreationTime *time.Time `json:"bootstrappingInfoCreationTime,omitempty"`
	UssList []UssListItem `json:"ussList,omitempty"`
	GbaType *GbaType `json:"gbaType,omitempty"`
	// IMS Private Identity of the UE
	Impi *string `json:"impi,omitempty"`
}

// NewBootstrappingInfoResponse instantiates a new BootstrappingInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBootstrappingInfoResponse(meKeyMaterial string) *BootstrappingInfoResponse {
	this := BootstrappingInfoResponse{}
	this.MeKeyMaterial = meKeyMaterial
	return &this
}

// NewBootstrappingInfoResponseWithDefaults instantiates a new BootstrappingInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBootstrappingInfoResponseWithDefaults() *BootstrappingInfoResponse {
	this := BootstrappingInfoResponse{}
	return &this
}

// GetMeKeyMaterial returns the MeKeyMaterial field value
func (o *BootstrappingInfoResponse) GetMeKeyMaterial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MeKeyMaterial
}

// GetMeKeyMaterialOk returns a tuple with the MeKeyMaterial field value
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoResponse) GetMeKeyMaterialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MeKeyMaterial, true
}

// SetMeKeyMaterial sets field value
func (o *BootstrappingInfoResponse) SetMeKeyMaterial(v string) {
	o.MeKeyMaterial = v
}

// GetUiccKeyMaterial returns the UiccKeyMaterial field value if set, zero value otherwise.
func (o *BootstrappingInfoResponse) GetUiccKeyMaterial() string {
	if o == nil || isNil(o.UiccKeyMaterial) {
		var ret string
		return ret
	}
	return *o.UiccKeyMaterial
}

// GetUiccKeyMaterialOk returns a tuple with the UiccKeyMaterial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoResponse) GetUiccKeyMaterialOk() (*string, bool) {
	if o == nil || isNil(o.UiccKeyMaterial) {
    return nil, false
	}
	return o.UiccKeyMaterial, true
}

// HasUiccKeyMaterial returns a boolean if a field has been set.
func (o *BootstrappingInfoResponse) HasUiccKeyMaterial() bool {
	if o != nil && !isNil(o.UiccKeyMaterial) {
		return true
	}

	return false
}

// SetUiccKeyMaterial gets a reference to the given string and assigns it to the UiccKeyMaterial field.
func (o *BootstrappingInfoResponse) SetUiccKeyMaterial(v string) {
	o.UiccKeyMaterial = &v
}

// GetKeyExpiryTime returns the KeyExpiryTime field value if set, zero value otherwise.
func (o *BootstrappingInfoResponse) GetKeyExpiryTime() time.Time {
	if o == nil || isNil(o.KeyExpiryTime) {
		var ret time.Time
		return ret
	}
	return *o.KeyExpiryTime
}

// GetKeyExpiryTimeOk returns a tuple with the KeyExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoResponse) GetKeyExpiryTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.KeyExpiryTime) {
    return nil, false
	}
	return o.KeyExpiryTime, true
}

// HasKeyExpiryTime returns a boolean if a field has been set.
func (o *BootstrappingInfoResponse) HasKeyExpiryTime() bool {
	if o != nil && !isNil(o.KeyExpiryTime) {
		return true
	}

	return false
}

// SetKeyExpiryTime gets a reference to the given time.Time and assigns it to the KeyExpiryTime field.
func (o *BootstrappingInfoResponse) SetKeyExpiryTime(v time.Time) {
	o.KeyExpiryTime = &v
}

// GetBootstrappingInfoCreationTime returns the BootstrappingInfoCreationTime field value if set, zero value otherwise.
func (o *BootstrappingInfoResponse) GetBootstrappingInfoCreationTime() time.Time {
	if o == nil || isNil(o.BootstrappingInfoCreationTime) {
		var ret time.Time
		return ret
	}
	return *o.BootstrappingInfoCreationTime
}

// GetBootstrappingInfoCreationTimeOk returns a tuple with the BootstrappingInfoCreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoResponse) GetBootstrappingInfoCreationTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.BootstrappingInfoCreationTime) {
    return nil, false
	}
	return o.BootstrappingInfoCreationTime, true
}

// HasBootstrappingInfoCreationTime returns a boolean if a field has been set.
func (o *BootstrappingInfoResponse) HasBootstrappingInfoCreationTime() bool {
	if o != nil && !isNil(o.BootstrappingInfoCreationTime) {
		return true
	}

	return false
}

// SetBootstrappingInfoCreationTime gets a reference to the given time.Time and assigns it to the BootstrappingInfoCreationTime field.
func (o *BootstrappingInfoResponse) SetBootstrappingInfoCreationTime(v time.Time) {
	o.BootstrappingInfoCreationTime = &v
}

// GetUssList returns the UssList field value if set, zero value otherwise.
func (o *BootstrappingInfoResponse) GetUssList() []UssListItem {
	if o == nil || isNil(o.UssList) {
		var ret []UssListItem
		return ret
	}
	return o.UssList
}

// GetUssListOk returns a tuple with the UssList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoResponse) GetUssListOk() ([]UssListItem, bool) {
	if o == nil || isNil(o.UssList) {
    return nil, false
	}
	return o.UssList, true
}

// HasUssList returns a boolean if a field has been set.
func (o *BootstrappingInfoResponse) HasUssList() bool {
	if o != nil && !isNil(o.UssList) {
		return true
	}

	return false
}

// SetUssList gets a reference to the given []UssListItem and assigns it to the UssList field.
func (o *BootstrappingInfoResponse) SetUssList(v []UssListItem) {
	o.UssList = v
}

// GetGbaType returns the GbaType field value if set, zero value otherwise.
func (o *BootstrappingInfoResponse) GetGbaType() GbaType {
	if o == nil || isNil(o.GbaType) {
		var ret GbaType
		return ret
	}
	return *o.GbaType
}

// GetGbaTypeOk returns a tuple with the GbaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoResponse) GetGbaTypeOk() (*GbaType, bool) {
	if o == nil || isNil(o.GbaType) {
    return nil, false
	}
	return o.GbaType, true
}

// HasGbaType returns a boolean if a field has been set.
func (o *BootstrappingInfoResponse) HasGbaType() bool {
	if o != nil && !isNil(o.GbaType) {
		return true
	}

	return false
}

// SetGbaType gets a reference to the given GbaType and assigns it to the GbaType field.
func (o *BootstrappingInfoResponse) SetGbaType(v GbaType) {
	o.GbaType = &v
}

// GetImpi returns the Impi field value if set, zero value otherwise.
func (o *BootstrappingInfoResponse) GetImpi() string {
	if o == nil || isNil(o.Impi) {
		var ret string
		return ret
	}
	return *o.Impi
}

// GetImpiOk returns a tuple with the Impi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BootstrappingInfoResponse) GetImpiOk() (*string, bool) {
	if o == nil || isNil(o.Impi) {
    return nil, false
	}
	return o.Impi, true
}

// HasImpi returns a boolean if a field has been set.
func (o *BootstrappingInfoResponse) HasImpi() bool {
	if o != nil && !isNil(o.Impi) {
		return true
	}

	return false
}

// SetImpi gets a reference to the given string and assigns it to the Impi field.
func (o *BootstrappingInfoResponse) SetImpi(v string) {
	o.Impi = &v
}

func (o BootstrappingInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["meKeyMaterial"] = o.MeKeyMaterial
	}
	if !isNil(o.UiccKeyMaterial) {
		toSerialize["uiccKeyMaterial"] = o.UiccKeyMaterial
	}
	if !isNil(o.KeyExpiryTime) {
		toSerialize["keyExpiryTime"] = o.KeyExpiryTime
	}
	if !isNil(o.BootstrappingInfoCreationTime) {
		toSerialize["bootstrappingInfoCreationTime"] = o.BootstrappingInfoCreationTime
	}
	if !isNil(o.UssList) {
		toSerialize["ussList"] = o.UssList
	}
	if !isNil(o.GbaType) {
		toSerialize["gbaType"] = o.GbaType
	}
	if !isNil(o.Impi) {
		toSerialize["impi"] = o.Impi
	}
	return json.Marshal(toSerialize)
}

type NullableBootstrappingInfoResponse struct {
	value *BootstrappingInfoResponse
	isSet bool
}

func (v NullableBootstrappingInfoResponse) Get() *BootstrappingInfoResponse {
	return v.value
}

func (v *NullableBootstrappingInfoResponse) Set(val *BootstrappingInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBootstrappingInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBootstrappingInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootstrappingInfoResponse(val *BootstrappingInfoResponse) *NullableBootstrappingInfoResponse {
	return &NullableBootstrappingInfoResponse{value: val, isSet: true}
}

func (v NullableBootstrappingInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootstrappingInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


