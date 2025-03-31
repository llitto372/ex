package main

import (
	"errors"
	"fmt"
)

type Ticket struct {
	ID            int
	PassengerName string
	Destination   string
}

type BookingSystem struct {
	ticketmap map[int]Ticket
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		ticketmap: make(map[int]Ticket),
	}
}

func (b *BookingSystem) BookTicket(id int, name, distination string) error {
	struc, ok := b.ticketmap[id]
	if !ok {
		struc.ID, struc.PassengerName, struc.Destination = id, name, distination
		b.ticketmap[id] = struc
		return nil
	}
	return errors.New("ID уже существует")

}
func (b *BookingSystem) CancelTicket(id int) error {
	_, ok := b.ticketmap[id]
	if !ok {
		return errors.New("Неверный ID")
	} else {
		delete(b.ticketmap, id)
		return nil
	}

}
func (b *BookingSystem) GetTicket(id int) (Ticket, error) {
	struc, ok := b.ticketmap[id]
	if !ok {
		return Ticket{}, errors.New("ID не найден")
	} else {
		return struc, nil
	}
}

func main() {
	bs := NewBookingSystem()
	bs.BookTicket(1, "Иван", "Москва")
	fmt.Println(bs.GetTicket(1))    // {1 Иван Москва}, nil
	fmt.Println(bs.CancelTicket(1)) // nil
	fmt.Println(bs.GetTicket(1))    // Ошибка: билет не найден
}
