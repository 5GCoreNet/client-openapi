# MbsServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsMediaComps** | [**map[string]MbsMediaCompRm**](MbsMediaCompRm.md) |  | 
**MbsSdfResPrio** | Pointer to [**ReservPriority**](ReservPriority.md) |  | [optional] 
**AfAppId** | Pointer to **string** | Contains an AF application identifier. | [optional] 
**MbsSessionAmbr** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewMbsServiceInfo

`func NewMbsServiceInfo(mbsMediaComps map[string]MbsMediaCompRm, ) *MbsServiceInfo`

NewMbsServiceInfo instantiates a new MbsServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsServiceInfoWithDefaults

`func NewMbsServiceInfoWithDefaults() *MbsServiceInfo`

NewMbsServiceInfoWithDefaults instantiates a new MbsServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsMediaComps

`func (o *MbsServiceInfo) GetMbsMediaComps() map[string]MbsMediaCompRm`

GetMbsMediaComps returns the MbsMediaComps field if non-nil, zero value otherwise.

### GetMbsMediaCompsOk

`func (o *MbsServiceInfo) GetMbsMediaCompsOk() (*map[string]MbsMediaCompRm, bool)`

GetMbsMediaCompsOk returns a tuple with the MbsMediaComps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsMediaComps

`func (o *MbsServiceInfo) SetMbsMediaComps(v map[string]MbsMediaCompRm)`

SetMbsMediaComps sets MbsMediaComps field to given value.


### GetMbsSdfResPrio

`func (o *MbsServiceInfo) GetMbsSdfResPrio() ReservPriority`

GetMbsSdfResPrio returns the MbsSdfResPrio field if non-nil, zero value otherwise.

### GetMbsSdfResPrioOk

`func (o *MbsServiceInfo) GetMbsSdfResPrioOk() (*ReservPriority, bool)`

GetMbsSdfResPrioOk returns a tuple with the MbsSdfResPrio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSdfResPrio

`func (o *MbsServiceInfo) SetMbsSdfResPrio(v ReservPriority)`

SetMbsSdfResPrio sets MbsSdfResPrio field to given value.

### HasMbsSdfResPrio

`func (o *MbsServiceInfo) HasMbsSdfResPrio() bool`

HasMbsSdfResPrio returns a boolean if a field has been set.

### GetAfAppId

`func (o *MbsServiceInfo) GetAfAppId() string`

GetAfAppId returns the AfAppId field if non-nil, zero value otherwise.

### GetAfAppIdOk

`func (o *MbsServiceInfo) GetAfAppIdOk() (*string, bool)`

GetAfAppIdOk returns a tuple with the AfAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfAppId

`func (o *MbsServiceInfo) SetAfAppId(v string)`

SetAfAppId sets AfAppId field to given value.

### HasAfAppId

`func (o *MbsServiceInfo) HasAfAppId() bool`

HasAfAppId returns a boolean if a field has been set.

### GetMbsSessionAmbr

`func (o *MbsServiceInfo) GetMbsSessionAmbr() string`

GetMbsSessionAmbr returns the MbsSessionAmbr field if non-nil, zero value otherwise.

### GetMbsSessionAmbrOk

`func (o *MbsServiceInfo) GetMbsSessionAmbrOk() (*string, bool)`

GetMbsSessionAmbrOk returns a tuple with the MbsSessionAmbr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionAmbr

`func (o *MbsServiceInfo) SetMbsSessionAmbr(v string)`

SetMbsSessionAmbr sets MbsSessionAmbr field to given value.

### HasMbsSessionAmbr

`func (o *MbsServiceInfo) HasMbsSessionAmbr() bool`

HasMbsSessionAmbr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


