# Go API client for Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import Nhss_imsSDM "//"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), Nhss_imsSDM.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), Nhss_imsSDM.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), Nhss_imsSDM.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), Nhss_imsSDM.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nhss-ims-sdm/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CSLocationRetrievalApi* | [**GetLocCsDomain**](docs/CSLocationRetrievalApi.md#getloccsdomain) | **Get** /{imsUeId}/access-data/cs-domain/location-data | Retrieve the location data in CS domain
*CSRNRetrievalApi* | [**GetCsrn**](docs/CSRNRetrievalApi.md#getcsrn) | **Get** /{imsUeId}/access-data/cs-domain/csrn | Retrieve the routeing number in CS domain
*CSUserStateInfoRetrievalApi* | [**GetCsUserStateInfo**](docs/CSUserStateInfoRetrievalApi.md#getcsuserstateinfo) | **Get** /{imsUeId}/access-data/cs-domain/user-state | Retrieve the user state information in CS domain
*ChargingInfoRetrievalApi* | [**GetChargingInfo**](docs/ChargingInfoRetrievalApi.md#getcharginginfo) | **Get** /{imsUeId}/ims-data/profile-data/charging-info | Retrieve the charging information for to the user
*DSAIRegistrationInformationApi* | [**GetDsaiInfo**](docs/DSAIRegistrationInformationApi.md#getdsaiinfo) | **Get** /{imsUeId}/service-data/dsai | Retrieve the DSAI information associated to an Application Server
*DeleteRepositoryDataApi* | [**DeleteRepositoryDataServInd**](docs/DeleteRepositoryDataApi.md#deleterepositorydataservind) | **Delete** /{imsUeId}/repository-data/{serviceIndication} | delete the Repository Data for a Service Indication
*DeleteSMSRegistrationInformationApi* | [**DeleteSmsRegistrationInfo**](docs/DeleteSMSRegistrationInformationApi.md#deletesmsregistrationinfo) | **Delete** /{imsUeId}/service-data/sms-registration-info | delete the SMS registration information
*IFCsRetrievalApi* | [**GetIfcs**](docs/IFCsRetrievalApi.md#getifcs) | **Get** /{imsUeId}/ims-data/profile-data/ifcs | Retrieve the Initial Filter Criteria for the associated IMS subscription
*IMEISVRetrievalApi* | [**GetIMEISVInfo**](docs/IMEISVRetrievalApi.md#getimeisvinfo) | **Get** /{imsUeId}/identities/imeisv | Retrieve the IMEISV information
*IMSProfileDataRetrievalApi* | [**GetProfileData**](docs/IMSProfileDataRetrievalApi.md#getprofiledata) | **Get** /{imsUeId}/ims-data/profile-data | Retrieve the complete IMS profile for a given IMS public identity (and public identities in the same IRS) 
*IPAddressInfoRetrievalApi* | [**GetIpAddressInfo**](docs/IPAddressInfoRetrievalApi.md#getipaddressinfo) | **Get** /{imsUeId}/access-data/ps-domain/ip-address | Retrieve the IP address information
*PSLocationRetrievalApi* | [**GetLocPsDomain**](docs/PSLocationRetrievalApi.md#getlocpsdomain) | **Get** /{imsUeId}/access-data/ps-domain/location-data | Retrieve the location data in PS domain
*PSUserStateInfoRetrievalApi* | [**GetPsUserStateInfo**](docs/PSUserStateInfoRetrievalApi.md#getpsuserstateinfo) | **Get** /{imsUeId}/access-data/ps-domain/user-state | Retrieve the user state information in PS domain
*PriorityInfoRetrievalApi* | [**GetPriorityInfo**](docs/PriorityInfoRetrievalApi.md#getpriorityinfo) | **Get** /{imsUeId}/ims-data/profile-data/priority-levels | Retrieve the service priority levels associated to the user
*ReachabilitySubscriptionDeletionApi* | [**UeReachUnsubscribe**](docs/ReachabilitySubscriptionDeletionApi.md#uereachunsubscribe) | **Delete** /{imsUeId}/access-data/ps-domain/ue-reach-subscriptions/{subscriptionId} | unsubscribe from notifications to UE reachability
*ReachabilitySubscriptionModificationApi* | [**UeReachSubsModify**](docs/ReachabilitySubscriptionModificationApi.md#uereachsubsmodify) | **Patch** /{imsUeId}/access-data/ps-domain/ue-reach-subscriptions/{subscriptionId} | modify the subscription
*ReferenceLocationInfoRetrievalApi* | [**GetReferenceLocationInfo**](docs/ReferenceLocationInfoRetrievalApi.md#getreferencelocationinfo) | **Get** /{imsUeId}/access-data/wireline-domain/reference-location | Retrieve the reference location information
*RegistrationStatusRetrievalApi* | [**GetRegistrationStatus**](docs/RegistrationStatusRetrievalApi.md#getregistrationstatus) | **Get** /{imsUeId}/ims-data/registration-status | Retrieve the registration status of a user
*RepositoryDataApi* | [**GetRepositoryDataServInd**](docs/RepositoryDataApi.md#getrepositorydataservind) | **Get** /{imsUeId}/repository-data/{serviceIndication} | Retrieve the repository data associated to an IMPU and service indication
*RepositoryDataListApi* | [**GetRepositoryDataServIndList**](docs/RepositoryDataListApi.md#getrepositorydataservindlist) | **Get** /{imsUeId}/repository-data | Retrieve the repository data associated to an IMPU and service indication list
*RetrievalOfAssociatedIMSPrivateIdentitiesApi* | [**GetImsPrivateIds**](docs/RetrievalOfAssociatedIMSPrivateIdentitiesApi.md#getimsprivateids) | **Get** /{imsUeId}/identities/private-identities | Retrieve the associated private identities to the IMS public identity included in the service request 
*RetrievalOfAssociatedIMSPublicIdentitiesApi* | [**GetImsAssocIds**](docs/RetrievalOfAssociatedIMSPublicIdentitiesApi.md#getimsassocids) | **Get** /{imsUeId}/identities/ims-associated-identities | Retrieve the associated identities to the IMS public identity included in the service request 
*RetrievalOfPSIActivationStateApi* | [**GetPsiState**](docs/RetrievalOfPSIActivationStateApi.md#getpsistate) | **Get** /{imsUeId}/service-data/psi-status | Retrieve the PSI activation state data
*RetrievalOfSharedDataApi* | [**GetSharedData**](docs/RetrievalOfSharedDataApi.md#getshareddata) | **Get** /shared-data | retrieve shared data
*RetrievalOfTheAssociatedMsisdnsApi* | [**GetMsisdns**](docs/RetrievalOfTheAssociatedMsisdnsApi.md#getmsisdns) | **Get** /{imsUeId}/identities/msisdns | retrieve the Msisdns associated to requested identity
*RetrievalOfTheSCSCFCapabilitiesForTheIMSSubscriptionApi* | [**GetScscfCapabilities**](docs/RetrievalOfTheSCSCFCapabilitiesForTheIMSSubscriptionApi.md#getscscfcapabilities) | **Get** /{imsUeId}/ims-data/location-data/scscf-capabilities | Retrieve the S-CSCF capabilities for the associated IMS subscription
*RetrievalOfTheSCSCFSelectionAssistanceInformationForTheIMSSubscriptionApi* | [**GetScscfSelectionAssistanceInfo**](docs/RetrievalOfTheSCSCFSelectionAssistanceInformationForTheIMSSubscriptionApi.md#getscscfselectionassistanceinfo) | **Get** /{imsUeId}/ims-data/location-data/scscf-selection-assistance-info | Retrieve the S-CSCF selection assistance info
*RetrievalOfUESRVCCCapabilityAndSTNSRApi* | [**GetSrvccData**](docs/RetrievalOfUESRVCCCapabilityAndSTNSRApi.md#getsrvccdata) | **Get** /{imsUeId}/srvcc-data | Retrieve the srvcc data
*SDMSubscriptionCreationApi* | [**ImsSdmSubscribe**](docs/SDMSubscriptionCreationApi.md#imssdmsubscribe) | **Post** /{imsUeId}/subscriptions | subscribe to notifications
*SDMSubscriptionDeletionApi* | [**ImsSdmUnsubscribe**](docs/SDMSubscriptionDeletionApi.md#imssdmunsubscribe) | **Delete** /{imsUeId}/subscriptions/{subscriptionId} | unsubscribe from notifications
*SDMSubscriptionModificationApi* | [**ImsSdmSubsModify**](docs/SDMSubscriptionModificationApi.md#imssdmsubsmodify) | **Patch** /{imsUeId}/subscriptions/{subscriptionId} | modify the subscription
*SMSRegistrationInformationApi* | [**GetSmsRegistrationInfo**](docs/SMSRegistrationInformationApi.md#getsmsregistrationinfo) | **Get** /{imsUeId}/service-data/sms-registration-info | Retrieve the SMS registration information associated to a user
*ServerNameRetrievalApi* | [**GetServerName**](docs/ServerNameRetrievalApi.md#getservername) | **Get** /{imsUeId}/ims-data/location-data/server-name | Retrieve the server name for the associated user
*ServiceTraceInfoRetrievalApi* | [**GetServiceTraceInfo**](docs/ServiceTraceInfoRetrievalApi.md#getservicetraceinfo) | **Get** /{imsUeId}/ims-data/profile-data/service-level-trace-information | Retrieve the IMS service level trace information for the associated user
*SubscriptionCreationForSharedDataApi* | [**SubscribeToSharedData**](docs/SubscriptionCreationForSharedDataApi.md#subscribetoshareddata) | **Post** /shared-data-subscriptions | subscribe to notifications for shared data
*SubscriptionDeletionForSharedDataApi* | [**UnsubscribeForSharedData**](docs/SubscriptionDeletionForSharedDataApi.md#unsubscribeforshareddata) | **Delete** /shared-data-subscriptions/{subscriptionId} | unsubscribe from notifications for shared data
*SubscriptionModificationApi* | [**ModifySharedDataSubs**](docs/SubscriptionModificationApi.md#modifyshareddatasubs) | **Patch** /shared-data-subscriptions/{subscriptionId} | modify the subscription
*TADSInfoRetrievalApi* | [**GetTadsInfo**](docs/TADSInfoRetrievalApi.md#gettadsinfo) | **Get** /{imsUeId}/access-data/ps-domain/tads-info | Retrieve the T-ADS information
*UEIPReachabilitySubscriptionCreationApi* | [**UeReachIpSubscribe**](docs/UEIPReachabilitySubscriptionCreationApi.md#uereachipsubscribe) | **Post** /{imsUeId}/access-data/ps-domain/ue-reach-subscriptions | subscribe to notifications of UE reachability
*UpdateDsaiStateApi* | [**UpdateDsaiState**](docs/UpdateDsaiStateApi.md#updatedsaistate) | **Patch** /{imsUeId}/service-data/dsai | Patch
*UpdatePSIStateDataApi* | [**UpdatePsiState**](docs/UpdatePSIStateDataApi.md#updatepsistate) | **Patch** /{imsUeId}/service-data/psi-status | Patch
*UpdateRepositoryDataApi* | [**UpdateRepositoryDataServInd**](docs/UpdateRepositoryDataApi.md#updaterepositorydataservind) | **Put** /{imsUeId}/repository-data/{serviceIndication} | Update the repository data associated to an IMPU and service indication
*UpdateSMSRegistrationInfoApi* | [**UpdateSmsRegistrationInfo**](docs/UpdateSMSRegistrationInfoApi.md#updatesmsregistrationinfo) | **Put** /{imsUeId}/service-data/sms-registration-info | Update the SMS registration information associated to a user
*UpdateSRVCCDataApi* | [**UpdateSrvccData**](docs/UpdateSRVCCDataApi.md#updatesrvccdata) | **Patch** /{imsUeId}/srvcc-data | Patch


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AccessType](docs/AccessType.md)
 - [AccessType1](docs/AccessType1.md)
 - [AccessTypeAnyOf](docs/AccessTypeAnyOf.md)
 - [ActivationState](docs/ActivationState.md)
 - [ActivationStateAnyOf](docs/ActivationStateAnyOf.md)
 - [AmfLocationData](docs/AmfLocationData.md)
 - [ApplicationServer](docs/ApplicationServer.md)
 - [CallReferenceInfo](docs/CallReferenceInfo.md)
 - [CellGlobalId](docs/CellGlobalId.md)
 - [ChangeItem](docs/ChangeItem.md)
 - [ChangeType](docs/ChangeType.md)
 - [ChangeTypeAnyOf](docs/ChangeTypeAnyOf.md)
 - [ChargingInfo](docs/ChargingInfo.md)
 - [CoreNetworkServiceAuthorization](docs/CoreNetworkServiceAuthorization.md)
 - [CreatedUeReachabilitySubscription](docs/CreatedUeReachabilitySubscription.md)
 - [CsLocation](docs/CsLocation.md)
 - [CsUserState](docs/CsUserState.md)
 - [CsgInformation](docs/CsgInformation.md)
 - [Csrn](docs/Csrn.md)
 - [DataSetName](docs/DataSetName.md)
 - [DataSetNameAnyOf](docs/DataSetNameAnyOf.md)
 - [DetectingNode](docs/DetectingNode.md)
 - [DetectingNodeAnyOf](docs/DetectingNodeAnyOf.md)
 - [DsaiTagInformation](docs/DsaiTagInformation.md)
 - [DsaiTagStatus](docs/DsaiTagStatus.md)
 - [Ecgi](docs/Ecgi.md)
 - [EutraLocation](docs/EutraLocation.md)
 - [GNbId](docs/GNbId.md)
 - [GeraLocation](docs/GeraLocation.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [HeaderSipRequest](docs/HeaderSipRequest.md)
 - [IdentityType](docs/IdentityType.md)
 - [IdentityTypeAnyOf](docs/IdentityTypeAnyOf.md)
 - [Ifc](docs/Ifc.md)
 - [Ifcs](docs/Ifcs.md)
 - [ImeiSvInformation](docs/ImeiSvInformation.md)
 - [ImsAssociatedIdentities](docs/ImsAssociatedIdentities.md)
 - [ImsLocationData](docs/ImsLocationData.md)
 - [ImsProfileData](docs/ImsProfileData.md)
 - [ImsRegistrationState](docs/ImsRegistrationState.md)
 - [ImsRegistrationStateAnyOf](docs/ImsRegistrationStateAnyOf.md)
 - [ImsRegistrationStatus](docs/ImsRegistrationStatus.md)
 - [ImsSdmSubscription](docs/ImsSdmSubscription.md)
 - [ImsServiceProfile](docs/ImsServiceProfile.md)
 - [ImsVoiceOverPsSessionSupport](docs/ImsVoiceOverPsSessionSupport.md)
 - [ImsVoiceOverPsSessionSupportAnyOf](docs/ImsVoiceOverPsSessionSupportAnyOf.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [IpSmGwAddress](docs/IpSmGwAddress.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [LocationAreaId](docs/LocationAreaId.md)
 - [MmeLocationData](docs/MmeLocationData.md)
 - [ModificationNotification](docs/ModificationNotification.md)
 - [MsisdnList](docs/MsisdnList.md)
 - [NFType](docs/NFType.md)
 - [NFTypeAnyOf](docs/NFTypeAnyOf.md)
 - [Ncgi](docs/Ncgi.md)
 - [NotifyItem](docs/NotifyItem.md)
 - [NrLocation](docs/NrLocation.md)
 - [PatchItem](docs/PatchItem.md)
 - [PatchOperation](docs/PatchOperation.md)
 - [PatchOperationAnyOf](docs/PatchOperationAnyOf.md)
 - [PatchResult](docs/PatchResult.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [PriorityLevels](docs/PriorityLevels.md)
 - [PrivateIdentities](docs/PrivateIdentities.md)
 - [PrivateIdentity](docs/PrivateIdentity.md)
 - [PrivateIdentityType](docs/PrivateIdentityType.md)
 - [PrivateIdentityTypeAnyOf](docs/PrivateIdentityTypeAnyOf.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [PsLocation](docs/PsLocation.md)
 - [PsUserState](docs/PsUserState.md)
 - [PsiActivationState](docs/PsiActivationState.md)
 - [PublicIdentifier](docs/PublicIdentifier.md)
 - [PublicIdentities](docs/PublicIdentities.md)
 - [PublicIdentity](docs/PublicIdentity.md)
 - [RatType](docs/RatType.md)
 - [RatTypeAnyOf](docs/RatTypeAnyOf.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [ReferenceLocationInformation](docs/ReferenceLocationInformation.md)
 - [RegistrationType](docs/RegistrationType.md)
 - [RegistrationTypeAnyOf](docs/RegistrationTypeAnyOf.md)
 - [ReportItem](docs/ReportItem.md)
 - [RepositoryData](docs/RepositoryData.md)
 - [RepositoryDataList](docs/RepositoryDataList.md)
 - [RequestDirection](docs/RequestDirection.md)
 - [RequestDirectionAnyOf](docs/RequestDirectionAnyOf.md)
 - [RequestedNode](docs/RequestedNode.md)
 - [RequestedNodeAnyOf](docs/RequestedNodeAnyOf.md)
 - [RoutingAreaId](docs/RoutingAreaId.md)
 - [ScscfCapabilityList](docs/ScscfCapabilityList.md)
 - [ScscfSelectionAssistanceInformation](docs/ScscfSelectionAssistanceInformation.md)
 - [SdpDescription](docs/SdpDescription.md)
 - [ServiceAreaId](docs/ServiceAreaId.md)
 - [ServiceInformation](docs/ServiceInformation.md)
 - [ServiceInformationAnyOf](docs/ServiceInformationAnyOf.md)
 - [ServiceLevelTraceInformation](docs/ServiceLevelTraceInformation.md)
 - [SgsnLocationData](docs/SgsnLocationData.md)
 - [SharedData](docs/SharedData.md)
 - [SmsRegistrationInfo](docs/SmsRegistrationInfo.md)
 - [Snssai](docs/Snssai.md)
 - [Spt](docs/Spt.md)
 - [SrvccCapability](docs/SrvccCapability.md)
 - [SrvccCapabilityAnyOf](docs/SrvccCapabilityAnyOf.md)
 - [SrvccData](docs/SrvccData.md)
 - [TadsInformation](docs/TadsInformation.md)
 - [Tai](docs/Tai.md)
 - [TriggerPoint](docs/TriggerPoint.md)
 - [TwanLocationData](docs/TwanLocationData.md)
 - [TypeOfCondition](docs/TypeOfCondition.md)
 - [TypeOfConditionAnyOf](docs/TypeOfConditionAnyOf.md)
 - [UeReachabilityNotification](docs/UeReachabilityNotification.md)
 - [UeReachabilitySubscription](docs/UeReachabilitySubscription.md)
 - [UserStateCs](docs/UserStateCs.md)
 - [UserStateCsAnyOf](docs/UserStateCsAnyOf.md)
 - [UserStatePs](docs/UserStatePs.md)
 - [UserStatePsAnyOf](docs/UserStatePsAnyOf.md)
 - [UtraLocation](docs/UtraLocation.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nhss-ims-sdm**: Access to the Nhss IMS Subscription Data Management API
 - **nhss-ims-sdm:registration-status:read**: Access to read the Registration Status resource
 - **nhss-ims-sdm:profile-data:read**: Access to read the Profile Data resource
 - **nhss-ims-sdm:priority-levels:read**: Access to read the Priority Levels resource
 - **nhss-ims-sdm:ifcs:read**: Access to read the Initial Filter Criteria resource
 - **nhss-ims-sdm:service-level-trace-information:read**: Access to read the Service Level Trace Information resource
 - **nhss-ims-sdm:server-name:read**: Access to read the Server Name resource
 - **nhss-ims-sdm:scscf-capabilities:read**: Access to read the S-CSCF Capabilities resource
 - **nhss-ims-sdm:ps-domain:location-data:read**: Access to read the PS-Domain Location Data resource
 - **nhss-ims-sdm:ps-domain:ip-address:read**: Access to read the PS-Domain IP Address resource
 - **nhss-ims-sdm:ps-domain:tads-info:read**: Access to read the PS-Domain TADS Info resource
 - **nhss-ims-sdm:ps-domain:ue-reach-subscriptions:create**: Access to create PS-Domain UE Reachability Subscriptions resources
 - **nhss-ims-sdm:ps-domain:ue-reach-subscriptions:modify**: Access to update/delete a PS-Domain UE Reachability Subscription resource
 - **nhss-ims-sdm:ps-domain:user-state:read**: Access to read the PS-Domain User State resource
 - **nhss-ims-sdm:cs-domain:location-data:read**: Access to read the CS-Domain Location Data resource
 - **nhss-ims-sdm:cs-domain:user-state:read**: Access to read the CS-Domain User State resource
 - **nhss-ims-sdm:cs-domain:csrn:read**: Access to read the CS-Domain CSRN resource
 - **nhss-ims-sdm:wireline-domain:reference-location:read**: Access to read the Wireline-Domain Reference Location resource
 - **nhss-ims-sdm:repository-data:modify**: Access to create/update/delete the Repository Data resource
 - **nhss-ims-sdm:repository-data:read**: Access to read the Repository Data resource
 - **nhss-ims-sdm:identities:read**: Access to read the Identities resource
 - **nhss-ims-sdm:srvcc:read**: Access to read the SRVCC resource
 - **nhss-ims-sdm:srvcc:modify**: Acess to update the SRVCC resource
 - **nhss-ims-sdm:psi-status:read**: Access to read the PSI Status resource
 - **nhss-ims-sdm:psi-status:modify**: Acess to update the PSI Status resource
 - **nhss-ims-sdm:dsai:read**: Acess to read the DSAI resource
 - **nhss-ims-sdm:dsai:modify**: Acess to update the DSAI resource
 - **nhss-ims-sdm:sms-registration-info:read**: Acess to read the SMS Registration Info resource
 - **nhss-ims-sdm:sms-registration-info:modify**: Acess to create/update/delete the SMS Registration Info resource
 - **nhss-ims-sdm:subscriptions:create**: Access to create Subscriptions resources
 - **nhss-ims-sdm:subscription:modify**: Access to update/delete a Subscription resource
 - **nhss-ims-sdm:shared-subscriptions:create**: Access to create a Shared-Data Subscriptions resource
 - **nhss-ims-sdm:shared-subscription:modify**: Access to update/delete a Shared-Data Subscription resource
 - **nhss-ims-sdm:shared-data:read**: Access to read the Shared-Data resource
 - **nhss-ims-sdm:charging-info:read**: Access to read the ChargingInfo resource

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


