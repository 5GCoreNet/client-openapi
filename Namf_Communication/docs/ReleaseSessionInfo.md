# ReleaseSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseSessionList** | **[]int32** |  | 
**ReleaseCause** | [**ReleaseCause**](ReleaseCause.md) |  | 

## Methods

### NewReleaseSessionInfo

`func NewReleaseSessionInfo(releaseSessionList []int32, releaseCause ReleaseCause, ) *ReleaseSessionInfo`

NewReleaseSessionInfo instantiates a new ReleaseSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseSessionInfoWithDefaults

`func NewReleaseSessionInfoWithDefaults() *ReleaseSessionInfo`

NewReleaseSessionInfoWithDefaults instantiates a new ReleaseSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseSessionList

`func (o *ReleaseSessionInfo) GetReleaseSessionList() []int32`

GetReleaseSessionList returns the ReleaseSessionList field if non-nil, zero value otherwise.

### GetReleaseSessionListOk

`func (o *ReleaseSessionInfo) GetReleaseSessionListOk() (*[]int32, bool)`

GetReleaseSessionListOk returns a tuple with the ReleaseSessionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSessionList

`func (o *ReleaseSessionInfo) SetReleaseSessionList(v []int32)`

SetReleaseSessionList sets ReleaseSessionList field to given value.


### GetReleaseCause

`func (o *ReleaseSessionInfo) GetReleaseCause() ReleaseCause`

GetReleaseCause returns the ReleaseCause field if non-nil, zero value otherwise.

### GetReleaseCauseOk

`func (o *ReleaseSessionInfo) GetReleaseCauseOk() (*ReleaseCause, bool)`

GetReleaseCauseOk returns a tuple with the ReleaseCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCause

`func (o *ReleaseSessionInfo) SetReleaseCause(v ReleaseCause)`

SetReleaseCause sets ReleaseCause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


