# NetworkStatusReportingNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscription** | **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | 
**NsiValue** | Pointer to **int32** | Unsigned integer with valid values between 0 and 31. The value 0 indicates that there is no congestion. The value 1 is the lowest congestion level and value 31 is the highest congestion level. | [optional] 
**NsiType** | Pointer to [**CongestionType**](CongestionType.md) |  | [optional] 

## Methods

### NewNetworkStatusReportingNotification

`func NewNetworkStatusReportingNotification(subscription string, ) *NetworkStatusReportingNotification`

NewNetworkStatusReportingNotification instantiates a new NetworkStatusReportingNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkStatusReportingNotificationWithDefaults

`func NewNetworkStatusReportingNotificationWithDefaults() *NetworkStatusReportingNotification`

NewNetworkStatusReportingNotificationWithDefaults instantiates a new NetworkStatusReportingNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscription

`func (o *NetworkStatusReportingNotification) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *NetworkStatusReportingNotification) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *NetworkStatusReportingNotification) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetNsiValue

`func (o *NetworkStatusReportingNotification) GetNsiValue() int32`

GetNsiValue returns the NsiValue field if non-nil, zero value otherwise.

### GetNsiValueOk

`func (o *NetworkStatusReportingNotification) GetNsiValueOk() (*int32, bool)`

GetNsiValueOk returns a tuple with the NsiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiValue

`func (o *NetworkStatusReportingNotification) SetNsiValue(v int32)`

SetNsiValue sets NsiValue field to given value.

### HasNsiValue

`func (o *NetworkStatusReportingNotification) HasNsiValue() bool`

HasNsiValue returns a boolean if a field has been set.

### GetNsiType

`func (o *NetworkStatusReportingNotification) GetNsiType() CongestionType`

GetNsiType returns the NsiType field if non-nil, zero value otherwise.

### GetNsiTypeOk

`func (o *NetworkStatusReportingNotification) GetNsiTypeOk() (*CongestionType, bool)`

GetNsiTypeOk returns a tuple with the NsiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsiType

`func (o *NetworkStatusReportingNotification) SetNsiType(v CongestionType)`

SetNsiType sets NsiType field to given value.

### HasNsiType

`func (o *NetworkStatusReportingNotification) HasNsiType() bool`

HasNsiType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


