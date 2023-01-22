/*
VAE_PC5ProvisioningRequirement

Testing IndividualPC5ProvisioningRequirementSubscriptionDocumentApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package VAE_PC5ProvisioningRequirement

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_VAE_PC5ProvisioningRequirement_IndividualPC5ProvisioningRequirementSubscriptionDocumentApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test IndividualPC5ProvisioningRequirementSubscriptionDocumentApiService DeletePC5ProvisioningRequirementSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.DeletePC5ProvisioningRequirementSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test IndividualPC5ProvisioningRequirementSubscriptionDocumentApiService ReadPC5ProvisioningRequirementSubscription", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var subscriptionId string

        resp, httpRes, err := apiClient.IndividualPC5ProvisioningRequirementSubscriptionDocumentApi.ReadPC5ProvisioningRequirementSubscription(context.Background(), subscriptionId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}