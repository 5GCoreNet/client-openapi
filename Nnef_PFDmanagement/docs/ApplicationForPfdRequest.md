# ApplicationForPfdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | String providing an application identifier. | 
**PfdTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewApplicationForPfdRequest

`func NewApplicationForPfdRequest(applicationId string, ) *ApplicationForPfdRequest`

NewApplicationForPfdRequest instantiates a new ApplicationForPfdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationForPfdRequestWithDefaults

`func NewApplicationForPfdRequestWithDefaults() *ApplicationForPfdRequest`

NewApplicationForPfdRequestWithDefaults instantiates a new ApplicationForPfdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationForPfdRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationForPfdRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationForPfdRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetPfdTimestamp

`func (o *ApplicationForPfdRequest) GetPfdTimestamp() time.Time`

GetPfdTimestamp returns the PfdTimestamp field if non-nil, zero value otherwise.

### GetPfdTimestampOk

`func (o *ApplicationForPfdRequest) GetPfdTimestampOk() (*time.Time, bool)`

GetPfdTimestampOk returns a tuple with the PfdTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfdTimestamp

`func (o *ApplicationForPfdRequest) SetPfdTimestamp(v time.Time)`

SetPfdTimestamp sets PfdTimestamp field to given value.

### HasPfdTimestamp

`func (o *ApplicationForPfdRequest) HasPfdTimestamp() bool`

HasPfdTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


