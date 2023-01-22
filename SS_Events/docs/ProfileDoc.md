# ProfileDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileInformation** | **string** | Profile information associated with the valUserId or valUeId. | 
**ValTgtUe** | [**ValTargetUe**](ValTargetUe.md) |  | 

## Methods

### NewProfileDoc

`func NewProfileDoc(profileInformation string, valTgtUe ValTargetUe, ) *ProfileDoc`

NewProfileDoc instantiates a new ProfileDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileDocWithDefaults

`func NewProfileDocWithDefaults() *ProfileDoc`

NewProfileDocWithDefaults instantiates a new ProfileDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileInformation

`func (o *ProfileDoc) GetProfileInformation() string`

GetProfileInformation returns the ProfileInformation field if non-nil, zero value otherwise.

### GetProfileInformationOk

`func (o *ProfileDoc) GetProfileInformationOk() (*string, bool)`

GetProfileInformationOk returns a tuple with the ProfileInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileInformation

`func (o *ProfileDoc) SetProfileInformation(v string)`

SetProfileInformation sets ProfileInformation field to given value.


### GetValTgtUe

`func (o *ProfileDoc) GetValTgtUe() ValTargetUe`

GetValTgtUe returns the ValTgtUe field if non-nil, zero value otherwise.

### GetValTgtUeOk

`func (o *ProfileDoc) GetValTgtUeOk() (*ValTargetUe, bool)`

GetValTgtUeOk returns a tuple with the ValTgtUe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValTgtUe

`func (o *ProfileDoc) SetValTgtUe(v ValTargetUe)`

SetValTgtUe sets ValTgtUe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


