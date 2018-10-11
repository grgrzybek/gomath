/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package numbers

import (
	"errors"
	"fmt"
	"strconv"
)

// Integer numbers: include Natural numbers, ZERO and Additive Inverses of Natural numbers
// https://en.wikipedia.org/wiki/Additive_inverse
//
// integer numbers are introduced from basic rules defined/found from natural numbers.
type Z struct {
	value int64

	fmt.Stringer
	// not declaring embedded interface allows us to validate the contract using this idiom:
	// var _ = ZOperations(&Z{})
	// when we declare embedded interface, this idiom doesn't validate our struct!
	// ZOperations
}

type ZOperations interface {
	Add(*Z) *Z
	Multiply(*Z) *Z
	Power(*Z) (*Z, error)
	Subtract(*Z) *Z
	Divide(*Z) (*Z, error)
	Root(*Z) (*Z, error)
	Logarithm(*Z) (*Z, error)
}

// NewZ Creates new ℤ from string
func NewZ(v string) *Z {
	z, _ := strconv.Atoi(v)
	res := &Z{value: int64(z)}

	return res
}

// DefZ Creates new ℤ as a result of subtracting two ℕs - definition of ℤ. Having ℤ defined and having
// the basic rules described for ℕ, we can implement the operations for ℤ
//
// if A < B, A is decreased to 0 and resulting (A - B) = (0 - (B - A)) is called "negative integer"
func DefZ(a *N, b *N) *Z {
	if a.value >= b.value {
		res, _ := a.Subtract(b)
		return &Z{value: int64(res.value)}
	}
	res, _ := b.Subtract(a)
	c := &N{value: ZERO.value}
	for ; c.value < res.value; c = c.addOne() {
	}
	return &Z{value: -int64(c.value)} // definition - putting "-" in front of natural number
}

func (z *Z) String() string {
	return fmt.Sprintf("%d", z.value)
}

// A + B:
//  - A >= 0, B >= 0: as in ℕ
//  - A >= 0, B < 0: A + (0 - |B|) = x -> (A + (0 - |B|)) + |B| = x + |B| -> A + ((0 - |B|) + |B|) = x + |B|
//    -> A = |B| + x -> x = A - |B|
//  - A < 0, B >= 0: A + B = B + A = B + (0 - |A|) = x -> (B + (0 - |A|)) + |A| = x + |A|
//    -> B + ((0 - |A|) + |A|) = |A| + x -> B = |A| + x -> x = B - |A|
//  - A < 0, B < 0: (0 - |A|) + (0 - |B|) = x -> (0 - |A|) + |A| + (0 - |B|) + |B| = x + |A| + |B|
//    -> ((0 - |A|) + |A|) + ((0 - |B|) + |B|) = x + |A| + |B| -> 0 = x + (|A| + |B|)
//    -> 0 = (|A| + |B|) + x -> x = 0 - (|A| + |B|) = - (|A| + |B|)
func (z *Z) Add(arg *Z) *Z {
	if z.value >= 0 && arg.value >= 0 {
		a := &N{value: uint64(z.value)}
		b := &N{value: uint64(arg.value)}
		c := a.Add(b)
		return &Z{value: int64(c.value)}
	} else if z.value >= 0 && arg.value < 0 {
		a := &N{value: uint64(z.value)}
		b := &N{value: uint64(-arg.value)} // get rid of "-" from negative integer to get ℕ
		return &Z{value: DefZ(a, b).value}
	} else if z.value < 0 && arg.value >= 0 {
		a := &N{value: uint64(-z.value)} // get rid of "-" from negative integer to get ℕ
		b := &N{value: uint64(arg.value)}
		return &Z{value: DefZ(b, a).value}
	} else {
		a := &N{value: uint64(-z.value)}   // get rid of "-" from negative integer to get ℕ
		b := &N{value: uint64(-arg.value)} // get rid of "-" from negative integer to get ℕ
		c := a.Add(b)
		return &Z{value: -int64(c.value)}
	}
}

// A * B:
//  - A >= 0, B >= 0: as in ℕ (N.Multiply)
//  - A >= 0, B < 0: A * (0 - |B|) = x -> (A * (0 - |B|)) + (A * |B|) = x + (A * |B|)
//    -> A * ((0 - |B|) + |B|) = x + (A * |B|) -> 0 = (A * |B|) + x -> x = 0 - (A * |B|) = - (A * |B|)
//  - A < 0, B >= 0: A * B = - (|A| * B) (as above)
//  - A > 0: -1 * A = -(|-1| * A) = -A
//  - -1 * -1: 0 = -1 * 0 = -1 * (1 - 1) = -1 * 1 + (-1 * -1) = -1 + (-1 * -1) = 0 -> (-1 * -1) = 0 + 1 -> -1 * -1 = 1
//  - A < 0, B < 0: (0 - |A|) * (0 - |B|) = x -> (-1 * |A|) * (-1 * |B|) = x -> -1 * |A| * -1 * |B| = x
//    -> x = (-1 * -1) * (|A| * |B|)) = 1 * (|A| * |B|) = |A| * |B|
func (z *Z) Multiply(arg *Z) *Z {
	if z.value >= 0 && arg.value >= 0 {
		a := &N{value: uint64(z.value)}
		b := &N{value: uint64(arg.value)}
		c := a.Multiply(b)
		return &Z{value: int64(c.value)}
	} else if z.value >= 0 && arg.value < 0 {
		a := &N{value: uint64(z.value)}
		b := &N{value: uint64(-arg.value)} // get rid of "-" from negative integer to get ℕ
		c := a.Multiply(b)
		return &Z{value: -int64(c.value)}
	} else if z.value < 0 && arg.value >= 0 {
		a := &N{value: uint64(-z.value)} // get rid of "-" from negative integer to get ℕ
		b := &N{value: uint64(arg.value)}
		c := a.Multiply(b)
		return &Z{value: -int64(c.value)}
	} else {
		a := &N{value: uint64(-z.value)}   // get rid of "-" from negative integer to get ℕ
		b := &N{value: uint64(-arg.value)} // get rid of "-" from negative integer to get ℕ
		c := a.Multiply(b)
		return &Z{value: int64(c.value)}
	}
}

// A ^ B:
//  - A >= 0, B >= 0: as in ℕ (N.Power)
//  - A < 0, B >= 0: as in ℕ but reimplemented with Z.Multiply
//  - B < 0: B = 0 - |B| -> B + |B| = 0 -> A ^ (B + |B|) = A ^ 0 -> A^B * A^|B| = 1 -> A^B = 1 / A^|B|
func (z *Z) Power(arg *Z) (*Z, error) {
	if z.value >= 0 && arg.value >= 0 {
		a := &N{value: uint64(z.value)} // get rid of "-" from negative integer to get ℕ
		b := &N{value: uint64(arg.value)}
		c := a.Power(b)
		return &Z{value: int64(c.value)}, nil
	} else if arg.value >= 0 {
		// reimplement from ℕ instead of delegate to ℕ
		res := NewZ("1")
		for i := int64(0); i < arg.value; i++ {
			res = res.Multiply(z)
		}
		return res, nil
	} else {
		return nil, errors.New("no solution in \u2124")
	}
}

// A - B = x -> B = A - x
//  - A >= 0, B >= 0: DefZ
//  - A >= 0, B < 0: A - (0 - |B|) = x -> A = x + (0 - |B|) -> A + |B| = x + ((0 - |B|) + |B|) -> x = A + |B|
//  - A < 0, B >= 0: (0 - |A|) - B = x -> 0 = x + (B + |A|) -> x = -(|A| + B)
//  - A < 0, B < 0: (0 - |A|) - (0 - |B|) = x -> (0 - |A|) = (0 - |B|) + x -> 0 = (0 - |B|) + x + |A|
//    -> 0 = ((0 - |B|) + |A|) + x -> x = 0 - (|A| + (0 - |B|)) 0 -> x = -(|A| - |B|)
func (z *Z) Subtract(arg *Z) *Z {
	if z.value >= 0 && arg.value >= 0 {
		return DefZ(&N{value: uint64(z.value)}, &N{value: uint64(arg.value)})
	} else if z.value >= 0 && arg.value < 0 {
		a := z
		b := &Z{value: -arg.value}
		c := a.Add(b)
		return &Z{value: int64(c.value)}
	} else if z.value < 0 && arg.value >= 0 {
		a := &Z{value: -z.value}
		b := arg
		c := a.Add(b)
		return &Z{value: -int64(c.value)}
	} else {
		a := &Z{value: -z.value}
		b := &Z{value: -arg.value}
		c := a.Subtract(b)
		return &Z{value: -int64(c.value)}
	}
}

func (z *Z) Divide(*Z) (*Z, error) {
	panic("implement me")
}

func (z *Z) Root(*Z) (*Z, error) {
	panic("implement me")
}

func (z *Z) Logarithm(*Z) (*Z, error) {
	panic("implement me")
}

var _ = fmt.Stringer(&Z{})
var _ = ZOperations(&Z{})
