output "resource_group_id" {
  value       = azurerm_resource_group.rg.id
  description = "ID of the provisioned Azure Resource Group"
}

output "vnet_id" {
  value       = azurerm_virtual_network.vnet.id
  description = "ID of the provisioned Azure Virtual Network"
}
