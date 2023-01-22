/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Nsmf_PDUSession

import (
	"encoding/json"
)

// TransferMtDataError Transfer MT Data Error Response
type TransferMtDataError struct {
	// String providing an URI formatted according to RFC 3986.
	Type *string `json:"type,omitempty"`
	Title *string `json:"title,omitempty"`
	Status *int32 `json:"status,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	Instance *string `json:"instance,omitempty"`
	// A machine-readable application error cause specific to this occurrence of the problem.  This IE should be present and provide application-related error information, if available. 
	Cause *string `json:"cause,omitempty"`
	InvalidParams []InvalidParam `json:"invalidParams,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	AccessTokenError *AccessTokenErr `json:"accessTokenError,omitempty"`
	AccessTokenRequest *AccessTokenReq `json:"accessTokenRequest,omitempty"`
	// Fully Qualified Domain Name
	NrfId *string `json:"nrfId,omitempty"`
	RemoteError *bool `json:"remoteError,omitempty"`
	// indicating a time in seconds.
	MaxWaitingTime *int32 `json:"maxWaitingTime,omitempty"`
}

// NewTransferMtDataError instantiates a new TransferMtDataError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferMtDataError() *TransferMtDataError {
	this := TransferMtDataError{}
	return &this
}

// NewTransferMtDataErrorWithDefaults instantiates a new TransferMtDataError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferMtDataErrorWithDefaults() *TransferMtDataError {
	this := TransferMtDataError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransferMtDataError) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransferMtDataError) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransferMtDataError) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TransferMtDataError) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TransferMtDataError) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TransferMtDataError) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TransferMtDataError) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TransferMtDataError) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *TransferMtDataError) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *TransferMtDataError) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *TransferMtDataError) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *TransferMtDataError) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *TransferMtDataError) GetInstance() string {
	if o == nil || isNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetInstanceOk() (*string, bool) {
	if o == nil || isNil(o.Instance) {
    return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *TransferMtDataError) HasInstance() bool {
	if o != nil && !isNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *TransferMtDataError) SetInstance(v string) {
	o.Instance = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *TransferMtDataError) GetCause() string {
	if o == nil || isNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetCauseOk() (*string, bool) {
	if o == nil || isNil(o.Cause) {
    return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *TransferMtDataError) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *TransferMtDataError) SetCause(v string) {
	o.Cause = &v
}

// GetInvalidParams returns the InvalidParams field value if set, zero value otherwise.
func (o *TransferMtDataError) GetInvalidParams() []InvalidParam {
	if o == nil || isNil(o.InvalidParams) {
		var ret []InvalidParam
		return ret
	}
	return o.InvalidParams
}

// GetInvalidParamsOk returns a tuple with the InvalidParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetInvalidParamsOk() ([]InvalidParam, bool) {
	if o == nil || isNil(o.InvalidParams) {
    return nil, false
	}
	return o.InvalidParams, true
}

// HasInvalidParams returns a boolean if a field has been set.
func (o *TransferMtDataError) HasInvalidParams() bool {
	if o != nil && !isNil(o.InvalidParams) {
		return true
	}

	return false
}

// SetInvalidParams gets a reference to the given []InvalidParam and assigns it to the InvalidParams field.
func (o *TransferMtDataError) SetInvalidParams(v []InvalidParam) {
	o.InvalidParams = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *TransferMtDataError) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *TransferMtDataError) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *TransferMtDataError) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAccessTokenError returns the AccessTokenError field value if set, zero value otherwise.
func (o *TransferMtDataError) GetAccessTokenError() AccessTokenErr {
	if o == nil || isNil(o.AccessTokenError) {
		var ret AccessTokenErr
		return ret
	}
	return *o.AccessTokenError
}

// GetAccessTokenErrorOk returns a tuple with the AccessTokenError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetAccessTokenErrorOk() (*AccessTokenErr, bool) {
	if o == nil || isNil(o.AccessTokenError) {
    return nil, false
	}
	return o.AccessTokenError, true
}

// HasAccessTokenError returns a boolean if a field has been set.
func (o *TransferMtDataError) HasAccessTokenError() bool {
	if o != nil && !isNil(o.AccessTokenError) {
		return true
	}

	return false
}

// SetAccessTokenError gets a reference to the given AccessTokenErr and assigns it to the AccessTokenError field.
func (o *TransferMtDataError) SetAccessTokenError(v AccessTokenErr) {
	o.AccessTokenError = &v
}

// GetAccessTokenRequest returns the AccessTokenRequest field value if set, zero value otherwise.
func (o *TransferMtDataError) GetAccessTokenRequest() AccessTokenReq {
	if o == nil || isNil(o.AccessTokenRequest) {
		var ret AccessTokenReq
		return ret
	}
	return *o.AccessTokenRequest
}

// GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetAccessTokenRequestOk() (*AccessTokenReq, bool) {
	if o == nil || isNil(o.AccessTokenRequest) {
    return nil, false
	}
	return o.AccessTokenRequest, true
}

// HasAccessTokenRequest returns a boolean if a field has been set.
func (o *TransferMtDataError) HasAccessTokenRequest() bool {
	if o != nil && !isNil(o.AccessTokenRequest) {
		return true
	}

	return false
}

// SetAccessTokenRequest gets a reference to the given AccessTokenReq and assigns it to the AccessTokenRequest field.
func (o *TransferMtDataError) SetAccessTokenRequest(v AccessTokenReq) {
	o.AccessTokenRequest = &v
}

// GetNrfId returns the NrfId field value if set, zero value otherwise.
func (o *TransferMtDataError) GetNrfId() string {
	if o == nil || isNil(o.NrfId) {
		var ret string
		return ret
	}
	return *o.NrfId
}

// GetNrfIdOk returns a tuple with the NrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetNrfIdOk() (*string, bool) {
	if o == nil || isNil(o.NrfId) {
    return nil, false
	}
	return o.NrfId, true
}

// HasNrfId returns a boolean if a field has been set.
func (o *TransferMtDataError) HasNrfId() bool {
	if o != nil && !isNil(o.NrfId) {
		return true
	}

	return false
}

// SetNrfId gets a reference to the given string and assigns it to the NrfId field.
func (o *TransferMtDataError) SetNrfId(v string) {
	o.NrfId = &v
}

// GetRemoteError returns the RemoteError field value if set, zero value otherwise.
func (o *TransferMtDataError) GetRemoteError() bool {
	if o == nil || isNil(o.RemoteError) {
		var ret bool
		return ret
	}
	return *o.RemoteError
}

// GetRemoteErrorOk returns a tuple with the RemoteError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetRemoteErrorOk() (*bool, bool) {
	if o == nil || isNil(o.RemoteError) {
    return nil, false
	}
	return o.RemoteError, true
}

// HasRemoteError returns a boolean if a field has been set.
func (o *TransferMtDataError) HasRemoteError() bool {
	if o != nil && !isNil(o.RemoteError) {
		return true
	}

	return false
}

// SetRemoteError gets a reference to the given bool and assigns it to the RemoteError field.
func (o *TransferMtDataError) SetRemoteError(v bool) {
	o.RemoteError = &v
}

// GetMaxWaitingTime returns the MaxWaitingTime field value if set, zero value otherwise.
func (o *TransferMtDataError) GetMaxWaitingTime() int32 {
	if o == nil || isNil(o.MaxWaitingTime) {
		var ret int32
		return ret
	}
	return *o.MaxWaitingTime
}

// GetMaxWaitingTimeOk returns a tuple with the MaxWaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferMtDataError) GetMaxWaitingTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxWaitingTime) {
    return nil, false
	}
	return o.MaxWaitingTime, true
}

// HasMaxWaitingTime returns a boolean if a field has been set.
func (o *TransferMtDataError) HasMaxWaitingTime() bool {
	if o != nil && !isNil(o.MaxWaitingTime) {
		return true
	}

	return false
}

// SetMaxWaitingTime gets a reference to the given int32 and assigns it to the MaxWaitingTime field.
func (o *TransferMtDataError) SetMaxWaitingTime(v int32) {
	o.MaxWaitingTime = &v
}

func (o TransferMtDataError) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.AccessTokenError) {
		toSerialize["accessTokenError"] = o.AccessTokenError
	}
	if !isNil(o.AccessTokenRequest) {
		toSerialize["accessTokenRequest"] = o.AccessTokenRequest
	}
	if !isNil(o.NrfId) {
		toSerialize["nrfId"] = o.NrfId
	}
	if !isNil(o.RemoteError) {
		toSerialize["remoteError"] = o.RemoteError
	}
	if !isNil(o.MaxWaitingTime) {
		toSerialize["maxWaitingTime"] = o.MaxWaitingTime
	}
	return json.Marshal(toSerialize)
}

type NullableTransferMtDataError struct {
	value *TransferMtDataError
	isSet bool
}

func (v NullableTransferMtDataError) Get() *TransferMtDataError {
	return v.value
}

func (v *NullableTransferMtDataError) Set(val *TransferMtDataError) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferMtDataError) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferMtDataError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferMtDataError(val *TransferMtDataError) *NullableTransferMtDataError {
	return &NullableTransferMtDataError{value: val, isSet: true}
}

func (v NullableTransferMtDataError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferMtDataError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


