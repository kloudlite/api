variable "agent-node-data" {
  type = map(object({
    name = string
    ip = string
  }))
  default = {

  }
}
variable "agent-nodes" {
  type = set(string)
  default = []
}
variable "cluster-id" {
  default = "{{ .ClusterID }}"
}
variable "do-token" {
  default = "{{ .DOToken }}"
}
variable "keys-path" {
  default = "{{ .KeysPath }}"
}
variable "do-image-id" {
  default = "{{ .DOImageID }}"
}
variable "region" {
  default = "{{ .Region }}"
}
