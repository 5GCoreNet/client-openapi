# ProseContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectDiscovery** | Pointer to [**UeAuth**](UeAuth.md) |  | [optional] 
**DirectComm** | Pointer to [**UeAuth**](UeAuth.md) |  | [optional] 
**L2Relay** | Pointer to [**UeAuth**](UeAuth.md) |  | [optional] 
**L3Relay** | Pointer to [**UeAuth**](UeAuth.md) |  | [optional] 
**L2Remote** | Pointer to [**UeAuth**](UeAuth.md) |  | [optional] 
**NrUePc5Ambr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**Pc5QoSPara** | Pointer to [**Pc5QoSPara**](Pc5QoSPara.md) |  | [optional] 

## Methods

### NewProseContext

`func NewProseContext() *ProseContext`

NewProseContext instantiates a new ProseContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProseContextWithDefaults

`func NewProseContextWithDefaults() *ProseContext`

NewProseContextWithDefaults instantiates a new ProseContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectDiscovery

`func (o *ProseContext) GetDirectDiscovery() UeAuth`

GetDirectDiscovery returns the DirectDiscovery field if non-nil, zero value otherwise.

### GetDirectDiscoveryOk

`func (o *ProseContext) GetDirectDiscoveryOk() (*UeAuth, bool)`

GetDirectDiscoveryOk returns a tuple with the DirectDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDiscovery

`func (o *ProseContext) SetDirectDiscovery(v UeAuth)`

SetDirectDiscovery sets DirectDiscovery field to given value.

### HasDirectDiscovery

`func (o *ProseContext) HasDirectDiscovery() bool`

HasDirectDiscovery returns a boolean if a field has been set.

### GetDirectComm

`func (o *ProseContext) GetDirectComm() UeAuth`

GetDirectComm returns the DirectComm field if non-nil, zero value otherwise.

### GetDirectCommOk

`func (o *ProseContext) GetDirectCommOk() (*UeAuth, bool)`

GetDirectCommOk returns a tuple with the DirectComm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectComm

`func (o *ProseContext) SetDirectComm(v UeAuth)`

SetDirectComm sets DirectComm field to given value.

### HasDirectComm

`func (o *ProseContext) HasDirectComm() bool`

HasDirectComm returns a boolean if a field has been set.

### GetL2Relay

`func (o *ProseContext) GetL2Relay() UeAuth`

GetL2Relay returns the L2Relay field if non-nil, zero value otherwise.

### GetL2RelayOk

`func (o *ProseContext) GetL2RelayOk() (*UeAuth, bool)`

GetL2RelayOk returns a tuple with the L2Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Relay

`func (o *ProseContext) SetL2Relay(v UeAuth)`

SetL2Relay sets L2Relay field to given value.

### HasL2Relay

`func (o *ProseContext) HasL2Relay() bool`

HasL2Relay returns a boolean if a field has been set.

### GetL3Relay

`func (o *ProseContext) GetL3Relay() UeAuth`

GetL3Relay returns the L3Relay field if non-nil, zero value otherwise.

### GetL3RelayOk

`func (o *ProseContext) GetL3RelayOk() (*UeAuth, bool)`

GetL3RelayOk returns a tuple with the L3Relay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL3Relay

`func (o *ProseContext) SetL3Relay(v UeAuth)`

SetL3Relay sets L3Relay field to given value.

### HasL3Relay

`func (o *ProseContext) HasL3Relay() bool`

HasL3Relay returns a boolean if a field has been set.

### GetL2Remote

`func (o *ProseContext) GetL2Remote() UeAuth`

GetL2Remote returns the L2Remote field if non-nil, zero value otherwise.

### GetL2RemoteOk

`func (o *ProseContext) GetL2RemoteOk() (*UeAuth, bool)`

GetL2RemoteOk returns a tuple with the L2Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL2Remote

`func (o *ProseContext) SetL2Remote(v UeAuth)`

SetL2Remote sets L2Remote field to given value.

### HasL2Remote

`func (o *ProseContext) HasL2Remote() bool`

HasL2Remote returns a boolean if a field has been set.

### GetNrUePc5Ambr

`func (o *ProseContext) GetNrUePc5Ambr() string`

GetNrUePc5Ambr returns the NrUePc5Ambr field if non-nil, zero value otherwise.

### GetNrUePc5AmbrOk

`func (o *ProseContext) GetNrUePc5AmbrOk() (*string, bool)`

GetNrUePc5AmbrOk returns a tuple with the NrUePc5Ambr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrUePc5Ambr

`func (o *ProseContext) SetNrUePc5Ambr(v string)`

SetNrUePc5Ambr sets NrUePc5Ambr field to given value.

### HasNrUePc5Ambr

`func (o *ProseContext) HasNrUePc5Ambr() bool`

HasNrUePc5Ambr returns a boolean if a field has been set.

### GetPc5QoSPara

`func (o *ProseContext) GetPc5QoSPara() Pc5QoSPara`

GetPc5QoSPara returns the Pc5QoSPara field if non-nil, zero value otherwise.

### GetPc5QoSParaOk

`func (o *ProseContext) GetPc5QoSParaOk() (*Pc5QoSPara, bool)`

GetPc5QoSParaOk returns a tuple with the Pc5QoSPara field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPc5QoSPara

`func (o *ProseContext) SetPc5QoSPara(v Pc5QoSPara)`

SetPc5QoSPara sets Pc5QoSPara field to given value.

### HasPc5QoSPara

`func (o *ProseContext) HasPc5QoSPara() bool`

HasPc5QoSPara returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


