# TrafficDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnn** | Pointer to **string** | String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \&quot;Label1.Label2.Label3\&quot;).  | [optional] 
**SNssai** | Pointer to [**Snssai**](Snssai.md) |  | [optional] 
**DddTrafficDescriptorList** | Pointer to [**[]DddTrafficDescriptor**](DddTrafficDescriptor.md) |  | [optional] 

## Methods

### NewTrafficDescriptor

`func NewTrafficDescriptor() *TrafficDescriptor`

NewTrafficDescriptor instantiates a new TrafficDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficDescriptorWithDefaults

`func NewTrafficDescriptorWithDefaults() *TrafficDescriptor`

NewTrafficDescriptorWithDefaults instantiates a new TrafficDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnn

`func (o *TrafficDescriptor) GetDnn() string`

GetDnn returns the Dnn field if non-nil, zero value otherwise.

### GetDnnOk

`func (o *TrafficDescriptor) GetDnnOk() (*string, bool)`

GetDnnOk returns a tuple with the Dnn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnn

`func (o *TrafficDescriptor) SetDnn(v string)`

SetDnn sets Dnn field to given value.

### HasDnn

`func (o *TrafficDescriptor) HasDnn() bool`

HasDnn returns a boolean if a field has been set.

### GetSNssai

`func (o *TrafficDescriptor) GetSNssai() Snssai`

GetSNssai returns the SNssai field if non-nil, zero value otherwise.

### GetSNssaiOk

`func (o *TrafficDescriptor) GetSNssaiOk() (*Snssai, bool)`

GetSNssaiOk returns a tuple with the SNssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSNssai

`func (o *TrafficDescriptor) SetSNssai(v Snssai)`

SetSNssai sets SNssai field to given value.

### HasSNssai

`func (o *TrafficDescriptor) HasSNssai() bool`

HasSNssai returns a boolean if a field has been set.

### GetDddTrafficDescriptorList

`func (o *TrafficDescriptor) GetDddTrafficDescriptorList() []DddTrafficDescriptor`

GetDddTrafficDescriptorList returns the DddTrafficDescriptorList field if non-nil, zero value otherwise.

### GetDddTrafficDescriptorListOk

`func (o *TrafficDescriptor) GetDddTrafficDescriptorListOk() (*[]DddTrafficDescriptor, bool)`

GetDddTrafficDescriptorListOk returns a tuple with the DddTrafficDescriptorList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDddTrafficDescriptorList

`func (o *TrafficDescriptor) SetDddTrafficDescriptorList(v []DddTrafficDescriptor)`

SetDddTrafficDescriptorList sets DddTrafficDescriptorList field to given value.

### HasDddTrafficDescriptorList

`func (o *TrafficDescriptor) HasDddTrafficDescriptorList() bool`

HasDddTrafficDescriptorList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


