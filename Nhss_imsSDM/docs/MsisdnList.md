# MsisdnList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BasicMsisdn** | **string** | String containing an additional or basic MSISDN | 
**AdditionalMsisdns** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMsisdnList

`func NewMsisdnList(basicMsisdn string, ) *MsisdnList`

NewMsisdnList instantiates a new MsisdnList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsisdnListWithDefaults

`func NewMsisdnListWithDefaults() *MsisdnList`

NewMsisdnListWithDefaults instantiates a new MsisdnList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasicMsisdn

`func (o *MsisdnList) GetBasicMsisdn() string`

GetBasicMsisdn returns the BasicMsisdn field if non-nil, zero value otherwise.

### GetBasicMsisdnOk

`func (o *MsisdnList) GetBasicMsisdnOk() (*string, bool)`

GetBasicMsisdnOk returns a tuple with the BasicMsisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicMsisdn

`func (o *MsisdnList) SetBasicMsisdn(v string)`

SetBasicMsisdn sets BasicMsisdn field to given value.


### GetAdditionalMsisdns

`func (o *MsisdnList) GetAdditionalMsisdns() []string`

GetAdditionalMsisdns returns the AdditionalMsisdns field if non-nil, zero value otherwise.

### GetAdditionalMsisdnsOk

`func (o *MsisdnList) GetAdditionalMsisdnsOk() (*[]string, bool)`

GetAdditionalMsisdnsOk returns a tuple with the AdditionalMsisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMsisdns

`func (o *MsisdnList) SetAdditionalMsisdns(v []string)`

SetAdditionalMsisdns sets AdditionalMsisdns field to given value.

### HasAdditionalMsisdns

`func (o *MsisdnList) HasAdditionalMsisdns() bool`

HasAdditionalMsisdns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


