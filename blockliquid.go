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

package MO

import (
	"reflect"
)

type blockLiquid struct {
	BaseBlock
	Lava  bool
	Level int `state:"level,0-15"`
}

func (l *blockLiquid) load(tag reflect.StructTag) {
	getBool := wrapTagBool(tag)
	l.Lava = getBool("lava", false)
	l.cullAgainst = false
	l.collidable = false
	if !l.Lava {
		l.translucent = true
	}
}

func (l *blockLiquid) LightReduction() int {
	if l.Lava {
		return 0
	}
	return 1
}

func (l *blockLiquid) LightEmitted() int {
	if l.Lava {
		return 15
	}
	return 0
}

func (l *blockLiquid) toData() int {
	return l.Level
}