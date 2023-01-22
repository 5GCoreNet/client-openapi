# UsedUnitContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**QuotaManagementIndicator** | Pointer to [**QuotaManagementIndicator**](QuotaManagementIndicator.md) |  | [optional] 
**Triggers** | Pointer to [**[]Trigger**](Trigger.md) |  | [optional] 
**TriggerTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**Time** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**TotalVolume** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**UplinkVolume** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**DownlinkVolume** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**ServiceSpecificUnits** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.  | [optional] 
**EventTimeStamps** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**LocalSequenceNumber** | **int32** |  | 
**PDUContainerInformation** | Pointer to [**PDUContainerInformation**](PDUContainerInformation.md) |  | [optional] 
**NSPAContainerInformation** | Pointer to [**NSPAContainerInformation**](NSPAContainerInformation.md) |  | [optional] 
**PC5ContainerInformation** | Pointer to [**PC5ContainerInformation**](PC5ContainerInformation.md) |  | [optional] 

## Methods

### NewUsedUnitContainer

`func NewUsedUnitContainer(localSequenceNumber int32, ) *UsedUnitContainer`

NewUsedUnitContainer instantiates a new UsedUnitContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsedUnitContainerWithDefaults

`func NewUsedUnitContainerWithDefaults() *UsedUnitContainer`

NewUsedUnitContainerWithDefaults instantiates a new UsedUnitContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *UsedUnitContainer) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *UsedUnitContainer) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *UsedUnitContainer) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *UsedUnitContainer) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetQuotaManagementIndicator

`func (o *UsedUnitContainer) GetQuotaManagementIndicator() QuotaManagementIndicator`

GetQuotaManagementIndicator returns the QuotaManagementIndicator field if non-nil, zero value otherwise.

### GetQuotaManagementIndicatorOk

`func (o *UsedUnitContainer) GetQuotaManagementIndicatorOk() (*QuotaManagementIndicator, bool)`

GetQuotaManagementIndicatorOk returns a tuple with the QuotaManagementIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaManagementIndicator

`func (o *UsedUnitContainer) SetQuotaManagementIndicator(v QuotaManagementIndicator)`

SetQuotaManagementIndicator sets QuotaManagementIndicator field to given value.

### HasQuotaManagementIndicator

`func (o *UsedUnitContainer) HasQuotaManagementIndicator() bool`

HasQuotaManagementIndicator returns a boolean if a field has been set.

### GetTriggers

`func (o *UsedUnitContainer) GetTriggers() []Trigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *UsedUnitContainer) GetTriggersOk() (*[]Trigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *UsedUnitContainer) SetTriggers(v []Trigger)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *UsedUnitContainer) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetTriggerTimestamp

`func (o *UsedUnitContainer) GetTriggerTimestamp() time.Time`

GetTriggerTimestamp returns the TriggerTimestamp field if non-nil, zero value otherwise.

### GetTriggerTimestampOk

`func (o *UsedUnitContainer) GetTriggerTimestampOk() (*time.Time, bool)`

GetTriggerTimestampOk returns a tuple with the TriggerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerTimestamp

`func (o *UsedUnitContainer) SetTriggerTimestamp(v time.Time)`

SetTriggerTimestamp sets TriggerTimestamp field to given value.

### HasTriggerTimestamp

`func (o *UsedUnitContainer) HasTriggerTimestamp() bool`

HasTriggerTimestamp returns a boolean if a field has been set.

### GetTime

`func (o *UsedUnitContainer) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UsedUnitContainer) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UsedUnitContainer) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *UsedUnitContainer) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTotalVolume

`func (o *UsedUnitContainer) GetTotalVolume() int32`

GetTotalVolume returns the TotalVolume field if non-nil, zero value otherwise.

### GetTotalVolumeOk

`func (o *UsedUnitContainer) GetTotalVolumeOk() (*int32, bool)`

GetTotalVolumeOk returns a tuple with the TotalVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVolume

`func (o *UsedUnitContainer) SetTotalVolume(v int32)`

SetTotalVolume sets TotalVolume field to given value.

### HasTotalVolume

`func (o *UsedUnitContainer) HasTotalVolume() bool`

HasTotalVolume returns a boolean if a field has been set.

### GetUplinkVolume

`func (o *UsedUnitContainer) GetUplinkVolume() int32`

GetUplinkVolume returns the UplinkVolume field if non-nil, zero value otherwise.

### GetUplinkVolumeOk

`func (o *UsedUnitContainer) GetUplinkVolumeOk() (*int32, bool)`

GetUplinkVolumeOk returns a tuple with the UplinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkVolume

`func (o *UsedUnitContainer) SetUplinkVolume(v int32)`

SetUplinkVolume sets UplinkVolume field to given value.

### HasUplinkVolume

`func (o *UsedUnitContainer) HasUplinkVolume() bool`

HasUplinkVolume returns a boolean if a field has been set.

### GetDownlinkVolume

`func (o *UsedUnitContainer) GetDownlinkVolume() int32`

GetDownlinkVolume returns the DownlinkVolume field if non-nil, zero value otherwise.

### GetDownlinkVolumeOk

`func (o *UsedUnitContainer) GetDownlinkVolumeOk() (*int32, bool)`

GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkVolume

`func (o *UsedUnitContainer) SetDownlinkVolume(v int32)`

SetDownlinkVolume sets DownlinkVolume field to given value.

### HasDownlinkVolume

`func (o *UsedUnitContainer) HasDownlinkVolume() bool`

HasDownlinkVolume returns a boolean if a field has been set.

### GetServiceSpecificUnits

`func (o *UsedUnitContainer) GetServiceSpecificUnits() int32`

GetServiceSpecificUnits returns the ServiceSpecificUnits field if non-nil, zero value otherwise.

### GetServiceSpecificUnitsOk

`func (o *UsedUnitContainer) GetServiceSpecificUnitsOk() (*int32, bool)`

GetServiceSpecificUnitsOk returns a tuple with the ServiceSpecificUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSpecificUnits

`func (o *UsedUnitContainer) SetServiceSpecificUnits(v int32)`

SetServiceSpecificUnits sets ServiceSpecificUnits field to given value.

### HasServiceSpecificUnits

`func (o *UsedUnitContainer) HasServiceSpecificUnits() bool`

HasServiceSpecificUnits returns a boolean if a field has been set.

### GetEventTimeStamps

`func (o *UsedUnitContainer) GetEventTimeStamps() []time.Time`

GetEventTimeStamps returns the EventTimeStamps field if non-nil, zero value otherwise.

### GetEventTimeStampsOk

`func (o *UsedUnitContainer) GetEventTimeStampsOk() (*[]time.Time, bool)`

GetEventTimeStampsOk returns a tuple with the EventTimeStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimeStamps

`func (o *UsedUnitContainer) SetEventTimeStamps(v []time.Time)`

SetEventTimeStamps sets EventTimeStamps field to given value.

### HasEventTimeStamps

`func (o *UsedUnitContainer) HasEventTimeStamps() bool`

HasEventTimeStamps returns a boolean if a field has been set.

### GetLocalSequenceNumber

`func (o *UsedUnitContainer) GetLocalSequenceNumber() int32`

GetLocalSequenceNumber returns the LocalSequenceNumber field if non-nil, zero value otherwise.

### GetLocalSequenceNumberOk

`func (o *UsedUnitContainer) GetLocalSequenceNumberOk() (*int32, bool)`

GetLocalSequenceNumberOk returns a tuple with the LocalSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSequenceNumber

`func (o *UsedUnitContainer) SetLocalSequenceNumber(v int32)`

SetLocalSequenceNumber sets LocalSequenceNumber field to given value.


### GetPDUContainerInformation

`func (o *UsedUnitContainer) GetPDUContainerInformation() PDUContainerInformation`

GetPDUContainerInformation returns the PDUContainerInformation field if non-nil, zero value otherwise.

### GetPDUContainerInformationOk

`func (o *UsedUnitContainer) GetPDUContainerInformationOk() (*PDUContainerInformation, bool)`

GetPDUContainerInformationOk returns a tuple with the PDUContainerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPDUContainerInformation

`func (o *UsedUnitContainer) SetPDUContainerInformation(v PDUContainerInformation)`

SetPDUContainerInformation sets PDUContainerInformation field to given value.

### HasPDUContainerInformation

`func (o *UsedUnitContainer) HasPDUContainerInformation() bool`

HasPDUContainerInformation returns a boolean if a field has been set.

### GetNSPAContainerInformation

`func (o *UsedUnitContainer) GetNSPAContainerInformation() NSPAContainerInformation`

GetNSPAContainerInformation returns the NSPAContainerInformation field if non-nil, zero value otherwise.

### GetNSPAContainerInformationOk

`func (o *UsedUnitContainer) GetNSPAContainerInformationOk() (*NSPAContainerInformation, bool)`

GetNSPAContainerInformationOk returns a tuple with the NSPAContainerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSPAContainerInformation

`func (o *UsedUnitContainer) SetNSPAContainerInformation(v NSPAContainerInformation)`

SetNSPAContainerInformation sets NSPAContainerInformation field to given value.

### HasNSPAContainerInformation

`func (o *UsedUnitContainer) HasNSPAContainerInformation() bool`

HasNSPAContainerInformation returns a boolean if a field has been set.

### GetPC5ContainerInformation

`func (o *UsedUnitContainer) GetPC5ContainerInformation() PC5ContainerInformation`

GetPC5ContainerInformation returns the PC5ContainerInformation field if non-nil, zero value otherwise.

### GetPC5ContainerInformationOk

`func (o *UsedUnitContainer) GetPC5ContainerInformationOk() (*PC5ContainerInformation, bool)`

GetPC5ContainerInformationOk returns a tuple with the PC5ContainerInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPC5ContainerInformation

`func (o *UsedUnitContainer) SetPC5ContainerInformation(v PC5ContainerInformation)`

SetPC5ContainerInformation sets PC5ContainerInformation field to given value.

### HasPC5ContainerInformation

`func (o *UsedUnitContainer) HasPC5ContainerInformation() bool`

HasPC5ContainerInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


