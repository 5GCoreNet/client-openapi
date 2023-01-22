# CollectiveBehaviourFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CollectiveBehaviourFilterType**](CollectiveBehaviourFilterType.md) |  | 
**Value** | **string** | Value of the parameter type as in the type attribute. | 
**ListOfUeInd** | Pointer to **bool** | Indicates whether request list of UE IDs that fulfill a collective behaviour within the area of interest. This attribute shall set to \&quot;true\&quot; if request the list of UE IDs, otherwise, set to \&quot;false\&quot;. May only be present and sets to \&quot;true\&quot; if \&quot;AfEvent\&quot; sets to \&quot;COLLECTIVE_BEHAVIOUR\&quot;.  | [optional] 

## Methods

### NewCollectiveBehaviourFilter

`func NewCollectiveBehaviourFilter(type_ CollectiveBehaviourFilterType, value string, ) *CollectiveBehaviourFilter`

NewCollectiveBehaviourFilter instantiates a new CollectiveBehaviourFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectiveBehaviourFilterWithDefaults

`func NewCollectiveBehaviourFilterWithDefaults() *CollectiveBehaviourFilter`

NewCollectiveBehaviourFilterWithDefaults instantiates a new CollectiveBehaviourFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CollectiveBehaviourFilter) GetType() CollectiveBehaviourFilterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectiveBehaviourFilter) GetTypeOk() (*CollectiveBehaviourFilterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectiveBehaviourFilter) SetType(v CollectiveBehaviourFilterType)`

SetType sets Type field to given value.


### GetValue

`func (o *CollectiveBehaviourFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CollectiveBehaviourFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CollectiveBehaviourFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetListOfUeInd

`func (o *CollectiveBehaviourFilter) GetListOfUeInd() bool`

GetListOfUeInd returns the ListOfUeInd field if non-nil, zero value otherwise.

### GetListOfUeIndOk

`func (o *CollectiveBehaviourFilter) GetListOfUeIndOk() (*bool, bool)`

GetListOfUeIndOk returns a tuple with the ListOfUeInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListOfUeInd

`func (o *CollectiveBehaviourFilter) SetListOfUeInd(v bool)`

SetListOfUeInd sets ListOfUeInd field to given value.

### HasListOfUeInd

`func (o *CollectiveBehaviourFilter) HasListOfUeInd() bool`

HasListOfUeInd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


