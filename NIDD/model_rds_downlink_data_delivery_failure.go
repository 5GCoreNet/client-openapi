/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDD

import (
	"encoding/json"
)

// RdsDownlinkDataDeliveryFailure Represents the failure delivery result for RDS.
type RdsDownlinkDataDeliveryFailure struct {
	// string providing an URI formatted according to IETF RFC 3986.
	Type *string `json:"type,omitempty"`
	// A short, human-readable summary of the problem type. It should not change from occurrence to occurrence of the problem.
	Title *string `json:"title,omitempty"`
	// The HTTP status code for this occurrence of the problem.
	Status *int32 `json:"status,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	Instance *string `json:"instance,omitempty"`
	// A machine-readable application error cause specific to this occurrence of the problem. This IE should be present and provide application-related error information, if available.
	Cause *string `json:"cause,omitempty"`
	// Description of invalid parameters, for a request rejected due to invalid parameters.
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	RequestedRetransmissionTime *time.Time `json:"requestedRetransmissionTime,omitempty"`
	// Indicates the serialization format(s) that are supported by the UE on the associated RDS port.
	SupportedUeFormats []SerializationFormat `json:"supportedUeFormats,omitempty"`
}

// NewRdsDownlinkDataDeliveryFailure instantiates a new RdsDownlinkDataDeliveryFailure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRdsDownlinkDataDeliveryFailure() *RdsDownlinkDataDeliveryFailure {
	this := RdsDownlinkDataDeliveryFailure{}
	return &this
}

// NewRdsDownlinkDataDeliveryFailureWithDefaults instantiates a new RdsDownlinkDataDeliveryFailure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRdsDownlinkDataDeliveryFailureWithDefaults() *RdsDownlinkDataDeliveryFailure {
	this := RdsDownlinkDataDeliveryFailure{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RdsDownlinkDataDeliveryFailure) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *RdsDownlinkDataDeliveryFailure) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *RdsDownlinkDataDeliveryFailure) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *RdsDownlinkDataDeliveryFailure) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetInstance() string {
	if o == nil || isNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetInstanceOk() (*string, bool) {
	if o == nil || isNil(o.Instance) {
    return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasInstance() bool {
	if o != nil && !isNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *RdsDownlinkDataDeliveryFailure) SetInstance(v string) {
	o.Instance = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetCause() string {
	if o == nil || isNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetCauseOk() (*string, bool) {
	if o == nil || isNil(o.Cause) {
    return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *RdsDownlinkDataDeliveryFailure) SetCause(v string) {
	o.Cause = &v
}

// GetInvalidParams returns the InvalidParams field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetInvalidParams() []InvalidParam {
	if o == nil || isNil(o.InvalidParams) {
		var ret []InvalidParam
		return ret
	}
	return o.InvalidParams
}

// GetInvalidParamsOk returns a tuple with the InvalidParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetInvalidParamsOk() ([]InvalidParam, bool) {
	if o == nil || isNil(o.InvalidParams) {
    return nil, false
	}
	return o.InvalidParams, true
}

// HasInvalidParams returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasInvalidParams() bool {
	if o != nil && !isNil(o.InvalidParams) {
		return true
	}

	return false
}

// SetInvalidParams gets a reference to the given []InvalidParam and assigns it to the InvalidParams field.
func (o *RdsDownlinkDataDeliveryFailure) SetInvalidParams(v []InvalidParam) {
	o.InvalidParams = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *RdsDownlinkDataDeliveryFailure) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetRequestedRetransmissionTime() time.Time {
	if o == nil || isNil(o.RequestedRetransmissionTime) {
		var ret time.Time
		return ret
	}
	return *o.RequestedRetransmissionTime
}

// GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetRequestedRetransmissionTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.RequestedRetransmissionTime) {
    return nil, false
	}
	return o.RequestedRetransmissionTime, true
}

// HasRequestedRetransmissionTime returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasRequestedRetransmissionTime() bool {
	if o != nil && !isNil(o.RequestedRetransmissionTime) {
		return true
	}

	return false
}

// SetRequestedRetransmissionTime gets a reference to the given time.Time and assigns it to the RequestedRetransmissionTime field.
func (o *RdsDownlinkDataDeliveryFailure) SetRequestedRetransmissionTime(v time.Time) {
	o.RequestedRetransmissionTime = &v
}

// GetSupportedUeFormats returns the SupportedUeFormats field value if set, zero value otherwise.
func (o *RdsDownlinkDataDeliveryFailure) GetSupportedUeFormats() []SerializationFormat {
	if o == nil || isNil(o.SupportedUeFormats) {
		var ret []SerializationFormat
		return ret
	}
	return o.SupportedUeFormats
}

// GetSupportedUeFormatsOk returns a tuple with the SupportedUeFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsDownlinkDataDeliveryFailure) GetSupportedUeFormatsOk() ([]SerializationFormat, bool) {
	if o == nil || isNil(o.SupportedUeFormats) {
    return nil, false
	}
	return o.SupportedUeFormats, true
}

// HasSupportedUeFormats returns a boolean if a field has been set.
func (o *RdsDownlinkDataDeliveryFailure) HasSupportedUeFormats() bool {
	if o != nil && !isNil(o.SupportedUeFormats) {
		return true
	}

	return false
}

// SetSupportedUeFormats gets a reference to the given []SerializationFormat and assigns it to the SupportedUeFormats field.
func (o *RdsDownlinkDataDeliveryFailure) SetSupportedUeFormats(v []SerializationFormat) {
	o.SupportedUeFormats = v
}

func (o RdsDownlinkDataDeliveryFailure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !isNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !isNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !isNil(o.InvalidParams) {
		toSerialize["invalidParams"] = o.InvalidParams
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.RequestedRetransmissionTime) {
		toSerialize["requestedRetransmissionTime"] = o.RequestedRetransmissionTime
	}
	if !isNil(o.SupportedUeFormats) {
		toSerialize["supportedUeFormats"] = o.SupportedUeFormats
	}
	return json.Marshal(toSerialize)
}

type NullableRdsDownlinkDataDeliveryFailure struct {
	value *RdsDownlinkDataDeliveryFailure
	isSet bool
}

func (v NullableRdsDownlinkDataDeliveryFailure) Get() *RdsDownlinkDataDeliveryFailure {
	return v.value
}

func (v *NullableRdsDownlinkDataDeliveryFailure) Set(val *RdsDownlinkDataDeliveryFailure) {
	v.value = val
	v.isSet = true
}

func (v NullableRdsDownlinkDataDeliveryFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableRdsDownlinkDataDeliveryFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRdsDownlinkDataDeliveryFailure(val *RdsDownlinkDataDeliveryFailure) *NullableRdsDownlinkDataDeliveryFailure {
	return &NullableRdsDownlinkDataDeliveryFailure{value: val, isSet: true}
}

func (v NullableRdsDownlinkDataDeliveryFailure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRdsDownlinkDataDeliveryFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


