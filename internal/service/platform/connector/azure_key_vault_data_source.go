package connector

import (
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceConnectorAzureKeyVault() *schema.Resource {
	resource := &schema.Resource{
		Description: "Datasource for looking up an Azure Cloud Provider Connector.",
		ReadContext: resourceConnectorAzureKeyVaultRead,

		Schema: map[string]*schema.Schema{
			"client_id": {
				Description: "Application ID of the Azure App.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"secret_key": {
				Description: "The Harness text secret with the Azure authentication key as its value.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tenant_id": {
				Description: "The Azure Active Directory (AAD) directory ID where you created your application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"vault_name": {
				Description: "The Azure Active Directory (AAD) directory ID where you created your application.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"subscription": {
				Description: "Azure Subscription ID.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_default": {
				Description: "Is default or not.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"azure_environment_type": {
				Description: "Specifies the Azure Environment type, which is AZURE by default. Can either be AZURE or AZURE_US_GOVERNMENT",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"delegate_selectors": {
				Description: "Connect using only the delegates which have these tags.",
				Type:        schema.TypeSet,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}

	helpers.SetMultiLevelDatasourceSchema(resource.Schema)

	return resource
}
