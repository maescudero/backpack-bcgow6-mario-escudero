package file

import (
	"fmt"
	"os"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	path string
}

func (f *File) Read() ([]service.Ticket, error) {
	return nil, nil
}

func (f *File) Write(ser service.Ticket) error {

	//f := file.File{Path: "/Users/MAESCUDERO/github/hackaton-go-bases-main/tickets.csv"}
	crearTicket := fmt.Sprintf("%v %s %s %s %s %v", ser.Id, ser.Names, ser.Email, ser.Destination, ser.Date, ser.Price)
	creacion, err := os.OpenFile(f.path, os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	_, e := creacion.Write([]byte(crearTicket))
	if e != nil {
		fmt.Println(err)
	}
	creacion.Close()
	return nil

}
