# EASProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasId** | **string** | Identifier of the EAS. | 
**EndPt** | [**EndPoint**](EndPoint.md) |  | 
**AcIds** | Pointer to **[]string** | Identities of application clients that are served by the EAS. | [optional] 
**ProvId** | Pointer to **string** | Identifier of the ASP that provides the EAS. | [optional] 
**Type** | Pointer to [**EASCategory**](EASCategory.md) |  | [optional] 
**Scheds** | Pointer to [**[]ScheduledCommunicationTime**](ScheduledCommunicationTime.md) | The availability schedule of the EAS. | [optional] 
**SvcArea** | Pointer to [**ServiceArea**](ServiceArea.md) |  | [optional] 
**SvcKpi** | Pointer to [**EASServiceKPI**](EASServiceKPI.md) |  | [optional] 
**PermLvl** | Pointer to [**[]PermissionLevel**](PermissionLevel.md) | level of service permissions supported by the EAS. | [optional] 
**EasFeats** | Pointer to **[]string** | Service specific features supported by EAS. | [optional] 
**AppLocs** | Pointer to [**[]RouteToLocation**](RouteToLocation.md) | List of DNAI(s) and the N6 traffic information associated with the EAS. | [optional] 
**SvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | The ACR scenarios supported by the EAS for service continuity. | [optional] 
**AvlRep** | Pointer to **int32** | Unsigned integer identifying a period of time in units of seconds. | [optional] 
**Status** | Pointer to **string** | EAS status information. | [optional] 

## Methods

### NewEASProfile

`func NewEASProfile(easId string, endPt EndPoint, ) *EASProfile`

NewEASProfile instantiates a new EASProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEASProfileWithDefaults

`func NewEASProfileWithDefaults() *EASProfile`

NewEASProfileWithDefaults instantiates a new EASProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasId

`func (o *EASProfile) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *EASProfile) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *EASProfile) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetEndPt

`func (o *EASProfile) GetEndPt() EndPoint`

GetEndPt returns the EndPt field if non-nil, zero value otherwise.

### GetEndPtOk

`func (o *EASProfile) GetEndPtOk() (*EndPoint, bool)`

GetEndPtOk returns a tuple with the EndPt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPt

`func (o *EASProfile) SetEndPt(v EndPoint)`

SetEndPt sets EndPt field to given value.


### GetAcIds

`func (o *EASProfile) GetAcIds() []string`

GetAcIds returns the AcIds field if non-nil, zero value otherwise.

### GetAcIdsOk

`func (o *EASProfile) GetAcIdsOk() (*[]string, bool)`

GetAcIdsOk returns a tuple with the AcIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcIds

`func (o *EASProfile) SetAcIds(v []string)`

SetAcIds sets AcIds field to given value.

### HasAcIds

`func (o *EASProfile) HasAcIds() bool`

HasAcIds returns a boolean if a field has been set.

### GetProvId

`func (o *EASProfile) GetProvId() string`

GetProvId returns the ProvId field if non-nil, zero value otherwise.

### GetProvIdOk

`func (o *EASProfile) GetProvIdOk() (*string, bool)`

GetProvIdOk returns a tuple with the ProvId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvId

`func (o *EASProfile) SetProvId(v string)`

SetProvId sets ProvId field to given value.

### HasProvId

`func (o *EASProfile) HasProvId() bool`

HasProvId returns a boolean if a field has been set.

### GetType

`func (o *EASProfile) GetType() EASCategory`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EASProfile) GetTypeOk() (*EASCategory, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EASProfile) SetType(v EASCategory)`

SetType sets Type field to given value.

### HasType

`func (o *EASProfile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetScheds

`func (o *EASProfile) GetScheds() []ScheduledCommunicationTime`

GetScheds returns the Scheds field if non-nil, zero value otherwise.

### GetSchedsOk

`func (o *EASProfile) GetSchedsOk() (*[]ScheduledCommunicationTime, bool)`

GetSchedsOk returns a tuple with the Scheds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheds

`func (o *EASProfile) SetScheds(v []ScheduledCommunicationTime)`

SetScheds sets Scheds field to given value.

### HasScheds

`func (o *EASProfile) HasScheds() bool`

HasScheds returns a boolean if a field has been set.

### GetSvcArea

`func (o *EASProfile) GetSvcArea() ServiceArea`

GetSvcArea returns the SvcArea field if non-nil, zero value otherwise.

### GetSvcAreaOk

`func (o *EASProfile) GetSvcAreaOk() (*ServiceArea, bool)`

GetSvcAreaOk returns a tuple with the SvcArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcArea

`func (o *EASProfile) SetSvcArea(v ServiceArea)`

SetSvcArea sets SvcArea field to given value.

### HasSvcArea

`func (o *EASProfile) HasSvcArea() bool`

HasSvcArea returns a boolean if a field has been set.

### GetSvcKpi

`func (o *EASProfile) GetSvcKpi() EASServiceKPI`

GetSvcKpi returns the SvcKpi field if non-nil, zero value otherwise.

### GetSvcKpiOk

`func (o *EASProfile) GetSvcKpiOk() (*EASServiceKPI, bool)`

GetSvcKpiOk returns a tuple with the SvcKpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcKpi

`func (o *EASProfile) SetSvcKpi(v EASServiceKPI)`

SetSvcKpi sets SvcKpi field to given value.

### HasSvcKpi

`func (o *EASProfile) HasSvcKpi() bool`

HasSvcKpi returns a boolean if a field has been set.

### GetPermLvl

`func (o *EASProfile) GetPermLvl() []PermissionLevel`

GetPermLvl returns the PermLvl field if non-nil, zero value otherwise.

### GetPermLvlOk

`func (o *EASProfile) GetPermLvlOk() (*[]PermissionLevel, bool)`

GetPermLvlOk returns a tuple with the PermLvl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermLvl

`func (o *EASProfile) SetPermLvl(v []PermissionLevel)`

SetPermLvl sets PermLvl field to given value.

### HasPermLvl

`func (o *EASProfile) HasPermLvl() bool`

HasPermLvl returns a boolean if a field has been set.

### GetEasFeats

`func (o *EASProfile) GetEasFeats() []string`

GetEasFeats returns the EasFeats field if non-nil, zero value otherwise.

### GetEasFeatsOk

`func (o *EASProfile) GetEasFeatsOk() (*[]string, bool)`

GetEasFeatsOk returns a tuple with the EasFeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasFeats

`func (o *EASProfile) SetEasFeats(v []string)`

SetEasFeats sets EasFeats field to given value.

### HasEasFeats

`func (o *EASProfile) HasEasFeats() bool`

HasEasFeats returns a boolean if a field has been set.

### GetAppLocs

`func (o *EASProfile) GetAppLocs() []RouteToLocation`

GetAppLocs returns the AppLocs field if non-nil, zero value otherwise.

### GetAppLocsOk

`func (o *EASProfile) GetAppLocsOk() (*[]RouteToLocation, bool)`

GetAppLocsOk returns a tuple with the AppLocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppLocs

`func (o *EASProfile) SetAppLocs(v []RouteToLocation)`

SetAppLocs sets AppLocs field to given value.

### HasAppLocs

`func (o *EASProfile) HasAppLocs() bool`

HasAppLocs returns a boolean if a field has been set.

### GetSvcContSupp

`func (o *EASProfile) GetSvcContSupp() []ACRScenario`

GetSvcContSupp returns the SvcContSupp field if non-nil, zero value otherwise.

### GetSvcContSuppOk

`func (o *EASProfile) GetSvcContSuppOk() (*[]ACRScenario, bool)`

GetSvcContSuppOk returns a tuple with the SvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcContSupp

`func (o *EASProfile) SetSvcContSupp(v []ACRScenario)`

SetSvcContSupp sets SvcContSupp field to given value.

### HasSvcContSupp

`func (o *EASProfile) HasSvcContSupp() bool`

HasSvcContSupp returns a boolean if a field has been set.

### GetAvlRep

`func (o *EASProfile) GetAvlRep() int32`

GetAvlRep returns the AvlRep field if non-nil, zero value otherwise.

### GetAvlRepOk

`func (o *EASProfile) GetAvlRepOk() (*int32, bool)`

GetAvlRepOk returns a tuple with the AvlRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvlRep

`func (o *EASProfile) SetAvlRep(v int32)`

SetAvlRep sets AvlRep field to given value.

### HasAvlRep

`func (o *EASProfile) HasAvlRep() bool`

HasAvlRep returns a boolean if a field has been set.

### GetStatus

`func (o *EASProfile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EASProfile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EASProfile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EASProfile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


