# N32fReformattedReqMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReformattedData** | [**FlatJweJson**](FlatJweJson.md) |  | 
**ModificationsBlock** | Pointer to [**[]FlatJwsJson**](FlatJwsJson.md) |  | [optional] 

## Methods

### NewN32fReformattedReqMsg

`func NewN32fReformattedReqMsg(reformattedData FlatJweJson, ) *N32fReformattedReqMsg`

NewN32fReformattedReqMsg instantiates a new N32fReformattedReqMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewN32fReformattedReqMsgWithDefaults

`func NewN32fReformattedReqMsgWithDefaults() *N32fReformattedReqMsg`

NewN32fReformattedReqMsgWithDefaults instantiates a new N32fReformattedReqMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReformattedData

`func (o *N32fReformattedReqMsg) GetReformattedData() FlatJweJson`

GetReformattedData returns the ReformattedData field if non-nil, zero value otherwise.

### GetReformattedDataOk

`func (o *N32fReformattedReqMsg) GetReformattedDataOk() (*FlatJweJson, bool)`

GetReformattedDataOk returns a tuple with the ReformattedData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReformattedData

`func (o *N32fReformattedReqMsg) SetReformattedData(v FlatJweJson)`

SetReformattedData sets ReformattedData field to given value.


### GetModificationsBlock

`func (o *N32fReformattedReqMsg) GetModificationsBlock() []FlatJwsJson`

GetModificationsBlock returns the ModificationsBlock field if non-nil, zero value otherwise.

### GetModificationsBlockOk

`func (o *N32fReformattedReqMsg) GetModificationsBlockOk() (*[]FlatJwsJson, bool)`

GetModificationsBlockOk returns a tuple with the ModificationsBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationsBlock

`func (o *N32fReformattedReqMsg) SetModificationsBlock(v []FlatJwsJson)`

SetModificationsBlock sets ModificationsBlock field to given value.

### HasModificationsBlock

`func (o *N32fReformattedReqMsg) HasModificationsBlock() bool`

HasModificationsBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


