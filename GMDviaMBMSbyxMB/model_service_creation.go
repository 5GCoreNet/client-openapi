/*
GMDviaMBMSbyxMB

API for Group Message Delivery via MBMS by xMB   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package GMDviaMBMSbyxMB

import (
	"encoding/json"
)

// ServiceCreation Represents an individual xMB Service resource.
type ServiceCreation struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId *string `json:"externalGroupId,omitempty"`
	// Identifies the MBMS User Service supplied by the SCEF.
	UserServiceId *string `json:"userServiceId,omitempty"`
	// The service class that service belongs to supplied by the SCEF.
	ServiceClass *string `json:"serviceClass,omitempty"`
	// List of language of the service content supplied by the SCEF.
	ServiceLanguages []string `json:"serviceLanguages,omitempty"`
	// List of Service Names supplied by the SCEF.
	ServiceNames []string `json:"serviceNames,omitempty"`
	// When set to 'true', the Content Provider indicates that the service is a Receive Only Mode service. This parameter is supplied by the SCEF.
	ReceiveOnlyMode *bool `json:"receiveOnlyMode,omitempty"`
	ServiceAnnouncementMode *ServiceAnnouncementMode `json:"serviceAnnouncementMode,omitempty"`
}

// NewServiceCreation instantiates a new ServiceCreation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCreation() *ServiceCreation {
	this := ServiceCreation{}
	return &this
}

// NewServiceCreationWithDefaults instantiates a new ServiceCreation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCreationWithDefaults() *ServiceCreation {
	this := ServiceCreation{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ServiceCreation) GetSelf() string {
	if o == nil || isNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetSelfOk() (*string, bool) {
	if o == nil || isNil(o.Self) {
    return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ServiceCreation) HasSelf() bool {
	if o != nil && !isNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ServiceCreation) SetSelf(v string) {
	o.Self = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ServiceCreation) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
    return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ServiceCreation) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ServiceCreation) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetExternalGroupId returns the ExternalGroupId field value if set, zero value otherwise.
func (o *ServiceCreation) GetExternalGroupId() string {
	if o == nil || isNil(o.ExternalGroupId) {
		var ret string
		return ret
	}
	return *o.ExternalGroupId
}

// GetExternalGroupIdOk returns a tuple with the ExternalGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetExternalGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalGroupId) {
    return nil, false
	}
	return o.ExternalGroupId, true
}

// HasExternalGroupId returns a boolean if a field has been set.
func (o *ServiceCreation) HasExternalGroupId() bool {
	if o != nil && !isNil(o.ExternalGroupId) {
		return true
	}

	return false
}

// SetExternalGroupId gets a reference to the given string and assigns it to the ExternalGroupId field.
func (o *ServiceCreation) SetExternalGroupId(v string) {
	o.ExternalGroupId = &v
}

// GetUserServiceId returns the UserServiceId field value if set, zero value otherwise.
func (o *ServiceCreation) GetUserServiceId() string {
	if o == nil || isNil(o.UserServiceId) {
		var ret string
		return ret
	}
	return *o.UserServiceId
}

// GetUserServiceIdOk returns a tuple with the UserServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetUserServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.UserServiceId) {
    return nil, false
	}
	return o.UserServiceId, true
}

// HasUserServiceId returns a boolean if a field has been set.
func (o *ServiceCreation) HasUserServiceId() bool {
	if o != nil && !isNil(o.UserServiceId) {
		return true
	}

	return false
}

// SetUserServiceId gets a reference to the given string and assigns it to the UserServiceId field.
func (o *ServiceCreation) SetUserServiceId(v string) {
	o.UserServiceId = &v
}

// GetServiceClass returns the ServiceClass field value if set, zero value otherwise.
func (o *ServiceCreation) GetServiceClass() string {
	if o == nil || isNil(o.ServiceClass) {
		var ret string
		return ret
	}
	return *o.ServiceClass
}

// GetServiceClassOk returns a tuple with the ServiceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetServiceClassOk() (*string, bool) {
	if o == nil || isNil(o.ServiceClass) {
    return nil, false
	}
	return o.ServiceClass, true
}

// HasServiceClass returns a boolean if a field has been set.
func (o *ServiceCreation) HasServiceClass() bool {
	if o != nil && !isNil(o.ServiceClass) {
		return true
	}

	return false
}

// SetServiceClass gets a reference to the given string and assigns it to the ServiceClass field.
func (o *ServiceCreation) SetServiceClass(v string) {
	o.ServiceClass = &v
}

// GetServiceLanguages returns the ServiceLanguages field value if set, zero value otherwise.
func (o *ServiceCreation) GetServiceLanguages() []string {
	if o == nil || isNil(o.ServiceLanguages) {
		var ret []string
		return ret
	}
	return o.ServiceLanguages
}

// GetServiceLanguagesOk returns a tuple with the ServiceLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetServiceLanguagesOk() ([]string, bool) {
	if o == nil || isNil(o.ServiceLanguages) {
    return nil, false
	}
	return o.ServiceLanguages, true
}

// HasServiceLanguages returns a boolean if a field has been set.
func (o *ServiceCreation) HasServiceLanguages() bool {
	if o != nil && !isNil(o.ServiceLanguages) {
		return true
	}

	return false
}

// SetServiceLanguages gets a reference to the given []string and assigns it to the ServiceLanguages field.
func (o *ServiceCreation) SetServiceLanguages(v []string) {
	o.ServiceLanguages = v
}

// GetServiceNames returns the ServiceNames field value if set, zero value otherwise.
func (o *ServiceCreation) GetServiceNames() []string {
	if o == nil || isNil(o.ServiceNames) {
		var ret []string
		return ret
	}
	return o.ServiceNames
}

// GetServiceNamesOk returns a tuple with the ServiceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetServiceNamesOk() ([]string, bool) {
	if o == nil || isNil(o.ServiceNames) {
    return nil, false
	}
	return o.ServiceNames, true
}

// HasServiceNames returns a boolean if a field has been set.
func (o *ServiceCreation) HasServiceNames() bool {
	if o != nil && !isNil(o.ServiceNames) {
		return true
	}

	return false
}

// SetServiceNames gets a reference to the given []string and assigns it to the ServiceNames field.
func (o *ServiceCreation) SetServiceNames(v []string) {
	o.ServiceNames = v
}

// GetReceiveOnlyMode returns the ReceiveOnlyMode field value if set, zero value otherwise.
func (o *ServiceCreation) GetReceiveOnlyMode() bool {
	if o == nil || isNil(o.ReceiveOnlyMode) {
		var ret bool
		return ret
	}
	return *o.ReceiveOnlyMode
}

// GetReceiveOnlyModeOk returns a tuple with the ReceiveOnlyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetReceiveOnlyModeOk() (*bool, bool) {
	if o == nil || isNil(o.ReceiveOnlyMode) {
    return nil, false
	}
	return o.ReceiveOnlyMode, true
}

// HasReceiveOnlyMode returns a boolean if a field has been set.
func (o *ServiceCreation) HasReceiveOnlyMode() bool {
	if o != nil && !isNil(o.ReceiveOnlyMode) {
		return true
	}

	return false
}

// SetReceiveOnlyMode gets a reference to the given bool and assigns it to the ReceiveOnlyMode field.
func (o *ServiceCreation) SetReceiveOnlyMode(v bool) {
	o.ReceiveOnlyMode = &v
}

// GetServiceAnnouncementMode returns the ServiceAnnouncementMode field value if set, zero value otherwise.
func (o *ServiceCreation) GetServiceAnnouncementMode() ServiceAnnouncementMode {
	if o == nil || isNil(o.ServiceAnnouncementMode) {
		var ret ServiceAnnouncementMode
		return ret
	}
	return *o.ServiceAnnouncementMode
}

// GetServiceAnnouncementModeOk returns a tuple with the ServiceAnnouncementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreation) GetServiceAnnouncementModeOk() (*ServiceAnnouncementMode, bool) {
	if o == nil || isNil(o.ServiceAnnouncementMode) {
    return nil, false
	}
	return o.ServiceAnnouncementMode, true
}

// HasServiceAnnouncementMode returns a boolean if a field has been set.
func (o *ServiceCreation) HasServiceAnnouncementMode() bool {
	if o != nil && !isNil(o.ServiceAnnouncementMode) {
		return true
	}

	return false
}

// SetServiceAnnouncementMode gets a reference to the given ServiceAnnouncementMode and assigns it to the ServiceAnnouncementMode field.
func (o *ServiceCreation) SetServiceAnnouncementMode(v ServiceAnnouncementMode) {
	o.ServiceAnnouncementMode = &v
}

func (o ServiceCreation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !isNil(o.ExternalGroupId) {
		toSerialize["externalGroupId"] = o.ExternalGroupId
	}
	if !isNil(o.UserServiceId) {
		toSerialize["userServiceId"] = o.UserServiceId
	}
	if !isNil(o.ServiceClass) {
		toSerialize["serviceClass"] = o.ServiceClass
	}
	if !isNil(o.ServiceLanguages) {
		toSerialize["serviceLanguages"] = o.ServiceLanguages
	}
	if !isNil(o.ServiceNames) {
		toSerialize["serviceNames"] = o.ServiceNames
	}
	if !isNil(o.ReceiveOnlyMode) {
		toSerialize["receiveOnlyMode"] = o.ReceiveOnlyMode
	}
	if !isNil(o.ServiceAnnouncementMode) {
		toSerialize["serviceAnnouncementMode"] = o.ServiceAnnouncementMode
	}
	return json.Marshal(toSerialize)
}

type NullableServiceCreation struct {
	value *ServiceCreation
	isSet bool
}

func (v NullableServiceCreation) Get() *ServiceCreation {
	return v.value
}

func (v *NullableServiceCreation) Set(val *ServiceCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceCreation(val *ServiceCreation) *NullableServiceCreation {
	return &NullableServiceCreation{value: val, isSet: true}
}

func (v NullableServiceCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

