# CipheringDataSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipheringSetID** | **int32** | Ciphering Data Set Identifier. | 
**CipheringKey** | **string** | Ciphering Key. | 
**C0** | **string** | First component of the initial ciphering counter. | 
**LtePosSibTypes** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**NrPosSibTypes** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 
**ValidityStartTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**ValidityDuration** | **int32** | Validity Duration of the Ciphering Data Set. | 
**TaiList** | Pointer to **string** | string with format &#39;bytes&#39; as defined in OpenAPI | [optional] 

## Methods

### NewCipheringDataSet

`func NewCipheringDataSet(cipheringSetID int32, cipheringKey string, c0 string, validityStartTime time.Time, validityDuration int32, ) *CipheringDataSet`

NewCipheringDataSet instantiates a new CipheringDataSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCipheringDataSetWithDefaults

`func NewCipheringDataSetWithDefaults() *CipheringDataSet`

NewCipheringDataSetWithDefaults instantiates a new CipheringDataSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipheringSetID

`func (o *CipheringDataSet) GetCipheringSetID() int32`

GetCipheringSetID returns the CipheringSetID field if non-nil, zero value otherwise.

### GetCipheringSetIDOk

`func (o *CipheringDataSet) GetCipheringSetIDOk() (*int32, bool)`

GetCipheringSetIDOk returns a tuple with the CipheringSetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipheringSetID

`func (o *CipheringDataSet) SetCipheringSetID(v int32)`

SetCipheringSetID sets CipheringSetID field to given value.


### GetCipheringKey

`func (o *CipheringDataSet) GetCipheringKey() string`

GetCipheringKey returns the CipheringKey field if non-nil, zero value otherwise.

### GetCipheringKeyOk

`func (o *CipheringDataSet) GetCipheringKeyOk() (*string, bool)`

GetCipheringKeyOk returns a tuple with the CipheringKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipheringKey

`func (o *CipheringDataSet) SetCipheringKey(v string)`

SetCipheringKey sets CipheringKey field to given value.


### GetC0

`func (o *CipheringDataSet) GetC0() string`

GetC0 returns the C0 field if non-nil, zero value otherwise.

### GetC0Ok

`func (o *CipheringDataSet) GetC0Ok() (*string, bool)`

GetC0Ok returns a tuple with the C0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetC0

`func (o *CipheringDataSet) SetC0(v string)`

SetC0 sets C0 field to given value.


### GetLtePosSibTypes

`func (o *CipheringDataSet) GetLtePosSibTypes() string`

GetLtePosSibTypes returns the LtePosSibTypes field if non-nil, zero value otherwise.

### GetLtePosSibTypesOk

`func (o *CipheringDataSet) GetLtePosSibTypesOk() (*string, bool)`

GetLtePosSibTypesOk returns a tuple with the LtePosSibTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtePosSibTypes

`func (o *CipheringDataSet) SetLtePosSibTypes(v string)`

SetLtePosSibTypes sets LtePosSibTypes field to given value.

### HasLtePosSibTypes

`func (o *CipheringDataSet) HasLtePosSibTypes() bool`

HasLtePosSibTypes returns a boolean if a field has been set.

### GetNrPosSibTypes

`func (o *CipheringDataSet) GetNrPosSibTypes() string`

GetNrPosSibTypes returns the NrPosSibTypes field if non-nil, zero value otherwise.

### GetNrPosSibTypesOk

`func (o *CipheringDataSet) GetNrPosSibTypesOk() (*string, bool)`

GetNrPosSibTypesOk returns a tuple with the NrPosSibTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNrPosSibTypes

`func (o *CipheringDataSet) SetNrPosSibTypes(v string)`

SetNrPosSibTypes sets NrPosSibTypes field to given value.

### HasNrPosSibTypes

`func (o *CipheringDataSet) HasNrPosSibTypes() bool`

HasNrPosSibTypes returns a boolean if a field has been set.

### GetValidityStartTime

`func (o *CipheringDataSet) GetValidityStartTime() time.Time`

GetValidityStartTime returns the ValidityStartTime field if non-nil, zero value otherwise.

### GetValidityStartTimeOk

`func (o *CipheringDataSet) GetValidityStartTimeOk() (*time.Time, bool)`

GetValidityStartTimeOk returns a tuple with the ValidityStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityStartTime

`func (o *CipheringDataSet) SetValidityStartTime(v time.Time)`

SetValidityStartTime sets ValidityStartTime field to given value.


### GetValidityDuration

`func (o *CipheringDataSet) GetValidityDuration() int32`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *CipheringDataSet) GetValidityDurationOk() (*int32, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *CipheringDataSet) SetValidityDuration(v int32)`

SetValidityDuration sets ValidityDuration field to given value.


### GetTaiList

`func (o *CipheringDataSet) GetTaiList() string`

GetTaiList returns the TaiList field if non-nil, zero value otherwise.

### GetTaiListOk

`func (o *CipheringDataSet) GetTaiListOk() (*string, bool)`

GetTaiListOk returns a tuple with the TaiList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaiList

`func (o *CipheringDataSet) SetTaiList(v string)`

SetTaiList sets TaiList field to given value.

### HasTaiList

`func (o *CipheringDataSet) HasTaiList() bool`

HasTaiList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


