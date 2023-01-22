# LocUpdateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gpsi** | **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | 
**LocInfo** | [**LocationInfo**](LocationInfo.md) |  | 
**LcsQosClass** | [**LcsQosClass**](LcsQosClass.md) |  | 
**SvcId** | Pointer to **string** | Contains the service identity | [optional] 
**SuppFeat** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 

## Methods

### NewLocUpdateData

`func NewLocUpdateData(gpsi string, locInfo LocationInfo, lcsQosClass LcsQosClass, suppFeat string, ) *LocUpdateData`

NewLocUpdateData instantiates a new LocUpdateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocUpdateDataWithDefaults

`func NewLocUpdateDataWithDefaults() *LocUpdateData`

NewLocUpdateDataWithDefaults instantiates a new LocUpdateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGpsi

`func (o *LocUpdateData) GetGpsi() string`

GetGpsi returns the Gpsi field if non-nil, zero value otherwise.

### GetGpsiOk

`func (o *LocUpdateData) GetGpsiOk() (*string, bool)`

GetGpsiOk returns a tuple with the Gpsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsi

`func (o *LocUpdateData) SetGpsi(v string)`

SetGpsi sets Gpsi field to given value.


### GetLocInfo

`func (o *LocUpdateData) GetLocInfo() LocationInfo`

GetLocInfo returns the LocInfo field if non-nil, zero value otherwise.

### GetLocInfoOk

`func (o *LocUpdateData) GetLocInfoOk() (*LocationInfo, bool)`

GetLocInfoOk returns a tuple with the LocInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocInfo

`func (o *LocUpdateData) SetLocInfo(v LocationInfo)`

SetLocInfo sets LocInfo field to given value.


### GetLcsQosClass

`func (o *LocUpdateData) GetLcsQosClass() LcsQosClass`

GetLcsQosClass returns the LcsQosClass field if non-nil, zero value otherwise.

### GetLcsQosClassOk

`func (o *LocUpdateData) GetLcsQosClassOk() (*LcsQosClass, bool)`

GetLcsQosClassOk returns a tuple with the LcsQosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLcsQosClass

`func (o *LocUpdateData) SetLcsQosClass(v LcsQosClass)`

SetLcsQosClass sets LcsQosClass field to given value.


### GetSvcId

`func (o *LocUpdateData) GetSvcId() string`

GetSvcId returns the SvcId field if non-nil, zero value otherwise.

### GetSvcIdOk

`func (o *LocUpdateData) GetSvcIdOk() (*string, bool)`

GetSvcIdOk returns a tuple with the SvcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcId

`func (o *LocUpdateData) SetSvcId(v string)`

SetSvcId sets SvcId field to given value.

### HasSvcId

`func (o *LocUpdateData) HasSvcId() bool`

HasSvcId returns a boolean if a field has been set.

### GetSuppFeat

`func (o *LocUpdateData) GetSuppFeat() string`

GetSuppFeat returns the SuppFeat field if non-nil, zero value otherwise.

### GetSuppFeatOk

`func (o *LocUpdateData) GetSuppFeatOk() (*string, bool)`

GetSuppFeatOk returns a tuple with the SuppFeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppFeat

`func (o *LocUpdateData) SetSuppFeat(v string)`

SetSuppFeat sets SuppFeat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


