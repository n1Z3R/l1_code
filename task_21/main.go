package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

type Phone struct {
}

func (p *Phone) connectToCharge(c ChargingDevice) {
	fmt.Println("Connecting to charge")
	c.chargeFromSocket()
}

type ChargingDevice interface {
	chargeFromSocket()
}

type UsbcWithAdapter struct {
}

func (u *UsbcWithAdapter) chargeFromSocket() {
	fmt.Println("Charging from socket")
}

type Usbc struct {
}

func (u *Usbc) charge() {
	fmt.Println("Charging")
}

// Реализуется адаптер чтобы обьект соответствовал интерфейсу
type Adapter struct {
	usbc *Usbc
}

func (a *Adapter) chargeFromSocket() {
	fmt.Println("inputing in socket")
	a.usbc.charge()
}

func main() {
	ph := Phone{}

	usbcwa := &UsbcWithAdapter{}
	usbc := &Usbc{}

	// Оборачиваем usbc в адаптер
	adapter := &Adapter{usbc: usbc}

	ph.connectToCharge(usbcwa)
	ph.connectToCharge(adapter)
}
