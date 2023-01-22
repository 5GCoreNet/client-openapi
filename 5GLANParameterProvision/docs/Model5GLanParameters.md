# Model5GLanParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExterGroupId** | **string** | string containing a local identifier followed by \&quot;@\&quot; and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \&quot;@\&quot; characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information. | 
**Gpsis** | **map[string]string** | Contains the list of 5G VN Group members, each member is identified by GPSI. Any string value can be used as a key of the map.  | 
**Dnn** | **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | 
**AaaIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**AaaIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**AaaUsgs** | Pointer to [**[]AaaUsage**](AaaUsage.md) |  | [optional] 
**MtcProviderId** | Pointer to **string** | String uniquely identifying MTC provider information. | [optional] 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**SessionType** | [**PduSessionType**](PduSessionType.md) |  | 
**SessionTypes** | Pointer to [**[]PduSessionType**](PduSessionType.md) | Further allowed PDU Session types. | [optional] 
**AppDesps** | [**map[string]AppDescriptor**](AppDescriptor.md) | Describes the operation systems and the corresponding applications for each operation systems. The key of map is osId. | 

## Methods

### NewModel5GLanParameters

`func NewModel5GLanParameters(exterGroupId string, gpsis map[string]string, dnn string, snssai Snssai, sessionType PduSessionType, appDesps map[string]AppDescriptor, ) *Model5GLanParameters`

NewModel5GLanParameters instantiates a new Model5GLanParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel5GLanParametersWithDefaults

`func NewModel5GLanParametersWithDefaults() *Model5GLanParameters`

NewModel5GLanParametersWithDefaults instantiates a new Model5GLanParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExterGroupId

`func (o *Model5GLanParameters) GetExterGroupId() string`

GetExterGroupId returns the ExterGroupId field if non-nil, zero value otherwise.

### GetExterGroupIdOk

`func (o *Model5GLanParameters) GetExterGroupIdOk() (*string, bool)`

GetExterGroupIdOk returns a tuple with the ExterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExterGroupId

`func (o *Model5GLanParameters) SetExterGroupId(v string)`

SetExterGroupId sets ExterGroupId field to given value.


### GetGpsis

`func (o *Model5GLanParameters) GetGpsis() map[string]string`

GetGpsis returns the Gpsis field if non-nil, zero value otherwise.

### GetGpsisOk

`func (o *Model5GLanParameters) GetGpsisOk() (*map[string]string, bool)`

GetGpsisOk returns a tuple with the Gpsis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpsis

`func (o *Model5GLanParameters) SetGpsis(v map[string]string)`

SetGpsis sets Gpsis field to given value.


### GetDnn

`func (o *Model5GLanParameters) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *Model5GLanParameters) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *Model5GLanParameters) SetDnn(v string)`

SetDnn sets Dnn field to given value.


### GetAaaIpv4Addr

`func (o *Model5GLanParameters) GetAaaIpv4Addr() string`

GetAaaIpv4Addr returns the AaaIpv4Addr field if non-nil, zero value otherwise.

### GetAaaIpv4AddrOk

`func (o *Model5GLanParameters) GetAaaIpv4AddrOk() (*string, bool)`

GetAaaIpv4AddrOk returns a tuple with the AaaIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaaIpv4Addr

`func (o *Model5GLanParameters) SetAaaIpv4Addr(v string)`

SetAaaIpv4Addr sets AaaIpv4Addr field to given value.

### HasAaaIpv4Addr

`func (o *Model5GLanParameters) HasAaaIpv4Addr() bool`

HasAaaIpv4Addr returns a boolean if a field has been set.

### GetAaaIpv6Addr

`func (o *Model5GLanParameters) GetAaaIpv6Addr() Ipv6Addr`

GetAaaIpv6Addr returns the AaaIpv6Addr field if non-nil, zero value otherwise.

### GetAaaIpv6AddrOk

`func (o *Model5GLanParameters) GetAaaIpv6AddrOk() (*Ipv6Addr, bool)`

GetAaaIpv6AddrOk returns a tuple with the AaaIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaaIpv6Addr

`func (o *Model5GLanParameters) SetAaaIpv6Addr(v Ipv6Addr)`

SetAaaIpv6Addr sets AaaIpv6Addr field to given value.

### HasAaaIpv6Addr

`func (o *Model5GLanParameters) HasAaaIpv6Addr() bool`

HasAaaIpv6Addr returns a boolean if a field has been set.

### GetAaaUsgs

`func (o *Model5GLanParameters) GetAaaUsgs() []AaaUsage`

GetAaaUsgs returns the AaaUsgs field if non-nil, zero value otherwise.

### GetAaaUsgsOk

`func (o *Model5GLanParameters) GetAaaUsgsOk() (*[]AaaUsage, bool)`

GetAaaUsgsOk returns a tuple with the AaaUsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAaaUsgs

`func (o *Model5GLanParameters) SetAaaUsgs(v []AaaUsage)`

SetAaaUsgs sets AaaUsgs field to given value.

### HasAaaUsgs

`func (o *Model5GLanParameters) HasAaaUsgs() bool`

HasAaaUsgs returns a boolean if a field has been set.

### GetMtcProviderId

`func (o *Model5GLanParameters) GetMtcProviderId() string`

GetMtcProviderId returns the MtcProviderId field if non-nil, zero value otherwise.

### GetMtcProviderIdOk

`func (o *Model5GLanParameters) GetMtcProviderIdOk() (*string, bool)`

GetMtcProviderIdOk returns a tuple with the MtcProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtcProviderId

`func (o *Model5GLanParameters) SetMtcProviderId(v string)`

SetMtcProviderId sets MtcProviderId field to given value.

### HasMtcProviderId

`func (o *Model5GLanParameters) HasMtcProviderId() bool`

HasMtcProviderId returns a boolean if a field has been set.

### GetSnssai

`func (o *Model5GLanParameters) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *Model5GLanParameters) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *Model5GLanParameters) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetSessionType

`func (o *Model5GLanParameters) GetSessionType() PduSessionType`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *Model5GLanParameters) GetSessionTypeOk() (*PduSessionType, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *Model5GLanParameters) SetSessionType(v PduSessionType)`

SetSessionType sets SessionType field to given value.


### GetSessionTypes

`func (o *Model5GLanParameters) GetSessionTypes() []PduSessionType`

GetSessionTypes returns the SessionTypes field if non-nil, zero value otherwise.

### GetSessionTypesOk

`func (o *Model5GLanParameters) GetSessionTypesOk() (*[]PduSessionType, bool)`

GetSessionTypesOk returns a tuple with the SessionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTypes

`func (o *Model5GLanParameters) SetSessionTypes(v []PduSessionType)`

SetSessionTypes sets SessionTypes field to given value.

### HasSessionTypes

`func (o *Model5GLanParameters) HasSessionTypes() bool`

HasSessionTypes returns a boolean if a field has been set.

### GetAppDesps

`func (o *Model5GLanParameters) GetAppDesps() map[string]AppDescriptor`

GetAppDesps returns the AppDesps field if non-nil, zero value otherwise.

### GetAppDespsOk

`func (o *Model5GLanParameters) GetAppDespsOk() (*map[string]AppDescriptor, bool)`

GetAppDespsOk returns a tuple with the AppDesps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDesps

`func (o *Model5GLanParameters) SetAppDesps(v map[string]AppDescriptor)`

SetAppDesps sets AppDesps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


