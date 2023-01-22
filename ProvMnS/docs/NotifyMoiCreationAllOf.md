# NotifyMoiCreationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**SourceIndicator** | Pointer to [**SourceIndicator**](SourceIndicator.md) |  | [optional] 
**AttributeList** | Pointer to **map[string]interface{}** | The key of this map is the attribute name, and the value the attribute value. | [optional] 

## Methods

### NewNotifyMoiCreationAllOf

`func NewNotifyMoiCreationAllOf() *NotifyMoiCreationAllOf`

NewNotifyMoiCreationAllOf instantiates a new NotifyMoiCreationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyMoiCreationAllOfWithDefaults

`func NewNotifyMoiCreationAllOfWithDefaults() *NotifyMoiCreationAllOf`

NewNotifyMoiCreationAllOfWithDefaults instantiates a new NotifyMoiCreationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelatedNotifications

`func (o *NotifyMoiCreationAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyMoiCreationAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyMoiCreationAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyMoiCreationAllOf) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyMoiCreationAllOf) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyMoiCreationAllOf) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyMoiCreationAllOf) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyMoiCreationAllOf) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetSourceIndicator

`func (o *NotifyMoiCreationAllOf) GetSourceIndicator() SourceIndicator`

GetSourceIndicator returns the SourceIndicator field if non-nil, zero value otherwise.

### GetSourceIndicatorOk

`func (o *NotifyMoiCreationAllOf) GetSourceIndicatorOk() (*SourceIndicator, bool)`

GetSourceIndicatorOk returns a tuple with the SourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIndicator

`func (o *NotifyMoiCreationAllOf) SetSourceIndicator(v SourceIndicator)`

SetSourceIndicator sets SourceIndicator field to given value.

### HasSourceIndicator

`func (o *NotifyMoiCreationAllOf) HasSourceIndicator() bool`

HasSourceIndicator returns a boolean if a field has been set.

### GetAttributeList

`func (o *NotifyMoiCreationAllOf) GetAttributeList() map[string]interface{}`

GetAttributeList returns the AttributeList field if non-nil, zero value otherwise.

### GetAttributeListOk

`func (o *NotifyMoiCreationAllOf) GetAttributeListOk() (*map[string]interface{}, bool)`

GetAttributeListOk returns a tuple with the AttributeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeList

`func (o *NotifyMoiCreationAllOf) SetAttributeList(v map[string]interface{})`

SetAttributeList sets AttributeList field to given value.

### HasAttributeList

`func (o *NotifyMoiCreationAllOf) HasAttributeList() bool`

HasAttributeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


