# ReauthorizationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**RatingGroup** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.  | [optional] 
**QuotaManagementIndicator** | Pointer to [**QuotaManagementIndicator**](QuotaManagementIndicator.md) |  | [optional] 

## Methods

### NewReauthorizationDetails

`func NewReauthorizationDetails() *ReauthorizationDetails`

NewReauthorizationDetails instantiates a new ReauthorizationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReauthorizationDetailsWithDefaults

`func NewReauthorizationDetailsWithDefaults() *ReauthorizationDetails`

NewReauthorizationDetailsWithDefaults instantiates a new ReauthorizationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ReauthorizationDetails) GetServiceId() int32`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ReauthorizationDetails) GetServiceIdOk() (*int32, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ReauthorizationDetails) SetServiceId(v int32)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ReauthorizationDetails) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRatingGroup

`func (o *ReauthorizationDetails) GetRatingGroup() int32`

GetRatingGroup returns the RatingGroup field if non-nil, zero value otherwise.

### GetRatingGroupOk

`func (o *ReauthorizationDetails) GetRatingGroupOk() (*int32, bool)`

GetRatingGroupOk returns a tuple with the RatingGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingGroup

`func (o *ReauthorizationDetails) SetRatingGroup(v int32)`

SetRatingGroup sets RatingGroup field to given value.

### HasRatingGroup

`func (o *ReauthorizationDetails) HasRatingGroup() bool`

HasRatingGroup returns a boolean if a field has been set.

### GetQuotaManagementIndicator

`func (o *ReauthorizationDetails) GetQuotaManagementIndicator() QuotaManagementIndicator`

GetQuotaManagementIndicator returns the QuotaManagementIndicator field if non-nil, zero value otherwise.

### GetQuotaManagementIndicatorOk

`func (o *ReauthorizationDetails) GetQuotaManagementIndicatorOk() (*QuotaManagementIndicator, bool)`

GetQuotaManagementIndicatorOk returns a tuple with the QuotaManagementIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaManagementIndicator

`func (o *ReauthorizationDetails) SetQuotaManagementIndicator(v QuotaManagementIndicator)`

SetQuotaManagementIndicator sets QuotaManagementIndicator field to given value.

### HasQuotaManagementIndicator

`func (o *ReauthorizationDetails) HasQuotaManagementIndicator() bool`

HasQuotaManagementIndicator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


