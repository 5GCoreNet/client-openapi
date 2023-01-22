# ScpDomainRoutingInfoNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingInfo** | [**ScpDomainRoutingInformation**](ScpDomainRoutingInformation.md) |  | 
**LocalInd** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewScpDomainRoutingInfoNotification

`func NewScpDomainRoutingInfoNotification(routingInfo ScpDomainRoutingInformation, ) *ScpDomainRoutingInfoNotification`

NewScpDomainRoutingInfoNotification instantiates a new ScpDomainRoutingInfoNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScpDomainRoutingInfoNotificationWithDefaults

`func NewScpDomainRoutingInfoNotificationWithDefaults() *ScpDomainRoutingInfoNotification`

NewScpDomainRoutingInfoNotificationWithDefaults instantiates a new ScpDomainRoutingInfoNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutingInfo

`func (o *ScpDomainRoutingInfoNotification) GetRoutingInfo() ScpDomainRoutingInformation`

GetRoutingInfo returns the RoutingInfo field if non-nil, zero value otherwise.

### GetRoutingInfoOk

`func (o *ScpDomainRoutingInfoNotification) GetRoutingInfoOk() (*ScpDomainRoutingInformation, bool)`

GetRoutingInfoOk returns a tuple with the RoutingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingInfo

`func (o *ScpDomainRoutingInfoNotification) SetRoutingInfo(v ScpDomainRoutingInformation)`

SetRoutingInfo sets RoutingInfo field to given value.


### GetLocalInd

`func (o *ScpDomainRoutingInfoNotification) GetLocalInd() bool`

GetLocalInd returns the LocalInd field if non-nil, zero value otherwise.

### GetLocalIndOk

`func (o *ScpDomainRoutingInfoNotification) GetLocalIndOk() (*bool, bool)`

GetLocalIndOk returns a tuple with the LocalInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInd

`func (o *ScpDomainRoutingInfoNotification) SetLocalInd(v bool)`

SetLocalInd sets LocalInd field to given value.

### HasLocalInd

`func (o *ScpDomainRoutingInfoNotification) HasLocalInd() bool`

HasLocalInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


