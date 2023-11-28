package main

import routes "hexa-go/pkg"

func main() {
	app := routes.NewRoute().InitializeRouter()
	app.Listen(":3000")
}
