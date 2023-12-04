package main

import "fmt"

type songType interface{}

type songNode struct {
	id          int
	songContent songType
	prev, next  *songNode
}

type cachedSongs struct {
	valMap     map[int]*songNode
	head, tail *songNode
	cap, len   int
}

func InitCache(capacity int) cachedSongs {
	head := &songNode{-1, -1, nil, nil}
	tail := &songNode{-2, -1, head, nil}
	head.next = tail

	return cachedSongs{
		valMap: make(map[int]*songNode, capacity),
		head:   head,
		tail:   tail,
		cap:    capacity,
		len:    0,
	}
}

func (c *cachedSongs) get(key int) interface{} {
	cur, ok := c.valMap[key]
	if !ok {
		return -1
	}

	// id is already the head
	if c.head.next == cur {
		return cur.songContent
	}

	// re-order otherwise
	cur.prev.next = cur.next
	cur.next.prev = cur.prev

	cur.next = c.head.next
	cur.prev = c.head
	c.head.next.prev = cur
	c.head.next = cur
	return cur.songContent

}

func (c *cachedSongs) put(key int, song songType) {
	if c.get(key) != -1 { // the id is available
		c.head.next.songContent = song
		return
	}

	var cur *songNode
	if c.len >= c.cap {
		c.len--
		cur = c.tail.prev
		cur.prev.next = c.tail
		c.tail.prev = cur.prev
		delete(c.valMap, cur.id)
	} else {
		cur = &songNode{}
	}

	c.len++
	cur.songContent = song
	cur.id = key
	cur.next = c.head.next
	cur.prev = c.head
	c.head.next.prev = cur
	c.head.next = cur
	c.valMap[key] = cur
}

func main() {
	cached := InitCache(3)
	var song1 songType = "newSong1"
	var song2 songType = "newSong2"
	var song3 songType = "newSong3"
	var song4 songType = "newSong4"
	var song5 songType = "newSong5"

	cached.put(1, song1)
	cached.put(2, song2)
	cached.put(3, song3)
	fmt.Println("Fetching Song no 3 : ", cached.get(3))

	cached.put(4, song4)
	fmt.Println("Fetching Song no 1 : ", cached.get(1))

	cached.put(5, song5)
	fmt.Println("Fetching Song no 2 : ", cached.get(2))
}
