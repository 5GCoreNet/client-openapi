# MbsN2MessageTransferReqData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MbsSessionId** | [**MbsSessionId**](MbsSessionId.md) |  | 
**AreaSessionId** | Pointer to **int32** | Integer where the allowed values correspond to the value range of an unsigned 16-bit integer. | [optional] 
**N2MbsSmInfo** | [**N2MbsSmInfo**](N2MbsSmInfo.md) |  | 
**SupportedFeatures** | Pointer to **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | [optional] 

## Methods

### NewMbsN2MessageTransferReqData

`func NewMbsN2MessageTransferReqData(mbsSessionId MbsSessionId, n2MbsSmInfo N2MbsSmInfo, ) *MbsN2MessageTransferReqData`

NewMbsN2MessageTransferReqData instantiates a new MbsN2MessageTransferReqData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMbsN2MessageTransferReqDataWithDefaults

`func NewMbsN2MessageTransferReqDataWithDefaults() *MbsN2MessageTransferReqData`

NewMbsN2MessageTransferReqDataWithDefaults instantiates a new MbsN2MessageTransferReqData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMbsSessionId

`func (o *MbsN2MessageTransferReqData) GetMbsSessionId() MbsSessionId`

GetMbsSessionId returns the MbsSessionId field if non-nil, zero value otherwise.

### GetMbsSessionIdOk

`func (o *MbsN2MessageTransferReqData) GetMbsSessionIdOk() (*MbsSessionId, bool)`

GetMbsSessionIdOk returns a tuple with the MbsSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbsSessionId

`func (o *MbsN2MessageTransferReqData) SetMbsSessionId(v MbsSessionId)`

SetMbsSessionId sets MbsSessionId field to given value.


### GetAreaSessionId

`func (o *MbsN2MessageTransferReqData) GetAreaSessionId() int32`

GetAreaSessionId returns the AreaSessionId field if non-nil, zero value otherwise.

### GetAreaSessionIdOk

`func (o *MbsN2MessageTransferReqData) GetAreaSessionIdOk() (*int32, bool)`

GetAreaSessionIdOk returns a tuple with the AreaSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaSessionId

`func (o *MbsN2MessageTransferReqData) SetAreaSessionId(v int32)`

SetAreaSessionId sets AreaSessionId field to given value.

### HasAreaSessionId

`func (o *MbsN2MessageTransferReqData) HasAreaSessionId() bool`

HasAreaSessionId returns a boolean if a field has been set.

### GetN2MbsSmInfo

`func (o *MbsN2MessageTransferReqData) GetN2MbsSmInfo() N2MbsSmInfo`

GetN2MbsSmInfo returns the N2MbsSmInfo field if non-nil, zero value otherwise.

### GetN2MbsSmInfoOk

`func (o *MbsN2MessageTransferReqData) GetN2MbsSmInfoOk() (*N2MbsSmInfo, bool)`

GetN2MbsSmInfoOk returns a tuple with the N2MbsSmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN2MbsSmInfo

`func (o *MbsN2MessageTransferReqData) SetN2MbsSmInfo(v N2MbsSmInfo)`

SetN2MbsSmInfo sets N2MbsSmInfo field to given value.


### GetSupportedFeatures

`func (o *MbsN2MessageTransferReqData) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *MbsN2MessageTransferReqData) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *MbsN2MessageTransferReqData) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.

### HasSupportedFeatures

`func (o *MbsN2MessageTransferReqData) HasSupportedFeatures() bool`

HasSupportedFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


