package main

import (
	"fmt"

	"github.maescudero/backpack-bcgow6-mario-escudero/internal/service"
)

func main() {
	var tickets []service.Ticket
	// Funcion para obtener tickets del archivo csv
	service.NewBookings(tickets)

	datosReserva := service.Ticket{
		Id:          123,
		Names:       "mario",
		Email:       "marioesc@ml.ml",
		Destination: "malanzan",
		Date:        "21/05/1232",
		Price:       2421,
	}

	tickets = append(tickets, datosReserva)
	nuevaReserva := service.NewBookings(tickets)

	nuevaReserva.Create(datosReserva)

	fmt.Println("se agrego ticket")

}
