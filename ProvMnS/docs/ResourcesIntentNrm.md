# ResourcesIntentNrm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**UserLabel** | Pointer to **string** |  | [optional] 
**IntentExpectations** | Pointer to [**[]OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation**](OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation.md) |  | [optional] 
**IntentContexts** | Pointer to [**[]IntentContext**](IntentContext.md) |  | [optional] 
**IntentFulfilmentInfo** | Pointer to [**FulfilmentInfo**](FulfilmentInfo.md) |  | [optional] 

## Methods

### NewResourcesIntentNrm

`func NewResourcesIntentNrm(id NullableString, ) *ResourcesIntentNrm`

NewResourcesIntentNrm instantiates a new ResourcesIntentNrm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesIntentNrmWithDefaults

`func NewResourcesIntentNrmWithDefaults() *ResourcesIntentNrm`

NewResourcesIntentNrmWithDefaults instantiates a new ResourcesIntentNrm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcesIntentNrm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcesIntentNrm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcesIntentNrm) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ResourcesIntentNrm) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ResourcesIntentNrm) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *ResourcesIntentNrm) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *ResourcesIntentNrm) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *ResourcesIntentNrm) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *ResourcesIntentNrm) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *ResourcesIntentNrm) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *ResourcesIntentNrm) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *ResourcesIntentNrm) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *ResourcesIntentNrm) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *ResourcesIntentNrm) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *ResourcesIntentNrm) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *ResourcesIntentNrm) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *ResourcesIntentNrm) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetUserLabel

`func (o *ResourcesIntentNrm) GetUserLabel() string`

GetUserLabel returns the UserLabel field if non-nil, zero value otherwise.

### GetUserLabelOk

`func (o *ResourcesIntentNrm) GetUserLabelOk() (*string, bool)`

GetUserLabelOk returns a tuple with the UserLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLabel

`func (o *ResourcesIntentNrm) SetUserLabel(v string)`

SetUserLabel sets UserLabel field to given value.

### HasUserLabel

`func (o *ResourcesIntentNrm) HasUserLabel() bool`

HasUserLabel returns a boolean if a field has been set.

### GetIntentExpectations

`func (o *ResourcesIntentNrm) GetIntentExpectations() []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation`

GetIntentExpectations returns the IntentExpectations field if non-nil, zero value otherwise.

### GetIntentExpectationsOk

`func (o *ResourcesIntentNrm) GetIntentExpectationsOk() (*[]OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation, bool)`

GetIntentExpectationsOk returns a tuple with the IntentExpectations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentExpectations

`func (o *ResourcesIntentNrm) SetIntentExpectations(v []OneOfIntentExpectationRadioNetworkExpectationServiceSupportExpectation)`

SetIntentExpectations sets IntentExpectations field to given value.

### HasIntentExpectations

`func (o *ResourcesIntentNrm) HasIntentExpectations() bool`

HasIntentExpectations returns a boolean if a field has been set.

### GetIntentContexts

`func (o *ResourcesIntentNrm) GetIntentContexts() []IntentContext`

GetIntentContexts returns the IntentContexts field if non-nil, zero value otherwise.

### GetIntentContextsOk

`func (o *ResourcesIntentNrm) GetIntentContextsOk() (*[]IntentContext, bool)`

GetIntentContextsOk returns a tuple with the IntentContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentContexts

`func (o *ResourcesIntentNrm) SetIntentContexts(v []IntentContext)`

SetIntentContexts sets IntentContexts field to given value.

### HasIntentContexts

`func (o *ResourcesIntentNrm) HasIntentContexts() bool`

HasIntentContexts returns a boolean if a field has been set.

### GetIntentFulfilmentInfo

`func (o *ResourcesIntentNrm) GetIntentFulfilmentInfo() FulfilmentInfo`

GetIntentFulfilmentInfo returns the IntentFulfilmentInfo field if non-nil, zero value otherwise.

### GetIntentFulfilmentInfoOk

`func (o *ResourcesIntentNrm) GetIntentFulfilmentInfoOk() (*FulfilmentInfo, bool)`

GetIntentFulfilmentInfoOk returns a tuple with the IntentFulfilmentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntentFulfilmentInfo

`func (o *ResourcesIntentNrm) SetIntentFulfilmentInfo(v FulfilmentInfo)`

SetIntentFulfilmentInfo sets IntentFulfilmentInfo field to given value.

### HasIntentFulfilmentInfo

`func (o *ResourcesIntentNrm) HasIntentFulfilmentInfo() bool`

HasIntentFulfilmentInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


