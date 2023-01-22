# EnableUeReachabilityReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reachability** | [**UeReachability**](UeReachability.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**OldGuami** | Pointer to [**Guami**](Guami.md) |  | [optional] 
**ExtBufSupport** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEnableUeReachabilityReqData

`func NewEnableUeReachabilityReqData(reachability UeReachability, ) *EnableUeReachabilityReqData`

NewEnableUeReachabilityReqData instantiates a new EnableUeReachabilityReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableUeReachabilityReqDataWithDefaults

`func NewEnableUeReachabilityReqDataWithDefaults() *EnableUeReachabilityReqData`

NewEnableUeReachabilityReqDataWithDefaults instantiates a new EnableUeReachabilityReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReachability

`func (o *EnableUeReachabilityReqData) GetReachability() UeReachability`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *EnableUeReachabilityReqData) GetReachabilityOk() (*UeReachability, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *EnableUeReachabilityReqData) SetReachability(v UeReachability)`

SetReachability sets Reachability field to given value.


### GetSupportedFeatures

`func (o *EnableUeReachabilityReqData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EnableUeReachabilityReqData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EnableUeReachabilityReqData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *EnableUeReachabilityReqData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetOldGuami

`func (o *EnableUeReachabilityReqData) GetOldGuami() Guami`

GetOldGuami returns the OldGuami field if non-nil, zero value otherwise.

### GetOldGuamiOk

`func (o *EnableUeReachabilityReqData) GetOldGuamiOk() (*Guami, bool)`

GetOldGuamiOk returns a tuple with the OldGuami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldGuami

`func (o *EnableUeReachabilityReqData) SetOldGuami(v Guami)`

SetOldGuami sets OldGuami field to given value.

### HasOldGuami

`func (o *EnableUeReachabilityReqData) HasOldGuami() bool`

HasOldGuami returns a boolean if a field has been set.

### GetExtBufSupport

`func (o *EnableUeReachabilityReqData) GetExtBufSupport() bool`

GetExtBufSupport returns the ExtBufSupport field if non-nil, zero value otherwise.

### GetExtBufSupportOk

`func (o *EnableUeReachabilityReqData) GetExtBufSupportOk() (*bool, bool)`

GetExtBufSupportOk returns a tuple with the ExtBufSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtBufSupport

`func (o *EnableUeReachabilityReqData) SetExtBufSupport(v bool)`

SetExtBufSupport sets ExtBufSupport field to given value.

### HasExtBufSupport

`func (o *EnableUeReachabilityReqData) HasExtBufSupport() bool`

HasExtBufSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


