package custombadrequestconstraints

type BadRequestContraint struct {
	KeyName string
	ConstraintMessage string
}

var (
	BadRequest_AuthenticationProviderInvalid *BadRequestContraint = &BadRequestContraint{ KeyName: "provider", ConstraintMessage: "The provider is a enum and has to be one of our supported list." }
)