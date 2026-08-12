package routing

import (
	"context"
	"fmt"
	"log"
)

func GetRoute() {
	client := New()

	route, err := client.Calculate(
		context.Background(),

		Point{
			Latitude:  -12.9714,
			Longitude: -38.5014,
		},

		Point{
			Latitude:  -12.9777,
			Longitude: -38.5016,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Distância: %.2f km\n", route.Distance/1000)
	fmt.Printf("Duração: %.2f minutos\n", route.Duration/60)
}
