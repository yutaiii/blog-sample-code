package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type post struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// echoインスタンス
	e := echo.New()

	// Middleware
	// httpリクエストの情報をログに表示
	e.Use(middleware.Logger())
	// パニックを回復し、スタックトレースを表示
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/", getSample)
	e.GET("/sample", getResponseSample)
	e.POST("/post/:id", postSample)

	// サーバーをスタートさせる
	// サーバーのポート番号は指定できる
	e.Logger.Fatal(e.Start(":8080"))
}

// Get API
func getSample(c echo.Context) error {
	return c.String(http.StatusOK, "Get!")
}

func getResponseSample(c echo.Context) error {
	p := &post{
		ID:   1,
		Name: "taro",
	}
	return c.JSON(http.StatusOK, p)
}

// Post API
func postSample(c echo.Context) error {
	p := new(post)
	if err := c.Bind(p); err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusInternalServerError, "Error!")
	}

	// URLパラメーターはBindで入らない
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("err %v", err.Error())
		return c.String(http.StatusInternalServerError, "Error!")
	}
	p.ID = id
	msg := fmt.Sprintf("id: %v, name %v", p.ID, p.Name)
	return c.String(http.StatusOK, msg)
}
