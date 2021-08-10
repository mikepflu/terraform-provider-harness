---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_application Resource - terraform-provider-harness"
subcategory: ""
description: |-
  Resource for creating a Harness application
---

# harness_application (Resource)

Resource for creating a Harness application



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The name of the application

### Optional

- **description** (String) The application description
- **is_manual_trigger_authorized** (Boolean) When this is set to true, all manual triggers will require API Key authorization

### Read-Only

- **git_sync_connector_id** (String) The id of the git sync connector
- **git_sync_enabled** (Boolean) True if git sync is enabled on this application
- **id** (String) Unique identifier of the application

