/*
Unified Data Repository Service API file for structured data for exposure

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Exposure_Data

import (
	"encoding/json"
	"time"
)

// PduSessionManagementData Represents Session management data for a UE and a PDU session.
type PduSessionManagementData struct {
	PduSessionStatus *PduSessionStatus `json:"pduSessionStatus,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	PduSessionStatusTs *time.Time `json:"pduSessionStatusTs,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	Dnai *string `json:"dnai,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	DnaiTs *time.Time `json:"dnaiTs,omitempty"`
	N6TrafficRoutingInfo []RouteToLocation `json:"n6TrafficRoutingInfo,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	N6TrafficRoutingInfoTs *time.Time `json:"n6TrafficRoutingInfoTs,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr *string `json:"ipv4Addr,omitempty"`
	// UE IPv6 prefix.
	Ipv6Prefix []Ipv6Prefix `json:"ipv6Prefix,omitempty"`
	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`
	PduSessType *PduSessionType `json:"pduSessType,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	IpAddrTs *time.Time `json:"ipAddrTs,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId *int32 `json:"pduSessionId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	ResetIds []string `json:"resetIds,omitempty"`
}

// NewPduSessionManagementData instantiates a new PduSessionManagementData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPduSessionManagementData() *PduSessionManagementData {
	this := PduSessionManagementData{}
	return &this
}

// NewPduSessionManagementDataWithDefaults instantiates a new PduSessionManagementData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPduSessionManagementDataWithDefaults() *PduSessionManagementData {
	this := PduSessionManagementData{}
	return &this
}

// GetPduSessionStatus returns the PduSessionStatus field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetPduSessionStatus() PduSessionStatus {
	if o == nil || isNil(o.PduSessionStatus) {
		var ret PduSessionStatus
		return ret
	}
	return *o.PduSessionStatus
}

// GetPduSessionStatusOk returns a tuple with the PduSessionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetPduSessionStatusOk() (*PduSessionStatus, bool) {
	if o == nil || isNil(o.PduSessionStatus) {
    return nil, false
	}
	return o.PduSessionStatus, true
}

// HasPduSessionStatus returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasPduSessionStatus() bool {
	if o != nil && !isNil(o.PduSessionStatus) {
		return true
	}

	return false
}

// SetPduSessionStatus gets a reference to the given PduSessionStatus and assigns it to the PduSessionStatus field.
func (o *PduSessionManagementData) SetPduSessionStatus(v PduSessionStatus) {
	o.PduSessionStatus = &v
}

// GetPduSessionStatusTs returns the PduSessionStatusTs field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetPduSessionStatusTs() time.Time {
	if o == nil || isNil(o.PduSessionStatusTs) {
		var ret time.Time
		return ret
	}
	return *o.PduSessionStatusTs
}

// GetPduSessionStatusTsOk returns a tuple with the PduSessionStatusTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetPduSessionStatusTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.PduSessionStatusTs) {
    return nil, false
	}
	return o.PduSessionStatusTs, true
}

// HasPduSessionStatusTs returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasPduSessionStatusTs() bool {
	if o != nil && !isNil(o.PduSessionStatusTs) {
		return true
	}

	return false
}

// SetPduSessionStatusTs gets a reference to the given time.Time and assigns it to the PduSessionStatusTs field.
func (o *PduSessionManagementData) SetPduSessionStatusTs(v time.Time) {
	o.PduSessionStatusTs = &v
}

// GetDnai returns the Dnai field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetDnai() string {
	if o == nil || isNil(o.Dnai) {
		var ret string
		return ret
	}
	return *o.Dnai
}

// GetDnaiOk returns a tuple with the Dnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetDnaiOk() (*string, bool) {
	if o == nil || isNil(o.Dnai) {
    return nil, false
	}
	return o.Dnai, true
}

// HasDnai returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasDnai() bool {
	if o != nil && !isNil(o.Dnai) {
		return true
	}

	return false
}

// SetDnai gets a reference to the given string and assigns it to the Dnai field.
func (o *PduSessionManagementData) SetDnai(v string) {
	o.Dnai = &v
}

// GetDnaiTs returns the DnaiTs field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetDnaiTs() time.Time {
	if o == nil || isNil(o.DnaiTs) {
		var ret time.Time
		return ret
	}
	return *o.DnaiTs
}

// GetDnaiTsOk returns a tuple with the DnaiTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetDnaiTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.DnaiTs) {
    return nil, false
	}
	return o.DnaiTs, true
}

// HasDnaiTs returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasDnaiTs() bool {
	if o != nil && !isNil(o.DnaiTs) {
		return true
	}

	return false
}

// SetDnaiTs gets a reference to the given time.Time and assigns it to the DnaiTs field.
func (o *PduSessionManagementData) SetDnaiTs(v time.Time) {
	o.DnaiTs = &v
}

// GetN6TrafficRoutingInfo returns the N6TrafficRoutingInfo field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetN6TrafficRoutingInfo() []RouteToLocation {
	if o == nil || isNil(o.N6TrafficRoutingInfo) {
		var ret []RouteToLocation
		return ret
	}
	return o.N6TrafficRoutingInfo
}

// GetN6TrafficRoutingInfoOk returns a tuple with the N6TrafficRoutingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetN6TrafficRoutingInfoOk() ([]RouteToLocation, bool) {
	if o == nil || isNil(o.N6TrafficRoutingInfo) {
    return nil, false
	}
	return o.N6TrafficRoutingInfo, true
}

// HasN6TrafficRoutingInfo returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasN6TrafficRoutingInfo() bool {
	if o != nil && !isNil(o.N6TrafficRoutingInfo) {
		return true
	}

	return false
}

// SetN6TrafficRoutingInfo gets a reference to the given []RouteToLocation and assigns it to the N6TrafficRoutingInfo field.
func (o *PduSessionManagementData) SetN6TrafficRoutingInfo(v []RouteToLocation) {
	o.N6TrafficRoutingInfo = v
}

// GetN6TrafficRoutingInfoTs returns the N6TrafficRoutingInfoTs field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetN6TrafficRoutingInfoTs() time.Time {
	if o == nil || isNil(o.N6TrafficRoutingInfoTs) {
		var ret time.Time
		return ret
	}
	return *o.N6TrafficRoutingInfoTs
}

// GetN6TrafficRoutingInfoTsOk returns a tuple with the N6TrafficRoutingInfoTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetN6TrafficRoutingInfoTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.N6TrafficRoutingInfoTs) {
    return nil, false
	}
	return o.N6TrafficRoutingInfoTs, true
}

// HasN6TrafficRoutingInfoTs returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasN6TrafficRoutingInfoTs() bool {
	if o != nil && !isNil(o.N6TrafficRoutingInfoTs) {
		return true
	}

	return false
}

// SetN6TrafficRoutingInfoTs gets a reference to the given time.Time and assigns it to the N6TrafficRoutingInfoTs field.
func (o *PduSessionManagementData) SetN6TrafficRoutingInfoTs(v time.Time) {
	o.N6TrafficRoutingInfoTs = &v
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetIpv4Addr() string {
	if o == nil || isNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.Ipv4Addr) {
    return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasIpv4Addr() bool {
	if o != nil && !isNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *PduSessionManagementData) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Prefix returns the Ipv6Prefix field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetIpv6Prefix() []Ipv6Prefix {
	if o == nil || isNil(o.Ipv6Prefix) {
		var ret []Ipv6Prefix
		return ret
	}
	return o.Ipv6Prefix
}

// GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetIpv6PrefixOk() ([]Ipv6Prefix, bool) {
	if o == nil || isNil(o.Ipv6Prefix) {
    return nil, false
	}
	return o.Ipv6Prefix, true
}

// HasIpv6Prefix returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasIpv6Prefix() bool {
	if o != nil && !isNil(o.Ipv6Prefix) {
		return true
	}

	return false
}

// SetIpv6Prefix gets a reference to the given []Ipv6Prefix and assigns it to the Ipv6Prefix field.
func (o *PduSessionManagementData) SetIpv6Prefix(v []Ipv6Prefix) {
	o.Ipv6Prefix = v
}

// GetIpv6Addrs returns the Ipv6Addrs field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetIpv6Addrs() []Ipv6Addr {
	if o == nil || isNil(o.Ipv6Addrs) {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6Addrs
}

// GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || isNil(o.Ipv6Addrs) {
    return nil, false
	}
	return o.Ipv6Addrs, true
}

// HasIpv6Addrs returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasIpv6Addrs() bool {
	if o != nil && !isNil(o.Ipv6Addrs) {
		return true
	}

	return false
}

// SetIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the Ipv6Addrs field.
func (o *PduSessionManagementData) SetIpv6Addrs(v []Ipv6Addr) {
	o.Ipv6Addrs = v
}

// GetPduSessType returns the PduSessType field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetPduSessType() PduSessionType {
	if o == nil || isNil(o.PduSessType) {
		var ret PduSessionType
		return ret
	}
	return *o.PduSessType
}

// GetPduSessTypeOk returns a tuple with the PduSessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetPduSessTypeOk() (*PduSessionType, bool) {
	if o == nil || isNil(o.PduSessType) {
    return nil, false
	}
	return o.PduSessType, true
}

// HasPduSessType returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasPduSessType() bool {
	if o != nil && !isNil(o.PduSessType) {
		return true
	}

	return false
}

// SetPduSessType gets a reference to the given PduSessionType and assigns it to the PduSessType field.
func (o *PduSessionManagementData) SetPduSessType(v PduSessionType) {
	o.PduSessType = &v
}

// GetIpAddrTs returns the IpAddrTs field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetIpAddrTs() time.Time {
	if o == nil || isNil(o.IpAddrTs) {
		var ret time.Time
		return ret
	}
	return *o.IpAddrTs
}

// GetIpAddrTsOk returns a tuple with the IpAddrTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetIpAddrTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.IpAddrTs) {
    return nil, false
	}
	return o.IpAddrTs, true
}

// HasIpAddrTs returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasIpAddrTs() bool {
	if o != nil && !isNil(o.IpAddrTs) {
		return true
	}

	return false
}

// SetIpAddrTs gets a reference to the given time.Time and assigns it to the IpAddrTs field.
func (o *PduSessionManagementData) SetIpAddrTs(v time.Time) {
	o.IpAddrTs = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetDnn() string {
	if o == nil || isNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetDnnOk() (*string, bool) {
	if o == nil || isNil(o.Dnn) {
    return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasDnn() bool {
	if o != nil && !isNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PduSessionManagementData) SetDnn(v string) {
	o.Dnn = &v
}

// GetPduSessionId returns the PduSessionId field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetPduSessionId() int32 {
	if o == nil || isNil(o.PduSessionId) {
		var ret int32
		return ret
	}
	return *o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil || isNil(o.PduSessionId) {
    return nil, false
	}
	return o.PduSessionId, true
}

// HasPduSessionId returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasPduSessionId() bool {
	if o != nil && !isNil(o.PduSessionId) {
		return true
	}

	return false
}

// SetPduSessionId gets a reference to the given int32 and assigns it to the PduSessionId field.
func (o *PduSessionManagementData) SetPduSessionId(v int32) {
	o.PduSessionId = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *PduSessionManagementData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetResetIds returns the ResetIds field value if set, zero value otherwise.
func (o *PduSessionManagementData) GetResetIds() []string {
	if o == nil || isNil(o.ResetIds) {
		var ret []string
		return ret
	}
	return o.ResetIds
}

// GetResetIdsOk returns a tuple with the ResetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PduSessionManagementData) GetResetIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ResetIds) {
    return nil, false
	}
	return o.ResetIds, true
}

// HasResetIds returns a boolean if a field has been set.
func (o *PduSessionManagementData) HasResetIds() bool {
	if o != nil && !isNil(o.ResetIds) {
		return true
	}

	return false
}

// SetResetIds gets a reference to the given []string and assigns it to the ResetIds field.
func (o *PduSessionManagementData) SetResetIds(v []string) {
	o.ResetIds = v
}

func (o PduSessionManagementData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PduSessionStatus) {
		toSerialize["pduSessionStatus"] = o.PduSessionStatus
	}
	if !isNil(o.PduSessionStatusTs) {
		toSerialize["pduSessionStatusTs"] = o.PduSessionStatusTs
	}
	if !isNil(o.Dnai) {
		toSerialize["dnai"] = o.Dnai
	}
	if !isNil(o.DnaiTs) {
		toSerialize["dnaiTs"] = o.DnaiTs
	}
	if !isNil(o.N6TrafficRoutingInfo) {
		toSerialize["n6TrafficRoutingInfo"] = o.N6TrafficRoutingInfo
	}
	if !isNil(o.N6TrafficRoutingInfoTs) {
		toSerialize["n6TrafficRoutingInfoTs"] = o.N6TrafficRoutingInfoTs
	}
	if !isNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !isNil(o.Ipv6Prefix) {
		toSerialize["ipv6Prefix"] = o.Ipv6Prefix
	}
	if !isNil(o.Ipv6Addrs) {
		toSerialize["ipv6Addrs"] = o.Ipv6Addrs
	}
	if !isNil(o.PduSessType) {
		toSerialize["pduSessType"] = o.PduSessType
	}
	if !isNil(o.IpAddrTs) {
		toSerialize["ipAddrTs"] = o.IpAddrTs
	}
	if !isNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !isNil(o.PduSessionId) {
		toSerialize["pduSessionId"] = o.PduSessionId
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !isNil(o.ResetIds) {
		toSerialize["resetIds"] = o.ResetIds
	}
	return json.Marshal(toSerialize)
}

type NullablePduSessionManagementData struct {
	value *PduSessionManagementData
	isSet bool
}

func (v NullablePduSessionManagementData) Get() *PduSessionManagementData {
	return v.value
}

func (v *NullablePduSessionManagementData) Set(val *PduSessionManagementData) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionManagementData) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionManagementData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionManagementData(val *PduSessionManagementData) *NullablePduSessionManagementData {
	return &NullablePduSessionManagementData{value: val, isSet: true}
}

func (v NullablePduSessionManagementData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionManagementData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


