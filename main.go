package main

import (
	"project-1/config"
	"project-1/routes"
)

func main(){
config.KonekDB()
e:=routes.New()
e.Start(":3000")

}

