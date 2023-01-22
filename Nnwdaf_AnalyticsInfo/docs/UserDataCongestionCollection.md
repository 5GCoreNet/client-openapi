# UserDataCongestionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **string** | String providing an application identifier. | [optional] 
**IpTrafficFilter** | Pointer to [**FlowInfo**](FlowInfo.md) |  | [optional] 
**TimeInterv** | Pointer to [**TimeWindow**](TimeWindow.md) |  | [optional] 
**ThrputUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ThrputDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ThrputPkUl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 
**ThrputPkDl** | Pointer to **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | [optional] 

## Methods

### NewUserDataCongestionCollection

`func NewUserDataCongestionCollection() *UserDataCongestionCollection`

NewUserDataCongestionCollection instantiates a new UserDataCongestionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataCongestionCollectionWithDefaults

`func NewUserDataCongestionCollectionWithDefaults() *UserDataCongestionCollection`

NewUserDataCongestionCollectionWithDefaults instantiates a new UserDataCongestionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *UserDataCongestionCollection) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *UserDataCongestionCollection) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *UserDataCongestionCollection) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *UserDataCongestionCollection) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetIpTrafficFilter

`func (o *UserDataCongestionCollection) GetIpTrafficFilter() FlowInfo`

GetIpTrafficFilter returns the IpTrafficFilter field if non-nil, zero value otherwise.

### GetIpTrafficFilterOk

`func (o *UserDataCongestionCollection) GetIpTrafficFilterOk() (*FlowInfo, bool)`

GetIpTrafficFilterOk returns a tuple with the IpTrafficFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTrafficFilter

`func (o *UserDataCongestionCollection) SetIpTrafficFilter(v FlowInfo)`

SetIpTrafficFilter sets IpTrafficFilter field to given value.

### HasIpTrafficFilter

`func (o *UserDataCongestionCollection) HasIpTrafficFilter() bool`

HasIpTrafficFilter returns a boolean if a field has been set.

### GetTimeInterv

`func (o *UserDataCongestionCollection) GetTimeInterv() TimeWindow`

GetTimeInterv returns the TimeInterv field if non-nil, zero value otherwise.

### GetTimeIntervOk

`func (o *UserDataCongestionCollection) GetTimeIntervOk() (*TimeWindow, bool)`

GetTimeIntervOk returns a tuple with the TimeInterv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInterv

`func (o *UserDataCongestionCollection) SetTimeInterv(v TimeWindow)`

SetTimeInterv sets TimeInterv field to given value.

### HasTimeInterv

`func (o *UserDataCongestionCollection) HasTimeInterv() bool`

HasTimeInterv returns a boolean if a field has been set.

### GetThrputUl

`func (o *UserDataCongestionCollection) GetThrputUl() string`

GetThrputUl returns the ThrputUl field if non-nil, zero value otherwise.

### GetThrputUlOk

`func (o *UserDataCongestionCollection) GetThrputUlOk() (*string, bool)`

GetThrputUlOk returns a tuple with the ThrputUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrputUl

`func (o *UserDataCongestionCollection) SetThrputUl(v string)`

SetThrputUl sets ThrputUl field to given value.

### HasThrputUl

`func (o *UserDataCongestionCollection) HasThrputUl() bool`

HasThrputUl returns a boolean if a field has been set.

### GetThrputDl

`func (o *UserDataCongestionCollection) GetThrputDl() string`

GetThrputDl returns the ThrputDl field if non-nil, zero value otherwise.

### GetThrputDlOk

`func (o *UserDataCongestionCollection) GetThrputDlOk() (*string, bool)`

GetThrputDlOk returns a tuple with the ThrputDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrputDl

`func (o *UserDataCongestionCollection) SetThrputDl(v string)`

SetThrputDl sets ThrputDl field to given value.

### HasThrputDl

`func (o *UserDataCongestionCollection) HasThrputDl() bool`

HasThrputDl returns a boolean if a field has been set.

### GetThrputPkUl

`func (o *UserDataCongestionCollection) GetThrputPkUl() string`

GetThrputPkUl returns the ThrputPkUl field if non-nil, zero value otherwise.

### GetThrputPkUlOk

`func (o *UserDataCongestionCollection) GetThrputPkUlOk() (*string, bool)`

GetThrputPkUlOk returns a tuple with the ThrputPkUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrputPkUl

`func (o *UserDataCongestionCollection) SetThrputPkUl(v string)`

SetThrputPkUl sets ThrputPkUl field to given value.

### HasThrputPkUl

`func (o *UserDataCongestionCollection) HasThrputPkUl() bool`

HasThrputPkUl returns a boolean if a field has been set.

### GetThrputPkDl

`func (o *UserDataCongestionCollection) GetThrputPkDl() string`

GetThrputPkDl returns the ThrputPkDl field if non-nil, zero value otherwise.

### GetThrputPkDlOk

`func (o *UserDataCongestionCollection) GetThrputPkDlOk() (*string, bool)`

GetThrputPkDlOk returns a tuple with the ThrputPkDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrputPkDl

`func (o *UserDataCongestionCollection) SetThrputPkDl(v string)`

SetThrputPkDl sets ThrputPkDl field to given value.

### HasThrputPkDl

`func (o *UserDataCongestionCollection) HasThrputPkDl() bool`

HasThrputPkDl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


