variable "civo_kubernetes_size" {}
variable "civo_kubernetes_count" {}

# # Create Object Storage Credetials for Bucket to store backups
# data "civo_object_store_credential" "backup" {
#   name = "application-backup-bucket-credentials"
# }

# # Create Object Storage for Bucket to store backups
# resource "civo_object_store" "backup" {
#   name          = "application-backup-bucket"
#   max_size_gb   = 500
#   access_key_id = civo_object_store_credential.backup.access_key_id
# }

# # Create a managaed Database
# resource "civo_database" "custom_database" {
#   name    = "custom_database"
#   size    = element(data.civo_size.small.sizes, 0).name
#   nodes   = 2
#   engine  = element(data.civo_database_version.mysql.versions, 0).engine
#   version = element(data.civo_database_version.mysql.versions, 0).version
# }

# Create a firewall
resource "civo_firewall" "my-firewall" {
  name = "application-firewall"
}

# Create a firewall rule
resource "civo_firewall_rule" "kubernetes" {
  firewall_id = civo_firewall.my-firewall.id
  protocol    = "tcp"
  start_port  = "6443"
  end_port    = "6443"
  cidr        = ["0.0.0.0/0"]
  direction   = "ingress"
  label       = "kubernetes-api-server"
  action      = "allow"
}

# Create a cluster with k3s
resource "civo_kubernetes_cluster" "my-cluster" {
  name = "my-cluster"
  # applications = "Portainer,Linkerd:Linkerd & Jaeger"
  firewall_id  = civo_firewall.my-firewall.id
  cluster_type = "k3s"
  pools {
    label      = "myapplication" // Optional
    size       = var.civo_kubernetes_size
    node_count = var.civo_kubernetes_count
  }
}

output "k8skubeconfig" {
  sensitive = true
  value     = civo_kubernetes_cluster.my-cluster.kubeconfig
}
