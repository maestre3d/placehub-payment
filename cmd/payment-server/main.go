package main

import "github.com/maestre3d/placehub-payment/cmd/payment-server/bootstrap"

func main() {
	app := bootstrap.NewApp()
	app.Run()
}
