type Car struct {
	pos int
	speed int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	cars := make([]Car, n)
	for i := range position {
		cars[i] = Car{pos: position[i], speed: speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleetTimes := []float64{}

	for _, c := range cars {
		// Time for this car to reach target
		t := float64(target-c.pos) / float64(c.speed)

		// if this car can't catch up to fleet ahead, it starts a new fleet
		if len(fleetTimes) == 0 || t > fleetTimes[len(fleetTimes)-1] {
			fleetTimes = append(fleetTimes, t)
		} // else it joins existing fleet
	}

	return len(fleetTimes)
}
