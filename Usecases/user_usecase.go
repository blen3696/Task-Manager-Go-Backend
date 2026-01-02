package usecases

import (
	"errors"
	"github.com/blen/task_manager_api/Repositories"
	"github.com/blen/task_manager_api/domain"
)

type PasswordService interface {
	Hash(password string) (string, error)
	Compare(hash, password string) bool
}

type TokenService interface {
	Generate(userID, role string) (string, error)
}

type UserUsecase struct {
	repo     repositories.UserRepository
	password PasswordService
	token    TokenService
}

func NewUserUsecase(
	repo repositories.UserRepository,
	password PasswordService,
	token TokenService,
) *UserUsecase {
	return &UserUsecase{repo, password, token}
}

func (u *UserUsecase) Register(user *domain.User) (string, error) {
	hashed, err := u.password.Hash(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashed
	user.Role = "user"

	if err := u.repo.Create(user); err != nil {
		return "", err
	}

	return u.token.Generate(user.ID, user.Role)
}

func (u *UserUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !u.password.Compare(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return u.token.Generate(user.ID, user.Role)
}

func (u *UserUsecase) Promote(userID string) error {
	return u.repo.PromoteToAdmin(userID)
}
