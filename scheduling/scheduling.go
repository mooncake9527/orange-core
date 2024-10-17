package scheduling

import (
	"github.com/mooncake9527/npx/models"
	"github.com/mooncake9527/npx/scheduling/impl"
)

type SchedulingHandler interface {
	GetServiceNode(nodes []*models.ServiceNode, name string) *models.ServiceNode
}

func GetHandler(algorithm string) SchedulingHandler {
	algo := Algorithm(algorithm)
	switch algo {
	case AlgorithmRoundRobin:
		return impl.NewRoundRobinHandler()
	case AlgorithmRandom:
		return impl.NewRandomHandler()
	// case AlgorithmWeightedRandom:
	// 	return NewWeightedRandomHandler()
	// case AlgorithmIpHash:
	// 	return NewIpHashHandler()
	default:
		return impl.NewRoundRobinHandler()
	}
}

type Algorithm string

const (
	AlgorithmRandom     Algorithm = "random"
	AlgorithmRoundRobin Algorithm = "robin"
	// AlgorithmWeightedRandom Algorithm = "weight"
	// AlgorithmIpHash         Algorithm = "iphash"
)

func GetAlgorithm(name string) Algorithm {
	switch name {
	case "random":
		return AlgorithmRandom
	case "robin":
		return AlgorithmRoundRobin
	// case "weight":
	// 	return AlgorithmWeightedRandom
	// case "iphash":
	// 	return AlgorithmIpHash
	default:
		return AlgorithmRandom
	}
}
