/*
Copyright 2014 Workiva, LLC
Copyright 2018 Packet Clearing House

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package augmentedtree

import (
	"fmt"
	"math/big"
)

type dimension struct {
	low, high big.Int
}

type mockInterval struct {
	dimensions []*dimension
	id         uint64
}

func (mi *mockInterval) checkDimension(dimension uint64) {
	if dimension > uint64(len(mi.dimensions)) {
		panic(fmt.Sprintf(`Dimension: %d out of range.`, dimension))
	}
}

func (mi *mockInterval) LowAtDimension(dimension uint64) big.Int {
	return mi.dimensions[dimension-1].low
}

func (mi *mockInterval) HighAtDimension(dimension uint64) big.Int {
	return mi.dimensions[dimension-1].high
}

func (mi *mockInterval) OverlapsAtDimension(iv Interval, dimension uint64) bool {
	miHigh := mi.HighAtDimension(dimension)
	miLow := mi.LowAtDimension(dimension)
	ivHigh := iv.HighAtDimension(dimension)
	ivLow := iv.LowAtDimension(dimension)
	return miHigh.Cmp(&ivLow) > 0 && miLow.Cmp(&ivHigh) < 0
}

func (mi mockInterval) ID() uint64 {
	return mi.id
}

func constructSingleDimensionInterval(low, high int64, id uint64) *mockInterval {
	return &mockInterval{[]*dimension{&dimension{low: *big.NewInt(low), high: *big.NewInt(high)}}, id}
}

func constructMultiDimensionInterval(id uint64, dimensions ...*dimension) *mockInterval {
	return &mockInterval{dimensions: dimensions, id: id}
}

func constructDimension(low, high int64) *dimension {
	return &dimension{low: *big.NewInt(low), high: *big.NewInt(high)}
}
