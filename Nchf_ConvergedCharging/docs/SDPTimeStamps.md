# SDPTimeStamps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SDPOfferTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 
**SDPAnswerTimestamp** | Pointer to **time.Time** | string with format &#39;date-time&#39; as defined in OpenAPI. | [optional] 

## Methods

### NewSDPTimeStamps

`func NewSDPTimeStamps() *SDPTimeStamps`

NewSDPTimeStamps instantiates a new SDPTimeStamps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDPTimeStampsWithDefaults

`func NewSDPTimeStampsWithDefaults() *SDPTimeStamps`

NewSDPTimeStampsWithDefaults instantiates a new SDPTimeStamps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSDPOfferTimestamp

`func (o *SDPTimeStamps) GetSDPOfferTimestamp() time.Time`

GetSDPOfferTimestamp returns the SDPOfferTimestamp field if non-nil, zero value otherwise.

### GetSDPOfferTimestampOk

`func (o *SDPTimeStamps) GetSDPOfferTimestampOk() (*time.Time, bool)`

GetSDPOfferTimestampOk returns a tuple with the SDPOfferTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPOfferTimestamp

`func (o *SDPTimeStamps) SetSDPOfferTimestamp(v time.Time)`

SetSDPOfferTimestamp sets SDPOfferTimestamp field to given value.

### HasSDPOfferTimestamp

`func (o *SDPTimeStamps) HasSDPOfferTimestamp() bool`

HasSDPOfferTimestamp returns a boolean if a field has been set.

### GetSDPAnswerTimestamp

`func (o *SDPTimeStamps) GetSDPAnswerTimestamp() time.Time`

GetSDPAnswerTimestamp returns the SDPAnswerTimestamp field if non-nil, zero value otherwise.

### GetSDPAnswerTimestampOk

`func (o *SDPTimeStamps) GetSDPAnswerTimestampOk() (*time.Time, bool)`

GetSDPAnswerTimestampOk returns a tuple with the SDPAnswerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPAnswerTimestamp

`func (o *SDPTimeStamps) SetSDPAnswerTimestamp(v time.Time)`

SetSDPAnswerTimestamp sets SDPAnswerTimestamp field to given value.

### HasSDPAnswerTimestamp

`func (o *SDPTimeStamps) HasSDPAnswerTimestamp() bool`

HasSDPAnswerTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


