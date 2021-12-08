package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

type Board []int

func (b Board) CheckRow(xs []int) bool {
	result := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			idx := y*5 + x
			for _, x := range xs {
				if b[idx] == x {
					result++
				}
			}

		}
	}
	log.Println("check row", result)
	return result == 5
}

func (b Board) CheckColumn(xs []int) bool {
	result := 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			idx := y*5 + x
			for _, x := range xs {
				if b[idx] == x {
					result++
				}
			}
		}
	}
	log.Println("check column", result)
	return result == 5
}

type Boards []Board

func readBoards(br *bufio.Reader, nLines, nBoards int) (Boards, error) {
	board := make(Board, 0)
	boards := make(Boards, 0)

	for i := 0; i < nBoards; i++ {
		br.ReadString('\n')
		for y := 0; y < nLines; y++ {
			line, err := br.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					return boards, nil
				}
				return boards, err
			}
			digits := strings.Fields(line)
			for _, digit := range digits {
				n, err := strconv.Atoi(digit)
				if err != nil {
					return nil, err
				}
				board = append(board, n)
			}

		}
		boards = append(boards, board)
		board = nil
	}
	return boards, nil
}
