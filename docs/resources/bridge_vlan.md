# mikrotik_bridge_vlan (Resource)
Adds VLAN aware Layer2 forwarding and VLAN tag modifications within the bridge.

## Example Usage
```terraform
resource "mikrotik_bridge" "default" {
  name = "main"
}

resource "mikrotik_bridge_vlan" "testacc" {
  bridge   = mikrotik_bridge.default.name
  tagged   = ["ether2", "vlan30"]
  untagged = ["ether3"]
  vlan_ids = [10, 30]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `bridge` (String) The bridge interface which the respective VLAN entry is intended for.

### Optional

- `tagged` (List of String) Interface list with a VLAN tag adding action in egress.
- `untagged` (List of String) Interface list with a VLAN tag removing action in egress.
- `vlan_ids` (List of Number) The list of VLAN IDs for certain port configuration. Ranges are not supported yet.

### Read-Only

- `id` (String) The ID of this resource.

## Import
Import is supported using the following syntax:
```shell
# import with id of bridge vlan
terraform import mikrotik_bridge_vlan.default "*2"
```