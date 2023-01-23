/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// PatchItem it contains information on data to be changed.
type PatchItem struct {
	Op PatchOperation `json:"op"`
	// contains a JSON pointer value (as defined in IETF RFC 6901) that references a location of a resource on which the patch operation shall be performed. 
	Path string `json:"path"`
	// indicates the path of the source JSON element (according to JSON Pointer syntax) being moved or copied to the location indicated by the \"path\" attribute. 
	From *string `json:"from,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// NewPatchItem instantiates a new PatchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchItem(op PatchOperation, path string) *PatchItem {
	this := PatchItem{}
	this.Op = op
	this.Path = path
	return &this
}

// NewPatchItemWithDefaults instantiates a new PatchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchItemWithDefaults() *PatchItem {
	this := PatchItem{}
	return &this
}

// GetOp returns the Op field value
func (o *PatchItem) GetOp() PatchOperation {
	if o == nil {
		var ret PatchOperation
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchItem) GetOpOk() (*PatchOperation, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchItem) SetOp(v PatchOperation) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchItem) GetPathOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchItem) SetPath(v string) {
	o.Path = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PatchItem) GetFrom() string {
	if o == nil || isNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchItem) GetFromOk() (*string, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PatchItem) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PatchItem) SetFrom(v string) {
	o.From = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchItem) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchItem) GetValueOk() (*interface{}, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchItem) HasValue() bool {
	if o != nil && isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *PatchItem) SetValue(v interface{}) {
	o.Value = v
}

func (o PatchItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["op"] = o.Op
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePatchItem struct {
	value *PatchItem
	isSet bool
}

func (v NullablePatchItem) Get() *PatchItem {
	return v.value
}

func (v *NullablePatchItem) Set(val *PatchItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchItem(val *PatchItem) *NullablePatchItem {
	return &NullablePatchItem{value: val, isSet: true}
}

func (v NullablePatchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


