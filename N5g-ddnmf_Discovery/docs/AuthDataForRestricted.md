# AuthDataForRestricted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProseQueryCodes** | **[]string** |  | 
**ProseRespCode** | **string** | Contains the ProSe Response Code. | 
**ValidityTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 

## Methods

### NewAuthDataForRestricted

`func NewAuthDataForRestricted(proseQueryCodes []string, proseRespCode string, validityTime time.Time, ) *AuthDataForRestricted`

NewAuthDataForRestricted instantiates a new AuthDataForRestricted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthDataForRestrictedWithDefaults

`func NewAuthDataForRestrictedWithDefaults() *AuthDataForRestricted`

NewAuthDataForRestrictedWithDefaults instantiates a new AuthDataForRestricted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProseQueryCodes

`func (o *AuthDataForRestricted) GetProseQueryCodes() []string`

GetProseQueryCodes returns the ProseQueryCodes field if non-nil, zero value otherwise.

### GetProseQueryCodesOk

`func (o *AuthDataForRestricted) GetProseQueryCodesOk() (*[]string, bool)`

GetProseQueryCodesOk returns a tuple with the ProseQueryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseQueryCodes

`func (o *AuthDataForRestricted) SetProseQueryCodes(v []string)`

SetProseQueryCodes sets ProseQueryCodes field to given value.


### GetProseRespCode

`func (o *AuthDataForRestricted) GetProseRespCode() string`

GetProseRespCode returns the ProseRespCode field if non-nil, zero value otherwise.

### GetProseRespCodeOk

`func (o *AuthDataForRestricted) GetProseRespCodeOk() (*string, bool)`

GetProseRespCodeOk returns a tuple with the ProseRespCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseRespCode

`func (o *AuthDataForRestricted) SetProseRespCode(v string)`

SetProseRespCode sets ProseRespCode field to given value.


### GetValidityTime

`func (o *AuthDataForRestricted) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *AuthDataForRestricted) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *AuthDataForRestricted) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


