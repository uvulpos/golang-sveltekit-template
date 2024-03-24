variable "civo_kubernetes_size" {}
variable "civo_kubernetes_count" {}

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
