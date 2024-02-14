func findOrder(numCourses int, prerequisites [][]int) []int {
    var result []int
    
    var queue []int
    
    indegree := make(map[int][]int)
    adjacent := make(map[int][]int)
    count := 0
    
    for _,  prerequisite := range  prerequisites {
        src :=  prerequisite[1]
        dst :=  prerequisite[0]
        
        indegree[dst] = append(indegree[dst], src)
        adjacent[src] = append(adjacent[src], dst)
    }
    
    for i := 0; i < numCourses; i++ {
        if len(indegree[i]) == 0 {
            enqueue(&queue, i)
            result = append(result, i)
        }
    }
    
    for len(queue) > 0 {
        dequeuedEle := dequeue(&queue)
        
        for _, vertex := range adjacent[dequeuedEle] {
            tmp := indegree[vertex]
            remove(&tmp, dequeuedEle)
            indegree[vertex] = tmp

            if len(indegree[vertex]) == 0 {
                enqueue(&queue, vertex)
                result = append(result, vertex)
            }
        }
        
        count += 1
    }

    if count == numCourses {
        return result
    }
    
    return []int{}
}

func remove(lst *[]int, removedEle int) {
	if lst == nil {
		panic("nil pointer")
	}

	if len(*lst) == 0 {
		panic("empty list")
	}

	for idx, num := range *lst {
		if num == removedEle {
			*lst = append((*lst)[:idx], (*lst)[idx+1:]...)
		}
	}
}

func enqueue(queue *[]int, newEle int) {
    if queue == nil {
        panic("nil pointer")
    }
    
    *queue = append(*queue, newEle)
}

func dequeue(queue *[]int) int {
    if queue == nil {
        panic("nil pointer")
    }
    
    if len(*queue) == 0 {
        panic("empty queue")
    }
    
    dequeuedEle := (*queue)[0]
    
    *queue = (*queue)[1:]
    
    return dequeuedEle
}
