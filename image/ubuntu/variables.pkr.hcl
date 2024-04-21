variable "cpu" {
  type    = string
  default = "4"
}

variable "disk_size" {
  type    = string
  default = "10000"
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

variable "name" {
  type    = string
  default = "jammy"
}

variable "ram" {
  type    = string
  default = "4096"
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