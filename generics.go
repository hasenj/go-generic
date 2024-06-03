package generic

import (
	"testing"
)

func Last[T any](list []T) T {
	return list[len(list)-1]
}

func OneOf[T comparable](value T, matches []T) bool {
	for _, m := range matches {
		if value == m {
			return true
		}
	}
	return false
}

// Append adds an element to the list and returns its index
func Append[T any](list *[]T, items ...T) int {
	*list = append(*list, items...)
	return len(*list) - 1
}

// AllocAppend appends a zero item to the list and returns a pointer to the added item
func AllocAppend[T any](list *[]T) *T {
	var t T
	idx := Append(list, t)
	return &(*list)[idx]
}

func ShrinkTo[T any](list *[]T, toLen int) {
	if len(*list) > toLen {
		*list = (*list)[:toLen]
	}
}

func ResetSlice[T any](list *[]T) {
	ShrinkTo(list, 0)
}

func InsertAt[T any](list *[]T, idx int, items ...T) {
	Append(list, items...)
	count := len(items)
	copy((*list)[idx+count:], (*list)[idx:])
	copy((*list)[idx:idx+count], items)
}

func RemoveAt[T any](list *[]T, idx int, count int) {
	var zero T
	copy((*list)[idx:], (*list)[idx+count:])
	for i := len(*list) - count; i < len(*list); i++ {
		(*list)[i] = zero
	}
	ShrinkTo(list, len(*list)-count)
}

func SlicesEqual[T comparable](list1 []T, list2 []T) bool {
	if len(list1) != len(list2) {
		return false
	}
	for i := range list1 {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func InitMap[K comparable, V any](m *map[K]V) {
	*m = make(map[K]V)
}

func InitSlice[T any](m *[]T) {
	*m = make([]T, 0)
}

func GrowSlice[T any](m *[]T, length int) {
	if cap(*m) < length {
		// the only thing we can do here AFACT is to allocate new size and copy
		newlist := make([]T, length)
		copy(newlist, *m)
		*m = newlist
	} else if len(*m) < length { // but within capacity
		*m = (*m)[:length]
	} else { // desired length is less than actual length; do nothing
	}
}

func EnsureSliceNotNil[T any](m *[]T) {
	if *m == nil {
		InitSlice(m)
	}
}

func EnsureMapNotNil[K comparable, V any](m *map[K]V) {
	if *m == nil {
		InitMap(m)
	}
}

// Slice simplifies the syntax of a struct literal by removing the ugly curly braces
func Slice[T any](items ...T) []T {
	return items
}

func CappedLength[T any](slice []T, maxLen int, slack int) []T {
	if len(slice) > maxLen+slack {
		return slice[:maxLen]
	} else {
		return slice
	}
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func MustOK(err error) {
	if err != nil {
		panic(err)
	}
}

func MustTrue(b bool, msg error) {
	if !b {
		panic(msg)
	}
}

func MustNotNil[T any](p *T) *T {
	if p == nil {
		panic("Nil")
	}
	return p
}

// Reset a thing to its zero value
func Reset[T any](ptr *T) {
	var zero T
	*ptr = zero
}

func TryAndLog[T any](value T, err error) T {
	if err != nil {
		LogError(err)
	}
	return value
}

type numeric interface {
	~int | ~int32 | ~float64 | ~float32 | ~uint8
}

func Max[T numeric](a T, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Min[T numeric](a T, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}

func Assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func TestExpect(t *testing.T, cond bool, msg string) bool {
	if !cond {
		t.Error(msg)
	}
	return cond
}

func TestExpectf(t *testing.T, cond bool, format string, args ...any) bool {
	if !cond {
		t.Errorf(format, args...)
	}
	return cond
}

func Clone[T any](s []T) []T {
	out := make([]T, len(s), cap(s))
	copy(out, s)
	return out
}

func Reverse[T any](a []T) {
	// from https://github.com/golang/go/wiki/SliceTricks#reversing
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
}

type int_enum interface {
	~int | ~byte
}

func IntAbs[T int_enum](a T) int {
	if a < 0 {
		return int(-a)
	} else {
		return int(a)
	}
}

func IndexOf[T comparable](list []T, item T) int {
	for idx := range list {
		if list[idx] == item {
			return idx
		}
	}
	return -1
}

// is `item` in the range [start, end)
func IsBetween[T numeric](start, item, end T) bool {
	return start <= item && item < end
}

func Clamp[T numeric](start T, item *T, end T) {
	if *item < start {
		*item = start
	}
	if *item > end {
		*item = end
	}
}

type Inty interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint64 | ~uint32 | ~uint16 | ~uint8
}

type Floaty interface {
	~float32 | ~float64
}

func MulF[T Inty, F Floaty](n T, f F) T {
	return T(F(n) * f)
}

// (a * b) / c
func MulDiv[T Inty](a, b, c T) T {
	return T(float64(a*b) / float64(c))
}

type TypedBucket[T any] struct {
	Items      [4 * 1024]T
	Next       int
	NextBucket *TypedBucket[T]
}

type TypedArena[T any] struct {
	First   *TypedBucket[T]
	Current *TypedBucket[T]
}

func NewTypedArena[T any]() *TypedArena[T] {
	bucket := new(TypedBucket[T])
	return &TypedArena[T]{
		First:   bucket,
		Current: bucket,
	}
}

func (a *TypedArena[T]) Allocate() *T {
	if a.Current.Next >= len(a.Current.Items) {
		a.Current.NextBucket = new(TypedBucket[T])
		a.Current = a.Current.NextBucket
	}
	item := &a.Current.Items[a.Current.Next]
	a.Current.Next++
	return item
}

func (a *TypedArena[T]) Iterate(visitFn func(index int, item *T) bool) {
	bucket := a.First
	index := 0
	for bucket != nil {
		slice := bucket.Items[:bucket.Next]
		for bucketItemIndex := range slice {
			cont := visitFn(index, &bucket.Items[bucketItemIndex])
			if !cont {
				return
			}
			index++
		}
		bucket = bucket.NextBucket
	}
}

func (a *TypedArena[T]) IterateBuckets(visitFn func(items []T)) {
	bucket := a.First
	for bucket != nil {
		slice := bucket.Items[:bucket.Next]
		visitFn(slice)
		bucket = bucket.NextBucket
	}
}

type Set[T comparable] struct {
	Map map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		Map: make(map[T]bool),
	}
}

func NewSetFrom[T comparable](items []T) *Set[T] {
	set := NewSet[T]()
	set.Add(items...)
	return set
}

func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.Map[item] = true
	}
}

func (s *Set[T]) Has(item T) bool {
	exists, _ := s.Map[item]
	return exists
}

func (s *Set[T]) Remove(item T) {
	delete(s.Map, item)
}

func SetsEqual[T comparable](s1 *Set[T], s2 *Set[T]) bool {
	for item := range s1.Map {
		if !s2.Has(item) {
			return false
		}
	}
	for item := range s2.Map {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Get a map entry with a function to create it if not existing
func MapEntry[K comparable, V any](m map[K]V, key K, fn func(k K) V) V {
	item, found := m[key]
	if !found {
		item = fn(key)
		m[key] = item
	}
	return item
}
