package main

import "gin_frame/routers"

func main() {
	router := routers.SetupRouter()

	router.Run()
}
