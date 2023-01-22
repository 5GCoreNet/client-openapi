# SupplementaryService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupplementaryServiceType** | Pointer to [**SupplementaryServiceType**](SupplementaryServiceType.md) |  | [optional] 
**SupplementaryServiceMode** | Pointer to [**SupplementaryServiceMode**](SupplementaryServiceMode.md) |  | [optional] 
**NumberOfDiversions** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**AssociatedPartyAddress** | Pointer to **string** |  | [optional] 
**ConferenceId** | Pointer to **string** |  | [optional] 
**ParticipantActionType** | Pointer to [**ParticipantActionType**](ParticipantActionType.md) |  | [optional] 
**ChangeTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**NumberOfParticipants** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**CUGInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewSupplementaryService

`func NewSupplementaryService() *SupplementaryService`

NewSupplementaryService instantiates a new SupplementaryService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupplementaryServiceWithDefaults

`func NewSupplementaryServiceWithDefaults() *SupplementaryService`

NewSupplementaryServiceWithDefaults instantiates a new SupplementaryService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupplementaryServiceType

`func (o *SupplementaryService) GetSupplementaryServiceType() SupplementaryServiceType`

GetSupplementaryServiceType returns the SupplementaryServiceType field if non-nil, zero value otherwise.

### GetSupplementaryServiceTypeOk

`func (o *SupplementaryService) GetSupplementaryServiceTypeOk() (*SupplementaryServiceType, bool)`

GetSupplementaryServiceTypeOk returns a tuple with the SupplementaryServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryServiceType

`func (o *SupplementaryService) SetSupplementaryServiceType(v SupplementaryServiceType)`

SetSupplementaryServiceType sets SupplementaryServiceType field to given value.

### HasSupplementaryServiceType

`func (o *SupplementaryService) HasSupplementaryServiceType() bool`

HasSupplementaryServiceType returns a boolean if a field has been set.

### GetSupplementaryServiceMode

`func (o *SupplementaryService) GetSupplementaryServiceMode() SupplementaryServiceMode`

GetSupplementaryServiceMode returns the SupplementaryServiceMode field if non-nil, zero value otherwise.

### GetSupplementaryServiceModeOk

`func (o *SupplementaryService) GetSupplementaryServiceModeOk() (*SupplementaryServiceMode, bool)`

GetSupplementaryServiceModeOk returns a tuple with the SupplementaryServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplementaryServiceMode

`func (o *SupplementaryService) SetSupplementaryServiceMode(v SupplementaryServiceMode)`

SetSupplementaryServiceMode sets SupplementaryServiceMode field to given value.

### HasSupplementaryServiceMode

`func (o *SupplementaryService) HasSupplementaryServiceMode() bool`

HasSupplementaryServiceMode returns a boolean if a field has been set.

### GetNumberOfDiversions

`func (o *SupplementaryService) GetNumberOfDiversions() int32`

GetNumberOfDiversions returns the NumberOfDiversions field if non-nil, zero value otherwise.

### GetNumberOfDiversionsOk

`func (o *SupplementaryService) GetNumberOfDiversionsOk() (*int32, bool)`

GetNumberOfDiversionsOk returns a tuple with the NumberOfDiversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDiversions

`func (o *SupplementaryService) SetNumberOfDiversions(v int32)`

SetNumberOfDiversions sets NumberOfDiversions field to given value.

### HasNumberOfDiversions

`func (o *SupplementaryService) HasNumberOfDiversions() bool`

HasNumberOfDiversions returns a boolean if a field has been set.

### GetAssociatedPartyAddress

`func (o *SupplementaryService) GetAssociatedPartyAddress() string`

GetAssociatedPartyAddress returns the AssociatedPartyAddress field if non-nil, zero value otherwise.

### GetAssociatedPartyAddressOk

`func (o *SupplementaryService) GetAssociatedPartyAddressOk() (*string, bool)`

GetAssociatedPartyAddressOk returns a tuple with the AssociatedPartyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPartyAddress

`func (o *SupplementaryService) SetAssociatedPartyAddress(v string)`

SetAssociatedPartyAddress sets AssociatedPartyAddress field to given value.

### HasAssociatedPartyAddress

`func (o *SupplementaryService) HasAssociatedPartyAddress() bool`

HasAssociatedPartyAddress returns a boolean if a field has been set.

### GetConferenceId

`func (o *SupplementaryService) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *SupplementaryService) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *SupplementaryService) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *SupplementaryService) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetParticipantActionType

`func (o *SupplementaryService) GetParticipantActionType() ParticipantActionType`

GetParticipantActionType returns the ParticipantActionType field if non-nil, zero value otherwise.

### GetParticipantActionTypeOk

`func (o *SupplementaryService) GetParticipantActionTypeOk() (*ParticipantActionType, bool)`

GetParticipantActionTypeOk returns a tuple with the ParticipantActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantActionType

`func (o *SupplementaryService) SetParticipantActionType(v ParticipantActionType)`

SetParticipantActionType sets ParticipantActionType field to given value.

### HasParticipantActionType

`func (o *SupplementaryService) HasParticipantActionType() bool`

HasParticipantActionType returns a boolean if a field has been set.

### GetChangeTime

`func (o *SupplementaryService) GetChangeTime() time.Time`

GetChangeTime returns the ChangeTime field if non-nil, zero value otherwise.

### GetChangeTimeOk

`func (o *SupplementaryService) GetChangeTimeOk() (*time.Time, bool)`

GetChangeTimeOk returns a tuple with the ChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeTime

`func (o *SupplementaryService) SetChangeTime(v time.Time)`

SetChangeTime sets ChangeTime field to given value.

### HasChangeTime

`func (o *SupplementaryService) HasChangeTime() bool`

HasChangeTime returns a boolean if a field has been set.

### GetNumberOfParticipants

`func (o *SupplementaryService) GetNumberOfParticipants() int32`

GetNumberOfParticipants returns the NumberOfParticipants field if non-nil, zero value otherwise.

### GetNumberOfParticipantsOk

`func (o *SupplementaryService) GetNumberOfParticipantsOk() (*int32, bool)`

GetNumberOfParticipantsOk returns a tuple with the NumberOfParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfParticipants

`func (o *SupplementaryService) SetNumberOfParticipants(v int32)`

SetNumberOfParticipants sets NumberOfParticipants field to given value.

### HasNumberOfParticipants

`func (o *SupplementaryService) HasNumberOfParticipants() bool`

HasNumberOfParticipants returns a boolean if a field has been set.

### GetCUGInformation

`func (o *SupplementaryService) GetCUGInformation() string`

GetCUGInformation returns the CUGInformation field if non-nil, zero value otherwise.

### GetCUGInformationOk

`func (o *SupplementaryService) GetCUGInformationOk() (*string, bool)`

GetCUGInformationOk returns a tuple with the CUGInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCUGInformation

`func (o *SupplementaryService) SetCUGInformation(v string)`

SetCUGInformation sets CUGInformation field to given value.

### HasCUGInformation

`func (o *SupplementaryService) HasCUGInformation() bool`

HasCUGInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


