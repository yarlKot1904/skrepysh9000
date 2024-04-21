packer {
  required_plugins {
    qemu = {
      version = "~> 1"
      source  = "github.com/hashicorp/qemu"
    }
  }
}

source "qemu" "jammy" {
  accelerator      = "kvm"
  disk_compression = true
  disk_interface   = "virtio"
  disk_image       = true
  http_directory   = "http"
  cpus             = var.cpu
  memory           = var.ram
  disk_size        = var.disk_size
  format           = var.format
  headless         = var.headless
  iso_checksum     = var.iso_checksum
  iso_url          = var.iso_url
  net_device       = "virtio-net-pci"
  output_directory = "artifacts/qemu/${var.name}${var.version}"
  qemuargs         = [
    ["-cdrom", "cidata.iso"]
  ]
  communicator     = "ssh"
  shutdown_command = "echo '${var.ssh_password}' | sudo -S shutdown -P now"
  ssh_password     = var.ssh_password
  ssh_username     = var.ssh_username
  ssh_port         = 22
  ssh_timeout      = "1m"
}

build {
  sources = ["source.qemu.jammy"]

  provisioner "shell" {
    execute_command = "{{ .Vars }} sudo -E bash '{{ .Path }}'"
    inline          = ["sudo apt update", "sudo apt install python3"]
  }

  post-processor "shell-local" {
    environment_vars = ["IMAGE_NAME=${var.name}", "IMAGE_VERSION=${var.version}", "IMAGE_FORMAT=${var.format}"]
    script           = "scripts/prepare-image.sh"
  }
}