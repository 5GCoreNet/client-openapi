/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"os"
)

// CreateUEContext201Response struct for CreateUEContext201Response
type CreateUEContext201Response struct {
	JsonData *UeContextCreatedData `json:"jsonData,omitempty"`
	BinaryDataN2Information **os.File `json:"binaryDataN2Information,omitempty"`
	BinaryDataN2InformationExt1 **os.File `json:"binaryDataN2InformationExt1,omitempty"`
	BinaryDataN2InformationExt2 **os.File `json:"binaryDataN2InformationExt2,omitempty"`
	BinaryDataN2InformationExt3 **os.File `json:"binaryDataN2InformationExt3,omitempty"`
	BinaryDataN2InformationExt4 **os.File `json:"binaryDataN2InformationExt4,omitempty"`
	BinaryDataN2InformationExt5 **os.File `json:"binaryDataN2InformationExt5,omitempty"`
	BinaryDataN2InformationExt6 **os.File `json:"binaryDataN2InformationExt6,omitempty"`
	BinaryDataN2InformationExt7 **os.File `json:"binaryDataN2InformationExt7,omitempty"`
	BinaryDataN2InformationExt8 **os.File `json:"binaryDataN2InformationExt8,omitempty"`
	BinaryDataN2InformationExt9 **os.File `json:"binaryDataN2InformationExt9,omitempty"`
	BinaryDataN2InformationExt10 **os.File `json:"binaryDataN2InformationExt10,omitempty"`
	BinaryDataN2InformationExt11 **os.File `json:"binaryDataN2InformationExt11,omitempty"`
	BinaryDataN2InformationExt12 **os.File `json:"binaryDataN2InformationExt12,omitempty"`
	BinaryDataN2InformationExt13 **os.File `json:"binaryDataN2InformationExt13,omitempty"`
	BinaryDataN2InformationExt14 **os.File `json:"binaryDataN2InformationExt14,omitempty"`
	BinaryDataN2InformationExt15 **os.File `json:"binaryDataN2InformationExt15,omitempty"`
}

// NewCreateUEContext201Response instantiates a new CreateUEContext201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUEContext201Response() *CreateUEContext201Response {
	this := CreateUEContext201Response{}
	return &this
}

// NewCreateUEContext201ResponseWithDefaults instantiates a new CreateUEContext201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUEContext201ResponseWithDefaults() *CreateUEContext201Response {
	this := CreateUEContext201Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetJsonData() UeContextCreatedData {
	if o == nil || isNil(o.JsonData) {
		var ret UeContextCreatedData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetJsonDataOk() (*UeContextCreatedData, bool) {
	if o == nil || isNil(o.JsonData) {
    return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasJsonData() bool {
	if o != nil && !isNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given UeContextCreatedData and assigns it to the JsonData field.
func (o *CreateUEContext201Response) SetJsonData(v UeContextCreatedData) {
	o.JsonData = &v
}

// GetBinaryDataN2Information returns the BinaryDataN2Information field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2Information() *os.File {
	if o == nil || isNil(o.BinaryDataN2Information) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information
}

// GetBinaryDataN2InformationOk returns a tuple with the BinaryDataN2Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationOk() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2Information) {
    return nil, false
	}
	return o.BinaryDataN2Information, true
}

// HasBinaryDataN2Information returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2Information() bool {
	if o != nil && !isNil(o.BinaryDataN2Information) {
		return true
	}

	return false
}

// SetBinaryDataN2Information gets a reference to the given *os.File and assigns it to the BinaryDataN2Information field.
func (o *CreateUEContext201Response) SetBinaryDataN2Information(v *os.File) {
	o.BinaryDataN2Information = &v
}

// GetBinaryDataN2InformationExt1 returns the BinaryDataN2InformationExt1 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt1() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt1) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt1
}

// GetBinaryDataN2InformationExt1Ok returns a tuple with the BinaryDataN2InformationExt1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt1Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt1) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt1, true
}

// HasBinaryDataN2InformationExt1 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt1() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt1) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt1 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt1 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt1(v *os.File) {
	o.BinaryDataN2InformationExt1 = &v
}

// GetBinaryDataN2InformationExt2 returns the BinaryDataN2InformationExt2 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt2() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt2) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt2
}

// GetBinaryDataN2InformationExt2Ok returns a tuple with the BinaryDataN2InformationExt2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt2Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt2) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt2, true
}

// HasBinaryDataN2InformationExt2 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt2() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt2) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt2 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt2 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt2(v *os.File) {
	o.BinaryDataN2InformationExt2 = &v
}

// GetBinaryDataN2InformationExt3 returns the BinaryDataN2InformationExt3 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt3() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt3) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt3
}

// GetBinaryDataN2InformationExt3Ok returns a tuple with the BinaryDataN2InformationExt3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt3Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt3) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt3, true
}

// HasBinaryDataN2InformationExt3 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt3() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt3) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt3 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt3 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt3(v *os.File) {
	o.BinaryDataN2InformationExt3 = &v
}

// GetBinaryDataN2InformationExt4 returns the BinaryDataN2InformationExt4 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt4() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt4) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt4
}

// GetBinaryDataN2InformationExt4Ok returns a tuple with the BinaryDataN2InformationExt4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt4Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt4) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt4, true
}

// HasBinaryDataN2InformationExt4 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt4() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt4) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt4 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt4 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt4(v *os.File) {
	o.BinaryDataN2InformationExt4 = &v
}

// GetBinaryDataN2InformationExt5 returns the BinaryDataN2InformationExt5 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt5() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt5) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt5
}

// GetBinaryDataN2InformationExt5Ok returns a tuple with the BinaryDataN2InformationExt5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt5Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt5) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt5, true
}

// HasBinaryDataN2InformationExt5 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt5() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt5) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt5 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt5 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt5(v *os.File) {
	o.BinaryDataN2InformationExt5 = &v
}

// GetBinaryDataN2InformationExt6 returns the BinaryDataN2InformationExt6 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt6() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt6) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt6
}

// GetBinaryDataN2InformationExt6Ok returns a tuple with the BinaryDataN2InformationExt6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt6Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt6) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt6, true
}

// HasBinaryDataN2InformationExt6 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt6() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt6) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt6 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt6 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt6(v *os.File) {
	o.BinaryDataN2InformationExt6 = &v
}

// GetBinaryDataN2InformationExt7 returns the BinaryDataN2InformationExt7 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt7() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt7) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt7
}

// GetBinaryDataN2InformationExt7Ok returns a tuple with the BinaryDataN2InformationExt7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt7Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt7) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt7, true
}

// HasBinaryDataN2InformationExt7 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt7() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt7) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt7 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt7 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt7(v *os.File) {
	o.BinaryDataN2InformationExt7 = &v
}

// GetBinaryDataN2InformationExt8 returns the BinaryDataN2InformationExt8 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt8() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt8) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt8
}

// GetBinaryDataN2InformationExt8Ok returns a tuple with the BinaryDataN2InformationExt8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt8Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt8) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt8, true
}

// HasBinaryDataN2InformationExt8 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt8() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt8) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt8 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt8 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt8(v *os.File) {
	o.BinaryDataN2InformationExt8 = &v
}

// GetBinaryDataN2InformationExt9 returns the BinaryDataN2InformationExt9 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt9() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt9) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt9
}

// GetBinaryDataN2InformationExt9Ok returns a tuple with the BinaryDataN2InformationExt9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt9Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt9) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt9, true
}

// HasBinaryDataN2InformationExt9 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt9() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt9) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt9 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt9 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt9(v *os.File) {
	o.BinaryDataN2InformationExt9 = &v
}

// GetBinaryDataN2InformationExt10 returns the BinaryDataN2InformationExt10 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt10() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt10) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt10
}

// GetBinaryDataN2InformationExt10Ok returns a tuple with the BinaryDataN2InformationExt10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt10Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt10) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt10, true
}

// HasBinaryDataN2InformationExt10 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt10() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt10) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt10 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt10 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt10(v *os.File) {
	o.BinaryDataN2InformationExt10 = &v
}

// GetBinaryDataN2InformationExt11 returns the BinaryDataN2InformationExt11 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt11() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt11) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt11
}

// GetBinaryDataN2InformationExt11Ok returns a tuple with the BinaryDataN2InformationExt11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt11Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt11) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt11, true
}

// HasBinaryDataN2InformationExt11 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt11() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt11) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt11 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt11 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt11(v *os.File) {
	o.BinaryDataN2InformationExt11 = &v
}

// GetBinaryDataN2InformationExt12 returns the BinaryDataN2InformationExt12 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt12() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt12) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt12
}

// GetBinaryDataN2InformationExt12Ok returns a tuple with the BinaryDataN2InformationExt12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt12Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt12) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt12, true
}

// HasBinaryDataN2InformationExt12 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt12() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt12) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt12 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt12 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt12(v *os.File) {
	o.BinaryDataN2InformationExt12 = &v
}

// GetBinaryDataN2InformationExt13 returns the BinaryDataN2InformationExt13 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt13() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt13) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt13
}

// GetBinaryDataN2InformationExt13Ok returns a tuple with the BinaryDataN2InformationExt13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt13Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt13) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt13, true
}

// HasBinaryDataN2InformationExt13 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt13() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt13) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt13 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt13 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt13(v *os.File) {
	o.BinaryDataN2InformationExt13 = &v
}

// GetBinaryDataN2InformationExt14 returns the BinaryDataN2InformationExt14 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt14() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt14) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt14
}

// GetBinaryDataN2InformationExt14Ok returns a tuple with the BinaryDataN2InformationExt14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt14Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt14) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt14, true
}

// HasBinaryDataN2InformationExt14 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt14() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt14) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt14 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt14 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt14(v *os.File) {
	o.BinaryDataN2InformationExt14 = &v
}

// GetBinaryDataN2InformationExt15 returns the BinaryDataN2InformationExt15 field value if set, zero value otherwise.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt15() *os.File {
	if o == nil || isNil(o.BinaryDataN2InformationExt15) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2InformationExt15
}

// GetBinaryDataN2InformationExt15Ok returns a tuple with the BinaryDataN2InformationExt15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUEContext201Response) GetBinaryDataN2InformationExt15Ok() (**os.File, bool) {
	if o == nil || isNil(o.BinaryDataN2InformationExt15) {
    return nil, false
	}
	return o.BinaryDataN2InformationExt15, true
}

// HasBinaryDataN2InformationExt15 returns a boolean if a field has been set.
func (o *CreateUEContext201Response) HasBinaryDataN2InformationExt15() bool {
	if o != nil && !isNil(o.BinaryDataN2InformationExt15) {
		return true
	}

	return false
}

// SetBinaryDataN2InformationExt15 gets a reference to the given *os.File and assigns it to the BinaryDataN2InformationExt15 field.
func (o *CreateUEContext201Response) SetBinaryDataN2InformationExt15(v *os.File) {
	o.BinaryDataN2InformationExt15 = &v
}

func (o CreateUEContext201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !isNil(o.BinaryDataN2Information) {
		toSerialize["binaryDataN2Information"] = o.BinaryDataN2Information
	}
	if !isNil(o.BinaryDataN2InformationExt1) {
		toSerialize["binaryDataN2InformationExt1"] = o.BinaryDataN2InformationExt1
	}
	if !isNil(o.BinaryDataN2InformationExt2) {
		toSerialize["binaryDataN2InformationExt2"] = o.BinaryDataN2InformationExt2
	}
	if !isNil(o.BinaryDataN2InformationExt3) {
		toSerialize["binaryDataN2InformationExt3"] = o.BinaryDataN2InformationExt3
	}
	if !isNil(o.BinaryDataN2InformationExt4) {
		toSerialize["binaryDataN2InformationExt4"] = o.BinaryDataN2InformationExt4
	}
	if !isNil(o.BinaryDataN2InformationExt5) {
		toSerialize["binaryDataN2InformationExt5"] = o.BinaryDataN2InformationExt5
	}
	if !isNil(o.BinaryDataN2InformationExt6) {
		toSerialize["binaryDataN2InformationExt6"] = o.BinaryDataN2InformationExt6
	}
	if !isNil(o.BinaryDataN2InformationExt7) {
		toSerialize["binaryDataN2InformationExt7"] = o.BinaryDataN2InformationExt7
	}
	if !isNil(o.BinaryDataN2InformationExt8) {
		toSerialize["binaryDataN2InformationExt8"] = o.BinaryDataN2InformationExt8
	}
	if !isNil(o.BinaryDataN2InformationExt9) {
		toSerialize["binaryDataN2InformationExt9"] = o.BinaryDataN2InformationExt9
	}
	if !isNil(o.BinaryDataN2InformationExt10) {
		toSerialize["binaryDataN2InformationExt10"] = o.BinaryDataN2InformationExt10
	}
	if !isNil(o.BinaryDataN2InformationExt11) {
		toSerialize["binaryDataN2InformationExt11"] = o.BinaryDataN2InformationExt11
	}
	if !isNil(o.BinaryDataN2InformationExt12) {
		toSerialize["binaryDataN2InformationExt12"] = o.BinaryDataN2InformationExt12
	}
	if !isNil(o.BinaryDataN2InformationExt13) {
		toSerialize["binaryDataN2InformationExt13"] = o.BinaryDataN2InformationExt13
	}
	if !isNil(o.BinaryDataN2InformationExt14) {
		toSerialize["binaryDataN2InformationExt14"] = o.BinaryDataN2InformationExt14
	}
	if !isNil(o.BinaryDataN2InformationExt15) {
		toSerialize["binaryDataN2InformationExt15"] = o.BinaryDataN2InformationExt15
	}
	return json.Marshal(toSerialize)
}

type NullableCreateUEContext201Response struct {
	value *CreateUEContext201Response
	isSet bool
}

func (v NullableCreateUEContext201Response) Get() *CreateUEContext201Response {
	return v.value
}

func (v *NullableCreateUEContext201Response) Set(val *CreateUEContext201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUEContext201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUEContext201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUEContext201Response(val *CreateUEContext201Response) *NullableCreateUEContext201Response {
	return &NullableCreateUEContext201Response{value: val, isSet: true}
}

func (v NullableCreateUEContext201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUEContext201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


