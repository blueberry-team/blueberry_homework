// internal/domain/user_role.go
package enum

type UserRole string

const (
    UserRoleBoss   UserRole = "boss"
    UserRoleWorker UserRole = "worker"
)

func IsUserRoleValid(role UserRole) bool {
	return role == UserRoleBoss || role == UserRoleWorker
}
