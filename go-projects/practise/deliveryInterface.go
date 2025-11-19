package main

type Delivery interface {
	Deliver()
}
type BikeDelivery struct {
	BikeNumber string
}
type DroneDelivery struct {
	DroneID string
}
type TruckDelivery struct {
	TruckNumber string
}

func (d BikeDelivery) Deliver() {
	println("Delivering by Bike:", d.BikeNumber)
}
func (d DroneDelivery) Deliver() {
	println("Delivering by Drone:", d.DroneID)
}
func (d TruckDelivery) Deliver() {
	println("Delivering by Truck:", d.TruckNumber)
}
func main() {
	var d1 Delivery = BikeDelivery{BikeNumber: "B123"}
	var d2 Delivery = DroneDelivery{DroneID: "D456"}
	var d3 Delivery = TruckDelivery{TruckNumber: "T789"}
	d1.Deliver()
	d2.Deliver()
	d3.Deliver()
}
