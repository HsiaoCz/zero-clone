package main

import (
	"flag"
	"fmt"

	"github.com/HsiaoCz/zero-clone/product_service/internal/config"
	"github.com/HsiaoCz/zero-clone/product_service/internal/handler"
	"github.com/HsiaoCz/zero-clone/product_service/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/productservice-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
