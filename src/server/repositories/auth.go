package repositories

import (
	"context"
	"time"

	"github.com/sheodox/wellread/db"
	"github.com/sheodox/wellread/query"
)

type UserEntity struct {
	Id             int       `db:"id"`
	ProviderId     string    `db:"provider_id"`
	FirebaseUserId string    `db:"firebase_user_id"`
	Email          string    `db:"email"`
	DisplayName    string    `db:"display_name"`
	CreatedAt      time.Time `db:"created_at"`
}

type UserAuthEntity struct {
	ProviderId     string    `db:"provider_id"`
	FirebaseUserId string    `db:"firebase_user_id"`
	Email          string    `db:"email"`
	DisplayName    string    `db:"name"`
	CreatedAt      time.Time `db:"created_at"`
}

type AuthRepository struct {
	queries *query.Queries
	ctx     context.Context
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{db.Queries, context.Background()}
}

func (a *AuthRepository) Get(userId int) (query.User, error) {
	return a.queries.GetUser(a.ctx, int32(userId))
}

func (a *AuthRepository) Add(userAuth UserAuthEntity) (query.User, error) {
	_, err := a.queries.GetUserByFirebaseId(a.ctx, userAuth.FirebaseUserId)

	if err != nil {
		return a.queries.AddUser(a.ctx, query.AddUserParams{
			ProviderID:     userAuth.ProviderId,
			FirebaseUserID: userAuth.FirebaseUserId,
			Email:          userAuth.Email,
			DisplayName:    userAuth.DisplayName,
			CreatedAt:      time.Now(),
		})
	} else {
		// todo update user if anything in userAuth is different than in user
	}

	return a.queries.GetUserByFirebaseId(a.ctx, userAuth.FirebaseUserId)
}
