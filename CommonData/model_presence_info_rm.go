/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// PresenceInfoRm This data type is defined in the same way as the 'PresenceInfo' data type, but with the OpenAPI 'nullable: true' property.  If the additionalPraId IE is present, this IE shall state the presence information of the UE for the individual PRA identified by the additionalPraId IE;  If the additionalPraId IE is not present, this IE shall state the presence information of the UE for the PRA identified by the praId IE.  
type PresenceInfoRm struct {
	// Represents an identifier of the Presence Reporting Area (see clause 28.10 of  3GPP TS 23.003. This IE shall be present  if the Area of Interest subscribed or reported is a Presence Reporting Area or a Set of Core Network predefined Presence Reporting Areas. When present, it shall be encoded as a string representing an integer in the following ranges: - 0 to 8 388 607 for UE-dedicated PRA - 8 388 608 to 16 777 215 for Core Network predefined PRA Examples: PRA ID 123 is encoded as \"123\" PRA ID 11 238 660 is encoded as \"11238660\" 
	PraId *string `json:"praId,omitempty"`
	// This IE may be present if the praId IE is present and if it contains a PRA identifier referring to a set of Core Network predefined Presence Reporting Areas. When present, this IE shall contain a PRA Identifier of an individual PRA within the Set of Core Network predefined Presence Reporting Areas indicated by the praId IE. 
	AdditionalPraId *string `json:"additionalPraId,omitempty"`
	PresenceState *PresenceState `json:"presenceState,omitempty"`
	// Represents the list of tracking areas that constitutes the area. This IE shall be present if the subscription or the event report  is for tracking UE presence in the tracking areas. For non 3GPP access the TAI shall be the N3GPP TAI. 
	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`
	// Represents the list of EUTRAN cell Ids that constitutes the area. This IE shall be present if the Area of Interest subscribed is a list of EUTRAN cell Ids. 
	EcgiList []Ecgi `json:"ecgiList,omitempty"`
	// Represents the list of NR cell Ids that constitutes the area. This IE shall be present if the Area of Interest subscribed is a list of NR cell Ids. 
	NcgiList []Ncgi `json:"ncgiList,omitempty"`
	// Represents the list of NG RAN node identifiers that constitutes the area. This IE shall be present if the Area of Interest subscribed is a list of NG RAN node identifiers. 
	GlobalRanNodeIdList []GlobalRanNodeId `json:"globalRanNodeIdList,omitempty"`
	// Represents the list of eNodeB identifiers that constitutes the area. This IE shall be present if the Area of Interest subscribed is a list of eNodeB identifiers. 
	GlobaleNbIdList []GlobalRanNodeId `json:"globaleNbIdList,omitempty"`
}

// NewPresenceInfoRm instantiates a new PresenceInfoRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPresenceInfoRm() *PresenceInfoRm {
	this := PresenceInfoRm{}
	return &this
}

// NewPresenceInfoRmWithDefaults instantiates a new PresenceInfoRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPresenceInfoRmWithDefaults() *PresenceInfoRm {
	this := PresenceInfoRm{}
	return &this
}

// GetPraId returns the PraId field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetPraId() string {
	if o == nil || isNil(o.PraId) {
		var ret string
		return ret
	}
	return *o.PraId
}

// GetPraIdOk returns a tuple with the PraId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetPraIdOk() (*string, bool) {
	if o == nil || isNil(o.PraId) {
    return nil, false
	}
	return o.PraId, true
}

// HasPraId returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasPraId() bool {
	if o != nil && !isNil(o.PraId) {
		return true
	}

	return false
}

// SetPraId gets a reference to the given string and assigns it to the PraId field.
func (o *PresenceInfoRm) SetPraId(v string) {
	o.PraId = &v
}

// GetAdditionalPraId returns the AdditionalPraId field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetAdditionalPraId() string {
	if o == nil || isNil(o.AdditionalPraId) {
		var ret string
		return ret
	}
	return *o.AdditionalPraId
}

// GetAdditionalPraIdOk returns a tuple with the AdditionalPraId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetAdditionalPraIdOk() (*string, bool) {
	if o == nil || isNil(o.AdditionalPraId) {
    return nil, false
	}
	return o.AdditionalPraId, true
}

// HasAdditionalPraId returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasAdditionalPraId() bool {
	if o != nil && !isNil(o.AdditionalPraId) {
		return true
	}

	return false
}

// SetAdditionalPraId gets a reference to the given string and assigns it to the AdditionalPraId field.
func (o *PresenceInfoRm) SetAdditionalPraId(v string) {
	o.AdditionalPraId = &v
}

// GetPresenceState returns the PresenceState field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetPresenceState() PresenceState {
	if o == nil || isNil(o.PresenceState) {
		var ret PresenceState
		return ret
	}
	return *o.PresenceState
}

// GetPresenceStateOk returns a tuple with the PresenceState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetPresenceStateOk() (*PresenceState, bool) {
	if o == nil || isNil(o.PresenceState) {
    return nil, false
	}
	return o.PresenceState, true
}

// HasPresenceState returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasPresenceState() bool {
	if o != nil && !isNil(o.PresenceState) {
		return true
	}

	return false
}

// SetPresenceState gets a reference to the given PresenceState and assigns it to the PresenceState field.
func (o *PresenceInfoRm) SetPresenceState(v PresenceState) {
	o.PresenceState = &v
}

// GetTrackingAreaList returns the TrackingAreaList field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetTrackingAreaList() []Tai {
	if o == nil || isNil(o.TrackingAreaList) {
		var ret []Tai
		return ret
	}
	return o.TrackingAreaList
}

// GetTrackingAreaListOk returns a tuple with the TrackingAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetTrackingAreaListOk() ([]Tai, bool) {
	if o == nil || isNil(o.TrackingAreaList) {
    return nil, false
	}
	return o.TrackingAreaList, true
}

// HasTrackingAreaList returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasTrackingAreaList() bool {
	if o != nil && !isNil(o.TrackingAreaList) {
		return true
	}

	return false
}

// SetTrackingAreaList gets a reference to the given []Tai and assigns it to the TrackingAreaList field.
func (o *PresenceInfoRm) SetTrackingAreaList(v []Tai) {
	o.TrackingAreaList = v
}

// GetEcgiList returns the EcgiList field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetEcgiList() []Ecgi {
	if o == nil || isNil(o.EcgiList) {
		var ret []Ecgi
		return ret
	}
	return o.EcgiList
}

// GetEcgiListOk returns a tuple with the EcgiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetEcgiListOk() ([]Ecgi, bool) {
	if o == nil || isNil(o.EcgiList) {
    return nil, false
	}
	return o.EcgiList, true
}

// HasEcgiList returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasEcgiList() bool {
	if o != nil && !isNil(o.EcgiList) {
		return true
	}

	return false
}

// SetEcgiList gets a reference to the given []Ecgi and assigns it to the EcgiList field.
func (o *PresenceInfoRm) SetEcgiList(v []Ecgi) {
	o.EcgiList = v
}

// GetNcgiList returns the NcgiList field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetNcgiList() []Ncgi {
	if o == nil || isNil(o.NcgiList) {
		var ret []Ncgi
		return ret
	}
	return o.NcgiList
}

// GetNcgiListOk returns a tuple with the NcgiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetNcgiListOk() ([]Ncgi, bool) {
	if o == nil || isNil(o.NcgiList) {
    return nil, false
	}
	return o.NcgiList, true
}

// HasNcgiList returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasNcgiList() bool {
	if o != nil && !isNil(o.NcgiList) {
		return true
	}

	return false
}

// SetNcgiList gets a reference to the given []Ncgi and assigns it to the NcgiList field.
func (o *PresenceInfoRm) SetNcgiList(v []Ncgi) {
	o.NcgiList = v
}

// GetGlobalRanNodeIdList returns the GlobalRanNodeIdList field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetGlobalRanNodeIdList() []GlobalRanNodeId {
	if o == nil || isNil(o.GlobalRanNodeIdList) {
		var ret []GlobalRanNodeId
		return ret
	}
	return o.GlobalRanNodeIdList
}

// GetGlobalRanNodeIdListOk returns a tuple with the GlobalRanNodeIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetGlobalRanNodeIdListOk() ([]GlobalRanNodeId, bool) {
	if o == nil || isNil(o.GlobalRanNodeIdList) {
    return nil, false
	}
	return o.GlobalRanNodeIdList, true
}

// HasGlobalRanNodeIdList returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasGlobalRanNodeIdList() bool {
	if o != nil && !isNil(o.GlobalRanNodeIdList) {
		return true
	}

	return false
}

// SetGlobalRanNodeIdList gets a reference to the given []GlobalRanNodeId and assigns it to the GlobalRanNodeIdList field.
func (o *PresenceInfoRm) SetGlobalRanNodeIdList(v []GlobalRanNodeId) {
	o.GlobalRanNodeIdList = v
}

// GetGlobaleNbIdList returns the GlobaleNbIdList field value if set, zero value otherwise.
func (o *PresenceInfoRm) GetGlobaleNbIdList() []GlobalRanNodeId {
	if o == nil || isNil(o.GlobaleNbIdList) {
		var ret []GlobalRanNodeId
		return ret
	}
	return o.GlobaleNbIdList
}

// GetGlobaleNbIdListOk returns a tuple with the GlobaleNbIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PresenceInfoRm) GetGlobaleNbIdListOk() ([]GlobalRanNodeId, bool) {
	if o == nil || isNil(o.GlobaleNbIdList) {
    return nil, false
	}
	return o.GlobaleNbIdList, true
}

// HasGlobaleNbIdList returns a boolean if a field has been set.
func (o *PresenceInfoRm) HasGlobaleNbIdList() bool {
	if o != nil && !isNil(o.GlobaleNbIdList) {
		return true
	}

	return false
}

// SetGlobaleNbIdList gets a reference to the given []GlobalRanNodeId and assigns it to the GlobaleNbIdList field.
func (o *PresenceInfoRm) SetGlobaleNbIdList(v []GlobalRanNodeId) {
	o.GlobaleNbIdList = v
}

func (o PresenceInfoRm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PraId) {
		toSerialize["praId"] = o.PraId
	}
	if !isNil(o.AdditionalPraId) {
		toSerialize["additionalPraId"] = o.AdditionalPraId
	}
	if !isNil(o.PresenceState) {
		toSerialize["presenceState"] = o.PresenceState
	}
	if !isNil(o.TrackingAreaList) {
		toSerialize["trackingAreaList"] = o.TrackingAreaList
	}
	if !isNil(o.EcgiList) {
		toSerialize["ecgiList"] = o.EcgiList
	}
	if !isNil(o.NcgiList) {
		toSerialize["ncgiList"] = o.NcgiList
	}
	if !isNil(o.GlobalRanNodeIdList) {
		toSerialize["globalRanNodeIdList"] = o.GlobalRanNodeIdList
	}
	if !isNil(o.GlobaleNbIdList) {
		toSerialize["globaleNbIdList"] = o.GlobaleNbIdList
	}
	return json.Marshal(toSerialize)
}

type NullablePresenceInfoRm struct {
	value *PresenceInfoRm
	isSet bool
}

func (v NullablePresenceInfoRm) Get() *PresenceInfoRm {
	return v.value
}

func (v *NullablePresenceInfoRm) Set(val *PresenceInfoRm) {
	v.value = val
	v.isSet = true
}

func (v NullablePresenceInfoRm) IsSet() bool {
	return v.isSet
}

func (v *NullablePresenceInfoRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresenceInfoRm(val *PresenceInfoRm) *NullablePresenceInfoRm {
	return &NullablePresenceInfoRm{value: val, isSet: true}
}

func (v NullablePresenceInfoRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresenceInfoRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


