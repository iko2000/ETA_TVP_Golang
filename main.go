package main

import (
"fmt"
"main/services"
"main/routes"
)


func main() {

	fmt.Println("Microservice beta");
	
	routes.Servers();
    services.FeedBacktaker();
}