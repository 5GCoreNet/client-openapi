# UeACRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Supi** | **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | 
**AnType** | [**AccessType**](AccessType.md) |  | 
**AcuOperationList** | [**[]AcuOperationItem**](AcuOperationItem.md) |  | 
**AdditionalAnType** | Pointer to [**AccessType**](AccessType.md) |  | [optional] 

## Methods

### NewUeACRequestInfo

`func NewUeACRequestInfo(supi string, anType AccessType, acuOperationList []AcuOperationItem, ) *UeACRequestInfo`

NewUeACRequestInfo instantiates a new UeACRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUeACRequestInfoWithDefaults

`func NewUeACRequestInfoWithDefaults() *UeACRequestInfo`

NewUeACRequestInfoWithDefaults instantiates a new UeACRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupi

`func (o *UeACRequestInfo) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *UeACRequestInfo) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *UeACRequestInfo) SetSupi(v string)`

SetSupi sets Supi field to given value.


### GetAnType

`func (o *UeACRequestInfo) GetAnType() AccessType`

GetAnType returns the AnType field if non-nil, zero value otherwise.

### GetAnTypeOk

`func (o *UeACRequestInfo) GetAnTypeOk() (*AccessType, bool)`

GetAnTypeOk returns a tuple with the AnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnType

`func (o *UeACRequestInfo) SetAnType(v AccessType)`

SetAnType sets AnType field to given value.


### GetAcuOperationList

`func (o *UeACRequestInfo) GetAcuOperationList() []AcuOperationItem`

GetAcuOperationList returns the AcuOperationList field if non-nil, zero value otherwise.

### GetAcuOperationListOk

`func (o *UeACRequestInfo) GetAcuOperationListOk() (*[]AcuOperationItem, bool)`

GetAcuOperationListOk returns a tuple with the AcuOperationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcuOperationList

`func (o *UeACRequestInfo) SetAcuOperationList(v []AcuOperationItem)`

SetAcuOperationList sets AcuOperationList field to given value.


### GetAdditionalAnType

`func (o *UeACRequestInfo) GetAdditionalAnType() AccessType`

GetAdditionalAnType returns the AdditionalAnType field if non-nil, zero value otherwise.

### GetAdditionalAnTypeOk

`func (o *UeACRequestInfo) GetAdditionalAnTypeOk() (*AccessType, bool)`

GetAdditionalAnTypeOk returns a tuple with the AdditionalAnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalAnType

`func (o *UeACRequestInfo) SetAdditionalAnType(v AccessType)`

SetAdditionalAnType sets AdditionalAnType field to given value.

### HasAdditionalAnType

`func (o *UeACRequestInfo) HasAdditionalAnType() bool`

HasAdditionalAnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


