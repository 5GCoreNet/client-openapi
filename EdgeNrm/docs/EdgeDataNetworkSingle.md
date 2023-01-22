# EdgeDataNetworkSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**EdnIdentifier** | Pointer to **string** |  | [optional] 
**EDNConnectionInfo** | Pointer to [**EDNConnectionInfo**](EDNConnectionInfo.md) |  | [optional] 
**EASFunction** | Pointer to [**[]EASFunctionSingle**](EASFunctionSingle.md) |  | [optional] 
**EESFunction** | Pointer to [**[]EESFunctionSingle**](EESFunctionSingle.md) |  | [optional] 

## Methods

### NewEdgeDataNetworkSingle

`func NewEdgeDataNetworkSingle(id NullableString, ) *EdgeDataNetworkSingle`

NewEdgeDataNetworkSingle instantiates a new EdgeDataNetworkSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEdgeDataNetworkSingleWithDefaults

`func NewEdgeDataNetworkSingleWithDefaults() *EdgeDataNetworkSingle`

NewEdgeDataNetworkSingleWithDefaults instantiates a new EdgeDataNetworkSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EdgeDataNetworkSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EdgeDataNetworkSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EdgeDataNetworkSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *EdgeDataNetworkSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *EdgeDataNetworkSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *EdgeDataNetworkSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *EdgeDataNetworkSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *EdgeDataNetworkSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *EdgeDataNetworkSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *EdgeDataNetworkSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *EdgeDataNetworkSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *EdgeDataNetworkSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *EdgeDataNetworkSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *EdgeDataNetworkSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *EdgeDataNetworkSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *EdgeDataNetworkSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *EdgeDataNetworkSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetEdnIdentifier

`func (o *EdgeDataNetworkSingle) GetEdnIdentifier() string`

GetEdnIdentifier returns the EdnIdentifier field if non-nil, zero value otherwise.

### GetEdnIdentifierOk

`func (o *EdgeDataNetworkSingle) GetEdnIdentifierOk() (*string, bool)`

GetEdnIdentifierOk returns a tuple with the EdnIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdnIdentifier

`func (o *EdgeDataNetworkSingle) SetEdnIdentifier(v string)`

SetEdnIdentifier sets EdnIdentifier field to given value.

### HasEdnIdentifier

`func (o *EdgeDataNetworkSingle) HasEdnIdentifier() bool`

HasEdnIdentifier returns a boolean if a field has been set.

### GetEDNConnectionInfo

`func (o *EdgeDataNetworkSingle) GetEDNConnectionInfo() EDNConnectionInfo`

GetEDNConnectionInfo returns the EDNConnectionInfo field if non-nil, zero value otherwise.

### GetEDNConnectionInfoOk

`func (o *EdgeDataNetworkSingle) GetEDNConnectionInfoOk() (*EDNConnectionInfo, bool)`

GetEDNConnectionInfoOk returns a tuple with the EDNConnectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEDNConnectionInfo

`func (o *EdgeDataNetworkSingle) SetEDNConnectionInfo(v EDNConnectionInfo)`

SetEDNConnectionInfo sets EDNConnectionInfo field to given value.

### HasEDNConnectionInfo

`func (o *EdgeDataNetworkSingle) HasEDNConnectionInfo() bool`

HasEDNConnectionInfo returns a boolean if a field has been set.

### GetEASFunction

`func (o *EdgeDataNetworkSingle) GetEASFunction() []EASFunctionSingle`

GetEASFunction returns the EASFunction field if non-nil, zero value otherwise.

### GetEASFunctionOk

`func (o *EdgeDataNetworkSingle) GetEASFunctionOk() (*[]EASFunctionSingle, bool)`

GetEASFunctionOk returns a tuple with the EASFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEASFunction

`func (o *EdgeDataNetworkSingle) SetEASFunction(v []EASFunctionSingle)`

SetEASFunction sets EASFunction field to given value.

### HasEASFunction

`func (o *EdgeDataNetworkSingle) HasEASFunction() bool`

HasEASFunction returns a boolean if a field has been set.

### GetEESFunction

`func (o *EdgeDataNetworkSingle) GetEESFunction() []EESFunctionSingle`

GetEESFunction returns the EESFunction field if non-nil, zero value otherwise.

### GetEESFunctionOk

`func (o *EdgeDataNetworkSingle) GetEESFunctionOk() (*[]EESFunctionSingle, bool)`

GetEESFunctionOk returns a tuple with the EESFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEESFunction

`func (o *EdgeDataNetworkSingle) SetEESFunction(v []EESFunctionSingle)`

SetEESFunction sets EESFunction field to given value.

### HasEESFunction

`func (o *EdgeDataNetworkSingle) HasEESFunction() bool`

HasEESFunction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


