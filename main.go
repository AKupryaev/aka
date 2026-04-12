package main

import (
	"errors"
	"fmt"
)

// Структура билета
type Ticket struct {
	ID            int
	PassengerName string
	Destination   string
}

// Система бронирования
type BookingSystem struct {
	tickets map[int]Ticket
}

// Конструктор
func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		tickets: make(map[int]Ticket),
	}
}

// Бронирование билета
func (bs *BookingSystem) BookTicket(id int, name, destination string) error {
	if _, exists := bs.tickets[id]; exists {
		return errors.New("билет с таким ID уже существует")
	}

	bs.tickets[id] = Ticket{
		ID:            id,
		PassengerName: name,
		Destination:   destination,
	}

	return nil
}

// Отмена билета
func (bs *BookingSystem) CancelTicket(id int) error {
	if _, exists := bs.tickets[id]; !exists {
		return errors.New("билет не найден")
	}

	delete(bs.tickets, id)
	return nil
}

// Получение билета
func (bs *BookingSystem) GetTicket(id int) (Ticket, error) {
	ticket, exists := bs.tickets[id]
	if !exists {
		return Ticket{}, errors.New("билет не найден")
	}

	return ticket, nil
}

func main() {
	bs := NewBookingSystem()

	fmt.Println(bs.BookTicket(1, "Иван", "Москва")) // nil

	t, err := bs.GetTicket(1)
	fmt.Println(t, err) // {1 Иван Москва} <nil>

	fmt.Println(bs.CancelTicket(1)) // nil

	t, err = bs.GetTicket(1)
	fmt.Println(t, err) // {} билет не найден
}
