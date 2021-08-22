package main

import (
	"log"
	"fmt"
)

func region(r,c int) int {
	//0 1 2
	//3 4 5
	//6 7 8
	region := 3 * (r/3)
	region += c/3
	return region
}

type Board struct {
	cell [9][9]uint

	/* use a bit array (1..9), representing used numbers
	   to check the "sudoku-constraints"...
	   "a number can only be used once in each column, row or field"
	*/
	row, col, region [9]int
}

func (b *Board) print() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(b.cell[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (b *Board) possible(num uint, r int, c int) bool {
	t := 1 << num
	return b.row[r] & t == 0 && b.col[c] & t == 0 && b.region[region(r,c)] & t == 0
}

func (b *Board) insert(num uint, r int, c int) {
	if !b.possible(num, r ,c) {
		log.Fatal("bad number on insert", num, r, c)
	}
	b.cell[r][c] = num
	var bit = 1 << num
	b.row[r] |= bit
	b.col[c] |= bit
	b.region[region(r,c)] |= bit
}

func (b *Board) remove(r int, c int) {
	var bit = 1 << b.cell[r][c]
	b.cell[r][c] = 0
	b.row[r] ^= bit
	b.col[c] ^= bit
	b.region[region(r,c)] ^= bit
}

//read a board from stdin
func readBoard() Board {
	var n uint
	var b Board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			x, _ := fmt.Scanf("%d", &n)
			if x != 1  || n > 9 {
				log.Fatal("whoopsie, bad input")
			}
			if n != 0 {
				b.insert(n, i, j)
			}
		}
	}
	return b
}

func (b *Board) find_empty_spot() (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.cell[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func (b *Board) solve(itercnt *int) bool {
	//simple brute force; just recurse on an empty spot
	r,c,found := b.find_empty_spot()
	if !found {
		return true //a full board means we have a solution
	}
	for n := uint(1); n <= 9; n++ {
		if b.possible(n, r, c) {
			b.insert(n, r, c)
			*itercnt++
			if b.solve(itercnt) {
				return true
			}
			b.remove(r, c)
		}
	}
	return false
}

func main() {
	var count int
	var b Board = readBoard()
	b.print()
	result := b.solve(&count)
	fmt.Println(result, "(", count, ")")
	b.print()
}
