---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_platform_connector_prometheus Resource - terraform-provider-harness"
subcategory: ""
description: |-
  Resource for creating a Prometheus connector.
---

# harness_platform_connector_prometheus (Resource)

Resource for creating a Prometheus connector.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `identifier` (String) Unique identifier of the resource.
- `name` (String) Name of the resource.
- `url` (String) Url of the Prometheus server.

### Optional

- `delegate_selectors` (Set of String) Connect using only the delegates which have these tags.
- `description` (String) Description of the resource.
- `git_sync` (Block List, Max: 1) (see [below for nested schema](#nestedblock--git_sync))
- `id` (String) The ID of this resource.
- `org_id` (String) Unique identifier of the organization.
- `project_id` (String) Unique identifier of the project.
- `tags` (Set of String) Tags to associate with the resource.

<a id="nestedblock--git_sync"></a>
### Nested Schema for `git_sync`

Optional:

- `branch` (String) The git branch to use for the resource.
- `file_path` (String) The file path of the resource.
- `repo_id` (String) Git sync config Id.
- `root_folder` (String) The root folder path of the resource.

Read-Only:

- `object_id` (String) The object id of the resource.

