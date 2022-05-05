package middleware

import (
	"42tokyo-road-to-dojo-go/pkg/database/dao"
	"42tokyo-road-to-dojo-go/pkg/database/object"
	"context"
	"fmt"
	"net/http"
)

var (
	contextKey         = new(struct{})
	ErrUnAuthenticated = fmt.Errorf("unauthenticated")
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(d dao.Dao, nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//ctx := r.Context()
		//if ctx == nil {
		//	ctx = context.Background()
		//}

		// TODO: implement here
		token := r.Header.Get("x-token")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user, err := d.NewUser().FindByToken(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		nextFunc(w, r.WithContext(context.WithValue(r.Context(), contextKey, user)))
	}
}

func UserOf(r *http.Request) (*object.User, error) {
	v, ok := r.Context().Value(contextKey).(*object.User)
	if ok {
		return v, nil
	} else {
		return v, ErrUnAuthenticated
	}
}
