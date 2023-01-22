# FeasibilityCheckAndReservationJobSingleAllOfAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**NullableOneOfSliceProfileServiceProfile**](oneOf&lt;SliceProfile,ServiceProfile&gt;.md) |  | [optional] 
**ResourceReservation** | Pointer to **bool** | An attribute represents MnS consumer&#39;s requirements for resource reservation. | [optional] 
**RecommendationRequest** | Pointer to **bool** | An attribute represents MnS consumer&#39;s request for recommended network slice related requirements. | [optional] 
**RequestedReservationExpiration** | Pointer to **string** | An attribute which specifes MnS consuner&#39;s requirements for the validity period of the resource reservation. | [optional] 
**ProcessMonitor** | Pointer to [**ProcessMonitor**](ProcessMonitor.md) |  | [optional] 
**FeasibilityResult** | Pointer to [**FeasibilityResult**](FeasibilityResult.md) |  | [optional] 
**InFeasibleReason** | Pointer to **string** | An attribute that specifies the additional reason information if the feasibility check result is infeasible.The detailed ENUM value is FFS.  | [optional] 
**ResourceReservationStatus** | Pointer to [**ResourceReservationStatus**](ResourceReservationStatus.md) |  | [optional] 
**ReservationFailureReason** | Pointer to **string** | An attribute that specifies the additional reason information if the reservation is failed.  | [optional] 
**ReservationExpiration** | Pointer to **string** | An attribute which specifes the actual validity period of the resource reservation. | [optional] 
**RecommendedRequirements** | Pointer to **string** | An attribute that specifies the recommended network slicing related requirements (i.e. ServiceProfile and SliceProfile information) which can be supported by the MnS producer. | [optional] 

## Methods

### NewFeasibilityCheckAndReservationJobSingleAllOfAttributes

`func NewFeasibilityCheckAndReservationJobSingleAllOfAttributes() *FeasibilityCheckAndReservationJobSingleAllOfAttributes`

NewFeasibilityCheckAndReservationJobSingleAllOfAttributes instantiates a new FeasibilityCheckAndReservationJobSingleAllOfAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeasibilityCheckAndReservationJobSingleAllOfAttributesWithDefaults

`func NewFeasibilityCheckAndReservationJobSingleAllOfAttributesWithDefaults() *FeasibilityCheckAndReservationJobSingleAllOfAttributes`

NewFeasibilityCheckAndReservationJobSingleAllOfAttributesWithDefaults instantiates a new FeasibilityCheckAndReservationJobSingleAllOfAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProfile() OneOfSliceProfileServiceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProfileOk() (*OneOfSliceProfileServiceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetProfile(v OneOfSliceProfileServiceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### SetProfileNil

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetProfileNil(b bool)`

 SetProfileNil sets the value for Profile to be an explicit nil

### UnsetProfile
`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) UnsetProfile()`

UnsetProfile ensures that no value is present for Profile, not even an explicit nil
### GetResourceReservation

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservation() bool`

GetResourceReservation returns the ResourceReservation field if non-nil, zero value otherwise.

### GetResourceReservationOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservationOk() (*bool, bool)`

GetResourceReservationOk returns a tuple with the ResourceReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceReservation

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetResourceReservation(v bool)`

SetResourceReservation sets ResourceReservation field to given value.

### HasResourceReservation

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasResourceReservation() bool`

HasResourceReservation returns a boolean if a field has been set.

### GetRecommendationRequest

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendationRequest() bool`

GetRecommendationRequest returns the RecommendationRequest field if non-nil, zero value otherwise.

### GetRecommendationRequestOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendationRequestOk() (*bool, bool)`

GetRecommendationRequestOk returns a tuple with the RecommendationRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendationRequest

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetRecommendationRequest(v bool)`

SetRecommendationRequest sets RecommendationRequest field to given value.

### HasRecommendationRequest

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasRecommendationRequest() bool`

HasRecommendationRequest returns a boolean if a field has been set.

### GetRequestedReservationExpiration

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRequestedReservationExpiration() string`

GetRequestedReservationExpiration returns the RequestedReservationExpiration field if non-nil, zero value otherwise.

### GetRequestedReservationExpirationOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRequestedReservationExpirationOk() (*string, bool)`

GetRequestedReservationExpirationOk returns a tuple with the RequestedReservationExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReservationExpiration

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetRequestedReservationExpiration(v string)`

SetRequestedReservationExpiration sets RequestedReservationExpiration field to given value.

### HasRequestedReservationExpiration

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasRequestedReservationExpiration() bool`

HasRequestedReservationExpiration returns a boolean if a field has been set.

### GetProcessMonitor

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProcessMonitor() ProcessMonitor`

GetProcessMonitor returns the ProcessMonitor field if non-nil, zero value otherwise.

### GetProcessMonitorOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetProcessMonitorOk() (*ProcessMonitor, bool)`

GetProcessMonitorOk returns a tuple with the ProcessMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessMonitor

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetProcessMonitor(v ProcessMonitor)`

SetProcessMonitor sets ProcessMonitor field to given value.

### HasProcessMonitor

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasProcessMonitor() bool`

HasProcessMonitor returns a boolean if a field has been set.

### GetFeasibilityResult

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetFeasibilityResult() FeasibilityResult`

GetFeasibilityResult returns the FeasibilityResult field if non-nil, zero value otherwise.

### GetFeasibilityResultOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetFeasibilityResultOk() (*FeasibilityResult, bool)`

GetFeasibilityResultOk returns a tuple with the FeasibilityResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeasibilityResult

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetFeasibilityResult(v FeasibilityResult)`

SetFeasibilityResult sets FeasibilityResult field to given value.

### HasFeasibilityResult

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasFeasibilityResult() bool`

HasFeasibilityResult returns a boolean if a field has been set.

### GetInFeasibleReason

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetInFeasibleReason() string`

GetInFeasibleReason returns the InFeasibleReason field if non-nil, zero value otherwise.

### GetInFeasibleReasonOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetInFeasibleReasonOk() (*string, bool)`

GetInFeasibleReasonOk returns a tuple with the InFeasibleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFeasibleReason

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetInFeasibleReason(v string)`

SetInFeasibleReason sets InFeasibleReason field to given value.

### HasInFeasibleReason

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasInFeasibleReason() bool`

HasInFeasibleReason returns a boolean if a field has been set.

### GetResourceReservationStatus

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservationStatus() ResourceReservationStatus`

GetResourceReservationStatus returns the ResourceReservationStatus field if non-nil, zero value otherwise.

### GetResourceReservationStatusOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetResourceReservationStatusOk() (*ResourceReservationStatus, bool)`

GetResourceReservationStatusOk returns a tuple with the ResourceReservationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceReservationStatus

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetResourceReservationStatus(v ResourceReservationStatus)`

SetResourceReservationStatus sets ResourceReservationStatus field to given value.

### HasResourceReservationStatus

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasResourceReservationStatus() bool`

HasResourceReservationStatus returns a boolean if a field has been set.

### GetReservationFailureReason

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationFailureReason() string`

GetReservationFailureReason returns the ReservationFailureReason field if non-nil, zero value otherwise.

### GetReservationFailureReasonOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationFailureReasonOk() (*string, bool)`

GetReservationFailureReasonOk returns a tuple with the ReservationFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationFailureReason

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetReservationFailureReason(v string)`

SetReservationFailureReason sets ReservationFailureReason field to given value.

### HasReservationFailureReason

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasReservationFailureReason() bool`

HasReservationFailureReason returns a boolean if a field has been set.

### GetReservationExpiration

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationExpiration() string`

GetReservationExpiration returns the ReservationExpiration field if non-nil, zero value otherwise.

### GetReservationExpirationOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetReservationExpirationOk() (*string, bool)`

GetReservationExpirationOk returns a tuple with the ReservationExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationExpiration

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetReservationExpiration(v string)`

SetReservationExpiration sets ReservationExpiration field to given value.

### HasReservationExpiration

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasReservationExpiration() bool`

HasReservationExpiration returns a boolean if a field has been set.

### GetRecommendedRequirements

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendedRequirements() string`

GetRecommendedRequirements returns the RecommendedRequirements field if non-nil, zero value otherwise.

### GetRecommendedRequirementsOk

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) GetRecommendedRequirementsOk() (*string, bool)`

GetRecommendedRequirementsOk returns a tuple with the RecommendedRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedRequirements

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) SetRecommendedRequirements(v string)`

SetRecommendedRequirements sets RecommendedRequirements field to given value.

### HasRecommendedRequirements

`func (o *FeasibilityCheckAndReservationJobSingleAllOfAttributes) HasRecommendedRequirements() bool`

HasRecommendedRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


