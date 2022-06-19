package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}
type SkipList struct {
	key   string
	value int
	size  int
	next  []*SkipList
}
type AVL_Tree struct {
	k                   string
	v, balance          int
	left, right, parent *AVL_Tree
}
type Tree struct {
	tree *AVL_Tree
}

func InitSkipList(m int) *SkipList {
	l := new(SkipList)
	l.size = m
	l.next = make([]*SkipList, m)
	for i := 0; i < l.size; i++ {
		l.next[i] = nil
	}
	return l
}
func (l *SkipList) Skip(k string) []*SkipList {
	x := l
	p := make([]*SkipList, l.size)
	for i := l.size - 1; i >= 0; i-- {
		for x.next[i] != nil && k < x.next[i].key {
			x = x.next[i]
		}
		p[i] = x
	}
	return p
}
func (l *SkipList) Lookup(s string) (x int, exist bool) {
	p := l.Skip(s)
	X := p[0].Succ()
	if X == nil || X.key != s {
		return 0, false
	} else {
		return X.value, true
	}
	return
}
func (x *SkipList) Succ() *SkipList {
	return x.next[0]
}
func (l *SkipList) Assign(s string, x int) {
	p := l.Skip(s)
	X := new(SkipList)
	X.next = make([]*SkipList, l.size)
	X.value = x
	X.key = s
	r := rand.Int() * 2
	var i int
	i = 0
	for (i < l.size) && (r%2 == 0) {
		X.next[i] = p[i].next[i]
		p[i].next[i] = X
		i++
		r /= 2
	}
	for i < l.size {
		X.next[i] = nil
		i++
	}
}
func (t *Tree) ReplaceNode(x, y *AVL_Tree) {
	if x == t.tree {
		t.tree = y
	} else {
		p := x.parent
		if y != nil {
			y.parent = p
		}
		if p.left == x {
			p.left = y
		} else {
			p.right = y
		}
	}
}
func (t *Tree) RotateLeft(x *AVL_Tree) {
	y := x.right
	t.ReplaceNode(x, y)
	b := y.left
	if b != nil {
		b.parent = x
	}
	x.right = b
	x.parent = y
	y.left = x
	x.balance--
	if y.balance > 0 {
		x.balance -= y.balance
	}
	y.balance--
	if x.balance < 0 {
		y.balance += x.balance
	}
}
func (t *Tree) RotateRight(x *AVL_Tree) {
	y := x.left
	t.ReplaceNode(x, y)
	b := y.right
	if b != nil {
		b.parent = x
	}
	y.right = x
	x.parent = y
	x.left = b
	x.balance++
	if y.balance < 0 {
		x.balance -= y.balance
	}
	y.balance++
	if x.balance > 0 {
		y.balance += x.balance
	}
	return
}
func Descend(t *Tree, k string) *AVL_Tree {
	x := t.tree
	for x != nil && x.k != k {
		if k < x.k {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}
func Insert(t *Tree, k string, v int) *AVL_Tree {
	y := new(AVL_Tree)
	y.k = k
	y.v = v
	y.parent = nil
	y.left = nil
	y.right = nil
	if t.tree == nil {
		t.tree = y
	} else {
		x := t.tree
		for {
			if k < x.k {
				if x.left == nil {
					x.left = y
					y.parent = x
					break
				}
				x = x.left
			} else {
				if x.right == nil {
					x.right = y
					y.parent = x
					break
				}
				x = x.right
			}
		}
	}
	return y
}
func (t *Tree) Assign(k string, v int) {
	a := Insert(t, k, v)
	a.balance = 0
	for {
		x := a.parent
		if x == nil {
			break
		}
		if a == x.left {
			x.balance--
			if x.balance == 0 {
				break
			}
			if x.balance == -2 {
				if a.balance == 1 {
					t.RotateLeft(a)
				}
				t.RotateRight(x)
				break
			}
		} else {
			x.balance++
			if x.balance == 0 {
				break
			}
			if x.balance == 2 {
				if a.balance == -1 {
					t.RotateRight(a)
				}
				t.RotateLeft(x)
				break
			}
		}
		a = x
	}
}
func (t *Tree) Lookup(s string) (int, bool) {
	x := Descend(t, s)
	if x == nil {
		return 0, false
	}
	return x.v, true
}
func lex(sentence string, arr AssocArray) []int {
	ans := make([]int, 0, len(sentence))
	s := strings.Fields(sentence)
	count := 0
	for i := 0; i < len(s); i++ {
		k, r := arr.Lookup(s[i])
		if r {
			ans = append(ans, k+1)
		} else {
			arr.Assign(s[i], count)
			ans = append(ans, count+1)
			count++
		}
	}
	return ans
}
func main() {
	in := bufio.NewReader(os.Stdin)
	example, _ := in.ReadString('\n')
	var avl *Tree = new(Tree)
	avl_res := lex(example, avl)
	for _, x := range avl_res {
		fmt.Printf("%d ", x)
	}
}
