package resource

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/testGoa/design/media"
)

// /actionsの定義をする
var _ = Resource("action", func() {
	// actionsリソースのベースパス
	BasePath("/action")
	/*
	   リソースに対してどういった操作を行うか定義する
	   add リソースを追加する
	   list    リソースをリストで取得する
	   delete  リソースを削除する
	   上記のような感じで定義すればおｋです。
	*/
	Action("ping", func() {
		// アクションの説明
		Description("サーバーとの導通確認")
		Routing(
			// エンドポイント -> GET http://localhost/api/v1/actions/pingになる
			GET("/ping"),
		)
		// 返したいレスポンス
		// 200 OK + MessageTypeで定義しているMediaType
		Response(OK, MessageMedia)
		// 400 BadRequest + ErrorMediaというデフォルトで容易されているMediaType
		// 足りないパラメーターなどがあれば自動的に返される
		Response(BadRequest, ErrorMedia)
	})
})

// Swaggerをローカルで実行するめの定義
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})
