# ConsentRevoked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UcPurpose** | [**UcPurpose**](UcPurpose.md) |  | 
**ExternalId** | Pointer to **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clause 4.6.2 of 3GPP TS 23.682 for more information. | [optional] 
**UeId** | Pointer to **string** | String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier&#x3D; \&quot;extid-&#39;extid&#39;, where &#39;extid&#39;  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.   | [optional] 

## Methods

### NewConsentRevoked

`func NewConsentRevoked(ucPurpose UcPurpose, ) *ConsentRevoked`

NewConsentRevoked instantiates a new ConsentRevoked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentRevokedWithDefaults

`func NewConsentRevokedWithDefaults() *ConsentRevoked`

NewConsentRevokedWithDefaults instantiates a new ConsentRevoked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUcPurpose

`func (o *ConsentRevoked) GetUcPurpose() UcPurpose`

GetUcPurpose returns the UcPurpose field if non-nil, zero value otherwise.

### GetUcPurposeOk

`func (o *ConsentRevoked) GetUcPurposeOk() (*UcPurpose, bool)`

GetUcPurposeOk returns a tuple with the UcPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcPurpose

`func (o *ConsentRevoked) SetUcPurpose(v UcPurpose)`

SetUcPurpose sets UcPurpose field to given value.


### GetExternalId

`func (o *ConsentRevoked) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ConsentRevoked) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ConsentRevoked) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ConsentRevoked) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetUeId

`func (o *ConsentRevoked) GetUeId() string`

GetUeId returns the UeId field if non-nil, zero value otherwise.

### GetUeIdOk

`func (o *ConsentRevoked) GetUeIdOk() (*string, bool)`

GetUeIdOk returns a tuple with the UeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeId

`func (o *ConsentRevoked) SetUeId(v string)`

SetUeId sets UeId field to given value.

### HasUeId

`func (o *ConsentRevoked) HasUeId() bool`

HasUeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


