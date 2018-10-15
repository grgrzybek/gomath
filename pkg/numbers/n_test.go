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

func TestPower(t *testing.T) {
	fmt.Printf("0^0: %s\n", ZERO.Power(&ZERO))
	fmt.Printf("1^1: %s\n", NewN("1").Power(NewN("1")))
	fmt.Printf("0^3: %s\n", ZERO.Power(NewN("3")))
	fmt.Printf("3^0: %s\n", NewN("3").Power(&ZERO))
	fmt.Printf("2^16: %s\n", NewN("2").Power(NewN("16")))
	fmt.Printf("10^6: %s\n", NewN("10").Power(NewN("6")))
}

func TestSubtract(t *testing.T) {
	if n, e := NewN("145").Subtract(NewN("22")); e == nil {
		fmt.Printf("145-22: %s\n", n)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if n, e := NewN("42").Subtract(NewN("42")); e == nil {
		fmt.Printf("42-42: %s\n", n)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if n, e := NewN("42").Subtract(NewN("43")); e == nil {
		fmt.Printf("42-43: %s\n", n)
	} else {
		fmt.Printf("42-43: %s\n", fmt.Errorf("%s", e))
	}
}

func TestDivide(t *testing.T) {
	if n, _, e := NewN("140").Divide(NewN("10")); e == nil {
		fmt.Printf("140/10: %s\n", n)
	} else {
		fmt.Printf("140/10: %s\n", fmt.Errorf("%s", e))
	}
	if n, _, e := NewN("140").Divide(NewN("11")); e == nil {
		fmt.Printf("140/11: %s\n", n)
	} else {
		fmt.Printf("140/11: %s\n", fmt.Errorf("%s", e))
	}
	if n, _, e := NewN("42").Divide(NewN("42")); e == nil {
		fmt.Printf("42/42: %s\n", n)
	} else {
		fmt.Printf("42/42: %s\n", fmt.Errorf("%s", e))
	}
	if n, _, e := NewN("0").Divide(NewN("42")); e == nil {
		fmt.Printf("0/42: %s\n", n)
	} else {
		fmt.Printf("0/42: %s\n", fmt.Errorf("%s", e))
	}
	if n, _, e := NewN("1").Divide(NewN("0")); e == nil {
		fmt.Printf("1/0: %s\n", n)
	} else {
		fmt.Printf("1/0: %s\n", fmt.Errorf("%s", e))
	}
}

func TestDivideR(t *testing.T) {
	divideR("140", "10")
	divideR("140", "11")
	divideR("42", "42")
	divideR("0", "42")
	divideR("1", "0")
	divideR("25", "10")
	divideR("10", "3")
	divideR("3", "10")
}

func TestRoot(t *testing.T) {
	if n, e := NewN("1").Root(NewN("1")); e == nil {
		fmt.Printf("1√1: %s\n", n)
	} else {
		fmt.Printf("1√1: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("1").Root(NewN("4")); e == nil {
		fmt.Printf("1√4: %s\n", n)
	} else {
		fmt.Printf("1√4: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("4").Root(NewN("1")); e == nil {
		fmt.Printf("4√1: %s\n", n)
	} else {
		fmt.Printf("4√1: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("2").Root(NewN("9")); e == nil {
		fmt.Printf("2√9: %s\n", n)
	} else {
		fmt.Printf("2√9: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("4").Root(NewN("16")); e == nil {
		fmt.Printf("4√16: %s\n", n)
	} else {
		fmt.Printf("4√16: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("16").Root(NewN("65536")); e == nil {
		fmt.Printf("16√65536: %s\n", n)
	} else {
		fmt.Printf("16√65536: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("2").Root(NewN("65536")); e == nil {
		fmt.Printf("2√65536: %s\n", n)
	} else {
		fmt.Printf("2√65536: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("3").Root(NewN("0")); e == nil {
		fmt.Printf("3√0: %s\n", n)
	} else {
		fmt.Printf("3√0: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("0").Root(NewN("3")); e == nil {
		fmt.Printf("0√3: %s\n", n)
	} else {
		fmt.Printf("0√3: %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("0").Root(NewN("0")); e == nil {
		fmt.Printf("0√0: %s\n", n)
	} else {
		fmt.Printf("0√0: %s\n", fmt.Errorf("%s", e))
	}
}

func TestLogarithm(t *testing.T) {
	if n, e := NewN("1").Logarithm(NewN("1")); e == nil {
		fmt.Printf("log 1 (1): %s\n", n)
	} else {
		fmt.Printf("log 1 (1): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("1").Logarithm(NewN("4")); e == nil {
		fmt.Printf("log 1 (4): %s\n", n)
	} else {
		fmt.Printf("log 1 (4): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("4").Logarithm(NewN("1")); e == nil {
		fmt.Printf("log 4 (1): %s\n", n)
	} else {
		fmt.Printf("log 4 (1): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("2").Logarithm(NewN("9")); e == nil {
		fmt.Printf("log 2 (9): %s\n", n)
	} else {
		fmt.Printf("log 2 (9): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("3").Logarithm(NewN("9")); e == nil {
		fmt.Printf("log 3 (9): %s\n", n)
	} else {
		fmt.Printf("log 3 (9): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("4").Logarithm(NewN("16")); e == nil {
		fmt.Printf("log 4 (16): %s\n", n)
	} else {
		fmt.Printf("log 4 (16): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("16").Logarithm(NewN("65536")); e == nil {
		fmt.Printf("log 16 (65536): %s\n", n)
	} else {
		fmt.Printf("log 16 (65536): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("2").Logarithm(NewN("65536")); e == nil {
		fmt.Printf("log 2 (65536): %s\n", n)
	} else {
		fmt.Printf("log 2 (65536): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("3").Logarithm(NewN("0")); e == nil {
		fmt.Printf("log 3 (0): %s\n", n)
	} else {
		fmt.Printf("log 3 (0): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("0").Logarithm(NewN("3")); e == nil {
		fmt.Printf("log 0 (3): %s\n", n)
	} else {
		fmt.Printf("log 0 (3): %s\n", fmt.Errorf("%s", e))
	}
	if n, e := NewN("0").Logarithm(NewN("0")); e == nil {
		fmt.Printf("log 0 (0): %s\n", n)
	} else {
		fmt.Printf("log 0 (0): %s\n", fmt.Errorf("%s", e))
	}
}

func display(s fmt.Stringer) {
	// fmt.Printf("%s\n", s.String())
	// fmt.Printf can handle fmt.Stringer interface
	fmt.Printf("%s\n", s)
	// glog.Info(s)
}

func divideR(a string, b string) {
	z, r, e := NewN(a).DivideR(NewN(b))
	if z != nil {
		fmt.Printf("%s/%s: %s r %s\n", a, b, z, r)
	} else {
		fmt.Printf("%s/%s: %s\n", a, b, fmt.Errorf("%s", e))
	}
}
