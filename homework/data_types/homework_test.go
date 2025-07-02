package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

// go test -v homework_test.go

func ToLittleEndianSimple(number uint32) uint32 {
	return (number << 24) | (number << 8 & 0x00FF0000) | (number >> 8 & 0x0000FF00) | (number >> 24)
}

type number interface {
	uint16 | uint32 | uint64
}

func ToLittleEndian[T number](number T) T {
	size := int(unsafe.Sizeof(number))

	resBytes := make([]byte, size)

	srcPtr := unsafe.Pointer(&number)
	srcBytes := unsafe.Slice((*byte)(srcPtr), size)

	for i := 0; i < size; i++ {
		resBytes[i] = srcBytes[size-1-i]
	}

	var result T
	dstPtr := unsafe.Pointer(&result)
	dstBytes := unsafe.Slice((*byte)(dstPtr), size)

	copy(dstBytes, resBytes)

	return result
}

type testCase[T number] struct {
	name   string
	number T
	result T
}

func TestConversion(t *testing.T) {
	tests := []testCase[uint32]{
		{
			name:   "test case #1",
			number: 0x00000000,
			result: 0x00000000,
		},
		{
			name:   "test case #2",
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		{
			name:   "test case #3",
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		{
			name:   "test case #4",
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		{
			name:   "test case #5",
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToLittleEndian[uint32](test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestToLittleEndian_Uint16(t *testing.T) {
	tests := []testCase[uint16]{
		{
			name:   "test case #1",
			number: 0x1234,
			result: 0x3412,
		},
		{
			name:   "test case #2",
			number: 0xABCD,
			result: 0xCDAB,
		},
		{
			name:   "test case #3",
			number: 0x0001,
			result: 0x0100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToLittleEndian[uint16](test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestToLittleEndian_Uint64(t *testing.T) {
	tests := []testCase[uint64]{
		{
			name:   "test case #1",
			number: 0x1122334455667788,
			result: 0x8877665544332211,
		},
		{
			name:   "test case #2",
			number: 0xA1B2C3D4E5F60718,
			result: 0x1807F6E5D4C3B2A1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToLittleEndian[uint64](test.number)
			assert.Equal(t, test.result, result)
		})
	}
}

func TestConversionSimple(t *testing.T) {
	tests := []testCase[uint32]{
		{
			name:   "test case #1",
			number: 0x00000000,
			result: 0x00000000,
		},
		{
			name:   "test case #2",
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		{
			name:   "test case #3",
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		{
			name:   "test case #4",
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		{
			name:   "test case #5",
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToLittleEndianSimple(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
