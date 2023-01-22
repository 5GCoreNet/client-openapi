/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package StreamingDataMnS

import (
	"encoding/json"
)

// ProducerIdType DN of the streaming data reporting MnS producer.
type ProducerIdType struct {
}

// NewProducerIdType instantiates a new ProducerIdType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProducerIdType() *ProducerIdType {
	this := ProducerIdType{}
	return &this
}

// NewProducerIdTypeWithDefaults instantiates a new ProducerIdType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProducerIdTypeWithDefaults() *ProducerIdType {
	this := ProducerIdType{}
	return &this
}

func (o ProducerIdType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	return json.Marshal(toSerialize)
}

type NullableProducerIdType struct {
	value *ProducerIdType
	isSet bool
}

func (v NullableProducerIdType) Get() *ProducerIdType {
	return v.value
}

func (v *NullableProducerIdType) Set(val *ProducerIdType) {
	v.value = val
	v.isSet = true
}

func (v NullableProducerIdType) IsSet() bool {
	return v.isSet
}

func (v *NullableProducerIdType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProducerIdType(val *ProducerIdType) *NullableProducerIdType {
	return &NullableProducerIdType{value: val, isSet: true}
}

func (v NullableProducerIdType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProducerIdType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

