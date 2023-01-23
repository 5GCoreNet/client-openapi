/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
)

// MLTrainingRequestSingleAllOf struct for MLTrainingRequestSingleAllOf
type MLTrainingRequestSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewMLTrainingRequestSingleAllOf instantiates a new MLTrainingRequestSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLTrainingRequestSingleAllOf() *MLTrainingRequestSingleAllOf {
	this := MLTrainingRequestSingleAllOf{}
	return &this
}

// NewMLTrainingRequestSingleAllOfWithDefaults instantiates a new MLTrainingRequestSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLTrainingRequestSingleAllOfWithDefaults() *MLTrainingRequestSingleAllOf {
	this := MLTrainingRequestSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MLTrainingRequestSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingRequestSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
    return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MLTrainingRequestSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *MLTrainingRequestSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o MLTrainingRequestSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableMLTrainingRequestSingleAllOf struct {
	value *MLTrainingRequestSingleAllOf
	isSet bool
}

func (v NullableMLTrainingRequestSingleAllOf) Get() *MLTrainingRequestSingleAllOf {
	return v.value
}

func (v *NullableMLTrainingRequestSingleAllOf) Set(val *MLTrainingRequestSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMLTrainingRequestSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMLTrainingRequestSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLTrainingRequestSingleAllOf(val *MLTrainingRequestSingleAllOf) *NullableMLTrainingRequestSingleAllOf {
	return &NullableMLTrainingRequestSingleAllOf{value: val, isSet: true}
}

func (v NullableMLTrainingRequestSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLTrainingRequestSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


