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
	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
)

// BinaryColumn is a column of binary data.
type BinaryColumn struct {
	// Name of the column.
	name string
	// Data of the column.
	data []*[]byte
}

// MakeBinaryColumn creates a new Binary column.
func MakeBinaryColumn(name string, data *[]byte) BinaryColumn {
	return BinaryColumn{
		name: name,
		data: []*[]byte{data},
	}
}

// NewBinaryBuilder creates and initializes a new BinaryBuilder for the column.
func (c *BinaryColumn) NewBinaryBuilder(allocator *memory.GoAllocator) *array.BinaryBuilder {
	builder := array.NewBinaryBuilder(allocator, arrow.BinaryTypes.Binary)
	builder.Reserve(len(c.data))
	for _, v := range c.data {
		if v == nil {
			builder.AppendNull()
		} else {
			builder.Append(*v)
		}
	}
	c.Clear()
	return builder
}

// Name returns the name of the column.
func (c *BinaryColumn) Name() string {
	return c.name
}

// Push adds a new value to the column.
func (c *BinaryColumn) Push(data *[]byte) {
	c.data = append(c.data, data)
}

// Len returns the number of values in the column.
func (c *BinaryColumn) Len() int {
	return len(c.data)
}

// Clear clears the bool data in the column but keep the original memory buffer allocated.
func (c *BinaryColumn) Clear() {
	c.data = c.data[:0]
}

// MakeBinarySchemaField creates a Binary schema field.
func (c *BinaryColumn) MakeBinarySchemaField() arrow.Field {
	return arrow.Field{Name: c.name, Type: arrow.BinaryTypes.Binary}
}