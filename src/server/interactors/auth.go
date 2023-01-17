package interactors

import (
	"context"
	"log"
	"time"

	_ "firebase.google.com/go/auth"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

	"github.com/sheodox/wellread/query"
	"github.com/sheodox/wellread/repositories"
)

type AuthInteractor struct {
	repo        *repositories.AuthRepository
	firebaseApp *firebase.App
}

func NewAuthInteractor() *AuthInteractor {
	opt := option.WithCredentialsFile("./service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing firebase app: %v", err)
	}

	return &AuthInteractor{repositories.Auth, app}
}

func (a *AuthInteractor) Login(idToken string) (query.User, error) {
	ctx := context.Background()
	userEntity := query.User{}

	client, err := a.firebaseApp.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting firebase auth client: %v\n", err)
	}

	token, err := client.VerifyIDTokenAndCheckRevoked(ctx, idToken)
	if err != nil {
		return userEntity, err
	}

	user, err := client.GetUser(ctx, token.UID)

	if err != nil {
		return userEntity, err
	}

	return a.repo.Add(repositories.UserAuthEntity{
		FirebaseUserId: user.UID,
		CreatedAt:      time.UnixMilli(user.UserMetadata.CreationTimestamp),
		ProviderId:     user.ProviderID,
		DisplayName:    user.DisplayName,
		Email:          user.Email,
	})
}

func (a *AuthInteractor) GetUser(userId int) (query.User, error) {
	return a.repo.Get(userId)
}
