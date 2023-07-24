package main

import (
	"os"

	"git.isfk.cn/isfk/protoc-gen-echo/example"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

type myHandler struct {
	example.ExampleService_EchoClientHandlerImpl
	log *slog.Logger
}

func NewMyHandler(log *slog.Logger) *myHandler {
	return &myHandler{log: log}
}

func main() {
	e := echo.New()

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	handler := example.NewExampleService_EchoServerHandler(NewMyHandler(log))

	e.GET("/hello", handler.Hello)
	e.GET("/say", handler.Say)
	e.Start(":1111")
}

func (h myHandler) Hello(args *example.HelloRequest) (*example.HelloResponse, error) {
	h.log.Info("打印参数", slog.Any("args", args))
	return &example.HelloResponse{Msg: args.Name}, nil
}

func (h myHandler) Say(args *example.SayRequest) (*example.SayResponse, error) {
	h.log.Info("打印参数", slog.Any("args", args))
	return &example.SayResponse{Msg: args.Name}, nil
}
