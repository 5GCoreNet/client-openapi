# UePerLocationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UeCount** | **int32** | Identifies the number of UEs. | 
**ExternalIds** | Pointer to **[]string** | Each element uniquely identifies a user. | [optional] 
**Msisdns** | Pointer to **[]string** | Each element identifies the MS internal PSTN/ISDN number allocated for a UE. | [optional] 
**ServLevelDevIds** | Pointer to **[]string** | Each element uniquely identifies a UAV. | [optional] 

## Methods

### NewUePerLocationReport

`func NewUePerLocationReport(ueCount int32, ) *UePerLocationReport`

NewUePerLocationReport instantiates a new UePerLocationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUePerLocationReportWithDefaults

`func NewUePerLocationReportWithDefaults() *UePerLocationReport`

NewUePerLocationReportWithDefaults instantiates a new UePerLocationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUeCount

`func (o *UePerLocationReport) GetUeCount() int32`

GetUeCount returns the UeCount field if non-nil, zero value otherwise.

### GetUeCountOk

`func (o *UePerLocationReport) GetUeCountOk() (*int32, bool)`

GetUeCountOk returns a tuple with the UeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeCount

`func (o *UePerLocationReport) SetUeCount(v int32)`

SetUeCount sets UeCount field to given value.


### GetExternalIds

`func (o *UePerLocationReport) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *UePerLocationReport) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *UePerLocationReport) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *UePerLocationReport) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.

### GetMsisdns

`func (o *UePerLocationReport) GetMsisdns() []string`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *UePerLocationReport) GetMsisdnsOk() (*[]string, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *UePerLocationReport) SetMsisdns(v []string)`

SetMsisdns sets Msisdns field to given value.

### HasMsisdns

`func (o *UePerLocationReport) HasMsisdns() bool`

HasMsisdns returns a boolean if a field has been set.

### GetServLevelDevIds

`func (o *UePerLocationReport) GetServLevelDevIds() []string`

GetServLevelDevIds returns the ServLevelDevIds field if non-nil, zero value otherwise.

### GetServLevelDevIdsOk

`func (o *UePerLocationReport) GetServLevelDevIdsOk() (*[]string, bool)`

GetServLevelDevIdsOk returns a tuple with the ServLevelDevIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServLevelDevIds

`func (o *UePerLocationReport) SetServLevelDevIds(v []string)`

SetServLevelDevIds sets ServLevelDevIds field to given value.

### HasServLevelDevIds

`func (o *UePerLocationReport) HasServLevelDevIds() bool`

HasServLevelDevIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


