# TadsInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoiceOverPsSessionSupport** | [**ImsVoiceOverPsSessionSupport**](ImsVoiceOverPsSessionSupport.md) |  | 
**AccessType** | Pointer to [**AccessType1**](AccessType1.md) |  | [optional] 
**RatType** | Pointer to [**RatType**](RatType.md) |  | [optional] 
**LastUeActivityTime** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewTadsInformation

`func NewTadsInformation(voiceOverPsSessionSupport ImsVoiceOverPsSessionSupport, ) *TadsInformation`

NewTadsInformation instantiates a new TadsInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTadsInformationWithDefaults

`func NewTadsInformationWithDefaults() *TadsInformation`

NewTadsInformationWithDefaults instantiates a new TadsInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoiceOverPsSessionSupport

`func (o *TadsInformation) GetVoiceOverPsSessionSupport() ImsVoiceOverPsSessionSupport`

GetVoiceOverPsSessionSupport returns the VoiceOverPsSessionSupport field if non-nil, zero value otherwise.

### GetVoiceOverPsSessionSupportOk

`func (o *TadsInformation) GetVoiceOverPsSessionSupportOk() (*ImsVoiceOverPsSessionSupport, bool)`

GetVoiceOverPsSessionSupportOk returns a tuple with the VoiceOverPsSessionSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceOverPsSessionSupport

`func (o *TadsInformation) SetVoiceOverPsSessionSupport(v ImsVoiceOverPsSessionSupport)`

SetVoiceOverPsSessionSupport sets VoiceOverPsSessionSupport field to given value.


### GetAccessType

`func (o *TadsInformation) GetAccessType() AccessType1`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *TadsInformation) GetAccessTypeOk() (*AccessType1, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *TadsInformation) SetAccessType(v AccessType1)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *TadsInformation) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetRatType

`func (o *TadsInformation) GetRatType() RatType`

GetRatType returns the RatType field if non-nil, zero value otherwise.

### GetRatTypeOk

`func (o *TadsInformation) GetRatTypeOk() (*RatType, bool)`

GetRatTypeOk returns a tuple with the RatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatType

`func (o *TadsInformation) SetRatType(v RatType)`

SetRatType sets RatType field to given value.

### HasRatType

`func (o *TadsInformation) HasRatType() bool`

HasRatType returns a boolean if a field has been set.

### GetLastUeActivityTime

`func (o *TadsInformation) GetLastUeActivityTime() time.Time`

GetLastUeActivityTime returns the LastUeActivityTime field if non-nil, zero value otherwise.

### GetLastUeActivityTimeOk

`func (o *TadsInformation) GetLastUeActivityTimeOk() (*time.Time, bool)`

GetLastUeActivityTimeOk returns a tuple with the LastUeActivityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUeActivityTime

`func (o *TadsInformation) SetLastUeActivityTime(v time.Time)`

SetLastUeActivityTime sets LastUeActivityTime field to given value.

### HasLastUeActivityTime

`func (o *TadsInformation) HasLastUeActivityTime() bool`

HasLastUeActivityTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


