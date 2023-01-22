/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ProvMnS

import (
	"encoding/json"
)

// FeasibilityCheckAndReservationJobSingleAllOfAttributes struct for FeasibilityCheckAndReservationJobSingleAllOfAttributes
type FeasibilityCheckAndReservationJobSingleAllOfAttributes struct {
	Profile NullableOneOfSliceProfileServiceProfile `json:"profile,omitempty"`
	// An attribute represents MnS consumer's requirements for resource reservation.
	ResourceReservation *bool `json:"resourceReservation,omitempty"`
	// An attribute represents MnS consumer's request for recommended network slice related requirements.
	RecommendationRequest *bool `json:"recommendationRequest,omitempty"`
	// An attribute which specifes MnS consuner's requirements for the validity period of the resource reservation.
	RequestedReservationExpiration *string `json:"requestedReservationExpiration,omitempty"`
	ProcessMonitor *ProcessMonitor `json:"processMonitor,omitempty"`
	FeasibilityResult *FeasibilityResult `json:"feasibilityResult,omitempty"`
	// An attribute that specifies the additional reason information if the feasibility check result is infeasible.The detailed ENUM value is FFS. 
	InFeasibleReason *string `json:"inFeasibleReason,omitempty"`
	ResourceReservationStatus *ResourceReservationStatus `json:"resourceReservationStatus,omitempty"`
	// An attribute that specifies the additional reason information if the reservation is failed. 
	ReservationFailureReason *string `json:"reservationFailureReason,omitempty"`
	// An attribute which specifes the actual validity period of the resource reservation.
	ReservationExpiration *string `json:"reservationExpiration,omitempty"`
	// An attribute that specifies the recommended network slicing related requirements (i.e. ServiceProfile and SliceProfile information) which can be supported by the MnS producer.
	RecommendedRequirements *string `json:"recommendedRequirements,omitempty"`
}

// NewFeasibilityCheckAndReservationJobSingleAllOfAttributes instantiates a new FeasibilityCheckAndReservationJobSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeasibilityCheckAndReservationJobSingleAllOfAttributes() *FeasibilityCheckAndReservationJobSingleAllOfAttributes {
	this := FeasibilityCheckAndReservationJobSingleAllOfAttributes{}
	return &this
}

// NewFeasibilityCheckAndReservationJobSingleAllOfAttributesWithDefaults instantiates a new FeasibilityCheckAndReservationJobSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeasibilityCheckAndReservationJobSingleAllOfAttributesWithDefaults() *FeasibilityCheckAndReservationJobSingleAllOfAttributes {
	this := FeasibilityCheckAndReservationJobSingleAllOfAttributes{}
	return &this
}

// GetProfile returns the Profile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProfile() OneOfSliceProfileServiceProfile {
	if o == nil || isNil(o.Profile.Get()) {
		var ret OneOfSliceProfileServiceProfile
		return ret
	}
	return *o.Profile.Get()
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProfileOk() (*OneOfSliceProfileServiceProfile, bool) {
	if o == nil {
    return nil, false
	}
	return o.Profile.Get(), o.Profile.IsSet()
}

// HasProfile returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasProfile() bool {
	if o != nil && o.Profile.IsSet() {
		return true
	}

	return false
}

// SetProfile gets a reference to the given NullableOneOfSliceProfileServiceProfile and assigns it to the Profile field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetProfile(v OneOfSliceProfileServiceProfile) {
	o.Profile.Set(&v)
}
// SetProfileNil sets the value for Profile to be an explicit nil
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetProfileNil() {
	o.Profile.Set(nil)
}

// UnsetProfile ensures that no value is present for Profile, not even an explicit nil
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) UnsetProfile() {
	o.Profile.Unset()
}

// GetResourceReservation returns the ResourceReservation field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservation() bool {
	if o == nil || isNil(o.ResourceReservation) {
		var ret bool
		return ret
	}
	return *o.ResourceReservation
}

// GetResourceReservationOk returns a tuple with the ResourceReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservationOk() (*bool, bool) {
	if o == nil || isNil(o.ResourceReservation) {
    return nil, false
	}
	return o.ResourceReservation, true
}

// HasResourceReservation returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasResourceReservation() bool {
	if o != nil && !isNil(o.ResourceReservation) {
		return true
	}

	return false
}

// SetResourceReservation gets a reference to the given bool and assigns it to the ResourceReservation field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetResourceReservation(v bool) {
	o.ResourceReservation = &v
}

// GetRecommendationRequest returns the RecommendationRequest field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendationRequest() bool {
	if o == nil || isNil(o.RecommendationRequest) {
		var ret bool
		return ret
	}
	return *o.RecommendationRequest
}

// GetRecommendationRequestOk returns a tuple with the RecommendationRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendationRequestOk() (*bool, bool) {
	if o == nil || isNil(o.RecommendationRequest) {
    return nil, false
	}
	return o.RecommendationRequest, true
}

// HasRecommendationRequest returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasRecommendationRequest() bool {
	if o != nil && !isNil(o.RecommendationRequest) {
		return true
	}

	return false
}

// SetRecommendationRequest gets a reference to the given bool and assigns it to the RecommendationRequest field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetRecommendationRequest(v bool) {
	o.RecommendationRequest = &v
}

// GetRequestedReservationExpiration returns the RequestedReservationExpiration field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRequestedReservationExpiration() string {
	if o == nil || isNil(o.RequestedReservationExpiration) {
		var ret string
		return ret
	}
	return *o.RequestedReservationExpiration
}

// GetRequestedReservationExpirationOk returns a tuple with the RequestedReservationExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRequestedReservationExpirationOk() (*string, bool) {
	if o == nil || isNil(o.RequestedReservationExpiration) {
    return nil, false
	}
	return o.RequestedReservationExpiration, true
}

// HasRequestedReservationExpiration returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasRequestedReservationExpiration() bool {
	if o != nil && !isNil(o.RequestedReservationExpiration) {
		return true
	}

	return false
}

// SetRequestedReservationExpiration gets a reference to the given string and assigns it to the RequestedReservationExpiration field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetRequestedReservationExpiration(v string) {
	o.RequestedReservationExpiration = &v
}

// GetProcessMonitor returns the ProcessMonitor field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProcessMonitor() ProcessMonitor {
	if o == nil || isNil(o.ProcessMonitor) {
		var ret ProcessMonitor
		return ret
	}
	return *o.ProcessMonitor
}

// GetProcessMonitorOk returns a tuple with the ProcessMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProcessMonitorOk() (*ProcessMonitor, bool) {
	if o == nil || isNil(o.ProcessMonitor) {
    return nil, false
	}
	return o.ProcessMonitor, true
}

// HasProcessMonitor returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasProcessMonitor() bool {
	if o != nil && !isNil(o.ProcessMonitor) {
		return true
	}

	return false
}

// SetProcessMonitor gets a reference to the given ProcessMonitor and assigns it to the ProcessMonitor field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetProcessMonitor(v ProcessMonitor) {
	o.ProcessMonitor = &v
}

// GetFeasibilityResult returns the FeasibilityResult field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetFeasibilityResult() FeasibilityResult {
	if o == nil || isNil(o.FeasibilityResult) {
		var ret FeasibilityResult
		return ret
	}
	return *o.FeasibilityResult
}

// GetFeasibilityResultOk returns a tuple with the FeasibilityResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetFeasibilityResultOk() (*FeasibilityResult, bool) {
	if o == nil || isNil(o.FeasibilityResult) {
    return nil, false
	}
	return o.FeasibilityResult, true
}

// HasFeasibilityResult returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasFeasibilityResult() bool {
	if o != nil && !isNil(o.FeasibilityResult) {
		return true
	}

	return false
}

// SetFeasibilityResult gets a reference to the given FeasibilityResult and assigns it to the FeasibilityResult field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetFeasibilityResult(v FeasibilityResult) {
	o.FeasibilityResult = &v
}

// GetInFeasibleReason returns the InFeasibleReason field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetInFeasibleReason() string {
	if o == nil || isNil(o.InFeasibleReason) {
		var ret string
		return ret
	}
	return *o.InFeasibleReason
}

// GetInFeasibleReasonOk returns a tuple with the InFeasibleReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetInFeasibleReasonOk() (*string, bool) {
	if o == nil || isNil(o.InFeasibleReason) {
    return nil, false
	}
	return o.InFeasibleReason, true
}

// HasInFeasibleReason returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasInFeasibleReason() bool {
	if o != nil && !isNil(o.InFeasibleReason) {
		return true
	}

	return false
}

// SetInFeasibleReason gets a reference to the given string and assigns it to the InFeasibleReason field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetInFeasibleReason(v string) {
	o.InFeasibleReason = &v
}

// GetResourceReservationStatus returns the ResourceReservationStatus field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservationStatus() ResourceReservationStatus {
	if o == nil || isNil(o.ResourceReservationStatus) {
		var ret ResourceReservationStatus
		return ret
	}
	return *o.ResourceReservationStatus
}

// GetResourceReservationStatusOk returns a tuple with the ResourceReservationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservationStatusOk() (*ResourceReservationStatus, bool) {
	if o == nil || isNil(o.ResourceReservationStatus) {
    return nil, false
	}
	return o.ResourceReservationStatus, true
}

// HasResourceReservationStatus returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasResourceReservationStatus() bool {
	if o != nil && !isNil(o.ResourceReservationStatus) {
		return true
	}

	return false
}

// SetResourceReservationStatus gets a reference to the given ResourceReservationStatus and assigns it to the ResourceReservationStatus field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetResourceReservationStatus(v ResourceReservationStatus) {
	o.ResourceReservationStatus = &v
}

// GetReservationFailureReason returns the ReservationFailureReason field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationFailureReason() string {
	if o == nil || isNil(o.ReservationFailureReason) {
		var ret string
		return ret
	}
	return *o.ReservationFailureReason
}

// GetReservationFailureReasonOk returns a tuple with the ReservationFailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationFailureReasonOk() (*string, bool) {
	if o == nil || isNil(o.ReservationFailureReason) {
    return nil, false
	}
	return o.ReservationFailureReason, true
}

// HasReservationFailureReason returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasReservationFailureReason() bool {
	if o != nil && !isNil(o.ReservationFailureReason) {
		return true
	}

	return false
}

// SetReservationFailureReason gets a reference to the given string and assigns it to the ReservationFailureReason field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetReservationFailureReason(v string) {
	o.ReservationFailureReason = &v
}

// GetReservationExpiration returns the ReservationExpiration field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationExpiration() string {
	if o == nil || isNil(o.ReservationExpiration) {
		var ret string
		return ret
	}
	return *o.ReservationExpiration
}

// GetReservationExpirationOk returns a tuple with the ReservationExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationExpirationOk() (*string, bool) {
	if o == nil || isNil(o.ReservationExpiration) {
    return nil, false
	}
	return o.ReservationExpiration, true
}

// HasReservationExpiration returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasReservationExpiration() bool {
	if o != nil && !isNil(o.ReservationExpiration) {
		return true
	}

	return false
}

// SetReservationExpiration gets a reference to the given string and assigns it to the ReservationExpiration field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetReservationExpiration(v string) {
	o.ReservationExpiration = &v
}

// GetRecommendedRequirements returns the RecommendedRequirements field value if set, zero value otherwise.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendedRequirements() string {
	if o == nil || isNil(o.RecommendedRequirements) {
		var ret string
		return ret
	}
	return *o.RecommendedRequirements
}

// GetRecommendedRequirementsOk returns a tuple with the RecommendedRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendedRequirementsOk() (*string, bool) {
	if o == nil || isNil(o.RecommendedRequirements) {
    return nil, false
	}
	return o.RecommendedRequirements, true
}

// HasRecommendedRequirements returns a boolean if a field has been set.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasRecommendedRequirements() bool {
	if o != nil && !isNil(o.RecommendedRequirements) {
		return true
	}

	return false
}

// SetRecommendedRequirements gets a reference to the given string and assigns it to the RecommendedRequirements field.
func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetRecommendedRequirements(v string) {
	o.RecommendedRequirements = &v
}

func (o FeasibilityCheckAndReservationJobSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Profile.IsSet() {
		toSerialize["profile"] = o.Profile.Get()
	}
	if !isNil(o.ResourceReservation) {
		toSerialize["resourceReservation"] = o.ResourceReservation
	}
	if !isNil(o.RecommendationRequest) {
		toSerialize["recommendationRequest"] = o.RecommendationRequest
	}
	if !isNil(o.RequestedReservationExpiration) {
		toSerialize["requestedReservationExpiration"] = o.RequestedReservationExpiration
	}
	if !isNil(o.ProcessMonitor) {
		toSerialize["processMonitor"] = o.ProcessMonitor
	}
	if !isNil(o.FeasibilityResult) {
		toSerialize["feasibilityResult"] = o.FeasibilityResult
	}
	if !isNil(o.InFeasibleReason) {
		toSerialize["inFeasibleReason"] = o.InFeasibleReason
	}
	if !isNil(o.ResourceReservationStatus) {
		toSerialize["resourceReservationStatus"] = o.ResourceReservationStatus
	}
	if !isNil(o.ReservationFailureReason) {
		toSerialize["reservationFailureReason"] = o.ReservationFailureReason
	}
	if !isNil(o.ReservationExpiration) {
		toSerialize["reservationExpiration"] = o.ReservationExpiration
	}
	if !isNil(o.RecommendedRequirements) {
		toSerialize["recommendedRequirements"] = o.RecommendedRequirements
	}
	return json.Marshal(toSerialize)
}

type NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes struct {
	value *FeasibilityCheckAndReservationJobSingleAllOfAttributes
	isSet bool
}

func (v NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes) Get() *FeasibilityCheckAndReservationJobSingleAllOfAttributes {
	return v.value
}

func (v *NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes) Set(val *FeasibilityCheckAndReservationJobSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeasibilityCheckAndReservationJobSingleAllOfAttributes(val *FeasibilityCheckAndReservationJobSingleAllOfAttributes) *NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes {
	return &NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeasibilityCheckAndReservationJobSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


