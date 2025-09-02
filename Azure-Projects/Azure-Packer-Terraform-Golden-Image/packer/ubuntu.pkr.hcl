# packer/ubuntu.pkr.hcl
packer {
  required_plugins {
    azure = {
      source  = "github.com/hashicorp/azure"
      version = ">= 1.0.0"
    }
  }
}

variable "client_id" {}
variable "client_secret" {}
variable "subscription_id" {}
variable "tenant_id" {}
variable "resource_group" { default = "rg-packer-demo" }
variable "location"       { default = "eastus" }
variable "vm_size"       { default = "Standard_B1s" }

source "azure-arm" "ubuntu" {
  subscription_id                    = var.subscription_id
  client_id                          = var.client_id
  client_secret                      = var.client_secret
  tenant_id                          = var.tenant_id
  build_resource_group_name          = var.resource_group
  os_type                            = "Linux"
  image_publisher                    = "Canonical"
  image_offer                        = "0001-com-ubuntu-server-jammy"
  image_sku                          = "22_04-lts"
  vm_size                            = var.vm_size

  managed_image_name                 = "ubuntu-golden"
  managed_image_resource_group_name  = var.resource_group
}

build {
  sources = ["source.azure-arm.ubuntu"]

  # Wait for cloud-init to finish + use sudo for apt
  provisioner "shell" {
    pause_before = "30s"
    inline = [
      # ensure cloud-init is done (won't fail the build if command missing)
      "sudo cloud-init status --wait || true",

      # be defensive about stale apt/dpkg locks
      "sudo rm -f /var/lib/dpkg/lock-frontend /var/lib/apt/lists/lock || true",

      # retry wrapper in case mirrors are flaky / locks still around
      "for i in {1..10}; do sudo apt-get update -y && break || (echo 'retrying apt-get update'; sleep 10); done",

      # noninteractive upgrade to avoid prompts
      "sudo DEBIAN_FRONTEND=noninteractive apt-get -yq upgrade",

      "echo '✅ Baseline image built successfully'"
    ]
  }
}