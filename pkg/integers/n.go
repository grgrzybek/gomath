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
package integers

import (
	"fmt"
	"strconv"
)

var (
	// this is the only Integer number we know initially
	ZERO = N{value: 0}
)

// Natural numbers including ZERO
// We suppose that we already know what integers are, what zero is, and what it means to increase a number by one unit.
type N struct {
	value uint64

	// mark N as "implementing fmt.Stringer" using anonymous field / embedded type
	// https://golang.org/doc/effective_go.html#embedding
	// When we embed a type, the methods of that type become methods of the outer type, but when
	// they are invoked the receiver of the method is the inner type, not the outer one.
	fmt.Stringer
	NOperations
}

type NOperations interface {
	Add(*N) *N
	Multiply(*N) *N
}

// Override promoted methods, by default it's just n.String() -> n.Stringer.String() and is causing NPE
func (n *N) String() string {
	return fmt.Sprintf("%d", n.value)
}

// If we start with a certain number a, an integer, and we count successively one unit b times,
// the number we arrive at we call a+b, and that defines addition of integers.
func (n *N) Add(arg *N) *N {
	res := n
	for i := uint64(0); i < arg.value; i++ {
		res = res.addOne()
	}
	return res
}

// Once we have defined addition, then we can consider this: if we start with nothing and add a to it, b times in
// succession, we call the result multiplication of integers; we call it b times a.
func (n *N) Multiply(arg *N) *N {
	res := &ZERO
	for i := uint64(0); i < arg.value; i++ {
		res = res.Add(n)
	}
	return res
}

// and we only know how to "add 1" - find "next" number
func (n *N) addOne() *N {
	res := &N{value: n.value}
	res.value++
	return res
}

// NewN Creates new N from string
func NewN(v string) *N {
	value, _ := strconv.Atoi(v)
	res := &ZERO
	for i := 0; i < value; i++ {
		res = res.addOne()
	}
	return res
}

// validation of interface implementation
var _ = fmt.Stringer(&N{})
var _ = NOperations(&N{})

// Integer numbers: include Natural numbers, ZERO and Additive Inverses of Natural numbers
// https://en.wikipedia.org/wiki/Additive_inverse
type Z struct {
	value int64

	fmt.Stringer
}

func (z *Z) String() string {
	return fmt.Sprintf("%d", z.value)
}
