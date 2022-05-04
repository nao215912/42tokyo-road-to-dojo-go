package middleware

import (
	"context"
	"fmt"
	"net/http"
)

var (
	contextKey         = new(struct{})
	ErrUnAuthenticated = fmt.Errorf("unauthenticated")
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//ctx := r.Context()
		//if ctx == nil {
		//	ctx = context.Background()
		//}

		// TODO: implement here
		token := r.Header.Get("x-token")
		nextFunc(w, r.WithContext(context.WithValue(r.Context(), contextKey, token)))
	}
}

func UserOf(r *http.Request) (string, error) {
	v, ok := r.Context().Value(contextKey).(string)
	if ok {
		return v, nil
	} else {
		return "", ErrUnAuthenticated
	}
}
