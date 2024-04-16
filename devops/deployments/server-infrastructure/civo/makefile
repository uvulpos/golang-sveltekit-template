apply:
	terraform init \
		-backend-config='bucket=iac-storage' \
		-backend-config='key=terraform.tfstate' \
		-backend-config='access_key=' \
		-backend-config='secret_key='
	@terraform apply -auto-approve
	@terraform output -raw k8skubeconfig > ../../.ci-pipeline/k8s.civo.config
	@chmod 600 ../../.ci-pipeline/k8s.civo.config
	
destroy:
	@terraform init
	@terraform destroy -auto-approve


