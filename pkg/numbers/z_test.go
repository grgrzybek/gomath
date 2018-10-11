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

func TestSubtractN(t *testing.T) {
	if n, z := NewN("42").Subtract(NewN("43")); z == nil {
		fmt.Printf("42-43: %s\n", n)
	} else {
		fmt.Printf("42-43: %s\n", z)
	}
	if n, z := NewN("6").Subtract(NewN("1211")); z == nil {
		fmt.Printf("6-1211: %s\n", n)
	} else {
		fmt.Printf("6-1211: %s\n", z)
	}
}

func TestAddZ(t *testing.T) {
	fmt.Printf("1 + -3: %s\n", NewZ("1").Add(NewZ("-3")))
	fmt.Printf("1 + 3: %s\n", NewZ("1").Add(NewZ("3")))
	fmt.Printf("-1 + 3: %s\n", NewZ("-1").Add(NewZ("3")))
	fmt.Printf("-10 + 3: %s\n", NewZ("-10").Add(NewZ("3")))
	fmt.Printf("-10 + -3: %s\n", NewZ("-10").Add(NewZ("-3")))
}

func TestMultiplyZ(t *testing.T) {
	fmt.Printf("1 * -3: %s\n", NewZ("1").Multiply(NewZ("-3")))
	fmt.Printf("1 * 3: %s\n", NewZ("1").Multiply(NewZ("3")))
	fmt.Printf("-1 * 3: %s\n", NewZ("-1").Multiply(NewZ("3")))
	fmt.Printf("-10 * 3: %s\n", NewZ("-10").Multiply(NewZ("3")))
	fmt.Printf("-10 * -3: %s\n", NewZ("-10").Multiply(NewZ("-3")))
}

func TestPowerZ(t *testing.T) {
	if z, e := NewZ("0").Power(NewZ("0")); e == nil {
		fmt.Printf("0^0: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("1").Power(NewZ("1")); e == nil {
		fmt.Printf("1^1: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("0").Power(NewZ("3")); e == nil {
		fmt.Printf("0^3: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("3").Power(NewZ("0")); e == nil {
		fmt.Printf("3^0: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("2").Power(NewZ("16")); e == nil {
		fmt.Printf("2^16: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("10").Power(NewZ("6")); e == nil {
		fmt.Printf("10^6: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}

	if z, e := NewZ("-1").Power(NewZ("1")); e == nil {
		fmt.Printf("-1^1: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("-3").Power(NewZ("0")); e == nil {
		fmt.Printf("-3^0: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("-2").Power(NewZ("16")); e == nil {
		fmt.Printf("-2^16: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("-10").Power(NewZ("6")); e == nil {
		fmt.Printf("-10^6: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}
	if z, e := NewZ("-10").Power(NewZ("5")); e == nil {
		fmt.Printf("-10^5: %s\n", z)
	} else {
		_ = fmt.Errorf("%s\n", e)
	}

	if z, e := NewZ("1").Power(NewZ("-1")); e == nil {
		fmt.Printf("1^-1: %s\n", z)
	} else {
		fmt.Printf("1^-1: %s\n", fmt.Errorf("%s", e))
	}
	if z, e := NewZ("-1").Power(NewZ("-1")); e == nil {
		fmt.Printf("-1^-1: %s\n", z)
	} else {
		fmt.Printf("-1^-1: %s\n", fmt.Errorf("%s", e))
	}
}

func TestSubtractZ(t *testing.T) {
	fmt.Printf("1 - -3: %s\n", NewZ("1").Subtract(NewZ("-3")))
	fmt.Printf("1 - 3: %s\n", NewZ("1").Subtract(NewZ("3")))
	fmt.Printf("-1 - 3: %s\n", NewZ("-1").Subtract(NewZ("3")))
	fmt.Printf("-10 - 3: %s\n", NewZ("-10").Subtract(NewZ("3")))
	fmt.Printf("-10 - -3: %s\n", NewZ("-10").Subtract(NewZ("-3")))
}
