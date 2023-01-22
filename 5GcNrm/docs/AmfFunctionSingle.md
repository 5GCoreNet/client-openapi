# AmfFunctionSingle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** |  | 
**ObjectClass** | Pointer to **string** |  | [optional] 
**ObjectInstance** | Pointer to **string** |  | [optional] 
**VsDataContainer** | Pointer to [**[]VsDataContainerSingle**](VsDataContainerSingle.md) |  | [optional] 
**Attributes** | Pointer to [**ManagedFunctionAttr**](ManagedFunction-Attr.md) |  | [optional] 
**PerfMetricJob** | Pointer to [**[]PerfMetricJobSingle**](PerfMetricJobSingle.md) |  | [optional] 
**ThresholdMonitor** | Pointer to [**[]ThresholdMonitorSingle**](ThresholdMonitorSingle.md) |  | [optional] 
**ManagedNFService** | Pointer to [**[]ManagedNFServiceSingle**](ManagedNFServiceSingle.md) |  | [optional] 
**TraceJob** | Pointer to [**[]TraceJobSingle**](TraceJobSingle.md) |  | [optional] 
**EPN2** | Pointer to [**[]EPN2Single**](EPN2Single.md) |  | [optional] 
**EPN8** | Pointer to [**[]EPN8Single**](EPN8Single.md) |  | [optional] 
**EPN11** | Pointer to [**[]EPN11Single**](EPN11Single.md) |  | [optional] 
**EPN12** | Pointer to [**[]EPN12Single**](EPN12Single.md) |  | [optional] 
**EPN14** | Pointer to [**[]EPN14Single**](EPN14Single.md) |  | [optional] 
**EPN15** | Pointer to [**[]EPN15Single**](EPN15Single.md) |  | [optional] 
**EPN17** | Pointer to [**[]EPN17Single**](EPN17Single.md) |  | [optional] 
**EPN20** | Pointer to [**[]EPN20Single**](EPN20Single.md) |  | [optional] 
**EPN22** | Pointer to [**[]EPN22Single**](EPN22Single.md) |  | [optional] 
**EPN26** | Pointer to [**[]EPN26Single**](EPN26Single.md) |  | [optional] 
**EP_NLS** | Pointer to [**[]EPNLSSingle**](EPNLSSingle.md) |  | [optional] 
**EP_NLG** | Pointer to [**[]EPNLGSingle**](EPNLGSingle.md) |  | [optional] 

## Methods

### NewAmfFunctionSingle

`func NewAmfFunctionSingle(id NullableString, ) *AmfFunctionSingle`

NewAmfFunctionSingle instantiates a new AmfFunctionSingle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmfFunctionSingleWithDefaults

`func NewAmfFunctionSingleWithDefaults() *AmfFunctionSingle`

NewAmfFunctionSingleWithDefaults instantiates a new AmfFunctionSingle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AmfFunctionSingle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AmfFunctionSingle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AmfFunctionSingle) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *AmfFunctionSingle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AmfFunctionSingle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetObjectClass

`func (o *AmfFunctionSingle) GetObjectClass() string`

GetObjectClass returns the ObjectClass field if non-nil, zero value otherwise.

### GetObjectClassOk

`func (o *AmfFunctionSingle) GetObjectClassOk() (*string, bool)`

GetObjectClassOk returns a tuple with the ObjectClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectClass

`func (o *AmfFunctionSingle) SetObjectClass(v string)`

SetObjectClass sets ObjectClass field to given value.

### HasObjectClass

`func (o *AmfFunctionSingle) HasObjectClass() bool`

HasObjectClass returns a boolean if a field has been set.

### GetObjectInstance

`func (o *AmfFunctionSingle) GetObjectInstance() string`

GetObjectInstance returns the ObjectInstance field if non-nil, zero value otherwise.

### GetObjectInstanceOk

`func (o *AmfFunctionSingle) GetObjectInstanceOk() (*string, bool)`

GetObjectInstanceOk returns a tuple with the ObjectInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectInstance

`func (o *AmfFunctionSingle) SetObjectInstance(v string)`

SetObjectInstance sets ObjectInstance field to given value.

### HasObjectInstance

`func (o *AmfFunctionSingle) HasObjectInstance() bool`

HasObjectInstance returns a boolean if a field has been set.

### GetVsDataContainer

`func (o *AmfFunctionSingle) GetVsDataContainer() []VsDataContainerSingle`

GetVsDataContainer returns the VsDataContainer field if non-nil, zero value otherwise.

### GetVsDataContainerOk

`func (o *AmfFunctionSingle) GetVsDataContainerOk() (*[]VsDataContainerSingle, bool)`

GetVsDataContainerOk returns a tuple with the VsDataContainer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsDataContainer

`func (o *AmfFunctionSingle) SetVsDataContainer(v []VsDataContainerSingle)`

SetVsDataContainer sets VsDataContainer field to given value.

### HasVsDataContainer

`func (o *AmfFunctionSingle) HasVsDataContainer() bool`

HasVsDataContainer returns a boolean if a field has been set.

### GetAttributes

`func (o *AmfFunctionSingle) GetAttributes() ManagedFunctionAttr`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AmfFunctionSingle) GetAttributesOk() (*ManagedFunctionAttr, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AmfFunctionSingle) SetAttributes(v ManagedFunctionAttr)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AmfFunctionSingle) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetPerfMetricJob

`func (o *AmfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle`

GetPerfMetricJob returns the PerfMetricJob field if non-nil, zero value otherwise.

### GetPerfMetricJobOk

`func (o *AmfFunctionSingle) GetPerfMetricJobOk() (*[]PerfMetricJobSingle, bool)`

GetPerfMetricJobOk returns a tuple with the PerfMetricJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfMetricJob

`func (o *AmfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle)`

SetPerfMetricJob sets PerfMetricJob field to given value.

### HasPerfMetricJob

`func (o *AmfFunctionSingle) HasPerfMetricJob() bool`

HasPerfMetricJob returns a boolean if a field has been set.

### GetThresholdMonitor

`func (o *AmfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle`

GetThresholdMonitor returns the ThresholdMonitor field if non-nil, zero value otherwise.

### GetThresholdMonitorOk

`func (o *AmfFunctionSingle) GetThresholdMonitorOk() (*[]ThresholdMonitorSingle, bool)`

GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdMonitor

`func (o *AmfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle)`

SetThresholdMonitor sets ThresholdMonitor field to given value.

### HasThresholdMonitor

`func (o *AmfFunctionSingle) HasThresholdMonitor() bool`

HasThresholdMonitor returns a boolean if a field has been set.

### GetManagedNFService

`func (o *AmfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle`

GetManagedNFService returns the ManagedNFService field if non-nil, zero value otherwise.

### GetManagedNFServiceOk

`func (o *AmfFunctionSingle) GetManagedNFServiceOk() (*[]ManagedNFServiceSingle, bool)`

GetManagedNFServiceOk returns a tuple with the ManagedNFService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedNFService

`func (o *AmfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle)`

SetManagedNFService sets ManagedNFService field to given value.

### HasManagedNFService

`func (o *AmfFunctionSingle) HasManagedNFService() bool`

HasManagedNFService returns a boolean if a field has been set.

### GetTraceJob

`func (o *AmfFunctionSingle) GetTraceJob() []TraceJobSingle`

GetTraceJob returns the TraceJob field if non-nil, zero value otherwise.

### GetTraceJobOk

`func (o *AmfFunctionSingle) GetTraceJobOk() (*[]TraceJobSingle, bool)`

GetTraceJobOk returns a tuple with the TraceJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceJob

`func (o *AmfFunctionSingle) SetTraceJob(v []TraceJobSingle)`

SetTraceJob sets TraceJob field to given value.

### HasTraceJob

`func (o *AmfFunctionSingle) HasTraceJob() bool`

HasTraceJob returns a boolean if a field has been set.

### GetEPN2

`func (o *AmfFunctionSingle) GetEPN2() []EPN2Single`

GetEPN2 returns the EPN2 field if non-nil, zero value otherwise.

### GetEPN2Ok

`func (o *AmfFunctionSingle) GetEPN2Ok() (*[]EPN2Single, bool)`

GetEPN2Ok returns a tuple with the EPN2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN2

`func (o *AmfFunctionSingle) SetEPN2(v []EPN2Single)`

SetEPN2 sets EPN2 field to given value.

### HasEPN2

`func (o *AmfFunctionSingle) HasEPN2() bool`

HasEPN2 returns a boolean if a field has been set.

### GetEPN8

`func (o *AmfFunctionSingle) GetEPN8() []EPN8Single`

GetEPN8 returns the EPN8 field if non-nil, zero value otherwise.

### GetEPN8Ok

`func (o *AmfFunctionSingle) GetEPN8Ok() (*[]EPN8Single, bool)`

GetEPN8Ok returns a tuple with the EPN8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN8

`func (o *AmfFunctionSingle) SetEPN8(v []EPN8Single)`

SetEPN8 sets EPN8 field to given value.

### HasEPN8

`func (o *AmfFunctionSingle) HasEPN8() bool`

HasEPN8 returns a boolean if a field has been set.

### GetEPN11

`func (o *AmfFunctionSingle) GetEPN11() []EPN11Single`

GetEPN11 returns the EPN11 field if non-nil, zero value otherwise.

### GetEPN11Ok

`func (o *AmfFunctionSingle) GetEPN11Ok() (*[]EPN11Single, bool)`

GetEPN11Ok returns a tuple with the EPN11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN11

`func (o *AmfFunctionSingle) SetEPN11(v []EPN11Single)`

SetEPN11 sets EPN11 field to given value.

### HasEPN11

`func (o *AmfFunctionSingle) HasEPN11() bool`

HasEPN11 returns a boolean if a field has been set.

### GetEPN12

`func (o *AmfFunctionSingle) GetEPN12() []EPN12Single`

GetEPN12 returns the EPN12 field if non-nil, zero value otherwise.

### GetEPN12Ok

`func (o *AmfFunctionSingle) GetEPN12Ok() (*[]EPN12Single, bool)`

GetEPN12Ok returns a tuple with the EPN12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN12

`func (o *AmfFunctionSingle) SetEPN12(v []EPN12Single)`

SetEPN12 sets EPN12 field to given value.

### HasEPN12

`func (o *AmfFunctionSingle) HasEPN12() bool`

HasEPN12 returns a boolean if a field has been set.

### GetEPN14

`func (o *AmfFunctionSingle) GetEPN14() []EPN14Single`

GetEPN14 returns the EPN14 field if non-nil, zero value otherwise.

### GetEPN14Ok

`func (o *AmfFunctionSingle) GetEPN14Ok() (*[]EPN14Single, bool)`

GetEPN14Ok returns a tuple with the EPN14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN14

`func (o *AmfFunctionSingle) SetEPN14(v []EPN14Single)`

SetEPN14 sets EPN14 field to given value.

### HasEPN14

`func (o *AmfFunctionSingle) HasEPN14() bool`

HasEPN14 returns a boolean if a field has been set.

### GetEPN15

`func (o *AmfFunctionSingle) GetEPN15() []EPN15Single`

GetEPN15 returns the EPN15 field if non-nil, zero value otherwise.

### GetEPN15Ok

`func (o *AmfFunctionSingle) GetEPN15Ok() (*[]EPN15Single, bool)`

GetEPN15Ok returns a tuple with the EPN15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN15

`func (o *AmfFunctionSingle) SetEPN15(v []EPN15Single)`

SetEPN15 sets EPN15 field to given value.

### HasEPN15

`func (o *AmfFunctionSingle) HasEPN15() bool`

HasEPN15 returns a boolean if a field has been set.

### GetEPN17

`func (o *AmfFunctionSingle) GetEPN17() []EPN17Single`

GetEPN17 returns the EPN17 field if non-nil, zero value otherwise.

### GetEPN17Ok

`func (o *AmfFunctionSingle) GetEPN17Ok() (*[]EPN17Single, bool)`

GetEPN17Ok returns a tuple with the EPN17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN17

`func (o *AmfFunctionSingle) SetEPN17(v []EPN17Single)`

SetEPN17 sets EPN17 field to given value.

### HasEPN17

`func (o *AmfFunctionSingle) HasEPN17() bool`

HasEPN17 returns a boolean if a field has been set.

### GetEPN20

`func (o *AmfFunctionSingle) GetEPN20() []EPN20Single`

GetEPN20 returns the EPN20 field if non-nil, zero value otherwise.

### GetEPN20Ok

`func (o *AmfFunctionSingle) GetEPN20Ok() (*[]EPN20Single, bool)`

GetEPN20Ok returns a tuple with the EPN20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN20

`func (o *AmfFunctionSingle) SetEPN20(v []EPN20Single)`

SetEPN20 sets EPN20 field to given value.

### HasEPN20

`func (o *AmfFunctionSingle) HasEPN20() bool`

HasEPN20 returns a boolean if a field has been set.

### GetEPN22

`func (o *AmfFunctionSingle) GetEPN22() []EPN22Single`

GetEPN22 returns the EPN22 field if non-nil, zero value otherwise.

### GetEPN22Ok

`func (o *AmfFunctionSingle) GetEPN22Ok() (*[]EPN22Single, bool)`

GetEPN22Ok returns a tuple with the EPN22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN22

`func (o *AmfFunctionSingle) SetEPN22(v []EPN22Single)`

SetEPN22 sets EPN22 field to given value.

### HasEPN22

`func (o *AmfFunctionSingle) HasEPN22() bool`

HasEPN22 returns a boolean if a field has been set.

### GetEPN26

`func (o *AmfFunctionSingle) GetEPN26() []EPN26Single`

GetEPN26 returns the EPN26 field if non-nil, zero value otherwise.

### GetEPN26Ok

`func (o *AmfFunctionSingle) GetEPN26Ok() (*[]EPN26Single, bool)`

GetEPN26Ok returns a tuple with the EPN26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPN26

`func (o *AmfFunctionSingle) SetEPN26(v []EPN26Single)`

SetEPN26 sets EPN26 field to given value.

### HasEPN26

`func (o *AmfFunctionSingle) HasEPN26() bool`

HasEPN26 returns a boolean if a field has been set.

### GetEP_NLS

`func (o *AmfFunctionSingle) GetEP_NLS() []EPNLSSingle`

GetEP_NLS returns the EP_NLS field if non-nil, zero value otherwise.

### GetEP_NLSOk

`func (o *AmfFunctionSingle) GetEP_NLSOk() (*[]EPNLSSingle, bool)`

GetEP_NLSOk returns a tuple with the EP_NLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEP_NLS

`func (o *AmfFunctionSingle) SetEP_NLS(v []EPNLSSingle)`

SetEP_NLS sets EP_NLS field to given value.

### HasEP_NLS

`func (o *AmfFunctionSingle) HasEP_NLS() bool`

HasEP_NLS returns a boolean if a field has been set.

### GetEP_NLG

`func (o *AmfFunctionSingle) GetEP_NLG() []EPNLGSingle`

GetEP_NLG returns the EP_NLG field if non-nil, zero value otherwise.

### GetEP_NLGOk

`func (o *AmfFunctionSingle) GetEP_NLGOk() (*[]EPNLGSingle, bool)`

GetEP_NLGOk returns a tuple with the EP_NLG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEP_NLG

`func (o *AmfFunctionSingle) SetEP_NLG(v []EPNLGSingle)`

SetEP_NLG sets EP_NLG field to given value.

### HasEP_NLG

`func (o *AmfFunctionSingle) HasEP_NLG() bool`

HasEP_NLG returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


