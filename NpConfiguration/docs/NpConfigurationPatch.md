# NpConfigurationPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumLatency** | Pointer to **NullableInt32** | Unsigned integer identifying a period of time in units of seconds with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**MaximumResponseTime** | Pointer to **NullableInt32** | Unsigned integer identifying a period of time in units of seconds with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**SuggestedNumberOfDlPackets** | Pointer to **NullableInt32** | This parameter may be included to identify the number of packets that the serving gateway shall buffer in case that the UE is not reachable. | [optional] 
**GroupReportGuardTime** | Pointer to **NullableInt32** | Unsigned integer identifying a period of time in units of seconds with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**ValidityTime** | Pointer to **NullableTime** | string with format \&quot;date-time\&quot; as defined in OpenAPI with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**NotificationDestination** | Pointer to **string** | string formatted according to IETF RFC 3986 identifying a referenced resource. | [optional] 

## Methods

### NewNpConfigurationPatch

`func NewNpConfigurationPatch() *NpConfigurationPatch`

NewNpConfigurationPatch instantiates a new NpConfigurationPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNpConfigurationPatchWithDefaults

`func NewNpConfigurationPatchWithDefaults() *NpConfigurationPatch`

NewNpConfigurationPatchWithDefaults instantiates a new NpConfigurationPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumLatency

`func (o *NpConfigurationPatch) GetMaximumLatency() int32`

GetMaximumLatency returns the MaximumLatency field if non-nil, zero value otherwise.

### GetMaximumLatencyOk

`func (o *NpConfigurationPatch) GetMaximumLatencyOk() (*int32, bool)`

GetMaximumLatencyOk returns a tuple with the MaximumLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumLatency

`func (o *NpConfigurationPatch) SetMaximumLatency(v int32)`

SetMaximumLatency sets MaximumLatency field to given value.

### HasMaximumLatency

`func (o *NpConfigurationPatch) HasMaximumLatency() bool`

HasMaximumLatency returns a boolean if a field has been set.

### SetMaximumLatencyNil

`func (o *NpConfigurationPatch) SetMaximumLatencyNil(b bool)`

 SetMaximumLatencyNil sets the value for MaximumLatency to be an explicit nil

### UnsetMaximumLatency
`func (o *NpConfigurationPatch) UnsetMaximumLatency()`

UnsetMaximumLatency ensures that no value is present for MaximumLatency, not even an explicit nil
### GetMaximumResponseTime

`func (o *NpConfigurationPatch) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *NpConfigurationPatch) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *NpConfigurationPatch) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.

### HasMaximumResponseTime

`func (o *NpConfigurationPatch) HasMaximumResponseTime() bool`

HasMaximumResponseTime returns a boolean if a field has been set.

### SetMaximumResponseTimeNil

`func (o *NpConfigurationPatch) SetMaximumResponseTimeNil(b bool)`

 SetMaximumResponseTimeNil sets the value for MaximumResponseTime to be an explicit nil

### UnsetMaximumResponseTime
`func (o *NpConfigurationPatch) UnsetMaximumResponseTime()`

UnsetMaximumResponseTime ensures that no value is present for MaximumResponseTime, not even an explicit nil
### GetSuggestedNumberOfDlPackets

`func (o *NpConfigurationPatch) GetSuggestedNumberOfDlPackets() int32`

GetSuggestedNumberOfDlPackets returns the SuggestedNumberOfDlPackets field if non-nil, zero value otherwise.

### GetSuggestedNumberOfDlPacketsOk

`func (o *NpConfigurationPatch) GetSuggestedNumberOfDlPacketsOk() (*int32, bool)`

GetSuggestedNumberOfDlPacketsOk returns a tuple with the SuggestedNumberOfDlPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedNumberOfDlPackets

`func (o *NpConfigurationPatch) SetSuggestedNumberOfDlPackets(v int32)`

SetSuggestedNumberOfDlPackets sets SuggestedNumberOfDlPackets field to given value.

### HasSuggestedNumberOfDlPackets

`func (o *NpConfigurationPatch) HasSuggestedNumberOfDlPackets() bool`

HasSuggestedNumberOfDlPackets returns a boolean if a field has been set.

### SetSuggestedNumberOfDlPacketsNil

`func (o *NpConfigurationPatch) SetSuggestedNumberOfDlPacketsNil(b bool)`

 SetSuggestedNumberOfDlPacketsNil sets the value for SuggestedNumberOfDlPackets to be an explicit nil

### UnsetSuggestedNumberOfDlPackets
`func (o *NpConfigurationPatch) UnsetSuggestedNumberOfDlPackets()`

UnsetSuggestedNumberOfDlPackets ensures that no value is present for SuggestedNumberOfDlPackets, not even an explicit nil
### GetGroupReportGuardTime

`func (o *NpConfigurationPatch) GetGroupReportGuardTime() int32`

GetGroupReportGuardTime returns the GroupReportGuardTime field if non-nil, zero value otherwise.

### GetGroupReportGuardTimeOk

`func (o *NpConfigurationPatch) GetGroupReportGuardTimeOk() (*int32, bool)`

GetGroupReportGuardTimeOk returns a tuple with the GroupReportGuardTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupReportGuardTime

`func (o *NpConfigurationPatch) SetGroupReportGuardTime(v int32)`

SetGroupReportGuardTime sets GroupReportGuardTime field to given value.

### HasGroupReportGuardTime

`func (o *NpConfigurationPatch) HasGroupReportGuardTime() bool`

HasGroupReportGuardTime returns a boolean if a field has been set.

### SetGroupReportGuardTimeNil

`func (o *NpConfigurationPatch) SetGroupReportGuardTimeNil(b bool)`

 SetGroupReportGuardTimeNil sets the value for GroupReportGuardTime to be an explicit nil

### UnsetGroupReportGuardTime
`func (o *NpConfigurationPatch) UnsetGroupReportGuardTime()`

UnsetGroupReportGuardTime ensures that no value is present for GroupReportGuardTime, not even an explicit nil
### GetValidityTime

`func (o *NpConfigurationPatch) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *NpConfigurationPatch) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *NpConfigurationPatch) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.

### HasValidityTime

`func (o *NpConfigurationPatch) HasValidityTime() bool`

HasValidityTime returns a boolean if a field has been set.

### SetValidityTimeNil

`func (o *NpConfigurationPatch) SetValidityTimeNil(b bool)`

 SetValidityTimeNil sets the value for ValidityTime to be an explicit nil

### UnsetValidityTime
`func (o *NpConfigurationPatch) UnsetValidityTime()`

UnsetValidityTime ensures that no value is present for ValidityTime, not even an explicit nil
### GetNotificationDestination

`func (o *NpConfigurationPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *NpConfigurationPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *NpConfigurationPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *NpConfigurationPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


