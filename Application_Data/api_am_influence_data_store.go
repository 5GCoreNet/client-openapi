/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Application_Data

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)


// AMInfluenceDataStoreApiService AMInfluenceDataStoreApi service
type AMInfluenceDataStoreApiService service

type ApiReadAmInfluenceDataRequest struct {
	ctx context.Context
	ApiService *AMInfluenceDataStoreApiService
	amInfluenceIds *[]string
	dnns *[]string
	snssais *[]Snssai
	dnnSnssaiInfos *[]DnnSnssaiInformation
	internalGroupIds *[]string
	supis *[]string
	anyUe *bool
	suppFeat *string
}

// Each element identifies a service.
func (r ApiReadAmInfluenceDataRequest) AmInfluenceIds(amInfluenceIds []string) ApiReadAmInfluenceDataRequest {
	r.amInfluenceIds = &amInfluenceIds
	return r
}

// Each element identifies a DNN.
func (r ApiReadAmInfluenceDataRequest) Dnns(dnns []string) ApiReadAmInfluenceDataRequest {
	r.dnns = &dnns
	return r
}

// Each element identifies a slice.
func (r ApiReadAmInfluenceDataRequest) Snssais(snssais []Snssai) ApiReadAmInfluenceDataRequest {
	r.snssais = &snssais
	return r
}

// Each element identifies a combination of (DNN, S-NSSAI).
func (r ApiReadAmInfluenceDataRequest) DnnSnssaiInfos(dnnSnssaiInfos []DnnSnssaiInformation) ApiReadAmInfluenceDataRequest {
	r.dnnSnssaiInfos = &dnnSnssaiInfos
	return r
}

// Each element identifies a group of users.
func (r ApiReadAmInfluenceDataRequest) InternalGroupIds(internalGroupIds []string) ApiReadAmInfluenceDataRequest {
	r.internalGroupIds = &internalGroupIds
	return r
}

// Each element identifies the user.
func (r ApiReadAmInfluenceDataRequest) Supis(supis []string) ApiReadAmInfluenceDataRequest {
	r.supis = &supis
	return r
}

// Indicates whether the request is for any UE.
func (r ApiReadAmInfluenceDataRequest) AnyUe(anyUe bool) ApiReadAmInfluenceDataRequest {
	r.anyUe = &anyUe
	return r
}

// Supported Features
func (r ApiReadAmInfluenceDataRequest) SuppFeat(suppFeat string) ApiReadAmInfluenceDataRequest {
	r.suppFeat = &suppFeat
	return r
}

func (r ApiReadAmInfluenceDataRequest) Execute() ([]AmInfluData, *http.Response, error) {
	return r.ApiService.ReadAmInfluenceDataExecute(r)
}

/*
ReadAmInfluenceData Retrieve AM Influence Data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReadAmInfluenceDataRequest
*/
func (a *AMInfluenceDataStoreApiService) ReadAmInfluenceData(ctx context.Context) ApiReadAmInfluenceDataRequest {
	return ApiReadAmInfluenceDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []AmInfluData
func (a *AMInfluenceDataStoreApiService) ReadAmInfluenceDataExecute(r ApiReadAmInfluenceDataRequest) ([]AmInfluData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AmInfluData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AMInfluenceDataStoreApiService.ReadAmInfluenceData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application-data/am-influence-data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.amInfluenceIds != nil {
		t := *r.amInfluenceIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("am-influence-ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("am-influence-ids", parameterToString(t, "multi"))
		}
	}
	if r.dnns != nil {
		t := *r.dnns
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("dnns", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("dnns", parameterToString(t, "multi"))
		}
	}
	if r.snssais != nil {
		localVarQueryParams.Add("snssais", parameterToString(*r.snssais, "csv"))
	}
	if r.dnnSnssaiInfos != nil {
		localVarQueryParams.Add("dnn-snssai-infos", parameterToString(*r.dnnSnssaiInfos, "csv"))
	}
	if r.internalGroupIds != nil {
		t := *r.internalGroupIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("internal-group-ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("internal-group-ids", parameterToString(t, "multi"))
		}
	}
	if r.supis != nil {
		t := *r.supis
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("supis", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("supis", parameterToString(t, "multi"))
		}
	}
	if r.anyUe != nil {
		localVarQueryParams.Add("any-ue", parameterToString(*r.anyUe, ""))
	}
	if r.suppFeat != nil {
		localVarQueryParams.Add("supp-feat", parameterToString(*r.suppFeat, ""))
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
