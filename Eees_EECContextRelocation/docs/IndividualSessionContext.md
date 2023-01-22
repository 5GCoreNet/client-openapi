# IndividualSessionContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EasId** | **string** | Identifier of the EAS providing the application services. | 
**EndPt** | [**EndPoint**](EndPoint.md) |  | 
**AcId** | Pointer to **string** | Identifier of the AC for which the service session information is provided. | [optional] 

## Methods

### NewIndividualSessionContext

`func NewIndividualSessionContext(easId string, endPt EndPoint, ) *IndividualSessionContext`

NewIndividualSessionContext instantiates a new IndividualSessionContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndividualSessionContextWithDefaults

`func NewIndividualSessionContextWithDefaults() *IndividualSessionContext`

NewIndividualSessionContextWithDefaults instantiates a new IndividualSessionContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEasId

`func (o *IndividualSessionContext) GetEasId() string`

GetEasId returns the EasId field if non-nil, zero value otherwise.

### GetEasIdOk

`func (o *IndividualSessionContext) GetEasIdOk() (*string, bool)`

GetEasIdOk returns a tuple with the EasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasId

`func (o *IndividualSessionContext) SetEasId(v string)`

SetEasId sets EasId field to given value.


### GetEndPt

`func (o *IndividualSessionContext) GetEndPt() EndPoint`

GetEndPt returns the EndPt field if non-nil, zero value otherwise.

### GetEndPtOk

`func (o *IndividualSessionContext) GetEndPtOk() (*EndPoint, bool)`

GetEndPtOk returns a tuple with the EndPt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPt

`func (o *IndividualSessionContext) SetEndPt(v EndPoint)`

SetEndPt sets EndPt field to given value.


### GetAcId

`func (o *IndividualSessionContext) GetAcId() string`

GetAcId returns the AcId field if non-nil, zero value otherwise.

### GetAcIdOk

`func (o *IndividualSessionContext) GetAcIdOk() (*string, bool)`

GetAcIdOk returns a tuple with the AcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcId

`func (o *IndividualSessionContext) SetAcId(v string)`

SetAcId sets AcId field to given value.

### HasAcId

`func (o *IndividualSessionContext) HasAcId() bool`

HasAcId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


