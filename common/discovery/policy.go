package discovery

type Policy int

const (
	_          Policy = iota // 0
	Ordered                  // 1
	Random                   // 2
	RoundRobin               // 3
)
