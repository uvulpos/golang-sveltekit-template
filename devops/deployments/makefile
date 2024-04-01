.PHONY: base-applications

base-applications:
	@(cd ./base-applications/ ; make run)

infrastructure-civo:
	@(cd server-infrastructure/civo-k3s/ ; make apply)
	$(MAKE) base-applications

destroy-infrastructure-civo:
	@(cd server-infrastructure/civo-k3s/ ; make destroy)

# infrastructure-hetzner:
# infrastructure-azure:
