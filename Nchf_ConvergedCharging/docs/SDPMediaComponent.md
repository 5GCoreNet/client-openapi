# SDPMediaComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SDPMediaName** | Pointer to **string** |  | [optional] 
**SDPMediaDescription** | Pointer to **[]string** |  | [optional] 
**LocalGWInsertedIndication** | Pointer to **bool** |  | [optional] 
**IpRealmDefaultIndication** | Pointer to **bool** |  | [optional] 
**TranscoderInsertedIndication** | Pointer to **bool** |  | [optional] 
**MediaInitiatorFlag** | Pointer to [**MediaInitiatorFlag**](MediaInitiatorFlag.md) |  | [optional] 
**MediaInitiatorParty** | Pointer to **string** |  | [optional] 
**ThreeGPPChargingId** | Pointer to **string** |  | [optional] 
**AccessNetworkChargingIdentifierValue** | Pointer to **string** |  | [optional] 
**SDPType** | Pointer to [**SDPType**](SDPType.md) |  | [optional] 

## Methods

### NewSDPMediaComponent

`func NewSDPMediaComponent() *SDPMediaComponent`

NewSDPMediaComponent instantiates a new SDPMediaComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDPMediaComponentWithDefaults

`func NewSDPMediaComponentWithDefaults() *SDPMediaComponent`

NewSDPMediaComponentWithDefaults instantiates a new SDPMediaComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSDPMediaName

`func (o *SDPMediaComponent) GetSDPMediaName() string`

GetSDPMediaName returns the SDPMediaName field if non-nil, zero value otherwise.

### GetSDPMediaNameOk

`func (o *SDPMediaComponent) GetSDPMediaNameOk() (*string, bool)`

GetSDPMediaNameOk returns a tuple with the SDPMediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPMediaName

`func (o *SDPMediaComponent) SetSDPMediaName(v string)`

SetSDPMediaName sets SDPMediaName field to given value.

### HasSDPMediaName

`func (o *SDPMediaComponent) HasSDPMediaName() bool`

HasSDPMediaName returns a boolean if a field has been set.

### GetSDPMediaDescription

`func (o *SDPMediaComponent) GetSDPMediaDescription() []string`

GetSDPMediaDescription returns the SDPMediaDescription field if non-nil, zero value otherwise.

### GetSDPMediaDescriptionOk

`func (o *SDPMediaComponent) GetSDPMediaDescriptionOk() (*[]string, bool)`

GetSDPMediaDescriptionOk returns a tuple with the SDPMediaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPMediaDescription

`func (o *SDPMediaComponent) SetSDPMediaDescription(v []string)`

SetSDPMediaDescription sets SDPMediaDescription field to given value.

### HasSDPMediaDescription

`func (o *SDPMediaComponent) HasSDPMediaDescription() bool`

HasSDPMediaDescription returns a boolean if a field has been set.

### GetLocalGWInsertedIndication

`func (o *SDPMediaComponent) GetLocalGWInsertedIndication() bool`

GetLocalGWInsertedIndication returns the LocalGWInsertedIndication field if non-nil, zero value otherwise.

### GetLocalGWInsertedIndicationOk

`func (o *SDPMediaComponent) GetLocalGWInsertedIndicationOk() (*bool, bool)`

GetLocalGWInsertedIndicationOk returns a tuple with the LocalGWInsertedIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGWInsertedIndication

`func (o *SDPMediaComponent) SetLocalGWInsertedIndication(v bool)`

SetLocalGWInsertedIndication sets LocalGWInsertedIndication field to given value.

### HasLocalGWInsertedIndication

`func (o *SDPMediaComponent) HasLocalGWInsertedIndication() bool`

HasLocalGWInsertedIndication returns a boolean if a field has been set.

### GetIpRealmDefaultIndication

`func (o *SDPMediaComponent) GetIpRealmDefaultIndication() bool`

GetIpRealmDefaultIndication returns the IpRealmDefaultIndication field if non-nil, zero value otherwise.

### GetIpRealmDefaultIndicationOk

`func (o *SDPMediaComponent) GetIpRealmDefaultIndicationOk() (*bool, bool)`

GetIpRealmDefaultIndicationOk returns a tuple with the IpRealmDefaultIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRealmDefaultIndication

`func (o *SDPMediaComponent) SetIpRealmDefaultIndication(v bool)`

SetIpRealmDefaultIndication sets IpRealmDefaultIndication field to given value.

### HasIpRealmDefaultIndication

`func (o *SDPMediaComponent) HasIpRealmDefaultIndication() bool`

HasIpRealmDefaultIndication returns a boolean if a field has been set.

### GetTranscoderInsertedIndication

`func (o *SDPMediaComponent) GetTranscoderInsertedIndication() bool`

GetTranscoderInsertedIndication returns the TranscoderInsertedIndication field if non-nil, zero value otherwise.

### GetTranscoderInsertedIndicationOk

`func (o *SDPMediaComponent) GetTranscoderInsertedIndicationOk() (*bool, bool)`

GetTranscoderInsertedIndicationOk returns a tuple with the TranscoderInsertedIndication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscoderInsertedIndication

`func (o *SDPMediaComponent) SetTranscoderInsertedIndication(v bool)`

SetTranscoderInsertedIndication sets TranscoderInsertedIndication field to given value.

### HasTranscoderInsertedIndication

`func (o *SDPMediaComponent) HasTranscoderInsertedIndication() bool`

HasTranscoderInsertedIndication returns a boolean if a field has been set.

### GetMediaInitiatorFlag

`func (o *SDPMediaComponent) GetMediaInitiatorFlag() MediaInitiatorFlag`

GetMediaInitiatorFlag returns the MediaInitiatorFlag field if non-nil, zero value otherwise.

### GetMediaInitiatorFlagOk

`func (o *SDPMediaComponent) GetMediaInitiatorFlagOk() (*MediaInitiatorFlag, bool)`

GetMediaInitiatorFlagOk returns a tuple with the MediaInitiatorFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInitiatorFlag

`func (o *SDPMediaComponent) SetMediaInitiatorFlag(v MediaInitiatorFlag)`

SetMediaInitiatorFlag sets MediaInitiatorFlag field to given value.

### HasMediaInitiatorFlag

`func (o *SDPMediaComponent) HasMediaInitiatorFlag() bool`

HasMediaInitiatorFlag returns a boolean if a field has been set.

### GetMediaInitiatorParty

`func (o *SDPMediaComponent) GetMediaInitiatorParty() string`

GetMediaInitiatorParty returns the MediaInitiatorParty field if non-nil, zero value otherwise.

### GetMediaInitiatorPartyOk

`func (o *SDPMediaComponent) GetMediaInitiatorPartyOk() (*string, bool)`

GetMediaInitiatorPartyOk returns a tuple with the MediaInitiatorParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaInitiatorParty

`func (o *SDPMediaComponent) SetMediaInitiatorParty(v string)`

SetMediaInitiatorParty sets MediaInitiatorParty field to given value.

### HasMediaInitiatorParty

`func (o *SDPMediaComponent) HasMediaInitiatorParty() bool`

HasMediaInitiatorParty returns a boolean if a field has been set.

### GetThreeGPPChargingId

`func (o *SDPMediaComponent) GetThreeGPPChargingId() string`

GetThreeGPPChargingId returns the ThreeGPPChargingId field if non-nil, zero value otherwise.

### GetThreeGPPChargingIdOk

`func (o *SDPMediaComponent) GetThreeGPPChargingIdOk() (*string, bool)`

GetThreeGPPChargingIdOk returns a tuple with the ThreeGPPChargingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeGPPChargingId

`func (o *SDPMediaComponent) SetThreeGPPChargingId(v string)`

SetThreeGPPChargingId sets ThreeGPPChargingId field to given value.

### HasThreeGPPChargingId

`func (o *SDPMediaComponent) HasThreeGPPChargingId() bool`

HasThreeGPPChargingId returns a boolean if a field has been set.

### GetAccessNetworkChargingIdentifierValue

`func (o *SDPMediaComponent) GetAccessNetworkChargingIdentifierValue() string`

GetAccessNetworkChargingIdentifierValue returns the AccessNetworkChargingIdentifierValue field if non-nil, zero value otherwise.

### GetAccessNetworkChargingIdentifierValueOk

`func (o *SDPMediaComponent) GetAccessNetworkChargingIdentifierValueOk() (*string, bool)`

GetAccessNetworkChargingIdentifierValueOk returns a tuple with the AccessNetworkChargingIdentifierValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessNetworkChargingIdentifierValue

`func (o *SDPMediaComponent) SetAccessNetworkChargingIdentifierValue(v string)`

SetAccessNetworkChargingIdentifierValue sets AccessNetworkChargingIdentifierValue field to given value.

### HasAccessNetworkChargingIdentifierValue

`func (o *SDPMediaComponent) HasAccessNetworkChargingIdentifierValue() bool`

HasAccessNetworkChargingIdentifierValue returns a boolean if a field has been set.

### GetSDPType

`func (o *SDPMediaComponent) GetSDPType() SDPType`

GetSDPType returns the SDPType field if non-nil, zero value otherwise.

### GetSDPTypeOk

`func (o *SDPMediaComponent) GetSDPTypeOk() (*SDPType, bool)`

GetSDPTypeOk returns a tuple with the SDPType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSDPType

`func (o *SDPMediaComponent) SetSDPType(v SDPType)`

SetSDPType sets SDPType field to given value.

### HasSDPType

`func (o *SDPMediaComponent) HasSDPType() bool`

HasSDPType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


