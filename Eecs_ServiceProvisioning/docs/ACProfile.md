# ACProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcId** | **string** | Identity of the AC. | 
**AcType** | Pointer to **string** | The category or type of AC. | [optional] 
**PrefEcsps** | Pointer to **[]string** | Indicates to the ECS which ECSPs are preferred for the AC. | [optional] 
**AcSchedule** | Pointer to [**ScheduledCommunicationTime**](ScheduledCommunicationTime.md) |  | [optional] 
**ExpAcGeoServArea** | Pointer to [**LocationArea5G**](LocationArea5G.md) |  | [optional] 
**AcSvcContSupp** | Pointer to [**[]ACRScenario**](ACRScenario.md) | Profiles of ACs for which the EEC provides edge enabling services. | [optional] 
**Eass** | Pointer to [**[]EasDetail**](EasDetail.md) | List of EAS information. | [optional] 

## Methods

### NewACProfile

`func NewACProfile(acId string, ) *ACProfile`

NewACProfile instantiates a new ACProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACProfileWithDefaults

`func NewACProfileWithDefaults() *ACProfile`

NewACProfileWithDefaults instantiates a new ACProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcId

`func (o *ACProfile) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *ACProfile) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *ACProfile) SetAcId(v string)`

SetAcId sets AcId field to given value.


### GetAcType

`func (o *ACProfile) GetAcType() string`

GetAcType returns the AcType field if non-nil, zero value otherwise.

### GetAcTypeOk

`func (o *ACProfile) GetAcTypeOk() (*string, bool)`

GetAcTypeOk returns a tuple with the AcType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcType

`func (o *ACProfile) SetAcType(v string)`

SetAcType sets AcType field to given value.

### HasAcType

`func (o *ACProfile) HasAcType() bool`

HasAcType returns a boolean if a field has been set.

### GetPrefEcsps

`func (o *ACProfile) GetPrefEcsps() []string`

GetPrefEcsps returns the PrefEcsps field if non-nil, zero value otherwise.

### GetPrefEcspsOk

`func (o *ACProfile) GetPrefEcspsOk() (*[]string, bool)`

GetPrefEcspsOk returns a tuple with the PrefEcsps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefEcsps

`func (o *ACProfile) SetPrefEcsps(v []string)`

SetPrefEcsps sets PrefEcsps field to given value.

### HasPrefEcsps

`func (o *ACProfile) HasPrefEcsps() bool`

HasPrefEcsps returns a boolean if a field has been set.

### GetAcSchedule

`func (o *ACProfile) GetAcSchedule() ScheduledCommunicationTime`

GetAcSchedule returns the AcSchedule field if non-nil, zero value otherwise.

### GetAcScheduleOk

`func (o *ACProfile) GetAcScheduleOk() (*ScheduledCommunicationTime, bool)`

GetAcScheduleOk returns a tuple with the AcSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcSchedule

`func (o *ACProfile) SetAcSchedule(v ScheduledCommunicationTime)`

SetAcSchedule sets AcSchedule field to given value.

### HasAcSchedule

`func (o *ACProfile) HasAcSchedule() bool`

HasAcSchedule returns a boolean if a field has been set.

### GetExpAcGeoServArea

`func (o *ACProfile) GetExpAcGeoServArea() LocationArea5G`

GetExpAcGeoServArea returns the ExpAcGeoServArea field if non-nil, zero value otherwise.

### GetExpAcGeoServAreaOk

`func (o *ACProfile) GetExpAcGeoServAreaOk() (*LocationArea5G, bool)`

GetExpAcGeoServAreaOk returns a tuple with the ExpAcGeoServArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpAcGeoServArea

`func (o *ACProfile) SetExpAcGeoServArea(v LocationArea5G)`

SetExpAcGeoServArea sets ExpAcGeoServArea field to given value.

### HasExpAcGeoServArea

`func (o *ACProfile) HasExpAcGeoServArea() bool`

HasExpAcGeoServArea returns a boolean if a field has been set.

### GetAcSvcContSupp

`func (o *ACProfile) GetAcSvcContSupp() []ACRScenario`

GetAcSvcContSupp returns the AcSvcContSupp field if non-nil, zero value otherwise.

### GetAcSvcContSuppOk

`func (o *ACProfile) GetAcSvcContSuppOk() (*[]ACRScenario, bool)`

GetAcSvcContSuppOk returns a tuple with the AcSvcContSupp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcSvcContSupp

`func (o *ACProfile) SetAcSvcContSupp(v []ACRScenario)`

SetAcSvcContSupp sets AcSvcContSupp field to given value.

### HasAcSvcContSupp

`func (o *ACProfile) HasAcSvcContSupp() bool`

HasAcSvcContSupp returns a boolean if a field has been set.

### GetEass

`func (o *ACProfile) GetEass() []EasDetail`

GetEass returns the Eass field if non-nil, zero value otherwise.

### GetEassOk

`func (o *ACProfile) GetEassOk() (*[]EasDetail, bool)`

GetEassOk returns a tuple with the Eass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEass

`func (o *ACProfile) SetEass(v []EasDetail)`

SetEass sets Eass field to given value.

### HasEass

`func (o *ACProfile) HasEass() bool`

HasEass returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


