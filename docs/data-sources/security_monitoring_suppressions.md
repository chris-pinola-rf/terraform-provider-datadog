---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "datadog_security_monitoring_suppressions Data Source - terraform-provider-datadog"
subcategory: ""
description: |-
  Use this data source to retrieve information about existing suppression rules, and use them in other resources.
---

# datadog_security_monitoring_suppressions (Data Source)

Use this data source to retrieve information about existing suppression rules, and use them in other resources.



<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `suppression_ids` (List of String) List of IDs of suppressions
- `suppressions` (List of Object) List of suppressions (see [below for nested schema](#nestedatt--suppressions))

<a id="nestedatt--suppressions"></a>
### Nested Schema for `suppressions`

Read-Only:

- `description` (String)
- `enabled` (Boolean)
- `expiration_date` (String)
- `id` (String)
- `name` (String)
- `rule_query` (String)
- `suppression_query` (String)