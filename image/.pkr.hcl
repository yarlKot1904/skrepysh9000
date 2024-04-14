source "qemu" "ubuntu-2204-amd64" {
  iso_url          = "https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64-disk-kvm.img"
  iso_checksum     = "sha256:83e1c1fbbc73bc8fa740c6175b3e74d9fa907de06111826dba8349c6daf43395"
  accelerator      = "kvm"
  memory           = "10000"
  disk_image       = false
  output_directory = "output-ubuntu-2204-amd64-qemu"
  disk_interface   = "virtio"
  format           = "qcow2"
  net_device       = "virtio-net"
  boot_wait        = "3s"
  http_directory   = "http-server"
  shutdown_command = "sudo shutdown -h now"
  ssh_username     = "ubuntu"
  ssh_password     = "ubuntu"
  ssh_timeout      = "60m"
}

build {
  sources = ["source.qemu.ubuntu-2204-amd64"]

  provisioner "shell" {
    script = "provisioning/init.sh"
  }

  post-processor "vagrant" {
    output = "ubuntu-2204-amd64.box"
  }
}

