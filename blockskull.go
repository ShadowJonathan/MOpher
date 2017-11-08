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
	"./encoding/nbt"
	"./type/direction"
)

type skullType int

const (
	skullSkeleton skullType = iota
	skullWitherSkeleton
	skullZombie
	skullPlayer
	skullCreeper
)

type skullComponent struct {
	SkullType skullType
	Rotation  int
	Facing    direction.Type
	Owner     string
	position  Position
}

func (s *skullComponent) CanHandleAction(action int) bool {
	return action == 4
}

func (s *skullComponent) Deserilize(tag *nbt.Compound) {
	t, ok := tag.Items["SkullType"].(int8)
	if !ok {
		return
	}
	s.SkullType = skullType(t)
	rot, ok := tag.Items["Rot"].(int8)
	if !ok {
		return
	}
	s.Rotation = int(rot)

	if s.SkullType != skullPlayer {
		return
	}

	owner, ok := tag.Items["Owner"].(*nbt.Compound)
	if !ok {
		return
	}
	props, ok := owner.Items["Properties"].(*nbt.Compound)
	if !ok {
		return
	}
	tex, ok := props.Items["textures"].(*nbt.List)
	if !ok || tex.Type != nbt.TagCompound || len(tex.Elements) < 1 {
		return
	}
}
