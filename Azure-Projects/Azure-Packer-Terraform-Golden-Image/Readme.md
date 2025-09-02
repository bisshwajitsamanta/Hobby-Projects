# Golden Image + Terraform VM Deployment on Azure

This project demonstrates how to:
- Register Azure subscription providers
- Build a golden Ubuntu image with **Packer**
- Store Terraform state in an **Azure Storage Account**
- Deploy a VM with **Terraform** using the golden image
- Authenticate securely with **Azure CLI** (no client secrets committed)

---

## ⚙️ Prerequisites

- [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli)
- [Packer](https://developer.hashicorp.com/packer/downloads)
- [Terraform](https://developer.hashicorp.com/terraform/downloads)
- A valid Azure subscription (e.g. Visual Studio subscription or personal free account)

---

## 🔐 Authenticate with Azure

Login with the correct account:

```bash
az login
az account set --subscription <YOUR_SUBSCRIPTION_ID>
export ARM_SUBSCRIPTION_ID="<YOUR_SUBSCRIPTION_ID>"
az account show --query "{id:id, name:name, user:user}" -o jsonc
```

## 🛠️ One-time Setup Script

Run the helper scripts to:

- Create a Resource Group for Packer
- Register `Microsoft.Storage` provider
- Create a Storage Account + container for Terraform state
- Generate `.env.azure` with reusable environment variables

```bash
./init-azure-packer.sh rg-packer-demo eastus
./setup-tf-backend.sh rg-packer-demo eastus
```
This creates:

- `spn.json` → Service Principal creds for Packer (**ignored by Git**)
- `.env.azure` → Handy environment exports, safe for local use
- `backend.hcl` → Terraform remote backend config

## 🔑 Generate SSH Keys
If you don’t already have an SSH key:

```bash
ssh-keygen -t ed25519 -C "azure-terraform" -f ~/.ssh/id_ed25519
```

Keys are created in:

- ~/.ssh/id_ed25519 (private key)
- ~/.ssh/id_ed25519.pub (public key)


## 🖼️ Build Golden Image with Packer

From the packer/ folder:
``` bash
packer init .
packer build \
  -var "client_id=$ARM_CLIENT_ID" \
  -var "client_secret=$ARM_CLIENT_SECRET" \
  -var "subscription_id=$ARM_SUBSCRIPTION_ID" \
  -var "tenant_id=$ARM_TENANT_ID" \
  -var "resource_group=rg-packer-demo" \
  -var "location=eastus" \
  -var "vm_size=Standard_B1s" \
  ubuntu.pkr.hcl
```
This produces a managed image named ubuntu-golden in `rg-packer-demo`.

## 🌐 Deploy VM with Terraform

- Re-initialize Terraform backend:

    ```bash
    cd terraform
    terraform init -reconfigure -backend-config=../backend.hcl
    ```

- Apply resources:

    ``` bash

    terraform apply \
    -var "subscription_id=$ARM_SUBSCRIPTION_ID" \
    -var "resource_group_image=rg-packer-demo" \
    -var "location=eastus" \
    -var "image_name=ubuntu-golden" \
    -var "admin_username=azureuser" \
    -var "ssh_public_key=$(cat ~/.ssh/id_ed25519.pub)"

    ```
- Terraform will output the VM’s public IP:
    ``` bash
    Apply complete!
    Outputs:
    vm_public_ip = "20.xx.xx.xx"
    ```
- SSH into the VM:
    ``` bash
    ssh -i ~/.ssh/id_ed25519 azureuser@$(terraform output -raw vm_public_ip)
    ```
## 🧹 Cleanup

- To destroy Terraform resources:
    ``` bash
        terraform destroy -auto-approve
    ```
- To delete the Packer Resource Group and state backend:
    ``` bash
        az group delete -n rg-packer-demo -y
    ```

## 📂 Git Hygiene
- Add sensitive/local files to .gitignore:
    ``` bash
        spn.json
        .env.azure
        backend.hcl
        terraform.tfstate
        *.tfstate*
        .terraform/
    ```
   - Never commit your subscription ID, client secrets, or SSH private keys.

## ✅ Summary

You now have:
- A golden image built by Packer
- A VM deployed by Terraform using that image
- A remote backend for state stored securely in Azure Storage
- CLI-based authentication (no secrets in GitHub)