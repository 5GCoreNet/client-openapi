/*
CAPIF_API_Invoker_Management_API

API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Invoker_Management_API

import (
	"encoding/json"
)

// APIInvokerEnrolmentDetailsPatch Represents an API Invoker's enrolment details to be updated.
type APIInvokerEnrolmentDetailsPatch struct {
	OnboardingInformation *OnboardingInformation `json:"onboardingInformation,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// The list of service APIs that the API Invoker is allowed to invoke
	ApiList []ServiceAPIDescription `json:"apiList,omitempty"`
	// Generic information related to the API invoker such as details of the device or the application. 
	ApiInvokerInformation *string `json:"apiInvokerInformation,omitempty"`
}

// NewAPIInvokerEnrolmentDetailsPatch instantiates a new APIInvokerEnrolmentDetailsPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIInvokerEnrolmentDetailsPatch() *APIInvokerEnrolmentDetailsPatch {
	this := APIInvokerEnrolmentDetailsPatch{}
	return &this
}

// NewAPIInvokerEnrolmentDetailsPatchWithDefaults instantiates a new APIInvokerEnrolmentDetailsPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIInvokerEnrolmentDetailsPatchWithDefaults() *APIInvokerEnrolmentDetailsPatch {
	this := APIInvokerEnrolmentDetailsPatch{}
	return &this
}

// GetOnboardingInformation returns the OnboardingInformation field value if set, zero value otherwise.
func (o *APIInvokerEnrolmentDetailsPatch) GetOnboardingInformation() OnboardingInformation {
	if o == nil || isNil(o.OnboardingInformation) {
		var ret OnboardingInformation
		return ret
	}
	return *o.OnboardingInformation
}

// GetOnboardingInformationOk returns a tuple with the OnboardingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIInvokerEnrolmentDetailsPatch) GetOnboardingInformationOk() (*OnboardingInformation, bool) {
	if o == nil || isNil(o.OnboardingInformation) {
    return nil, false
	}
	return o.OnboardingInformation, true
}

// HasOnboardingInformation returns a boolean if a field has been set.
func (o *APIInvokerEnrolmentDetailsPatch) HasOnboardingInformation() bool {
	if o != nil && !isNil(o.OnboardingInformation) {
		return true
	}

	return false
}

// SetOnboardingInformation gets a reference to the given OnboardingInformation and assigns it to the OnboardingInformation field.
func (o *APIInvokerEnrolmentDetailsPatch) SetOnboardingInformation(v OnboardingInformation) {
	o.OnboardingInformation = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *APIInvokerEnrolmentDetailsPatch) GetNotificationDestination() string {
	if o == nil || isNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIInvokerEnrolmentDetailsPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || isNil(o.NotificationDestination) {
    return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *APIInvokerEnrolmentDetailsPatch) HasNotificationDestination() bool {
	if o != nil && !isNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *APIInvokerEnrolmentDetailsPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetApiList returns the ApiList field value if set, zero value otherwise.
func (o *APIInvokerEnrolmentDetailsPatch) GetApiList() []ServiceAPIDescription {
	if o == nil || isNil(o.ApiList) {
		var ret []ServiceAPIDescription
		return ret
	}
	return o.ApiList
}

// GetApiListOk returns a tuple with the ApiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIInvokerEnrolmentDetailsPatch) GetApiListOk() ([]ServiceAPIDescription, bool) {
	if o == nil || isNil(o.ApiList) {
    return nil, false
	}
	return o.ApiList, true
}

// HasApiList returns a boolean if a field has been set.
func (o *APIInvokerEnrolmentDetailsPatch) HasApiList() bool {
	if o != nil && !isNil(o.ApiList) {
		return true
	}

	return false
}

// SetApiList gets a reference to the given []ServiceAPIDescription and assigns it to the ApiList field.
func (o *APIInvokerEnrolmentDetailsPatch) SetApiList(v []ServiceAPIDescription) {
	o.ApiList = v
}

// GetApiInvokerInformation returns the ApiInvokerInformation field value if set, zero value otherwise.
func (o *APIInvokerEnrolmentDetailsPatch) GetApiInvokerInformation() string {
	if o == nil || isNil(o.ApiInvokerInformation) {
		var ret string
		return ret
	}
	return *o.ApiInvokerInformation
}

// GetApiInvokerInformationOk returns a tuple with the ApiInvokerInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIInvokerEnrolmentDetailsPatch) GetApiInvokerInformationOk() (*string, bool) {
	if o == nil || isNil(o.ApiInvokerInformation) {
    return nil, false
	}
	return o.ApiInvokerInformation, true
}

// HasApiInvokerInformation returns a boolean if a field has been set.
func (o *APIInvokerEnrolmentDetailsPatch) HasApiInvokerInformation() bool {
	if o != nil && !isNil(o.ApiInvokerInformation) {
		return true
	}

	return false
}

// SetApiInvokerInformation gets a reference to the given string and assigns it to the ApiInvokerInformation field.
func (o *APIInvokerEnrolmentDetailsPatch) SetApiInvokerInformation(v string) {
	o.ApiInvokerInformation = &v
}

func (o APIInvokerEnrolmentDetailsPatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OnboardingInformation) {
		toSerialize["onboardingInformation"] = o.OnboardingInformation
	}
	if !isNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !isNil(o.ApiList) {
		toSerialize["apiList"] = o.ApiList
	}
	if !isNil(o.ApiInvokerInformation) {
		toSerialize["apiInvokerInformation"] = o.ApiInvokerInformation
	}
	return json.Marshal(toSerialize)
}

type NullableAPIInvokerEnrolmentDetailsPatch struct {
	value *APIInvokerEnrolmentDetailsPatch
	isSet bool
}

func (v NullableAPIInvokerEnrolmentDetailsPatch) Get() *APIInvokerEnrolmentDetailsPatch {
	return v.value
}

func (v *NullableAPIInvokerEnrolmentDetailsPatch) Set(val *APIInvokerEnrolmentDetailsPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIInvokerEnrolmentDetailsPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIInvokerEnrolmentDetailsPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIInvokerEnrolmentDetailsPatch(val *APIInvokerEnrolmentDetailsPatch) *NullableAPIInvokerEnrolmentDetailsPatch {
	return &NullableAPIInvokerEnrolmentDetailsPatch{value: val, isSet: true}
}

func (v NullableAPIInvokerEnrolmentDetailsPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIInvokerEnrolmentDetailsPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


