# LocationSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**RevocationNotifUri** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**LocGran** | Pointer to [**Accuracy**](Accuracy.md) |  | [optional] 
**LocQos** | Pointer to [**LocationQoS**](LocationQoS.md) |  | [optional] 

## Methods

### NewLocationSubscriptionPatch

`func NewLocationSubscriptionPatch() *LocationSubscriptionPatch`

NewLocationSubscriptionPatch instantiates a new LocationSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationSubscriptionPatchWithDefaults

`func NewLocationSubscriptionPatchWithDefaults() *LocationSubscriptionPatch`

NewLocationSubscriptionPatchWithDefaults instantiates a new LocationSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventReq

`func (o *LocationSubscriptionPatch) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *LocationSubscriptionPatch) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *LocationSubscriptionPatch) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *LocationSubscriptionPatch) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.

### GetExpTime

`func (o *LocationSubscriptionPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *LocationSubscriptionPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *LocationSubscriptionPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *LocationSubscriptionPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *LocationSubscriptionPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *LocationSubscriptionPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *LocationSubscriptionPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *LocationSubscriptionPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### GetRevocationNotifUri

`func (o *LocationSubscriptionPatch) GetRevocationNotifUri() string`

GetRevocationNotifUri returns the RevocationNotifUri field if non-nil, zero value otherwise.

### GetRevocationNotifUriOk

`func (o *LocationSubscriptionPatch) GetRevocationNotifUriOk() (*string, bool)`

GetRevocationNotifUriOk returns a tuple with the RevocationNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationNotifUri

`func (o *LocationSubscriptionPatch) SetRevocationNotifUri(v string)`

SetRevocationNotifUri sets RevocationNotifUri field to given value.

### HasRevocationNotifUri

`func (o *LocationSubscriptionPatch) HasRevocationNotifUri() bool`

HasRevocationNotifUri returns a boolean if a field has been set.

### GetLocGran

`func (o *LocationSubscriptionPatch) GetLocGran() Accuracy`

GetLocGran returns the LocGran field if non-nil, zero value otherwise.

### GetLocGranOk

`func (o *LocationSubscriptionPatch) GetLocGranOk() (*Accuracy, bool)`

GetLocGranOk returns a tuple with the LocGran field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocGran

`func (o *LocationSubscriptionPatch) SetLocGran(v Accuracy)`

SetLocGran sets LocGran field to given value.

### HasLocGran

`func (o *LocationSubscriptionPatch) HasLocGran() bool`

HasLocGran returns a boolean if a field has been set.

### GetLocQos

`func (o *LocationSubscriptionPatch) GetLocQos() LocationQoS`

GetLocQos returns the LocQos field if non-nil, zero value otherwise.

### GetLocQosOk

`func (o *LocationSubscriptionPatch) GetLocQosOk() (*LocationQoS, bool)`

GetLocQosOk returns a tuple with the LocQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocQos

`func (o *LocationSubscriptionPatch) SetLocQos(v LocationQoS)`

SetLocQos sets LocQos field to given value.

### HasLocQos

`func (o *LocationSubscriptionPatch) HasLocQos() bool`

HasLocQos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


