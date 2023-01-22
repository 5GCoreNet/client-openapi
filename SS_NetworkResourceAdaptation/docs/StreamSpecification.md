# StreamSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SrcMacAddr** | **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | 
**DstMacAddr** | **string** | String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.  | 

## Methods

### NewStreamSpecification

`func NewStreamSpecification(srcMacAddr string, dstMacAddr string, ) *StreamSpecification`

NewStreamSpecification instantiates a new StreamSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSpecificationWithDefaults

`func NewStreamSpecificationWithDefaults() *StreamSpecification`

NewStreamSpecificationWithDefaults instantiates a new StreamSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrcMacAddr

`func (o *StreamSpecification) GetSrcMacAddr() string`

GetSrcMacAddr returns the SrcMacAddr field if non-nil, zero value otherwise.

### GetSrcMacAddrOk

`func (o *StreamSpecification) GetSrcMacAddrOk() (*string, bool)`

GetSrcMacAddrOk returns a tuple with the SrcMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMacAddr

`func (o *StreamSpecification) SetSrcMacAddr(v string)`

SetSrcMacAddr sets SrcMacAddr field to given value.


### GetDstMacAddr

`func (o *StreamSpecification) GetDstMacAddr() string`

GetDstMacAddr returns the DstMacAddr field if non-nil, zero value otherwise.

### GetDstMacAddrOk

`func (o *StreamSpecification) GetDstMacAddrOk() (*string, bool)`

GetDstMacAddrOk returns a tuple with the DstMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstMacAddr

`func (o *StreamSpecification) SetDstMacAddr(v string)`

SetDstMacAddr sets DstMacAddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


