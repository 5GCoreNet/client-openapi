# MsisdnLessMoSmsNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedFeatures** | **string** | A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \&quot;0\&quot; to \&quot;9\&quot;,  \&quot;a\&quot; to \&quot;f\&quot; or \&quot;A\&quot; to \&quot;F\&quot; and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.  | 
**Sms** | **string** | String with format \&quot;byte\&quot; as defined in OpenAPI Specification, i.e, base64-encoded characters. | 
**ExternalId** | **string** | External identifier has the form username@realm. | 
**ApplicationPort** | **int32** | Unsigned integer with valid values between 0 and 65535. | 

## Methods

### NewMsisdnLessMoSmsNotification

`func NewMsisdnLessMoSmsNotification(supportedFeatures string, sms string, externalId string, applicationPort int32, ) *MsisdnLessMoSmsNotification`

NewMsisdnLessMoSmsNotification instantiates a new MsisdnLessMoSmsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsisdnLessMoSmsNotificationWithDefaults

`func NewMsisdnLessMoSmsNotificationWithDefaults() *MsisdnLessMoSmsNotification`

NewMsisdnLessMoSmsNotificationWithDefaults instantiates a new MsisdnLessMoSmsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedFeatures

`func (o *MsisdnLessMoSmsNotification) GetSupportedFeatures() string`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *MsisdnLessMoSmsNotification) GetSupportedFeaturesOk() (*string, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *MsisdnLessMoSmsNotification) SetSupportedFeatures(v string)`

SetSupportedFeatures sets SupportedFeatures field to given value.


### GetSms

`func (o *MsisdnLessMoSmsNotification) GetSms() string`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *MsisdnLessMoSmsNotification) GetSmsOk() (*string, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *MsisdnLessMoSmsNotification) SetSms(v string)`

SetSms sets Sms field to given value.


### GetExternalId

`func (o *MsisdnLessMoSmsNotification) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MsisdnLessMoSmsNotification) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MsisdnLessMoSmsNotification) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetApplicationPort

`func (o *MsisdnLessMoSmsNotification) GetApplicationPort() int32`

GetApplicationPort returns the ApplicationPort field if non-nil, zero value otherwise.

### GetApplicationPortOk

`func (o *MsisdnLessMoSmsNotification) GetApplicationPortOk() (*int32, bool)`

GetApplicationPortOk returns a tuple with the ApplicationPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPort

`func (o *MsisdnLessMoSmsNotification) SetApplicationPort(v int32)`

SetApplicationPort sets ApplicationPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


