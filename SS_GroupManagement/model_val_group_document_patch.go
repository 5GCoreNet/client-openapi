/*
SS_GroupManagement

API for SEAL Group management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package SS_GroupManagement

import (
	"encoding/json"
)

// VALGroupDocumentPatch Represents details of the partial update of VAL group document information.
type VALGroupDocumentPatch struct {
	// The text description of the VAL group.
	GrpDesc *string `json:"grpDesc,omitempty"`
	// The list of VAL User IDs or VAL UE IDs, which are members of the VAL group.
	Members []ValTargetUe `json:"members,omitempty"`
	// Configuration data for the VAL group.
	ValGrpConf *string `json:"valGrpConf,omitempty"`
	// The list of VAL services enabled on the group.
	ValServiceIds []string `json:"valServiceIds,omitempty"`
	LocInfo *LocationInfo `json:"locInfo,omitempty"`
	AddLocInfo *LocationArea5G `json:"addLocInfo,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExtGrpId *string `json:"extGrpId,omitempty"`
	Com5GLanType *PduSessionType `json:"com5GLanType,omitempty"`
}

// NewVALGroupDocumentPatch instantiates a new VALGroupDocumentPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVALGroupDocumentPatch() *VALGroupDocumentPatch {
	this := VALGroupDocumentPatch{}
	return &this
}

// NewVALGroupDocumentPatchWithDefaults instantiates a new VALGroupDocumentPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVALGroupDocumentPatchWithDefaults() *VALGroupDocumentPatch {
	this := VALGroupDocumentPatch{}
	return &this
}

// GetGrpDesc returns the GrpDesc field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetGrpDesc() string {
	if o == nil || isNil(o.GrpDesc) {
		var ret string
		return ret
	}
	return *o.GrpDesc
}

// GetGrpDescOk returns a tuple with the GrpDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetGrpDescOk() (*string, bool) {
	if o == nil || isNil(o.GrpDesc) {
    return nil, false
	}
	return o.GrpDesc, true
}

// HasGrpDesc returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasGrpDesc() bool {
	if o != nil && !isNil(o.GrpDesc) {
		return true
	}

	return false
}

// SetGrpDesc gets a reference to the given string and assigns it to the GrpDesc field.
func (o *VALGroupDocumentPatch) SetGrpDesc(v string) {
	o.GrpDesc = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetMembers() []ValTargetUe {
	if o == nil || isNil(o.Members) {
		var ret []ValTargetUe
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetMembersOk() ([]ValTargetUe, bool) {
	if o == nil || isNil(o.Members) {
    return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasMembers() bool {
	if o != nil && !isNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []ValTargetUe and assigns it to the Members field.
func (o *VALGroupDocumentPatch) SetMembers(v []ValTargetUe) {
	o.Members = v
}

// GetValGrpConf returns the ValGrpConf field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetValGrpConf() string {
	if o == nil || isNil(o.ValGrpConf) {
		var ret string
		return ret
	}
	return *o.ValGrpConf
}

// GetValGrpConfOk returns a tuple with the ValGrpConf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetValGrpConfOk() (*string, bool) {
	if o == nil || isNil(o.ValGrpConf) {
    return nil, false
	}
	return o.ValGrpConf, true
}

// HasValGrpConf returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasValGrpConf() bool {
	if o != nil && !isNil(o.ValGrpConf) {
		return true
	}

	return false
}

// SetValGrpConf gets a reference to the given string and assigns it to the ValGrpConf field.
func (o *VALGroupDocumentPatch) SetValGrpConf(v string) {
	o.ValGrpConf = &v
}

// GetValServiceIds returns the ValServiceIds field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetValServiceIds() []string {
	if o == nil || isNil(o.ValServiceIds) {
		var ret []string
		return ret
	}
	return o.ValServiceIds
}

// GetValServiceIdsOk returns a tuple with the ValServiceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetValServiceIdsOk() ([]string, bool) {
	if o == nil || isNil(o.ValServiceIds) {
    return nil, false
	}
	return o.ValServiceIds, true
}

// HasValServiceIds returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasValServiceIds() bool {
	if o != nil && !isNil(o.ValServiceIds) {
		return true
	}

	return false
}

// SetValServiceIds gets a reference to the given []string and assigns it to the ValServiceIds field.
func (o *VALGroupDocumentPatch) SetValServiceIds(v []string) {
	o.ValServiceIds = v
}

// GetLocInfo returns the LocInfo field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetLocInfo() LocationInfo {
	if o == nil || isNil(o.LocInfo) {
		var ret LocationInfo
		return ret
	}
	return *o.LocInfo
}

// GetLocInfoOk returns a tuple with the LocInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetLocInfoOk() (*LocationInfo, bool) {
	if o == nil || isNil(o.LocInfo) {
    return nil, false
	}
	return o.LocInfo, true
}

// HasLocInfo returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasLocInfo() bool {
	if o != nil && !isNil(o.LocInfo) {
		return true
	}

	return false
}

// SetLocInfo gets a reference to the given LocationInfo and assigns it to the LocInfo field.
func (o *VALGroupDocumentPatch) SetLocInfo(v LocationInfo) {
	o.LocInfo = &v
}

// GetAddLocInfo returns the AddLocInfo field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetAddLocInfo() LocationArea5G {
	if o == nil || isNil(o.AddLocInfo) {
		var ret LocationArea5G
		return ret
	}
	return *o.AddLocInfo
}

// GetAddLocInfoOk returns a tuple with the AddLocInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetAddLocInfoOk() (*LocationArea5G, bool) {
	if o == nil || isNil(o.AddLocInfo) {
    return nil, false
	}
	return o.AddLocInfo, true
}

// HasAddLocInfo returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasAddLocInfo() bool {
	if o != nil && !isNil(o.AddLocInfo) {
		return true
	}

	return false
}

// SetAddLocInfo gets a reference to the given LocationArea5G and assigns it to the AddLocInfo field.
func (o *VALGroupDocumentPatch) SetAddLocInfo(v LocationArea5G) {
	o.AddLocInfo = &v
}

// GetExtGrpId returns the ExtGrpId field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetExtGrpId() string {
	if o == nil || isNil(o.ExtGrpId) {
		var ret string
		return ret
	}
	return *o.ExtGrpId
}

// GetExtGrpIdOk returns a tuple with the ExtGrpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetExtGrpIdOk() (*string, bool) {
	if o == nil || isNil(o.ExtGrpId) {
    return nil, false
	}
	return o.ExtGrpId, true
}

// HasExtGrpId returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasExtGrpId() bool {
	if o != nil && !isNil(o.ExtGrpId) {
		return true
	}

	return false
}

// SetExtGrpId gets a reference to the given string and assigns it to the ExtGrpId field.
func (o *VALGroupDocumentPatch) SetExtGrpId(v string) {
	o.ExtGrpId = &v
}

// GetCom5GLanType returns the Com5GLanType field value if set, zero value otherwise.
func (o *VALGroupDocumentPatch) GetCom5GLanType() PduSessionType {
	if o == nil || isNil(o.Com5GLanType) {
		var ret PduSessionType
		return ret
	}
	return *o.Com5GLanType
}

// GetCom5GLanTypeOk returns a tuple with the Com5GLanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VALGroupDocumentPatch) GetCom5GLanTypeOk() (*PduSessionType, bool) {
	if o == nil || isNil(o.Com5GLanType) {
    return nil, false
	}
	return o.Com5GLanType, true
}

// HasCom5GLanType returns a boolean if a field has been set.
func (o *VALGroupDocumentPatch) HasCom5GLanType() bool {
	if o != nil && !isNil(o.Com5GLanType) {
		return true
	}

	return false
}

// SetCom5GLanType gets a reference to the given PduSessionType and assigns it to the Com5GLanType field.
func (o *VALGroupDocumentPatch) SetCom5GLanType(v PduSessionType) {
	o.Com5GLanType = &v
}

func (o VALGroupDocumentPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GrpDesc) {
		toSerialize["grpDesc"] = o.GrpDesc
	}
	if !isNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !isNil(o.ValGrpConf) {
		toSerialize["valGrpConf"] = o.ValGrpConf
	}
	if !isNil(o.ValServiceIds) {
		toSerialize["valServiceIds"] = o.ValServiceIds
	}
	if !isNil(o.LocInfo) {
		toSerialize["locInfo"] = o.LocInfo
	}
	if !isNil(o.AddLocInfo) {
		toSerialize["addLocInfo"] = o.AddLocInfo
	}
	if !isNil(o.ExtGrpId) {
		toSerialize["extGrpId"] = o.ExtGrpId
	}
	if !isNil(o.Com5GLanType) {
		toSerialize["com5GLanType"] = o.Com5GLanType
	}
	return json.Marshal(toSerialize)
}

type NullableVALGroupDocumentPatch struct {
	value *VALGroupDocumentPatch
	isSet bool
}

func (v NullableVALGroupDocumentPatch) Get() *VALGroupDocumentPatch {
	return v.value
}

func (v *NullableVALGroupDocumentPatch) Set(val *VALGroupDocumentPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableVALGroupDocumentPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableVALGroupDocumentPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVALGroupDocumentPatch(val *VALGroupDocumentPatch) *NullableVALGroupDocumentPatch {
	return &NullableVALGroupDocumentPatch{value: val, isSet: true}
}

func (v NullableVALGroupDocumentPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVALGroupDocumentPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


