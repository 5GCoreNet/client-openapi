# AmInfluDataPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppIds** | Pointer to **[]string** | Identifies one or more applications. | [optional] 
**DnnSnssaiInfos** | Pointer to [**[]DnnSnssaiInformation**](DnnSnssaiInformation.md) | Identifies one or more DNN, S-NSSAI combinations. | [optional] 
**EvSubs** | Pointer to [**[]AmInfluEvent**](AmInfluEvent.md) | List of AM related events for which a subscription is required. | [optional] 
**Headers** | Pointer to **[]string** | Contains the headers provisioned by the NEF. | [optional] 
**ThruReq** | Pointer to **NullableBool** | Indicates whether high throughput is desired for the indicated UE traffic. | [optional] 
**NotifUri** | Pointer to **NullableString** | String providing an URI formatted according to RFC 3986 with the OpenAPI &#39;nullable: true&#39; property.  | [optional] 
**NotifCorrId** | Pointer to **NullableString** | Notification correlation identifier. | [optional] 
**CovReq** | Pointer to [**[]ServiceAreaCoverageInfo**](ServiceAreaCoverageInfo.md) | Indicates the service area coverage requirement. | [optional] 

## Methods

### NewAmInfluDataPatch

`func NewAmInfluDataPatch() *AmInfluDataPatch`

NewAmInfluDataPatch instantiates a new AmInfluDataPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmInfluDataPatchWithDefaults

`func NewAmInfluDataPatchWithDefaults() *AmInfluDataPatch`

NewAmInfluDataPatchWithDefaults instantiates a new AmInfluDataPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppIds

`func (o *AmInfluDataPatch) GetAppIds() []string`

GetAppIds returns the AppIds field if non-nil, zero value otherwise.

### GetAppIdsOk

`func (o *AmInfluDataPatch) GetAppIdsOk() (*[]string, bool)`

GetAppIdsOk returns a tuple with the AppIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppIds

`func (o *AmInfluDataPatch) SetAppIds(v []string)`

SetAppIds sets AppIds field to given value.

### HasAppIds

`func (o *AmInfluDataPatch) HasAppIds() bool`

HasAppIds returns a boolean if a field has been set.

### SetAppIdsNil

`func (o *AmInfluDataPatch) SetAppIdsNil(b bool)`

 SetAppIdsNil sets the value for AppIds to be an explicit nil

### UnsetAppIds
`func (o *AmInfluDataPatch) UnsetAppIds()`

UnsetAppIds ensures that no value is present for AppIds, not even an explicit nil
### GetDnnSnssaiInfos

`func (o *AmInfluDataPatch) GetDnnSnssaiInfos() []DnnSnssaiInformation`

GetDnnSnssaiInfos returns the DnnSnssaiInfos field if non-nil, zero value otherwise.

### GetDnnSnssaiInfosOk

`func (o *AmInfluDataPatch) GetDnnSnssaiInfosOk() (*[]DnnSnssaiInformation, bool)`

GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnSnssaiInfos

`func (o *AmInfluDataPatch) SetDnnSnssaiInfos(v []DnnSnssaiInformation)`

SetDnnSnssaiInfos sets DnnSnssaiInfos field to given value.

### HasDnnSnssaiInfos

`func (o *AmInfluDataPatch) HasDnnSnssaiInfos() bool`

HasDnnSnssaiInfos returns a boolean if a field has been set.

### SetDnnSnssaiInfosNil

`func (o *AmInfluDataPatch) SetDnnSnssaiInfosNil(b bool)`

 SetDnnSnssaiInfosNil sets the value for DnnSnssaiInfos to be an explicit nil

### UnsetDnnSnssaiInfos
`func (o *AmInfluDataPatch) UnsetDnnSnssaiInfos()`

UnsetDnnSnssaiInfos ensures that no value is present for DnnSnssaiInfos, not even an explicit nil
### GetEvSubs

`func (o *AmInfluDataPatch) GetEvSubs() []AmInfluEvent`

GetEvSubs returns the EvSubs field if non-nil, zero value otherwise.

### GetEvSubsOk

`func (o *AmInfluDataPatch) GetEvSubsOk() (*[]AmInfluEvent, bool)`

GetEvSubsOk returns a tuple with the EvSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubs

`func (o *AmInfluDataPatch) SetEvSubs(v []AmInfluEvent)`

SetEvSubs sets EvSubs field to given value.

### HasEvSubs

`func (o *AmInfluDataPatch) HasEvSubs() bool`

HasEvSubs returns a boolean if a field has been set.

### SetEvSubsNil

`func (o *AmInfluDataPatch) SetEvSubsNil(b bool)`

 SetEvSubsNil sets the value for EvSubs to be an explicit nil

### UnsetEvSubs
`func (o *AmInfluDataPatch) UnsetEvSubs()`

UnsetEvSubs ensures that no value is present for EvSubs, not even an explicit nil
### GetHeaders

`func (o *AmInfluDataPatch) GetHeaders() []string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AmInfluDataPatch) GetHeadersOk() (*[]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AmInfluDataPatch) SetHeaders(v []string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AmInfluDataPatch) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetThruReq

`func (o *AmInfluDataPatch) GetThruReq() bool`

GetThruReq returns the ThruReq field if non-nil, zero value otherwise.

### GetThruReqOk

`func (o *AmInfluDataPatch) GetThruReqOk() (*bool, bool)`

GetThruReqOk returns a tuple with the ThruReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThruReq

`func (o *AmInfluDataPatch) SetThruReq(v bool)`

SetThruReq sets ThruReq field to given value.

### HasThruReq

`func (o *AmInfluDataPatch) HasThruReq() bool`

HasThruReq returns a boolean if a field has been set.

### SetThruReqNil

`func (o *AmInfluDataPatch) SetThruReqNil(b bool)`

 SetThruReqNil sets the value for ThruReq to be an explicit nil

### UnsetThruReq
`func (o *AmInfluDataPatch) UnsetThruReq()`

UnsetThruReq ensures that no value is present for ThruReq, not even an explicit nil
### GetNotifUri

`func (o *AmInfluDataPatch) GetNotifUri() string`

GetNotifUri returns the NotifUri field if non-nil, zero value otherwise.

### GetNotifUriOk

`func (o *AmInfluDataPatch) GetNotifUriOk() (*string, bool)`

GetNotifUriOk returns a tuple with the NotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifUri

`func (o *AmInfluDataPatch) SetNotifUri(v string)`

SetNotifUri sets NotifUri field to given value.

### HasNotifUri

`func (o *AmInfluDataPatch) HasNotifUri() bool`

HasNotifUri returns a boolean if a field has been set.

### SetNotifUriNil

`func (o *AmInfluDataPatch) SetNotifUriNil(b bool)`

 SetNotifUriNil sets the value for NotifUri to be an explicit nil

### UnsetNotifUri
`func (o *AmInfluDataPatch) UnsetNotifUri()`

UnsetNotifUri ensures that no value is present for NotifUri, not even an explicit nil
### GetNotifCorrId

`func (o *AmInfluDataPatch) GetNotifCorrId() string`

GetNotifCorrId returns the NotifCorrId field if non-nil, zero value otherwise.

### GetNotifCorrIdOk

`func (o *AmInfluDataPatch) GetNotifCorrIdOk() (*string, bool)`

GetNotifCorrIdOk returns a tuple with the NotifCorrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifCorrId

`func (o *AmInfluDataPatch) SetNotifCorrId(v string)`

SetNotifCorrId sets NotifCorrId field to given value.

### HasNotifCorrId

`func (o *AmInfluDataPatch) HasNotifCorrId() bool`

HasNotifCorrId returns a boolean if a field has been set.

### SetNotifCorrIdNil

`func (o *AmInfluDataPatch) SetNotifCorrIdNil(b bool)`

 SetNotifCorrIdNil sets the value for NotifCorrId to be an explicit nil

### UnsetNotifCorrId
`func (o *AmInfluDataPatch) UnsetNotifCorrId()`

UnsetNotifCorrId ensures that no value is present for NotifCorrId, not even an explicit nil
### GetCovReq

`func (o *AmInfluDataPatch) GetCovReq() []ServiceAreaCoverageInfo`

GetCovReq returns the CovReq field if non-nil, zero value otherwise.

### GetCovReqOk

`func (o *AmInfluDataPatch) GetCovReqOk() (*[]ServiceAreaCoverageInfo, bool)`

GetCovReqOk returns a tuple with the CovReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovReq

`func (o *AmInfluDataPatch) SetCovReq(v []ServiceAreaCoverageInfo)`

SetCovReq sets CovReq field to given value.

### HasCovReq

`func (o *AmInfluDataPatch) HasCovReq() bool`

HasCovReq returns a boolean if a field has been set.

### SetCovReqNil

`func (o *AmInfluDataPatch) SetCovReqNil(b bool)`

 SetCovReqNil sets the value for CovReq to be an explicit nil

### UnsetCovReq
`func (o *AmInfluDataPatch) UnsetCovReq()`

UnsetCovReq ensures that no value is present for CovReq, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


