// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package column

import (
	"github.com/apache/arrow/go/v9/arrow"
	"github.com/apache/arrow/go/v9/arrow/array"
	"github.com/apache/arrow/go/v9/arrow/memory"
)

// I8Column is a column of int8 data.
type I8Column struct {
	// name of the column.
	name string
	// data of the column.
	data []*int8
}

// U8Column is a column of int8 data.
type I16Column struct {
	// name of the column.
	name string
	// data of the column.
	data []*int16
}

// I32Column is a column of int32 data.
type I32Column struct {
	// name of the column.
	name string
	// data of the column.
	data []*int32
}

// I64Column is a column of int64 data.
type I64Column struct {
	// name of the column.
	name string
	// data of the column.
	data []*int64
}

// MakeI8Column creates a new I8 column.
func MakeI8Column(name string, data *int8) I8Column {
	return I8Column{
		name: name,
		data: []*int8{data},
	}
}

// MakeI16Column creates a new I16 column.
func MakeI16Column(name string, data *int16) I16Column {
	return I16Column{
		name: name,
		data: []*int16{data},
	}
}

// MakeI32Column creates a new I32 column.
func MakeI32Column(name string, data *int32) I32Column {
	return I32Column{
		name: name,
		data: []*int32{data},
	}
}

// MakeI64Column creates a new I64 column.
func MakeI64Column(name string, data *int64) I64Column {
	return I64Column{
		name: name,
		data: []*int64{data},
	}
}

func MakeI64sColumn(name string, data []*int64) I64Column {
	return I64Column{
		name: name,
		data: data,
	}
}

// Name returns the name of the column.
func (c *I8Column) Name() string {
	return c.name
}

// Push adds a new value to the column.
func (c *I8Column) Push(data *int8) {
	c.data = append(c.data, data)
}

// Len returns the number of values in the column.
func (c *I8Column) Len() int {
	return len(c.data)
}

// NewI8SchemaField creates a I8 schema field.
func (c *I8Column) NewI8SchemaField() *arrow.Field {
	return &arrow.Field{Name: c.name, Type: arrow.PrimitiveTypes.Int8}
}

// NewI8Builder creates and initializes a new Int8Builder for the column.
func (c *I8Column) NewI8Builder(allocator *memory.GoAllocator) *array.Int8Builder {
	builder := array.NewInt8Builder(allocator)
	builder.Reserve(len(c.data))
	for _, v := range c.data {
		if v == nil {
			builder.AppendNull()
		} else {
			builder.UnsafeAppend(*v)
		}
	}
	c.Clear()
	return builder
}

// Clear clears the int8 data in the column but keep the original memory buffer allocated.
func (c *I8Column) Clear() {
	c.data = c.data[:0]
}

// Name returns the name of the column.
func (c *I16Column) Name() string {
	return c.name
}

// Push adds a new value to the column.
func (c *I16Column) Push(data *int16) {
	c.data = append(c.data, data)
}

// Len returns the number of values in the column.
func (c *I16Column) Len() int {
	return len(c.data)
}

// NewI16SchemaField creates a I16 schema field.
func (c *I16Column) NewI16SchemaField() *arrow.Field {
	return &arrow.Field{Name: c.name, Type: arrow.PrimitiveTypes.Int16}
}

// NewI16Builder creates and initializes a new Iint16Builder for the column.
func (c *I16Column) NewI16Builder(allocator *memory.GoAllocator) *array.Int16Builder {
	builder := array.NewInt16Builder(allocator)
	builder.Reserve(len(c.data))
	for _, v := range c.data {
		if v == nil {
			builder.AppendNull()
		} else {
			builder.UnsafeAppend(*v)
		}
	}
	c.Clear()
	return builder
}

// Clear clears the int16 data in the column but keep the original memory buffer allocated.
func (c *I16Column) Clear() {
	c.data = c.data[:0]
}

// Name returns the name of the column.
func (c *I32Column) Name() string {
	return c.name
}

// Push adds a new value to the column.
func (c *I32Column) Push(data *int32) {
	c.data = append(c.data, data)
}

// Len returns the number of values in the column.
func (c *I32Column) Len() int {
	return len(c.data)
}

// Clear clears the int32 data in the column but keep the original memory buffer allocated.
func (c *I32Column) Clear() {
	c.data = c.data[:0]
}

// NewI32SchemaField creates a I32 schema field.
func (c *I32Column) NewI32SchemaField() *arrow.Field {
	return &arrow.Field{Name: c.name, Type: arrow.PrimitiveTypes.Int32}
}

// NewI32Builder creates and initializes a new Iint32Builder for the column.
func (c *I32Column) NewI32Builder(allocator *memory.GoAllocator) *array.Int32Builder {
	builder := array.NewInt32Builder(allocator)
	builder.Reserve(len(c.data))
	for _, v := range c.data {
		if v == nil {
			builder.AppendNull()
		} else {
			builder.UnsafeAppend(*v)
		}
	}
	c.Clear()
	return builder
}

// Name returns the name of the column.
func (c *I64Column) Name() string {
	return c.name
}

// Push adds a new value to the column.
func (c *I64Column) Push(data *int64) {
	c.data = append(c.data, data)
}

// Len returns the number of values in the column.
func (c *I64Column) Len() int {
	return len(c.data)
}

// Clear clears the int64 data in the column but keep the original memory buffer allocated.
func (c *I64Column) Clear() {
	c.data = c.data[:0]
}

// NewI64SchemaField creates a I64 schema field.
func (c *I64Column) NewI64SchemaField() *arrow.Field {
	return &arrow.Field{Name: c.name, Type: arrow.PrimitiveTypes.Int64}
}

// NewI64Builder creates and initializes a new Iint64Builder for the column.
func (c *I64Column) NewI64Builder(allocator *memory.GoAllocator) *array.Int64Builder {
	builder := array.NewInt64Builder(allocator)
	builder.Reserve(len(c.data))
	for _, v := range c.data {
		if v == nil {
			builder.AppendNull()
		} else {
			builder.UnsafeAppend(*v)
		}
	}
	c.Clear()
	return builder
}
