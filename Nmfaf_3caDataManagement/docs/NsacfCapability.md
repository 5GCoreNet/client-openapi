# NsacfCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportUeSAC** | Pointer to **bool** | Indicates the service capability of the NSACF to monitor and control the number of registered UEs per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported  | [optional] [default to false]
**SupportPduSAC** | Pointer to **bool** | Indicates the service capability of the NSACF to monitor and control the number of established PDU sessions per network slice for the network slice that is subject to NSAC   true: Supported   false (default): Not Supported  | [optional] [default to false]

## Methods

### NewNsacfCapability

`func NewNsacfCapability() *NsacfCapability`

NewNsacfCapability instantiates a new NsacfCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsacfCapabilityWithDefaults

`func NewNsacfCapabilityWithDefaults() *NsacfCapability`

NewNsacfCapabilityWithDefaults instantiates a new NsacfCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportUeSAC

`func (o *NsacfCapability) GetSupportUeSAC() bool`

GetSupportUeSAC returns the SupportUeSAC field if non-nil, zero value otherwise.

### GetSupportUeSACOk

`func (o *NsacfCapability) GetSupportUeSACOk() (*bool, bool)`

GetSupportUeSACOk returns a tuple with the SupportUeSAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUeSAC

`func (o *NsacfCapability) SetSupportUeSAC(v bool)`

SetSupportUeSAC sets SupportUeSAC field to given value.

### HasSupportUeSAC

`func (o *NsacfCapability) HasSupportUeSAC() bool`

HasSupportUeSAC returns a boolean if a field has been set.

### GetSupportPduSAC

`func (o *NsacfCapability) GetSupportPduSAC() bool`

GetSupportPduSAC returns the SupportPduSAC field if non-nil, zero value otherwise.

### GetSupportPduSACOk

`func (o *NsacfCapability) GetSupportPduSACOk() (*bool, bool)`

GetSupportPduSACOk returns a tuple with the SupportPduSAC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPduSAC

`func (o *NsacfCapability) SetSupportPduSAC(v bool)`

SetSupportPduSAC sets SupportPduSAC field to given value.

### HasSupportPduSAC

`func (o *NsacfCapability) HasSupportPduSAC() bool`

HasSupportPduSAC returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


