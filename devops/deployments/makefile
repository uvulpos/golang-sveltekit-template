.PHONY: base-applications

base-applications:
	@(cd ./base-applications/ ; make run)

infrastructure-civo:
	@(cd server-infrastructure/civo/ ; make apply)

# $(MAKE) base-applications

destroy-infrastructure-civo:
	@(cd server-infrastructure/civo/ ; make destroy)

# infrastructure-hetzner:
# infrastructure-azure:
