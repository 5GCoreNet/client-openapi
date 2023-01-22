# SliceMbrRm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uplink** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 
**Downlink** | **string** | String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \&quot;K\&quot; is used to represent the standard symbol \&quot;k\&quot;.  | 

## Methods

### NewSliceMbrRm

`func NewSliceMbrRm(uplink string, downlink string, ) *SliceMbrRm`

NewSliceMbrRm instantiates a new SliceMbrRm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceMbrRmWithDefaults

`func NewSliceMbrRmWithDefaults() *SliceMbrRm`

NewSliceMbrRmWithDefaults instantiates a new SliceMbrRm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplink

`func (o *SliceMbrRm) GetUplink() string`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *SliceMbrRm) GetUplinkOk() (*string, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *SliceMbrRm) SetUplink(v string)`

SetUplink sets Uplink field to given value.


### GetDownlink

`func (o *SliceMbrRm) GetDownlink() string`

GetDownlink returns the Downlink field if non-nil, zero value otherwise.

### GetDownlinkOk

`func (o *SliceMbrRm) GetDownlinkOk() (*string, bool)`

GetDownlinkOk returns a tuple with the Downlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlink

`func (o *SliceMbrRm) SetDownlink(v string)`

SetDownlink sets Downlink field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


