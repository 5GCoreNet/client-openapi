/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
)

// N3gaLocation Contains the Non-3GPP access user location.
type N3gaLocation struct {
	N3gppTai *Tai `json:"n3gppTai,omitempty"`
	// This IE shall contain the N3IWF identifier received over NGAP and shall be encoded as a  string of hexadecimal characters. Each character in the string shall take a value of \"0\"  to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant  character representing the 4 most significant bits of the N3IWF ID shall appear first in  the string, and the character representing the 4 least significant bit of the N3IWF ID  shall appear last in the string.  
	N3IwfId *string `json:"n3IwfId,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	UeIpv4Addr *string `json:"ueIpv4Addr,omitempty"`
	UeIpv6Addr *Ipv6Addr `json:"ueIpv6Addr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber *int32 `json:"portNumber,omitempty"`
	Protocol *TransportProtocol `json:"protocol,omitempty"`
	TnapId *TnapId `json:"tnapId,omitempty"`
	TwapId *TwapId `json:"twapId,omitempty"`
	HfcNodeId *HfcNodeId `json:"hfcNodeId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	Gli *string `json:"gli,omitempty"`
	W5gbanLineType *LineType `json:"w5gbanLineType,omitempty"`
	// Global Cable Identifier uniquely identifying the connection between the 5G-CRG or FN-CRG to the 5GS. See clause 28.15.4 of 3GPP TS 23.003. This shall be encoded as a string per clause 28.15.4 of 3GPP TS 23.003, and compliant with the syntax specified  in clause 2.2  of IETF RFC 7542 for the username part of a NAI. The GCI value is specified in CableLabs WR-TR-5WWC-ARCH. 
	Gci *string `json:"gci,omitempty"`
}

// NewN3gaLocation instantiates a new N3gaLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN3gaLocation() *N3gaLocation {
	this := N3gaLocation{}
	return &this
}

// NewN3gaLocationWithDefaults instantiates a new N3gaLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN3gaLocationWithDefaults() *N3gaLocation {
	this := N3gaLocation{}
	return &this
}

// GetN3gppTai returns the N3gppTai field value if set, zero value otherwise.
func (o *N3gaLocation) GetN3gppTai() Tai {
	if o == nil || isNil(o.N3gppTai) {
		var ret Tai
		return ret
	}
	return *o.N3gppTai
}

// GetN3gppTaiOk returns a tuple with the N3gppTai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetN3gppTaiOk() (*Tai, bool) {
	if o == nil || isNil(o.N3gppTai) {
    return nil, false
	}
	return o.N3gppTai, true
}

// HasN3gppTai returns a boolean if a field has been set.
func (o *N3gaLocation) HasN3gppTai() bool {
	if o != nil && !isNil(o.N3gppTai) {
		return true
	}

	return false
}

// SetN3gppTai gets a reference to the given Tai and assigns it to the N3gppTai field.
func (o *N3gaLocation) SetN3gppTai(v Tai) {
	o.N3gppTai = &v
}

// GetN3IwfId returns the N3IwfId field value if set, zero value otherwise.
func (o *N3gaLocation) GetN3IwfId() string {
	if o == nil || isNil(o.N3IwfId) {
		var ret string
		return ret
	}
	return *o.N3IwfId
}

// GetN3IwfIdOk returns a tuple with the N3IwfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetN3IwfIdOk() (*string, bool) {
	if o == nil || isNil(o.N3IwfId) {
    return nil, false
	}
	return o.N3IwfId, true
}

// HasN3IwfId returns a boolean if a field has been set.
func (o *N3gaLocation) HasN3IwfId() bool {
	if o != nil && !isNil(o.N3IwfId) {
		return true
	}

	return false
}

// SetN3IwfId gets a reference to the given string and assigns it to the N3IwfId field.
func (o *N3gaLocation) SetN3IwfId(v string) {
	o.N3IwfId = &v
}

// GetUeIpv4Addr returns the UeIpv4Addr field value if set, zero value otherwise.
func (o *N3gaLocation) GetUeIpv4Addr() string {
	if o == nil || isNil(o.UeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.UeIpv4Addr
}

// GetUeIpv4AddrOk returns a tuple with the UeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetUeIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.UeIpv4Addr) {
    return nil, false
	}
	return o.UeIpv4Addr, true
}

// HasUeIpv4Addr returns a boolean if a field has been set.
func (o *N3gaLocation) HasUeIpv4Addr() bool {
	if o != nil && !isNil(o.UeIpv4Addr) {
		return true
	}

	return false
}

// SetUeIpv4Addr gets a reference to the given string and assigns it to the UeIpv4Addr field.
func (o *N3gaLocation) SetUeIpv4Addr(v string) {
	o.UeIpv4Addr = &v
}

// GetUeIpv6Addr returns the UeIpv6Addr field value if set, zero value otherwise.
func (o *N3gaLocation) GetUeIpv6Addr() Ipv6Addr {
	if o == nil || isNil(o.UeIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.UeIpv6Addr
}

// GetUeIpv6AddrOk returns a tuple with the UeIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetUeIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.UeIpv6Addr) {
    return nil, false
	}
	return o.UeIpv6Addr, true
}

// HasUeIpv6Addr returns a boolean if a field has been set.
func (o *N3gaLocation) HasUeIpv6Addr() bool {
	if o != nil && !isNil(o.UeIpv6Addr) {
		return true
	}

	return false
}

// SetUeIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the UeIpv6Addr field.
func (o *N3gaLocation) SetUeIpv6Addr(v Ipv6Addr) {
	o.UeIpv6Addr = &v
}

// GetPortNumber returns the PortNumber field value if set, zero value otherwise.
func (o *N3gaLocation) GetPortNumber() int32 {
	if o == nil || isNil(o.PortNumber) {
		var ret int32
		return ret
	}
	return *o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetPortNumberOk() (*int32, bool) {
	if o == nil || isNil(o.PortNumber) {
    return nil, false
	}
	return o.PortNumber, true
}

// HasPortNumber returns a boolean if a field has been set.
func (o *N3gaLocation) HasPortNumber() bool {
	if o != nil && !isNil(o.PortNumber) {
		return true
	}

	return false
}

// SetPortNumber gets a reference to the given int32 and assigns it to the PortNumber field.
func (o *N3gaLocation) SetPortNumber(v int32) {
	o.PortNumber = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *N3gaLocation) GetProtocol() TransportProtocol {
	if o == nil || isNil(o.Protocol) {
		var ret TransportProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetProtocolOk() (*TransportProtocol, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *N3gaLocation) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given TransportProtocol and assigns it to the Protocol field.
func (o *N3gaLocation) SetProtocol(v TransportProtocol) {
	o.Protocol = &v
}

// GetTnapId returns the TnapId field value if set, zero value otherwise.
func (o *N3gaLocation) GetTnapId() TnapId {
	if o == nil || isNil(o.TnapId) {
		var ret TnapId
		return ret
	}
	return *o.TnapId
}

// GetTnapIdOk returns a tuple with the TnapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetTnapIdOk() (*TnapId, bool) {
	if o == nil || isNil(o.TnapId) {
    return nil, false
	}
	return o.TnapId, true
}

// HasTnapId returns a boolean if a field has been set.
func (o *N3gaLocation) HasTnapId() bool {
	if o != nil && !isNil(o.TnapId) {
		return true
	}

	return false
}

// SetTnapId gets a reference to the given TnapId and assigns it to the TnapId field.
func (o *N3gaLocation) SetTnapId(v TnapId) {
	o.TnapId = &v
}

// GetTwapId returns the TwapId field value if set, zero value otherwise.
func (o *N3gaLocation) GetTwapId() TwapId {
	if o == nil || isNil(o.TwapId) {
		var ret TwapId
		return ret
	}
	return *o.TwapId
}

// GetTwapIdOk returns a tuple with the TwapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetTwapIdOk() (*TwapId, bool) {
	if o == nil || isNil(o.TwapId) {
    return nil, false
	}
	return o.TwapId, true
}

// HasTwapId returns a boolean if a field has been set.
func (o *N3gaLocation) HasTwapId() bool {
	if o != nil && !isNil(o.TwapId) {
		return true
	}

	return false
}

// SetTwapId gets a reference to the given TwapId and assigns it to the TwapId field.
func (o *N3gaLocation) SetTwapId(v TwapId) {
	o.TwapId = &v
}

// GetHfcNodeId returns the HfcNodeId field value if set, zero value otherwise.
func (o *N3gaLocation) GetHfcNodeId() HfcNodeId {
	if o == nil || isNil(o.HfcNodeId) {
		var ret HfcNodeId
		return ret
	}
	return *o.HfcNodeId
}

// GetHfcNodeIdOk returns a tuple with the HfcNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetHfcNodeIdOk() (*HfcNodeId, bool) {
	if o == nil || isNil(o.HfcNodeId) {
    return nil, false
	}
	return o.HfcNodeId, true
}

// HasHfcNodeId returns a boolean if a field has been set.
func (o *N3gaLocation) HasHfcNodeId() bool {
	if o != nil && !isNil(o.HfcNodeId) {
		return true
	}

	return false
}

// SetHfcNodeId gets a reference to the given HfcNodeId and assigns it to the HfcNodeId field.
func (o *N3gaLocation) SetHfcNodeId(v HfcNodeId) {
	o.HfcNodeId = &v
}

// GetGli returns the Gli field value if set, zero value otherwise.
func (o *N3gaLocation) GetGli() string {
	if o == nil || isNil(o.Gli) {
		var ret string
		return ret
	}
	return *o.Gli
}

// GetGliOk returns a tuple with the Gli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetGliOk() (*string, bool) {
	if o == nil || isNil(o.Gli) {
    return nil, false
	}
	return o.Gli, true
}

// HasGli returns a boolean if a field has been set.
func (o *N3gaLocation) HasGli() bool {
	if o != nil && !isNil(o.Gli) {
		return true
	}

	return false
}

// SetGli gets a reference to the given string and assigns it to the Gli field.
func (o *N3gaLocation) SetGli(v string) {
	o.Gli = &v
}

// GetW5gbanLineType returns the W5gbanLineType field value if set, zero value otherwise.
func (o *N3gaLocation) GetW5gbanLineType() LineType {
	if o == nil || isNil(o.W5gbanLineType) {
		var ret LineType
		return ret
	}
	return *o.W5gbanLineType
}

// GetW5gbanLineTypeOk returns a tuple with the W5gbanLineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetW5gbanLineTypeOk() (*LineType, bool) {
	if o == nil || isNil(o.W5gbanLineType) {
    return nil, false
	}
	return o.W5gbanLineType, true
}

// HasW5gbanLineType returns a boolean if a field has been set.
func (o *N3gaLocation) HasW5gbanLineType() bool {
	if o != nil && !isNil(o.W5gbanLineType) {
		return true
	}

	return false
}

// SetW5gbanLineType gets a reference to the given LineType and assigns it to the W5gbanLineType field.
func (o *N3gaLocation) SetW5gbanLineType(v LineType) {
	o.W5gbanLineType = &v
}

// GetGci returns the Gci field value if set, zero value otherwise.
func (o *N3gaLocation) GetGci() string {
	if o == nil || isNil(o.Gci) {
		var ret string
		return ret
	}
	return *o.Gci
}

// GetGciOk returns a tuple with the Gci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gaLocation) GetGciOk() (*string, bool) {
	if o == nil || isNil(o.Gci) {
    return nil, false
	}
	return o.Gci, true
}

// HasGci returns a boolean if a field has been set.
func (o *N3gaLocation) HasGci() bool {
	if o != nil && !isNil(o.Gci) {
		return true
	}

	return false
}

// SetGci gets a reference to the given string and assigns it to the Gci field.
func (o *N3gaLocation) SetGci(v string) {
	o.Gci = &v
}

func (o N3gaLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.N3gppTai) {
		toSerialize["n3gppTai"] = o.N3gppTai
	}
	if !isNil(o.N3IwfId) {
		toSerialize["n3IwfId"] = o.N3IwfId
	}
	if !isNil(o.UeIpv4Addr) {
		toSerialize["ueIpv4Addr"] = o.UeIpv4Addr
	}
	if !isNil(o.UeIpv6Addr) {
		toSerialize["ueIpv6Addr"] = o.UeIpv6Addr
	}
	if !isNil(o.PortNumber) {
		toSerialize["portNumber"] = o.PortNumber
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.TnapId) {
		toSerialize["tnapId"] = o.TnapId
	}
	if !isNil(o.TwapId) {
		toSerialize["twapId"] = o.TwapId
	}
	if !isNil(o.HfcNodeId) {
		toSerialize["hfcNodeId"] = o.HfcNodeId
	}
	if !isNil(o.Gli) {
		toSerialize["gli"] = o.Gli
	}
	if !isNil(o.W5gbanLineType) {
		toSerialize["w5gbanLineType"] = o.W5gbanLineType
	}
	if !isNil(o.Gci) {
		toSerialize["gci"] = o.Gci
	}
	return json.Marshal(toSerialize)
}

type NullableN3gaLocation struct {
	value *N3gaLocation
	isSet bool
}

func (v NullableN3gaLocation) Get() *N3gaLocation {
	return v.value
}

func (v *NullableN3gaLocation) Set(val *N3gaLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableN3gaLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableN3gaLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN3gaLocation(val *N3gaLocation) *NullableN3gaLocation {
	return &NullableN3gaLocation{value: val, isSet: true}
}

func (v NullableN3gaLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN3gaLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


