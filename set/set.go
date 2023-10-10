package set

type Set[ItemsType comparable] struct {
	items map[ItemsType]bool
}

// NewSet - конструктор Set
func NewSet[ItemsType comparable]() *Set[ItemsType] {
	s := &Set[ItemsType]{}

	s.items = make(map[ItemsType]bool)

	return s
}

// Add добавляет элемент в множество
func (s *Set[ItemsType]) Add(t ItemsType) {
	s.items[t] = true
}

// Clear очищает множество
func (s *Set[ItemsType]) Clear() {
	s.items = make(map[ItemsType]bool)
}

// Remove удаляет элемент по значению
func (s *Set[ItemsType]) Remove(item ItemsType) {
	delete(s.items, item)
}

// Contains проверяет на вхождение в множество элемент
func (s *Set[ItemsType]) Contains(item ItemsType) bool {
	_, exists := s.items[item]

	return exists
}

// Items возвращает элементы множества
func (s *Set[ItemsType]) Items() []ItemsType {
	items := make([]ItemsType, 0, len(s.items))

	for item := range s.items {
		items = append(items, item)
	}

	return items
}

// IntersectWith возвращает пересечение с множеством other
func (s *Set[ItemsType]) IntersectWith(other *Set[ItemsType]) *Set[ItemsType] {
	intersection := NewSet[ItemsType]()

	for item := range s.items {
		if other.Contains(item) {
			intersection.Add(item)
		}
	}

	return intersection
}
