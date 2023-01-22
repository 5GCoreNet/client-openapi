/*
Nnef_Authentication

NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nnef_Authentication

import (
	"encoding/json"
)

// UAVAuthInfo UAV auth data
type UAVAuthInfo struct {
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi"`
	ServiceLevelId string `json:"serviceLevelId"`
	// String providing an URI formatted according to RFC 3986.
	AuthNotificationURI *string `json:"authNotificationURI,omitempty"`
	IpAddr *IpAddr `json:"ipAddr,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	Pei *string `json:"pei,omitempty"`
	AuthServerAddress *string `json:"authServerAddress,omitempty"`
	AuthMsg *RefToBinaryData `json:"authMsg,omitempty"`
	UeLocInfo *UserLocation `json:"ueLocInfo,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	SNssai *ExtSnssai `json:"sNssai,omitempty"`
	NfType NFType `json:"nfType"`
}

// NewUAVAuthInfo instantiates a new UAVAuthInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUAVAuthInfo(gpsi string, serviceLevelId string, nfType NFType) *UAVAuthInfo {
	this := UAVAuthInfo{}
	this.Gpsi = gpsi
	this.ServiceLevelId = serviceLevelId
	this.NfType = nfType
	return &this
}

// NewUAVAuthInfoWithDefaults instantiates a new UAVAuthInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUAVAuthInfoWithDefaults() *UAVAuthInfo {
	this := UAVAuthInfo{}
	return &this
}

// GetGpsi returns the Gpsi field value
func (o *UAVAuthInfo) GetGpsi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetGpsiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Gpsi, true
}

// SetGpsi sets field value
func (o *UAVAuthInfo) SetGpsi(v string) {
	o.Gpsi = v
}

// GetServiceLevelId returns the ServiceLevelId field value
func (o *UAVAuthInfo) GetServiceLevelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceLevelId
}

// GetServiceLevelIdOk returns a tuple with the ServiceLevelId field value
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetServiceLevelIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ServiceLevelId, true
}

// SetServiceLevelId sets field value
func (o *UAVAuthInfo) SetServiceLevelId(v string) {
	o.ServiceLevelId = v
}

// GetAuthNotificationURI returns the AuthNotificationURI field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetAuthNotificationURI() string {
	if o == nil || isNil(o.AuthNotificationURI) {
		var ret string
		return ret
	}
	return *o.AuthNotificationURI
}

// GetAuthNotificationURIOk returns a tuple with the AuthNotificationURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetAuthNotificationURIOk() (*string, bool) {
	if o == nil || isNil(o.AuthNotificationURI) {
    return nil, false
	}
	return o.AuthNotificationURI, true
}

// HasAuthNotificationURI returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasAuthNotificationURI() bool {
	if o != nil && !isNil(o.AuthNotificationURI) {
		return true
	}

	return false
}

// SetAuthNotificationURI gets a reference to the given string and assigns it to the AuthNotificationURI field.
func (o *UAVAuthInfo) SetAuthNotificationURI(v string) {
	o.AuthNotificationURI = &v
}

// GetIpAddr returns the IpAddr field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetIpAddr() IpAddr {
	if o == nil || isNil(o.IpAddr) {
		var ret IpAddr
		return ret
	}
	return *o.IpAddr
}

// GetIpAddrOk returns a tuple with the IpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetIpAddrOk() (*IpAddr, bool) {
	if o == nil || isNil(o.IpAddr) {
    return nil, false
	}
	return o.IpAddr, true
}

// HasIpAddr returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasIpAddr() bool {
	if o != nil && !isNil(o.IpAddr) {
		return true
	}

	return false
}

// SetIpAddr gets a reference to the given IpAddr and assigns it to the IpAddr field.
func (o *UAVAuthInfo) SetIpAddr(v IpAddr) {
	o.IpAddr = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetPei() string {
	if o == nil || isNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetPeiOk() (*string, bool) {
	if o == nil || isNil(o.Pei) {
    return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasPei() bool {
	if o != nil && !isNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *UAVAuthInfo) SetPei(v string) {
	o.Pei = &v
}

// GetAuthServerAddress returns the AuthServerAddress field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetAuthServerAddress() string {
	if o == nil || isNil(o.AuthServerAddress) {
		var ret string
		return ret
	}
	return *o.AuthServerAddress
}

// GetAuthServerAddressOk returns a tuple with the AuthServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetAuthServerAddressOk() (*string, bool) {
	if o == nil || isNil(o.AuthServerAddress) {
    return nil, false
	}
	return o.AuthServerAddress, true
}

// HasAuthServerAddress returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasAuthServerAddress() bool {
	if o != nil && !isNil(o.AuthServerAddress) {
		return true
	}

	return false
}

// SetAuthServerAddress gets a reference to the given string and assigns it to the AuthServerAddress field.
func (o *UAVAuthInfo) SetAuthServerAddress(v string) {
	o.AuthServerAddress = &v
}

// GetAuthMsg returns the AuthMsg field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetAuthMsg() RefToBinaryData {
	if o == nil || isNil(o.AuthMsg) {
		var ret RefToBinaryData
		return ret
	}
	return *o.AuthMsg
}

// GetAuthMsgOk returns a tuple with the AuthMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetAuthMsgOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.AuthMsg) {
    return nil, false
	}
	return o.AuthMsg, true
}

// HasAuthMsg returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasAuthMsg() bool {
	if o != nil && !isNil(o.AuthMsg) {
		return true
	}

	return false
}

// SetAuthMsg gets a reference to the given RefToBinaryData and assigns it to the AuthMsg field.
func (o *UAVAuthInfo) SetAuthMsg(v RefToBinaryData) {
	o.AuthMsg = &v
}

// GetUeLocInfo returns the UeLocInfo field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetUeLocInfo() UserLocation {
	if o == nil || isNil(o.UeLocInfo) {
		var ret UserLocation
		return ret
	}
	return *o.UeLocInfo
}

// GetUeLocInfoOk returns a tuple with the UeLocInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetUeLocInfoOk() (*UserLocation, bool) {
	if o == nil || isNil(o.UeLocInfo) {
    return nil, false
	}
	return o.UeLocInfo, true
}

// HasUeLocInfo returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasUeLocInfo() bool {
	if o != nil && !isNil(o.UeLocInfo) {
		return true
	}

	return false
}

// SetUeLocInfo gets a reference to the given UserLocation and assigns it to the UeLocInfo field.
func (o *UAVAuthInfo) SetUeLocInfo(v UserLocation) {
	o.UeLocInfo = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *UAVAuthInfo) SetDnn(v string) {
	o.Dnn = &v
}

// GetSNssai returns the SNssai field value if set, zero value otherwise.
func (o *UAVAuthInfo) GetSNssai() ExtSnssai {
	if o == nil || isNil(o.SNssai) {
		var ret ExtSnssai
		return ret
	}
	return *o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil || isNil(o.SNssai) {
    return nil, false
	}
	return o.SNssai, true
}

// HasSNssai returns a boolean if a field has been set.
func (o *UAVAuthInfo) HasSNssai() bool {
	if o != nil && !isNil(o.SNssai) {
		return true
	}

	return false
}

// SetSNssai gets a reference to the given ExtSnssai and assigns it to the SNssai field.
func (o *UAVAuthInfo) SetSNssai(v ExtSnssai) {
	o.SNssai = &v
}

// GetNfType returns the NfType field value
func (o *UAVAuthInfo) GetNfType() NFType {
	if o == nil {
		var ret NFType
		return ret
	}

	return o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value
// and a boolean to check if the value has been set.
func (o *UAVAuthInfo) GetNfTypeOk() (*NFType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NfType, true
}

// SetNfType sets field value
func (o *UAVAuthInfo) SetNfType(v NFType) {
	o.NfType = v
}

func (o UAVAuthInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gpsi"] = o.Gpsi
	}
	if true {
		toSerialize["serviceLevelId"] = o.ServiceLevelId
	}
	if !isNil(o.AuthNotificationURI) {
		toSerialize["authNotificationURI"] = o.AuthNotificationURI
	}
	if !isNil(o.IpAddr) {
		toSerialize["ipAddr"] = o.IpAddr
	}
	if !isNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if !isNil(o.AuthServerAddress) {
		toSerialize["authServerAddress"] = o.AuthServerAddress
	}
	if !isNil(o.AuthMsg) {
		toSerialize["authMsg"] = o.AuthMsg
	}
	if !isNil(o.UeLocInfo) {
		toSerialize["ueLocInfo"] = o.UeLocInfo
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.SNssai) {
		toSerialize["sNssai"] = o.SNssai
	}
	if true {
		toSerialize["nfType"] = o.NfType
	}
	return json.Marshal(toSerialize)
}

type NullableUAVAuthInfo struct {
	value *UAVAuthInfo
	isSet bool
}

func (v NullableUAVAuthInfo) Get() *UAVAuthInfo {
	return v.value
}

func (v *NullableUAVAuthInfo) Set(val *UAVAuthInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUAVAuthInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUAVAuthInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUAVAuthInfo(val *UAVAuthInfo) *NullableUAVAuthInfo {
	return &NullableUAVAuthInfo{value: val, isSet: true}
}

func (v NullableUAVAuthInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUAVAuthInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


