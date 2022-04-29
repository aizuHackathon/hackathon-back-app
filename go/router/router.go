package router

import (
	"app/handler"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IV1 interface {
		withNone(e *echo.Echo)
		// 緒方の編集箇所
		//以下の関数をデモ的に追加したので以下にも記述した。
		GetMongon(e *echo.Echo)
	}

	V1 struct {
		userHandler handler.IUser
	}
)

func Init(e *echo.Echo) {
	var r IV1 = &V1{
		userHandler: handler.NewUser(),
	}

	r.withNone(e)

	// 簡易版
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	// 緒方の編集箇所
	//　デモ用でお手本を基に作成し、自分の手で動かしてみた。
	e.GET("/mongon", func(c echo.Context) error{
		return c.String(http.StatusOK, "mongon")
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func (r V1) withNone(e *echo.Echo) {
	e.GET("/users", r.userHandler.Index)
}

// 緒方の編集箇所
// 実際にdbにアクセスしてカラムのデータを取得したい
func (r V1) GetMongon(e *echo.Echo) {
	e.GET("/GetMongon", r.userHandler.Index)
}