/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)


type ServiceParameterDataStoreApi interface {

	/*
	ReadServiceParameterData Retrieve Service Parameter Data

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiReadServiceParameterDataRequest
	*/
	ReadServiceParameterData(ctx context.Context) ApiReadServiceParameterDataRequest

	// ReadServiceParameterDataExecute executes the request
	//  @return []ServiceParameterData
	ReadServiceParameterDataExecute(r ApiReadServiceParameterDataRequest) ([]ServiceParameterData, *http.Response, error)
}

// ServiceParameterDataStoreApiService ServiceParameterDataStoreApi service
type ServiceParameterDataStoreApiService service

type ApiReadServiceParameterDataRequest struct {
	ctx context.Context
	ApiService ServiceParameterDataStoreApi
	serviceParamIds *[]string
	dnns *[]string
	snssais *[]Snssai
	internalGroupIds *[]string
	supis *[]string
	ueIpv4s *[]string
	ueIpv6s *[]Ipv6Addr
	ueMacs *[]string
	anyUe *bool
	suppFeat *string
}

// Each element identifies a service.
func (r ApiReadServiceParameterDataRequest) ServiceParamIds(serviceParamIds []string) ApiReadServiceParameterDataRequest {
	r.serviceParamIds = &serviceParamIds
	return r
}

// Each element identifies a DNN.
func (r ApiReadServiceParameterDataRequest) Dnns(dnns []string) ApiReadServiceParameterDataRequest {
	r.dnns = &dnns
	return r
}

// Each element identifies a slice.
func (r ApiReadServiceParameterDataRequest) Snssais(snssais []Snssai) ApiReadServiceParameterDataRequest {
	r.snssais = &snssais
	return r
}

// Each element identifies a group of users.
func (r ApiReadServiceParameterDataRequest) InternalGroupIds(internalGroupIds []string) ApiReadServiceParameterDataRequest {
	r.internalGroupIds = &internalGroupIds
	return r
}

// Each element identifies the user.
func (r ApiReadServiceParameterDataRequest) Supis(supis []string) ApiReadServiceParameterDataRequest {
	r.supis = &supis
	return r
}

// Each element identifies the user.
func (r ApiReadServiceParameterDataRequest) UeIpv4s(ueIpv4s []string) ApiReadServiceParameterDataRequest {
	r.ueIpv4s = &ueIpv4s
	return r
}

// Each element identifies the user.
func (r ApiReadServiceParameterDataRequest) UeIpv6s(ueIpv6s []Ipv6Addr) ApiReadServiceParameterDataRequest {
	r.ueIpv6s = &ueIpv6s
	return r
}

// Each element identifies the user.
func (r ApiReadServiceParameterDataRequest) UeMacs(ueMacs []string) ApiReadServiceParameterDataRequest {
	r.ueMacs = &ueMacs
	return r
}

// Indicates whether the request is for any UE.
func (r ApiReadServiceParameterDataRequest) AnyUe(anyUe bool) ApiReadServiceParameterDataRequest {
	r.anyUe = &anyUe
	return r
}

// Supported Features
func (r ApiReadServiceParameterDataRequest) SuppFeat(suppFeat string) ApiReadServiceParameterDataRequest {
	r.suppFeat = &suppFeat
	return r
}

func (r ApiReadServiceParameterDataRequest) Execute() ([]ServiceParameterData, *http.Response, error) {
	return r.ApiService.ReadServiceParameterDataExecute(r)
}

/*
ReadServiceParameterData Retrieve Service Parameter Data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReadServiceParameterDataRequest
*/
func (a *ServiceParameterDataStoreApiService) ReadServiceParameterData(ctx context.Context) ApiReadServiceParameterDataRequest {
	return ApiReadServiceParameterDataRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ServiceParameterData
func (a *ServiceParameterDataStoreApiService) ReadServiceParameterDataExecute(r ApiReadServiceParameterDataRequest) ([]ServiceParameterData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ServiceParameterData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ServiceParameterDataStoreApiService.ReadServiceParameterData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/application-data/serviceParamData"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.serviceParamIds != nil {
		t := *r.serviceParamIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("service-param-ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("service-param-ids", parameterToString(t, "multi"))
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
	if r.ueIpv4s != nil {
		t := *r.ueIpv4s
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ue-ipv4s", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ue-ipv4s", parameterToString(t, "multi"))
		}
	}
	if r.ueIpv6s != nil {
		t := *r.ueIpv6s
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ue-ipv6s", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ue-ipv6s", parameterToString(t, "multi"))
		}
	}
	if r.ueMacs != nil {
		t := *r.ueMacs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ue-macs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ue-macs", parameterToString(t, "multi"))
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
