//autor: Adriano Carniel Benin - 173464
//https://github.com/adrianob/kruskal
//executável fornecido para Mac OS X, rodar 'go build grafo.go' para compilar em outro sistema
//ex de saida:
//lista de adjacencia
//[[0 1] [1 0 3] [2 5] [3 4 1] [4 3 5] [5 2 4]]
//array de array de inteiros onde o primeiro numero de cada array indica o vértice e os números restantes os nodos adjacentes
//lista de pesos
//[[2 5 1] [3 4 1] [0 1 2] [1 3 2] [4 5 2]]
//lista de pesos onde os primeiros 2 inteiros indicam os vertices conectados e o terceiro número indica o peso da aresta

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type AdjList [][]int
type WList [][]int

type Graph struct {
	adjacency_list AdjList
	weight_list    WList
}

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
func findSlice(re *regexp.Regexp, slice []string) (i int) {
	for i, v := range slice {
		if (re.FindStringIndex(v) != nil) && i != 0 {
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

//read text file into a graph struct
func read_graph() (g Graph) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Arquivo de dados: ")
	text, _, _ := reader.ReadLine()
	content, err := ioutil.ReadFile(string(text))

	if err != nil {
		fmt.Println("arquivo nao existe")
		return
	}

	lines := strings.Split(string(content), "\n")

	//search for beggining of weight list
	re := regexp.MustCompile("(?m)^\\($")
	adjacency_list := convertInput(lines[0:findSlice(re, lines)])
	weight_list := convertInput(lines[findSlice(re, lines):len(lines)])

	g = Graph{adjacency_list, weight_list}
	return g
}

//check for connection with DFS
func (a_list AdjList) ConnectedBool(v int, v2 int, discovered []bool) bool {
	adj_list := a_list[v][1:len(a_list[v])] //get all  connected vertices

	if v == v2 {
		return true
	}
	discovered[v] = true
	for _, vertex := range adj_list {
		if !discovered[vertex] {
			a_list.ConnectedBool(vertex, v2, discovered)
		}
	}

	return false
}

func (a_list AdjList) Connected(v int, v2 int) bool {
	discovered := make([]bool, len(a_list))

	return a_list.ConnectedBool(v, v2, discovered)
}

func (g Graph) Kruskal() (spanning_tree Graph, total_weight int) {
	var forest AdjList
	var weight_list WList
	total_weight = 0
	inserted_vertices := 0
	vertices_size := len(g.adjacency_list)

	sort.Sort(g.weight_list)
	//separate graph into forests of 1 vertex
	for _, v := range g.adjacency_list {
		forest = append(forest, append(make([]int, 0), v[0]))
	}

	for _, v := range g.weight_list {
		//if number of edges = vertices - 1, tree is spanning
		if inserted_vertices == (vertices_size - 1) {
			return Graph{forest, weight_list}, total_weight
		}
		if !forest.Connected(v[0], v[1]) {
			total_weight += v[2]
			weight_list = append(weight_list, v)
			//insert both ends of vertice on adjacency list
			forest[v[0]] = append(forest[v[0]], v[1])
			forest[v[1]] = append(forest[v[1]], v[0])
			inserted_vertices++
		}

	}

	return Graph{forest, weight_list}, total_weight

}

func main() {
	graph := read_graph()
	var spanning_tree Graph
	var total_weight int

	spanning_tree, total_weight = graph.Kruskal()
	fmt.Println("arvore fornecida")
	fmt.Println(graph.adjacency_list)
	fmt.Println("arvore geradora minima - lista de adjacência")
	fmt.Println(spanning_tree.adjacency_list)
	fmt.Println("arvore geradora minima - lista de pesos")
	fmt.Println(spanning_tree.weight_list)
	fmt.Println("peso total")
	fmt.Println(total_weight)

}
