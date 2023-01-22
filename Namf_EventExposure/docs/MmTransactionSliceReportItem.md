# MmTransactionSliceReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**Transactions** | **int32** |  | 

## Methods

### NewMmTransactionSliceReportItem

`func NewMmTransactionSliceReportItem(timestamp time.Time, transactions int32, ) *MmTransactionSliceReportItem`

NewMmTransactionSliceReportItem instantiates a new MmTransactionSliceReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMmTransactionSliceReportItemWithDefaults

`func NewMmTransactionSliceReportItemWithDefaults() *MmTransactionSliceReportItem`

NewMmTransactionSliceReportItemWithDefaults instantiates a new MmTransactionSliceReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *MmTransactionSliceReportItem) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *MmTransactionSliceReportItem) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *MmTransactionSliceReportItem) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.

### HasSnssai

`func (o *MmTransactionSliceReportItem) HasSnssai() bool`

HasSnssai returns a boolean if a field has been set.

### GetTimestamp

`func (o *MmTransactionSliceReportItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MmTransactionSliceReportItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MmTransactionSliceReportItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTransactions

`func (o *MmTransactionSliceReportItem) GetTransactions() int32`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *MmTransactionSliceReportItem) GetTransactionsOk() (*int32, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *MmTransactionSliceReportItem) SetTransactions(v int32)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


