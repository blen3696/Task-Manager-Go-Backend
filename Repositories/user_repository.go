package repositories

import "github.com/blen/task_manager_api/domain"

type UserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	PromoteToAdmin(userID string) error
}
