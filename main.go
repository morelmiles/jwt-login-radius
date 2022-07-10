package main

import (
	"github.com/morelmiles/jwt-login-radius/config"
	"github.com/morelmiles/jwt-login-radius/routes"
)

func main() {
	config.GetDB()
	routes.Routes()
}
