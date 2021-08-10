package provider

import (
	"fmt"
	"testing"

	"github.com/harness-io/harness-go-sdk/harness/api/cac"
	"github.com/harness-io/harness-go-sdk/harness/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/require"
)

func TestAccResourceAWSLambdaService(t *testing.T) {

	var (
		name               = fmt.Sprintf("%s_%s", t.Name(), utils.RandStringBytes(12))
		description        = "some description"
		updatedDescription = "updated description"
		resourceName       = "harness_service_aws_lambda.test"
	)

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		CheckDestroy:      testAccServiceDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAWSLambdaService(name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAWSLambdaServiceExists(t, resourceName, name, description),
				),
			},
			{
				Config: testAccResourceAWSLambdaService(name, updatedDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					testAccCheckAWSLambdaServiceExists(t, resourceName, name, updatedDescription),
				),
			},
		},
	})
}

func testAccCheckAWSLambdaServiceExists(t *testing.T, resourceName string, name string, description string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		svc, err := testAccGetService(resourceName, state)
		require.NoError(t, err)
		require.NotNil(t, svc)
		require.Equal(t, name, svc.Name)
		require.Equal(t, cac.ArtifactTypes.AWSLambda, svc.ArtifactType)
		require.Equal(t, cac.DeploymentTypes.AWSLambda, svc.DeploymentType)
		require.Equal(t, description, svc.Description)

		return nil
	}
}

func testAccResourceAWSLambdaService(name string, description string) string {
	return fmt.Sprintf(`
		resource "harness_application" "test" {
			name = "%[1]s"
		}

		resource "harness_service_aws_lambda" "test" {
			app_id = harness_application.test.id
			name = "%[1]s"
			description = "%[2]s"
		}

`, name, description)
}