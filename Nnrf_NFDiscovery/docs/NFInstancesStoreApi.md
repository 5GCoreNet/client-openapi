# \NFInstancesStoreApi

All URIs are relative to *https://example.com/nnrf-disc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchNFInstances**](NFInstancesStoreApi.md#SearchNFInstances) | **Get** /nf-instances | Search a collection of NF Instances



## SearchNFInstances

> SearchResult SearchNFInstances(ctx).TargetNfType(targetNfType).RequesterNfType(requesterNfType).AcceptEncoding(acceptEncoding).PreferredCollocatedNfTypes(preferredCollocatedNfTypes).RequesterNfInstanceId(requesterNfInstanceId).ServiceNames(serviceNames).RequesterNfInstanceFqdn(requesterNfInstanceFqdn).TargetPlmnList(targetPlmnList).RequesterPlmnList(requesterPlmnList).TargetNfInstanceId(targetNfInstanceId).TargetNfFqdn(targetNfFqdn).HnrfUri(hnrfUri).Snssais(snssais).RequesterSnssais(requesterSnssais).PlmnSpecificSnssaiList(plmnSpecificSnssaiList).RequesterPlmnSpecificSnssaiList(requesterPlmnSpecificSnssaiList).Dnn(dnn).Ipv4Index(ipv4Index).Ipv6Index(ipv6Index).NsiList(nsiList).SmfServingArea(smfServingArea).MbsmfServingArea(mbsmfServingArea).Tai(tai).AmfRegionId(amfRegionId).AmfSetId(amfSetId).Guami(guami).Supi(supi).UeIpv4Address(ueIpv4Address).IpDomain(ipDomain).UeIpv6Prefix(ueIpv6Prefix).PgwInd(pgwInd).PreferredPgwInd(preferredPgwInd).Pgw(pgw).PgwIp(pgwIp).Gpsi(gpsi).ExternalGroupIdentity(externalGroupIdentity).InternalGroupIdentity(internalGroupIdentity).PfdData(pfdData).DataSet(dataSet).RoutingIndicator(routingIndicator).GroupIdList(groupIdList).DnaiList(dnaiList).PduSessionTypes(pduSessionTypes).EventIdList(eventIdList).NwdafEventList(nwdafEventList).SupportedFeatures(supportedFeatures).UpfIwkEpsInd(upfIwkEpsInd).ChfSupportedPlmn(chfSupportedPlmn).PreferredLocality(preferredLocality).ExtPreferredLocality(extPreferredLocality).AccessType(accessType).Limit(limit).RequiredFeatures(requiredFeatures).ComplexQuery(complexQuery).MaxPayloadSize(maxPayloadSize).MaxPayloadSizeExt(maxPayloadSizeExt).AtsssCapability(atsssCapability).UpfUeIpAddrInd(upfUeIpAddrInd).ClientType(clientType).LmfId(lmfId).AnNodeType(anNodeType).RatType(ratType).PreferredTai(preferredTai).PreferredNfInstances(preferredNfInstances).IfNoneMatch(ifNoneMatch).TargetSnpn(targetSnpn).RequesterSnpnList(requesterSnpnList).AfEeData(afEeData).WAgfInfo(wAgfInfo).TngfInfo(tngfInfo).TwifInfo(twifInfo).TargetNfSetId(targetNfSetId).TargetNfServiceSetId(targetNfServiceSetId).NefId(nefId).NotificationType(notificationType).N1MsgClass(n1MsgClass).N2InfoClass(n2InfoClass).ServingScope(servingScope).Imsi(imsi).ImsPrivateIdentity(imsPrivateIdentity).ImsPublicIdentity(imsPublicIdentity).Msisdn(msisdn).PreferredApiVersions(preferredApiVersions).V2xSupportInd(v2xSupportInd).RedundantGtpu(redundantGtpu).RedundantTransport(redundantTransport).Ipups(ipups).ScpDomainList(scpDomainList).AddressDomain(addressDomain).Ipv4Addr(ipv4Addr).Ipv6Prefix(ipv6Prefix).ServedNfSetId(servedNfSetId).RemotePlmnId(remotePlmnId).RemoteSnpnId(remoteSnpnId).DataForwarding(dataForwarding).PreferredFullPlmn(preferredFullPlmn).RequesterFeatures(requesterFeatures).RealmId(realmId).StorageId(storageId).VsmfSupportInd(vsmfSupportInd).IsmfSupportInd(ismfSupportInd).NrfDiscUri(nrfDiscUri).PreferredVendorSpecificFeatures(preferredVendorSpecificFeatures).PreferredVendorSpecificNfFeatures(preferredVendorSpecificNfFeatures).RequiredPfcpFeatures(requiredPfcpFeatures).HomePubKeyId(homePubKeyId).ProseSupportInd(proseSupportInd).AnalyticsAggregationInd(analyticsAggregationInd).ServingNfSetId(servingNfSetId).ServingNfType(servingNfType).MlAnalyticsInfoList(mlAnalyticsInfoList).AnalyticsMetadataProvInd(analyticsMetadataProvInd).NsacfCapability(nsacfCapability).MbsSessionIdList(mbsSessionIdList).AreaSessionId(areaSessionId).GmlcNumber(gmlcNumber).UpfN6Ip(upfN6Ip).TaiList(taiList).PreferencesPrecedence(preferencesPrecedence).SupportOnboardingCapability(supportOnboardingCapability).UasNfFunctionalityInd(uasNfFunctionalityInd).V2xCapability(v2xCapability).ProseCapability(proseCapability).SharedDataId(sharedDataId).TargetHni(targetHni).TargetNwResolution(targetNwResolution).ExcludeNfinstList(excludeNfinstList).ExcludeNfservinstList(excludeNfservinstList).ExcludeNfservicesetList(excludeNfservicesetList).ExcludeNfsetList(excludeNfsetList).PreferredAnalyticsDelays(preferredAnalyticsDelays).Execute()

Search a collection of NF Instances

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/5GCoreNet/client-openapi/Nnrf_NFDiscovery"
)

func main() {
    targetNfType := *openapiclient.NewNFType() // NFType | Type of the target NF
    requesterNfType := *openapiclient.NewNFType() // NFType | Type of the requester NF
    acceptEncoding := "acceptEncoding_example" // string | Accept-Encoding, described in IETF RFC 7231 (optional)
    preferredCollocatedNfTypes := []openapiclient.CollocatedNfType{*openapiclient.NewCollocatedNfType()} // []CollocatedNfType | collocated NF types that candidate NFs should preferentially support (optional)
    requesterNfInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | NfInstanceId of the requester NF (optional)
    serviceNames := []openapiclient.ServiceName{*openapiclient.NewServiceName()} // []ServiceName | Names of the services offered by the NF (optional)
    requesterNfInstanceFqdn := "requesterNfInstanceFqdn_example" // string | FQDN of the requester NF (optional)
    targetPlmnList := []openapiclient.PlmnId{*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")} // []PlmnId | Id of the PLMN of either the target NF, or in SNPN scenario the Credentials Holder in the PLMN  (optional)
    requesterPlmnList := []openapiclient.PlmnId{*openapiclient.NewPlmnId("Mcc_example", "Mnc_example")} // []PlmnId | Id of the PLMN where the NF issuing the Discovery request is located (optional)
    targetNfInstanceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Identity of the NF instance being discovered (optional)
    targetNfFqdn := "targetNfFqdn_example" // string | FQDN of the NF instance being discovered (optional)
    hnrfUri := "hnrfUri_example" // string | Uri of the home NRF (optional)
    snssais := []openapiclient.Snssai{*openapiclient.NewSnssai(int32(123))} // []Snssai | Slice info of the target NF (optional)
    requesterSnssais := []openapiclient.ExtSnssai{*openapiclient.NewExtSnssai(int32(123))} // []ExtSnssai | Slice info of the requester NF (optional)
    plmnSpecificSnssaiList := []openapiclient.PlmnSnssai{*openapiclient.NewPlmnSnssai(*openapiclient.NewPlmnId("Mcc_example", "Mnc_example"), []openapiclient.ExtSnssai{*openapiclient.NewExtSnssai(int32(123))})} // []PlmnSnssai | PLMN specific Slice info of the target NF (optional)
    requesterPlmnSpecificSnssaiList := []openapiclient.PlmnSnssai{*openapiclient.NewPlmnSnssai(, []openapiclient.ExtSnssai{})} // []PlmnSnssai | PLMN-specific slice info of the NF issuing the Discovery request (optional)
    dnn := "dnn_example" // string | Dnn supported by the BSF, SMF or UPF (optional)
    ipv4Index := *openapiclient.NewIpIndex() // IpIndex | The IPv4 Index supported by the candidate UPF. (optional)
    ipv6Index := *openapiclient.NewIpIndex() // IpIndex | The IPv6 Index supported by the candidate UPF. (optional)
    nsiList := []string{"Inner_example"} // []string | NSI IDs that are served by the services being discovered (optional)
    smfServingArea := "smfServingArea_example" // string |  (optional)
    mbsmfServingArea := "mbsmfServingArea_example" // string |  (optional)
    tai := map[string][]openapiclient.Tai{ ... } // Tai | Tracking Area Identity (optional)
    amfRegionId := "amfRegionId_example" // string | AMF Region Identity (optional)
    amfSetId := "amfSetId_example" // string | AMF Set Identity (optional)
    guami := map[string][]openapiclient.Guami{ ... } // Guami | Guami used to search for an appropriate AMF (optional)
    supi := "supi_example" // string | SUPI of the user (optional)
    ueIpv4Address := "ueIpv4Address_example" // string | IPv4 address of the UE (optional)
    ipDomain := "ipDomain_example" // string | IP domain of the UE, which supported by BSF (optional)
    ueIpv6Prefix := "ueIpv6Prefix_example" // Ipv6Prefix | IPv6 prefix of the UE (optional)
    pgwInd := true // bool | Combined PGW-C and SMF or a standalone SMF (optional)
    preferredPgwInd := true // bool | Indicates combined PGW-C+SMF or standalone SMF are preferred (optional)
    pgw := "pgw_example" // string | PGW FQDN of a combined PGW-C and SMF (optional)
    pgwIp := map[string][]openapiclient.IpAddr{ ... } // IpAddr | PGW IP Address of a combined PGW-C and SMF (optional)
    gpsi := "gpsi_example" // string | GPSI of the user (optional)
    externalGroupIdentity := "externalGroupIdentity_example" // string | external group identifier of the user (optional)
    internalGroupIdentity := "internalGroupIdentity_example" // string | internal group identifier of the user (optional)
    pfdData := map[string][]openapiclient.PfdData{ ... } // PfdData | PFD data (optional)
    dataSet := *openapiclient.NewDataSetId() // DataSetId | data set supported by the NF (optional)
    routingIndicator := "routingIndicator_example" // string | routing indicator in SUCI (optional)
    groupIdList := []string{"Inner_example"} // []string | Group IDs of the NFs being discovered (optional)
    dnaiList := []string{"Inner_example"} // []string | Data network access identifiers of the NFs being discovered (optional)
    pduSessionTypes := []openapiclient.PduSessionType{*openapiclient.NewPduSessionType()} // []PduSessionType | list of PDU Session Type required to be supported by the target NF (optional)
    eventIdList := []openapiclient.EventId{*openapiclient.NewEventId()} // []EventId | Analytics event(s) requested to be supported by the Nnwdaf_AnalyticsInfo service  (optional)
    nwdafEventList := []openapiclient.NwdafEvent{*openapiclient.NewNwdafEvent()} // []NwdafEvent | Analytics event(s) requested to be supported by the Nnwdaf_EventsSubscription service.  (optional)
    supportedFeatures := "supportedFeatures_example" // string | Features required to be supported by the target NF (optional)
    upfIwkEpsInd := true // bool | UPF supporting interworking with EPS or not (optional)
    chfSupportedPlmn := map[string][]openapiclient.PlmnId{ ... } // PlmnId | PLMN ID supported by a CHF (optional)
    preferredLocality := "preferredLocality_example" // string | preferred target NF location (optional)
    extPreferredLocality := map[string][]openapiclient.LocalityDescription{"key": []openapiclient.LocalityDescription{*openapiclient.NewLocalityDescription(*openapiclient.NewLocalityType(), "LocalityValue_example")}} // map[string][]LocalityDescription | preferred target NF location A map (list of key-value pairs) where the key of the map represents the relative priority, for the requester, of each locality description among the list of locality descriptions in this query parameter, encoded as \"1\" (highest priority\"), \"2\", \"3\", …,  \"n\" (lowest priority)  (optional)
    accessType := openapiclient.AccessType("3GPP_ACCESS") // AccessType | AccessType supported by the target NF (optional)
    limit := int32(56) // int32 | Maximum number of NFProfiles to return in the response (optional)
    requiredFeatures := []string{"Inner_example"} // []string | Features required to be supported by the target NF (optional)
    complexQuery := openapiclient.ComplexQuery{Cnf: openapiclient.NewCnf([]openapiclient.CnfUnit{*openapiclient.NewCnfUnit([]openapiclient.Atom{*openapiclient.NewAtom("Attr_example", interface{}(123))})})} // ComplexQuery | the complex query condition expression (optional)
    maxPayloadSize := int32(56) // int32 | Maximum payload size of the response expressed in kilo octets (optional) (default to 124)
    maxPayloadSizeExt := int32(56) // int32 | Extended query for maximum payload size of the response expressed in kilo octets  (optional) (default to 124)
    atsssCapability := map[string][]openapiclient.AtsssCapability{ ... } // AtsssCapability | ATSSS Capability (optional)
    upfUeIpAddrInd := true // bool | UPF supporting allocating UE IP addresses/prefixes (optional)
    clientType := *openapiclient.NewExternalClientType() // ExternalClientType | Requested client type served by the NF (optional)
    lmfId := "lmfId_example" // string | LMF identification to be discovered (optional)
    anNodeType := *openapiclient.NewAnNodeType() // AnNodeType | Requested AN node type served by the NF (optional)
    ratType := *openapiclient.NewRatType() // RatType | Requested RAT type served by the NF (optional)
    preferredTai := map[string][]openapiclient.Tai{ ... } // Tai | preferred Tracking Area Identity (optional)
    preferredNfInstances := []string{"Inner_example"} // []string | preferred NF Instances (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Validator for conditional requests, as described in IETF RFC 7232, 3.2 (optional)
    targetSnpn := map[string][]openapiclient.PlmnIdNid{ ... } // PlmnIdNid | Target SNPN Identity, or the Credentials Holder in the SNPN (optional)
    requesterSnpnList := []openapiclient.PlmnIdNid{*openapiclient.NewPlmnIdNid("Mcc_example", "Mnc_example")} // []PlmnIdNid | SNPN ID(s) of the NF instance issuing the Discovery request (optional)
    afEeData := map[string][]openapiclient.AfEventExposureData{ ... } // AfEventExposureData | NEF exposured by the AF (optional)
    wAgfInfo := map[string][]openapiclient.WAgfInfo{ ... } // WAgfInfo | UPF collocated with W-AGF (optional)
    tngfInfo := map[string][]openapiclient.TngfInfo{ ... } // TngfInfo | UPF collocated with TNGF (optional)
    twifInfo := map[string][]openapiclient.TwifInfo{ ... } // TwifInfo | UPF collocated with TWIF (optional)
    targetNfSetId := "targetNfSetId_example" // string | Target NF Set ID (optional)
    targetNfServiceSetId := "targetNfServiceSetId_example" // string | Target NF Service Set ID (optional)
    nefId := "nefId_example" // string | NEF ID (optional)
    notificationType := *openapiclient.NewNotificationType() // NotificationType | Notification Type (optional)
    n1MsgClass := *openapiclient.NewN1MessageClass() // N1MessageClass | N1 Message Class (optional)
    n2InfoClass := *openapiclient.NewN2InformationClass() // N2InformationClass | N2 Information Class (optional)
    servingScope := []string{"Inner_example"} // []string | areas that can be served by the target NF (optional)
    imsi := "imsi_example" // string | IMSI of the requester UE to search for an appropriate NF (e.g. HSS) (optional)
    imsPrivateIdentity := "imsPrivateIdentity_example" // string | IMPI of the requester UE to search for a target HSS (optional)
    imsPublicIdentity := "imsPublicIdentity_example" // string | IMS Public Identity of the requester UE to search for a target HSS (optional)
    msisdn := "msisdn_example" // string | MSISDN of the requester UE to search for a target HSS (optional)
    preferredApiVersions := map[string]string{"key": "Inner_example"} // map[string]string | Preferred API version of the services to be discovered (optional)
    v2xSupportInd := true // bool | PCF supports V2X (optional)
    redundantGtpu := true // bool | UPF supports redundant gtp-u to be discovered (optional)
    redundantTransport := true // bool | UPF supports redundant transport path to be discovered (optional)
    ipups := true // bool | UPF which is configured for IPUPS functionality to be discovered (optional)
    scpDomainList := []string{"Inner_example"} // []string | SCP domains the target SCP or SEPP belongs to (optional)
    addressDomain := "addressDomain_example" // string | Address domain reachable through the SCP (optional)
    ipv4Addr := "ipv4Addr_example" // string | IPv4 address reachable through the SCP (optional)
    ipv6Prefix := "ipv6Prefix_example" // Ipv6Prefix | IPv6 prefix reachable through the SCP (optional)
    servedNfSetId := "servedNfSetId_example" // string | NF Set ID served by the SCP (optional)
    remotePlmnId := map[string][]openapiclient.PlmnId{ ... } // PlmnId | Id of the PLMN reachable through the SCP or SEPP (optional)
    remoteSnpnId := map[string][]openapiclient.PlmnIdNid{ ... } // PlmnIdNid | Id of the SNPN reachable through the SCP or SEPP (optional)
    dataForwarding := true // bool | UPF Instance(s) configured for data forwarding are requested (optional)
    preferredFullPlmn := true // bool | NF Instance(s) serving the full PLMN are preferred (optional)
    requesterFeatures := "requesterFeatures_example" // string | Features supported by the NF Service Consumer that is invoking the Nnrf_NFDiscovery service  (optional)
    realmId := "realmId_example" // string | realm-id to search for an appropriate UDSF (optional)
    storageId := "storageId_example" // string | storage-id to search for an appropriate UDSF (optional)
    vsmfSupportInd := true // bool | V-SMF capability supported by the target NF instance(s) (optional)
    ismfSupportInd := true // bool | I-SMF capability supported by the target NF instance(s) (optional)
    nrfDiscUri := "nrfDiscUri_example" // string | Uri of the NRF holding the NF profile of a target NF Instance (optional)
    preferredVendorSpecificFeatures := map[string]map[string][]VendorSpecificFeature{"key": map[string][]openapiclient.VendorSpecificFeature{"key": []openapiclient.VendorSpecificFeature{*openapiclient.NewVendorSpecificFeature("FeatureName_example", "FeatureVersion_example")}}} // map[string]map[string][]VendorSpecificFeature | Preferred vendor specific features of the services to be discovered (optional)
    preferredVendorSpecificNfFeatures := map[string][]openapiclient.VendorSpecificFeature{"key": []openapiclient.VendorSpecificFeature{*openapiclient.NewVendorSpecificFeature("FeatureName_example", "FeatureVersion_example")}} // map[string][]VendorSpecificFeature | Preferred vendor specific features of the network function to be discovered (optional)
    requiredPfcpFeatures := "requiredPfcpFeatures_example" // string | PFCP features required to be supported by the target UPF (optional)
    homePubKeyId := int32(56) // int32 | Indicates the Home Network Public Key ID which shall be able to be served by the NF instance  (optional)
    proseSupportInd := true // bool | PCF supports ProSe Capability (optional)
    analyticsAggregationInd := true // bool | analytics aggregation is supported by NWDAF or not (optional)
    servingNfSetId := "servingNfSetId_example" // string | NF Set Id served by target NF (optional)
    servingNfType := *openapiclient.NewNFType() // NFType | NF type served by the target NF (optional)
    mlAnalyticsInfoList := []openapiclient.MlAnalyticsInfo{*openapiclient.NewMlAnalyticsInfo()} // []MlAnalyticsInfo | Lisf of ML Analytics Filter information of Nnwdaf_MLModelProvision service (optional)
    analyticsMetadataProvInd := true // bool | analytics matadata provisioning is supported by NWDAF or not (optional)
    nsacfCapability := map[string][]openapiclient.NsacfCapability{ ... } // NsacfCapability | the service capability supported by the target NSACF (optional)
    mbsSessionIdList := []openapiclient.MbsSessionId{*openapiclient.NewMbsSessionId()} // []MbsSessionId | List of MBS Session ID(s) (optional)
    areaSessionId := int32(56) // int32 | Area Session ID (optional)
    gmlcNumber := "gmlcNumber_example" // string | The GMLC Number supported by the GMLC (optional)
    upfN6Ip := map[string][]openapiclient.IpAddr{ ... } // IpAddr | N6 IP address of PSA UPF supported by the EASDF (optional)
    taiList := []openapiclient.Tai{*openapiclient.NewTai(, "Tac_example")} // []Tai | Tracking Area Identifiers of the NFs being discovered (optional)
    preferencesPrecedence := []string{"Inner_example"} // []string | Indicates the precedence of the preference query parameters (from higher to lower)  (optional)
    supportOnboardingCapability := true // bool | Indicating the support for onboarding. (optional) (default to false)
    uasNfFunctionalityInd := true // bool | UAS NF functionality is supported by NEF or not (optional)
    v2xCapability := map[string][]openapiclient.V2xCapability{ ... } // V2xCapability | indicates the V2X capability that the target PCF needs to support. (optional)
    proseCapability := map[string][]openapiclient.ProSeCapability{ ... } // ProSeCapability | indicates the ProSe capability that the target PCF needs to support. (optional)
    sharedDataId := "sharedDataId_example" // string | Identifier of shared data stored in the NF being discovered (optional)
    targetHni := "targetHni_example" // string | Home Network Identifier query. (optional)
    targetNwResolution := true // bool | Resolution of the identity of the target PLMN based on the GPSI of the UE (optional)
    excludeNfinstList := []string{"Inner_example"} // []string | NF Instance IDs to be excluded from the NF Discovery procedure (optional)
    excludeNfservinstList := []openapiclient.NfServiceInstance{openapiclient.NfServiceInstance{Interface{}: new(interface{})}} // []NfServiceInstance | NF service instance IDs to be excluded from the NF Discovery procedure (optional)
    excludeNfservicesetList := []string{"Inner_example"} // []string | NF Service Set IDs to be excluded from the NF Discovery procedure (optional)
    excludeNfsetList := []string{"Inner_example"} // []string | NF Set IDs to be excluded from the NF Discovery procedure (optional)
    preferredAnalyticsDelays := map[string]int32{"key": int32(123)} // map[string]int32 | Preferred analytics delays supported by the NWDAF to be discovered (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFInstancesStoreApi.SearchNFInstances(context.Background()).TargetNfType(targetNfType).RequesterNfType(requesterNfType).AcceptEncoding(acceptEncoding).PreferredCollocatedNfTypes(preferredCollocatedNfTypes).RequesterNfInstanceId(requesterNfInstanceId).ServiceNames(serviceNames).RequesterNfInstanceFqdn(requesterNfInstanceFqdn).TargetPlmnList(targetPlmnList).RequesterPlmnList(requesterPlmnList).TargetNfInstanceId(targetNfInstanceId).TargetNfFqdn(targetNfFqdn).HnrfUri(hnrfUri).Snssais(snssais).RequesterSnssais(requesterSnssais).PlmnSpecificSnssaiList(plmnSpecificSnssaiList).RequesterPlmnSpecificSnssaiList(requesterPlmnSpecificSnssaiList).Dnn(dnn).Ipv4Index(ipv4Index).Ipv6Index(ipv6Index).NsiList(nsiList).SmfServingArea(smfServingArea).MbsmfServingArea(mbsmfServingArea).Tai(tai).AmfRegionId(amfRegionId).AmfSetId(amfSetId).Guami(guami).Supi(supi).UeIpv4Address(ueIpv4Address).IpDomain(ipDomain).UeIpv6Prefix(ueIpv6Prefix).PgwInd(pgwInd).PreferredPgwInd(preferredPgwInd).Pgw(pgw).PgwIp(pgwIp).Gpsi(gpsi).ExternalGroupIdentity(externalGroupIdentity).InternalGroupIdentity(internalGroupIdentity).PfdData(pfdData).DataSet(dataSet).RoutingIndicator(routingIndicator).GroupIdList(groupIdList).DnaiList(dnaiList).PduSessionTypes(pduSessionTypes).EventIdList(eventIdList).NwdafEventList(nwdafEventList).SupportedFeatures(supportedFeatures).UpfIwkEpsInd(upfIwkEpsInd).ChfSupportedPlmn(chfSupportedPlmn).PreferredLocality(preferredLocality).ExtPreferredLocality(extPreferredLocality).AccessType(accessType).Limit(limit).RequiredFeatures(requiredFeatures).ComplexQuery(complexQuery).MaxPayloadSize(maxPayloadSize).MaxPayloadSizeExt(maxPayloadSizeExt).AtsssCapability(atsssCapability).UpfUeIpAddrInd(upfUeIpAddrInd).ClientType(clientType).LmfId(lmfId).AnNodeType(anNodeType).RatType(ratType).PreferredTai(preferredTai).PreferredNfInstances(preferredNfInstances).IfNoneMatch(ifNoneMatch).TargetSnpn(targetSnpn).RequesterSnpnList(requesterSnpnList).AfEeData(afEeData).WAgfInfo(wAgfInfo).TngfInfo(tngfInfo).TwifInfo(twifInfo).TargetNfSetId(targetNfSetId).TargetNfServiceSetId(targetNfServiceSetId).NefId(nefId).NotificationType(notificationType).N1MsgClass(n1MsgClass).N2InfoClass(n2InfoClass).ServingScope(servingScope).Imsi(imsi).ImsPrivateIdentity(imsPrivateIdentity).ImsPublicIdentity(imsPublicIdentity).Msisdn(msisdn).PreferredApiVersions(preferredApiVersions).V2xSupportInd(v2xSupportInd).RedundantGtpu(redundantGtpu).RedundantTransport(redundantTransport).Ipups(ipups).ScpDomainList(scpDomainList).AddressDomain(addressDomain).Ipv4Addr(ipv4Addr).Ipv6Prefix(ipv6Prefix).ServedNfSetId(servedNfSetId).RemotePlmnId(remotePlmnId).RemoteSnpnId(remoteSnpnId).DataForwarding(dataForwarding).PreferredFullPlmn(preferredFullPlmn).RequesterFeatures(requesterFeatures).RealmId(realmId).StorageId(storageId).VsmfSupportInd(vsmfSupportInd).IsmfSupportInd(ismfSupportInd).NrfDiscUri(nrfDiscUri).PreferredVendorSpecificFeatures(preferredVendorSpecificFeatures).PreferredVendorSpecificNfFeatures(preferredVendorSpecificNfFeatures).RequiredPfcpFeatures(requiredPfcpFeatures).HomePubKeyId(homePubKeyId).ProseSupportInd(proseSupportInd).AnalyticsAggregationInd(analyticsAggregationInd).ServingNfSetId(servingNfSetId).ServingNfType(servingNfType).MlAnalyticsInfoList(mlAnalyticsInfoList).AnalyticsMetadataProvInd(analyticsMetadataProvInd).NsacfCapability(nsacfCapability).MbsSessionIdList(mbsSessionIdList).AreaSessionId(areaSessionId).GmlcNumber(gmlcNumber).UpfN6Ip(upfN6Ip).TaiList(taiList).PreferencesPrecedence(preferencesPrecedence).SupportOnboardingCapability(supportOnboardingCapability).UasNfFunctionalityInd(uasNfFunctionalityInd).V2xCapability(v2xCapability).ProseCapability(proseCapability).SharedDataId(sharedDataId).TargetHni(targetHni).TargetNwResolution(targetNwResolution).ExcludeNfinstList(excludeNfinstList).ExcludeNfservinstList(excludeNfservinstList).ExcludeNfservicesetList(excludeNfservicesetList).ExcludeNfsetList(excludeNfsetList).PreferredAnalyticsDelays(preferredAnalyticsDelays).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFInstancesStoreApi.SearchNFInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchNFInstances`: SearchResult
    fmt.Fprintf(os.Stdout, "Response from `NFInstancesStoreApi.SearchNFInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNFInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **targetNfType** | [**NFType**](NFType.md) | Type of the target NF | 
 **requesterNfType** | [**NFType**](NFType.md) | Type of the requester NF | 
 **acceptEncoding** | **string** | Accept-Encoding, described in IETF RFC 7231 | 
 **preferredCollocatedNfTypes** | [**[]CollocatedNfType**](CollocatedNfType.md) | collocated NF types that candidate NFs should preferentially support | 
 **requesterNfInstanceId** | **string** | NfInstanceId of the requester NF | 
 **serviceNames** | [**[]ServiceName**](ServiceName.md) | Names of the services offered by the NF | 
 **requesterNfInstanceFqdn** | **string** | FQDN of the requester NF | 
 **targetPlmnList** | [**[]PlmnId**](PlmnId.md) | Id of the PLMN of either the target NF, or in SNPN scenario the Credentials Holder in the PLMN  | 
 **requesterPlmnList** | [**[]PlmnId**](PlmnId.md) | Id of the PLMN where the NF issuing the Discovery request is located | 
 **targetNfInstanceId** | **string** | Identity of the NF instance being discovered | 
 **targetNfFqdn** | **string** | FQDN of the NF instance being discovered | 
 **hnrfUri** | **string** | Uri of the home NRF | 
 **snssais** | [**[]Snssai**](Snssai.md) | Slice info of the target NF | 
 **requesterSnssais** | [**[]ExtSnssai**](ExtSnssai.md) | Slice info of the requester NF | 
 **plmnSpecificSnssaiList** | [**[]PlmnSnssai**](PlmnSnssai.md) | PLMN specific Slice info of the target NF | 
 **requesterPlmnSpecificSnssaiList** | [**[]PlmnSnssai**](PlmnSnssai.md) | PLMN-specific slice info of the NF issuing the Discovery request | 
 **dnn** | **string** | Dnn supported by the BSF, SMF or UPF | 
 **ipv4Index** | [**IpIndex**](IpIndex.md) | The IPv4 Index supported by the candidate UPF. | 
 **ipv6Index** | [**IpIndex**](IpIndex.md) | The IPv6 Index supported by the candidate UPF. | 
 **nsiList** | **[]string** | NSI IDs that are served by the services being discovered | 
 **smfServingArea** | **string** |  | 
 **mbsmfServingArea** | **string** |  | 
 **tai** | [**Tai**](Tai.md) | Tracking Area Identity | 
 **amfRegionId** | **string** | AMF Region Identity | 
 **amfSetId** | **string** | AMF Set Identity | 
 **guami** | [**Guami**](Guami.md) | Guami used to search for an appropriate AMF | 
 **supi** | **string** | SUPI of the user | 
 **ueIpv4Address** | **string** | IPv4 address of the UE | 
 **ipDomain** | **string** | IP domain of the UE, which supported by BSF | 
 **ueIpv6Prefix** | **Ipv6Prefix** | IPv6 prefix of the UE | 
 **pgwInd** | **bool** | Combined PGW-C and SMF or a standalone SMF | 
 **preferredPgwInd** | **bool** | Indicates combined PGW-C+SMF or standalone SMF are preferred | 
 **pgw** | **string** | PGW FQDN of a combined PGW-C and SMF | 
 **pgwIp** | [**IpAddr**](IpAddr.md) | PGW IP Address of a combined PGW-C and SMF | 
 **gpsi** | **string** | GPSI of the user | 
 **externalGroupIdentity** | **string** | external group identifier of the user | 
 **internalGroupIdentity** | **string** | internal group identifier of the user | 
 **pfdData** | [**PfdData**](PfdData.md) | PFD data | 
 **dataSet** | [**DataSetId**](DataSetId.md) | data set supported by the NF | 
 **routingIndicator** | **string** | routing indicator in SUCI | 
 **groupIdList** | **[]string** | Group IDs of the NFs being discovered | 
 **dnaiList** | **[]string** | Data network access identifiers of the NFs being discovered | 
 **pduSessionTypes** | [**[]PduSessionType**](PduSessionType.md) | list of PDU Session Type required to be supported by the target NF | 
 **eventIdList** | [**[]EventId**](EventId.md) | Analytics event(s) requested to be supported by the Nnwdaf_AnalyticsInfo service  | 
 **nwdafEventList** | [**[]NwdafEvent**](NwdafEvent.md) | Analytics event(s) requested to be supported by the Nnwdaf_EventsSubscription service.  | 
 **supportedFeatures** | **string** | Features required to be supported by the target NF | 
 **upfIwkEpsInd** | **bool** | UPF supporting interworking with EPS or not | 
 **chfSupportedPlmn** | [**PlmnId**](PlmnId.md) | PLMN ID supported by a CHF | 
 **preferredLocality** | **string** | preferred target NF location | 
 **extPreferredLocality** | [**map[string][]LocalityDescription**](array.md) | preferred target NF location A map (list of key-value pairs) where the key of the map represents the relative priority, for the requester, of each locality description among the list of locality descriptions in this query parameter, encoded as \&quot;1\&quot; (highest priority\&quot;), \&quot;2\&quot;, \&quot;3\&quot;, …,  \&quot;n\&quot; (lowest priority)  | 
 **accessType** | [**AccessType**](AccessType.md) | AccessType supported by the target NF | 
 **limit** | **int32** | Maximum number of NFProfiles to return in the response | 
 **requiredFeatures** | **[]string** | Features required to be supported by the target NF | 
 **complexQuery** | [**ComplexQuery**](ComplexQuery.md) | the complex query condition expression | 
 **maxPayloadSize** | **int32** | Maximum payload size of the response expressed in kilo octets | [default to 124]
 **maxPayloadSizeExt** | **int32** | Extended query for maximum payload size of the response expressed in kilo octets  | [default to 124]
 **atsssCapability** | [**AtsssCapability**](AtsssCapability.md) | ATSSS Capability | 
 **upfUeIpAddrInd** | **bool** | UPF supporting allocating UE IP addresses/prefixes | 
 **clientType** | [**ExternalClientType**](ExternalClientType.md) | Requested client type served by the NF | 
 **lmfId** | **string** | LMF identification to be discovered | 
 **anNodeType** | [**AnNodeType**](AnNodeType.md) | Requested AN node type served by the NF | 
 **ratType** | [**RatType**](RatType.md) | Requested RAT type served by the NF | 
 **preferredTai** | [**Tai**](Tai.md) | preferred Tracking Area Identity | 
 **preferredNfInstances** | **[]string** | preferred NF Instances | 
 **ifNoneMatch** | **string** | Validator for conditional requests, as described in IETF RFC 7232, 3.2 | 
 **targetSnpn** | [**PlmnIdNid**](PlmnIdNid.md) | Target SNPN Identity, or the Credentials Holder in the SNPN | 
 **requesterSnpnList** | [**[]PlmnIdNid**](PlmnIdNid.md) | SNPN ID(s) of the NF instance issuing the Discovery request | 
 **afEeData** | [**AfEventExposureData**](AfEventExposureData.md) | NEF exposured by the AF | 
 **wAgfInfo** | [**WAgfInfo**](WAgfInfo.md) | UPF collocated with W-AGF | 
 **tngfInfo** | [**TngfInfo**](TngfInfo.md) | UPF collocated with TNGF | 
 **twifInfo** | [**TwifInfo**](TwifInfo.md) | UPF collocated with TWIF | 
 **targetNfSetId** | **string** | Target NF Set ID | 
 **targetNfServiceSetId** | **string** | Target NF Service Set ID | 
 **nefId** | **string** | NEF ID | 
 **notificationType** | [**NotificationType**](NotificationType.md) | Notification Type | 
 **n1MsgClass** | [**N1MessageClass**](N1MessageClass.md) | N1 Message Class | 
 **n2InfoClass** | [**N2InformationClass**](N2InformationClass.md) | N2 Information Class | 
 **servingScope** | **[]string** | areas that can be served by the target NF | 
 **imsi** | **string** | IMSI of the requester UE to search for an appropriate NF (e.g. HSS) | 
 **imsPrivateIdentity** | **string** | IMPI of the requester UE to search for a target HSS | 
 **imsPublicIdentity** | **string** | IMS Public Identity of the requester UE to search for a target HSS | 
 **msisdn** | **string** | MSISDN of the requester UE to search for a target HSS | 
 **preferredApiVersions** | **map[string]string** | Preferred API version of the services to be discovered | 
 **v2xSupportInd** | **bool** | PCF supports V2X | 
 **redundantGtpu** | **bool** | UPF supports redundant gtp-u to be discovered | 
 **redundantTransport** | **bool** | UPF supports redundant transport path to be discovered | 
 **ipups** | **bool** | UPF which is configured for IPUPS functionality to be discovered | 
 **scpDomainList** | **[]string** | SCP domains the target SCP or SEPP belongs to | 
 **addressDomain** | **string** | Address domain reachable through the SCP | 
 **ipv4Addr** | **string** | IPv4 address reachable through the SCP | 
 **ipv6Prefix** | **Ipv6Prefix** | IPv6 prefix reachable through the SCP | 
 **servedNfSetId** | **string** | NF Set ID served by the SCP | 
 **remotePlmnId** | [**PlmnId**](PlmnId.md) | Id of the PLMN reachable through the SCP or SEPP | 
 **remoteSnpnId** | [**PlmnIdNid**](PlmnIdNid.md) | Id of the SNPN reachable through the SCP or SEPP | 
 **dataForwarding** | **bool** | UPF Instance(s) configured for data forwarding are requested | 
 **preferredFullPlmn** | **bool** | NF Instance(s) serving the full PLMN are preferred | 
 **requesterFeatures** | **string** | Features supported by the NF Service Consumer that is invoking the Nnrf_NFDiscovery service  | 
 **realmId** | **string** | realm-id to search for an appropriate UDSF | 
 **storageId** | **string** | storage-id to search for an appropriate UDSF | 
 **vsmfSupportInd** | **bool** | V-SMF capability supported by the target NF instance(s) | 
 **ismfSupportInd** | **bool** | I-SMF capability supported by the target NF instance(s) | 
 **nrfDiscUri** | **string** | Uri of the NRF holding the NF profile of a target NF Instance | 
 **preferredVendorSpecificFeatures** | [**map[string]map[string][]VendorSpecificFeature**](map.md) | Preferred vendor specific features of the services to be discovered | 
 **preferredVendorSpecificNfFeatures** | [**map[string][]VendorSpecificFeature**](array.md) | Preferred vendor specific features of the network function to be discovered | 
 **requiredPfcpFeatures** | **string** | PFCP features required to be supported by the target UPF | 
 **homePubKeyId** | **int32** | Indicates the Home Network Public Key ID which shall be able to be served by the NF instance  | 
 **proseSupportInd** | **bool** | PCF supports ProSe Capability | 
 **analyticsAggregationInd** | **bool** | analytics aggregation is supported by NWDAF or not | 
 **servingNfSetId** | **string** | NF Set Id served by target NF | 
 **servingNfType** | [**NFType**](NFType.md) | NF type served by the target NF | 
 **mlAnalyticsInfoList** | [**[]MlAnalyticsInfo**](MlAnalyticsInfo.md) | Lisf of ML Analytics Filter information of Nnwdaf_MLModelProvision service | 
 **analyticsMetadataProvInd** | **bool** | analytics matadata provisioning is supported by NWDAF or not | 
 **nsacfCapability** | [**NsacfCapability**](NsacfCapability.md) | the service capability supported by the target NSACF | 
 **mbsSessionIdList** | [**[]MbsSessionId**](MbsSessionId.md) | List of MBS Session ID(s) | 
 **areaSessionId** | **int32** | Area Session ID | 
 **gmlcNumber** | **string** | The GMLC Number supported by the GMLC | 
 **upfN6Ip** | [**IpAddr**](IpAddr.md) | N6 IP address of PSA UPF supported by the EASDF | 
 **taiList** | [**[]Tai**](Tai.md) | Tracking Area Identifiers of the NFs being discovered | 
 **preferencesPrecedence** | **[]string** | Indicates the precedence of the preference query parameters (from higher to lower)  | 
 **supportOnboardingCapability** | **bool** | Indicating the support for onboarding. | [default to false]
 **uasNfFunctionalityInd** | **bool** | UAS NF functionality is supported by NEF or not | 
 **v2xCapability** | [**V2xCapability**](V2xCapability.md) | indicates the V2X capability that the target PCF needs to support. | 
 **proseCapability** | [**ProSeCapability**](ProSeCapability.md) | indicates the ProSe capability that the target PCF needs to support. | 
 **sharedDataId** | **string** | Identifier of shared data stored in the NF being discovered | 
 **targetHni** | **string** | Home Network Identifier query. | 
 **targetNwResolution** | **bool** | Resolution of the identity of the target PLMN based on the GPSI of the UE | 
 **excludeNfinstList** | **[]string** | NF Instance IDs to be excluded from the NF Discovery procedure | 
 **excludeNfservinstList** | [**[]NfServiceInstance**](NfServiceInstance.md) | NF service instance IDs to be excluded from the NF Discovery procedure | 
 **excludeNfservicesetList** | **[]string** | NF Service Set IDs to be excluded from the NF Discovery procedure | 
 **excludeNfsetList** | **[]string** | NF Set IDs to be excluded from the NF Discovery procedure | 
 **preferredAnalyticsDelays** | **map[string]int32** | Preferred analytics delays supported by the NWDAF to be discovered | 

### Return type

[**SearchResult**](SearchResult.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

