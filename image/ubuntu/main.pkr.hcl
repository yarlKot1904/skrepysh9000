packer {
  required_plugins {
    qemu = {
      version = "~> 1"
      source  = "github.com/hashicorp/qemu"
    }
  }
}

source "qemu" "jammy" {
#   accelerator      = "kvm"
  disk_compression = true
  disk_interface   = "virtio"
  disk_image       = true
  http_directory   = ".packer-http"
  cpus             = var.cpu
  memory           = var.ram
  disk_size        = var.disk_size
  format           = var.format
  headless         = var.headless
  iso_checksum     = var.iso_checksum
  iso_url          = var.iso_url
  net_device       = "virtio-net-pci"
  output_directory = "artifacts/qemu/${var.output_dir}"
  vm_name          = "${var.name}"
  communicator     = "ssh"
  shutdown_command = "echo '${var.ssh_password}' | sudo -S shutdown -P now"
  ssh_password     = var.ssh_password
  ssh_username     = var.ssh_username
  ssh_port         = 22
  ssh_timeout      = "1h"
  boot_wait        = "3m"

  cd_files = ["cloud-init/user-data", "cloud-init/meta-data", "cloud-init/network-config"]
  cd_label = "cidata"

  qemuargs = [
    ["-serial", "mon:stdio"],
  ]
}

build {
  sources = ["source.qemu.jammy"]

  provisioner "file" {
    source      = "../../skrepysh-agent/bin/skrepysh-agent"
    destination = "/tmp/skrepysh-agent"
  }
  provisioner "file" {
    source      = "scripts/install-packages.sh"
    destination = "/tmp/"
  }
  provisioner "file" {
    source      = "scripts/provision.sh"
    destination = "/tmp/"
  }

  provisioner "shell" {
    inline = [
      "sudo apt update",
      "sudo apt upgrade",

      "sudo mv /tmp/skrepysh-agent /usr/local/bin/skrepysh-agent",
      "sudo /bin/bash /tmp/install-packages.sh",
      "sudo rm /tmp/install-packages.sh",
      "sudo /bin/bash /tmp/provision.sh",
      "sudo rm /tmp/provision.sh",
    ]
  }
}