variable "cpu" {
  type    = string
  default = "8"
}

variable "disk_size" {
  type    = string
  default = "10G"
}

variable "skip_resize_disk" {
  type    = bool
  default = false
}

variable "headless" {
  type    = string
  default = "true"
}

variable "iso_checksum" {
  type    = string
  default = "file:https://cloud-images.ubuntu.com/releases/22.04/release/SHA256SUMS"
}

variable "iso_url" {
  type    = string
  default = "https://cloud-images.ubuntu.com/releases/22.04/release/ubuntu-22.04-server-cloudimg-amd64.img"
}

variable "output_dir" {
  type    = string
  default = "jammy"
}

variable "name" {
  type    = string
  default = "ubuntu-22.04.qcow2"
}

variable "ram" {
  type    = string
  default = "8192"
}

variable "ssh_password" {
  type    = string
  default = "ubuntu"
}

variable "ssh_username" {
  type    = string
  default = "ubuntu"
}

variable "version" {
  type    = string
  default = ""
}

variable "format" {
  type    = string
  default = "qcow2"
}