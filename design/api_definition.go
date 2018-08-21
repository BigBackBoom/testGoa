package design

import (
	. "github.com/goadesign/goa/design/apidsl"
	_ "github.com/testGoa/design/resource"
)

var _ = API("goa simple sample", func() {
	// APIのタイトル
	Title("goa-simple-sample")
	// APIの説明
	Description("goaのサンプルです")
	// 作成者へのコンタクト情報
	Contact(func() {
		Name("bigbackboom")
		Email("kikuchik@istyle.co.jp")
	})

	// ホストの設定
	Host("localhost:8080")
	// 対応しているプロトコル定義、httpかhttpsまたはその両方
	Scheme("http")
	// 全てのエンドポイントのベースパス
	// /usersというエンドポイントがあったら、/api/v1/usersとなる
	BasePath("/api/v1")

	//　CORSポリシーの定義
	Origin("http://localhost:8080/swagger", func() {
		//クライアントに公開された1つ以上のヘッダー
		Expose("X-Time")
		// 1つまたは複数の許可されたHTTPメソッド
		Methods("GET", "POST", "PUT", "DELETE")
		//プリフライト要求応答をキャッシュする時間
		MaxAge(600)
		// Access-Control-Allow-Credentialsヘッダーを設定する
		Credentials()
	})
})
