// Copyright 2015 Matthew Collins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bit

// Set is a collection of booleans stored as bits
type Set []uint64

// NewSet allocates a new bit set that can store up to the
// passed number of bits.
func NewSet(size int) Set {
	return make(Set, (size+63)>>6)
}

// Set changes the value of the bit at the location.
func (s Set) Set(i int, v bool) {
	if v {
		s[i>>6] |= 1 << uint(i&0x3F)
	} else {
		s[i>>6] &= ^(1 << uint(i&0x3F))
	}
}

// Get returns the value of the bit at the location
func (s Set) Get(i int) bool {
	v := s[i>>6] & (1 << uint(i&0x3F))
	return v != 0
}
