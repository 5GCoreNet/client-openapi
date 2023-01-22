# EventDataProcessingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventDataProcessingConfigurationId** | **string** | String chosen by the 5GMS AF to serve as an identifier in a resource URI. | 
**EventId** | [**AfEvent**](AfEvent.md) |  | 
**AuthorizationUrl** | Pointer to **string** | Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986. | [optional] 
**DataAccessProfiles** | [**[]DataAccessProfile**](DataAccessProfile.md) |  | 

## Methods

### NewEventDataProcessingConfiguration

`func NewEventDataProcessingConfiguration(eventDataProcessingConfigurationId string, eventId AfEvent, dataAccessProfiles []DataAccessProfile, ) *EventDataProcessingConfiguration`

NewEventDataProcessingConfiguration instantiates a new EventDataProcessingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventDataProcessingConfigurationWithDefaults

`func NewEventDataProcessingConfigurationWithDefaults() *EventDataProcessingConfiguration`

NewEventDataProcessingConfigurationWithDefaults instantiates a new EventDataProcessingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventDataProcessingConfigurationId

`func (o *EventDataProcessingConfiguration) GetEventDataProcessingConfigurationId() string`

GetEventDataProcessingConfigurationId returns the EventDataProcessingConfigurationId field if non-nil, zero value otherwise.

### GetEventDataProcessingConfigurationIdOk

`func (o *EventDataProcessingConfiguration) GetEventDataProcessingConfigurationIdOk() (*string, bool)`

GetEventDataProcessingConfigurationIdOk returns a tuple with the EventDataProcessingConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDataProcessingConfigurationId

`func (o *EventDataProcessingConfiguration) SetEventDataProcessingConfigurationId(v string)`

SetEventDataProcessingConfigurationId sets EventDataProcessingConfigurationId field to given value.


### GetEventId

`func (o *EventDataProcessingConfiguration) GetEventId() AfEvent`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventDataProcessingConfiguration) GetEventIdOk() (*AfEvent, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventDataProcessingConfiguration) SetEventId(v AfEvent)`

SetEventId sets EventId field to given value.


### GetAuthorizationUrl

`func (o *EventDataProcessingConfiguration) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *EventDataProcessingConfiguration) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *EventDataProcessingConfiguration) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.

### HasAuthorizationUrl

`func (o *EventDataProcessingConfiguration) HasAuthorizationUrl() bool`

HasAuthorizationUrl returns a boolean if a field has been set.

### GetDataAccessProfiles

`func (o *EventDataProcessingConfiguration) GetDataAccessProfiles() []DataAccessProfile`

GetDataAccessProfiles returns the DataAccessProfiles field if non-nil, zero value otherwise.

### GetDataAccessProfilesOk

`func (o *EventDataProcessingConfiguration) GetDataAccessProfilesOk() (*[]DataAccessProfile, bool)`

GetDataAccessProfilesOk returns a tuple with the DataAccessProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAccessProfiles

`func (o *EventDataProcessingConfiguration) SetDataAccessProfiles(v []DataAccessProfile)`

SetDataAccessProfiles sets DataAccessProfiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


