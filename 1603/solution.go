package solution

type ParkingSystem struct {
	bigLimit    int
	mediumLimit int
	smallLimit  int
	bigCount    int
	mediumCount int
	smallCount  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		bigLimit:    big,
		mediumLimit: medium,
		smallLimit:  small,
		bigCount:    0,
		mediumCount: 0,
		smallCount:  0,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.bigLimit > this.bigCount {
		this.bigCount++
		return true
	} else if carType == 2 && this.mediumLimit > this.mediumCount {
		this.mediumCount++
		return true
	} else if carType == 3 && this.smallLimit > this.smallCount {
		this.smallCount++
		return true
	} else {
		return false
	}
}
