# AmTerminationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppAmContextId** | **string** | Contains the Individual application AM context resource identifier related to the termination notification. | 
**TermCause** | [**AmTerminationCause**](AmTerminationCause.md) |  | 

## Methods

### NewAmTerminationInfo

`func NewAmTerminationInfo(appAmContextId string, termCause AmTerminationCause, ) *AmTerminationInfo`

NewAmTerminationInfo instantiates a new AmTerminationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmTerminationInfoWithDefaults

`func NewAmTerminationInfoWithDefaults() *AmTerminationInfo`

NewAmTerminationInfoWithDefaults instantiates a new AmTerminationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppAmContextId

`func (o *AmTerminationInfo) GetAppAmContextId() string`

GetAppAmContextId returns the AppAmContextId field if non-nil, zero value otherwise.

### GetAppAmContextIdOk

`func (o *AmTerminationInfo) GetAppAmContextIdOk() (*string, bool)`

GetAppAmContextIdOk returns a tuple with the AppAmContextId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppAmContextId

`func (o *AmTerminationInfo) SetAppAmContextId(v string)`

SetAppAmContextId sets AppAmContextId field to given value.


### GetTermCause

`func (o *AmTerminationInfo) GetTermCause() AmTerminationCause`

GetTermCause returns the TermCause field if non-nil, zero value otherwise.

### GetTermCauseOk

`func (o *AmTerminationInfo) GetTermCauseOk() (*AmTerminationCause, bool)`

GetTermCauseOk returns a tuple with the TermCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermCause

`func (o *AmTerminationInfo) SetTermCause(v AmTerminationCause)`

SetTermCause sets TermCause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


