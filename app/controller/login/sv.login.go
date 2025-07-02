package login

import (
	"app/app/model"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) checkPassword(hashed, plain string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}

func (s *Service) Login(ctx context.Context, user_number, password string) (*model.User, error) {
	user := new(model.User) // สร้าง user instance ก่อน

	err := s.db.NewSelect().Model(user).
		Where("user_number = ?", user_number).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !s.checkPassword(user.Password, password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
