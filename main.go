package main

import (
	"github.com/labstack/echo/v4"
)

// 定义接口 interface
type IExampleServiceEchoHandler interface {
	Hello(echo.Context) error
	Say(echo.Context) error
}

// 定义接口
type ExampleServiceEchoHandler struct {
	handler ExampleServiceHandler
}

func NewExampleServiceEchoHandler(handler ExampleServiceHandler) IExampleServiceEchoHandler {
	return &ExampleServiceEchoHandler{
		handler: handler,
	}
}

func (h ExampleServiceEchoHandler) Hello(c echo.Context) error {
	// return h.handler.hello(c)
	return nil
}

func (h ExampleServiceEchoHandler) Say(c echo.Context) error {
	// return h.handler.say(c)
	return nil
}

func ExampleServiceRegisterEchoRoutes(e *echo.Echo) {
	// h := NewExampleServiceEchoHandler(NewExampleServiceHandler())
	// e.GET("/hello", h.Hello)
	// e.GET("/say", h.Say)
}

//

type (
	ExampleServiceHandler struct{}
)

// func (e *ExampleServiceHandler) hello(c echo.Context) error {
// 	return c.JSON(http.StatusOK, "hello jack")
// }

// func (e *ExampleServiceHandler) say(c echo.Context) error {
// 	return c.JSON(http.StatusOK, "say it")
// }

func main() {
	e := echo.New()
	ExampleServiceRegisterEchoRoutes(e)
	e.GET("/abc", func(c echo.Context) error {
		return nil
	})
	if err := e.Start(":1111"); err != nil {
		panic(err)
	}
}
