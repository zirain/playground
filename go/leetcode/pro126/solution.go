package pro126

import (
	"math"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	res := [][]string{}
	ids := buildIdsMap(beginWord, wordList)
	if _, ok := ids[beginWord]; !ok {
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList) - 1
	}

	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}

	edges := buildEdges(wordList)
	cost := initCosts(len(wordList), beginWord, ids)

	queue := [][]int{[]int{ids[beginWord]}}
	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now)-1]
		if last == ids[endWord] {
			tmp := []string{}
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			res = append(res, tmp)
		} else {
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last] + 1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}

func initCosts(n int, beginWord string, ids map[string]int) []int {
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0
	return cost
}

func buildIdsMap(beginWord string, wordList []string) map[string]int {
	ids := map[string]int{}
	for i, word := range wordList {
		ids[word] = i
	}

	return ids
}

func buildEdges(wordList []string) [][]int {
	c := len(wordList)
	edges := make([][]int, c)
	for i := 0; i < c; i++ {
		for j := i + 1; j < c; j++ {
			if canTransform(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	return edges
}

func canTransform(from, to string) bool {
	if from == to || len(from) != len(to) {
		return false
	}

	diffChars := 0
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			diffChars++
		}

		if diffChars > 1 {
			return false
		}
	}

	if diffChars == 1 {
		return true
	} else {
		return false
	}
}
