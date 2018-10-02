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
	"testing"
)

func TestStringerInN(t *testing.T) {
	var n = N{value: 42}
	// we can access embedded interface implicitly - all that is required is "func (n *N) String() string {}"
	fmt.Printf("%s\n", n.String())
	// we can also access embedded interface directly, but it requires initialization
	// fmt.Printf("%s\n", n.Stringer.String())
	// this won't compile, since n is not an interface
	// display(n)
	// However this will compile!
	// see https://golang.org/doc/effective_go.html#pointers_vs_values
	// This rule arises because pointer methods can modify the receiver; invoking them on a value would cause the
	// method to receive a copy of the value, so any modifications would be discarded. The language therefore
	// disallows this mistake.
	display(&n)
	// this will compile as well
	display(fmt.Stringer(&n))
}

func TestAddingOne(t *testing.T) {
	var zero = ZERO
	fmt.Printf("zero: %s\n", zero.String())
	fmt.Printf("two: %s\n", zero.addOne().addOne().String())
}

func TestAdding(t *testing.T) {
	var zero = ZERO
	fmt.Printf("zero: %s\n", zero.String())
	fmt.Printf("two: %s\n", zero.addOne().addOne().String())

	_42 := NewN("42")
	display(_42)
	fmt.Printf("fourty two: %s\n", _42)
	fmt.Printf("fourty two plus eighteen: %s\n", _42.Add(NewN("18")))
	fmt.Printf("fourty two: %s\n", _42)
}

func TestMultiplying(t *testing.T) {
	var zero = ZERO
	fmt.Printf("zero: %s\n", zero.String())
	fmt.Printf("zero x three: %s\n", zero.Multiply(NewN("3")))
	fmt.Printf("56 x 3: %s\n", NewN("56").Multiply(NewN("3")))
	fmt.Printf("56 x 121: %s\n", NewN("56").Multiply(NewN("121")))
}

func display(s fmt.Stringer) {
	// fmt.Printf("%s\n", s.String())
	// fmt.Printf can handle fmt.Stringer interface
	fmt.Printf("%s\n", s)
	// glog.Info(s)
}
