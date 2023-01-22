# AppAmContextRespData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAmContextId** | Pointer to **string** | Contains the AM Policy Events Subscription resource identifier related to the event notification. | [optional] 
**RepEvents** | [**[]AmEventNotification**](AmEventNotification.md) |  | 

## Methods

### NewAppAmContextRespData

`func NewAppAmContextRespData(repEvents []AmEventNotification, ) *AppAmContextRespData`

NewAppAmContextRespData instantiates a new AppAmContextRespData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAmContextRespDataWithDefaults

`func NewAppAmContextRespDataWithDefaults() *AppAmContextRespData`

NewAppAmContextRespDataWithDefaults instantiates a new AppAmContextRespData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAmContextId

`func (o *AppAmContextRespData) GetAppAmContextId() string`

GetAppAmContextId returns the AppAmContextId field if non-nil, zero value otherwise.

### GetAppAmContextIdOk

`func (o *AppAmContextRespData) GetAppAmContextIdOk() (*string, bool)`

GetAppAmContextIdOk returns a tuple with the AppAmContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAmContextId

`func (o *AppAmContextRespData) SetAppAmContextId(v string)`

SetAppAmContextId sets AppAmContextId field to given value.

### HasAppAmContextId

`func (o *AppAmContextRespData) HasAppAmContextId() bool`

HasAppAmContextId returns a boolean if a field has been set.

### GetRepEvents

`func (o *AppAmContextRespData) GetRepEvents() []AmEventNotification`

GetRepEvents returns the RepEvents field if non-nil, zero value otherwise.

### GetRepEventsOk

`func (o *AppAmContextRespData) GetRepEventsOk() (*[]AmEventNotification, bool)`

GetRepEventsOk returns a tuple with the RepEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepEvents

`func (o *AppAmContextRespData) SetRepEvents(v []AmEventNotification)`

SetRepEvents sets RepEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


