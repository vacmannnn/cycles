package cycles

const (
	NotVisited = iota
	InStack
	Visited
)

// Graph представляет ориентированный граф в виде списка смежности.
// T — тип вершины (должен быть comparable для использования как ключ map).
// Ключ — вершина, значение — слайс вершин, в которые ведут рёбра из этой вершины.
type Graph[T comparable] map[T][]T

// HasCycle определяет, есть ли в графе хотя бы один цикл.
// Используется обход в глубину с раскраской вершин (белый/серый/чёрный).
func (g Graph[T]) HasCycle() bool {
	if len(g) == 0 {
		return false
	}

	vertices := make(map[T]struct{})
	for v, neighbors := range g {
		vertices[v] = struct{}{}
		for _, u := range neighbors {
			vertices[u] = struct{}{}
		}
	}

	color := make(map[T]int)
	for v := range vertices {
		color[v] = NotVisited
	}

	var visit func(T) bool
	visit = func(v T) bool {
		color[v] = InStack
		for _, u := range g[v] {
			switch color[u] {
			case NotVisited:
				if visit(u) {
					return true
				}
			case InStack:
				return true
			case Visited:
			}
		}
		color[v] = Visited
		return false
	}

	for v := range vertices {
		if color[v] == NotVisited && visit(v) {
			return true
		}
	}
	return false
}
