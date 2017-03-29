package handler

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	"github.com/micro/go-micro/errors"
	"github.com/bketelsen/microclass/module4/userservice/db"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"

	account "github.com/bketelsen/microclass/module4/userservice/proto/account"
)

const (
	x = "cruft123"
)

var (
	alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func random(i int) string {
	bytes := make([]byte, i)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		return string(bytes)
	}
	return "ughwhy?!!!"
}

type Account struct{}

func (s *Account) Create(ctx context.Context, req *account.CreateRequest, rsp *account.CreateResponse) error {
	salt := random(16)
	h, err := bcrypt.GenerateFromPassword([]byte(x+salt+req.Password), 10)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.user.Create", err.Error())
	}
	pp := base64.StdEncoding.EncodeToString(h)

	req.User.Username = strings.ToLower(req.User.Username)
	req.User.Email = strings.ToLower(req.User.Email)
	return db.Create(req.User, salt, pp)
}

func (s *Account) Read(ctx context.Context, req *account.ReadRequest, rsp *account.ReadResponse) error {
	user, err := db.Read(req.Id)
	if err != nil {
		return err
	}
	rsp.User = user
	return nil
}

func (s *Account) Update(ctx context.Context, req *account.UpdateRequest, rsp *account.UpdateResponse) error {
	req.User.Username = strings.ToLower(req.User.Username)
	req.User.Email = strings.ToLower(req.User.Email)
	return db.Update(req.User)
}

func (s *Account) Delete(ctx context.Context, req *account.DeleteRequest, rsp *account.DeleteResponse) error {
	return db.Delete(req.Id)
}

func (s *Account) Search(ctx context.Context, req *account.SearchRequest, rsp *account.SearchResponse) error {
	users, err := db.Search(req.Username, req.Email, req.Limit, req.Offset)
	if err != nil {
		return err
	}
	rsp.Users = users
	return nil
}

func (s *Account) UpdatePassword(ctx context.Context, req *account.UpdatePasswordRequest, rsp *account.UpdatePasswordResponse) error {
	usr, err := db.Read(req.UserId)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.user.updatepassword", err.Error())
	}

	salt, hashed, err := db.SaltAndPassword(usr.Username, usr.Email)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.user.updatepassword", err.Error())
	}

	hh, err := base64.StdEncoding.DecodeString(hashed)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.user.updatepassword", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword(hh, []byte(x+salt+req.OldPassword)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.updatepassword", err.Error())
	}

	salt = random(16)
	h, err := bcrypt.GenerateFromPassword([]byte(x+salt+req.NewPassword), 10)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.user.updatepassword", err.Error())
	}
	pp := base64.StdEncoding.EncodeToString(h)

	if err := db.UpdatePassword(req.UserId, salt, pp); err != nil {
		return errors.InternalServerError("go.micro.srv.user.updatepassword", err.Error())
	}
	return nil
}

func (s *Account) Login(ctx context.Context, req *account.LoginRequest, rsp *account.LoginResponse) error {
	username := strings.ToLower(req.Username)
	email := strings.ToLower(req.Email)

	salt, hashed, err := db.SaltAndPassword(username, email)
	if err != nil {
		return err
	}

	hh, err := base64.StdEncoding.DecodeString(hashed)
	if err != nil {
		return errors.InternalServerError("go.micro.srv.user.Login", err.Error())
	}

	if err := bcrypt.CompareHashAndPassword(hh, []byte(x+salt+req.Password)); err != nil {
		return errors.Unauthorized("go.micro.srv.user.login", err.Error())
	}
	// save session
	sess := &account.Session{
		Id:       random(128),
		Username: username,
		Created:  time.Now().Unix(),
		Expires:  time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	if err := db.CreateSession(sess); err != nil {
		return errors.InternalServerError("go.micro.srv.user.Login", err.Error())
	}
	rsp.Session = sess
	return nil
}

func (s *Account) Logout(ctx context.Context, req *account.LogoutRequest, rsp *account.LogoutResponse) error {
	return db.DeleteSession(req.SessionId)
}

func (s *Account) ReadSession(ctx context.Context, req *account.ReadSessionRequest, rsp *account.ReadSessionResponse) error {
	sess, err := db.ReadSession(req.SessionId)
	if err != nil {
		return err
	}
	rsp.Session = sess
	return nil
}

