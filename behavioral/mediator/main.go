package main

func main() {
	stationManager := newStationManager()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}

	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
