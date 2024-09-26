package impl

import (
	"github.com/mooncake9527/orange-core/models"
	"math/rand"
	"time"
)

type RandomHandler struct {
	r *rand.Rand
}

func NewRandomHandler() *RandomHandler {
	return &RandomHandler{
		r: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (rh *RandomHandler) GetServiceNode(nodes []*models.ServiceNode, name string) *models.ServiceNode {
	if len(nodes) == 0 {
		return nil
	}
	for i := 0; i < len(nodes); i++ {
		idx := rh.r.Intn(len(nodes))
		if nodes[idx].Enable() {
			return nodes[idx]
		}
	}
	return nil
}
