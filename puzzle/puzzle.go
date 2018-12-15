package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var solution string
var maxIdx int

// Node structure
type treeNode struct {
	idx       int     // Node id
	parentIdx int     // Node parent id
	move      [2]byte // Move that creates the node from parent
	rep       []byte  // Node board representation
}

func moveType1(newRep []byte, rep []byte, idx1 int, idx2 int, idx3 int) {
	copy(newRep[:], rep[:])
	newRep[idx1] = rep[idx2]
	newRep[idx3] = 'x'
}

func moveType2(newRep []byte, rep []byte, idx1 int, idx2 int, idx3 int, idx4 int, idx5 int) {
	copy(newRep[:], rep[:])
	newRep[idx1] = rep[idx2]
	newRep[idx3] = rep[idx2]
	newRep[idx4] = 'x'
	newRep[idx5] = 'x'
}

// Create node in tree
func createNode(nodeList []treeNode, nodeListIdx *int, repList []string, repListIdx *int, move [2]byte, newRep []byte, parentNodeIdx int) int {
	repeatedRep := 0
	retVal := -1
	// Check if board (representation) is repeated to avoid dummy nodes
	for k := 0; k < *repListIdx; k++ {
		if string(newRep[:]) == repList[*repListIdx] {
			repeatedRep = 1
			break
		}
	}
	if repeatedRep == 0 {
		repList[*repListIdx] = string(newRep[:])
		nodeList[*nodeListIdx].rep = make([]byte, 20)
		nodeList[*nodeListIdx].parentIdx = parentNodeIdx
		nodeList[*nodeListIdx].idx = *nodeListIdx
		copy(nodeList[*nodeListIdx].move[:], move[:])
		copy(nodeList[*nodeListIdx].rep[:], newRep[:])
		// Check for single exit move possibility
		if (newRep[17] == 'b') && (newRep[18] == 'b') {
			solution = string(nodeList[*nodeListIdx].move[:]) + "bD"
			retVal = nodeList[*nodeListIdx].parentIdx
		} else if (newRep[17] == 'x') && (newRep[18] == 'x') { // Check for double exit move possibility
			if (newRep[13] == 'b') && (newRep[14] == 'b') {
				solution = string(nodeList[*nodeListIdx].move[:]) + "BD"
				retVal = nodeList[*nodeListIdx].parentIdx
			}
		}
		// Check memory limits
		if *repListIdx >= maxIdx {
			retVal = -2
		} else {
			(*nodeListIdx)++
			(*repListIdx)++
		}
	}
	return retVal
}

// Recursive function to expand each tree level
func expand(strtIdx int, nodeList []treeNode, nodeListIdx *int, repList []string, repListIdx *int) int {
	var idx [2]int
	newRep := make([]byte, 20)
	ret := -1
	endIdx := *nodeListIdx

	for p := strtIdx; p < endIdx; p++ {
		// Extract index of empty space
		repStr := string(nodeList[p].rep[:20])
		idx[0] = strings.IndexByte(repStr, 'x')
		idx[1] = strings.LastIndexByte(repStr, 'x')

		// Loop over empty spaces for possible moves
		for n := 0; n < 2; n++ {
			idxX := idx[n] / 4
			idxY := idx[n] % 4

			var testIdx [4]int
			var testMove [4]byte
			numChk := 0

			// Test tile up
			if idxX != 0 {
				testIdx[numChk] = (idxX-1)*4 + idxY
				testMove[numChk] = 'D'
				numChk++
			}
			// Test tile down
			if idxX != 4 {
				testIdx[numChk] = (idxX+1)*4 + idxY
				testMove[numChk] = 'U'
				numChk++
			}
			// Test tile left
			if idxY != 0 {
				testIdx[numChk] = idxX*4 + (idxY - 1)
				testMove[numChk] = 'R'
				numChk++
			}
			// Test tile right
			if idxY != 3 {
				testIdx[numChk] = idxX*4 + (idxY + 1)
				testMove[numChk] = 'L'
				numChk++
			}
			// Add tree nodes based on possible tile moves
			for k := 0; k < numChk; k++ {
				switch nodeList[p].rep[testIdx[k]] {
				// Possible 1X1 single tile moves
				case 'g', 'h', 'i', 'j':
					moveType1(newRep, nodeList[p].rep[:], idx[n], testIdx[k], testIdx[k])
					ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[testIdx[k]], testMove[k]}, newRep, nodeList[p].idx)
				// Possible single space 1X2 moves
				case 'e':
					switch testMove[k] {
					case 'R':
						moveType1(newRep, nodeList[p].rep[:], idx[n], testIdx[k], idx[n]-2)
						ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[testIdx[k]], testMove[k]}, newRep, nodeList[p].idx)
					case 'L':
						moveType1(newRep, nodeList[p].rep[:], idx[n], testIdx[k], idx[n]+2)
						ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[testIdx[k]], testMove[k]}, newRep, nodeList[p].idx)
					}
				// Possible single space 2X1 moves
				case 'a', 'd', 'c', 'f':
					switch testMove[k] {
					case 'D':
						moveType1(newRep, nodeList[p].rep[:], idx[n], testIdx[k], idx[n]-8)
						ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[testIdx[k]], testMove[k]}, newRep, nodeList[p].idx)
					case 'U':
						moveType1(newRep, nodeList[p].rep[:], idx[n], testIdx[k], idx[n]+8)
						ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[testIdx[k]], testMove[k]}, newRep, nodeList[p].idx)
					}
				// Possible double space moves
				case 'x':
					if n == 0 {
						switch testMove[k] {
						case 'U':
							// Possile 2X1, 2X2 and 1X1 tile moves
							if idxY > 0 {
								tileIdx := idxX*4 + (idxY - 1)
								switch nodeList[p].rep[tileIdx] {
								case 'a', 'd', 'c', 'f':
									if nodeList[p].rep[tileIdx] == nodeList[p].rep[tileIdx+4] {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]-1, idx[1]-1)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'R'}, newRep, nodeList[p].idx)
									}
								case 'b':
									if nodeList[p].rep[tileIdx+4] == 'b' {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx+4, idx[1], idx[0]-2, idx[1]-2)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{'b', 'R'}, newRep, nodeList[p].idx)
									}
								}
								if ret > -1 {
									return ret
								}
							}
							if idxY < 3 {
								tileIdx := idxX*4 + (idxY + 1)
								switch nodeList[p].rep[tileIdx] {
								case 'a', 'd', 'c', 'f':
									if nodeList[p].rep[tileIdx] == nodeList[p].rep[tileIdx+4] {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]+1, idx[1]+1)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'L'}, newRep, nodeList[p].idx)
									}
								case 'b':
									if nodeList[p].rep[tileIdx+4] == 'b' {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx+4, idx[1], idx[0]+2, idx[1]+2)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{'b', 'L'}, newRep, nodeList[p].idx)
									}
								}
								if ret > -1 {
									return ret
								}
							}
							if idxX > 0 {
								tileIdx := (idxX-1)*4 + idxY
								switch nodeList[p].rep[tileIdx] {
								case 'a', 'd', 'c', 'f':
									moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]-8, idx[1]-8)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'D'}, newRep, nodeList[p].idx)
								case 'g', 'h', 'i', 'j':
									moveType1(newRep, nodeList[p].rep[:], idx[1], testIdx[k], idx[1]-8)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{byte(unicode.ToUpper(rune(nodeList[p].rep[tileIdx]))), 'D'}, newRep, nodeList[p].idx)
								}
								if ret > -1 {
									return ret
								}
							}
							if idxX < 3 {
								tileIdx := (idxX+2)*4 + idxY
								switch nodeList[p].rep[tileIdx] {
								case 'a', 'd', 'c', 'f':
									moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]+8, idx[1]+8)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'U'}, newRep, nodeList[p].idx)
								case 'g', 'h', 'i', 'j':
									moveType1(newRep, nodeList[p].rep[:], idx[0], tileIdx, idx[0]+8)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{byte(unicode.ToUpper(rune(nodeList[p].rep[tileIdx]))), 'U'}, newRep, nodeList[p].idx)
								}
								if ret > -1 {
									return ret
								}
							}
						// Possible 1X2, 2X2 and 1X1 tile moves
						case 'L':
							if idxY > 0 {
								tileIdx := idxX*4 + (idxY - 1)
								switch nodeList[p].rep[tileIdx] {
								case 'e':
									moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]-2, idx[1]-2)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'R'}, newRep, nodeList[p].idx)
								case 'g', 'h', 'i', 'j':
									moveType1(newRep, nodeList[p].rep[:], idx[1], tileIdx, idx[1]-2)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{byte(unicode.ToUpper(rune(nodeList[p].rep[tileIdx]))), 'R'}, newRep, nodeList[p].idx)
								}
								if ret > -1 {
									return ret
								}
							}
							if idxY < 2 {
								tileIdx := idxX*4 + (idxY + 2)
								switch nodeList[p].rep[tileIdx] {
								case 'e':
									moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]+2, idx[1]+2)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'L'}, newRep, nodeList[p].idx)
								case 'g', 'h', 'i', 'j':
									moveType1(newRep, nodeList[p].rep[:], idx[0], tileIdx, idx[0]+2)
									ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{byte(unicode.ToUpper(rune(nodeList[p].rep[tileIdx]))), 'L'}, newRep, nodeList[p].idx)
								}
								if ret > -1 {
									return ret
								}
							}
							if idxX > 0 {
								tileIdx := (idxX-1)*4 + idxY
								switch nodeList[p].rep[tileIdx] {
								case 'e':
									if nodeList[p].rep[tileIdx] == nodeList[p].rep[tileIdx+1] {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]-4, idx[1]-4)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'D'}, newRep, nodeList[p].idx)
									}
								case 'b':
									if nodeList[p].rep[tileIdx] == nodeList[p].rep[tileIdx+1] {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]-8, idx[1]-8)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'D'}, newRep, nodeList[p].idx)
									}
								}
								if ret > -1 {
									return ret
								}
							}
							if idxX < 4 {
								tileIdx := (idxX+1)*4 + idxY
								switch nodeList[p].rep[tileIdx] {
								case 'e':
									if nodeList[p].rep[tileIdx] == nodeList[p].rep[tileIdx+1] {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]+4, idx[1]+4)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'U'}, newRep, nodeList[p].idx)
									}
								case 'b':
									if nodeList[p].rep[tileIdx] == nodeList[p].rep[tileIdx+1] {
										moveType2(newRep, nodeList[p].rep, idx[0], tileIdx, idx[1], idx[0]+8, idx[1]+8)
										ret = createNode(nodeList, nodeListIdx, repList, repListIdx, [2]byte{nodeList[p].rep[tileIdx], 'U'}, newRep, nodeList[p].idx)
									}
								}
								if ret > -1 {
									return ret
								}
							}
						}
					}
				}
				if (ret > -1) || (ret == -2) {
					return ret
				}
			}
		}
	}
	// Recursive call untill solution is found
	solParentIdx := expand(endIdx, nodeList, nodeListIdx, repList, repListIdx)
	if solParentIdx < 0 {
		return solParentIdx
	}
	// Recursive solution forming
	solution = string(nodeList[solParentIdx].move[:]) + solution
	return nodeList[solParentIdx].parentIdx
}

func main() {

	// Max memory assigned for tree nodes
	maxIdx = 10000

	// Initialization
	nodeList := make([]treeNode, maxIdx+1)
	repList := make([]string, maxIdx+1)
	nodeListIdx := 1
	repListIdx := 1
	repList[0] = os.Args[1]
	nodeList[0].idx = 0
	nodeList[0].parentIdx = -1
	nodeList[0].rep = make([]byte, 20)
	copy(nodeList[0].rep[:], repList[0])

	if expand(0, nodeList, &nodeListIdx, repList, &repListIdx) == -2 {
		fmt.Println("No solution found within memory limits")
	} else {
		fmt.Printf("Number of tree nodes: %d\n", repListIdx)
		fmt.Println("Optimal solution: ", solution)
	}
}
