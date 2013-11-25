package vlist

import (
       "fmt"
)

type VList struct {
     base * VBlock
     offset int
}

type VBlock struct {
     base * VBlock
     block []VElement
}

type VElement interface {}

func EmptyList() (VList){
	block := make([]VElement, 1)
	Vblock := VBlock{base: nil,
     	    	         block: block}
     return VList{base: &Vblock, offset: 0}
}

func NewList(first VElement) (VList){
	block := make([]VElement, 1)
	block[0] = first
	Vblock := VBlock{base: nil,
     	    	         block: block}
     return VList{base: &Vblock, offset: 0}
}

func (l VList) Cons(next VElement) VList {
	if l.offset == 0{
		dim := len(l.base.block)*2
		blo := VBlock{base: l.base,
		              block: make([]VElement, dim)}
		blo.block[dim-1] = next
		return VList{base: &blo,
		             offset: dim-1}
	}else{
		l.base.block[l.offset-1] = next
		return VList{base: l.base, offset: l.offset-1}
	}
}

func (l VList) String() string{
	if l.Len() == 0 {return "[]"}
	str := "["
	i := l.offset
	b := l.base
	for {
		if i < len(b.block) {
			str += fmt.Sprintf("%v", b.block[i])
                        i++
		}else{
			if b.base != nil {
				b = b.base
				i = 0
			}else {
				return str + "]"}
		}
	}
}

func (l VList) First() VElement{
	return l.base.block[l.offset]
}

func (l VList) Nth(i int) VElement{
	if i > (len(l.base.block)*2-l.offset-2){
		panic("Index out of collection")
	}else{
		l := l.base
		for i >= len(l.block) {
			i -= len(l.block)
			l = l.base}
		return l.block[i]}
}

func (l VList) Len() int {
	if (l.offset == 0 && l.base.block[0] == nil) {return 0}
	return len(l.base.block)*2-l.offset-1
}

func (l VList) Rest() VList{
	if l.Len() == 1 {
		return EmptyList()}
	if l.offset == len(l.base.block)-1 {
		return VList{base: l.base.base, offset: 0}
	}else{
		return VList{base: l.base, offset: l.offset+1}}
}