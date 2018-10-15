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
)

// Rational numbers ℚ - needed to define negative power or division in ℤ
type Q struct {
	a int64
	b int64

	fmt.Stringer
}

// NewQ Creates new ℚ from string
func NewQ(v string) *Q {
	var a, b int64
	_, e := fmt.Sscanf(v, "%d/%d", &a, &b)
	if e == nil {
		q, _ := (&Q{a: int64(a), b: int64(b)}).GCD()
		return q
	}

	panic(fmt.Errorf("%s", e))
}

// DefQ creates new ℚ as a result of dividing two ℤs - definition of ℚ.
//
// if A < B, B is decreased (using division by A) to 1 and resulting (A / B) is called "rational number"
func DefQ(a *Z, b *Z) *Q {
	return &Q{a: a.value, b: b.value} // definition - by division of integer numbers
}

type QOperations interface {
}

// Trim tries to minimize nominator and denominator
// A = Q0 * B + R0
// B = Q1 * R0 + R1
// R0 = Q2 * R1 + R2
// R1 = Q3 * R2 + R3
// ...
func (q *Q) GCD() (*Q, error) {
	r0 := q.a
	r1 := q.b
	if r0 < 0 {
		r0 = -r0
	}
	if r1 < 0 {
		r1 = -r1
	}
	var gcd int64 = 1
	for {
		if r1 == 0 {
			gcd = 1
		}
		q, r, e := (&Z{value: r0}).DivideR(&Z{value: r1})
		if q != nil && r != nil {
			if r.value == 0 {
				gcd = r1
				break
			}
			r0 = r1
			r1 = r.value
		} else {
			return nil, e
		}
	}

	gcdz := &Z{value: gcd}
	az, _, _ := (&Z{value: q.a}).Divide(gcdz)
	bz, _, _ := (&Z{value: q.b}).Divide(gcdz)
	if az.value < 0 && bz.value < 0 {
		az.value = -az.value
		bz.value = -bz.value
	}
	return &Q{a: az.value, b: bz.value}, nil
}

func (q *Q) String() string {
	_q, _ := q.GCD()
	return fmt.Sprintf("%d/%d", _q.a, _q.b)
}

var _ = fmt.Stringer(&Q{})
var _ = QOperations(&Q{})
