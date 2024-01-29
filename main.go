package main

import "OJ_sandbox/router"

func main() {
	r := router.Router()
	r.Run(":8081")
}
