package server

import (
	"42tokyo-road-to-dojo-go/pkg/database/dao"
	"42tokyo-road-to-dojo-go/pkg/http/middleware"
	"log"
	"net/http"

	"42tokyo-road-to-dojo-go/pkg/server/handler"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {

	d, err := dao.NewDao()
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}

	/* ===== URLマッピングを行う ===== */
	http.HandleFunc("/setting/get", get(handler.HandleSettingGet()))

	http.HandleFunc("/user/create", post(handler.HandleUserCreate(d)))
	http.HandleFunc("/user/get", get(middleware.Authenticate(d, handler.HandleUserGet())))
	http.HandleFunc("/user/update", post(middleware.Authenticate(d, handler.HandleUserUpdate(d))))

	http.HandleFunc("/game/finish", post(middleware.Authenticate(d, handler.HandleGameFinish())))

	http.HandleFunc("/gacha/draw", post(middleware.Authenticate(d, handler.HandleGachaDraw())))

	http.HandleFunc("/ranking/list", get(middleware.Authenticate(d, handler.HandleRankingList())))

	http.HandleFunc("/collection/list", get(middleware.Authenticate(d, handler.HandleCollectionList())))

	// TODO: 認証を行うmiddlewareを実装する
	// middlewareは pkg/http/middleware パッケージを利用する
	// http.HandleFunc("/user/get",
	//   get(middleware.Authenticate(handler.HandleUserGet())))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}
		apiFunc(writer, request)
	}
}
