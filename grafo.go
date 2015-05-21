//autor: Adriano Carniel Benin - 173464
package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type AdjList [][]int
type WList [][]int

//sort interface
func (w WList) Len() int {
	return len(w)
}

func (w WList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WList) Less(i, j int) bool {
	return w[i][2] < w[j][2]
}

//return index of str in slice, -1 if not found
func findSlice(slice []string, str string) (i int) {
	for i, v := range slice {
		if v == str {
			return i
		}
	}
	return -1
}

func convertInput(lines []string) [][]int {
	var vertices_list []int
	var graph [][]int

	for _, v := range lines {
		if len(v) > 1 {
			for _, v := range strings.Fields(v[1 : len(v)-1]) {
				i, _ := strconv.Atoi(v)
				vertices_list = append(vertices_list, i)
			}
			graph = append(graph, vertices_list)
			vertices_list = nil
		}
	}

	return graph
}

//read text file into a adjacency and weight arrays
func read_graph() (adjacency_list AdjList, weight_list WList) {
	content, err := ioutil.ReadFile("grafo.txt")

	if err != nil {
		fmt.Println("arquivo nao existe")
		return
	}

	lines := strings.Split(string(content), "\n")

	adjacency_list = convertInput(lines[0:findSlice(lines, "")])
	weight_list = convertInput(lines[findSlice(lines, ""):len(lines)])

	return adjacency_list, weight_list
}

func (a_list AdjList) Spanning() bool {
	if len(a_list) > 0 {
		dfs := a_list.DFS(a_list[0][0])
		for _, v := range dfs {
			if v == false {
				return false
			}
		}
	}

	return true
}

func DFSBool(a_list AdjList, v int, discovered []bool) []bool {
	adj_list := a_list[v][1:len(a_list[v])] //get all  connected vertices

	discovered[v] = true
	for _, vertex := range adj_list {
		if !discovered[vertex] {
			DFSBool(a_list, vertex, discovered)
		}
	}

	return discovered
}

func (a_list AdjList) DFS(v int) []bool {
	discovered := make([]bool, len(a_list))

	return DFSBool(a_list, v, discovered)

}

func (a_list AdjList) Connected(v1 int, v2 int) bool {
	v1_connected := a_list.DFS(v1)
	return v1_connected[v2]
}

//receives sorted weight list
func (a_list AdjList) Kruskal(weight_list WList) AdjList {
	var forest AdjList
	inserted_vertices := 0
	vertices_size := len(a_list)

	//separate graph into forests of 1 vertex
	for _, v := range a_list {
		forest = append(forest, append(make([]int, 0), v[0]))
	}

	for _, v := range weight_list {
		//if number of edges = vertices - 1 tree is spanning
		if inserted_vertices == (vertices_size - 1) {
			return forest
		}
		if !forest.Connected(v[0], v[1]) {
			forest[v[0]] = append(forest[v[0]], v[1])
			forest[v[1]] = append(forest[v[1]], v[0])
			inserted_vertices++
		}

	}

	return forest

}

func main() {
	adjacency_list, weight_list := read_graph()
	var spanning_tree AdjList

	sort.Sort(weight_list)
	spanning_tree = adjacency_list.Kruskal(weight_list)
	fmt.Println("arvore fornecida")
	fmt.Println(adjacency_list)
	fmt.Println("arvore geradora minima")
	fmt.Println(spanning_tree)

}
