package impl

import (
	"github.com/mooncake9527/npx/models"
)

type RoundRobinHandler struct {
	cur map[string]int
}

func NewRoundRobinHandler() *RoundRobinHandler {
	return &RoundRobinHandler{
		cur: make(map[string]int),
	}
}

func (r *RoundRobinHandler) GetServiceNode(nodes []*models.ServiceNode, name string) *models.ServiceNode {
	if len(nodes) == 0 {
		return nil
	}
	for i := 0; i < len(nodes); i++ {
		if idx, ok := r.cur[name]; ok {
			useIdx := idx % len(nodes)
			r.cur[name] = useIdx + 1
			if nodes[useIdx].Enable() {
				return nodes[useIdx]
			}
		} else {
			r.cur[name] = 0
			if nodes[0].Enable() {
				return nodes[0]
			}
		}
	}
	return nil
}
