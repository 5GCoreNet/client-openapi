# OnboardingNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | **bool** | Set to \&quot;true\&quot; indicate successful on-boarding. Otherwise set to \&quot;false\&quot; | 
**ResourceLocation** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**ApiInvokerEnrolmentDetails** | Pointer to [**APIInvokerEnrolmentDetails**](APIInvokerEnrolmentDetails.md) |  | [optional] 
**ApiList** | Pointer to [**[]ServiceAPIDescription**](ServiceAPIDescription.md) | The list of service APIs that the API Invoker is allowed to invoke | [optional] 

## Methods

### NewOnboardingNotification

`func NewOnboardingNotification(result bool, ) *OnboardingNotification`

NewOnboardingNotification instantiates a new OnboardingNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingNotificationWithDefaults

`func NewOnboardingNotificationWithDefaults() *OnboardingNotification`

NewOnboardingNotificationWithDefaults instantiates a new OnboardingNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *OnboardingNotification) GetResult() bool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OnboardingNotification) GetResultOk() (*bool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OnboardingNotification) SetResult(v bool)`

SetResult sets Result field to given value.


### GetResourceLocation

`func (o *OnboardingNotification) GetResourceLocation() string`

GetResourceLocation returns the ResourceLocation field if non-nil, zero value otherwise.

### GetResourceLocationOk

`func (o *OnboardingNotification) GetResourceLocationOk() (*string, bool)`

GetResourceLocationOk returns a tuple with the ResourceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLocation

`func (o *OnboardingNotification) SetResourceLocation(v string)`

SetResourceLocation sets ResourceLocation field to given value.

### HasResourceLocation

`func (o *OnboardingNotification) HasResourceLocation() bool`

HasResourceLocation returns a boolean if a field has been set.

### GetApiInvokerEnrolmentDetails

`func (o *OnboardingNotification) GetApiInvokerEnrolmentDetails() APIInvokerEnrolmentDetails`

GetApiInvokerEnrolmentDetails returns the ApiInvokerEnrolmentDetails field if non-nil, zero value otherwise.

### GetApiInvokerEnrolmentDetailsOk

`func (o *OnboardingNotification) GetApiInvokerEnrolmentDetailsOk() (*APIInvokerEnrolmentDetails, bool)`

GetApiInvokerEnrolmentDetailsOk returns a tuple with the ApiInvokerEnrolmentDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerEnrolmentDetails

`func (o *OnboardingNotification) SetApiInvokerEnrolmentDetails(v APIInvokerEnrolmentDetails)`

SetApiInvokerEnrolmentDetails sets ApiInvokerEnrolmentDetails field to given value.

### HasApiInvokerEnrolmentDetails

`func (o *OnboardingNotification) HasApiInvokerEnrolmentDetails() bool`

HasApiInvokerEnrolmentDetails returns a boolean if a field has been set.

### GetApiList

`func (o *OnboardingNotification) GetApiList() []ServiceAPIDescription`

GetApiList returns the ApiList field if non-nil, zero value otherwise.

### GetApiListOk

`func (o *OnboardingNotification) GetApiListOk() (*[]ServiceAPIDescription, bool)`

GetApiListOk returns a tuple with the ApiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiList

`func (o *OnboardingNotification) SetApiList(v []ServiceAPIDescription)`

SetApiList sets ApiList field to given value.

### HasApiList

`func (o *OnboardingNotification) HasApiList() bool`

HasApiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


