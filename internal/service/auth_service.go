package service

//go:generate mockgen -source $GOFILE -destination ../../mock/service/mock_$GOFILE -package mock$GOPACKAGE

import (
	"context"
	"errors"

	"github.com/elangreza14/moviefestival/internal/domain"
	"github.com/elangreza14/moviefestival/internal/dto"
	"github.com/jackc/pgx/v5"
)

type (
	userRepo interface {
		Create(ctx context.Context, entities ...domain.User) error
		Get(ctx context.Context, by string, val any, columns ...string) (*domain.User, error)
	}

	tokenRepo interface {
		Create(ctx context.Context, entities ...domain.Token) error
		Get(ctx context.Context, by string, val any, columns ...string) (*domain.Token, error)
	}

	AuthService struct {
		UserRepo  userRepo
		TokenRepo tokenRepo
		config    ServiceConfig
	}
)

func NewAuthService(userRepo userRepo, tokenRepo tokenRepo, config ServiceConfig) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		TokenRepo: tokenRepo,
		config:    config,
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

	user, err = domain.NewUser(req.Email, req.Password, req.Name)
	if err != nil {
		return err
	}

	err = as.UserRepo.Create(ctx, *user)
	if err != nil {
		return err
	}

	return nil
}

func (as *AuthService) LoginUser(ctx context.Context, req dto.LoginPayload) (*dto.LoginResponse, error) {
	user, err := as.UserRepo.Get(ctx, "email", req.Email, "id", "password", "permission")
	if err != nil {
		return nil, err
	}

	ok := user.IsPasswordValid(req.Password)
	if !ok {
		return nil, errors.New("password not valid")
	}

	token, err := as.TokenRepo.Get(ctx, "user_id", user.ID, "token")
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, dto.ErrorNotFound{Entity: "user"}
		}
		return nil, err
	}

	user.LoadPermissions()

	if token != nil {
		_, err = token.IsTokenValid([]byte("test"))
		if err == nil {
			return &dto.LoginResponse{
				Token:       token.Token,
				Permissions: user.Permissions,
			}, nil
		}
	}

	token, err = domain.NewToken([]byte("test"), user.ID, "LOGIN")
	if err != nil {
		return nil, err
	}

	err = as.TokenRepo.Create(ctx, *token)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token:       token.Token,
		Permissions: user.Permissions,
	}, nil
}

func (as *AuthService) ProcessToken(ctx context.Context, reqToken string) (*domain.User, error) {
	token := &domain.Token{Token: reqToken}

	id, err := token.IsTokenValid([]byte(as.config.TokenSecret))
	if err != nil {
		return nil, err
	}

	token, err = as.TokenRepo.Get(ctx, "id", id, "user_id")
	if err != nil {
		return nil, err
	}

	return as.UserRepo.Get(ctx, "id", token.UserID, "id", "email", "permission")
}
