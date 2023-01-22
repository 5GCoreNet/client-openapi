# TransactionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | 
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**AppIds** | Pointer to **[]string** |  | [optional] 
**TransacMetrics** | Pointer to [**[]TransactionMetric**](TransactionMetric.md) |  | [optional] 

## Methods

### NewTransactionInfo

`func NewTransactionInfo(transaction int32, ) *TransactionInfo`

NewTransactionInfo instantiates a new TransactionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInfoWithDefaults

`func NewTransactionInfoWithDefaults() *TransactionInfo`

NewTransactionInfoWithDefaults instantiates a new TransactionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *TransactionInfo) GetTransaction() int32`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *TransactionInfo) GetTransactionOk() (*int32, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *TransactionInfo) SetTransaction(v int32)`

SetTransaction sets Transaction field to given value.


### GetSnssai

`func (o *TransactionInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *TransactionInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *TransactionInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *TransactionInfo) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetAppIds

`func (o *TransactionInfo) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *TransactionInfo) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *TransactionInfo) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *TransactionInfo) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### GetTransacMetrics

`func (o *TransactionInfo) GetTransacMetrics() []TransactionMetric`

GetTransacMetrics returns the TransacMetrics field if non-nil, zero value otherwise.

### GetTransacMetricsOk

`func (o *TransactionInfo) GetTransacMetricsOk() (*[]TransactionMetric, bool)`

GetTransacMetricsOk returns a tuple with the TransacMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransacMetrics

`func (o *TransactionInfo) SetTransacMetrics(v []TransactionMetric)`

SetTransacMetrics sets TransacMetrics field to given value.

### HasTransacMetrics

`func (o *TransactionInfo) HasTransacMetrics() bool`

HasTransacMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


