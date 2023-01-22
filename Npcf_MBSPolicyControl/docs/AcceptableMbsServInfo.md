# AcceptableMbsServInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccMbsServInfo** | Pointer to [**map[string]MbsMediaComp**](MbsMediaComp.md) |  | [optional] 
**AccMaxMbsBw** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewAcceptableMbsServInfo

`func NewAcceptableMbsServInfo() *AcceptableMbsServInfo`

NewAcceptableMbsServInfo instantiates a new AcceptableMbsServInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptableMbsServInfoWithDefaults

`func NewAcceptableMbsServInfoWithDefaults() *AcceptableMbsServInfo`

NewAcceptableMbsServInfoWithDefaults instantiates a new AcceptableMbsServInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccMbsServInfo

`func (o *AcceptableMbsServInfo) GetAccMbsServInfo() map[string]MbsMediaComp`

GetAccMbsServInfo returns the AccMbsServInfo field if non-nil, zero value otherwise.

### GetAccMbsServInfoOk

`func (o *AcceptableMbsServInfo) GetAccMbsServInfoOk() (*map[string]MbsMediaComp, bool)`

GetAccMbsServInfoOk returns a tuple with the AccMbsServInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccMbsServInfo

`func (o *AcceptableMbsServInfo) SetAccMbsServInfo(v map[string]MbsMediaComp)`

SetAccMbsServInfo sets AccMbsServInfo field to given value.

### HasAccMbsServInfo

`func (o *AcceptableMbsServInfo) HasAccMbsServInfo() bool`

HasAccMbsServInfo returns a boolean if a field has been set.

### GetAccMaxMbsBw

`func (o *AcceptableMbsServInfo) GetAccMaxMbsBw() string`

GetAccMaxMbsBw returns the AccMaxMbsBw field if non-nil, zero value otherwise.

### GetAccMaxMbsBwOk

`func (o *AcceptableMbsServInfo) GetAccMaxMbsBwOk() (*string, bool)`

GetAccMaxMbsBwOk returns a tuple with the AccMaxMbsBw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccMaxMbsBw

`func (o *AcceptableMbsServInfo) SetAccMaxMbsBw(v string)`

SetAccMaxMbsBw sets AccMaxMbsBw field to given value.

### HasAccMaxMbsBw

`func (o *AcceptableMbsServInfo) HasAccMaxMbsBw() bool`

HasAccMaxMbsBw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


