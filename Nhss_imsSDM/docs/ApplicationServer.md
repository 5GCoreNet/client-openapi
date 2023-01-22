# ApplicationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsUri** | **string** |  | 
**SessionContinue** | Pointer to **bool** |  | [optional] 
**ServiceInfoList** | Pointer to [**[]ServiceInformation**](ServiceInformation.md) |  | [optional] 

## Methods

### NewApplicationServer

`func NewApplicationServer(asUri string, ) *ApplicationServer`

NewApplicationServer instantiates a new ApplicationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationServerWithDefaults

`func NewApplicationServerWithDefaults() *ApplicationServer`

NewApplicationServerWithDefaults instantiates a new ApplicationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsUri

`func (o *ApplicationServer) GetAsUri() string`

GetAsUri returns the AsUri field if non-nil, zero value otherwise.

### GetAsUriOk

`func (o *ApplicationServer) GetAsUriOk() (*string, bool)`

GetAsUriOk returns a tuple with the AsUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsUri

`func (o *ApplicationServer) SetAsUri(v string)`

SetAsUri sets AsUri field to given value.


### GetSessionContinue

`func (o *ApplicationServer) GetSessionContinue() bool`

GetSessionContinue returns the SessionContinue field if non-nil, zero value otherwise.

### GetSessionContinueOk

`func (o *ApplicationServer) GetSessionContinueOk() (*bool, bool)`

GetSessionContinueOk returns a tuple with the SessionContinue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionContinue

`func (o *ApplicationServer) SetSessionContinue(v bool)`

SetSessionContinue sets SessionContinue field to given value.

### HasSessionContinue

`func (o *ApplicationServer) HasSessionContinue() bool`

HasSessionContinue returns a boolean if a field has been set.

### GetServiceInfoList

`func (o *ApplicationServer) GetServiceInfoList() []ServiceInformation`

GetServiceInfoList returns the ServiceInfoList field if non-nil, zero value otherwise.

### GetServiceInfoListOk

`func (o *ApplicationServer) GetServiceInfoListOk() (*[]ServiceInformation, bool)`

GetServiceInfoListOk returns a tuple with the ServiceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInfoList

`func (o *ApplicationServer) SetServiceInfoList(v []ServiceInformation)`

SetServiceInfoList sets ServiceInfoList field to given value.

### HasServiceInfoList

`func (o *ApplicationServer) HasServiceInfoList() bool`

HasServiceInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


