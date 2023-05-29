package main

import "pancakaki/utils/server"

func main() {
	if err := server.Run(); err != nil {
		panic(err)
	}
}
