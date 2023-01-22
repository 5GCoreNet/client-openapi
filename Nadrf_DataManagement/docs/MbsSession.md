# MbsSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**MbsAreaSessions** | Pointer to [**map[string]MbsServiceAreaInfo**](MbsServiceAreaInfo.md) | A map (list of key-value pairs) where the key identifies an areaSessionId | [optional] 

## Methods

### NewMbsSession

`func NewMbsSession(mbsSessionId MbsSessionId, ) *MbsSession`

NewMbsSession instantiates a new MbsSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsSessionWithDefaults

`func NewMbsSessionWithDefaults() *MbsSession`

NewMbsSessionWithDefaults instantiates a new MbsSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *MbsSession) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MbsSession) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MbsSession) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetMbsAreaSessions

`func (o *MbsSession) GetMbsAreaSessions() map[string]MbsServiceAreaInfo`

GetMbsAreaSessions returns the MbsAreaSessions field if non-nil, zero value otherwise.

### GetMbsAreaSessionsOk

`func (o *MbsSession) GetMbsAreaSessionsOk() (*map[string]MbsServiceAreaInfo, bool)`

GetMbsAreaSessionsOk returns a tuple with the MbsAreaSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsAreaSessions

`func (o *MbsSession) SetMbsAreaSessions(v map[string]MbsServiceAreaInfo)`

SetMbsAreaSessions sets MbsAreaSessions field to given value.

### HasMbsAreaSessions

`func (o *MbsSession) HasMbsAreaSessions() bool`

HasMbsAreaSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


