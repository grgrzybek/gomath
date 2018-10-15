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

/*
 * https://unicode-table.com/en/blocks/letterlike-symbols/
 * Natural numbers: ℕ, u2115
 * Integer numbers: ℤ, u2124
 * Rational numbers: ℚ, u211A
 * Real numbers: ℝ, u211D
 * Complex numbers: ℂ, u2102
 */

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	// this is the only Integer number we know initially
	ZERO = N{value: 0}
)

// Natural numbers including ZERO
// We suppose that we already know what integers are, what zero is, and what it means to increase a number by one unit.
//
// Basic Rules for addition, multiplication and raising to a power (steming from definition):
//  - (a) a+b = b+a
//  - (d) a+(b+c) = (a+b)+c
//  - (b) a*b = b*a
//  - (c) a*(b+c) = a*b + a*c
//  - (e) (a*b)*c = a*(b*c)
//  - (f) (a*b)^c = a^c * b^c
//  - (g) a^b * a^c = a^(b+c)
//  - (h) (a^b)^c = a^(b*c)
//  - (i) a+0 = a
//  - (j) a*1 = a
//  - (k) a^1 = a
type N struct {
	value uint64

	// mark ℕ as "implementing fmt.Stringer" using anonymous field / embedded type
	// https://golang.org/doc/effective_go.html#embedding
	// When we embed a type, the methods of that type become methods of the outer type, but when
	// they are invoked the receiver of the method is the inner type, not the outer one.
	fmt.Stringer
	// NOperations
}

type NOperations interface {
	Add(*N) *N
	Multiply(*N) *N
	Power(*N) *N
	Subtract(*N) (*N, *Z)
	Divide(*N) (*N, *Q, error)
	Root(*N) (*N, error)
	Logarithm(*N) (*N, error)
}

// NewN Creates new ℕ from string
func NewN(v string) *N {
	value, _ := strconv.Atoi(v)
	res := &ZERO
	for i := 0; i < value; i++ {
		res = res.addOne()
	}
	return res
}

// Override promoted methods, by default it's just n.String() -> n.Stringer.String() and is causing NPE
func (n *N) String() string {
	return fmt.Sprintf("%d", n.value)
}

// "Addition": Start with integer A and increase it by 1, B times to get "A + B"
func (n *N) Add(arg *N) *N {
	res := n
	for i := uint64(0); i < arg.value; i++ {
		res = res.addOne()
	}
	return res
}

// "Multiplication": (requires definition of "Addition") Start with ZERO and add A to it B times to get "A * B"
func (n *N) Multiply(arg *N) *N {
	res := &ZERO
	for i := uint64(0); i < arg.value; i++ {
		res = res.Add(n)
	}
	return res
}

// "Raising to power": (requires definition of "Multiplication") Start with ONE and multiply it by A, B times
// to get "A^B"
func (n *N) Power(arg *N) *N {
	res := NewN("1")
	for i := uint64(0); i < arg.value; i++ {
		res = res.Multiply(n)
	}
	return res
}

// "Subtraction": Assuming A and C are given, we want to find B that "A + B = C". Then B is defined as "C - A"
func (n *N) Subtract(arg *N) (*N, *Z) {
	if n.value < arg.value {
		// no solution in ℕ, switching to ℤ and returning ℤ by definition
		return nil, DefZ(n, arg)
	}

	res := &ZERO
	for {
		if arg.Add(res).value == n.value {
			return res, nil
		}
		res = res.addOne()
	}

	return res, nil
}

// "Division": Assuming A and C are given, we want to find B that "A * B = C". Then B is defined as "C / A"
func (n *N) Divide(arg *N) (*N, *Q, error) {
	if arg.value == 0 {
		return nil, nil, errors.New("can't divide by ZERO")
	}

	res := &ZERO
	for {
		if arg.Multiply(res).value == n.value {
			return res, nil, nil
		}
		if arg.Multiply(res).value > n.value {
			break
		}
		res = res.addOne()
	}

	// immediately delegate to ℚ
	return nil, DefQ(&Z{value: int64(n.value)}, &Z{value: int64(arg.value)}), nil
}

// Division with Remainder: A = QB + R and 0 <= R < |B|
func (n *N) DivideR(arg *N) (*N, *N, error) {
	if arg.value == 0 {
		return nil, nil, errors.New("can't divide by ZERO")
	}

	res := &ZERO
	for {
		if arg.Multiply(res).value == n.value {
			return res, &ZERO, nil
		}
		if arg.Multiply(res).value > n.value {
			break
		}
		res = res.addOne()
	}

	res1, _ := res.Subtract(NewN("1"))
	rem, _ := n.Subtract(res1.Multiply(arg))
	return res1, rem, nil
}

// "Root": Assuming A and C are given, we want to find B that "B ^ A = C", Then B is defined as "Ath√C".
// n is root's degree, arg is the argument, result is a number we need to raise to power n, to get arg
func (n *N) Root(arg *N) (*N, error) {
	if n.value == 0 {
		return nil, errors.New("can't take ZEROth root")
	}

	res := NewN("0")
	for {
		if res.Power(n).value == arg.value {
			return res, nil
		}
		if res.Power(n).value > arg.value {
			break
		}
		res = res.addOne()
	}

	return nil, errors.New("no solution in \u2115")
}

// "Logarithm": Assuming A and C are given, we want to find B that "A ^ B = C", Then B is defined as "log A (C)".
// n is base, arg is the argument, result is a power to which we need to raise n, to get arg
func (n *N) Logarithm(arg *N) (*N, error) {
	if arg.value == 0 {
		return nil, errors.New("can't take logarithm from ZERO")
	}
	if n.value == 0 {
		return nil, errors.New("can't take logarithm with base ZERO")
	}
	if n.value == 1 {
		return nil, errors.New("can't take logarithm with base ONE")
	}

	res := NewN("0")
	for {
		if n.Power(res).value == arg.value {
			return res, nil
		}
		if n.Power(res).value > arg.value {
			break
		}
		res = res.addOne()
	}

	return nil, errors.New("no solution in \u2115")
}

// and we only know how to "add 1" - find "next" number
func (n *N) addOne() *N {
	res := &N{value: n.value}
	res.value++
	return res
}

// validation of interface implementation
var _ = fmt.Stringer(&N{})
var _ = NOperations(&N{})
