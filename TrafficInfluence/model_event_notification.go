/*
3gpp-traffic-influence

API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TrafficInfluence

import (
	"encoding/json"
)

// EventNotification Represents a traffic influence event notification.
type EventNotification struct {
	// Identifies an NEF Northbound interface transaction, generated by the AF.
	AfTransId *string `json:"afTransId,omitempty"`
	DnaiChgType DnaiChangeType `json:"dnaiChgType"`
	SourceTrafficRoute NullableRouteToLocation `json:"sourceTrafficRoute,omitempty"`
	SubscribedEvent SubscribedEvent `json:"subscribedEvent"`
	TargetTrafficRoute NullableRouteToLocation `json:"targetTrafficRoute,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	SourceDnai *string `json:"sourceDnai,omitempty"`
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	TargetDnai *string `json:"targetDnai,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	SrcUeIpv4Addr *string `json:"srcUeIpv4Addr,omitempty"`
	SrcUeIpv6Prefix *Ipv6Prefix `json:"srcUeIpv6Prefix,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	TgtUeIpv4Addr *string `json:"tgtUeIpv4Addr,omitempty"`
	TgtUeIpv6Prefix *Ipv6Prefix `json:"tgtUeIpv6Prefix,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	UeMac *string `json:"ueMac,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	AfAckUri *string `json:"afAckUri,omitempty"`
}

// NewEventNotification instantiates a new EventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventNotification(dnaiChgType DnaiChangeType, subscribedEvent SubscribedEvent) *EventNotification {
	this := EventNotification{}
	this.DnaiChgType = dnaiChgType
	this.SubscribedEvent = subscribedEvent
	return &this
}

// NewEventNotificationWithDefaults instantiates a new EventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventNotificationWithDefaults() *EventNotification {
	this := EventNotification{}
	return &this
}

// GetAfTransId returns the AfTransId field value if set, zero value otherwise.
func (o *EventNotification) GetAfTransId() string {
	if o == nil || isNil(o.AfTransId) {
		var ret string
		return ret
	}
	return *o.AfTransId
}

// GetAfTransIdOk returns a tuple with the AfTransId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetAfTransIdOk() (*string, bool) {
	if o == nil || isNil(o.AfTransId) {
    return nil, false
	}
	return o.AfTransId, true
}

// HasAfTransId returns a boolean if a field has been set.
func (o *EventNotification) HasAfTransId() bool {
	if o != nil && !isNil(o.AfTransId) {
		return true
	}

	return false
}

// SetAfTransId gets a reference to the given string and assigns it to the AfTransId field.
func (o *EventNotification) SetAfTransId(v string) {
	o.AfTransId = &v
}

// GetDnaiChgType returns the DnaiChgType field value
func (o *EventNotification) GetDnaiChgType() DnaiChangeType {
	if o == nil {
		var ret DnaiChangeType
		return ret
	}

	return o.DnaiChgType
}

// GetDnaiChgTypeOk returns a tuple with the DnaiChgType field value
// and a boolean to check if the value has been set.
func (o *EventNotification) GetDnaiChgTypeOk() (*DnaiChangeType, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DnaiChgType, true
}

// SetDnaiChgType sets field value
func (o *EventNotification) SetDnaiChgType(v DnaiChangeType) {
	o.DnaiChgType = v
}

// GetSourceTrafficRoute returns the SourceTrafficRoute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventNotification) GetSourceTrafficRoute() RouteToLocation {
	if o == nil || isNil(o.SourceTrafficRoute.Get()) {
		var ret RouteToLocation
		return ret
	}
	return *o.SourceTrafficRoute.Get()
}

// GetSourceTrafficRouteOk returns a tuple with the SourceTrafficRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventNotification) GetSourceTrafficRouteOk() (*RouteToLocation, bool) {
	if o == nil {
    return nil, false
	}
	return o.SourceTrafficRoute.Get(), o.SourceTrafficRoute.IsSet()
}

// HasSourceTrafficRoute returns a boolean if a field has been set.
func (o *EventNotification) HasSourceTrafficRoute() bool {
	if o != nil && o.SourceTrafficRoute.IsSet() {
		return true
	}

	return false
}

// SetSourceTrafficRoute gets a reference to the given NullableRouteToLocation and assigns it to the SourceTrafficRoute field.
func (o *EventNotification) SetSourceTrafficRoute(v RouteToLocation) {
	o.SourceTrafficRoute.Set(&v)
}
// SetSourceTrafficRouteNil sets the value for SourceTrafficRoute to be an explicit nil
func (o *EventNotification) SetSourceTrafficRouteNil() {
	o.SourceTrafficRoute.Set(nil)
}

// UnsetSourceTrafficRoute ensures that no value is present for SourceTrafficRoute, not even an explicit nil
func (o *EventNotification) UnsetSourceTrafficRoute() {
	o.SourceTrafficRoute.Unset()
}

// GetSubscribedEvent returns the SubscribedEvent field value
func (o *EventNotification) GetSubscribedEvent() SubscribedEvent {
	if o == nil {
		var ret SubscribedEvent
		return ret
	}

	return o.SubscribedEvent
}

// GetSubscribedEventOk returns a tuple with the SubscribedEvent field value
// and a boolean to check if the value has been set.
func (o *EventNotification) GetSubscribedEventOk() (*SubscribedEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SubscribedEvent, true
}

// SetSubscribedEvent sets field value
func (o *EventNotification) SetSubscribedEvent(v SubscribedEvent) {
	o.SubscribedEvent = v
}

// GetTargetTrafficRoute returns the TargetTrafficRoute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EventNotification) GetTargetTrafficRoute() RouteToLocation {
	if o == nil || isNil(o.TargetTrafficRoute.Get()) {
		var ret RouteToLocation
		return ret
	}
	return *o.TargetTrafficRoute.Get()
}

// GetTargetTrafficRouteOk returns a tuple with the TargetTrafficRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EventNotification) GetTargetTrafficRouteOk() (*RouteToLocation, bool) {
	if o == nil {
    return nil, false
	}
	return o.TargetTrafficRoute.Get(), o.TargetTrafficRoute.IsSet()
}

// HasTargetTrafficRoute returns a boolean if a field has been set.
func (o *EventNotification) HasTargetTrafficRoute() bool {
	if o != nil && o.TargetTrafficRoute.IsSet() {
		return true
	}

	return false
}

// SetTargetTrafficRoute gets a reference to the given NullableRouteToLocation and assigns it to the TargetTrafficRoute field.
func (o *EventNotification) SetTargetTrafficRoute(v RouteToLocation) {
	o.TargetTrafficRoute.Set(&v)
}
// SetTargetTrafficRouteNil sets the value for TargetTrafficRoute to be an explicit nil
func (o *EventNotification) SetTargetTrafficRouteNil() {
	o.TargetTrafficRoute.Set(nil)
}

// UnsetTargetTrafficRoute ensures that no value is present for TargetTrafficRoute, not even an explicit nil
func (o *EventNotification) UnsetTargetTrafficRoute() {
	o.TargetTrafficRoute.Unset()
}

// GetSourceDnai returns the SourceDnai field value if set, zero value otherwise.
func (o *EventNotification) GetSourceDnai() string {
	if o == nil || isNil(o.SourceDnai) {
		var ret string
		return ret
	}
	return *o.SourceDnai
}

// GetSourceDnaiOk returns a tuple with the SourceDnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetSourceDnaiOk() (*string, bool) {
	if o == nil || isNil(o.SourceDnai) {
    return nil, false
	}
	return o.SourceDnai, true
}

// HasSourceDnai returns a boolean if a field has been set.
func (o *EventNotification) HasSourceDnai() bool {
	if o != nil && !isNil(o.SourceDnai) {
		return true
	}

	return false
}

// SetSourceDnai gets a reference to the given string and assigns it to the SourceDnai field.
func (o *EventNotification) SetSourceDnai(v string) {
	o.SourceDnai = &v
}

// GetTargetDnai returns the TargetDnai field value if set, zero value otherwise.
func (o *EventNotification) GetTargetDnai() string {
	if o == nil || isNil(o.TargetDnai) {
		var ret string
		return ret
	}
	return *o.TargetDnai
}

// GetTargetDnaiOk returns a tuple with the TargetDnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetTargetDnaiOk() (*string, bool) {
	if o == nil || isNil(o.TargetDnai) {
    return nil, false
	}
	return o.TargetDnai, true
}

// HasTargetDnai returns a boolean if a field has been set.
func (o *EventNotification) HasTargetDnai() bool {
	if o != nil && !isNil(o.TargetDnai) {
		return true
	}

	return false
}

// SetTargetDnai gets a reference to the given string and assigns it to the TargetDnai field.
func (o *EventNotification) SetTargetDnai(v string) {
	o.TargetDnai = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *EventNotification) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *EventNotification) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *EventNotification) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSrcUeIpv4Addr returns the SrcUeIpv4Addr field value if set, zero value otherwise.
func (o *EventNotification) GetSrcUeIpv4Addr() string {
	if o == nil || isNil(o.SrcUeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.SrcUeIpv4Addr
}

// GetSrcUeIpv4AddrOk returns a tuple with the SrcUeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetSrcUeIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.SrcUeIpv4Addr) {
    return nil, false
	}
	return o.SrcUeIpv4Addr, true
}

// HasSrcUeIpv4Addr returns a boolean if a field has been set.
func (o *EventNotification) HasSrcUeIpv4Addr() bool {
	if o != nil && !isNil(o.SrcUeIpv4Addr) {
		return true
	}

	return false
}

// SetSrcUeIpv4Addr gets a reference to the given string and assigns it to the SrcUeIpv4Addr field.
func (o *EventNotification) SetSrcUeIpv4Addr(v string) {
	o.SrcUeIpv4Addr = &v
}

// GetSrcUeIpv6Prefix returns the SrcUeIpv6Prefix field value if set, zero value otherwise.
func (o *EventNotification) GetSrcUeIpv6Prefix() Ipv6Prefix {
	if o == nil || isNil(o.SrcUeIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.SrcUeIpv6Prefix
}

// GetSrcUeIpv6PrefixOk returns a tuple with the SrcUeIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetSrcUeIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.SrcUeIpv6Prefix) {
    return nil, false
	}
	return o.SrcUeIpv6Prefix, true
}

// HasSrcUeIpv6Prefix returns a boolean if a field has been set.
func (o *EventNotification) HasSrcUeIpv6Prefix() bool {
	if o != nil && !isNil(o.SrcUeIpv6Prefix) {
		return true
	}

	return false
}

// SetSrcUeIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the SrcUeIpv6Prefix field.
func (o *EventNotification) SetSrcUeIpv6Prefix(v Ipv6Prefix) {
	o.SrcUeIpv6Prefix = &v
}

// GetTgtUeIpv4Addr returns the TgtUeIpv4Addr field value if set, zero value otherwise.
func (o *EventNotification) GetTgtUeIpv4Addr() string {
	if o == nil || isNil(o.TgtUeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.TgtUeIpv4Addr
}

// GetTgtUeIpv4AddrOk returns a tuple with the TgtUeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetTgtUeIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.TgtUeIpv4Addr) {
    return nil, false
	}
	return o.TgtUeIpv4Addr, true
}

// HasTgtUeIpv4Addr returns a boolean if a field has been set.
func (o *EventNotification) HasTgtUeIpv4Addr() bool {
	if o != nil && !isNil(o.TgtUeIpv4Addr) {
		return true
	}

	return false
}

// SetTgtUeIpv4Addr gets a reference to the given string and assigns it to the TgtUeIpv4Addr field.
func (o *EventNotification) SetTgtUeIpv4Addr(v string) {
	o.TgtUeIpv4Addr = &v
}

// GetTgtUeIpv6Prefix returns the TgtUeIpv6Prefix field value if set, zero value otherwise.
func (o *EventNotification) GetTgtUeIpv6Prefix() Ipv6Prefix {
	if o == nil || isNil(o.TgtUeIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.TgtUeIpv6Prefix
}

// GetTgtUeIpv6PrefixOk returns a tuple with the TgtUeIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetTgtUeIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.TgtUeIpv6Prefix) {
    return nil, false
	}
	return o.TgtUeIpv6Prefix, true
}

// HasTgtUeIpv6Prefix returns a boolean if a field has been set.
func (o *EventNotification) HasTgtUeIpv6Prefix() bool {
	if o != nil && !isNil(o.TgtUeIpv6Prefix) {
		return true
	}

	return false
}

// SetTgtUeIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the TgtUeIpv6Prefix field.
func (o *EventNotification) SetTgtUeIpv6Prefix(v Ipv6Prefix) {
	o.TgtUeIpv6Prefix = &v
}

// GetUeMac returns the UeMac field value if set, zero value otherwise.
func (o *EventNotification) GetUeMac() string {
	if o == nil || isNil(o.UeMac) {
		var ret string
		return ret
	}
	return *o.UeMac
}

// GetUeMacOk returns a tuple with the UeMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetUeMacOk() (*string, bool) {
	if o == nil || isNil(o.UeMac) {
    return nil, false
	}
	return o.UeMac, true
}

// HasUeMac returns a boolean if a field has been set.
func (o *EventNotification) HasUeMac() bool {
	if o != nil && !isNil(o.UeMac) {
		return true
	}

	return false
}

// SetUeMac gets a reference to the given string and assigns it to the UeMac field.
func (o *EventNotification) SetUeMac(v string) {
	o.UeMac = &v
}

// GetAfAckUri returns the AfAckUri field value if set, zero value otherwise.
func (o *EventNotification) GetAfAckUri() string {
	if o == nil || isNil(o.AfAckUri) {
		var ret string
		return ret
	}
	return *o.AfAckUri
}

// GetAfAckUriOk returns a tuple with the AfAckUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventNotification) GetAfAckUriOk() (*string, bool) {
	if o == nil || isNil(o.AfAckUri) {
    return nil, false
	}
	return o.AfAckUri, true
}

// HasAfAckUri returns a boolean if a field has been set.
func (o *EventNotification) HasAfAckUri() bool {
	if o != nil && !isNil(o.AfAckUri) {
		return true
	}

	return false
}

// SetAfAckUri gets a reference to the given string and assigns it to the AfAckUri field.
func (o *EventNotification) SetAfAckUri(v string) {
	o.AfAckUri = &v
}

func (o EventNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AfTransId) {
		toSerialize["afTransId"] = o.AfTransId
	}
	if true {
		toSerialize["dnaiChgType"] = o.DnaiChgType
	}
	if o.SourceTrafficRoute.IsSet() {
		toSerialize["sourceTrafficRoute"] = o.SourceTrafficRoute.Get()
	}
	if true {
		toSerialize["subscribedEvent"] = o.SubscribedEvent
	}
	if o.TargetTrafficRoute.IsSet() {
		toSerialize["targetTrafficRoute"] = o.TargetTrafficRoute.Get()
	}
	if !isNil(o.SourceDnai) {
		toSerialize["sourceDnai"] = o.SourceDnai
	}
	if !isNil(o.TargetDnai) {
		toSerialize["targetDnai"] = o.TargetDnai
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.SrcUeIpv4Addr) {
		toSerialize["srcUeIpv4Addr"] = o.SrcUeIpv4Addr
	}
	if !isNil(o.SrcUeIpv6Prefix) {
		toSerialize["srcUeIpv6Prefix"] = o.SrcUeIpv6Prefix
	}
	if !isNil(o.TgtUeIpv4Addr) {
		toSerialize["tgtUeIpv4Addr"] = o.TgtUeIpv4Addr
	}
	if !isNil(o.TgtUeIpv6Prefix) {
		toSerialize["tgtUeIpv6Prefix"] = o.TgtUeIpv6Prefix
	}
	if !isNil(o.UeMac) {
		toSerialize["ueMac"] = o.UeMac
	}
	if !isNil(o.AfAckUri) {
		toSerialize["afAckUri"] = o.AfAckUri
	}
	return json.Marshal(toSerialize)
}

type NullableEventNotification struct {
	value *EventNotification
	isSet bool
}

func (v NullableEventNotification) Get() *EventNotification {
	return v.value
}

func (v *NullableEventNotification) Set(val *EventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotification(val *EventNotification) *NullableEventNotification {
	return &NullableEventNotification{value: val, isSet: true}
}

func (v NullableEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


