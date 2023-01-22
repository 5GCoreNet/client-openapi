# Ifc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **int32** |  | 
**Trigger** | Pointer to [**TriggerPoint**](TriggerPoint.md) |  | [optional] 
**AppServer** | [**ApplicationServer**](ApplicationServer.md) |  | 

## Methods

### NewIfc

`func NewIfc(priority int32, appServer ApplicationServer, ) *Ifc`

NewIfc instantiates a new Ifc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIfcWithDefaults

`func NewIfcWithDefaults() *Ifc`

NewIfcWithDefaults instantiates a new Ifc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *Ifc) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Ifc) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Ifc) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetTrigger

`func (o *Ifc) GetTrigger() TriggerPoint`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Ifc) GetTriggerOk() (*TriggerPoint, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Ifc) SetTrigger(v TriggerPoint)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *Ifc) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetAppServer

`func (o *Ifc) GetAppServer() ApplicationServer`

GetAppServer returns the AppServer field if non-nil, zero value otherwise.

### GetAppServerOk

`func (o *Ifc) GetAppServerOk() (*ApplicationServer, bool)`

GetAppServerOk returns a tuple with the AppServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppServer

`func (o *Ifc) SetAppServer(v ApplicationServer)`

SetAppServer sets AppServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


