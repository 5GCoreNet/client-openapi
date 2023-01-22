# SipAuthenticationInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CscfServerName** | **string** |  | 
**SipAuthenticationScheme** | [**SipAuthenticationScheme**](SipAuthenticationScheme.md) |  | 
**SipNumberAuthItems** | Pointer to **int32** | Indicates the number of requested SIP authentication items | [optional] 
**ResynchronizationInfo** | Pointer to [**ResynchronizationInfo**](ResynchronizationInfo.md) |  | [optional] 

## Methods

### NewSipAuthenticationInfoRequest

`func NewSipAuthenticationInfoRequest(cscfServerName string, sipAuthenticationScheme SipAuthenticationScheme, ) *SipAuthenticationInfoRequest`

NewSipAuthenticationInfoRequest instantiates a new SipAuthenticationInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSipAuthenticationInfoRequestWithDefaults

`func NewSipAuthenticationInfoRequestWithDefaults() *SipAuthenticationInfoRequest`

NewSipAuthenticationInfoRequestWithDefaults instantiates a new SipAuthenticationInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCscfServerName

`func (o *SipAuthenticationInfoRequest) GetCscfServerName() string`

GetCscfServerName returns the CscfServerName field if non-nil, zero value otherwise.

### GetCscfServerNameOk

`func (o *SipAuthenticationInfoRequest) GetCscfServerNameOk() (*string, bool)`

GetCscfServerNameOk returns a tuple with the CscfServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCscfServerName

`func (o *SipAuthenticationInfoRequest) SetCscfServerName(v string)`

SetCscfServerName sets CscfServerName field to given value.


### GetSipAuthenticationScheme

`func (o *SipAuthenticationInfoRequest) GetSipAuthenticationScheme() SipAuthenticationScheme`

GetSipAuthenticationScheme returns the SipAuthenticationScheme field if non-nil, zero value otherwise.

### GetSipAuthenticationSchemeOk

`func (o *SipAuthenticationInfoRequest) GetSipAuthenticationSchemeOk() (*SipAuthenticationScheme, bool)`

GetSipAuthenticationSchemeOk returns a tuple with the SipAuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthenticationScheme

`func (o *SipAuthenticationInfoRequest) SetSipAuthenticationScheme(v SipAuthenticationScheme)`

SetSipAuthenticationScheme sets SipAuthenticationScheme field to given value.


### GetSipNumberAuthItems

`func (o *SipAuthenticationInfoRequest) GetSipNumberAuthItems() int32`

GetSipNumberAuthItems returns the SipNumberAuthItems field if non-nil, zero value otherwise.

### GetSipNumberAuthItemsOk

`func (o *SipAuthenticationInfoRequest) GetSipNumberAuthItemsOk() (*int32, bool)`

GetSipNumberAuthItemsOk returns a tuple with the SipNumberAuthItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipNumberAuthItems

`func (o *SipAuthenticationInfoRequest) SetSipNumberAuthItems(v int32)`

SetSipNumberAuthItems sets SipNumberAuthItems field to given value.

### HasSipNumberAuthItems

`func (o *SipAuthenticationInfoRequest) HasSipNumberAuthItems() bool`

HasSipNumberAuthItems returns a boolean if a field has been set.

### GetResynchronizationInfo

`func (o *SipAuthenticationInfoRequest) GetResynchronizationInfo() ResynchronizationInfo`

GetResynchronizationInfo returns the ResynchronizationInfo field if non-nil, zero value otherwise.

### GetResynchronizationInfoOk

`func (o *SipAuthenticationInfoRequest) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool)`

GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResynchronizationInfo

`func (o *SipAuthenticationInfoRequest) SetResynchronizationInfo(v ResynchronizationInfo)`

SetResynchronizationInfo sets ResynchronizationInfo field to given value.

### HasResynchronizationInfo

`func (o *SipAuthenticationInfoRequest) HasResynchronizationInfo() bool`

HasResynchronizationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


