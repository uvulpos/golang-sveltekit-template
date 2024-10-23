package sessionlocals

type SessionLocals string

const (
	UserUUID     SessionLocals = "user-uuid"
	Permissions  SessionLocals = "permissions"
	SessionUUID  SessionLocals = "session-uuid"
	AuthProvider SessionLocals = "auth-provider"
)
