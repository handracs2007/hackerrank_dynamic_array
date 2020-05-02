// https://www.hackerrank.com/challenges/dynamic-array/problem?isFullScreen=true

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dynamicArray(n int, queries [][]int) []int {
	// initialise the empty sequence(s)
	var seq = make([][]int, n)

	for i := 0; i < n; i++ {
		seq[i] = make([]int, 0)
	}

	// initialise the return variable
	var result = make([]int, 0)
	var lastAnswer = 0

	// run through all the queries
	for i := 0; i < len(queries); i++ {
		var qType, qX, qY = queries[i][0], queries[i][1], queries[i][2]
		var idx = (qX ^ lastAnswer) % n

		// append or update the array
		switch qType {
		case 1:
			// append
			seq[idx] = append(seq[idx], qY)
		case 2:
			// update
			var size = len(seq[idx])
			var idx2 = qY % size
			lastAnswer = seq[idx][idx2]
			result = append(result, lastAnswer)
		}
	}

	return result
}

func main() {
	var reader = bufio.NewReader(os.Stdin)

	var strNumOfSeq, _ = reader.ReadString(' ')
	var numOfSeq, _ = strconv.Atoi(strings.TrimSpace(strNumOfSeq))

	var strNumOfQuery, _ = reader.ReadString('\n')
	var numOfQuery, _ = strconv.Atoi(strings.TrimSpace(strNumOfQuery))

	var queries = make([][]int, numOfQuery)

	for i := 0; i < numOfQuery; i++ {
		var strQueryType, _ = reader.ReadString(' ')
		var strX, _ = reader.ReadString(' ')
		var strY, _ = reader.ReadString('\n')

		queries[i] = make([]int, 3)
		queries[i][0], _ = strconv.Atoi(strings.TrimSpace(strQueryType))
		queries[i][1], _ = strconv.Atoi(strings.TrimSpace(strX))
		queries[i][2], _ = strconv.Atoi(strings.TrimSpace(strY))
	}

	var result = dynamicArray(numOfSeq, queries)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
