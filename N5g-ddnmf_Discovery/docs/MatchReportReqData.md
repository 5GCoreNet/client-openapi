# MatchReportReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**ProseAppCodes** | Pointer to **[]string** |  | [optional] 
**MoniteredPlmnId** | Pointer to [**PlmnId**](PlmnId.md) |  | [optional] 

## Methods

### NewMatchReportReqData

`func NewMatchReportReqData(discType DiscoveryType, ) *MatchReportReqData`

NewMatchReportReqData instantiates a new MatchReportReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchReportReqDataWithDefaults

`func NewMatchReportReqDataWithDefaults() *MatchReportReqData`

NewMatchReportReqDataWithDefaults instantiates a new MatchReportReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *MatchReportReqData) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *MatchReportReqData) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *MatchReportReqData) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetProseAppCodes

`func (o *MatchReportReqData) GetProseAppCodes() []string`

GetProseAppCodes returns the ProseAppCodes field if non-nil, zero value otherwise.

### GetProseAppCodesOk

`func (o *MatchReportReqData) GetProseAppCodesOk() (*[]string, bool)`

GetProseAppCodesOk returns a tuple with the ProseAppCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppCodes

`func (o *MatchReportReqData) SetProseAppCodes(v []string)`

SetProseAppCodes sets ProseAppCodes field to given value.

### HasProseAppCodes

`func (o *MatchReportReqData) HasProseAppCodes() bool`

HasProseAppCodes returns a boolean if a field has been set.

### GetMoniteredPlmnId

`func (o *MatchReportReqData) GetMoniteredPlmnId() PlmnId`

GetMoniteredPlmnId returns the MoniteredPlmnId field if non-nil, zero value otherwise.

### GetMoniteredPlmnIdOk

`func (o *MatchReportReqData) GetMoniteredPlmnIdOk() (*PlmnId, bool)`

GetMoniteredPlmnIdOk returns a tuple with the MoniteredPlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoniteredPlmnId

`func (o *MatchReportReqData) SetMoniteredPlmnId(v PlmnId)`

SetMoniteredPlmnId sets MoniteredPlmnId field to given value.

### HasMoniteredPlmnId

`func (o *MatchReportReqData) HasMoniteredPlmnId() bool`

HasMoniteredPlmnId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


