/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_EventExposure

import (
	"encoding/json"
)

// EventSubscription Represents a subscription to a single event.
type EventSubscription struct {
	Event SmfEvent `json:"event"`
	DnaiChgType *DnaiChangeType `json:"dnaiChgType,omitempty"`
	DddTraDescriptors []DddTrafficDescriptor `json:"dddTraDescriptors,omitempty"`
	DddStati []DlDataDeliveryStatus `json:"dddStati,omitempty"`
	AppIds []string `json:"appIds,omitempty"`
	TargetPeriod *TimeWindow `json:"targetPeriod,omitempty"`
	// Indicates the subscription for UE transaction dispersion collectionon, if it is included and set to \"true\". Default value is \"false\". 
	TransacDispInd *bool `json:"transacDispInd,omitempty"`
	// Indicates Session Management Transaction metrics.
	TransacMetrics []TransactionMetric `json:"transacMetrics,omitempty"`
	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty"`
}

// NewEventSubscription instantiates a new EventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSubscription(event SmfEvent) *EventSubscription {
	this := EventSubscription{}
	this.Event = event
	return &this
}

// NewEventSubscriptionWithDefaults instantiates a new EventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSubscriptionWithDefaults() *EventSubscription {
	this := EventSubscription{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventSubscription) GetEvent() SmfEvent {
	if o == nil {
		var ret SmfEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetEventOk() (*SmfEvent, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventSubscription) SetEvent(v SmfEvent) {
	o.Event = v
}

// GetDnaiChgType returns the DnaiChgType field value if set, zero value otherwise.
func (o *EventSubscription) GetDnaiChgType() DnaiChangeType {
	if o == nil || isNil(o.DnaiChgType) {
		var ret DnaiChangeType
		return ret
	}
	return *o.DnaiChgType
}

// GetDnaiChgTypeOk returns a tuple with the DnaiChgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetDnaiChgTypeOk() (*DnaiChangeType, bool) {
	if o == nil || isNil(o.DnaiChgType) {
    return nil, false
	}
	return o.DnaiChgType, true
}

// HasDnaiChgType returns a boolean if a field has been set.
func (o *EventSubscription) HasDnaiChgType() bool {
	if o != nil && !isNil(o.DnaiChgType) {
		return true
	}

	return false
}

// SetDnaiChgType gets a reference to the given DnaiChangeType and assigns it to the DnaiChgType field.
func (o *EventSubscription) SetDnaiChgType(v DnaiChangeType) {
	o.DnaiChgType = &v
}

// GetDddTraDescriptors returns the DddTraDescriptors field value if set, zero value otherwise.
func (o *EventSubscription) GetDddTraDescriptors() []DddTrafficDescriptor {
	if o == nil || isNil(o.DddTraDescriptors) {
		var ret []DddTrafficDescriptor
		return ret
	}
	return o.DddTraDescriptors
}

// GetDddTraDescriptorsOk returns a tuple with the DddTraDescriptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetDddTraDescriptorsOk() ([]DddTrafficDescriptor, bool) {
	if o == nil || isNil(o.DddTraDescriptors) {
    return nil, false
	}
	return o.DddTraDescriptors, true
}

// HasDddTraDescriptors returns a boolean if a field has been set.
func (o *EventSubscription) HasDddTraDescriptors() bool {
	if o != nil && !isNil(o.DddTraDescriptors) {
		return true
	}

	return false
}

// SetDddTraDescriptors gets a reference to the given []DddTrafficDescriptor and assigns it to the DddTraDescriptors field.
func (o *EventSubscription) SetDddTraDescriptors(v []DddTrafficDescriptor) {
	o.DddTraDescriptors = v
}

// GetDddStati returns the DddStati field value if set, zero value otherwise.
func (o *EventSubscription) GetDddStati() []DlDataDeliveryStatus {
	if o == nil || isNil(o.DddStati) {
		var ret []DlDataDeliveryStatus
		return ret
	}
	return o.DddStati
}

// GetDddStatiOk returns a tuple with the DddStati field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetDddStatiOk() ([]DlDataDeliveryStatus, bool) {
	if o == nil || isNil(o.DddStati) {
    return nil, false
	}
	return o.DddStati, true
}

// HasDddStati returns a boolean if a field has been set.
func (o *EventSubscription) HasDddStati() bool {
	if o != nil && !isNil(o.DddStati) {
		return true
	}

	return false
}

// SetDddStati gets a reference to the given []DlDataDeliveryStatus and assigns it to the DddStati field.
func (o *EventSubscription) SetDddStati(v []DlDataDeliveryStatus) {
	o.DddStati = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *EventSubscription) GetAppIds() []string {
	if o == nil || isNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetAppIdsOk() ([]string, bool) {
	if o == nil || isNil(o.AppIds) {
    return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *EventSubscription) HasAppIds() bool {
	if o != nil && !isNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *EventSubscription) SetAppIds(v []string) {
	o.AppIds = v
}

// GetTargetPeriod returns the TargetPeriod field value if set, zero value otherwise.
func (o *EventSubscription) GetTargetPeriod() TimeWindow {
	if o == nil || isNil(o.TargetPeriod) {
		var ret TimeWindow
		return ret
	}
	return *o.TargetPeriod
}

// GetTargetPeriodOk returns a tuple with the TargetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetTargetPeriodOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.TargetPeriod) {
    return nil, false
	}
	return o.TargetPeriod, true
}

// HasTargetPeriod returns a boolean if a field has been set.
func (o *EventSubscription) HasTargetPeriod() bool {
	if o != nil && !isNil(o.TargetPeriod) {
		return true
	}

	return false
}

// SetTargetPeriod gets a reference to the given TimeWindow and assigns it to the TargetPeriod field.
func (o *EventSubscription) SetTargetPeriod(v TimeWindow) {
	o.TargetPeriod = &v
}

// GetTransacDispInd returns the TransacDispInd field value if set, zero value otherwise.
func (o *EventSubscription) GetTransacDispInd() bool {
	if o == nil || isNil(o.TransacDispInd) {
		var ret bool
		return ret
	}
	return *o.TransacDispInd
}

// GetTransacDispIndOk returns a tuple with the TransacDispInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetTransacDispIndOk() (*bool, bool) {
	if o == nil || isNil(o.TransacDispInd) {
    return nil, false
	}
	return o.TransacDispInd, true
}

// HasTransacDispInd returns a boolean if a field has been set.
func (o *EventSubscription) HasTransacDispInd() bool {
	if o != nil && !isNil(o.TransacDispInd) {
		return true
	}

	return false
}

// SetTransacDispInd gets a reference to the given bool and assigns it to the TransacDispInd field.
func (o *EventSubscription) SetTransacDispInd(v bool) {
	o.TransacDispInd = &v
}

// GetTransacMetrics returns the TransacMetrics field value if set, zero value otherwise.
func (o *EventSubscription) GetTransacMetrics() []TransactionMetric {
	if o == nil || isNil(o.TransacMetrics) {
		var ret []TransactionMetric
		return ret
	}
	return o.TransacMetrics
}

// GetTransacMetricsOk returns a tuple with the TransacMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetTransacMetricsOk() ([]TransactionMetric, bool) {
	if o == nil || isNil(o.TransacMetrics) {
    return nil, false
	}
	return o.TransacMetrics, true
}

// HasTransacMetrics returns a boolean if a field has been set.
func (o *EventSubscription) HasTransacMetrics() bool {
	if o != nil && !isNil(o.TransacMetrics) {
		return true
	}

	return false
}

// SetTransacMetrics gets a reference to the given []TransactionMetric and assigns it to the TransacMetrics field.
func (o *EventSubscription) SetTransacMetrics(v []TransactionMetric) {
	o.TransacMetrics = v
}

// GetUeIpAddr returns the UeIpAddr field value if set, zero value otherwise.
func (o *EventSubscription) GetUeIpAddr() IpAddr {
	if o == nil || isNil(o.UeIpAddr) {
		var ret IpAddr
		return ret
	}
	return *o.UeIpAddr
}

// GetUeIpAddrOk returns a tuple with the UeIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription) GetUeIpAddrOk() (*IpAddr, bool) {
	if o == nil || isNil(o.UeIpAddr) {
    return nil, false
	}
	return o.UeIpAddr, true
}

// HasUeIpAddr returns a boolean if a field has been set.
func (o *EventSubscription) HasUeIpAddr() bool {
	if o != nil && !isNil(o.UeIpAddr) {
		return true
	}

	return false
}

// SetUeIpAddr gets a reference to the given IpAddr and assigns it to the UeIpAddr field.
func (o *EventSubscription) SetUeIpAddr(v IpAddr) {
	o.UeIpAddr = &v
}

func (o EventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event"] = o.Event
	}
	if !isNil(o.DnaiChgType) {
		toSerialize["dnaiChgType"] = o.DnaiChgType
	}
	if !isNil(o.DddTraDescriptors) {
		toSerialize["dddTraDescriptors"] = o.DddTraDescriptors
	}
	if !isNil(o.DddStati) {
		toSerialize["dddStati"] = o.DddStati
	}
	if !isNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !isNil(o.TargetPeriod) {
		toSerialize["targetPeriod"] = o.TargetPeriod
	}
	if !isNil(o.TransacDispInd) {
		toSerialize["transacDispInd"] = o.TransacDispInd
	}
	if !isNil(o.TransacMetrics) {
		toSerialize["transacMetrics"] = o.TransacMetrics
	}
	if !isNil(o.UeIpAddr) {
		toSerialize["ueIpAddr"] = o.UeIpAddr
	}
	return json.Marshal(toSerialize)
}

type NullableEventSubscription struct {
	value *EventSubscription
	isSet bool
}

func (v NullableEventSubscription) Get() *EventSubscription {
	return v.value
}

func (v *NullableEventSubscription) Set(val *EventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSubscription(val *EventSubscription) *NullableEventSubscription {
	return &NullableEventSubscription{value: val, isSet: true}
}

func (v NullableEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


