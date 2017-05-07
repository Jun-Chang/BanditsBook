package algorithms

import (
	"math/rand"
	"sync"
)

type EpsilonGreedy struct {
	counts  []int64
	values  []float64
	epsilon float64
	seed    int64
	sync.RWMutex
}

func NewEpsilonGreedy(nArms int, epsilon float64, seed int64) *EpsilonGreedy {
	return &EpsilonGreedy{
		counts:  make([]int64, nArms),
		values:  make([]float64, nArms),
		epsilon: epsilon,
		seed:    seed,
	}
}

func (e *EpsilonGreedy) SelectArm() int {
	e.RLock()
	defer e.RUnlock()

	rnd := rand.New(rand.NewSource(e.seed))

	if rnd.Float64() > e.epsilon {
		i, _ := maxFloat64(e.values)
		return i
	}

	return rnd.Intn(len(e.values))
}

func (e *EpsilonGreedy) Update(chosen int, reward float64) {
	e.Lock()
	defer e.Unlock()

	e.counts[chosen]++

	n := float64(e.counts[chosen])
	org := e.values[chosen]
	e.values[chosen] = ((n-1)/n)*org + (1/n)*reward
}
