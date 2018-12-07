package auth

import "github.com/dgrijalva/jwt-go"

type Auth struct {
	signingKey []byte
}

func New() *Auth {
	auth := new(Auth)

	return auth
}

func (a *Auth) Middleware() *JWTMiddleware {
	return middleware(MiddlewareOptions{
		Extractor: FromFirst(
			FromAuthHeader,
			FromParameter("key"),
			FromFormValue("key"),
		),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return a.signingKey, nil
		},
		SigningMethod:       jwt.SigningMethodHS256,
		CredentialsOptional: true,
	})
}
