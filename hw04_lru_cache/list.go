package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	FrontItem *ListItem
	BackItem  *ListItem
	Length    int
}

func (l list) Len() int {
	return l.Length
}

func (l list) Front() *ListItem {
	return l.FrontItem
}

func (l list) Back() *ListItem {
	return l.BackItem
}

func (l *list) fillEmptyList(v interface{}) *ListItem {
	listItem := ListItem{Value: v}
	l.FrontItem = &listItem
	l.BackItem = &listItem
	return &listItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	var listItem *ListItem
	if l.Len() == 0 {
		listItem = l.fillEmptyList(v)
	} else {
		oldFrontItem := l.FrontItem
		listItem = &ListItem{Value: v, Next: oldFrontItem}
		oldFrontItem.Prev = listItem
		l.FrontItem = listItem
	}
	l.Length++
	return listItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	var listItem *ListItem
	if l.Len() == 0 {
		listItem = l.fillEmptyList(v)
	} else {
		oldBackItem := l.BackItem
		listItem = &ListItem{Value: v, Prev: oldBackItem}
		oldBackItem.Next = listItem
		l.BackItem = listItem
	}
	l.Length++
	return listItem
}

func (l *list) Remove(i *ListItem) {
	switch i {
	case l.Front():
		l.FrontItem = i.Next
		l.FrontItem.Prev = nil
	case l.Back():
		l.BackItem = i.Prev
		l.BackItem.Next = nil
	default:
		prev := i.Prev
		next := i.Next
		prev.Next = next
		next.Prev = prev
	}
	l.Length--
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.Length++
	i.Prev = nil
	i.Next = l.FrontItem
	l.FrontItem.Prev = i
	l.FrontItem = i
}

func NewList() List {
	return new(list)
}
