package media

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// レスポンスデータの定義
// MediaTypeに名前をつけます
var MessageMedia = MediaType("application/hello+json", func() {
	// 説明
	Description("example")
	// どのような値があるか（複数定義出来る）
	Attributes(func() {
		// idはInteger型
		Attribute("id", Integer, "id", func() {
			// 返すレスポンスの例
			Example(1)
		})

		// messageはString型
		Attribute("message", String, "message", func() {
			// 返すレスポンスの例
			Example("hello, world")
		})

		// レスポンスに必須な要素（基本は全て必須にした方が楽）
		Required("id")
		Required("message")
	})
	// 返すレスポンスのフォーマット（別記事で紹介予定）
	View("default", func() {
		Attribute("id")
		Attribute("message")
	})
})
