package main

import (
	"fmt"
	"sort"
	"time"
)

// 节点的结构体
type Node struct {
	element      []int
	left_length  int
	right_length int
	left         int
	right        int
}

// 边的结构体
type Edge struct {
	destination int
	weight      int
}

// 邻接表的结构体
type Graph struct {
	adjacencyList map[int][]Edge
}

func main() {

	fmt.Println("20250101")

	// 定义图结构
	graph := Graph{
		adjacencyList: make(map[int][]Edge),
	}

	//定义边集
	var Edges []int

	// 添加边
	graph.addEdge(1, 7, 2, &Edges)
	graph.addEdge(2, 8, 3, &Edges)
	graph.addEdge(3, 9, 4, &Edges)
	graph.addEdge(4, 6, 5, &Edges)
	graph.addEdge(5, 3, 6, &Edges)
	graph.addEdge(6, 4, 1, &Edges)
	graph.addEdge(1, 1, 4, &Edges)
	graph.addEdge(2, 5, 5, &Edges)
	graph.addEdge(6, 2, 3, &Edges)

	// 打印邻接表
	fmt.Println("***************解析模块***************")
	graph.printGraph()

	// 获取节点数量
	node_num := len(graph.adjacencyList)
	fmt.Printf("\n节点数量:%d\n", node_num)

	// 获取边的数量
	edge_num := 0
	for i := 1; i <= node_num; i++ {
		edge_num += len(graph.adjacencyList[i])
	}
	edge_num /= 2
	fmt.Printf("边的数量:%d\n\n", edge_num)

	// 创建节点集合,也就是制作数据库
	var nodes []Node

	// 遍历每个节点
	for i := 1; i <= node_num; i++ {
		// 遍历每个节点的左右边的所有可能性
		for l := 0; l < len(graph.adjacencyList[i])-1; l++ {
			for r := 1; r < len(graph.adjacencyList[i]); r++ {
				if l < r {
					for k := 0; k < 1; k++ {
						ele_temp := []int{i}
						nodes = append(nodes, Node{
							element:      ele_temp,
							left_length:  graph.adjacencyList[i][l].weight,
							right_length: graph.adjacencyList[i][r].weight,
							left:         graph.adjacencyList[i][l].destination,
							right:        graph.adjacencyList[i][r].destination,
						})
					}
				}
			}
		}
	}

	// 打印数据体
	// fmt.Println("数据体")
	// fmt.Println("编号|内编号|左边长|右边长|左节点|右节点")
	// for i, n := range nodes {
	// 	fmt.Printf("%d\t%d\t\t%d\t\t%d\t%d\t\t%d\n", i, n.element, n.left_length, n.right_length, n.left, n.right)
	// }
	fmt.Println("***************映射模块***************")
	fmt.Println("数据库")
	fmt.Println()
	// fmt.Println("编号|内编号|左边长|右边长|左节点|右节点")
	for _, n := range nodes {
		fmt.Printf("x_{%d%d%d},", n.left, n.element[0], n.right)
	}
	fmt.Println()

	// //生成探针库
	// for i := 0; i < len(nodes); i++ {
	// 	for j := i + 1; j < len(nodes); j++ {
	// 		if nodes[i].left == nodes[j].element[0] && nodes[j].right == nodes[i].element[0] {
	// 			fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].left, "},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].right, "}},")
	// 		} else if nodes[i].right == nodes[j].element[0] && nodes[j].right == nodes[i].element[0] {
	// 			fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].right, "},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].right, "}},")
	// 		} else if nodes[i].left == nodes[j].element[0] && nodes[j].left == nodes[i].element[0] {
	// 			fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].left, "},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].left, "}},")
	// 		} else if nodes[i].right == nodes[j].element[0] && nodes[j].left == nodes[i].element[0] {
	// 			fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].right, "},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].left, "}},")
	// 		}
	// 	}
	// }

	//生成探针库
	fmt.Println("探针库")
	fmt.Println()
	for i := 0; i < len(nodes); i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].left == nodes[j].element[0] && nodes[j].right == nodes[i].element[0] {
				//fmt.Println()
				//fmt.Println("代表连接", nodes[i].element[0], nodes[j].element[0], "的边")
				//fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].left,
				//	"},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].right, "}},")

				fmt.Println()
				fmt.Println("x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].left,
					"},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].right, "},")
				
			} else if nodes[i].right == nodes[j].element[0] && nodes[j].right == nodes[i].element[0] {
				//fmt.Println()
				//fmt.Println("代表连接", nodes[i].element[0], nodes[j].element[0], "的边")
				//fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].right,
				//	"},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].right, "}},")

				fmt.Println()
				fmt.Println("代表连接", nodes[i].element[0], nodes[j].element[0], "的边")
				fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].right,
					"},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].right, "}},")
			} else if nodes[i].left == nodes[j].element[0] && nodes[j].left == nodes[i].element[0] {
				fmt.Println()
				fmt.Println("代表连接", nodes[i].element[0], nodes[j].element[0], "的边")
				fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].left,
					"},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].left, "}},")
			} else if nodes[i].right == nodes[j].element[0] && nodes[j].left == nodes[i].element[0] {
				fmt.Println()
				fmt.Println("代表连接", nodes[i].element[0], nodes[j].element[0], "的边")
				fmt.Println("\\overline{x_{", nodes[i].left, nodes[i].element[0], nodes[i].right, "}^{", nodes[i].right,
					"},x_{", nodes[j].left, nodes[j].element[0], nodes[j].right, "}^{", nodes[j].left, "}},")
			}
		}
	}

	// 处理探针
	sort.Ints(Edges)

	// 打印边集
	fmt.Println("探针")
	for _, b := range Edges {
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	fmt.Println()

	// 定义操作池子
	pool := make([][]Node, 10)

	// 准备工作,把初始数据放进池子
	for _, node := range nodes {
		pool[node.left_length] = append(pool[node.left_length], node)
		pool[node.right_length] = append(pool[node.right_length], node)
	}
	// 输出池子
	// fmt.Println("操作池")
	// for i, p := range pool {
	// 	fmt.Printf("第%d个小池子\n", i)
	// 	for _, n := range p {
	// 		fmt.Println(n)
	// 	}
	// }

	//开始计时
	start := time.Now()

	//新产生的数据体
	var new_data []Node
	//取出一个探针
	for i := 0; i < len(Edges); i++ {

		// 当前的探针
		curr_p := Edges[i]
		fmt.Printf("\n\n\n现在把探针%d放入池子\n", Edges[i])

		// 探针和单个数据体的结合
		// 当首次可能产生环时，进行补环操作
		if i >= (node_num - 1) {
			for j := 0; j < len(pool[curr_p]); j++ {
				if is_close_loop(pool[curr_p][j], curr_p, node_num) {
					fmt.Println("获得环", pool[curr_p][j])
					fmt.Println("这是进行探针", curr_p, "操作时出现的")
					goto end
				}
			}
		}

		//在计算平台内的双重循环
		for l := 0; l < len(pool[curr_p])-1; l++ {
			for r := l + 1; r < len(pool[curr_p]); r++ {

				//左左连接
				if is_left_to_left(pool[curr_p][l], pool[curr_p][r], curr_p) {
					if !hasDuplicate(pool[curr_p][l].element, pool[curr_p][r].element) {
						node_temp := newnode_left_to_left(pool[curr_p][l], pool[curr_p][r])
						if detection(node_temp, node_num) {
							//fmt.Println(pool[curr_p][j], pool[curr_p][k], "左与左连接产生一个新的结构", node_temp)
							pool[node_temp.left_length] = append(pool[node_temp.left_length], node_temp)
							pool[node_temp.right_length] = append(pool[node_temp.right_length], node_temp)
							if !is_in_nodes(node_temp, new_data) {
								new_data = append(new_data, node_temp)
							}
						}
					} else {
						//fmt.Println(pool[curr_p][j], pool[curr_p][k], "有重合元素,不予连接", pool[curr_p][j].element, pool[curr_p][k].element)
					}
				} else if is_right_to_left(pool[curr_p][l], pool[curr_p][r], curr_p) {
					//fmt.Print("右和左相等")
					//fmt.Println(pool[curr_p][j], pool[curr_p][k])
					if !hasDuplicate(pool[curr_p][l].element, pool[curr_p][r].element) {
						node_temp := newnode_right_to_left(pool[curr_p][l], pool[curr_p][r])
						if detection(node_temp, node_num) {
							//fmt.Println(pool[curr_p][j], pool[curr_p][k], "左与右连接产生一个新的结构", node_temp)
							pool[node_temp.left_length] = append(pool[node_temp.left_length], node_temp)
							pool[node_temp.right_length] = append(pool[node_temp.right_length], node_temp)
							if !is_in_nodes(node_temp, new_data) {
								new_data = append(new_data, node_temp)
							}
						}
					} else {
						//fmt.Println(pool[curr_p][j], pool[curr_p][k], "有重合元素,不予连接", pool[curr_p][j].element, pool[curr_p][k].element)
					}
				} else if is_left_to_right(pool[curr_p][l], pool[curr_p][r], curr_p) {
					//fmt.Print("左和右相等")
					//fmt.Println(pool[curr_p][j], pool[curr_p][k])
					if !hasDuplicate(pool[curr_p][l].element, pool[curr_p][r].element) {
						node_temp := newnode_left_to_right(pool[curr_p][l], pool[curr_p][r])
						if detection(node_temp, node_num) {
							//fmt.Println(pool[curr_p][j], pool[curr_p][k], "左与右连接产生一个新的结构", node_temp)
							pool[node_temp.left_length] = append(pool[node_temp.left_length], node_temp)
							pool[node_temp.right_length] = append(pool[node_temp.right_length], node_temp)
							if !is_in_nodes(node_temp, new_data) {
								new_data = append(new_data, node_temp)
							}
						}
					} else {
						//fmt.Println(pool[curr_p][j], pool[curr_p][k], "有重合元素,不予连接", pool[curr_p][j].element, pool[curr_p][k].element)
					}
				} else if is_right_to_right(pool[curr_p][l], pool[curr_p][r], curr_p) {
					//fmt.Print("右和右相等")
					//fmt.Println(pool[curr_p][j], pool[curr_p][k])
					if !hasDuplicate(pool[curr_p][l].element, pool[curr_p][r].element) {
						node_temp := newnode_right_to_right(pool[curr_p][l], pool[curr_p][r])
						if detection(node_temp, node_num) {
							//fmt.Println(pool[curr_p][j], pool[curr_p][k], "左与左连接产生一个新的结构", node_temp)
							pool[node_temp.left_length] = append(pool[node_temp.left_length], node_temp)
							pool[node_temp.right_length] = append(pool[node_temp.right_length], node_temp)
							if !is_in_nodes(node_temp, new_data) {
								new_data = append(new_data, node_temp)
							}
						}
					} else {
						//fmt.Println(pool[curr_p][j], pool[curr_p][k], "有重合元素,不予连接", pool[curr_p][j].element, pool[curr_p][k].element)
					}
				}
			}
		}
		// 输出池子
		// fmt.Println("操作池")
		// for i, p := range pool {
		// 	fmt.Printf("第%d个小池子\n", i)
		// 	for _, n := range p {
		// 		fmt.Println(n)
		// 	}
		// }
		// fmt.Println()
		fmt.Println("当前计算平台中的数据有")
		for _, n := range new_data {
			fmt.Print("x_{", n.left, n.element, n.right, "},")
		}
		fmt.Print()
	}

end:
	fmt.Println("程序结束")
	elapsed := time.Since(start)
	fmt.Printf("程序运行时间: %s\n", elapsed)
}

// 添加无向图的边
func (g *Graph) addEdge(source, weight, destination int, Edges *[]int) {

	// 添加 source -> destination 的边
	edge := Edge{destination, weight}
	g.adjacencyList[source] = append(g.adjacencyList[source], edge)

	// 添加 destination -> source 的边（因为是无向图）
	edge = Edge{source, weight}
	g.adjacencyList[destination] = append(g.adjacencyList[destination], edge)

	*Edges = append(*Edges, weight)
	// fmt.Println("子函数打印边集")
	// for _, b := range *Edges {
	// 	fmt.Printf("%d ", b)
	// }
	// fmt.Println()

}

// 打印邻接表
func (g *Graph) printGraph() {
	fmt.Println("图的邻接表:(目标节点,路径权值)")
	for vertex, edges := range g.adjacencyList {
		fmt.Printf("顶点 %d 的相邻顶点: ", vertex)
		fmt.Println()
		for _, edge := range edges {
			fmt.Printf("(%d, %d) ", edge.destination, edge.weight)
		}
		fmt.Println()
	}
}

// 判断两个切片是否含有重合
func hasDuplicate(slice1, slice2 []int) bool {
	for _, value1 := range slice1 {
		for _, value2 := range slice2 {
			if value1 == value2 {
				return true // 一旦找到重复元素，立即返回true
			}
		}
	}
	return false // 循环结束后仍未找到重复元素，返回false
}

// 检测函数
func detection(node Node, node_num int) bool {
	//fmt.Println("判断", node, "是否合规")
	// 如果产生了像是 3<-1 2->3这样的结构，也就是小环，直接丢弃
	if node.left == node.right {
		if len(node.element) < node_num-1 {
			//fmt.Println("形成了小环", node, "直接丢弃")
			return false
		}
	}
	// 如果产生了像是 3<-1 2 3->2这样的结构，也就是半环，直接丢弃
	for i := 1; i < len(node.element)-1; i++ {
		if node.element[i] == node.left || node.element[i] == node.right {
			//fmt.Println("形成了半环", node, "直接丢弃")
			return false
		}
	}
	//fmt.Println("经判断", node, "合规")
	return true
}

// 判断两个切片是否相等的函数
func slicesAreEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, value := range slice1 {
		if value != slice2[i] {
			return false
		}
	}
	return true
}

// 判断探针是否能够与闭合这个环
func is_close_loop(loop Node, probe int, node_num int) bool {

	if len(loop.element) == node_num &&
		loop.left_length == loop.right_length &&
		loop.right_length == probe &&
		loop.left == loop.element[len(loop.element)-1] &&
		loop.right == loop.element[0] {
		return true
	} else {
		return false
	}
}

// 左左连接
func is_left_to_left(left Node, right Node, probe int) bool {
	if left.left_length == right.left_length &&
		left.left_length == probe &&
		left.left == right.element[0] &&
		right.left == left.element[0] {
		return true
	} else {
		return false
	}

}

// 左右连接
func is_left_to_right(left Node, right Node, probe int) bool {
	if left.left_length == right.right_length &&
		left.left_length == probe &&
		left.left == right.element[len(right.element)-1] &&
		right.right == left.element[0] {
		return true
	} else {
		return false
	}

}

// 右右连接
func is_right_to_right(left Node, right Node, probe int) bool {
	if left.right_length == right.right_length &&
		left.right_length == probe &&
		left.right == right.element[len(right.element)-1] &&
		right.right == left.element[len(left.element)-1] {
		return true
	} else {
		return false
	}

}

// 右左连接
func is_right_to_left(left Node, right Node, probe int) bool {
	if left.right_length == right.left_length &&
		right.left_length == probe &&
		left.right == right.element[0] &&
		right.left == left.element[len(left.element)-1] {
		return true
	} else {
		return false
	}

}

// 左左连接产生新节点
func newnode_left_to_left(left Node, right Node) Node {

	var ele_temp []int
	for l := len(left.element) - 1; l > -1; l-- {
		ele_temp = append(ele_temp, left.element[l])
	}
	ele_temp = append(ele_temp, right.element...)

	node_temp := Node{
		element:      ele_temp,
		left_length:  left.right_length,
		right_length: right.right_length,
		left:         left.right,
		right:        right.right,
	}
	return node_temp
}

// 左右连接产生新节点
func newnode_left_to_right(left Node, right Node) Node {

	// 产生一个新的结构
	// 创建新节点的数据时要遵循一定的顺序
	var ele_temp []int
	for l := len(left.element) - 1; l > -1; l-- {
		ele_temp = append(ele_temp, left.element[l])
	}
	for r := len(right.element) - 1; r > -1; r-- {
		ele_temp = append(ele_temp, right.element[r])
	}
	node_temp := Node{
		element:      ele_temp,
		left_length:  left.right_length,
		right_length: right.left_length,
		left:         left.right,
		right:        right.left,
	}
	return node_temp
}

// 右右连接产生新节点
func newnode_right_to_right(left Node, right Node) Node {

	var ele_temp []int
	ele_temp = append(ele_temp, left.element...)

	for r := len(right.element) - 1; r > -1; r-- {
		ele_temp = append(ele_temp, right.element[r])
	}
	node_temp := Node{
		element:      ele_temp,
		left_length:  left.left_length,
		right_length: right.left_length,
		left:         left.left,
		right:        right.left,
	}
	return node_temp
}

// 右左连接产生新节点
func newnode_right_to_left(left Node, right Node) Node {

	var ele_temp []int
	ele_temp = append(ele_temp, left.element...)
	ele_temp = append(ele_temp, right.element...)
	node_temp := Node{
		element:      ele_temp,
		left_length:  left.left_length,
		right_length: right.right_length,
		left:         left.left,
		right:        right.right,
	}
	return node_temp
}

// 判断两个数组是否相等
func is_two_arr_equ(ele_1 []int, ele_2 []int) bool {
	if len(ele_1) != len(ele_2) {
		return false
	} else {
		for i := 0; i < len(ele_1); i++ {
			if ele_1[i] != ele_2[i] {
				return false
			}
		}
	}

	return true
}

// 判断新产生的node是否在node数组中
func is_in_nodes(node Node, nodes []Node) bool {
	for i := 0; i < len(nodes); i++ {
		if is_two_arr_equ(node.element, nodes[i].element) {
			if node.left == nodes[i].left && node.right == nodes[i].right {
				return true
			} else if node.left == nodes[i].right && node.right == nodes[i].left {
				return true
			}
		}
	}
	return false
}
