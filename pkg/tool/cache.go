package tool

type caches [][]vertex

func (c caches) Size() int {
	size := 0
	for i := 0; i < len(c); i++ {
		size += len(c[i])
	}
	return size
}

func (c caches) isEmpty() bool {
	for i := 0; i < len(c); i++ {
		if len(c[i]) > 0 {
			return false
		}
	}
	return true
}

func (c *caches) Add(n, x, y int) {
	//findIdx := c.Find(n, x, y)
	insertIdx := c.FindFirstGE(n, x, y)
	//if findIdx == -1 {
	v := vertex{x, y}
	if insertIdx == -1 {
		for i := len(*c); i < n+1; i++ {
			*c = append(*c, make([]vertex, 0))
		}
		(*c)[n] = append((*c)[n], v)
	} else if insertIdx > -1 && !c.isEqual(n, insertIdx, x, y) {
		//(*c)[n] = append((*c)[n][:insertIdx], append([]vertex{vertex{x, y}}, (*c)[n][insertIdx:]...)...)
		(*c)[n] = append((*c)[n], v)
		for i := len((*c)[n]) - 1; i > insertIdx; i-- {
			(*c)[n][i] = (*c)[n][i-1]
		}
		(*c)[n][insertIdx] = v
	}
}

func (c *caches) Del(n, x, y int) {
	if len(*c) > n {
		idx := -1
		for i, v := range (*c)[n] {
			if v.x == x && v.y == y {
				idx = i
				break
			}
		}
		if idx > -1 {
			(*c)[n] = append((*c)[n][:idx], (*c)[n][idx+1:]...)
		}
	}
}

func (c caches) Find(n, x, y int) int {
	if len(c) <= n {
		return -1
	}
	m := len(c[n])
	return c.bSearchEqualInternally(n, 0, m-1, x, y)
}

func (c caches) FindFirstGE(n, x, y int) int {
	if len(c) <= n {
		return -1
	}
	m := len(c[n])
	return c.bSearchFirstGreatEqualInternally(n, 0, m-1, x, y)
}

func (c caches) bSearchEqualInternally(n, low, high, x, y int) int {
	if low > high {
		return -1
	}

	mid := low + ((high - low) >> 1)
	if c.isEqual(n, mid, x, y) {
		return mid
	}

	if c.isGreatThan(n, mid, x, y) {
		return c.bSearchEqualInternally(n, low, mid-1, x, y)
	}

	return c.bSearchEqualInternally(n, mid+1, high, x, y)
}

func (c caches) bSearchFirstGreatEqualInternally(n, low, high, x, y int) int {
	if low > high {
		return -1
	}

	if c.isGreatEqual(n, low, x, y) {
		return low
	}
	mid := low + ((high - low) >> 1)
	if c.isLessThan(n, mid, x, y) {
		return c.bSearchFirstGreatEqualInternally(n, mid+1, high, x, y)
	}
	return c.bSearchFirstGreatEqualInternally(n, low, mid, x, y)
}

func (c caches) isEqual(imgIdx, vertexIdx, x, y int) bool {
	return len(c) > imgIdx && len(c[imgIdx]) > vertexIdx && c[imgIdx][vertexIdx].x == x && c[imgIdx][vertexIdx].y == y
}

func (c caches) isGreatThan(imgIdx, vertexIdx, x, y int) bool {
	return len(c) > imgIdx && len(c[imgIdx]) > vertexIdx && c[imgIdx][vertexIdx].x > x || (c[imgIdx][vertexIdx].x == x && c[imgIdx][vertexIdx].y > y)
}

func (c caches) isGreatEqual(imgIdx, vertexIdx, x, y int) bool {
	return len(c) > imgIdx && len(c[imgIdx]) > vertexIdx && c[imgIdx][vertexIdx].x > x || (c[imgIdx][vertexIdx].x == x && c[imgIdx][vertexIdx].y >= y)
}

func (c caches) isLessThan(imgIdx, vertexIdx, x, y int) bool {
	return len(c) <= imgIdx || len(c[imgIdx]) <= vertexIdx || c[imgIdx][vertexIdx].x < x || (c[imgIdx][vertexIdx].x == x && c[imgIdx][vertexIdx].y < y)
}

func (c caches) isLessEqual(imgIdx, vertexIdx, x, y int) bool {
	return len(c) <= imgIdx || len(c[imgIdx]) <= vertexIdx || c[imgIdx][vertexIdx].x < x || (c[imgIdx][vertexIdx].x == x && c[imgIdx][vertexIdx].y <= y)
}

func (c caches) FindLastLE(n, x, y int) int {
	if len(c) <= n {
		return -1
	}
	m := len(c[n])
	return c.bSearchLastLessEqualInternally(n, 0, m-1, x, y)
}

func (c caches) bSearchLastLessEqualInternally(n, low, high, x, y int) int {
	if low > high {
		return -1
	}

	if c.isLessEqual(n, high, x, y) {
		return high
	}

	mid := low + ((high - low) >> 1)
	if c.isGreatThan(n, mid, x, y) {
		return c.bSearchLastLessEqualInternally(n, low, mid-1, x, y)
	} else {
		return c.bSearchLastLessEqualInternally(n, mid, high-1, x, y)
	}

}
