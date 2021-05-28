// MIT License
//
// Copyright (c) 2016-2021 GACHAIN
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package converter

import (
	"strconv"
	"strings"
)

// FillLeft is filling slice
func FillLeft(slice []byte) []byte {
	if len(slice) >= 32 {
		return slice
	}
	return append(make([]byte, 32-len(slice)), slice...)
}

// AddressToString converts int64 address to chain address as XXXX-...-XXXX.
func AddressToString(address int64) (ret string) {
	num := strconv.FormatUint(uint64(address), 10)
	val := []byte(strings.Repeat("0", 20-len(num)) + num)

	for i := 0; i < 4; i++ {
		ret += string(val[i*4:(i+1)*4]) + `-`
	}
	ret += string(val[16:])
	return
}
