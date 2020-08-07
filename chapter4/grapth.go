package main

graph = make(map[string]map[string]bool)

func addEdge(from, to string){
	edge := graph[from]
	if edge == nil {
		ed := make(map[string]bool)
		graph[from] = ed 
	}
	ed[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
