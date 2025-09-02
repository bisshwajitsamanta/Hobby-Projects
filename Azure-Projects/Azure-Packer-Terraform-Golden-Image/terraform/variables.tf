variable "location" {
  description = "Azure region for the VM/resources"
  default     = "eastus"
}

variable "resource_group_image" {
  description = "Resource group where Packer stored the managed image"
  default     = "rg-packer-demo"
}

variable "image_name" {
  description = "Managed image name produced by Packer"
  default     = "ubuntu-golden"
}

variable "admin_username" {
  default = "azureuser"
}

variable "ssh_public_key" {
  description = "Your SSH public key (paste content of ~/.ssh/id_rsa.pub or similar)"
}

variable "subscription_id" {
  description = "Azure Subscription ID"
}