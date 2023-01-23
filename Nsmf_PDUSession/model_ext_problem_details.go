/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// ExtProblemDetails Extended Problem Details
type ExtProblemDetails struct {
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
}

// NewExtProblemDetails instantiates a new ExtProblemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtProblemDetails() *ExtProblemDetails {
	this := ExtProblemDetails{}
	return &this
}

// NewExtProblemDetailsWithDefaults instantiates a new ExtProblemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtProblemDetailsWithDefaults() *ExtProblemDetails {
	this := ExtProblemDetails{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExtProblemDetails) SetType(v string) {
	o.Type = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ExtProblemDetails) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ExtProblemDetails) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetDetail() string {
	if o == nil || isNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetDetailOk() (*string, bool) {
	if o == nil || isNil(o.Detail) {
    return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasDetail() bool {
	if o != nil && !isNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ExtProblemDetails) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetInstance() string {
	if o == nil || isNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetInstanceOk() (*string, bool) {
	if o == nil || isNil(o.Instance) {
    return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasInstance() bool {
	if o != nil && !isNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ExtProblemDetails) SetInstance(v string) {
	o.Instance = &v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetCause() string {
	if o == nil || isNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetCauseOk() (*string, bool) {
	if o == nil || isNil(o.Cause) {
    return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *ExtProblemDetails) SetCause(v string) {
	o.Cause = &v
}

// GetInvalidParams returns the InvalidParams field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetInvalidParams() []InvalidParam {
	if o == nil || isNil(o.InvalidParams) {
		var ret []InvalidParam
		return ret
	}
	return o.InvalidParams
}

// GetInvalidParamsOk returns a tuple with the InvalidParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetInvalidParamsOk() ([]InvalidParam, bool) {
	if o == nil || isNil(o.InvalidParams) {
    return nil, false
	}
	return o.InvalidParams, true
}

// HasInvalidParams returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasInvalidParams() bool {
	if o != nil && !isNil(o.InvalidParams) {
		return true
	}

	return false
}

// SetInvalidParams gets a reference to the given []InvalidParam and assigns it to the InvalidParams field.
func (o *ExtProblemDetails) SetInvalidParams(v []InvalidParam) {
	o.InvalidParams = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ExtProblemDetails) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetAccessTokenError returns the AccessTokenError field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetAccessTokenError() AccessTokenErr {
	if o == nil || isNil(o.AccessTokenError) {
		var ret AccessTokenErr
		return ret
	}
	return *o.AccessTokenError
}

// GetAccessTokenErrorOk returns a tuple with the AccessTokenError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetAccessTokenErrorOk() (*AccessTokenErr, bool) {
	if o == nil || isNil(o.AccessTokenError) {
    return nil, false
	}
	return o.AccessTokenError, true
}

// HasAccessTokenError returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasAccessTokenError() bool {
	if o != nil && !isNil(o.AccessTokenError) {
		return true
	}

	return false
}

// SetAccessTokenError gets a reference to the given AccessTokenErr and assigns it to the AccessTokenError field.
func (o *ExtProblemDetails) SetAccessTokenError(v AccessTokenErr) {
	o.AccessTokenError = &v
}

// GetAccessTokenRequest returns the AccessTokenRequest field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetAccessTokenRequest() AccessTokenReq {
	if o == nil || isNil(o.AccessTokenRequest) {
		var ret AccessTokenReq
		return ret
	}
	return *o.AccessTokenRequest
}

// GetAccessTokenRequestOk returns a tuple with the AccessTokenRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetAccessTokenRequestOk() (*AccessTokenReq, bool) {
	if o == nil || isNil(o.AccessTokenRequest) {
    return nil, false
	}
	return o.AccessTokenRequest, true
}

// HasAccessTokenRequest returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasAccessTokenRequest() bool {
	if o != nil && !isNil(o.AccessTokenRequest) {
		return true
	}

	return false
}

// SetAccessTokenRequest gets a reference to the given AccessTokenReq and assigns it to the AccessTokenRequest field.
func (o *ExtProblemDetails) SetAccessTokenRequest(v AccessTokenReq) {
	o.AccessTokenRequest = &v
}

// GetNrfId returns the NrfId field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetNrfId() string {
	if o == nil || isNil(o.NrfId) {
		var ret string
		return ret
	}
	return *o.NrfId
}

// GetNrfIdOk returns a tuple with the NrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetNrfIdOk() (*string, bool) {
	if o == nil || isNil(o.NrfId) {
    return nil, false
	}
	return o.NrfId, true
}

// HasNrfId returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasNrfId() bool {
	if o != nil && !isNil(o.NrfId) {
		return true
	}

	return false
}

// SetNrfId gets a reference to the given string and assigns it to the NrfId field.
func (o *ExtProblemDetails) SetNrfId(v string) {
	o.NrfId = &v
}

// GetRemoteError returns the RemoteError field value if set, zero value otherwise.
func (o *ExtProblemDetails) GetRemoteError() bool {
	if o == nil || isNil(o.RemoteError) {
		var ret bool
		return ret
	}
	return *o.RemoteError
}

// GetRemoteErrorOk returns a tuple with the RemoteError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtProblemDetails) GetRemoteErrorOk() (*bool, bool) {
	if o == nil || isNil(o.RemoteError) {
    return nil, false
	}
	return o.RemoteError, true
}

// HasRemoteError returns a boolean if a field has been set.
func (o *ExtProblemDetails) HasRemoteError() bool {
	if o != nil && !isNil(o.RemoteError) {
		return true
	}

	return false
}

// SetRemoteError gets a reference to the given bool and assigns it to the RemoteError field.
func (o *ExtProblemDetails) SetRemoteError(v bool) {
	o.RemoteError = &v
}

func (o ExtProblemDetails) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableExtProblemDetails struct {
	value *ExtProblemDetails
	isSet bool
}

func (v NullableExtProblemDetails) Get() *ExtProblemDetails {
	return v.value
}

func (v *NullableExtProblemDetails) Set(val *ExtProblemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableExtProblemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableExtProblemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtProblemDetails(val *ExtProblemDetails) *NullableExtProblemDetails {
	return &NullableExtProblemDetails{value: val, isSet: true}
}

func (v NullableExtProblemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtProblemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


