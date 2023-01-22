# RoamingStatusUpdateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Imsi** | **string** |  | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 

## Methods

### NewRoamingStatusUpdateInfo

`func NewRoamingStatusUpdateInfo(imsi string, plmnId PlmnId, ) *RoamingStatusUpdateInfo`

NewRoamingStatusUpdateInfo instantiates a new RoamingStatusUpdateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoamingStatusUpdateInfoWithDefaults

`func NewRoamingStatusUpdateInfoWithDefaults() *RoamingStatusUpdateInfo`

NewRoamingStatusUpdateInfoWithDefaults instantiates a new RoamingStatusUpdateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImsi

`func (o *RoamingStatusUpdateInfo) GetImsi() string`

GetImsi returns the Imsi field if non-nil, zero value otherwise.

### GetImsiOk

`func (o *RoamingStatusUpdateInfo) GetImsiOk() (*string, bool)`

GetImsiOk returns a tuple with the Imsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImsi

`func (o *RoamingStatusUpdateInfo) SetImsi(v string)`

SetImsi sets Imsi field to given value.


### GetPlmnId

`func (o *RoamingStatusUpdateInfo) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *RoamingStatusUpdateInfo) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *RoamingStatusUpdateInfo) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


