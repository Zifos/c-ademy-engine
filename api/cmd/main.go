package main

import "c-ademy/api"

func main() {
	e := api.GetRouter()

	e.Logger.Fatal(e.Start(":1323"))

}
