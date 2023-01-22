# TraceData1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceRef** | **string** | Trace Reference (see 3GPP TS 32.422).It shall be encoded as the concatenation of MCC, MNC and Trace ID as follows: &#39;MCC&#39;&lt;MNC&#39;-&#39;Trace ID&#39;The Trace ID shall be encoded as a 3 octet string in hexadecimal representation. Each character in the Trace ID string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits of the Trace ID shall appear first  in the string, and the character representing the 4 least significant bit of the Trace ID shall appear last in the string.  | 
**TraceDepth** | [**TraceDepth**](TraceDepth.md) |  | 
**NeTypeList** | **string** | List of NE Types (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string.Octets shall be coded according to 3GPP TS 32.422.  | 
**EventList** | **string** | Triggering events (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the 4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422.  | 
**CollectionEntityIpv4Addr** | Pointer to **string** | String identifying a IPv4 address formatted in the &#39;dotted decimal&#39; notation as defined in RFC 1166.  | [optional] 
**CollectionEntityIpv6Addr** | Pointer to [**Ipv6Addr**](Ipv6Addr.md) |  | [optional] 
**InterfaceList** | Pointer to **string** | List of Interfaces (see 3GPP TS 32.422).It shall be encoded as an octet string in hexadecimal representation. Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;, \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent 4 bits. The most significant character representing the 4 most significant bits shall appear first in the string, and the character representing the  4 least significant bit shall appear last in the string. Octets shall be coded according to 3GPP TS 32.422. If this attribute is not present, all the interfaces applicable to the list of NE types indicated in the neTypeList attribute should be traced.  | [optional] 

## Methods

### NewTraceData1

`func NewTraceData1(traceRef string, traceDepth TraceDepth, neTypeList string, eventList string, ) *TraceData1`

NewTraceData1 instantiates a new TraceData1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceData1WithDefaults

`func NewTraceData1WithDefaults() *TraceData1`

NewTraceData1WithDefaults instantiates a new TraceData1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceRef

`func (o *TraceData1) GetTraceRef() string`

GetTraceRef returns the TraceRef field if non-nil, zero value otherwise.

### GetTraceRefOk

`func (o *TraceData1) GetTraceRefOk() (*string, bool)`

GetTraceRefOk returns a tuple with the TraceRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceRef

`func (o *TraceData1) SetTraceRef(v string)`

SetTraceRef sets TraceRef field to given value.


### GetTraceDepth

`func (o *TraceData1) GetTraceDepth() TraceDepth`

GetTraceDepth returns the TraceDepth field if non-nil, zero value otherwise.

### GetTraceDepthOk

`func (o *TraceData1) GetTraceDepthOk() (*TraceDepth, bool)`

GetTraceDepthOk returns a tuple with the TraceDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceDepth

`func (o *TraceData1) SetTraceDepth(v TraceDepth)`

SetTraceDepth sets TraceDepth field to given value.


### GetNeTypeList

`func (o *TraceData1) GetNeTypeList() string`

GetNeTypeList returns the NeTypeList field if non-nil, zero value otherwise.

### GetNeTypeListOk

`func (o *TraceData1) GetNeTypeListOk() (*string, bool)`

GetNeTypeListOk returns a tuple with the NeTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeTypeList

`func (o *TraceData1) SetNeTypeList(v string)`

SetNeTypeList sets NeTypeList field to given value.


### GetEventList

`func (o *TraceData1) GetEventList() string`

GetEventList returns the EventList field if non-nil, zero value otherwise.

### GetEventListOk

`func (o *TraceData1) GetEventListOk() (*string, bool)`

GetEventListOk returns a tuple with the EventList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventList

`func (o *TraceData1) SetEventList(v string)`

SetEventList sets EventList field to given value.


### GetCollectionEntityIpv4Addr

`func (o *TraceData1) GetCollectionEntityIpv4Addr() string`

GetCollectionEntityIpv4Addr returns the CollectionEntityIpv4Addr field if non-nil, zero value otherwise.

### GetCollectionEntityIpv4AddrOk

`func (o *TraceData1) GetCollectionEntityIpv4AddrOk() (*string, bool)`

GetCollectionEntityIpv4AddrOk returns a tuple with the CollectionEntityIpv4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionEntityIpv4Addr

`func (o *TraceData1) SetCollectionEntityIpv4Addr(v string)`

SetCollectionEntityIpv4Addr sets CollectionEntityIpv4Addr field to given value.

### HasCollectionEntityIpv4Addr

`func (o *TraceData1) HasCollectionEntityIpv4Addr() bool`

HasCollectionEntityIpv4Addr returns a boolean if a field has been set.

### GetCollectionEntityIpv6Addr

`func (o *TraceData1) GetCollectionEntityIpv6Addr() Ipv6Addr`

GetCollectionEntityIpv6Addr returns the CollectionEntityIpv6Addr field if non-nil, zero value otherwise.

### GetCollectionEntityIpv6AddrOk

`func (o *TraceData1) GetCollectionEntityIpv6AddrOk() (*Ipv6Addr, bool)`

GetCollectionEntityIpv6AddrOk returns a tuple with the CollectionEntityIpv6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionEntityIpv6Addr

`func (o *TraceData1) SetCollectionEntityIpv6Addr(v Ipv6Addr)`

SetCollectionEntityIpv6Addr sets CollectionEntityIpv6Addr field to given value.

### HasCollectionEntityIpv6Addr

`func (o *TraceData1) HasCollectionEntityIpv6Addr() bool`

HasCollectionEntityIpv6Addr returns a boolean if a field has been set.

### GetInterfaceList

`func (o *TraceData1) GetInterfaceList() string`

GetInterfaceList returns the InterfaceList field if non-nil, zero value otherwise.

### GetInterfaceListOk

`func (o *TraceData1) GetInterfaceListOk() (*string, bool)`

GetInterfaceListOk returns a tuple with the InterfaceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceList

`func (o *TraceData1) SetInterfaceList(v string)`

SetInterfaceList sets InterfaceList field to given value.

### HasInterfaceList

`func (o *TraceData1) HasInterfaceList() bool`

HasInterfaceList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

