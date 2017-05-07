package algorithms

import (
	"math"
	"sync"
)

type UCB1 struct {
	counts []int64
	values []float64
	sync.RWMutex
}

func NewUCB1(nArms int, seed int64) *Softmax {
	return &Softmax{
		counts: make([]int64, nArms),
		values: make([]float64, nArms),
		seed:   seed,
	}
}

func (u *UCB1) SelectArm() int {
	u.RLock()
	defer u.RUnlock()

	for i := range u.counts {
		if u.counts[i] == 0 {
			return i
		}
	}

	ucbValues := make([]float64, len(u.values))
	t := sumInt64(u.counts)
	for i := range u.values {
		bonus := math.Sqrt((2 * math.Log(float64(t))) / float64(u.counts[i]))
		ucbValues[i] = u.values[i] + bonus
	}

	i, _ := maxFloat64(ucbValues)
	return i
}

func (u *UCB1) Update(chosen int, reward float64) {
	u.Lock()
	defer u.Unlock()

	u.counts[chosen]++

	n := float64(u.counts[chosen])
	org := u.values[chosen]
	u.values[chosen] = ((n-1)/n)*org + (1/n)*reward
}
