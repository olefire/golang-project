package auth

import (
	"backend/internal/models"
	"context"
	"fmt"
)

type Repository interface {
	SignUpUser(ctx context.Context, signUpInput *models.User) (string, error)
	SignInUser(ctx context.Context, signInInput *models.SignInInput) (string, error)
}

type Deps struct {
	AuthRepo Repository
}

type Service struct {
	Deps
}

func NewService(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}

func (s *Service) SignUpUser(ctx context.Context, user *models.User) (string, error) {
	insertedId, err := s.AuthRepo.SignUpUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("can`t sign up user: %w", err)
	}

	return insertedId, err
}

//func (s *Service) SignUpUser(user *models.SignUpInput) (string, error) {
//	user.CreatedAt = time.Now()
//	user.UpdatedAt = user.CreatedAt
//	user.Verified = false
//	user.Role = "user"
//
//	hashedPassword, _ := auth.HashPassword(user.Password)
//	user.Password = hashedPassword
//	res, err := uc.collection.InsertOne(uc.ctx, &user)
//
//	if err != nil {
//		var er mongo.WriteException
//		if errors.As(err, &er) && er.WriteErrors[0].Code == 11000 {
//			return "", errors.New("user with that email already exist")
//		}
//		return nil, err
//	}
//
//	Create a unique index for the email field
//opt := options.Index()
//opt.SetUnique(true)
//index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}
//
//if _, err := uc.collection.Indexes().CreateOne(uc.ctx, index); err != nil {
//	return nil, errors.New("could not create index for email")
//}
//
//var newUser *models.UserDBResponse
//query := bson.M{"_id": res.InsertedID}
//
//err = uc.collection.FindOne(uc.ctx, query).Decode(&newUser)
//if err != nil {
//	return nil, err
//}
//
//return newUser, nil
//}

func (s *Service) SignInUser(context.Context, *models.SignInInput) (string, error) {
	return "", nil
}

func (s *Service) LogoutUser() {}
