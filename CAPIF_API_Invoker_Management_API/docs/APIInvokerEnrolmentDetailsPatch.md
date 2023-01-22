# APIInvokerEnrolmentDetailsPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnboardingInformation** | Pointer to [**OnboardingInformation**](OnboardingInformation.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**ApiList** | Pointer to [**[]ServiceAPIDescription**](ServiceAPIDescription.md) | The list of service APIs that the API Invoker is allowed to invoke | [optional] 
**ApiInvokerInformation** | Pointer to **string** | Generic information related to the API invoker such as details of the device or the application.  | [optional] 

## Methods

### NewAPIInvokerEnrolmentDetailsPatch

`func NewAPIInvokerEnrolmentDetailsPatch() *APIInvokerEnrolmentDetailsPatch`

NewAPIInvokerEnrolmentDetailsPatch instantiates a new APIInvokerEnrolmentDetailsPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIInvokerEnrolmentDetailsPatchWithDefaults

`func NewAPIInvokerEnrolmentDetailsPatchWithDefaults() *APIInvokerEnrolmentDetailsPatch`

NewAPIInvokerEnrolmentDetailsPatchWithDefaults instantiates a new APIInvokerEnrolmentDetailsPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnboardingInformation

`func (o *APIInvokerEnrolmentDetailsPatch) GetOnboardingInformation() OnboardingInformation`

GetOnboardingInformation returns the OnboardingInformation field if non-nil, zero value otherwise.

### GetOnboardingInformationOk

`func (o *APIInvokerEnrolmentDetailsPatch) GetOnboardingInformationOk() (*OnboardingInformation, bool)`

GetOnboardingInformationOk returns a tuple with the OnboardingInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingInformation

`func (o *APIInvokerEnrolmentDetailsPatch) SetOnboardingInformation(v OnboardingInformation)`

SetOnboardingInformation sets OnboardingInformation field to given value.

### HasOnboardingInformation

`func (o *APIInvokerEnrolmentDetailsPatch) HasOnboardingInformation() bool`

HasOnboardingInformation returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *APIInvokerEnrolmentDetailsPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *APIInvokerEnrolmentDetailsPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *APIInvokerEnrolmentDetailsPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *APIInvokerEnrolmentDetailsPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetApiList

`func (o *APIInvokerEnrolmentDetailsPatch) GetApiList() []ServiceAPIDescription`

GetApiList returns the ApiList field if non-nil, zero value otherwise.

### GetApiListOk

`func (o *APIInvokerEnrolmentDetailsPatch) GetApiListOk() (*[]ServiceAPIDescription, bool)`

GetApiListOk returns a tuple with the ApiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiList

`func (o *APIInvokerEnrolmentDetailsPatch) SetApiList(v []ServiceAPIDescription)`

SetApiList sets ApiList field to given value.

### HasApiList

`func (o *APIInvokerEnrolmentDetailsPatch) HasApiList() bool`

HasApiList returns a boolean if a field has been set.

### GetApiInvokerInformation

`func (o *APIInvokerEnrolmentDetailsPatch) GetApiInvokerInformation() string`

GetApiInvokerInformation returns the ApiInvokerInformation field if non-nil, zero value otherwise.

### GetApiInvokerInformationOk

`func (o *APIInvokerEnrolmentDetailsPatch) GetApiInvokerInformationOk() (*string, bool)`

GetApiInvokerInformationOk returns a tuple with the ApiInvokerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerInformation

`func (o *APIInvokerEnrolmentDetailsPatch) SetApiInvokerInformation(v string)`

SetApiInvokerInformation sets ApiInvokerInformation field to given value.

### HasApiInvokerInformation

`func (o *APIInvokerEnrolmentDetailsPatch) HasApiInvokerInformation() bool`

HasApiInvokerInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


