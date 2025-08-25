package repository

import (
	"context"
	"database/sql"
	"errors"


	"github.com/sema/apiAuth/db"
	"github.com/sema/apiAuth/internal/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func GetUserByEmail(ctx context.Context, email string)(*models.User, error){

	var u models.User
	email = strings.ToLower(email)
	err :=db.DB.QueryRowContext(ctx, "SELECT id,email,senha FROM usuario.vw_credencial WHERE email = $1", email).Scan( &u.ID,&u.Email, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errors.ErrUnsupported
		}
		return nil, err
	}
	
	return &u, nil
}



func CreateUser(ctx context.Context, email, password string) error{
	email = strings.ToLower(email)
	
	_,err := GetUserByEmail(ctx, email)
	if err == nil{
		return errors.New("usuario já existe")
	}
	
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.DB.ExecContext(ctx, "INSERT INTO usuario.login (email, senha) VALUES ($1, $2)", email, hash)
	return err
}

func RecoverPassword(ctx context.Context, email string) error{
	return nil
}