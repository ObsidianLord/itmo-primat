package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func buildTable(root *Item, s string, m map[rune]string) {

	if root.left == nil && root.right == nil && root.value != 0 {
		m[root.value] = s
		return
	}

	buildTable(root.left, s+"0", m)
	buildTable(root.right, s+"1", m)
}

func huffman(items map[rune]int) map[rune]string {

	pq := make(PriorityQueue, len(items))

	i := 0
	for value, priority := range items {
		rating[i] = value
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	sort.Sort(custom(rating))
	heap.Init(&pq)

	var root *Item

	for pq.Len() > 1 {
		x, y := heap.Pop(&pq).(*Item), heap.Pop(&pq).(*Item)
		f := &Item{
			value:    0,
			priority: x.priority + y.priority,
			left:     x,
			right:    y,
		}

		root = f
		heap.Push(&pq, f)
	}

	table := make(map[rune]string)
	buildTable(root, "", table)

	fmt.Println("Код Хаффмана:")
	for _, v := range rating {
		fmt.Printf("%s\t%.4f\t%s\t%d\n", string(v), float64(items[v])/float64(fileLength), table[v], len(table[v]))
	}

	return table
}
