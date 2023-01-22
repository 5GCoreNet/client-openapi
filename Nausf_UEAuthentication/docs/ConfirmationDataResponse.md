# ConfirmationDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthResult** | [**AuthResult**](AuthResult.md) |  | 
**Supi** | Pointer to **string** | String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \&quot;imsi-&lt;imsi&gt;\&quot;, where &lt;imsi&gt; shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \&quot;nai-&lt;nai&gt;, where &lt;nai&gt; shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \&quot;gci-&lt;gci&gt;\&quot;, where &lt;gci&gt; shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \&quot;gli-&lt;gli&gt;\&quot;, where &lt;gli&gt; shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \&quot;lower-with-hyphen\&quot; naming convention    defined in 3GPP TS 29.501.  | [optional] 
**Kseaf** | Pointer to **string** | Contains the Kseaf. | [optional] 
**PvsInfo** | Pointer to [**[]ServerAddressingInfo**](ServerAddressingInfo.md) |  | [optional] 

## Methods

### NewConfirmationDataResponse

`func NewConfirmationDataResponse(authResult AuthResult, ) *ConfirmationDataResponse`

NewConfirmationDataResponse instantiates a new ConfirmationDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmationDataResponseWithDefaults

`func NewConfirmationDataResponseWithDefaults() *ConfirmationDataResponse`

NewConfirmationDataResponseWithDefaults instantiates a new ConfirmationDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthResult

`func (o *ConfirmationDataResponse) GetAuthResult() AuthResult`

GetAuthResult returns the AuthResult field if non-nil, zero value otherwise.

### GetAuthResultOk

`func (o *ConfirmationDataResponse) GetAuthResultOk() (*AuthResult, bool)`

GetAuthResultOk returns a tuple with the AuthResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthResult

`func (o *ConfirmationDataResponse) SetAuthResult(v AuthResult)`

SetAuthResult sets AuthResult field to given value.


### GetSupi

`func (o *ConfirmationDataResponse) GetSupi() string`

GetSupi returns the Supi field if non-nil, zero value otherwise.

### GetSupiOk

`func (o *ConfirmationDataResponse) GetSupiOk() (*string, bool)`

GetSupiOk returns a tuple with the Supi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupi

`func (o *ConfirmationDataResponse) SetSupi(v string)`

SetSupi sets Supi field to given value.

### HasSupi

`func (o *ConfirmationDataResponse) HasSupi() bool`

HasSupi returns a boolean if a field has been set.

### GetKseaf

`func (o *ConfirmationDataResponse) GetKseaf() string`

GetKseaf returns the Kseaf field if non-nil, zero value otherwise.

### GetKseafOk

`func (o *ConfirmationDataResponse) GetKseafOk() (*string, bool)`

GetKseafOk returns a tuple with the Kseaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKseaf

`func (o *ConfirmationDataResponse) SetKseaf(v string)`

SetKseaf sets Kseaf field to given value.

### HasKseaf

`func (o *ConfirmationDataResponse) HasKseaf() bool`

HasKseaf returns a boolean if a field has been set.

### GetPvsInfo

`func (o *ConfirmationDataResponse) GetPvsInfo() []ServerAddressingInfo`

GetPvsInfo returns the PvsInfo field if non-nil, zero value otherwise.

### GetPvsInfoOk

`func (o *ConfirmationDataResponse) GetPvsInfoOk() (*[]ServerAddressingInfo, bool)`

GetPvsInfoOk returns a tuple with the PvsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvsInfo

`func (o *ConfirmationDataResponse) SetPvsInfo(v []ServerAddressingInfo)`

SetPvsInfo sets PvsInfo field to given value.

### HasPvsInfo

`func (o *ConfirmationDataResponse) HasPvsInfo() bool`

HasPvsInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


