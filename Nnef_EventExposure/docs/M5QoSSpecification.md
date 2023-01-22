# M5QoSSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarBwDlBitRate** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**MarBwUlBitRate** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**MinDesBwDlBitRate** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MinDesBwUlBitRate** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**MirBwDlBitRate** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**MirBwUlBitRate** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**DesLatency** | Pointer to **int32** |  | [optional] 
**DesLoss** | Pointer to **int32** |  | [optional] 

## Methods

### NewM5QoSSpecification

`func NewM5QoSSpecification(marBwDlBitRate string, marBwUlBitRate string, mirBwDlBitRate string, mirBwUlBitRate string, ) *M5QoSSpecification`

NewM5QoSSpecification instantiates a new M5QoSSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM5QoSSpecificationWithDefaults

`func NewM5QoSSpecificationWithDefaults() *M5QoSSpecification`

NewM5QoSSpecificationWithDefaults instantiates a new M5QoSSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarBwDlBitRate

`func (o *M5QoSSpecification) GetMarBwDlBitRate() string`

GetMarBwDlBitRate returns the MarBwDlBitRate field if non-nil, zero value otherwise.

### GetMarBwDlBitRateOk

`func (o *M5QoSSpecification) GetMarBwDlBitRateOk() (*string, bool)`

GetMarBwDlBitRateOk returns a tuple with the MarBwDlBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwDlBitRate

`func (o *M5QoSSpecification) SetMarBwDlBitRate(v string)`

SetMarBwDlBitRate sets MarBwDlBitRate field to given value.


### GetMarBwUlBitRate

`func (o *M5QoSSpecification) GetMarBwUlBitRate() string`

GetMarBwUlBitRate returns the MarBwUlBitRate field if non-nil, zero value otherwise.

### GetMarBwUlBitRateOk

`func (o *M5QoSSpecification) GetMarBwUlBitRateOk() (*string, bool)`

GetMarBwUlBitRateOk returns a tuple with the MarBwUlBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarBwUlBitRate

`func (o *M5QoSSpecification) SetMarBwUlBitRate(v string)`

SetMarBwUlBitRate sets MarBwUlBitRate field to given value.


### GetMinDesBwDlBitRate

`func (o *M5QoSSpecification) GetMinDesBwDlBitRate() string`

GetMinDesBwDlBitRate returns the MinDesBwDlBitRate field if non-nil, zero value otherwise.

### GetMinDesBwDlBitRateOk

`func (o *M5QoSSpecification) GetMinDesBwDlBitRateOk() (*string, bool)`

GetMinDesBwDlBitRateOk returns a tuple with the MinDesBwDlBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDesBwDlBitRate

`func (o *M5QoSSpecification) SetMinDesBwDlBitRate(v string)`

SetMinDesBwDlBitRate sets MinDesBwDlBitRate field to given value.

### HasMinDesBwDlBitRate

`func (o *M5QoSSpecification) HasMinDesBwDlBitRate() bool`

HasMinDesBwDlBitRate returns a boolean if a field has been set.

### GetMinDesBwUlBitRate

`func (o *M5QoSSpecification) GetMinDesBwUlBitRate() string`

GetMinDesBwUlBitRate returns the MinDesBwUlBitRate field if non-nil, zero value otherwise.

### GetMinDesBwUlBitRateOk

`func (o *M5QoSSpecification) GetMinDesBwUlBitRateOk() (*string, bool)`

GetMinDesBwUlBitRateOk returns a tuple with the MinDesBwUlBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDesBwUlBitRate

`func (o *M5QoSSpecification) SetMinDesBwUlBitRate(v string)`

SetMinDesBwUlBitRate sets MinDesBwUlBitRate field to given value.

### HasMinDesBwUlBitRate

`func (o *M5QoSSpecification) HasMinDesBwUlBitRate() bool`

HasMinDesBwUlBitRate returns a boolean if a field has been set.

### GetMirBwDlBitRate

`func (o *M5QoSSpecification) GetMirBwDlBitRate() string`

GetMirBwDlBitRate returns the MirBwDlBitRate field if non-nil, zero value otherwise.

### GetMirBwDlBitRateOk

`func (o *M5QoSSpecification) GetMirBwDlBitRateOk() (*string, bool)`

GetMirBwDlBitRateOk returns a tuple with the MirBwDlBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwDlBitRate

`func (o *M5QoSSpecification) SetMirBwDlBitRate(v string)`

SetMirBwDlBitRate sets MirBwDlBitRate field to given value.


### GetMirBwUlBitRate

`func (o *M5QoSSpecification) GetMirBwUlBitRate() string`

GetMirBwUlBitRate returns the MirBwUlBitRate field if non-nil, zero value otherwise.

### GetMirBwUlBitRateOk

`func (o *M5QoSSpecification) GetMirBwUlBitRateOk() (*string, bool)`

GetMirBwUlBitRateOk returns a tuple with the MirBwUlBitRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirBwUlBitRate

`func (o *M5QoSSpecification) SetMirBwUlBitRate(v string)`

SetMirBwUlBitRate sets MirBwUlBitRate field to given value.


### GetDesLatency

`func (o *M5QoSSpecification) GetDesLatency() int32`

GetDesLatency returns the DesLatency field if non-nil, zero value otherwise.

### GetDesLatencyOk

`func (o *M5QoSSpecification) GetDesLatencyOk() (*int32, bool)`

GetDesLatencyOk returns a tuple with the DesLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesLatency

`func (o *M5QoSSpecification) SetDesLatency(v int32)`

SetDesLatency sets DesLatency field to given value.

### HasDesLatency

`func (o *M5QoSSpecification) HasDesLatency() bool`

HasDesLatency returns a boolean if a field has been set.

### GetDesLoss

`func (o *M5QoSSpecification) GetDesLoss() int32`

GetDesLoss returns the DesLoss field if non-nil, zero value otherwise.

### GetDesLossOk

`func (o *M5QoSSpecification) GetDesLossOk() (*int32, bool)`

GetDesLossOk returns a tuple with the DesLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesLoss

`func (o *M5QoSSpecification) SetDesLoss(v int32)`

SetDesLoss sets DesLoss field to given value.

### HasDesLoss

`func (o *M5QoSSpecification) HasDesLoss() bool`

HasDesLoss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


