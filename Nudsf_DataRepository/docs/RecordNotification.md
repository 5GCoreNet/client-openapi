# RecordNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Descriptor** | [**NotificationDescription**](NotificationDescription.md) |  | 
**Meta** | [**RecordMeta**](RecordMeta.md) |  | 
**Blocks** | Pointer to **[]interface{}** | list of opaque Block&#39;s in this Record | [optional] 

## Methods

### NewRecordNotification

`func NewRecordNotification(descriptor NotificationDescription, meta RecordMeta, ) *RecordNotification`

NewRecordNotification instantiates a new RecordNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordNotificationWithDefaults

`func NewRecordNotificationWithDefaults() *RecordNotification`

NewRecordNotificationWithDefaults instantiates a new RecordNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescriptor

`func (o *RecordNotification) GetDescriptor() NotificationDescription`

GetDescriptor returns the Descriptor field if non-nil, zero value otherwise.

### GetDescriptorOk

`func (o *RecordNotification) GetDescriptorOk() (*NotificationDescription, bool)`

GetDescriptorOk returns a tuple with the Descriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptor

`func (o *RecordNotification) SetDescriptor(v NotificationDescription)`

SetDescriptor sets Descriptor field to given value.


### GetMeta

`func (o *RecordNotification) GetMeta() RecordMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RecordNotification) GetMetaOk() (*RecordMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RecordNotification) SetMeta(v RecordMeta)`

SetMeta sets Meta field to given value.


### GetBlocks

`func (o *RecordNotification) GetBlocks() []interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *RecordNotification) GetBlocksOk() (*[]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *RecordNotification) SetBlocks(v []interface{})`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *RecordNotification) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


