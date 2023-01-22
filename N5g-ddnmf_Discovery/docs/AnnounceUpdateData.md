# AnnounceUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscType** | [**DiscoveryType**](DiscoveryType.md) |  | 
**ValidityTime** | **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | 
**ProseAppCode** | Pointer to **string** | Contains the ProSe Application Code. | [optional] 

## Methods

### NewAnnounceUpdateData

`func NewAnnounceUpdateData(discType DiscoveryType, validityTime time.Time, ) *AnnounceUpdateData`

NewAnnounceUpdateData instantiates a new AnnounceUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnounceUpdateDataWithDefaults

`func NewAnnounceUpdateDataWithDefaults() *AnnounceUpdateData`

NewAnnounceUpdateDataWithDefaults instantiates a new AnnounceUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscType

`func (o *AnnounceUpdateData) GetDiscType() DiscoveryType`

GetDiscType returns the DiscType field if non-nil, zero value otherwise.

### GetDiscTypeOk

`func (o *AnnounceUpdateData) GetDiscTypeOk() (*DiscoveryType, bool)`

GetDiscTypeOk returns a tuple with the DiscType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscType

`func (o *AnnounceUpdateData) SetDiscType(v DiscoveryType)`

SetDiscType sets DiscType field to given value.


### GetValidityTime

`func (o *AnnounceUpdateData) GetValidityTime() time.Time`

GetValidityTime returns the ValidityTime field if non-nil, zero value otherwise.

### GetValidityTimeOk

`func (o *AnnounceUpdateData) GetValidityTimeOk() (*time.Time, bool)`

GetValidityTimeOk returns a tuple with the ValidityTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityTime

`func (o *AnnounceUpdateData) SetValidityTime(v time.Time)`

SetValidityTime sets ValidityTime field to given value.


### GetProseAppCode

`func (o *AnnounceUpdateData) GetProseAppCode() string`

GetProseAppCode returns the ProseAppCode field if non-nil, zero value otherwise.

### GetProseAppCodeOk

`func (o *AnnounceUpdateData) GetProseAppCodeOk() (*string, bool)`

GetProseAppCodeOk returns a tuple with the ProseAppCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProseAppCode

`func (o *AnnounceUpdateData) SetProseAppCode(v string)`

SetProseAppCode sets ProseAppCode field to given value.

### HasProseAppCode

`func (o *AnnounceUpdateData) HasProseAppCode() bool`

HasProseAppCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


