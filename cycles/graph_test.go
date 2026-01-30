package cycles

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle_EmptyGraph(t *testing.T) {
	g := Graph[int]{}
	assert.False(t, g.HasCycle(), "пустой граф не должен содержать циклов")
}

func TestHasCycle_SingleVertexNoEdge(t *testing.T) {
	g := Graph[int]{0: {}}
	assert.False(t, g.HasCycle(), "одна вершина без рёбер — циклов нет")
}

func TestHasCycle_SingleVertexSelfLoop(t *testing.T) {
	g := Graph[int]{0: {0}}
	assert.True(t, g.HasCycle(), "петля 0→0 должна давать цикл")
}

func TestHasCycle_SimpleCycle(t *testing.T) {
	g := Graph[int]{
		0: {1},
		1: {2},
		2: {0},
	}
	assert.True(t, g.HasCycle(), "граф 0→1→2→0 содержит цикл")
}

func TestHasCycle_Acyclic(t *testing.T) {
	g := Graph[int]{
		0: {1},
		1: {2},
	}
	assert.False(t, g.HasCycle(), "граф 0→1→2 ациклический")
}

func TestHasCycle_DisconnectedAcyclic(t *testing.T) {
	g := Graph[int]{
		0: {1},
		1: {},
		2: {3},
		3: {},
	}
	assert.False(t, g.HasCycle(), "две компоненты, обе ациклические")
}

func TestHasCycle_DisconnectedWithCycle(t *testing.T) {
	g := Graph[int]{
		0: {1},
		1: {},
		2: {3},
		3: {2},
	}
	assert.True(t, g.HasCycle(), "вторая компонента 2→3→2 содержит цикл")
}

func TestHasCycle_VertexNotInAdjacencyList(t *testing.T) {
	// Вершина 2 только в списке смежности (куда ведёт ребро), но не как ключ
	g := Graph[int]{
		0: {1},
		1: {2},
	}
	assert.False(t, g.HasCycle(), "цепочка 0→1→2 без обратных рёбер — циклов нет")
}

func TestHasCycle_MultipleCycles(t *testing.T) {
	g := Graph[int]{
		0: {1, 2},
		1: {2},
		2: {0},
	}
	assert.True(t, g.HasCycle(), "граф содержит цикл (например 0→2→0)")
}

func TestHasCycle_DAG(t *testing.T) {
	g := Graph[int]{
		0: {1, 2},
		1: {3},
		2: {3},
		3: {},
	}
	assert.False(t, g.HasCycle(), "ориентированный ациклический граф (алмаз) не должен содержать циклов")
}
