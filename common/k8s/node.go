package k8s

import "context"

type Node struct {
	*Watch
	nodeType       int
	watchNodeTypes []int
	ctx            context.Context
	ctxCancel      context.CancelFunc
}

func NewNode() *Node {

}
