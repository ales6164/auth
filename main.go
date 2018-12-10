package auth

import "github.com/dgrijalva/jwt-go"

type Auth struct {
	*Options
}

type Options struct {
	SigningKey          []byte
	Extractors          []TokenExtractor
	CredentialsOptional bool
	SigningMethod       jwt.SigningMethod
}

func New(opt *Options) *Auth {
	auth := &Auth{Options: opt}
	return auth
}

func (a *Auth) Middleware() *JWTMiddleware {
	return middleware(MiddlewareOptions{
		Extractor: FromFirst(a.Extractors...),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return a.SigningKey, nil
		},
		SigningMethod:       a.SigningMethod,
		CredentialsOptional: a.CredentialsOptional,
	})
}

// when project selected - switch namespace
// when on some project endpoint check for project namespace??

// group/user entity access
// kind datastore is protected by general scope rules - anyone with kind.ReadWrite scope can read/write to that kind
// collection is a datastore data container separated with a namespace
// collections are dynamic, it's rules are stored with auth
// collections can be created dynamically - for a kind entry; creator get's collection.FullControl scope
// collection is a wrapper and is handled with middleware - sets context namespace if user has access

// kinds inside collections can be created, edited, deleted
// keys contain namespace (collection id) so API must check user has scope to edit kind within the collection
// creating entries inside collection?
// how to identify collection access?

// /{collection}/{kind}/... ?
// this would work for collections as is user data, projects, groups
