# N32fReformattedRspMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReformattedData** | [**FlatJweJson**](FlatJweJson.md) |  | 
**ModificationsBlock** | Pointer to [**[]FlatJwsJson**](FlatJwsJson.md) |  | [optional] 

## Methods

### NewN32fReformattedRspMsg

`func NewN32fReformattedRspMsg(reformattedData FlatJweJson, ) *N32fReformattedRspMsg`

NewN32fReformattedRspMsg instantiates a new N32fReformattedRspMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN32fReformattedRspMsgWithDefaults

`func NewN32fReformattedRspMsgWithDefaults() *N32fReformattedRspMsg`

NewN32fReformattedRspMsgWithDefaults instantiates a new N32fReformattedRspMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReformattedData

`func (o *N32fReformattedRspMsg) GetReformattedData() FlatJweJson`

GetReformattedData returns the ReformattedData field if non-nil, zero value otherwise.

### GetReformattedDataOk

`func (o *N32fReformattedRspMsg) GetReformattedDataOk() (*FlatJweJson, bool)`

GetReformattedDataOk returns a tuple with the ReformattedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReformattedData

`func (o *N32fReformattedRspMsg) SetReformattedData(v FlatJweJson)`

SetReformattedData sets ReformattedData field to given value.


### GetModificationsBlock

`func (o *N32fReformattedRspMsg) GetModificationsBlock() []FlatJwsJson`

GetModificationsBlock returns the ModificationsBlock field if non-nil, zero value otherwise.

### GetModificationsBlockOk

`func (o *N32fReformattedRspMsg) GetModificationsBlockOk() (*[]FlatJwsJson, bool)`

GetModificationsBlockOk returns a tuple with the ModificationsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationsBlock

`func (o *N32fReformattedRspMsg) SetModificationsBlock(v []FlatJwsJson)`

SetModificationsBlock sets ModificationsBlock field to given value.

### HasModificationsBlock

`func (o *N32fReformattedRspMsg) HasModificationsBlock() bool`

HasModificationsBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


