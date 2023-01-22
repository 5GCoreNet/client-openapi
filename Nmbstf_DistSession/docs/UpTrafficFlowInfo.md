# UpTrafficFlowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestIpAddr** | [**IpAddr**](IpAddr.md) |  | 
**PortNumber** | Pointer to **int32** | Unsigned Integer, i.e. only value 0 and integers above 0 are permissible. | [optional] 

## Methods

### NewUpTrafficFlowInfo

`func NewUpTrafficFlowInfo(destIpAddr IpAddr, ) *UpTrafficFlowInfo`

NewUpTrafficFlowInfo instantiates a new UpTrafficFlowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpTrafficFlowInfoWithDefaults

`func NewUpTrafficFlowInfoWithDefaults() *UpTrafficFlowInfo`

NewUpTrafficFlowInfoWithDefaults instantiates a new UpTrafficFlowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestIpAddr

`func (o *UpTrafficFlowInfo) GetDestIpAddr() IpAddr`

GetDestIpAddr returns the DestIpAddr field if non-nil, zero value otherwise.

### GetDestIpAddrOk

`func (o *UpTrafficFlowInfo) GetDestIpAddrOk() (*IpAddr, bool)`

GetDestIpAddrOk returns a tuple with the DestIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIpAddr

`func (o *UpTrafficFlowInfo) SetDestIpAddr(v IpAddr)`

SetDestIpAddr sets DestIpAddr field to given value.


### GetPortNumber

`func (o *UpTrafficFlowInfo) GetPortNumber() int32`

GetPortNumber returns the PortNumber field if non-nil, zero value otherwise.

### GetPortNumberOk

`func (o *UpTrafficFlowInfo) GetPortNumberOk() (*int32, bool)`

GetPortNumberOk returns a tuple with the PortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNumber

`func (o *UpTrafficFlowInfo) SetPortNumber(v int32)`

SetPortNumber sets PortNumber field to given value.

### HasPortNumber

`func (o *UpTrafficFlowInfo) HasPortNumber() bool`

HasPortNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


