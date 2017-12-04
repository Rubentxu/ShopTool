provider "azurerm" {

}


data "azurerm_platform_image" "test" {
  location  = "West Europe"
  publisher = "Canonical"
  offer     = "UbuntuServer"
  sku       = "16.04-LTS"
}

data "azurerm_public_ip" "datasourceip" {
    name = "testPublicIp"
    resource_group_name = "shoptoolgroup"
    depends_on = ["azurerm_virtual_machine.shoptoolvm"]
}

resource "azurerm_resource_group" "shoptoolgroup" {
    name     = "ShopToolGroup"
    location = "West Europe"

    tags {
        environment = "ShopTool tag"
    }
}

resource "azurerm_virtual_network" "shoptoolnetwork" {
    name                = "shoptoolVnet"
    address_space       = ["10.0.0.0/16"]
    location            = "West Europe"
    resource_group_name = "${azurerm_resource_group.shoptoolgroup.name}"

    tags {
        environment = "ShopTool tag"
    }
}

resource "azurerm_subnet" "shoptoolsubnet" {
    name                 = "shoptoolSubnet"
    resource_group_name  = "${azurerm_resource_group.shoptoolgroup.name}"
    virtual_network_name = "${azurerm_virtual_network.shoptoolnetwork.name}"
    address_prefix       = "10.0.2.0/24"
}

resource "azurerm_public_ip" "shoptoolpublicip" {
    name                         = "shoptoolPublicIP"
    location                     = "West Europe"
    resource_group_name          = "${azurerm_resource_group.shoptoolgroup.name}"
    public_ip_address_allocation = "dynamic"

    tags {
        environment = "ShopTool tag"
    }
}

resource "azurerm_network_security_group" "shoptoolpublicipnsg" {
    name                = "shoptoolNetworkSecurityGroup"
    location            = "West Europe"
    resource_group_name = "${azurerm_resource_group.shoptoolgroup.name}"

    security_rule {
        name                       = "SSH"
        priority                   = 1001
        direction                  = "Inbound"
        access                     = "Allow"
        protocol                   = "Tcp"
        source_port_range          = "*"
        destination_port_range     = "22"
        source_address_prefix      = "*"
        destination_address_prefix = "*"
    }

    tags {
        environment = "ShopTool tag"
    }
}

resource "azurerm_network_interface" "shoptoolnic" {
    name                = "shoptoolNIC"
    location            = "West Europe"
    resource_group_name = "${azurerm_resource_group.shoptoolgroup.name}"

    ip_configuration {
        name                          = "shoptoolNicConfiguration"
        subnet_id                     = "${azurerm_subnet.shoptoolsubnet.id}"
        private_ip_address_allocation = "dynamic"
        public_ip_address_id          = "${azurerm_public_ip.shoptoolpublicip.id}"
    }

    tags {
        environment = "ShopTool tag"
    }
}

resource "azurerm_virtual_machine" "shoptoolvm" {
    name                  = "shoptoolVM"
    location              = "West Europe"
    resource_group_name   = "${azurerm_resource_group.shoptoolgroup.name}"
    network_interface_ids = ["${azurerm_network_interface.shoptoolnic.id}"]
    vm_size               = "Standard_A2_v2"

    storage_os_disk {
        name              = "shoptoolOsDisk"
        caching           = "ReadWrite"
        create_option     = "FromImage"
        managed_disk_type = "Standard_LRS"
    }

    storage_image_reference {
        publisher = "Canonical"
        offer     = "UbuntuServer"
        sku       = "16.04.0-LTS"
        version   = "latest"
    }

    os_profile {
        computer_name  = "shoptoolvm"
        admin_username = "rubentxu"
    }

    os_profile_linux_config {
        disable_password_authentication = true
        ssh_keys {
            path     = "/home/rubentxu/.ssh/authorized_keys"
            key_data = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCtQjVRmpDSoVF/m9Dq7ahRahO8ftnYQm2GlV6ArNZ5AZUcbI8hvdJ4nKleUn/vHoR5TLbvxg5aCbohtq4Gn05yWjLkiya6SqHi2HTruNxJiOHnqzvKRRemVf/2rxkx91j4p1RfKY3RNWtf8zR2DJRdtEHJWJzmfSEOYmIBikL+KpBMkwc1sAgqOhecU9qKIpJKmgG1LbO5BzmnhL6FEeiGbpc1hi8zSEUqNFUpBqtyPBDkgCN9AxOcYPDob32iuq5KFClDN+KltOYip2WgF14kSkQU6qyglVfRUDWfZ/f5e2/BxNVe+e4M8JpVUR+niQHTIpJc3/aTEuXsBematRKd rubentxu@rubentxu-ubuntu17"
        }
    }

    tags {
        environment = "TShopTool tag"
    }
}


output "version" {
  value = "${data.azurerm_platform_image.test.version}"
}

output "ip_address" {
  value = "${data.azurerm_public_ip.datasourceip.ip_address}"
}
