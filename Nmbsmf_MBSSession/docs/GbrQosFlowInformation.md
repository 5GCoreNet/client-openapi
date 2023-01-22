# GbrQosFlowInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxFbrDl** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**GuaFbrDl** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**MaxPacketLossRateDl** | Pointer to **int32** | Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent.  | [optional] 

## Methods

### NewGbrQosFlowInformation

`func NewGbrQosFlowInformation(maxFbrDl string, guaFbrDl string, ) *GbrQosFlowInformation`

NewGbrQosFlowInformation instantiates a new GbrQosFlowInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGbrQosFlowInformationWithDefaults

`func NewGbrQosFlowInformationWithDefaults() *GbrQosFlowInformation`

NewGbrQosFlowInformationWithDefaults instantiates a new GbrQosFlowInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxFbrDl

`func (o *GbrQosFlowInformation) GetMaxFbrDl() string`

GetMaxFbrDl returns the MaxFbrDl field if non-nil, zero value otherwise.

### GetMaxFbrDlOk

`func (o *GbrQosFlowInformation) GetMaxFbrDlOk() (*string, bool)`

GetMaxFbrDlOk returns a tuple with the MaxFbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFbrDl

`func (o *GbrQosFlowInformation) SetMaxFbrDl(v string)`

SetMaxFbrDl sets MaxFbrDl field to given value.


### GetGuaFbrDl

`func (o *GbrQosFlowInformation) GetGuaFbrDl() string`

GetGuaFbrDl returns the GuaFbrDl field if non-nil, zero value otherwise.

### GetGuaFbrDlOk

`func (o *GbrQosFlowInformation) GetGuaFbrDlOk() (*string, bool)`

GetGuaFbrDlOk returns a tuple with the GuaFbrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaFbrDl

`func (o *GbrQosFlowInformation) SetGuaFbrDl(v string)`

SetGuaFbrDl sets GuaFbrDl field to given value.


### GetMaxPacketLossRateDl

`func (o *GbrQosFlowInformation) GetMaxPacketLossRateDl() int32`

GetMaxPacketLossRateDl returns the MaxPacketLossRateDl field if non-nil, zero value otherwise.

### GetMaxPacketLossRateDlOk

`func (o *GbrQosFlowInformation) GetMaxPacketLossRateDlOk() (*int32, bool)`

GetMaxPacketLossRateDlOk returns a tuple with the MaxPacketLossRateDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPacketLossRateDl

`func (o *GbrQosFlowInformation) SetMaxPacketLossRateDl(v int32)`

SetMaxPacketLossRateDl sets MaxPacketLossRateDl field to given value.

### HasMaxPacketLossRateDl

`func (o *GbrQosFlowInformation) HasMaxPacketLossRateDl() bool`

HasMaxPacketLossRateDl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


