# RdsDownlinkDataDeliveryFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**Title** | Pointer to **string** | A short, human-readable summary of the problem type. It should not change from occurrence to occurrence of the problem. | [optional] 
**Status** | Pointer to **int32** | The HTTP status code for this occurrence of the problem. | [optional] 
**Detail** | Pointer to **string** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**Instance** | Pointer to **string** | string providing an URI formatted according to IETF RFC 3986. | [optional] 
**Cause** | Pointer to **string** | A machine-readable application error cause specific to this occurrence of the problem. This IE should be present and provide application-related error information, if available. | [optional] 
**InvalidParams** | Pointer to [**[]InvalidParam**](InvalidParam.md) | Description of invalid parameters, for a request rejected due to invalid parameters. | [optional] 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 
**RequestedRetransmissionTime** | Pointer to **time.Time** | string with format \&quot;date-time\&quot; as defined in OpenAPI. | [optional] 
**SupportedUeFormats** | Pointer to [**[]SerializationFormat**](SerializationFormat.md) | Indicates the serialization format(s) that are supported by the UE on the associated RDS port. | [optional] 

## Methods

### NewRdsDownlinkDataDeliveryFailure

`func NewRdsDownlinkDataDeliveryFailure() *RdsDownlinkDataDeliveryFailure`

NewRdsDownlinkDataDeliveryFailure instantiates a new RdsDownlinkDataDeliveryFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsDownlinkDataDeliveryFailureWithDefaults

`func NewRdsDownlinkDataDeliveryFailureWithDefaults() *RdsDownlinkDataDeliveryFailure`

NewRdsDownlinkDataDeliveryFailureWithDefaults instantiates a new RdsDownlinkDataDeliveryFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RdsDownlinkDataDeliveryFailure) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RdsDownlinkDataDeliveryFailure) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RdsDownlinkDataDeliveryFailure) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RdsDownlinkDataDeliveryFailure) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *RdsDownlinkDataDeliveryFailure) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RdsDownlinkDataDeliveryFailure) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RdsDownlinkDataDeliveryFailure) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RdsDownlinkDataDeliveryFailure) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetStatus

`func (o *RdsDownlinkDataDeliveryFailure) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RdsDownlinkDataDeliveryFailure) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RdsDownlinkDataDeliveryFailure) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RdsDownlinkDataDeliveryFailure) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDetail

`func (o *RdsDownlinkDataDeliveryFailure) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *RdsDownlinkDataDeliveryFailure) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *RdsDownlinkDataDeliveryFailure) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *RdsDownlinkDataDeliveryFailure) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *RdsDownlinkDataDeliveryFailure) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *RdsDownlinkDataDeliveryFailure) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *RdsDownlinkDataDeliveryFailure) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *RdsDownlinkDataDeliveryFailure) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetCause

`func (o *RdsDownlinkDataDeliveryFailure) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *RdsDownlinkDataDeliveryFailure) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *RdsDownlinkDataDeliveryFailure) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *RdsDownlinkDataDeliveryFailure) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetInvalidParams

`func (o *RdsDownlinkDataDeliveryFailure) GetInvalidParams() []InvalidParam`

GetInvalidParams returns the InvalidParams field if non-nil, zero value otherwise.

### GetInvalidParamsOk

`func (o *RdsDownlinkDataDeliveryFailure) GetInvalidParamsOk() (*[]InvalidParam, bool)`

GetInvalidParamsOk returns a tuple with the InvalidParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidParams

`func (o *RdsDownlinkDataDeliveryFailure) SetInvalidParams(v []InvalidParam)`

SetInvalidParams sets InvalidParams field to given value.

### HasInvalidParams

`func (o *RdsDownlinkDataDeliveryFailure) HasInvalidParams() bool`

HasInvalidParams returns a boolean if a field has been set.

### GetSupportedFeatures

`func (o *RdsDownlinkDataDeliveryFailure) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *RdsDownlinkDataDeliveryFailure) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *RdsDownlinkDataDeliveryFailure) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *RdsDownlinkDataDeliveryFailure) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.

### GetRequestedRetransmissionTime

`func (o *RdsDownlinkDataDeliveryFailure) GetRequestedRetransmissionTime() time.Time`

GetRequestedRetransmissionTime returns the RequestedRetransmissionTime field if non-nil, zero value otherwise.

### GetRequestedRetransmissionTimeOk

`func (o *RdsDownlinkDataDeliveryFailure) GetRequestedRetransmissionTimeOk() (*time.Time, bool)`

GetRequestedRetransmissionTimeOk returns a tuple with the RequestedRetransmissionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedRetransmissionTime

`func (o *RdsDownlinkDataDeliveryFailure) SetRequestedRetransmissionTime(v time.Time)`

SetRequestedRetransmissionTime sets RequestedRetransmissionTime field to given value.

### HasRequestedRetransmissionTime

`func (o *RdsDownlinkDataDeliveryFailure) HasRequestedRetransmissionTime() bool`

HasRequestedRetransmissionTime returns a boolean if a field has been set.

### GetSupportedUeFormats

`func (o *RdsDownlinkDataDeliveryFailure) GetSupportedUeFormats() []SerializationFormat`

GetSupportedUeFormats returns the SupportedUeFormats field if non-nil, zero value otherwise.

### GetSupportedUeFormatsOk

`func (o *RdsDownlinkDataDeliveryFailure) GetSupportedUeFormatsOk() (*[]SerializationFormat, bool)`

GetSupportedUeFormatsOk returns a tuple with the SupportedUeFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedUeFormats

`func (o *RdsDownlinkDataDeliveryFailure) SetSupportedUeFormats(v []SerializationFormat)`

SetSupportedUeFormats sets SupportedUeFormats field to given value.

### HasSupportedUeFormats

`func (o *RdsDownlinkDataDeliveryFailure) HasSupportedUeFormats() bool`

HasSupportedUeFormats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


