# ACFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcTypes** | Pointer to **[]string** | List of AC Types or categories. | [optional] 
**EcspIds** | Pointer to **[]string** | The list of identifiers of ECSPs associated with the EEC. | [optional] 
**AcIds** | Pointer to **[]string** | List of identifiers of ACs to be matched. | [optional] 
**SvcArea** | Pointer to [**ServiceArea**](ServiceArea.md) |  | [optional] 
**MaxAcKpi** | Pointer to [**ACServiceKPIs**](ACServiceKPIs.md) |  | [optional] 
**MinAcKpi** | Pointer to [**ACServiceKPIs**](ACServiceKPIs.md) |  | [optional] 
**OpSchds** | Pointer to [**[]ScheduledCommunicationTime**](ScheduledCommunicationTime.md) | Operation schedule of EAS to be matched with operation schedule of the AC. | [optional] 
**UeIds** | Pointer to **[]string** | List of UE identifiers hosting the AC. | [optional] 
**LocInfs** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 

## Methods

### NewACFilters

`func NewACFilters() *ACFilters`

NewACFilters instantiates a new ACFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACFiltersWithDefaults

`func NewACFiltersWithDefaults() *ACFilters`

NewACFiltersWithDefaults instantiates a new ACFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcTypes

`func (o *ACFilters) GetAcTypes() []string`

GetAcTypes returns the AcTypes field if non-nil, zero value otherwise.

### GetAcTypesOk

`func (o *ACFilters) GetAcTypesOk() (*[]string, bool)`

GetAcTypesOk returns a tuple with the AcTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcTypes

`func (o *ACFilters) SetAcTypes(v []string)`

SetAcTypes sets AcTypes field to given value.

### HasAcTypes

`func (o *ACFilters) HasAcTypes() bool`

HasAcTypes returns a boolean if a field has been set.

### GetEcspIds

`func (o *ACFilters) GetEcspIds() []string`

GetEcspIds returns the EcspIds field if non-nil, zero value otherwise.

### GetEcspIdsOk

`func (o *ACFilters) GetEcspIdsOk() (*[]string, bool)`

GetEcspIdsOk returns a tuple with the EcspIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcspIds

`func (o *ACFilters) SetEcspIds(v []string)`

SetEcspIds sets EcspIds field to given value.

### HasEcspIds

`func (o *ACFilters) HasEcspIds() bool`

HasEcspIds returns a boolean if a field has been set.

### GetAcIds

`func (o *ACFilters) GetAcIds() []string`

GetAcIds returns the AcIds field if non-nil, zero value otherwise.

### GetAcIdsOk

`func (o *ACFilters) GetAcIdsOk() (*[]string, bool)`

GetAcIdsOk returns a tuple with the AcIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcIds

`func (o *ACFilters) SetAcIds(v []string)`

SetAcIds sets AcIds field to given value.

### HasAcIds

`func (o *ACFilters) HasAcIds() bool`

HasAcIds returns a boolean if a field has been set.

### GetSvcArea

`func (o *ACFilters) GetSvcArea() ServiceArea`

GetSvcArea returns the SvcArea field if non-nil, zero value otherwise.

### GetSvcAreaOk

`func (o *ACFilters) GetSvcAreaOk() (*ServiceArea, bool)`

GetSvcAreaOk returns a tuple with the SvcArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcArea

`func (o *ACFilters) SetSvcArea(v ServiceArea)`

SetSvcArea sets SvcArea field to given value.

### HasSvcArea

`func (o *ACFilters) HasSvcArea() bool`

HasSvcArea returns a boolean if a field has been set.

### GetMaxAcKpi

`func (o *ACFilters) GetMaxAcKpi() ACServiceKPIs`

GetMaxAcKpi returns the MaxAcKpi field if non-nil, zero value otherwise.

### GetMaxAcKpiOk

`func (o *ACFilters) GetMaxAcKpiOk() (*ACServiceKPIs, bool)`

GetMaxAcKpiOk returns a tuple with the MaxAcKpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAcKpi

`func (o *ACFilters) SetMaxAcKpi(v ACServiceKPIs)`

SetMaxAcKpi sets MaxAcKpi field to given value.

### HasMaxAcKpi

`func (o *ACFilters) HasMaxAcKpi() bool`

HasMaxAcKpi returns a boolean if a field has been set.

### GetMinAcKpi

`func (o *ACFilters) GetMinAcKpi() ACServiceKPIs`

GetMinAcKpi returns the MinAcKpi field if non-nil, zero value otherwise.

### GetMinAcKpiOk

`func (o *ACFilters) GetMinAcKpiOk() (*ACServiceKPIs, bool)`

GetMinAcKpiOk returns a tuple with the MinAcKpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAcKpi

`func (o *ACFilters) SetMinAcKpi(v ACServiceKPIs)`

SetMinAcKpi sets MinAcKpi field to given value.

### HasMinAcKpi

`func (o *ACFilters) HasMinAcKpi() bool`

HasMinAcKpi returns a boolean if a field has been set.

### GetOpSchds

`func (o *ACFilters) GetOpSchds() []ScheduledCommunicationTime`

GetOpSchds returns the OpSchds field if non-nil, zero value otherwise.

### GetOpSchdsOk

`func (o *ACFilters) GetOpSchdsOk() (*[]ScheduledCommunicationTime, bool)`

GetOpSchdsOk returns a tuple with the OpSchds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpSchds

`func (o *ACFilters) SetOpSchds(v []ScheduledCommunicationTime)`

SetOpSchds sets OpSchds field to given value.

### HasOpSchds

`func (o *ACFilters) HasOpSchds() bool`

HasOpSchds returns a boolean if a field has been set.

### GetUeIds

`func (o *ACFilters) GetUeIds() []string`

GetUeIds returns the UeIds field if non-nil, zero value otherwise.

### GetUeIdsOk

`func (o *ACFilters) GetUeIdsOk() (*[]string, bool)`

GetUeIdsOk returns a tuple with the UeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeIds

`func (o *ACFilters) SetUeIds(v []string)`

SetUeIds sets UeIds field to given value.

### HasUeIds

`func (o *ACFilters) HasUeIds() bool`

HasUeIds returns a boolean if a field has been set.

### GetLocInfs

`func (o *ACFilters) GetLocInfs() LocationArea5G`

GetLocInfs returns the LocInfs field if non-nil, zero value otherwise.

### GetLocInfsOk

`func (o *ACFilters) GetLocInfsOk() (*LocationArea5G, bool)`

GetLocInfsOk returns a tuple with the LocInfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfs

`func (o *ACFilters) SetLocInfs(v LocationArea5G)`

SetLocInfs sets LocInfs field to given value.

### HasLocInfs

`func (o *ACFilters) HasLocInfs() bool`

HasLocInfs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


