package service

import (
	"context"
	"errors"

	"github.com/elangreza14/moviefestival/dto"
	"github.com/elangreza14/moviefestival/model"
	"github.com/jackc/pgx/v5"
)

type (
	userRepo interface {
		Create(ctx context.Context, entities ...model.User) error
		Get(ctx context.Context, by string, val any, columns ...string) (*model.User, error)
	}

	tokenRepo interface {
		Create(ctx context.Context, entities ...model.Token) error
		Get(ctx context.Context, by string, val any, columns ...string) (*model.Token, error)
	}

	AuthService struct {
		UserRepo  userRepo
		TokenRepo tokenRepo
	}
)

func NewAuthService(userRepo userRepo, tokenRepo tokenRepo) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		TokenRepo: tokenRepo,
	}
}

func (as *AuthService) RegisterUser(ctx context.Context, req dto.RegisterPayload) error {
	user, err := as.UserRepo.Get(ctx, "email", req.Email, "id", "email")
	if err != nil && err != pgx.ErrNoRows {
		return err
	}

	if user != nil {
		return errors.New("email already exist")
	}

	user, err = model.NewUser(req.Email, req.Password, req.Name)
	if err != nil {
		return err
	}

	err = as.UserRepo.Create(ctx, *user)
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthService) LoginUser(ctx context.Context, req dto.LoginPayload) (string, error) {
	user, err := as.UserRepo.Get(ctx, "email", req.Email, "id", "password")
	if err != nil {
		return "", err
	}

	ok := user.IsPasswordValid(req.Password)
	if !ok {
		return "", errors.New("password not valid")
	}

	token, err := as.TokenRepo.Get(ctx, "user_id", user.ID, "token")
	if err != nil && err != pgx.ErrNoRows {
		return "", err
	}

	if token != nil {
		_, err = token.IsTokenValid([]byte("test"))
		if err == nil {
			return token.Token, nil
		}
	}

	token, err = model.NewToken([]byte("test"), user.ID, "LOGIN")
	if err != nil {
		return "", err
	}

	err = as.TokenRepo.Create(ctx, *token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}

func (as *AuthService) ProcessToken(ctx context.Context, reqToken string) (*model.User, error) {
	token := &model.Token{Token: reqToken}

	id, err := token.IsTokenValid([]byte("test"))
	if err != nil {
		return nil, err
	}

	token, err = as.TokenRepo.Get(ctx, "id", id, "user_id")
	if err != nil {
		return nil, err
	}

	return as.UserRepo.Get(ctx, "id", token.UserID)
}
