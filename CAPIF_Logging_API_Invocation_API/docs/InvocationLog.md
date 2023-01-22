# InvocationLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AefId** | **string** | Identity information of the API exposing function requesting logging of service API invocations  | 
**ApiInvokerId** | **string** | Identity of the API invoker which invoked the service API | 
**Logs** | [**[]Log**](Log.md) | Service API invocation log | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewInvocationLog

`func NewInvocationLog(aefId string, apiInvokerId string, logs []Log, ) *InvocationLog`

NewInvocationLog instantiates a new InvocationLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvocationLogWithDefaults

`func NewInvocationLogWithDefaults() *InvocationLog`

NewInvocationLogWithDefaults instantiates a new InvocationLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAefId

`func (o *InvocationLog) GetAefId() string`

GetAefId returns the AefId field if non-nil, zero value otherwise.

### GetAefIdOk

`func (o *InvocationLog) GetAefIdOk() (*string, bool)`

GetAefIdOk returns a tuple with the AefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAefId

`func (o *InvocationLog) SetAefId(v string)`

SetAefId sets AefId field to given value.


### GetApiInvokerId

`func (o *InvocationLog) GetApiInvokerId() string`

GetApiInvokerId returns the ApiInvokerId field if non-nil, zero value otherwise.

### GetApiInvokerIdOk

`func (o *InvocationLog) GetApiInvokerIdOk() (*string, bool)`

GetApiInvokerIdOk returns a tuple with the ApiInvokerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerId

`func (o *InvocationLog) SetApiInvokerId(v string)`

SetApiInvokerId sets ApiInvokerId field to given value.


### GetLogs

`func (o *InvocationLog) GetLogs() []Log`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *InvocationLog) GetLogsOk() (*[]Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *InvocationLog) SetLogs(v []Log)`

SetLogs sets Logs field to given value.


### GetSupportedFeatures

`func (o *InvocationLog) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *InvocationLog) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *InvocationLog) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *InvocationLog) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


