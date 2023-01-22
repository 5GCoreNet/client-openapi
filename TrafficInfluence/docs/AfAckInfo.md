# AfAckInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfTransId** | Pointer to **string** |  | [optional] 
**AckResult** | [**AfResultInfo**](AfResultInfo.md) |  | 
**Gpsi** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 

## Methods

### NewAfAckInfo

`func NewAfAckInfo(ackResult AfResultInfo, ) *AfAckInfo`

NewAfAckInfo instantiates a new AfAckInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAfAckInfoWithDefaults

`func NewAfAckInfoWithDefaults() *AfAckInfo`

NewAfAckInfoWithDefaults instantiates a new AfAckInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfTransId

`func (o *AfAckInfo) GetAfTransId() string`

GetAfTransId returns the AfTransId field if non-nil, zero value otherwise.

### GetAfTransIdOk

`func (o *AfAckInfo) GetAfTransIdOk() (*string, bool)`

GetAfTransIdOk returns a tuple with the AfTransId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfTransId

`func (o *AfAckInfo) SetAfTransId(v string)`

SetAfTransId sets AfTransId field to given value.

### HasAfTransId

`func (o *AfAckInfo) HasAfTransId() bool`

HasAfTransId returns a boolean if a field has been set.

### GetAckResult

`func (o *AfAckInfo) GetAckResult() AfResultInfo`

GetAckResult returns the AckResult field if non-nil, zero value otherwise.

### GetAckResultOk

`func (o *AfAckInfo) GetAckResultOk() (*AfResultInfo, bool)`

GetAckResultOk returns a tuple with the AckResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckResult

`func (o *AfAckInfo) SetAckResult(v AfResultInfo)`

SetAckResult sets AckResult field to given value.


### GetGpsi

`func (o *AfAckInfo) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *AfAckInfo) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *AfAckInfo) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.

### HasGpsi

`func (o *AfAckInfo) HasGpsi() bool`

HasGpsi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


