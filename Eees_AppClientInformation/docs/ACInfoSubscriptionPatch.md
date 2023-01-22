# ACInfoSubscriptionPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcFltrs** | Pointer to [**[]ACFilters**](ACFilters.md) | Filters to retrieve the information about specific ACs. | [optional] 
**ExpTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**EventReq** | Pointer to [**ReportingInformation**](ReportingInformation.md) |  | [optional] 
**NotificationDestination** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 

## Methods

### NewACInfoSubscriptionPatch

`func NewACInfoSubscriptionPatch() *ACInfoSubscriptionPatch`

NewACInfoSubscriptionPatch instantiates a new ACInfoSubscriptionPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACInfoSubscriptionPatchWithDefaults

`func NewACInfoSubscriptionPatchWithDefaults() *ACInfoSubscriptionPatch`

NewACInfoSubscriptionPatchWithDefaults instantiates a new ACInfoSubscriptionPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcFltrs

`func (o *ACInfoSubscriptionPatch) GetAcFltrs() []ACFilters`

GetAcFltrs returns the AcFltrs field if non-nil, zero value otherwise.

### GetAcFltrsOk

`func (o *ACInfoSubscriptionPatch) GetAcFltrsOk() (*[]ACFilters, bool)`

GetAcFltrsOk returns a tuple with the AcFltrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcFltrs

`func (o *ACInfoSubscriptionPatch) SetAcFltrs(v []ACFilters)`

SetAcFltrs sets AcFltrs field to given value.

### HasAcFltrs

`func (o *ACInfoSubscriptionPatch) HasAcFltrs() bool`

HasAcFltrs returns a boolean if a field has been set.

### GetExpTime

`func (o *ACInfoSubscriptionPatch) GetExpTime() time.Time`

GetExpTime returns the ExpTime field if non-nil, zero value otherwise.

### GetExpTimeOk

`func (o *ACInfoSubscriptionPatch) GetExpTimeOk() (*time.Time, bool)`

GetExpTimeOk returns a tuple with the ExpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpTime

`func (o *ACInfoSubscriptionPatch) SetExpTime(v time.Time)`

SetExpTime sets ExpTime field to given value.

### HasExpTime

`func (o *ACInfoSubscriptionPatch) HasExpTime() bool`

HasExpTime returns a boolean if a field has been set.

### GetEventReq

`func (o *ACInfoSubscriptionPatch) GetEventReq() ReportingInformation`

GetEventReq returns the EventReq field if non-nil, zero value otherwise.

### GetEventReqOk

`func (o *ACInfoSubscriptionPatch) GetEventReqOk() (*ReportingInformation, bool)`

GetEventReqOk returns a tuple with the EventReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventReq

`func (o *ACInfoSubscriptionPatch) SetEventReq(v ReportingInformation)`

SetEventReq sets EventReq field to given value.

### HasEventReq

`func (o *ACInfoSubscriptionPatch) HasEventReq() bool`

HasEventReq returns a boolean if a field has been set.

### GetNotificationDestination

`func (o *ACInfoSubscriptionPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *ACInfoSubscriptionPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *ACInfoSubscriptionPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *ACInfoSubscriptionPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


