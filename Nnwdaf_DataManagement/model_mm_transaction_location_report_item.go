/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// MmTransactionLocationReportItem UE MM Transaction Report Item per Location
type MmTransactionLocationReportItem struct {
	Tai *Tai `json:"tai,omitempty"`
	Ncgi *Ncgi `json:"ncgi,omitempty"`
	Ecgi *Ecgi `json:"ecgi,omitempty"`
	N3gaLocation *N3gaLocation `json:"n3gaLocation,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
	Transactions int32 `json:"transactions"`
}

// NewMmTransactionLocationReportItem instantiates a new MmTransactionLocationReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMmTransactionLocationReportItem(timestamp time.Time, transactions int32) *MmTransactionLocationReportItem {
	this := MmTransactionLocationReportItem{}
	this.Timestamp = timestamp
	this.Transactions = transactions
	return &this
}

// NewMmTransactionLocationReportItemWithDefaults instantiates a new MmTransactionLocationReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMmTransactionLocationReportItemWithDefaults() *MmTransactionLocationReportItem {
	this := MmTransactionLocationReportItem{}
	return &this
}

// GetTai returns the Tai field value if set, zero value otherwise.
func (o *MmTransactionLocationReportItem) GetTai() Tai {
	if o == nil || isNil(o.Tai) {
		var ret Tai
		return ret
	}
	return *o.Tai
}

// GetTaiOk returns a tuple with the Tai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmTransactionLocationReportItem) GetTaiOk() (*Tai, bool) {
	if o == nil || isNil(o.Tai) {
    return nil, false
	}
	return o.Tai, true
}

// HasTai returns a boolean if a field has been set.
func (o *MmTransactionLocationReportItem) HasTai() bool {
	if o != nil && !isNil(o.Tai) {
		return true
	}

	return false
}

// SetTai gets a reference to the given Tai and assigns it to the Tai field.
func (o *MmTransactionLocationReportItem) SetTai(v Tai) {
	o.Tai = &v
}

// GetNcgi returns the Ncgi field value if set, zero value otherwise.
func (o *MmTransactionLocationReportItem) GetNcgi() Ncgi {
	if o == nil || isNil(o.Ncgi) {
		var ret Ncgi
		return ret
	}
	return *o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmTransactionLocationReportItem) GetNcgiOk() (*Ncgi, bool) {
	if o == nil || isNil(o.Ncgi) {
    return nil, false
	}
	return o.Ncgi, true
}

// HasNcgi returns a boolean if a field has been set.
func (o *MmTransactionLocationReportItem) HasNcgi() bool {
	if o != nil && !isNil(o.Ncgi) {
		return true
	}

	return false
}

// SetNcgi gets a reference to the given Ncgi and assigns it to the Ncgi field.
func (o *MmTransactionLocationReportItem) SetNcgi(v Ncgi) {
	o.Ncgi = &v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *MmTransactionLocationReportItem) GetEcgi() Ecgi {
	if o == nil || isNil(o.Ecgi) {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmTransactionLocationReportItem) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || isNil(o.Ecgi) {
    return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *MmTransactionLocationReportItem) HasEcgi() bool {
	if o != nil && !isNil(o.Ecgi) {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *MmTransactionLocationReportItem) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

// GetN3gaLocation returns the N3gaLocation field value if set, zero value otherwise.
func (o *MmTransactionLocationReportItem) GetN3gaLocation() N3gaLocation {
	if o == nil || isNil(o.N3gaLocation) {
		var ret N3gaLocation
		return ret
	}
	return *o.N3gaLocation
}

// GetN3gaLocationOk returns a tuple with the N3gaLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MmTransactionLocationReportItem) GetN3gaLocationOk() (*N3gaLocation, bool) {
	if o == nil || isNil(o.N3gaLocation) {
    return nil, false
	}
	return o.N3gaLocation, true
}

// HasN3gaLocation returns a boolean if a field has been set.
func (o *MmTransactionLocationReportItem) HasN3gaLocation() bool {
	if o != nil && !isNil(o.N3gaLocation) {
		return true
	}

	return false
}

// SetN3gaLocation gets a reference to the given N3gaLocation and assigns it to the N3gaLocation field.
func (o *MmTransactionLocationReportItem) SetN3gaLocation(v N3gaLocation) {
	o.N3gaLocation = &v
}

// GetTimestamp returns the Timestamp field value
func (o *MmTransactionLocationReportItem) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MmTransactionLocationReportItem) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MmTransactionLocationReportItem) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTransactions returns the Transactions field value
func (o *MmTransactionLocationReportItem) GetTransactions() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *MmTransactionLocationReportItem) GetTransactionsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *MmTransactionLocationReportItem) SetTransactions(v int32) {
	o.Transactions = v
}

func (o MmTransactionLocationReportItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tai) {
		toSerialize["tai"] = o.Tai
	}
	if !isNil(o.Ncgi) {
		toSerialize["ncgi"] = o.Ncgi
	}
	if !isNil(o.Ecgi) {
		toSerialize["ecgi"] = o.Ecgi
	}
	if !isNil(o.N3gaLocation) {
		toSerialize["n3gaLocation"] = o.N3gaLocation
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["transactions"] = o.Transactions
	}
	return json.Marshal(toSerialize)
}

type NullableMmTransactionLocationReportItem struct {
	value *MmTransactionLocationReportItem
	isSet bool
}

func (v NullableMmTransactionLocationReportItem) Get() *MmTransactionLocationReportItem {
	return v.value
}

func (v *NullableMmTransactionLocationReportItem) Set(val *MmTransactionLocationReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMmTransactionLocationReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMmTransactionLocationReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMmTransactionLocationReportItem(val *MmTransactionLocationReportItem) *NullableMmTransactionLocationReportItem {
	return &NullableMmTransactionLocationReportItem{value: val, isSet: true}
}

func (v NullableMmTransactionLocationReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMmTransactionLocationReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


