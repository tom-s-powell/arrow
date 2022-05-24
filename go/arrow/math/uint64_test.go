// Code generated by type_test.go.tmpl. DO NOT EDIT.

// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package math_test

import (
	"testing"

	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/math"
	"github.com/apache/arrow/go/v9/arrow/memory"
	"github.com/stretchr/testify/assert"
)

func TestUint64Funcs_Sum(t *testing.T) {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	defer mem.AssertSize(t, 0)
	vec := makeArrayUint64(10000, mem)
	defer vec.Release()
	res := math.Uint64.Sum(vec)
	assert.Equal(t, res, uint64(49995000))
}

func TestUint64Funcs_SumEmpty(t *testing.T) {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	defer mem.AssertSize(t, 0)
	b := array.NewUint64Builder(mem)
	defer b.Release()
	vec := b.NewUint64Array()
	defer vec.Release()
	res := math.Uint64.Sum(vec)
	assert.Equal(t, res, uint64(0))
}

func makeArrayUint64(l int, mem memory.Allocator) *array.Uint64 {
	fb := array.NewUint64Builder(mem)
	defer fb.Release()
	fb.Reserve(l)
	for i := 0; i < l; i++ {
		fb.Append(uint64(i))
	}
	return fb.NewUint64Array()
}

func benchmarkUint64Funcs_Sum(b *testing.B, n int) {
	mem := memory.NewCheckedAllocator(memory.NewGoAllocator())
	defer mem.AssertSize(b, 0)
	vec := makeArrayUint64(n, mem)
	defer vec.Release()
	b.SetBytes(int64(vec.Len() * 8))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		math.Uint64.Sum(vec)
	}
}

func BenchmarkUint64Funcs_Sum_256(b *testing.B) {
	benchmarkUint64Funcs_Sum(b, 256)
}

func BenchmarkUint64Funcs_Sum_1024(b *testing.B) {
	benchmarkUint64Funcs_Sum(b, 1024)
}

func BenchmarkUint64Funcs_Sum_8192(b *testing.B) {
	benchmarkUint64Funcs_Sum(b, 8192)
}

func BenchmarkUint64Funcs_Sum_1000000(b *testing.B) {
	benchmarkUint64Funcs_Sum(b, 1e6)
}
