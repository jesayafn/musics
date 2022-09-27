package controller

import (
	"hash/fnv"
	"strconv"
)

func Id(stringInput string) (hash string) {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(stringInput))
	rawHash := algorithm.Sum32()
	hash = strconv.FormatUint(uint64(rawHash), 10)
	return hash
}
