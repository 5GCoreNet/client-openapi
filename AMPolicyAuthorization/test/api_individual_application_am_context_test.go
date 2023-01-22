/*
3gpp-am-policyauthorization

Testing IndividualApplicationAMContextApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package AMPolicyAuthorization

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_AMPolicyAuthorization_IndividualApplicationAMContextApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualApplicationAMContextApiService DeleteAppAmContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var appAmContextId string

        resp, httpRes, err := apiClient.IndividualApplicationAMContextApi.DeleteAppAmContext(context.Background(), afId, appAmContextId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationAMContextApiService GetAppAmContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var appAmContextId string

        resp, httpRes, err := apiClient.IndividualApplicationAMContextApi.GetAppAmContext(context.Background(), afId, appAmContextId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualApplicationAMContextApiService ModAppAmContext", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var afId string
        var appAmContextId string

        resp, httpRes, err := apiClient.IndividualApplicationAMContextApi.ModAppAmContext(context.Background(), afId, appAmContextId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
