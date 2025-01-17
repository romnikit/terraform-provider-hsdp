# hsdp_notification_subscriber

Looks up HSDP Notification subscriber resources

## Example usage

```hcl
data "hsdp_notification_subscriber" "subscriber" {
  subscriber_id =  "example-d8f5-4fe4-b486-29a7fd30c9ba"
}
```

## Argument reference

* `subscriber_id` - The subscriber ID

## Attribute reference

* `managing_organization_id` - (Required) The UUID of the IAM organization or tenant
* `managing_organization` - (Required) The name of IAM organization or tenant
* `subscriber_product_name` - (Required) The name of the product
* `subscriber_service_name` - (Required) The name of the subscriber service
* `subscriber_service_instance_name` - The name of a service instance, used to differentiate multiple copies of the same service used in an organization
* `subscriber_service_base_url` - The base URL of the subscriber
* `subscriber_service_path_url` - The URL extension of the subscriber
* `description` - Description of the subscriber application
