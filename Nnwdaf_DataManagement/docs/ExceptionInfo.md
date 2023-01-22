# ExceptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpTrafficFilter** | Pointer to [**FlowInfo**](FlowInfo.md) |  | [optional] 
**EthTrafficFilter** | Pointer to [**EthFlowDescription**](EthFlowDescription.md) |  | [optional] 
**Exceps** | [**[]Exception**](Exception.md) |  | 

## Methods

### NewExceptionInfo

`func NewExceptionInfo(exceps []Exception, ) *ExceptionInfo`

NewExceptionInfo instantiates a new ExceptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionInfoWithDefaults

`func NewExceptionInfoWithDefaults() *ExceptionInfo`

NewExceptionInfoWithDefaults instantiates a new ExceptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpTrafficFilter

`func (o *ExceptionInfo) GetIpTrafficFilter() FlowInfo`

GetIpTrafficFilter returns the IpTrafficFilter field if non-nil, zero value otherwise.

### GetIpTrafficFilterOk

`func (o *ExceptionInfo) GetIpTrafficFilterOk() (*FlowInfo, bool)`

GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTrafficFilter

`func (o *ExceptionInfo) SetIpTrafficFilter(v FlowInfo)`

SetIpTrafficFilter sets IpTrafficFilter field to given value.

### HasIpTrafficFilter

`func (o *ExceptionInfo) HasIpTrafficFilter() bool`

HasIpTrafficFilter returns a boolean if a field has been set.

### GetEthTrafficFilter

`func (o *ExceptionInfo) GetEthTrafficFilter() EthFlowDescription`

GetEthTrafficFilter returns the EthTrafficFilter field if non-nil, zero value otherwise.

### GetEthTrafficFilterOk

`func (o *ExceptionInfo) GetEthTrafficFilterOk() (*EthFlowDescription, bool)`

GetEthTrafficFilterOk returns a tuple with the EthTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthTrafficFilter

`func (o *ExceptionInfo) SetEthTrafficFilter(v EthFlowDescription)`

SetEthTrafficFilter sets EthTrafficFilter field to given value.

### HasEthTrafficFilter

`func (o *ExceptionInfo) HasEthTrafficFilter() bool`

HasEthTrafficFilter returns a boolean if a field has been set.

### GetExceps

`func (o *ExceptionInfo) GetExceps() []Exception`

GetExceps returns the Exceps field if non-nil, zero value otherwise.

### GetExcepsOk

`func (o *ExceptionInfo) GetExcepsOk() (*[]Exception, bool)`

GetExcepsOk returns a tuple with the Exceps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceps

`func (o *ExceptionInfo) SetExceps(v []Exception)`

SetExceps sets Exceps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


