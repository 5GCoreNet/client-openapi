# NotifyMoiAttributeValueChangesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelatedNotifications** | Pointer to [**[]CorrelatedNotification**](CorrelatedNotification.md) |  | [optional] 
**AdditionalText** | Pointer to **string** |  | [optional] 
**SourceIndicator** | Pointer to [**SourceIndicator**](SourceIndicator.md) |  | [optional] 
**AttributeListValueChanges** | **[]map[string]interface{}** | The first array item contains the attribute name value pairs with the new values, and the second array item the attribute name value pairs with the optional old values. | 

## Methods

### NewNotifyMoiAttributeValueChangesAllOf

`func NewNotifyMoiAttributeValueChangesAllOf(attributeListValueChanges []map[string]interface{}, ) *NotifyMoiAttributeValueChangesAllOf`

NewNotifyMoiAttributeValueChangesAllOf instantiates a new NotifyMoiAttributeValueChangesAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyMoiAttributeValueChangesAllOfWithDefaults

`func NewNotifyMoiAttributeValueChangesAllOfWithDefaults() *NotifyMoiAttributeValueChangesAllOf`

NewNotifyMoiAttributeValueChangesAllOfWithDefaults instantiates a new NotifyMoiAttributeValueChangesAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelatedNotifications

`func (o *NotifyMoiAttributeValueChangesAllOf) GetCorrelatedNotifications() []CorrelatedNotification`

GetCorrelatedNotifications returns the CorrelatedNotifications field if non-nil, zero value otherwise.

### GetCorrelatedNotificationsOk

`func (o *NotifyMoiAttributeValueChangesAllOf) GetCorrelatedNotificationsOk() (*[]CorrelatedNotification, bool)`

GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelatedNotifications

`func (o *NotifyMoiAttributeValueChangesAllOf) SetCorrelatedNotifications(v []CorrelatedNotification)`

SetCorrelatedNotifications sets CorrelatedNotifications field to given value.

### HasCorrelatedNotifications

`func (o *NotifyMoiAttributeValueChangesAllOf) HasCorrelatedNotifications() bool`

HasCorrelatedNotifications returns a boolean if a field has been set.

### GetAdditionalText

`func (o *NotifyMoiAttributeValueChangesAllOf) GetAdditionalText() string`

GetAdditionalText returns the AdditionalText field if non-nil, zero value otherwise.

### GetAdditionalTextOk

`func (o *NotifyMoiAttributeValueChangesAllOf) GetAdditionalTextOk() (*string, bool)`

GetAdditionalTextOk returns a tuple with the AdditionalText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalText

`func (o *NotifyMoiAttributeValueChangesAllOf) SetAdditionalText(v string)`

SetAdditionalText sets AdditionalText field to given value.

### HasAdditionalText

`func (o *NotifyMoiAttributeValueChangesAllOf) HasAdditionalText() bool`

HasAdditionalText returns a boolean if a field has been set.

### GetSourceIndicator

`func (o *NotifyMoiAttributeValueChangesAllOf) GetSourceIndicator() SourceIndicator`

GetSourceIndicator returns the SourceIndicator field if non-nil, zero value otherwise.

### GetSourceIndicatorOk

`func (o *NotifyMoiAttributeValueChangesAllOf) GetSourceIndicatorOk() (*SourceIndicator, bool)`

GetSourceIndicatorOk returns a tuple with the SourceIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIndicator

`func (o *NotifyMoiAttributeValueChangesAllOf) SetSourceIndicator(v SourceIndicator)`

SetSourceIndicator sets SourceIndicator field to given value.

### HasSourceIndicator

`func (o *NotifyMoiAttributeValueChangesAllOf) HasSourceIndicator() bool`

HasSourceIndicator returns a boolean if a field has been set.

### GetAttributeListValueChanges

`func (o *NotifyMoiAttributeValueChangesAllOf) GetAttributeListValueChanges() []map[string]interface{}`

GetAttributeListValueChanges returns the AttributeListValueChanges field if non-nil, zero value otherwise.

### GetAttributeListValueChangesOk

`func (o *NotifyMoiAttributeValueChangesAllOf) GetAttributeListValueChangesOk() (*[]map[string]interface{}, bool)`

GetAttributeListValueChangesOk returns a tuple with the AttributeListValueChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeListValueChanges

`func (o *NotifyMoiAttributeValueChangesAllOf) SetAttributeListValueChanges(v []map[string]interface{})`

SetAttributeListValueChanges sets AttributeListValueChanges field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


