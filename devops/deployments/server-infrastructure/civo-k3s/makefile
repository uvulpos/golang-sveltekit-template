apply:
	@terraform init
	@terraform apply -auto-approve
	@terraform output -raw k8skubeconfig > ../../.ci-pipeline/k8s.civo.config
	@chmod 600 ../../.ci-pipeline/k8s.civo.config
	
destroy:
	@terraform destroy -auto-approve