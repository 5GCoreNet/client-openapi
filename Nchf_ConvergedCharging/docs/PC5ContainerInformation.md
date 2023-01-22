# PC5ContainerInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CoverageInfoList** | Pointer to [**[]CoverageInfo**](CoverageInfo.md) |  | [optional] 
**RadioParameterSetInfoList** | Pointer to [**[]RadioParameterSetInfo**](RadioParameterSetInfo.md) |  | [optional] 
**TransmitterInfoList** | Pointer to [**[]TransmitterInfo**](TransmitterInfo.md) |  | [optional] 
**TimeOfFirstTransmission** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**TimeOfFirstReception** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewPC5ContainerInformation

`func NewPC5ContainerInformation() *PC5ContainerInformation`

NewPC5ContainerInformation instantiates a new PC5ContainerInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPC5ContainerInformationWithDefaults

`func NewPC5ContainerInformationWithDefaults() *PC5ContainerInformation`

NewPC5ContainerInformationWithDefaults instantiates a new PC5ContainerInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoverageInfoList

`func (o *PC5ContainerInformation) GetCoverageInfoList() []CoverageInfo`

GetCoverageInfoList returns the CoverageInfoList field if non-nil, zero value otherwise.

### GetCoverageInfoListOk

`func (o *PC5ContainerInformation) GetCoverageInfoListOk() (*[]CoverageInfo, bool)`

GetCoverageInfoListOk returns a tuple with the CoverageInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageInfoList

`func (o *PC5ContainerInformation) SetCoverageInfoList(v []CoverageInfo)`

SetCoverageInfoList sets CoverageInfoList field to given value.

### HasCoverageInfoList

`func (o *PC5ContainerInformation) HasCoverageInfoList() bool`

HasCoverageInfoList returns a boolean if a field has been set.

### GetRadioParameterSetInfoList

`func (o *PC5ContainerInformation) GetRadioParameterSetInfoList() []RadioParameterSetInfo`

GetRadioParameterSetInfoList returns the RadioParameterSetInfoList field if non-nil, zero value otherwise.

### GetRadioParameterSetInfoListOk

`func (o *PC5ContainerInformation) GetRadioParameterSetInfoListOk() (*[]RadioParameterSetInfo, bool)`

GetRadioParameterSetInfoListOk returns a tuple with the RadioParameterSetInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioParameterSetInfoList

`func (o *PC5ContainerInformation) SetRadioParameterSetInfoList(v []RadioParameterSetInfo)`

SetRadioParameterSetInfoList sets RadioParameterSetInfoList field to given value.

### HasRadioParameterSetInfoList

`func (o *PC5ContainerInformation) HasRadioParameterSetInfoList() bool`

HasRadioParameterSetInfoList returns a boolean if a field has been set.

### GetTransmitterInfoList

`func (o *PC5ContainerInformation) GetTransmitterInfoList() []TransmitterInfo`

GetTransmitterInfoList returns the TransmitterInfoList field if non-nil, zero value otherwise.

### GetTransmitterInfoListOk

`func (o *PC5ContainerInformation) GetTransmitterInfoListOk() (*[]TransmitterInfo, bool)`

GetTransmitterInfoListOk returns a tuple with the TransmitterInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitterInfoList

`func (o *PC5ContainerInformation) SetTransmitterInfoList(v []TransmitterInfo)`

SetTransmitterInfoList sets TransmitterInfoList field to given value.

### HasTransmitterInfoList

`func (o *PC5ContainerInformation) HasTransmitterInfoList() bool`

HasTransmitterInfoList returns a boolean if a field has been set.

### GetTimeOfFirstTransmission

`func (o *PC5ContainerInformation) GetTimeOfFirstTransmission() time.Time`

GetTimeOfFirstTransmission returns the TimeOfFirstTransmission field if non-nil, zero value otherwise.

### GetTimeOfFirstTransmissionOk

`func (o *PC5ContainerInformation) GetTimeOfFirstTransmissionOk() (*time.Time, bool)`

GetTimeOfFirstTransmissionOk returns a tuple with the TimeOfFirstTransmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfFirstTransmission

`func (o *PC5ContainerInformation) SetTimeOfFirstTransmission(v time.Time)`

SetTimeOfFirstTransmission sets TimeOfFirstTransmission field to given value.

### HasTimeOfFirstTransmission

`func (o *PC5ContainerInformation) HasTimeOfFirstTransmission() bool`

HasTimeOfFirstTransmission returns a boolean if a field has been set.

### GetTimeOfFirstReception

`func (o *PC5ContainerInformation) GetTimeOfFirstReception() time.Time`

GetTimeOfFirstReception returns the TimeOfFirstReception field if non-nil, zero value otherwise.

### GetTimeOfFirstReceptionOk

`func (o *PC5ContainerInformation) GetTimeOfFirstReceptionOk() (*time.Time, bool)`

GetTimeOfFirstReceptionOk returns a tuple with the TimeOfFirstReception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOfFirstReception

`func (o *PC5ContainerInformation) SetTimeOfFirstReception(v time.Time)`

SetTimeOfFirstReception sets TimeOfFirstReception field to given value.

### HasTimeOfFirstReception

`func (o *PC5ContainerInformation) HasTimeOfFirstReception() bool`

HasTimeOfFirstReception returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


