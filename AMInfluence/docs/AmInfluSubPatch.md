# AmInfluSubPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HighThruInd** | Pointer to **NullableBool** |  | [optional] 
**GeoAreas** | Pointer to [**[]GeographicArea**](GeographicArea.md) | Identifies geographic areas of the user where the request is applicable. | [optional] 
**PolicyDuration** | Pointer to **NullableInt32** | Unsigned integer identifying a period of time in units of seconds with \&quot;nullable&#x3D;true\&quot; property. | [optional] 
**DnnSnssaiInfos** | Pointer to [**[]DnnSnssaiInformation**](DnnSnssaiInformation.md) | Each of the element identifies a (DNN, S-NSSAI) combination. | [optional] 
**AfAppIds** | Pointer to **[]string** | Each of the element identifies an application. | [optional] 
**SubscribedEvents** | Pointer to [**[]AmInfluEvent**](AmInfluEvent.md) | Indicates one or more AM influence related events. | [optional] 
**NotificationDestination** | Pointer to **NullableString** | String formatted according to IETF RFC 3986 identifying a referenced resource, but with the nullable property set to true.  | [optional] 

## Methods

### NewAmInfluSubPatch

`func NewAmInfluSubPatch() *AmInfluSubPatch`

NewAmInfluSubPatch instantiates a new AmInfluSubPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmInfluSubPatchWithDefaults

`func NewAmInfluSubPatchWithDefaults() *AmInfluSubPatch`

NewAmInfluSubPatchWithDefaults instantiates a new AmInfluSubPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighThruInd

`func (o *AmInfluSubPatch) GetHighThruInd() bool`

GetHighThruInd returns the HighThruInd field if non-nil, zero value otherwise.

### GetHighThruIndOk

`func (o *AmInfluSubPatch) GetHighThruIndOk() (*bool, bool)`

GetHighThruIndOk returns a tuple with the HighThruInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighThruInd

`func (o *AmInfluSubPatch) SetHighThruInd(v bool)`

SetHighThruInd sets HighThruInd field to given value.

### HasHighThruInd

`func (o *AmInfluSubPatch) HasHighThruInd() bool`

HasHighThruInd returns a boolean if a field has been set.

### SetHighThruIndNil

`func (o *AmInfluSubPatch) SetHighThruIndNil(b bool)`

 SetHighThruIndNil sets the value for HighThruInd to be an explicit nil

### UnsetHighThruInd
`func (o *AmInfluSubPatch) UnsetHighThruInd()`

UnsetHighThruInd ensures that no value is present for HighThruInd, not even an explicit nil
### GetGeoAreas

`func (o *AmInfluSubPatch) GetGeoAreas() []GeographicArea`

GetGeoAreas returns the GeoAreas field if non-nil, zero value otherwise.

### GetGeoAreasOk

`func (o *AmInfluSubPatch) GetGeoAreasOk() (*[]GeographicArea, bool)`

GetGeoAreasOk returns a tuple with the GeoAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoAreas

`func (o *AmInfluSubPatch) SetGeoAreas(v []GeographicArea)`

SetGeoAreas sets GeoAreas field to given value.

### HasGeoAreas

`func (o *AmInfluSubPatch) HasGeoAreas() bool`

HasGeoAreas returns a boolean if a field has been set.

### SetGeoAreasNil

`func (o *AmInfluSubPatch) SetGeoAreasNil(b bool)`

 SetGeoAreasNil sets the value for GeoAreas to be an explicit nil

### UnsetGeoAreas
`func (o *AmInfluSubPatch) UnsetGeoAreas()`

UnsetGeoAreas ensures that no value is present for GeoAreas, not even an explicit nil
### GetPolicyDuration

`func (o *AmInfluSubPatch) GetPolicyDuration() int32`

GetPolicyDuration returns the PolicyDuration field if non-nil, zero value otherwise.

### GetPolicyDurationOk

`func (o *AmInfluSubPatch) GetPolicyDurationOk() (*int32, bool)`

GetPolicyDurationOk returns a tuple with the PolicyDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDuration

`func (o *AmInfluSubPatch) SetPolicyDuration(v int32)`

SetPolicyDuration sets PolicyDuration field to given value.

### HasPolicyDuration

`func (o *AmInfluSubPatch) HasPolicyDuration() bool`

HasPolicyDuration returns a boolean if a field has been set.

### SetPolicyDurationNil

`func (o *AmInfluSubPatch) SetPolicyDurationNil(b bool)`

 SetPolicyDurationNil sets the value for PolicyDuration to be an explicit nil

### UnsetPolicyDuration
`func (o *AmInfluSubPatch) UnsetPolicyDuration()`

UnsetPolicyDuration ensures that no value is present for PolicyDuration, not even an explicit nil
### GetDnnSnssaiInfos

`func (o *AmInfluSubPatch) GetDnnSnssaiInfos() []DnnSnssaiInformation`

GetDnnSnssaiInfos returns the DnnSnssaiInfos field if non-nil, zero value otherwise.

### GetDnnSnssaiInfosOk

`func (o *AmInfluSubPatch) GetDnnSnssaiInfosOk() (*[]DnnSnssaiInformation, bool)`

GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSnssaiInfos

`func (o *AmInfluSubPatch) SetDnnSnssaiInfos(v []DnnSnssaiInformation)`

SetDnnSnssaiInfos sets DnnSnssaiInfos field to given value.

### HasDnnSnssaiInfos

`func (o *AmInfluSubPatch) HasDnnSnssaiInfos() bool`

HasDnnSnssaiInfos returns a boolean if a field has been set.

### SetDnnSnssaiInfosNil

`func (o *AmInfluSubPatch) SetDnnSnssaiInfosNil(b bool)`

 SetDnnSnssaiInfosNil sets the value for DnnSnssaiInfos to be an explicit nil

### UnsetDnnSnssaiInfos
`func (o *AmInfluSubPatch) UnsetDnnSnssaiInfos()`

UnsetDnnSnssaiInfos ensures that no value is present for DnnSnssaiInfos, not even an explicit nil
### GetAfAppIds

`func (o *AmInfluSubPatch) GetAfAppIds() []string`

GetAfAppIds returns the AfAppIds field if non-nil, zero value otherwise.

### GetAfAppIdsOk

`func (o *AmInfluSubPatch) GetAfAppIdsOk() (*[]string, bool)`

GetAfAppIdsOk returns a tuple with the AfAppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppIds

`func (o *AmInfluSubPatch) SetAfAppIds(v []string)`

SetAfAppIds sets AfAppIds field to given value.

### HasAfAppIds

`func (o *AmInfluSubPatch) HasAfAppIds() bool`

HasAfAppIds returns a boolean if a field has been set.

### SetAfAppIdsNil

`func (o *AmInfluSubPatch) SetAfAppIdsNil(b bool)`

 SetAfAppIdsNil sets the value for AfAppIds to be an explicit nil

### UnsetAfAppIds
`func (o *AmInfluSubPatch) UnsetAfAppIds()`

UnsetAfAppIds ensures that no value is present for AfAppIds, not even an explicit nil
### GetSubscribedEvents

`func (o *AmInfluSubPatch) GetSubscribedEvents() []AmInfluEvent`

GetSubscribedEvents returns the SubscribedEvents field if non-nil, zero value otherwise.

### GetSubscribedEventsOk

`func (o *AmInfluSubPatch) GetSubscribedEventsOk() (*[]AmInfluEvent, bool)`

GetSubscribedEventsOk returns a tuple with the SubscribedEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedEvents

`func (o *AmInfluSubPatch) SetSubscribedEvents(v []AmInfluEvent)`

SetSubscribedEvents sets SubscribedEvents field to given value.

### HasSubscribedEvents

`func (o *AmInfluSubPatch) HasSubscribedEvents() bool`

HasSubscribedEvents returns a boolean if a field has been set.

### SetSubscribedEventsNil

`func (o *AmInfluSubPatch) SetSubscribedEventsNil(b bool)`

 SetSubscribedEventsNil sets the value for SubscribedEvents to be an explicit nil

### UnsetSubscribedEvents
`func (o *AmInfluSubPatch) UnsetSubscribedEvents()`

UnsetSubscribedEvents ensures that no value is present for SubscribedEvents, not even an explicit nil
### GetNotificationDestination

`func (o *AmInfluSubPatch) GetNotificationDestination() string`

GetNotificationDestination returns the NotificationDestination field if non-nil, zero value otherwise.

### GetNotificationDestinationOk

`func (o *AmInfluSubPatch) GetNotificationDestinationOk() (*string, bool)`

GetNotificationDestinationOk returns a tuple with the NotificationDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDestination

`func (o *AmInfluSubPatch) SetNotificationDestination(v string)`

SetNotificationDestination sets NotificationDestination field to given value.

### HasNotificationDestination

`func (o *AmInfluSubPatch) HasNotificationDestination() bool`

HasNotificationDestination returns a boolean if a field has been set.

### SetNotificationDestinationNil

`func (o *AmInfluSubPatch) SetNotificationDestinationNil(b bool)`

 SetNotificationDestinationNil sets the value for NotificationDestination to be an explicit nil

### UnsetNotificationDestination
`func (o *AmInfluSubPatch) UnsetNotificationDestination()`

UnsetNotificationDestination ensures that no value is present for NotificationDestination, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


