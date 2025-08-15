package mbpb

import (
	"sync"
)

// DAG 结构体表示有向无环图
type DAG struct {
	Id       string             `json:"id"`
	nodes    map[string]*PBNode // 存储图中的所有节点，键为节点名称
	edges    map[string]*Edge   // 边列表
	Children []*PBNode          `json:"children"`
	Edges    []*Edge            `json:"edges"` // 边列表
	mu       sync.Mutex         // 互斥锁，保护 Nodes 的并发访问
}

// Node 结构体表示图的节点
type PBNode struct {
	Id       string     `json:"id"`     // 节点名称
	BaseId   string     `json:"baseId"` // 原始id
	Type     SourceType `json:"type"`   // 类型
	InDegree int        `json:"-"`      // 入度
	OutEdges []*Edge    `json:"-"`      // 出边列表
}

func NewGraph() *Graph {
	return &Graph{}
}

// NewDAG 创建一个新的 DAG 实例
func NewDAG() *DAG {
	return &DAG{
		nodes: make(map[string]*PBNode),
		edges: make(map[string]*Edge),
	}
}

func (d *DAG) Nodes() map[string]*PBNode {
	return d.nodes
}

// AddNode 添加一个节点到 DAG 中
func (d *DAG) AddNode(id string, baseId string, typ SourceType) *PBNode {
	d.mu.Lock()
	defer d.mu.Unlock()

	if _, exists := d.nodes[id]; !exists {

		node := &PBNode{
			Id:     id,
			BaseId: baseId,
			Type:   typ,
		}

		d.nodes[id] = node
		return node
	}
	return nil
}

// AddEdge 添加一条边（有向连接）到 DAG 中
func (d *DAG) AddEdge(source, target string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	sourceNode, sourceExists := d.nodes[source]
	targetNode, targetExists := d.nodes[target]

	if !sourceExists || !targetExists {
		return
	}

	id := HashString(source, target)

	// 己存在返回
	_, ok := d.edges[id]
	if ok {
		return
	}

	edge := &Edge{
		Id:     id,
		Source: source,
		Target: target,
	}

	d.edges[id] = edge
	targetNode.InDegree++
	d.Edges = append(d.Edges, edge)

	sourceNode.OutEdges = append(sourceNode.OutEdges, edge)

}

// TopologicalSort 对 DAG 进行拓扑排序
func (d *DAG) TopologicalSort() []string {
	var result []string
	queue := make([]*PBNode, 0)

	// 首先将所有入度为 0 的节点加入队列
	for _, node := range d.nodes {
		if node.InDegree == 0 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.Id)

		// 遍历当前节点的所有出边
		for _, edge := range current.OutEdges {
			targetNode := d.nodes[edge.Target]
			targetNode.InDegree--
			if targetNode.InDegree == 0 {
				queue = append(queue, targetNode)
			}
		}
	}

	// 如果排序后的结果数量等于节点数量，则表示没有环
	if len(result) == len(d.nodes) {
		return result
	}

	// 否则，存在环
	return nil
}

// HasCycle 检查 DAG 是否有环
func (d *DAG) HasCycle() bool {
	visited := make(map[interface{}]bool)
	recStack := make(map[interface{}]bool)

	for nodeName := range d.nodes {
		if d.hasCycleUtil(nodeName, visited, recStack) {
			return true
		}
	}
	return false
}

// 辅助函数，递归检查是否有环
func (d *DAG) hasCycleUtil(nodeName string, visited, recStack map[interface{}]bool) bool {
	if !visited[nodeName] {
		visited[nodeName] = true
		recStack[nodeName] = true

		node := d.nodes[nodeName]
		for _, edge := range node.OutEdges {
			childName := edge.Target
			if !visited[childName] && d.hasCycleUtil(childName, visited, recStack) {
				return true
			} else if recStack[childName] {
				return true
			}
		}
	}

	recStack[nodeName] = false
	return false
}
