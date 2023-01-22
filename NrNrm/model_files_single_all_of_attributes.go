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

// FilesSingleAllOfAttributes struct for FilesSingleAllOfAttributes
type FilesSingleAllOfAttributes struct {
	NumberOfFiles *int32 `json:"numberOfFiles,omitempty"`
	JobRef *string `json:"jobRef,omitempty"`
	JobId *string `json:"jobId,omitempty"`
	File []FileSingle `json:"File,omitempty"`
}

// NewFilesSingleAllOfAttributes instantiates a new FilesSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilesSingleAllOfAttributes() *FilesSingleAllOfAttributes {
	this := FilesSingleAllOfAttributes{}
	return &this
}

// NewFilesSingleAllOfAttributesWithDefaults instantiates a new FilesSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilesSingleAllOfAttributesWithDefaults() *FilesSingleAllOfAttributes {
	this := FilesSingleAllOfAttributes{}
	return &this
}

// GetNumberOfFiles returns the NumberOfFiles field value if set, zero value otherwise.
func (o *FilesSingleAllOfAttributes) GetNumberOfFiles() int32 {
	if o == nil || isNil(o.NumberOfFiles) {
		var ret int32
		return ret
	}
	return *o.NumberOfFiles
}

// GetNumberOfFilesOk returns a tuple with the NumberOfFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSingleAllOfAttributes) GetNumberOfFilesOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfFiles) {
    return nil, false
	}
	return o.NumberOfFiles, true
}

// HasNumberOfFiles returns a boolean if a field has been set.
func (o *FilesSingleAllOfAttributes) HasNumberOfFiles() bool {
	if o != nil && !isNil(o.NumberOfFiles) {
		return true
	}

	return false
}

// SetNumberOfFiles gets a reference to the given int32 and assigns it to the NumberOfFiles field.
func (o *FilesSingleAllOfAttributes) SetNumberOfFiles(v int32) {
	o.NumberOfFiles = &v
}

// GetJobRef returns the JobRef field value if set, zero value otherwise.
func (o *FilesSingleAllOfAttributes) GetJobRef() string {
	if o == nil || isNil(o.JobRef) {
		var ret string
		return ret
	}
	return *o.JobRef
}

// GetJobRefOk returns a tuple with the JobRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSingleAllOfAttributes) GetJobRefOk() (*string, bool) {
	if o == nil || isNil(o.JobRef) {
    return nil, false
	}
	return o.JobRef, true
}

// HasJobRef returns a boolean if a field has been set.
func (o *FilesSingleAllOfAttributes) HasJobRef() bool {
	if o != nil && !isNil(o.JobRef) {
		return true
	}

	return false
}

// SetJobRef gets a reference to the given string and assigns it to the JobRef field.
func (o *FilesSingleAllOfAttributes) SetJobRef(v string) {
	o.JobRef = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *FilesSingleAllOfAttributes) GetJobId() string {
	if o == nil || isNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSingleAllOfAttributes) GetJobIdOk() (*string, bool) {
	if o == nil || isNil(o.JobId) {
    return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *FilesSingleAllOfAttributes) HasJobId() bool {
	if o != nil && !isNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *FilesSingleAllOfAttributes) SetJobId(v string) {
	o.JobId = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *FilesSingleAllOfAttributes) GetFile() []FileSingle {
	if o == nil || isNil(o.File) {
		var ret []FileSingle
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilesSingleAllOfAttributes) GetFileOk() ([]FileSingle, bool) {
	if o == nil || isNil(o.File) {
    return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *FilesSingleAllOfAttributes) HasFile() bool {
	if o != nil && !isNil(o.File) {
		return true
	}

	return false
}

// SetFile gets a reference to the given []FileSingle and assigns it to the File field.
func (o *FilesSingleAllOfAttributes) SetFile(v []FileSingle) {
	o.File = v
}

func (o FilesSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.NumberOfFiles) {
		toSerialize["numberOfFiles"] = o.NumberOfFiles
	}
	if !isNil(o.JobRef) {
		toSerialize["jobRef"] = o.JobRef
	}
	if !isNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	if !isNil(o.File) {
		toSerialize["File"] = o.File
	}
	return json.Marshal(toSerialize)
}

type NullableFilesSingleAllOfAttributes struct {
	value *FilesSingleAllOfAttributes
	isSet bool
}

func (v NullableFilesSingleAllOfAttributes) Get() *FilesSingleAllOfAttributes {
	return v.value
}

func (v *NullableFilesSingleAllOfAttributes) Set(val *FilesSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFilesSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFilesSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilesSingleAllOfAttributes(val *FilesSingleAllOfAttributes) *NullableFilesSingleAllOfAttributes {
	return &NullableFilesSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableFilesSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilesSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


