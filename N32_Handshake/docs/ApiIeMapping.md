# ApiIeMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiSignature** | [**ApiSignature**](ApiSignature.md) |  | 
**ApiMethod** | [**HttpMethod**](HttpMethod.md) |  | 
**IeList** | [**[]IeInfo**](IeInfo.md) |  | 

## Methods

### NewApiIeMapping

`func NewApiIeMapping(apiSignature ApiSignature, apiMethod HttpMethod, ieList []IeInfo, ) *ApiIeMapping`

NewApiIeMapping instantiates a new ApiIeMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiIeMappingWithDefaults

`func NewApiIeMappingWithDefaults() *ApiIeMapping`

NewApiIeMappingWithDefaults instantiates a new ApiIeMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiSignature

`func (o *ApiIeMapping) GetApiSignature() ApiSignature`

GetApiSignature returns the ApiSignature field if non-nil, zero value otherwise.

### GetApiSignatureOk

`func (o *ApiIeMapping) GetApiSignatureOk() (*ApiSignature, bool)`

GetApiSignatureOk returns a tuple with the ApiSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiSignature

`func (o *ApiIeMapping) SetApiSignature(v ApiSignature)`

SetApiSignature sets ApiSignature field to given value.


### GetApiMethod

`func (o *ApiIeMapping) GetApiMethod() HttpMethod`

GetApiMethod returns the ApiMethod field if non-nil, zero value otherwise.

### GetApiMethodOk

`func (o *ApiIeMapping) GetApiMethodOk() (*HttpMethod, bool)`

GetApiMethodOk returns a tuple with the ApiMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiMethod

`func (o *ApiIeMapping) SetApiMethod(v HttpMethod)`

SetApiMethod sets ApiMethod field to given value.


### GetIeList

`func (o *ApiIeMapping) GetIeList() []IeInfo`

GetIeList returns the IeList field if non-nil, zero value otherwise.

### GetIeListOk

`func (o *ApiIeMapping) GetIeListOk() (*[]IeInfo, bool)`

GetIeListOk returns a tuple with the IeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIeList

`func (o *ApiIeMapping) SetIeList(v []IeInfo)`

SetIeList sets IeList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


