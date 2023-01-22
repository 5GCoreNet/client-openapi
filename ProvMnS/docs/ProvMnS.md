# ProvMnS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubNetwork** | Pointer to [**[]SubNetworkSingle**](SubNetworkSingle.md) |  | [optional] 
**ManagedElement** | Pointer to [**[]ManagedElementSingle**](ManagedElementSingle.md) |  | [optional] 

## Methods

### NewProvMnS

`func NewProvMnS() *ProvMnS`

NewProvMnS instantiates a new ProvMnS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvMnSWithDefaults

`func NewProvMnSWithDefaults() *ProvMnS`

NewProvMnSWithDefaults instantiates a new ProvMnS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubNetwork

`func (o *ProvMnS) GetSubNetwork() []SubNetworkSingle`

GetSubNetwork returns the SubNetwork field if non-nil, zero value otherwise.

### GetSubNetworkOk

`func (o *ProvMnS) GetSubNetworkOk() (*[]SubNetworkSingle, bool)`

GetSubNetworkOk returns a tuple with the SubNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNetwork

`func (o *ProvMnS) SetSubNetwork(v []SubNetworkSingle)`

SetSubNetwork sets SubNetwork field to given value.

### HasSubNetwork

`func (o *ProvMnS) HasSubNetwork() bool`

HasSubNetwork returns a boolean if a field has been set.

### GetManagedElement

`func (o *ProvMnS) GetManagedElement() []ManagedElementSingle`

GetManagedElement returns the ManagedElement field if non-nil, zero value otherwise.

### GetManagedElementOk

`func (o *ProvMnS) GetManagedElementOk() (*[]ManagedElementSingle, bool)`

GetManagedElementOk returns a tuple with the ManagedElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedElement

`func (o *ProvMnS) SetManagedElement(v []ManagedElementSingle)`

SetManagedElement sets ManagedElement field to given value.

### HasManagedElement

`func (o *ProvMnS) HasManagedElement() bool`

HasManagedElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


