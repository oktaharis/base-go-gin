package exception

import (
	"errors"

	"github.com/rs/zerolog/log"
)

var (
	ErrBearerTokenInvalid = errors.New("format token bearer tidak sesuai")
	ErrUserConflict       = errors.New("akun pengguna sudah terdaftar")
	ErrUserNotFound       = errors.New("akun tidak ditemukan")
	ErrUserLoginFailed    = errors.New("username/password salah")
)

func LogError(err error, message string) {
	log.Error().Stack().Err(err).Msg(message)
}
