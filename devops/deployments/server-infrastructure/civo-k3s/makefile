run:
	@terraform apply
	@terraform output -raw k8skubeconfig > ../../.ci-pipeline/config
	
destroy:
	@terraform destroy