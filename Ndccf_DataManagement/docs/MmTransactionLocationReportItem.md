# MmTransactionLocationReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tai** | Pointer to [**Tai**](Tai.md) |  | [optional] 
**Ncgi** | Pointer to [**Ncgi**](Ncgi.md) |  | [optional] 
**Ecgi** | Pointer to [**Ecgi**](Ecgi.md) |  | [optional] 
**N3gaLocation** | Pointer to [**N3gaLocation**](N3gaLocation.md) |  | [optional] 
**Timestamp** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**Transactions** | **int32** |  | 

## Methods

### NewMmTransactionLocationReportItem

`func NewMmTransactionLocationReportItem(timestamp time.Time, transactions int32, ) *MmTransactionLocationReportItem`

NewMmTransactionLocationReportItem instantiates a new MmTransactionLocationReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMmTransactionLocationReportItemWithDefaults

`func NewMmTransactionLocationReportItemWithDefaults() *MmTransactionLocationReportItem`

NewMmTransactionLocationReportItemWithDefaults instantiates a new MmTransactionLocationReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTai

`func (o *MmTransactionLocationReportItem) GetTai() Tai`

GetTai returns the Tai field if non-nil, zero value otherwise.

### GetTaiOk

`func (o *MmTransactionLocationReportItem) GetTaiOk() (*Tai, bool)`

GetTaiOk returns a tuple with the Tai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTai

`func (o *MmTransactionLocationReportItem) SetTai(v Tai)`

SetTai sets Tai field to given value.

### HasTai

`func (o *MmTransactionLocationReportItem) HasTai() bool`

HasTai returns a boolean if a field has been set.

### GetNcgi

`func (o *MmTransactionLocationReportItem) GetNcgi() Ncgi`

GetNcgi returns the Ncgi field if non-nil, zero value otherwise.

### GetNcgiOk

`func (o *MmTransactionLocationReportItem) GetNcgiOk() (*Ncgi, bool)`

GetNcgiOk returns a tuple with the Ncgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcgi

`func (o *MmTransactionLocationReportItem) SetNcgi(v Ncgi)`

SetNcgi sets Ncgi field to given value.

### HasNcgi

`func (o *MmTransactionLocationReportItem) HasNcgi() bool`

HasNcgi returns a boolean if a field has been set.

### GetEcgi

`func (o *MmTransactionLocationReportItem) GetEcgi() Ecgi`

GetEcgi returns the Ecgi field if non-nil, zero value otherwise.

### GetEcgiOk

`func (o *MmTransactionLocationReportItem) GetEcgiOk() (*Ecgi, bool)`

GetEcgiOk returns a tuple with the Ecgi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcgi

`func (o *MmTransactionLocationReportItem) SetEcgi(v Ecgi)`

SetEcgi sets Ecgi field to given value.

### HasEcgi

`func (o *MmTransactionLocationReportItem) HasEcgi() bool`

HasEcgi returns a boolean if a field has been set.

### GetN3gaLocation

`func (o *MmTransactionLocationReportItem) GetN3gaLocation() N3gaLocation`

GetN3gaLocation returns the N3gaLocation field if non-nil, zero value otherwise.

### GetN3gaLocationOk

`func (o *MmTransactionLocationReportItem) GetN3gaLocationOk() (*N3gaLocation, bool)`

GetN3gaLocationOk returns a tuple with the N3gaLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN3gaLocation

`func (o *MmTransactionLocationReportItem) SetN3gaLocation(v N3gaLocation)`

SetN3gaLocation sets N3gaLocation field to given value.

### HasN3gaLocation

`func (o *MmTransactionLocationReportItem) HasN3gaLocation() bool`

HasN3gaLocation returns a boolean if a field has been set.

### GetTimestamp

`func (o *MmTransactionLocationReportItem) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MmTransactionLocationReportItem) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MmTransactionLocationReportItem) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTransactions

`func (o *MmTransactionLocationReportItem) GetTransactions() int32`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *MmTransactionLocationReportItem) GetTransactionsOk() (*int32, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *MmTransactionLocationReportItem) SetTransactions(v int32)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


