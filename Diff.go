package diff

import "github.com/Chara-X/util/slices"

func Diff[T comparable](from, to []T) []Change[T] {
	return slices.Reverse(diff(from, to, 0, 0, 0, map[[2]int][]Change[T]{}))
}
func diff[T comparable](from, to []T, i, j, offset int, cache map[[2]int][]Change[T]) []Change[T] {
	var changes []Change[T]
	if i == len(from) {
		for ; j < len(to); j++ {
			changes = append(changes, Change[T]{Kind: 0, Position: offset + i, From: *new(T), To: to[j]})
		}
		return changes
	}
	if j == len(to) {
		for ; i < len(from); i++ {
			changes = append(changes, Change[T]{Kind: 1, Position: offset + i, From: from[i], To: *new(T)})
		}
		return changes
	}
	if changes, ok := cache[[2]int{i, j}]; ok {
		return changes
	}
	if from[i] == to[j] {
		cache[[2]int{i, j}] = diff(from, to, i+1, j+1, offset, cache)
	} else {
		var options = [][]Change[T]{diff(from, to, i, j+1, offset+1, cache), diff(from, to, i+1, j, offset-1, cache), diff(from, to, i+1, j+1, offset, cache)}
		var index = slices.MinBy(options, func(e []Change[T]) int { return len(e) }, func(i int, e []Change[T]) int { return i })
		cache[[2]int{i, j}] = append(options[index], Change[T]{Kind: index, Position: offset + i, From: from[i], To: to[j]})
	}
	return cache[[2]int{i, j}]
}
