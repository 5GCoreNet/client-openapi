# OnboardingInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiInvokerPublicKey** | **string** | The API Invoker&#39;s public key | 
**ApiInvokerCertificate** | Pointer to **string** | The API Invoker&#39;s generic client certificate, provided by the CAPIF core function.  | [optional] 
**OnboardingSecret** | Pointer to **string** | The API Invoker&#39;s onboarding secret, provided by the CAPIF core function.  | [optional] 

## Methods

### NewOnboardingInformation

`func NewOnboardingInformation(apiInvokerPublicKey string, ) *OnboardingInformation`

NewOnboardingInformation instantiates a new OnboardingInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnboardingInformationWithDefaults

`func NewOnboardingInformationWithDefaults() *OnboardingInformation`

NewOnboardingInformationWithDefaults instantiates a new OnboardingInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiInvokerPublicKey

`func (o *OnboardingInformation) GetApiInvokerPublicKey() string`

GetApiInvokerPublicKey returns the ApiInvokerPublicKey field if non-nil, zero value otherwise.

### GetApiInvokerPublicKeyOk

`func (o *OnboardingInformation) GetApiInvokerPublicKeyOk() (*string, bool)`

GetApiInvokerPublicKeyOk returns a tuple with the ApiInvokerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerPublicKey

`func (o *OnboardingInformation) SetApiInvokerPublicKey(v string)`

SetApiInvokerPublicKey sets ApiInvokerPublicKey field to given value.


### GetApiInvokerCertificate

`func (o *OnboardingInformation) GetApiInvokerCertificate() string`

GetApiInvokerCertificate returns the ApiInvokerCertificate field if non-nil, zero value otherwise.

### GetApiInvokerCertificateOk

`func (o *OnboardingInformation) GetApiInvokerCertificateOk() (*string, bool)`

GetApiInvokerCertificateOk returns a tuple with the ApiInvokerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiInvokerCertificate

`func (o *OnboardingInformation) SetApiInvokerCertificate(v string)`

SetApiInvokerCertificate sets ApiInvokerCertificate field to given value.

### HasApiInvokerCertificate

`func (o *OnboardingInformation) HasApiInvokerCertificate() bool`

HasApiInvokerCertificate returns a boolean if a field has been set.

### GetOnboardingSecret

`func (o *OnboardingInformation) GetOnboardingSecret() string`

GetOnboardingSecret returns the OnboardingSecret field if non-nil, zero value otherwise.

### GetOnboardingSecretOk

`func (o *OnboardingInformation) GetOnboardingSecretOk() (*string, bool)`

GetOnboardingSecretOk returns a tuple with the OnboardingSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingSecret

`func (o *OnboardingInformation) SetOnboardingSecret(v string)`

SetOnboardingSecret sets OnboardingSecret field to given value.

### HasOnboardingSecret

`func (o *OnboardingInformation) HasOnboardingSecret() bool`

HasOnboardingSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


