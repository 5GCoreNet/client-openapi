/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nbsf_Management

import (
	"encoding/json"
)

// BsfSubscription Contains the event subscription data.
type BsfSubscription struct {
	// Contain te subscribed events.
	Events []BsfEvent `json:"events"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`
	// Notification Correlation ID assigned by the NF service consumer.
	NotifCorreId string `json:"notifCorreId"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	SnssaiDnnPairs *SnssaiDnnPair `json:"snssaiDnnPairs,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewBsfSubscription instantiates a new BsfSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBsfSubscription(events []BsfEvent, notifUri string, notifCorreId string, supi string) *BsfSubscription {
	this := BsfSubscription{}
	this.Events = events
	this.NotifUri = notifUri
	this.NotifCorreId = notifCorreId
	this.Supi = supi
	return &this
}

// NewBsfSubscriptionWithDefaults instantiates a new BsfSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBsfSubscriptionWithDefaults() *BsfSubscription {
	this := BsfSubscription{}
	return &this
}

// GetEvents returns the Events field value
func (o *BsfSubscription) GetEvents() []BsfEvent {
	if o == nil {
		var ret []BsfEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *BsfSubscription) GetEventsOk() ([]BsfEvent, bool) {
	if o == nil {
    return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *BsfSubscription) SetEvents(v []BsfEvent) {
	o.Events = v
}

// GetNotifUri returns the NotifUri field value
func (o *BsfSubscription) GetNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value
// and a boolean to check if the value has been set.
func (o *BsfSubscription) GetNotifUriOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifUri, true
}

// SetNotifUri sets field value
func (o *BsfSubscription) SetNotifUri(v string) {
	o.NotifUri = v
}

// GetNotifCorreId returns the NotifCorreId field value
func (o *BsfSubscription) GetNotifCorreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifCorreId
}

// GetNotifCorreIdOk returns a tuple with the NotifCorreId field value
// and a boolean to check if the value has been set.
func (o *BsfSubscription) GetNotifCorreIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.NotifCorreId, true
}

// SetNotifCorreId sets field value
func (o *BsfSubscription) SetNotifCorreId(v string) {
	o.NotifCorreId = v
}

// GetSupi returns the Supi field value
func (o *BsfSubscription) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *BsfSubscription) GetSupiOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *BsfSubscription) SetSupi(v string) {
	o.Supi = v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *BsfSubscription) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfSubscription) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
    return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *BsfSubscription) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *BsfSubscription) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetSnssaiDnnPairs returns the SnssaiDnnPairs field value if set, zero value otherwise.
func (o *BsfSubscription) GetSnssaiDnnPairs() SnssaiDnnPair {
	if o == nil || isNil(o.SnssaiDnnPairs) {
		var ret SnssaiDnnPair
		return ret
	}
	return *o.SnssaiDnnPairs
}

// GetSnssaiDnnPairsOk returns a tuple with the SnssaiDnnPairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfSubscription) GetSnssaiDnnPairsOk() (*SnssaiDnnPair, bool) {
	if o == nil || isNil(o.SnssaiDnnPairs) {
    return nil, false
	}
	return o.SnssaiDnnPairs, true
}

// HasSnssaiDnnPairs returns a boolean if a field has been set.
func (o *BsfSubscription) HasSnssaiDnnPairs() bool {
	if o != nil && !isNil(o.SnssaiDnnPairs) {
		return true
	}

	return false
}

// SetSnssaiDnnPairs gets a reference to the given SnssaiDnnPair and assigns it to the SnssaiDnnPairs field.
func (o *BsfSubscription) SetSnssaiDnnPairs(v SnssaiDnnPair) {
	o.SnssaiDnnPairs = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *BsfSubscription) GetSuppFeat() string {
	if o == nil || isNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BsfSubscription) GetSuppFeatOk() (*string, bool) {
	if o == nil || isNil(o.SuppFeat) {
    return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *BsfSubscription) HasSuppFeat() bool {
	if o != nil && !isNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *BsfSubscription) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o BsfSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["events"] = o.Events
	}
	if true {
		toSerialize["notifUri"] = o.NotifUri
	}
	if true {
		toSerialize["notifCorreId"] = o.NotifCorreId
	}
	if true {
		toSerialize["supi"] = o.Supi
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !isNil(o.SnssaiDnnPairs) {
		toSerialize["snssaiDnnPairs"] = o.SnssaiDnnPairs
	}
	if !isNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return json.Marshal(toSerialize)
}

type NullableBsfSubscription struct {
	value *BsfSubscription
	isSet bool
}

func (v NullableBsfSubscription) Get() *BsfSubscription {
	return v.value
}

func (v *NullableBsfSubscription) Set(val *BsfSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableBsfSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableBsfSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBsfSubscription(val *BsfSubscription) *NullableBsfSubscription {
	return &NullableBsfSubscription{value: val, isSet: true}
}

func (v NullableBsfSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBsfSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


