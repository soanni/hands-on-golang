package main

import "fmt"

type Trip struct {
	reservations []Reservation
}

func (t *Trip) CalculateCancellationFee() float64 {
	total := 0.0

	for _, r := range(t.reservations) {
		total += r.CalculateCancellationFee()
	}

	return total
}

func (t *Trip) AddReservation (r Reservation) {
	t.reservations = append(t.reservations, r)
}

type Reservation interface {
	GetReservationDate() string
	CalculateCancellationFee() float64
}

type HotelReservation interface {
	Reservation
	ChangeType()
}

type FlightReservation interface {
	Reservation
	AddExtraLuggageAllowance(peices int)
}

type HotelReservationImpl struct {
	reservationDate string
}

func (r HotelReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r HotelReservationImpl) CalculateCancellationFee() float64 {
	return 1.0
}

type FlightReservationImpl struct {
	reservationDate string
	luggageAllowed int
}

func (r FlightReservationImpl) GetReservationDate() string {
	return r.reservationDate
}

func (r FlightReservationImpl) CalculateCancellationFee() float64 {
	return 1.0
}

func NewReservation(vertical, reservationDate string) Reservation {
	switch(vertical) {
		case "flight":
			return FlightReservationImpl{reservationDate: reservationDate,}
		case "hotel":
			return HotelReservationImpl{reservationDate: reservationDate,}
		default:
			return nil
	}
}

type ReservationBuilder interface {
	Vertical(string) ReservationBuilder
	ReservationDate(string) ReservationBuilder
	Build() Reservation
}

type reservationBuilder struct {
	vertical string
	rdate string
}

func (r *reservationBuilder) Vertical(v string) ReservationBuilder {
	r.vertical = v
	return r
}

func (r *reservationBuilder) ReservationDate(date string) ReservationBuilder {
	r.rdate = date
	return r
}

func (r *reservationBuilder) Build() Reservation {
	var builtReservation Reservation

	switch r.vertical {
		case "flight":
			builtReservation = FlightReservationImpl{reservationDate: r.rdate}
		case "hotel":
			builtReservation = HotelReservationImpl{reservationDate: r.rdate}
	}

	return builtReservation
}

func NewReservationBuilder() ReservationBuilder {
	return &reservationBuilder{}
}

func main() {
	hotelReservation := NewReservation("hotel", "20190204")
	fmt.Println(hotelReservation)
	
	flightBuilder := NewReservationBuilder()
	flightBuilder.Vertical("flight")
	flightBuilder.ReservationDate("20190205")
	flightReservation := flightBuilder.Build()
	fmt.Println(flightReservation)
}
