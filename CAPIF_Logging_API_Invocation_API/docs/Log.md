# Log

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiId** | **string** | String identifying the API invoked. | 
**ApiName** | **string** | Name of the API which was invoked, it is set as {apiName} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122.  | 
**ApiVersion** | **string** | Version of the API which was invoked | 
**ResourceName** | **string** | Name of the specific resource invoked | 
**Uri** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**Protocol** | [**Protocol**](Protocol.md) |  | 
**Operation** | Pointer to [**Operation**](Operation.md) |  | [optional] 
**Result** | **string** | For HTTP protocol, it contains HTTP status code of the invocation | 
**InvocationTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**InvocationLatency** | Pointer to **int32** | Unsigned integer identifying a period of time in units of milliseconds. | [optional] 
**InputParameters** | Pointer to **interface{}** | List of input parameters. Can be any value - string, number, boolean, array or object.  | [optional] 
**OutputParameters** | Pointer to **interface{}** | List of output parameters. Can be any value - string, number, boolean, array or object.  | [optional] 
**SrcInterface** | Pointer to [**InterfaceDescription**](InterfaceDescription.md) |  | [optional] 
**DestInterface** | Pointer to [**InterfaceDescription**](InterfaceDescription.md) |  | [optional] 
**FwdInterface** | Pointer to **string** | It includes the node identifier (as defined in IETF RFC 7239 of all forwarding entities between the API invoker and the AEF, concatenated with comma and space, e.g. 192.0.2.43:80, unknown:_OBFport, 203.0.113.60  | [optional] 

## Methods

### NewLog

`func NewLog(apiId string, apiName string, apiVersion string, resourceName string, protocol Protocol, result string, ) *Log`

NewLog instantiates a new Log object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogWithDefaults

`func NewLogWithDefaults() *Log`

NewLogWithDefaults instantiates a new Log object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiId

`func (o *Log) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *Log) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *Log) SetApiId(v string)`

SetApiId sets ApiId field to given value.


### GetApiName

`func (o *Log) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *Log) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *Log) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetApiVersion

`func (o *Log) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Log) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Log) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetResourceName

`func (o *Log) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Log) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Log) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetUri

`func (o *Log) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Log) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Log) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Log) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetProtocol

`func (o *Log) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Log) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Log) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.


### GetOperation

`func (o *Log) GetOperation() Operation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Log) GetOperationOk() (*Operation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Log) SetOperation(v Operation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Log) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetResult

`func (o *Log) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Log) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Log) SetResult(v string)`

SetResult sets Result field to given value.


### GetInvocationTime

`func (o *Log) GetInvocationTime() time.Time`

GetInvocationTime returns the InvocationTime field if non-nil, zero value otherwise.

### GetInvocationTimeOk

`func (o *Log) GetInvocationTimeOk() (*time.Time, bool)`

GetInvocationTimeOk returns a tuple with the InvocationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationTime

`func (o *Log) SetInvocationTime(v time.Time)`

SetInvocationTime sets InvocationTime field to given value.

### HasInvocationTime

`func (o *Log) HasInvocationTime() bool`

HasInvocationTime returns a boolean if a field has been set.

### GetInvocationLatency

`func (o *Log) GetInvocationLatency() int32`

GetInvocationLatency returns the InvocationLatency field if non-nil, zero value otherwise.

### GetInvocationLatencyOk

`func (o *Log) GetInvocationLatencyOk() (*int32, bool)`

GetInvocationLatencyOk returns a tuple with the InvocationLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvocationLatency

`func (o *Log) SetInvocationLatency(v int32)`

SetInvocationLatency sets InvocationLatency field to given value.

### HasInvocationLatency

`func (o *Log) HasInvocationLatency() bool`

HasInvocationLatency returns a boolean if a field has been set.

### GetInputParameters

`func (o *Log) GetInputParameters() interface{}`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *Log) GetInputParametersOk() (*interface{}, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *Log) SetInputParameters(v interface{})`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *Log) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### SetInputParametersNil

`func (o *Log) SetInputParametersNil(b bool)`

 SetInputParametersNil sets the value for InputParameters to be an explicit nil

### UnsetInputParameters
`func (o *Log) UnsetInputParameters()`

UnsetInputParameters ensures that no value is present for InputParameters, not even an explicit nil
### GetOutputParameters

`func (o *Log) GetOutputParameters() interface{}`

GetOutputParameters returns the OutputParameters field if non-nil, zero value otherwise.

### GetOutputParametersOk

`func (o *Log) GetOutputParametersOk() (*interface{}, bool)`

GetOutputParametersOk returns a tuple with the OutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputParameters

`func (o *Log) SetOutputParameters(v interface{})`

SetOutputParameters sets OutputParameters field to given value.

### HasOutputParameters

`func (o *Log) HasOutputParameters() bool`

HasOutputParameters returns a boolean if a field has been set.

### SetOutputParametersNil

`func (o *Log) SetOutputParametersNil(b bool)`

 SetOutputParametersNil sets the value for OutputParameters to be an explicit nil

### UnsetOutputParameters
`func (o *Log) UnsetOutputParameters()`

UnsetOutputParameters ensures that no value is present for OutputParameters, not even an explicit nil
### GetSrcInterface

`func (o *Log) GetSrcInterface() InterfaceDescription`

GetSrcInterface returns the SrcInterface field if non-nil, zero value otherwise.

### GetSrcInterfaceOk

`func (o *Log) GetSrcInterfaceOk() (*InterfaceDescription, bool)`

GetSrcInterfaceOk returns a tuple with the SrcInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcInterface

`func (o *Log) SetSrcInterface(v InterfaceDescription)`

SetSrcInterface sets SrcInterface field to given value.

### HasSrcInterface

`func (o *Log) HasSrcInterface() bool`

HasSrcInterface returns a boolean if a field has been set.

### GetDestInterface

`func (o *Log) GetDestInterface() InterfaceDescription`

GetDestInterface returns the DestInterface field if non-nil, zero value otherwise.

### GetDestInterfaceOk

`func (o *Log) GetDestInterfaceOk() (*InterfaceDescription, bool)`

GetDestInterfaceOk returns a tuple with the DestInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestInterface

`func (o *Log) SetDestInterface(v InterfaceDescription)`

SetDestInterface sets DestInterface field to given value.

### HasDestInterface

`func (o *Log) HasDestInterface() bool`

HasDestInterface returns a boolean if a field has been set.

### GetFwdInterface

`func (o *Log) GetFwdInterface() string`

GetFwdInterface returns the FwdInterface field if non-nil, zero value otherwise.

### GetFwdInterfaceOk

`func (o *Log) GetFwdInterfaceOk() (*string, bool)`

GetFwdInterfaceOk returns a tuple with the FwdInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwdInterface

`func (o *Log) SetFwdInterface(v string)`

SetFwdInterface sets FwdInterface field to given value.

### HasFwdInterface

`func (o *Log) HasFwdInterface() bool`

HasFwdInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


