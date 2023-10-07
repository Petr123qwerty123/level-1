package set

type Item interface{}

type Set struct {
	items map[Item]bool
}

func NewSet() *Set {
	s := &Set{}

	s.items = make(map[Item]bool)

	return s
}

func (s *Set) Add(t Item) {
	s.items[t] = true
}

func (s *Set) Clear() {
	s.items = make(map[Item]bool)
}

func (s *Set) Remove(item Item) {
	delete(s.items, item)
}

func (s *Set) Contains(item Item) bool {
	_, exists := s.items[item]

	return exists
}

func (s *Set) Items() []Item {
	items := make([]Item, 0, len(s.items))

	for item := range s.items {
		items = append(items, item)
	}

	return items
}

func (s *Set) IntersectWith(other *Set) *Set {
	intersection := NewSet()

	for item := range s.items {
		if other.Contains(item) {
			intersection.Add(item)
		}
	}

	return intersection
}
