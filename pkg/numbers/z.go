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
	Power(*Z) (*Z, *Q, error)
	Subtract(*Z) *Z
	Divide(*Z) (*Z, *Q, error)
	Root(*Z) (*Z, *Q)
	Logarithm(*Z) (*Z, *Q)
}

// NewZ Creates new ℤ from string
func NewZ(v string) *Z {
	z, _ := strconv.Atoi(v)
	res := &Z{value: int64(z)}

	return res
}

// DefZ creates new ℤ as a result of subtracting two ℕs - definition of ℤ. Having ℤ defined and having
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
func (z *Z) Power(arg *Z) (*Z, *Q, error) {
	if z.value >= 0 && arg.value >= 0 {
		a := &N{value: uint64(z.value)}
		b := &N{value: uint64(arg.value)}
		c := a.Power(b)
		return &Z{value: int64(c.value)}, nil, nil
	} else if arg.value >= 0 {
		// reimplement from ℕ instead of delegate to ℕ
		res := NewZ("1")
		for i := int64(0); i < arg.value; i++ {
			res = res.Multiply(z)
		}
		return res, nil, nil
	} else {
		// possibly no solution in ℤ - delegating to Z.Divide which may switch to ℚ
		res, _, _ := z.Power(&Z{value: -arg.value})
		return (&Z{value: 1}).Divide(res)
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

// A / B:
//  - A >= 0, B >= 0: A / B (needs an implementaion, possible delegation to ℕ
//  - A >= 0, B < 0: -(A / |B|)
//  - A < 0, B >= 0: -(|A| / B)
//  - A < 0, B < 0: |A| / |B|
func (z *Z) Divide(arg *Z) (*Z, *Q, error) {
	if z.value >= 0 && arg.value >= 0 {
		a := &N{value: uint64(z.value)}
		b := &N{value: uint64(arg.value)}
		zres, qres, e := a.Divide(b)
		if zres != nil {
			return &Z{value: int64(zres.value)}, nil, nil
		}
		if qres != nil {
			return nil, qres, nil
		}
		return nil, nil, e
	} else if z.value >= 0 && arg.value < 0 {
		zres, qres, e := z.Divide(&Z{value: -arg.value})
		if zres != nil {
			zres.value = -zres.value
			return zres, nil, nil
		}
		if qres != nil {
			qres.a = -qres.a
			return nil, qres, nil
		}
		return nil, nil, e
	} else if z.value < 0 && arg.value >= 0 {
		zres, qres, e := (&Z{value: -z.value}).Divide(arg)
		if zres != nil {
			zres.value = -zres.value
			return zres, nil, nil
		}
		if qres != nil {
			qres.a = -qres.a
			return nil, qres, nil
		}
		return nil, nil, e
	} else {
		zres, qres, e := (&Z{value: -z.value}).Divide(&Z{value: -arg.value})
		if zres != nil {
			return zres, nil, nil
		}
		if qres != nil {
			return nil, qres, nil
		}
		return nil, nil, e
	}
	return nil, nil, nil
}

// A div B: division with remainder: A = QB + R and 0 <= R < |B|
// // C99 chooses the remainder with the same sign as the dividend A
//  - A >= 0, B >= 0: delegate to ℕ
//  - A >= 0, B < 0: -(A / |B|)
//  - A < 0, B >= 0: -(|A| / B)
//  - A < 0, B < 0: |A| / |B|
func (z *Z) DivideR(arg *Z) (*Z, *Z, error) {
	a := z.value
	b := arg.value
	negres := false
	negrem := false
	if a < 0 && b >= 0 || a >= 0 && b < 0 {
		negres = true
	}
	if a < 0 {
		negrem = true
		a = -a
	}
	if b < 0 {
		b = -b
	}
	av := &N{value: uint64(a)}
	bv := &N{value: uint64(b)}
	res, rem, e := av.DivideR(bv)
	if res != nil && rem != nil {
		if negrem {
			rem.value = -rem.value
		}
		if negres {
			res.value = -res.value
		}
		return &Z{value: int64(res.value)}, &Z{value: int64(rem.value)}, nil
	}
	return nil, nil, e
}

func (z *Z) Root(*Z) (*Z, *Q) {
	panic("implement me")
}

func (z *Z) Logarithm(*Z) (*Z, *Q) {
	panic("implement me")
}

var _ = fmt.Stringer(&Z{})
var _ = ZOperations(&Z{})
