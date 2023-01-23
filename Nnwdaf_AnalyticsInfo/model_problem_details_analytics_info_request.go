/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// ProblemDetailsAnalyticsInfoRequest Extends ProblemDetails to indicate more details why the analytics request is rejected. 
type ProblemDetailsAnalyticsInfoRequest struct {
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
	// indicating a time in seconds.
	RvWaitTime *int32 `json:"rvWaitTime,omitempty"`
}

// NewProblemDetailsAnalyticsInfoRequest instantiates a new ProblemDetailsAnalyticsInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemDetailsAnalyticsInfoRequest() *ProblemDetailsAnalyticsInfoRequest {
	this := ProblemDetailsAnalyticsInfoRequest{}
	return &this
}

// NewProblemDetailsAnalyticsInfoRequestWithDefaults instantiates a new ProblemDetailsAnalyticsInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemDetailsAnalyticsInfoRequestWithDefaults() *ProblemDetailsAnalyticsInfoRequest {
	this := ProblemDetailsAnalyticsInfoRequest{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetInstance() string {
	if o == nil || isNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetInstanceOk() (*string, bool) {
	if o == nil || isNil(o.Instance) {
    return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasInstance() bool {
	if o != nil && !isNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetInstance(v string) {
	o.Instance = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetCause() string {
	if o == nil || isNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetCauseOk() (*string, bool) {
	if o == nil || isNil(o.Cause) {
    return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetCause(v string) {
	o.Cause = &v
}

// GetInvalidParams returns the InvalidParams field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetInvalidParams() []InvalidParam {
	if o == nil || isNil(o.InvalidParams) {
		var ret []InvalidParam
		return ret
	}
	return o.InvalidParams
}

// GetInvalidParamsOk returns a tuple with the InvalidParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetInvalidParamsOk() ([]InvalidParam, bool) {
	if o == nil || isNil(o.InvalidParams) {
    return nil, false
	}
	return o.InvalidParams, true
}

// HasInvalidParams returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasInvalidParams() bool {
	if o != nil && !isNil(o.InvalidParams) {
		return true
	}

	return false
}

// SetInvalidParams gets a reference to the given []InvalidParam and assigns it to the InvalidParams field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetInvalidParams(v []InvalidParam) {
	o.InvalidParams = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAccessTokenError returns the AccessTokenError field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenError() AccessTokenErr {
	if o == nil || isNil(o.AccessTokenError) {
		var ret AccessTokenErr
		return ret
	}
	return *o.AccessTokenError
}

// GetAccessTokenErrorOk returns a tuple with the AccessTokenError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenErrorOk() (*AccessTokenErr, bool) {
	if o == nil || isNil(o.AccessTokenError) {
    return nil, false
	}
	return o.AccessTokenError, true
}

// HasAccessTokenError returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasAccessTokenError() bool {
	if o != nil && !isNil(o.AccessTokenError) {
		return true
	}

	return false
}

// SetAccessTokenError gets a reference to the given AccessTokenErr and assigns it to the AccessTokenError field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetAccessTokenError(v AccessTokenErr) {
	o.AccessTokenError = &v
}

// GetAccessTokenRequest returns the AccessTokenRequest field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenRequest() AccessTokenReq {
	if o == nil || isNil(o.AccessTokenRequest) {
		var ret AccessTokenReq
		return ret
	}
	return *o.AccessTokenRequest
}

// GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetAccessTokenRequestOk() (*AccessTokenReq, bool) {
	if o == nil || isNil(o.AccessTokenRequest) {
    return nil, false
	}
	return o.AccessTokenRequest, true
}

// HasAccessTokenRequest returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasAccessTokenRequest() bool {
	if o != nil && !isNil(o.AccessTokenRequest) {
		return true
	}

	return false
}

// SetAccessTokenRequest gets a reference to the given AccessTokenReq and assigns it to the AccessTokenRequest field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetAccessTokenRequest(v AccessTokenReq) {
	o.AccessTokenRequest = &v
}

// GetNrfId returns the NrfId field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetNrfId() string {
	if o == nil || isNil(o.NrfId) {
		var ret string
		return ret
	}
	return *o.NrfId
}

// GetNrfIdOk returns a tuple with the NrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetNrfIdOk() (*string, bool) {
	if o == nil || isNil(o.NrfId) {
    return nil, false
	}
	return o.NrfId, true
}

// HasNrfId returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasNrfId() bool {
	if o != nil && !isNil(o.NrfId) {
		return true
	}

	return false
}

// SetNrfId gets a reference to the given string and assigns it to the NrfId field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetNrfId(v string) {
	o.NrfId = &v
}

// GetRvWaitTime returns the RvWaitTime field value if set, zero value otherwise.
func (o *ProblemDetailsAnalyticsInfoRequest) GetRvWaitTime() int32 {
	if o == nil || isNil(o.RvWaitTime) {
		var ret int32
		return ret
	}
	return *o.RvWaitTime
}

// GetRvWaitTimeOk returns a tuple with the RvWaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) GetRvWaitTimeOk() (*int32, bool) {
	if o == nil || isNil(o.RvWaitTime) {
    return nil, false
	}
	return o.RvWaitTime, true
}

// HasRvWaitTime returns a boolean if a field has been set.
func (o *ProblemDetailsAnalyticsInfoRequest) HasRvWaitTime() bool {
	if o != nil && !isNil(o.RvWaitTime) {
		return true
	}

	return false
}

// SetRvWaitTime gets a reference to the given int32 and assigns it to the RvWaitTime field.
func (o *ProblemDetailsAnalyticsInfoRequest) SetRvWaitTime(v int32) {
	o.RvWaitTime = &v
}

func (o ProblemDetailsAnalyticsInfoRequest) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.RvWaitTime) {
		toSerialize["rvWaitTime"] = o.RvWaitTime
	}
	return json.Marshal(toSerialize)
}

type NullableProblemDetailsAnalyticsInfoRequest struct {
	value *ProblemDetailsAnalyticsInfoRequest
	isSet bool
}

func (v NullableProblemDetailsAnalyticsInfoRequest) Get() *ProblemDetailsAnalyticsInfoRequest {
	return v.value
}

func (v *NullableProblemDetailsAnalyticsInfoRequest) Set(val *ProblemDetailsAnalyticsInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemDetailsAnalyticsInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemDetailsAnalyticsInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemDetailsAnalyticsInfoRequest(val *ProblemDetailsAnalyticsInfoRequest) *NullableProblemDetailsAnalyticsInfoRequest {
	return &NullableProblemDetailsAnalyticsInfoRequest{value: val, isSet: true}
}

func (v NullableProblemDetailsAnalyticsInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemDetailsAnalyticsInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


