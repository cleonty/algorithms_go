package quick

func Sort(data []int) {
	size := len(data)
	if size < 2 {
		return
	}
	p := &Partitioner{}
	index := p.partition(data)
	Sort(data[0:index])
	Sort(data[index+1:])
}

type Partitioner struct {
	data      []int
	wallIndex int
	pivot     int
}

func (p *Partitioner) partition(data []int) int {
	p.data = data
	size := len(data)
	p.pivot = data[size-1]
	p.wallIndex = -1
	for currentIndex := 0; currentIndex < size; currentIndex++ {
		if data[currentIndex] <= p.pivot {
			p.moveWallAndExchange(currentIndex)
		}
	}
	return p.wallIndex
}

func (p *Partitioner) moveWallAndExchange(currentIndex int) {
	p.moveWall()
	p.data[p.wallIndex], p.data[currentIndex] = p.data[currentIndex], p.data[p.wallIndex]
}

func (p *Partitioner) moveWall() {
	p.wallIndex++
}
