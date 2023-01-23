/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


type NetworkSliceInformationDocumentApi interface {

	/*
	NSSelectionGet Retrieve the Network Slice Selection Information

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiNSSelectionGetRequest
	*/
	NSSelectionGet(ctx context.Context) ApiNSSelectionGetRequest

	// NSSelectionGetExecute executes the request
	//  @return AuthorizedNetworkSliceInfo
	NSSelectionGetExecute(r ApiNSSelectionGetRequest) (*AuthorizedNetworkSliceInfo, *http.Response, error)
}

// NetworkSliceInformationDocumentApiService NetworkSliceInformationDocumentApi service
type NetworkSliceInformationDocumentApiService service

type ApiNSSelectionGetRequest struct {
	ctx context.Context
	ApiService NetworkSliceInformationDocumentApi
	nfType *NFType
	nfId *string
	sliceInfoRequestForRegistration *SliceInfoForRegistration
	sliceInfoRequestForPduSession *SliceInfoForPDUSession
	sliceInfoRequestForUeCu *SliceInfoForUEConfigurationUpdate
	homePlmnId *PlmnId
	tai *Tai
	supportedFeatures *string
}

// NF type of the NF service consumer
func (r ApiNSSelectionGetRequest) NfType(nfType NFType) ApiNSSelectionGetRequest {
	r.nfType = &nfType
	return r
}

// NF Instance ID of the NF service consumer
func (r ApiNSSelectionGetRequest) NfId(nfId string) ApiNSSelectionGetRequest {
	r.nfId = &nfId
	return r
}

// Requested network slice information during Registration procedure
func (r ApiNSSelectionGetRequest) SliceInfoRequestForRegistration(sliceInfoRequestForRegistration SliceInfoForRegistration) ApiNSSelectionGetRequest {
	r.sliceInfoRequestForRegistration = &sliceInfoRequestForRegistration
	return r
}

// Requested network slice information during PDU session establishment procedure
func (r ApiNSSelectionGetRequest) SliceInfoRequestForPduSession(sliceInfoRequestForPduSession SliceInfoForPDUSession) ApiNSSelectionGetRequest {
	r.sliceInfoRequestForPduSession = &sliceInfoRequestForPduSession
	return r
}

// Requested network slice information during UE confuguration update procedure
func (r ApiNSSelectionGetRequest) SliceInfoRequestForUeCu(sliceInfoRequestForUeCu SliceInfoForUEConfigurationUpdate) ApiNSSelectionGetRequest {
	r.sliceInfoRequestForUeCu = &sliceInfoRequestForUeCu
	return r
}

// PLMN ID of the HPLMN
func (r ApiNSSelectionGetRequest) HomePlmnId(homePlmnId PlmnId) ApiNSSelectionGetRequest {
	r.homePlmnId = &homePlmnId
	return r
}

// TAI of the UE
func (r ApiNSSelectionGetRequest) Tai(tai Tai) ApiNSSelectionGetRequest {
	r.tai = &tai
	return r
}

// Features required to be supported by the NFs in the target slice instance
func (r ApiNSSelectionGetRequest) SupportedFeatures(supportedFeatures string) ApiNSSelectionGetRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiNSSelectionGetRequest) Execute() (*AuthorizedNetworkSliceInfo, *http.Response, error) {
	return r.ApiService.NSSelectionGetExecute(r)
}

/*
NSSelectionGet Retrieve the Network Slice Selection Information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNSSelectionGetRequest
*/
func (a *NetworkSliceInformationDocumentApiService) NSSelectionGet(ctx context.Context) ApiNSSelectionGetRequest {
	return ApiNSSelectionGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthorizedNetworkSliceInfo
func (a *NetworkSliceInformationDocumentApiService) NSSelectionGetExecute(r ApiNSSelectionGetRequest) (*AuthorizedNetworkSliceInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizedNetworkSliceInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkSliceInformationDocumentApiService.NSSelectionGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/network-slice-information"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nfType == nil {
		return localVarReturnValue, nil, reportError("nfType is required and must be specified")
	}
	if r.nfId == nil {
		return localVarReturnValue, nil, reportError("nfId is required and must be specified")
	}

	localVarQueryParams.Add("nf-type", parameterToString(*r.nfType, ""))
	localVarQueryParams.Add("nf-id", parameterToString(*r.nfId, ""))
	if r.sliceInfoRequestForRegistration != nil {
		localVarQueryParams.Add("slice-info-request-for-registration", parameterToString(*r.sliceInfoRequestForRegistration, ""))
	}
	if r.sliceInfoRequestForPduSession != nil {
		localVarQueryParams.Add("slice-info-request-for-pdu-session", parameterToString(*r.sliceInfoRequestForPduSession, ""))
	}
	if r.sliceInfoRequestForUeCu != nil {
		localVarQueryParams.Add("slice-info-request-for-ue-cu", parameterToString(*r.sliceInfoRequestForUeCu, ""))
	}
	if r.homePlmnId != nil {
		localVarQueryParams.Add("home-plmn-id", parameterToString(*r.homePlmnId, ""))
	}
	if r.tai != nil {
		localVarQueryParams.Add("tai", parameterToString(*r.tai, ""))
	}
	if r.supportedFeatures != nil {
		localVarQueryParams.Add("supported-features", parameterToString(*r.supportedFeatures, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 307 {
			var v RedirectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 308 {
			var v RedirectResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 414 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
