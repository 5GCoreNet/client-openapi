# EasDeploySubData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** |  | [optional] 
**DnnSnssaiInfos** | Pointer to [**[]DnnSnssaiInformation**](DnnSnssaiInformation.md) | Each of the element identifies a (DNN, S-NSSAI) combination. | [optional] 
**EventId** | [**EasEvent**](EasEvent.md) |  | 
**EventsNotifs** | Pointer to [**[]EasDeployInfoData**](EasDeployInfoData.md) | Represents the EAS Deployment Information changes event(s) to be reported. Shall only be present if the \&quot;immRep\&quot; attribute is included and sets to true, and the current status of EAS Deployment Information is available.  | [optional] 
**ImmRep** | Pointer to **bool** | Indication of immediate reporting. Set to true: requires the immediate reporting of the  current status of EAS Deployment Information, if available. Set to false (default): EAS  Deployment Information event report occurs when the event is met.  | [optional] 
**InterGroupId** | Pointer to **string** | String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.   | [optional] 
**NotifId** | **string** |  | 
**NotifUri** | **string** | String providing an URI formatted according to RFC 3986. | 

## Methods

### NewEasDeploySubData

`func NewEasDeploySubData(eventId EasEvent, notifId string, notifUri string, ) *EasDeploySubData`

NewEasDeploySubData instantiates a new EasDeploySubData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasDeploySubDataWithDefaults

`func NewEasDeploySubDataWithDefaults() *EasDeploySubData`

NewEasDeploySubDataWithDefaults instantiates a new EasDeploySubData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *EasDeploySubData) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *EasDeploySubData) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *EasDeploySubData) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *EasDeploySubData) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetDnnSnssaiInfos

`func (o *EasDeploySubData) GetDnnSnssaiInfos() []DnnSnssaiInformation`

GetDnnSnssaiInfos returns the DnnSnssaiInfos field if non-nil, zero value otherwise.

### GetDnnSnssaiInfosOk

`func (o *EasDeploySubData) GetDnnSnssaiInfosOk() (*[]DnnSnssaiInformation, bool)`

GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSnssaiInfos

`func (o *EasDeploySubData) SetDnnSnssaiInfos(v []DnnSnssaiInformation)`

SetDnnSnssaiInfos sets DnnSnssaiInfos field to given value.

### HasDnnSnssaiInfos

`func (o *EasDeploySubData) HasDnnSnssaiInfos() bool`

HasDnnSnssaiInfos returns a boolean if a field has been set.

### GetEventId

`func (o *EasDeploySubData) GetEventId() EasEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EasDeploySubData) GetEventIdOk() (*EasEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EasDeploySubData) SetEventId(v EasEvent)`

SetEventId sets EventId field to given value.


### GetEventsNotifs

`func (o *EasDeploySubData) GetEventsNotifs() []EasDeployInfoData`

GetEventsNotifs returns the EventsNotifs field if non-nil, zero value otherwise.

### GetEventsNotifsOk

`func (o *EasDeploySubData) GetEventsNotifsOk() (*[]EasDeployInfoData, bool)`

GetEventsNotifsOk returns a tuple with the EventsNotifs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsNotifs

`func (o *EasDeploySubData) SetEventsNotifs(v []EasDeployInfoData)`

SetEventsNotifs sets EventsNotifs field to given value.

### HasEventsNotifs

`func (o *EasDeploySubData) HasEventsNotifs() bool`

HasEventsNotifs returns a boolean if a field has been set.

### GetImmRep

`func (o *EasDeploySubData) GetImmRep() bool`

GetImmRep returns the ImmRep field if non-nil, zero value otherwise.

### GetImmRepOk

`func (o *EasDeploySubData) GetImmRepOk() (*bool, bool)`

GetImmRepOk returns a tuple with the ImmRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmRep

`func (o *EasDeploySubData) SetImmRep(v bool)`

SetImmRep sets ImmRep field to given value.

### HasImmRep

`func (o *EasDeploySubData) HasImmRep() bool`

HasImmRep returns a boolean if a field has been set.

### GetInterGroupId

`func (o *EasDeploySubData) GetInterGroupId() string`

GetInterGroupId returns the InterGroupId field if non-nil, zero value otherwise.

### GetInterGroupIdOk

`func (o *EasDeploySubData) GetInterGroupIdOk() (*string, bool)`

GetInterGroupIdOk returns a tuple with the InterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterGroupId

`func (o *EasDeploySubData) SetInterGroupId(v string)`

SetInterGroupId sets InterGroupId field to given value.

### HasInterGroupId

`func (o *EasDeploySubData) HasInterGroupId() bool`

HasInterGroupId returns a boolean if a field has been set.

### GetNotifId

`func (o *EasDeploySubData) GetNotifId() string`

GetNotifId returns the NotifId field if non-nil, zero value otherwise.

### GetNotifIdOk

`func (o *EasDeploySubData) GetNotifIdOk() (*string, bool)`

GetNotifIdOk returns a tuple with the NotifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifId

`func (o *EasDeploySubData) SetNotifId(v string)`

SetNotifId sets NotifId field to given value.


### GetNotifUri

`func (o *EasDeploySubData) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *EasDeploySubData) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *EasDeploySubData) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


