package diff

type Change[V any] struct {
	Kind     int
	Position int
	From     V
	To       V
}
