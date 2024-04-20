package generator

import (
	"github.com/bwmarrin/snowflake"
)

var (
	DefaultNumberNode int64 = 1
)

type SnowflakeNode struct {
	Snode *snowflake.Node
}

func New(numberNode int64) (*SnowflakeNode, error) {
	DefaultNumberNode = numberNode
	node, err := snowflake.NewNode(DefaultNumberNode)
	if err != nil {
		return &SnowflakeNode{}, err
	}
	return &SnowflakeNode{Snode: node}, nil
}
