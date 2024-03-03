# Application Deployment

here you can find all needed IaC Code to deploy your application to your kubernetes cluster

| Folder                | Purpose                                                                          |
| --------------------- | -------------------------------------------------------------------------------- |
| `cluster-deployment`  | Uses Terraform + Ansible to setup the K8s cluster                                |
| `core-infrastructure` | Deploys application dependencies via Helm as an ingress controller or a database |
| `services`            | Deploys the application services via Helm                                        |
