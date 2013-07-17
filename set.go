/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Package mapset implements a simple and generic set collection.
// Items stored within it are unordered and unique
// It supports typical set operations: membership testing, intersection, union, difference and symmetric difference

package mapset

import "fmt"
import "strings"

// The primary type that represents a set
type Set struct {
	set map[interface{}]_placeHolder
}

type _placeHolder struct{}

// Creates and returns a pointer to an empty set.
func NewSet() *Set {
	return &Set{make(map[interface{}]_placeHolder)}
}

// Adds an item to the current set if it doesn't already exist in the set.
func (set *Set) Add(i interface{}) bool {
	_, found := set.set[i]
	set.set[i] = _placeHolder{}
	return !found //False if it existed already
}

// Determines if a given item is already in the set.
func (set *Set) Contains(i interface{}) bool {
	if _, found := set.set[i]; found {
		return found //true if it existed already
	}
	return false
}

// Determines if every item in the other set is in this set.
func (set *Set) IsSubset(other *Set) bool {
	for key, _ := range set.set {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

// Determines if every item of this set is in the other set.
func (set *Set) IsSuperset(other *Set) bool {
	for key, _ := range other.set {
		if !set.Contains(key) {
			return false
		}
	}
	return true
}

// Returns a new set with all items in both sets.
func (set *Set) Union(other *Set) *Set {
	if set != nil && other != nil {
		unionedSet := NewSet()

		for key, _ := range set.set {
			unionedSet.Add(key)
		}
		for key, _ := range other.set {
			unionedSet.Add(key)
		}
		return unionedSet
	}
	return nil
}

// Returns a new set with items that exist only in both sets.
func (set *Set) Intersect(other *Set) *Set {
	if set != nil && other != nil {
		intersectedSet := NewSet()
		var smallerSet *Set = nil
		var largerSet *Set = nil

		//figure out the smaller of the two sets and loop on that one as an optimization.
		if set.Size() < other.Size() {
			smallerSet = set
			largerSet = other
		} else {
			smallerSet = other
			largerSet = set
		}

		for key, _ := range smallerSet.set {
			if largerSet.Contains(key) {
				intersectedSet.Add(key)
			}
		}
		return intersectedSet
	}
	return nil
}

// Returns a new set with items in the current set but not in the other set
func (set *Set) Difference(other *Set) *Set {
	if set != nil && other != nil {
		differencedSet := NewSet()

		for key, _ := range set.set {
			if !other.Contains(key) {
				differencedSet.Add(key)
			}
		}

		return differencedSet
	}
	return nil
}

// Returns a new set with items in the current set or the other set but not in both.
func (set *Set) SymmetricDifference(other *Set) *Set {
	if set != nil && other != nil {
		aDiff := set.Difference(other)
		bDiff := other.Difference(set)

		symDifferencedSet := aDiff.Union(bDiff)

		return symDifferencedSet
	}
	return nil
}

// Clears the entire set to be the empty set.
func (set *Set) Clear() {
	set.set = make(map[interface{}]_placeHolder)
}

// Allows the removal of a single item in the set.
func (set *Set) Remove(i interface{}) {
	delete(set.set, i)
}

// Size returns the how many items are currently in the set.
func (set *Set) Size() int {
	return len(set.set)
}

// Equal determines if two sets are equal to each other.
// If they both are the same size and have the same items they are considered equal.
// Order of items is not relevent for sets to be equal.
func (set *Set) Equal(other *Set) bool {
	if set != nil && other != nil {
		if !(set.Size() == other.Size()) {
			return false
		} else {
			for key, _ := range set.set {
				if !other.Contains(key) {
					return false
				}
			}
			return true
		}
	}
	return false
}

// Provides a convenient string representation of the current state of the set.
func (set *Set) String() string {
	items := make([]string, 0, len(set.set))

	for key := range set.set {
		items = append(items, fmt.Sprintf("%v", key))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

func (set *Set) Get() []interface{} {
	items := make([]interface{}, 0, len(set.set))
	for key := range set.set {
		items = append(items, key)
	}
	return items
}
