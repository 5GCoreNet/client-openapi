# LMInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValTgtUe** | [**ValTargetUe**](ValTargetUe.md) |  | 
**LocInfo** | [**LocationInfo**](LocationInfo.md) |  | 
**TimeStamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**ValSvcId** | Pointer to **string** | Identity of the VAL service | [optional] 

## Methods

### NewLMInformation

`func NewLMInformation(valTgtUe ValTargetUe, locInfo LocationInfo, ) *LMInformation`

NewLMInformation instantiates a new LMInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLMInformationWithDefaults

`func NewLMInformationWithDefaults() *LMInformation`

NewLMInformationWithDefaults instantiates a new LMInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValTgtUe

`func (o *LMInformation) GetValTgtUe() ValTargetUe`

GetValTgtUe returns the ValTgtUe field if non-nil, zero value otherwise.

### GetValTgtUeOk

`func (o *LMInformation) GetValTgtUeOk() (*ValTargetUe, bool)`

GetValTgtUeOk returns a tuple with the ValTgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUe

`func (o *LMInformation) SetValTgtUe(v ValTargetUe)`

SetValTgtUe sets ValTgtUe field to given value.


### GetLocInfo

`func (o *LMInformation) GetLocInfo() LocationInfo`

GetLocInfo returns the LocInfo field if non-nil, zero value otherwise.

### GetLocInfoOk

`func (o *LMInformation) GetLocInfoOk() (*LocationInfo, bool)`

GetLocInfoOk returns a tuple with the LocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfo

`func (o *LMInformation) SetLocInfo(v LocationInfo)`

SetLocInfo sets LocInfo field to given value.


### GetTimeStamp

`func (o *LMInformation) GetTimeStamp() time.Time`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *LMInformation) GetTimeStampOk() (*time.Time, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *LMInformation) SetTimeStamp(v time.Time)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *LMInformation) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetValSvcId

`func (o *LMInformation) GetValSvcId() string`

GetValSvcId returns the ValSvcId field if non-nil, zero value otherwise.

### GetValSvcIdOk

`func (o *LMInformation) GetValSvcIdOk() (*string, bool)`

GetValSvcIdOk returns a tuple with the ValSvcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValSvcId

`func (o *LMInformation) SetValSvcId(v string)`

SetValSvcId sets ValSvcId field to given value.

### HasValSvcId

`func (o *LMInformation) HasValSvcId() bool`

HasValSvcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


