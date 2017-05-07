package algorithms

import (
	"math"
	"math/rand"
	"sync"
)

type Softmax struct {
	counts      []int64
	values      []float64
	temperature float64
	seed        int64
	sync.RWMutex
}

func NewSoftmax(nArms int, temperature float64, seed int64) *Softmax {
	return &Softmax{
		counts:      make([]int64, nArms),
		values:      make([]float64, nArms),
		temperature: temperature,
		seed:        seed,
	}
}

func (s *Softmax) SelectArm() int {
	s.RLock()
	defer s.RUnlock()

	var z float64
	for _, v := range s.values {
		z += math.Exp(v / s.temperature)
	}

	probs := make([]float64, len(s.values))
	for i, v := range s.values {
		probs[i] = math.Exp(v/s.temperature) / z
	}

	return categoricalDraw(probs, s.seed)
}

func (s *Softmax) Update(chosen int, reward float64) {
	s.Lock()
	defer s.Unlock()

	s.counts[chosen]++

	n := float64(s.counts[chosen])
	org := s.values[chosen]
	s.values[chosen] = ((n-1)/n)*org + (1/n)*reward
}

func categoricalDraw(probs []float64, seed int64) int {
	z := rand.New(rand.NewSource(seed)).Float64()

	var cumProb float64
	for i, p := range probs {
		cumProb += p
		if cumProb > z {
			return i
		}
	}
	return len(probs) - 1
}
