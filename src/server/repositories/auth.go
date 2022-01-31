package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/sheodox/wellread/db"
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
	db *sqlx.DB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{db.Connection}
}

func (a *AuthRepository) Get(userId int) (*UserEntity, error) {
	user := UserEntity{}
	err := a.db.Get(&user, "select * from users where id=$1", userId)

	return &user, err
}

func (a *AuthRepository) Add(userAuth UserAuthEntity) (*UserEntity, error) {
	user := UserEntity{}
	err := a.db.Get(&user, "select * from users where firebase_user_id=$1", userAuth.FirebaseUserId)

	if err != nil {
		_, err = a.db.Exec(
			"insert into users (provider_id, firebase_user_id, email, display_name, created_at) values ($1, $2, $3, $4, $5)",
			userAuth.ProviderId, userAuth.FirebaseUserId, userAuth.Email, userAuth.DisplayName, userAuth.CreatedAt)

		if err != nil {
			return &user, err
		}
	} else {
		// todo update user if anything in userAuth is different than in user
	}

	err = a.db.Get(&user, "select * from users where firebase_user_id=$1", userAuth.FirebaseUserId)

	return &user, err
}
