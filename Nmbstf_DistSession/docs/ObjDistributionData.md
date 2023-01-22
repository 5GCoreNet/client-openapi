# ObjDistributionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjDistributionOperatingMode** | [**ObjDistributionOperatingMode**](ObjDistributionOperatingMode.md) |  | 
**ObjAcquisitionMethod** | [**ObjAcquisitionMethod**](ObjAcquisitionMethod.md) |  | 
**ObjAcquisitionIdsPull** | Pointer to **[]string** |  | [optional] 
**ObjAcquisitionIdPush** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ObjIngestBaseUrl** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ObjDistributionBaseUrl** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**ObjRepairBaseUrl** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 

## Methods

### NewObjDistributionData

`func NewObjDistributionData(objDistributionOperatingMode ObjDistributionOperatingMode, objAcquisitionMethod ObjAcquisitionMethod, ) *ObjDistributionData`

NewObjDistributionData instantiates a new ObjDistributionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjDistributionDataWithDefaults

`func NewObjDistributionDataWithDefaults() *ObjDistributionData`

NewObjDistributionDataWithDefaults instantiates a new ObjDistributionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjDistributionOperatingMode

`func (o *ObjDistributionData) GetObjDistributionOperatingMode() ObjDistributionOperatingMode`

GetObjDistributionOperatingMode returns the ObjDistributionOperatingMode field if non-nil, zero value otherwise.

### GetObjDistributionOperatingModeOk

`func (o *ObjDistributionData) GetObjDistributionOperatingModeOk() (*ObjDistributionOperatingMode, bool)`

GetObjDistributionOperatingModeOk returns a tuple with the ObjDistributionOperatingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDistributionOperatingMode

`func (o *ObjDistributionData) SetObjDistributionOperatingMode(v ObjDistributionOperatingMode)`

SetObjDistributionOperatingMode sets ObjDistributionOperatingMode field to given value.


### GetObjAcquisitionMethod

`func (o *ObjDistributionData) GetObjAcquisitionMethod() ObjAcquisitionMethod`

GetObjAcquisitionMethod returns the ObjAcquisitionMethod field if non-nil, zero value otherwise.

### GetObjAcquisitionMethodOk

`func (o *ObjDistributionData) GetObjAcquisitionMethodOk() (*ObjAcquisitionMethod, bool)`

GetObjAcquisitionMethodOk returns a tuple with the ObjAcquisitionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAcquisitionMethod

`func (o *ObjDistributionData) SetObjAcquisitionMethod(v ObjAcquisitionMethod)`

SetObjAcquisitionMethod sets ObjAcquisitionMethod field to given value.


### GetObjAcquisitionIdsPull

`func (o *ObjDistributionData) GetObjAcquisitionIdsPull() []string`

GetObjAcquisitionIdsPull returns the ObjAcquisitionIdsPull field if non-nil, zero value otherwise.

### GetObjAcquisitionIdsPullOk

`func (o *ObjDistributionData) GetObjAcquisitionIdsPullOk() (*[]string, bool)`

GetObjAcquisitionIdsPullOk returns a tuple with the ObjAcquisitionIdsPull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAcquisitionIdsPull

`func (o *ObjDistributionData) SetObjAcquisitionIdsPull(v []string)`

SetObjAcquisitionIdsPull sets ObjAcquisitionIdsPull field to given value.

### HasObjAcquisitionIdsPull

`func (o *ObjDistributionData) HasObjAcquisitionIdsPull() bool`

HasObjAcquisitionIdsPull returns a boolean if a field has been set.

### GetObjAcquisitionIdPush

`func (o *ObjDistributionData) GetObjAcquisitionIdPush() string`

GetObjAcquisitionIdPush returns the ObjAcquisitionIdPush field if non-nil, zero value otherwise.

### GetObjAcquisitionIdPushOk

`func (o *ObjDistributionData) GetObjAcquisitionIdPushOk() (*string, bool)`

GetObjAcquisitionIdPushOk returns a tuple with the ObjAcquisitionIdPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjAcquisitionIdPush

`func (o *ObjDistributionData) SetObjAcquisitionIdPush(v string)`

SetObjAcquisitionIdPush sets ObjAcquisitionIdPush field to given value.

### HasObjAcquisitionIdPush

`func (o *ObjDistributionData) HasObjAcquisitionIdPush() bool`

HasObjAcquisitionIdPush returns a boolean if a field has been set.

### GetObjIngestBaseUrl

`func (o *ObjDistributionData) GetObjIngestBaseUrl() string`

GetObjIngestBaseUrl returns the ObjIngestBaseUrl field if non-nil, zero value otherwise.

### GetObjIngestBaseUrlOk

`func (o *ObjDistributionData) GetObjIngestBaseUrlOk() (*string, bool)`

GetObjIngestBaseUrlOk returns a tuple with the ObjIngestBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjIngestBaseUrl

`func (o *ObjDistributionData) SetObjIngestBaseUrl(v string)`

SetObjIngestBaseUrl sets ObjIngestBaseUrl field to given value.

### HasObjIngestBaseUrl

`func (o *ObjDistributionData) HasObjIngestBaseUrl() bool`

HasObjIngestBaseUrl returns a boolean if a field has been set.

### GetObjDistributionBaseUrl

`func (o *ObjDistributionData) GetObjDistributionBaseUrl() string`

GetObjDistributionBaseUrl returns the ObjDistributionBaseUrl field if non-nil, zero value otherwise.

### GetObjDistributionBaseUrlOk

`func (o *ObjDistributionData) GetObjDistributionBaseUrlOk() (*string, bool)`

GetObjDistributionBaseUrlOk returns a tuple with the ObjDistributionBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjDistributionBaseUrl

`func (o *ObjDistributionData) SetObjDistributionBaseUrl(v string)`

SetObjDistributionBaseUrl sets ObjDistributionBaseUrl field to given value.

### HasObjDistributionBaseUrl

`func (o *ObjDistributionData) HasObjDistributionBaseUrl() bool`

HasObjDistributionBaseUrl returns a boolean if a field has been set.

### GetObjRepairBaseUrl

`func (o *ObjDistributionData) GetObjRepairBaseUrl() string`

GetObjRepairBaseUrl returns the ObjRepairBaseUrl field if non-nil, zero value otherwise.

### GetObjRepairBaseUrlOk

`func (o *ObjDistributionData) GetObjRepairBaseUrlOk() (*string, bool)`

GetObjRepairBaseUrlOk returns a tuple with the ObjRepairBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjRepairBaseUrl

`func (o *ObjDistributionData) SetObjRepairBaseUrl(v string)`

SetObjRepairBaseUrl sets ObjRepairBaseUrl field to given value.

### HasObjRepairBaseUrl

`func (o *ObjDistributionData) HasObjRepairBaseUrl() bool`

HasObjRepairBaseUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


