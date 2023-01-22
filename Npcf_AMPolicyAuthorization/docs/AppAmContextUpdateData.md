# AppAmContextUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermNotifUri** | Pointer to **string** | String providing an URI formatted according to RFC 3986. | [optional] 
**EvSubsc** | Pointer to [**NullableAmEventsSubscDataRm**](AmEventsSubscDataRm.md) |  | [optional] 
**Expiry** | Pointer to **NullableInt32** | indicating a time in seconds with OpenAPI defined &#39;nullable: true&#39; property. | [optional] 
**HighThruInd** | Pointer to **NullableBool** | Indicates whether high throughput is desired for the indicated UE traffic. | [optional] 
**CovReq** | Pointer to [**[]ServiceAreaCoverageInfo**](ServiceAreaCoverageInfo.md) | Identifies a list of Tracking Areas per serving network where service is allowed. | [optional] 
**AsTimeDisParam** | Pointer to [**NullableAsTimeDistributionParam**](AsTimeDistributionParam.md) |  | [optional] 

## Methods

### NewAppAmContextUpdateData

`func NewAppAmContextUpdateData() *AppAmContextUpdateData`

NewAppAmContextUpdateData instantiates a new AppAmContextUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAmContextUpdateDataWithDefaults

`func NewAppAmContextUpdateDataWithDefaults() *AppAmContextUpdateData`

NewAppAmContextUpdateDataWithDefaults instantiates a new AppAmContextUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermNotifUri

`func (o *AppAmContextUpdateData) GetTermNotifUri() string`

GetTermNotifUri returns the TermNotifUri field if non-nil, zero value otherwise.

### GetTermNotifUriOk

`func (o *AppAmContextUpdateData) GetTermNotifUriOk() (*string, bool)`

GetTermNotifUriOk returns a tuple with the TermNotifUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermNotifUri

`func (o *AppAmContextUpdateData) SetTermNotifUri(v string)`

SetTermNotifUri sets TermNotifUri field to given value.

### HasTermNotifUri

`func (o *AppAmContextUpdateData) HasTermNotifUri() bool`

HasTermNotifUri returns a boolean if a field has been set.

### GetEvSubsc

`func (o *AppAmContextUpdateData) GetEvSubsc() AmEventsSubscDataRm`

GetEvSubsc returns the EvSubsc field if non-nil, zero value otherwise.

### GetEvSubscOk

`func (o *AppAmContextUpdateData) GetEvSubscOk() (*AmEventsSubscDataRm, bool)`

GetEvSubscOk returns a tuple with the EvSubsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvSubsc

`func (o *AppAmContextUpdateData) SetEvSubsc(v AmEventsSubscDataRm)`

SetEvSubsc sets EvSubsc field to given value.

### HasEvSubsc

`func (o *AppAmContextUpdateData) HasEvSubsc() bool`

HasEvSubsc returns a boolean if a field has been set.

### SetEvSubscNil

`func (o *AppAmContextUpdateData) SetEvSubscNil(b bool)`

 SetEvSubscNil sets the value for EvSubsc to be an explicit nil

### UnsetEvSubsc
`func (o *AppAmContextUpdateData) UnsetEvSubsc()`

UnsetEvSubsc ensures that no value is present for EvSubsc, not even an explicit nil
### GetExpiry

`func (o *AppAmContextUpdateData) GetExpiry() int32`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AppAmContextUpdateData) GetExpiryOk() (*int32, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AppAmContextUpdateData) SetExpiry(v int32)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AppAmContextUpdateData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *AppAmContextUpdateData) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *AppAmContextUpdateData) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetHighThruInd

`func (o *AppAmContextUpdateData) GetHighThruInd() bool`

GetHighThruInd returns the HighThruInd field if non-nil, zero value otherwise.

### GetHighThruIndOk

`func (o *AppAmContextUpdateData) GetHighThruIndOk() (*bool, bool)`

GetHighThruIndOk returns a tuple with the HighThruInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighThruInd

`func (o *AppAmContextUpdateData) SetHighThruInd(v bool)`

SetHighThruInd sets HighThruInd field to given value.

### HasHighThruInd

`func (o *AppAmContextUpdateData) HasHighThruInd() bool`

HasHighThruInd returns a boolean if a field has been set.

### SetHighThruIndNil

`func (o *AppAmContextUpdateData) SetHighThruIndNil(b bool)`

 SetHighThruIndNil sets the value for HighThruInd to be an explicit nil

### UnsetHighThruInd
`func (o *AppAmContextUpdateData) UnsetHighThruInd()`

UnsetHighThruInd ensures that no value is present for HighThruInd, not even an explicit nil
### GetCovReq

`func (o *AppAmContextUpdateData) GetCovReq() []ServiceAreaCoverageInfo`

GetCovReq returns the CovReq field if non-nil, zero value otherwise.

### GetCovReqOk

`func (o *AppAmContextUpdateData) GetCovReqOk() (*[]ServiceAreaCoverageInfo, bool)`

GetCovReqOk returns a tuple with the CovReq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovReq

`func (o *AppAmContextUpdateData) SetCovReq(v []ServiceAreaCoverageInfo)`

SetCovReq sets CovReq field to given value.

### HasCovReq

`func (o *AppAmContextUpdateData) HasCovReq() bool`

HasCovReq returns a boolean if a field has been set.

### SetCovReqNil

`func (o *AppAmContextUpdateData) SetCovReqNil(b bool)`

 SetCovReqNil sets the value for CovReq to be an explicit nil

### UnsetCovReq
`func (o *AppAmContextUpdateData) UnsetCovReq()`

UnsetCovReq ensures that no value is present for CovReq, not even an explicit nil
### GetAsTimeDisParam

`func (o *AppAmContextUpdateData) GetAsTimeDisParam() AsTimeDistributionParam`

GetAsTimeDisParam returns the AsTimeDisParam field if non-nil, zero value otherwise.

### GetAsTimeDisParamOk

`func (o *AppAmContextUpdateData) GetAsTimeDisParamOk() (*AsTimeDistributionParam, bool)`

GetAsTimeDisParamOk returns a tuple with the AsTimeDisParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsTimeDisParam

`func (o *AppAmContextUpdateData) SetAsTimeDisParam(v AsTimeDistributionParam)`

SetAsTimeDisParam sets AsTimeDisParam field to given value.

### HasAsTimeDisParam

`func (o *AppAmContextUpdateData) HasAsTimeDisParam() bool`

HasAsTimeDisParam returns a boolean if a field has been set.

### SetAsTimeDisParamNil

`func (o *AppAmContextUpdateData) SetAsTimeDisParamNil(b bool)`

 SetAsTimeDisParamNil sets the value for AsTimeDisParam to be an explicit nil

### UnsetAsTimeDisParam
`func (o *AppAmContextUpdateData) UnsetAsTimeDisParam()`

UnsetAsTimeDisParam ensures that no value is present for AsTimeDisParam, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


