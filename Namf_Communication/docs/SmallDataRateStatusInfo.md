# SmallDataRateStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**SmallDataRateStatus** | [**SmallDataRateStatus**](SmallDataRateStatus.md) |  | 

## Methods

### NewSmallDataRateStatusInfo

`func NewSmallDataRateStatusInfo(snssai Snssai, dnn string, smallDataRateStatus SmallDataRateStatus, ) *SmallDataRateStatusInfo`

NewSmallDataRateStatusInfo instantiates a new SmallDataRateStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmallDataRateStatusInfoWithDefaults

`func NewSmallDataRateStatusInfoWithDefaults() *SmallDataRateStatusInfo`

NewSmallDataRateStatusInfoWithDefaults instantiates a new SmallDataRateStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnssai

`func (o *SmallDataRateStatusInfo) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SmallDataRateStatusInfo) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SmallDataRateStatusInfo) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetDnn

`func (o *SmallDataRateStatusInfo) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *SmallDataRateStatusInfo) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *SmallDataRateStatusInfo) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetSmallDataRateStatus

`func (o *SmallDataRateStatusInfo) GetSmallDataRateStatus() SmallDataRateStatus`

GetSmallDataRateStatus returns the SmallDataRateStatus field if non-nil, zero value otherwise.

### GetSmallDataRateStatusOk

`func (o *SmallDataRateStatusInfo) GetSmallDataRateStatusOk() (*SmallDataRateStatus, bool)`

GetSmallDataRateStatusOk returns a tuple with the SmallDataRateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmallDataRateStatus

`func (o *SmallDataRateStatusInfo) SetSmallDataRateStatus(v SmallDataRateStatus)`

SetSmallDataRateStatus sets SmallDataRateStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


