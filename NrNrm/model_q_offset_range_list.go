/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.7.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package NrNrm

import (
	"encoding/json"
)

// QOffsetRangeList struct for QOffsetRangeList
type QOffsetRangeList struct {
	RsrpOffsetSSB *QOffsetRange `json:"rsrpOffsetSSB,omitempty"`
	RsrqOffsetSSB *QOffsetRange `json:"rsrqOffsetSSB,omitempty"`
	SinrOffsetSSB *QOffsetRange `json:"sinrOffsetSSB,omitempty"`
	RsrpOffsetCSIRS *QOffsetRange `json:"rsrpOffsetCSI-RS,omitempty"`
	RsrqOffsetCSIRS *QOffsetRange `json:"rsrqOffsetCSI-RS,omitempty"`
	SinrOffsetCSIRS *QOffsetRange `json:"sinrOffsetCSI-RS,omitempty"`
}

// NewQOffsetRangeList instantiates a new QOffsetRangeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQOffsetRangeList() *QOffsetRangeList {
	this := QOffsetRangeList{}
	return &this
}

// NewQOffsetRangeListWithDefaults instantiates a new QOffsetRangeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQOffsetRangeListWithDefaults() *QOffsetRangeList {
	this := QOffsetRangeList{}
	return &this
}

// GetRsrpOffsetSSB returns the RsrpOffsetSSB field value if set, zero value otherwise.
func (o *QOffsetRangeList) GetRsrpOffsetSSB() QOffsetRange {
	if o == nil || isNil(o.RsrpOffsetSSB) {
		var ret QOffsetRange
		return ret
	}
	return *o.RsrpOffsetSSB
}

// GetRsrpOffsetSSBOk returns a tuple with the RsrpOffsetSSB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QOffsetRangeList) GetRsrpOffsetSSBOk() (*QOffsetRange, bool) {
	if o == nil || isNil(o.RsrpOffsetSSB) {
    return nil, false
	}
	return o.RsrpOffsetSSB, true
}

// HasRsrpOffsetSSB returns a boolean if a field has been set.
func (o *QOffsetRangeList) HasRsrpOffsetSSB() bool {
	if o != nil && !isNil(o.RsrpOffsetSSB) {
		return true
	}

	return false
}

// SetRsrpOffsetSSB gets a reference to the given QOffsetRange and assigns it to the RsrpOffsetSSB field.
func (o *QOffsetRangeList) SetRsrpOffsetSSB(v QOffsetRange) {
	o.RsrpOffsetSSB = &v
}

// GetRsrqOffsetSSB returns the RsrqOffsetSSB field value if set, zero value otherwise.
func (o *QOffsetRangeList) GetRsrqOffsetSSB() QOffsetRange {
	if o == nil || isNil(o.RsrqOffsetSSB) {
		var ret QOffsetRange
		return ret
	}
	return *o.RsrqOffsetSSB
}

// GetRsrqOffsetSSBOk returns a tuple with the RsrqOffsetSSB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QOffsetRangeList) GetRsrqOffsetSSBOk() (*QOffsetRange, bool) {
	if o == nil || isNil(o.RsrqOffsetSSB) {
    return nil, false
	}
	return o.RsrqOffsetSSB, true
}

// HasRsrqOffsetSSB returns a boolean if a field has been set.
func (o *QOffsetRangeList) HasRsrqOffsetSSB() bool {
	if o != nil && !isNil(o.RsrqOffsetSSB) {
		return true
	}

	return false
}

// SetRsrqOffsetSSB gets a reference to the given QOffsetRange and assigns it to the RsrqOffsetSSB field.
func (o *QOffsetRangeList) SetRsrqOffsetSSB(v QOffsetRange) {
	o.RsrqOffsetSSB = &v
}

// GetSinrOffsetSSB returns the SinrOffsetSSB field value if set, zero value otherwise.
func (o *QOffsetRangeList) GetSinrOffsetSSB() QOffsetRange {
	if o == nil || isNil(o.SinrOffsetSSB) {
		var ret QOffsetRange
		return ret
	}
	return *o.SinrOffsetSSB
}

// GetSinrOffsetSSBOk returns a tuple with the SinrOffsetSSB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QOffsetRangeList) GetSinrOffsetSSBOk() (*QOffsetRange, bool) {
	if o == nil || isNil(o.SinrOffsetSSB) {
    return nil, false
	}
	return o.SinrOffsetSSB, true
}

// HasSinrOffsetSSB returns a boolean if a field has been set.
func (o *QOffsetRangeList) HasSinrOffsetSSB() bool {
	if o != nil && !isNil(o.SinrOffsetSSB) {
		return true
	}

	return false
}

// SetSinrOffsetSSB gets a reference to the given QOffsetRange and assigns it to the SinrOffsetSSB field.
func (o *QOffsetRangeList) SetSinrOffsetSSB(v QOffsetRange) {
	o.SinrOffsetSSB = &v
}

// GetRsrpOffsetCSIRS returns the RsrpOffsetCSIRS field value if set, zero value otherwise.
func (o *QOffsetRangeList) GetRsrpOffsetCSIRS() QOffsetRange {
	if o == nil || isNil(o.RsrpOffsetCSIRS) {
		var ret QOffsetRange
		return ret
	}
	return *o.RsrpOffsetCSIRS
}

// GetRsrpOffsetCSIRSOk returns a tuple with the RsrpOffsetCSIRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QOffsetRangeList) GetRsrpOffsetCSIRSOk() (*QOffsetRange, bool) {
	if o == nil || isNil(o.RsrpOffsetCSIRS) {
    return nil, false
	}
	return o.RsrpOffsetCSIRS, true
}

// HasRsrpOffsetCSIRS returns a boolean if a field has been set.
func (o *QOffsetRangeList) HasRsrpOffsetCSIRS() bool {
	if o != nil && !isNil(o.RsrpOffsetCSIRS) {
		return true
	}

	return false
}

// SetRsrpOffsetCSIRS gets a reference to the given QOffsetRange and assigns it to the RsrpOffsetCSIRS field.
func (o *QOffsetRangeList) SetRsrpOffsetCSIRS(v QOffsetRange) {
	o.RsrpOffsetCSIRS = &v
}

// GetRsrqOffsetCSIRS returns the RsrqOffsetCSIRS field value if set, zero value otherwise.
func (o *QOffsetRangeList) GetRsrqOffsetCSIRS() QOffsetRange {
	if o == nil || isNil(o.RsrqOffsetCSIRS) {
		var ret QOffsetRange
		return ret
	}
	return *o.RsrqOffsetCSIRS
}

// GetRsrqOffsetCSIRSOk returns a tuple with the RsrqOffsetCSIRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QOffsetRangeList) GetRsrqOffsetCSIRSOk() (*QOffsetRange, bool) {
	if o == nil || isNil(o.RsrqOffsetCSIRS) {
    return nil, false
	}
	return o.RsrqOffsetCSIRS, true
}

// HasRsrqOffsetCSIRS returns a boolean if a field has been set.
func (o *QOffsetRangeList) HasRsrqOffsetCSIRS() bool {
	if o != nil && !isNil(o.RsrqOffsetCSIRS) {
		return true
	}

	return false
}

// SetRsrqOffsetCSIRS gets a reference to the given QOffsetRange and assigns it to the RsrqOffsetCSIRS field.
func (o *QOffsetRangeList) SetRsrqOffsetCSIRS(v QOffsetRange) {
	o.RsrqOffsetCSIRS = &v
}

// GetSinrOffsetCSIRS returns the SinrOffsetCSIRS field value if set, zero value otherwise.
func (o *QOffsetRangeList) GetSinrOffsetCSIRS() QOffsetRange {
	if o == nil || isNil(o.SinrOffsetCSIRS) {
		var ret QOffsetRange
		return ret
	}
	return *o.SinrOffsetCSIRS
}

// GetSinrOffsetCSIRSOk returns a tuple with the SinrOffsetCSIRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QOffsetRangeList) GetSinrOffsetCSIRSOk() (*QOffsetRange, bool) {
	if o == nil || isNil(o.SinrOffsetCSIRS) {
    return nil, false
	}
	return o.SinrOffsetCSIRS, true
}

// HasSinrOffsetCSIRS returns a boolean if a field has been set.
func (o *QOffsetRangeList) HasSinrOffsetCSIRS() bool {
	if o != nil && !isNil(o.SinrOffsetCSIRS) {
		return true
	}

	return false
}

// SetSinrOffsetCSIRS gets a reference to the given QOffsetRange and assigns it to the SinrOffsetCSIRS field.
func (o *QOffsetRangeList) SetSinrOffsetCSIRS(v QOffsetRange) {
	o.SinrOffsetCSIRS = &v
}

func (o QOffsetRangeList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RsrpOffsetSSB) {
		toSerialize["rsrpOffsetSSB"] = o.RsrpOffsetSSB
	}
	if !isNil(o.RsrqOffsetSSB) {
		toSerialize["rsrqOffsetSSB"] = o.RsrqOffsetSSB
	}
	if !isNil(o.SinrOffsetSSB) {
		toSerialize["sinrOffsetSSB"] = o.SinrOffsetSSB
	}
	if !isNil(o.RsrpOffsetCSIRS) {
		toSerialize["rsrpOffsetCSI-RS"] = o.RsrpOffsetCSIRS
	}
	if !isNil(o.RsrqOffsetCSIRS) {
		toSerialize["rsrqOffsetCSI-RS"] = o.RsrqOffsetCSIRS
	}
	if !isNil(o.SinrOffsetCSIRS) {
		toSerialize["sinrOffsetCSI-RS"] = o.SinrOffsetCSIRS
	}
	return json.Marshal(toSerialize)
}

type NullableQOffsetRangeList struct {
	value *QOffsetRangeList
	isSet bool
}

func (v NullableQOffsetRangeList) Get() *QOffsetRangeList {
	return v.value
}

func (v *NullableQOffsetRangeList) Set(val *QOffsetRangeList) {
	v.value = val
	v.isSet = true
}

func (v NullableQOffsetRangeList) IsSet() bool {
	return v.isSet
}

func (v *NullableQOffsetRangeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQOffsetRangeList(val *QOffsetRangeList) *NullableQOffsetRangeList {
	return &NullableQOffsetRangeList{value: val, isSet: true}
}

func (v NullableQOffsetRangeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQOffsetRangeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


