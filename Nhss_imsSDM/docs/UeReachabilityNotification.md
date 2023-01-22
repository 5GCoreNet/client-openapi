# UeReachabilityNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReachabilityIndicator** | **bool** |  | 
**DetectingNode** | [**DetectingNode**](DetectingNode.md) |  | 
**AccessType** | [**AccessType**](AccessType.md) |  | 

## Methods

### NewUeReachabilityNotification

`func NewUeReachabilityNotification(reachabilityIndicator bool, detectingNode DetectingNode, accessType AccessType, ) *UeReachabilityNotification`

NewUeReachabilityNotification instantiates a new UeReachabilityNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeReachabilityNotificationWithDefaults

`func NewUeReachabilityNotificationWithDefaults() *UeReachabilityNotification`

NewUeReachabilityNotificationWithDefaults instantiates a new UeReachabilityNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachabilityIndicator

`func (o *UeReachabilityNotification) GetReachabilityIndicator() bool`

GetReachabilityIndicator returns the ReachabilityIndicator field if non-nil, zero value otherwise.

### GetReachabilityIndicatorOk

`func (o *UeReachabilityNotification) GetReachabilityIndicatorOk() (*bool, bool)`

GetReachabilityIndicatorOk returns a tuple with the ReachabilityIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityIndicator

`func (o *UeReachabilityNotification) SetReachabilityIndicator(v bool)`

SetReachabilityIndicator sets ReachabilityIndicator field to given value.


### GetDetectingNode

`func (o *UeReachabilityNotification) GetDetectingNode() DetectingNode`

GetDetectingNode returns the DetectingNode field if non-nil, zero value otherwise.

### GetDetectingNodeOk

`func (o *UeReachabilityNotification) GetDetectingNodeOk() (*DetectingNode, bool)`

GetDetectingNodeOk returns a tuple with the DetectingNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectingNode

`func (o *UeReachabilityNotification) SetDetectingNode(v DetectingNode)`

SetDetectingNode sets DetectingNode field to given value.


### GetAccessType

`func (o *UeReachabilityNotification) GetAccessType() AccessType`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *UeReachabilityNotification) GetAccessTypeOk() (*AccessType, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *UeReachabilityNotification) SetAccessType(v AccessType)`

SetAccessType sets AccessType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


