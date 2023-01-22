/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Namf_Communication

import (
	"encoding/json"
)

// TraceData contains Trace control and configuration parameters.
type TraceData struct {
	// Trace Reference (see 3GPP TS 32.422).It shall be encoded as the concatenation of MCC, MNC and Trace ID as follows: 'MCC'<MNC'-'Trace ID'The Trace ID shall be encoded as a 3 octet string in hexadecimal representation. Each character in the Trace ID string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the Trace ID shall appear first  in the string, and the character representing the 4 least significant bit of the Trace ID shall appear last in the string. 
	TraceRef string `json:"traceRef"`
	TraceDepth TraceDepth `json:"traceDepth"`
	// List of NE Types (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string.Octets shall be coded according to 3GPP TS 32.422. 
	NeTypeList string `json:"neTypeList"`
	// Triggering events (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422. 
	EventList string `json:"eventList"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	CollectionEntityIpv4Addr *string `json:"collectionEntityIpv4Addr,omitempty"`
	CollectionEntityIpv6Addr *Ipv6Addr `json:"collectionEntityIpv6Addr,omitempty"`
	// List of Interfaces (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the  4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422. If this attribute is not present, all the interfaces applicable to the list of NE types indicated in the neTypeList attribute should be traced. 
	InterfaceList *string `json:"interfaceList,omitempty"`
}

// NewTraceData instantiates a new TraceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTraceData(traceRef string, traceDepth TraceDepth, neTypeList string, eventList string) *TraceData {
	this := TraceData{}
	this.TraceRef = traceRef
	this.TraceDepth = traceDepth
	this.NeTypeList = neTypeList
	this.EventList = eventList
	return &this
}

// NewTraceDataWithDefaults instantiates a new TraceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceDataWithDefaults() *TraceData {
	this := TraceData{}
	return &this
}

// GetTraceRef returns the TraceRef field value
func (o *TraceData) GetTraceRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceRef
}

// GetTraceRefOk returns a tuple with the TraceRef field value
// and a boolean to check if the value has been set.
func (o *TraceData) GetTraceRefOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TraceRef, true
}

// SetTraceRef sets field value
func (o *TraceData) SetTraceRef(v string) {
	o.TraceRef = v
}

// GetTraceDepth returns the TraceDepth field value
func (o *TraceData) GetTraceDepth() TraceDepth {
	if o == nil {
		var ret TraceDepth
		return ret
	}

	return o.TraceDepth
}

// GetTraceDepthOk returns a tuple with the TraceDepth field value
// and a boolean to check if the value has been set.
func (o *TraceData) GetTraceDepthOk() (*TraceDepth, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TraceDepth, true
}

// SetTraceDepth sets field value
func (o *TraceData) SetTraceDepth(v TraceDepth) {
	o.TraceDepth = v
}

// GetNeTypeList returns the NeTypeList field value
func (o *TraceData) GetNeTypeList() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NeTypeList
}

// GetNeTypeListOk returns a tuple with the NeTypeList field value
// and a boolean to check if the value has been set.
func (o *TraceData) GetNeTypeListOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NeTypeList, true
}

// SetNeTypeList sets field value
func (o *TraceData) SetNeTypeList(v string) {
	o.NeTypeList = v
}

// GetEventList returns the EventList field value
func (o *TraceData) GetEventList() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value
// and a boolean to check if the value has been set.
func (o *TraceData) GetEventListOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EventList, true
}

// SetEventList sets field value
func (o *TraceData) SetEventList(v string) {
	o.EventList = v
}

// GetCollectionEntityIpv4Addr returns the CollectionEntityIpv4Addr field value if set, zero value otherwise.
func (o *TraceData) GetCollectionEntityIpv4Addr() string {
	if o == nil || isNil(o.CollectionEntityIpv4Addr) {
		var ret string
		return ret
	}
	return *o.CollectionEntityIpv4Addr
}

// GetCollectionEntityIpv4AddrOk returns a tuple with the CollectionEntityIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceData) GetCollectionEntityIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.CollectionEntityIpv4Addr) {
    return nil, false
	}
	return o.CollectionEntityIpv4Addr, true
}

// HasCollectionEntityIpv4Addr returns a boolean if a field has been set.
func (o *TraceData) HasCollectionEntityIpv4Addr() bool {
	if o != nil && !isNil(o.CollectionEntityIpv4Addr) {
		return true
	}

	return false
}

// SetCollectionEntityIpv4Addr gets a reference to the given string and assigns it to the CollectionEntityIpv4Addr field.
func (o *TraceData) SetCollectionEntityIpv4Addr(v string) {
	o.CollectionEntityIpv4Addr = &v
}

// GetCollectionEntityIpv6Addr returns the CollectionEntityIpv6Addr field value if set, zero value otherwise.
func (o *TraceData) GetCollectionEntityIpv6Addr() Ipv6Addr {
	if o == nil || isNil(o.CollectionEntityIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.CollectionEntityIpv6Addr
}

// GetCollectionEntityIpv6AddrOk returns a tuple with the CollectionEntityIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceData) GetCollectionEntityIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.CollectionEntityIpv6Addr) {
    return nil, false
	}
	return o.CollectionEntityIpv6Addr, true
}

// HasCollectionEntityIpv6Addr returns a boolean if a field has been set.
func (o *TraceData) HasCollectionEntityIpv6Addr() bool {
	if o != nil && !isNil(o.CollectionEntityIpv6Addr) {
		return true
	}

	return false
}

// SetCollectionEntityIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the CollectionEntityIpv6Addr field.
func (o *TraceData) SetCollectionEntityIpv6Addr(v Ipv6Addr) {
	o.CollectionEntityIpv6Addr = &v
}

// GetInterfaceList returns the InterfaceList field value if set, zero value otherwise.
func (o *TraceData) GetInterfaceList() string {
	if o == nil || isNil(o.InterfaceList) {
		var ret string
		return ret
	}
	return *o.InterfaceList
}

// GetInterfaceListOk returns a tuple with the InterfaceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TraceData) GetInterfaceListOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceList) {
    return nil, false
	}
	return o.InterfaceList, true
}

// HasInterfaceList returns a boolean if a field has been set.
func (o *TraceData) HasInterfaceList() bool {
	if o != nil && !isNil(o.InterfaceList) {
		return true
	}

	return false
}

// SetInterfaceList gets a reference to the given string and assigns it to the InterfaceList field.
func (o *TraceData) SetInterfaceList(v string) {
	o.InterfaceList = &v
}

func (o TraceData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["traceRef"] = o.TraceRef
	}
	if true {
		toSerialize["traceDepth"] = o.TraceDepth
	}
	if true {
		toSerialize["neTypeList"] = o.NeTypeList
	}
	if true {
		toSerialize["eventList"] = o.EventList
	}
	if !isNil(o.CollectionEntityIpv4Addr) {
		toSerialize["collectionEntityIpv4Addr"] = o.CollectionEntityIpv4Addr
	}
	if !isNil(o.CollectionEntityIpv6Addr) {
		toSerialize["collectionEntityIpv6Addr"] = o.CollectionEntityIpv6Addr
	}
	if !isNil(o.InterfaceList) {
		toSerialize["interfaceList"] = o.InterfaceList
	}
	return json.Marshal(toSerialize)
}

type NullableTraceData struct {
	value *TraceData
	isSet bool
}

func (v NullableTraceData) Get() *TraceData {
	return v.value
}

func (v *NullableTraceData) Set(val *TraceData) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceData) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceData(val *TraceData) *NullableTraceData {
	return &NullableTraceData{value: val, isSet: true}
}

func (v NullableTraceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


