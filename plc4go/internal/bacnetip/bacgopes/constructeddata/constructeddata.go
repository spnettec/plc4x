/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package constructeddata

import (
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/comp"
	. "github.com/apache/plc4x/plc4go/internal/bacnetip/bacgopes/debugging"
)

var _debug = CreateDebugPrinter()

var _sequenceOfClasses map[any]struct{}
var _listOfClasses map[any]struct{}

func init() {
	_sequenceOfClasses = make(map[any]struct{})
	_listOfClasses = make(map[any]struct{})
}

// V2E accepts a function which takes an Arg and maps it to a ElementKlass
func V2E[T any](b func(arg Arg) (*T, error)) func(Args, KWArgs) (ElementKlass, error) {
	return func(args Args, kwArgs KWArgs) (ElementKlass, error) {
		var arg any
		if len(args) == 1 {
			arg = args[0]
		}
		r, err := b(arg)
		return any(r).(ElementKlass), err
	}
}

// Vs2E accepts a function which takes an Args and maps it to a ElementKlass
func Vs2E[T any](b func(args Args) (*T, error)) func(Args, KWArgs) (ElementKlass, error) {
	return func(args Args, kwArgs KWArgs) (ElementKlass, error) {
		r, err := b(args)
		return any(r).(ElementKlass), err
	}
}

// TODO: finish
func SequenceOfE[T any](b func(arg Arg) (*T, error)) func(Args, KWArgs) (ElementKlass, error) {
	panic("finish me")
}

func SequenceOfsE[T any](b func(args Args) (*T, error)) func(Args, KWArgs) (ElementKlass, error) {
	panic("finish me")
}

// TODO: finish // convert to kwArgs and check wtf we are doing here...
func ArrayOfE[T any](b func(arg Arg) (*T, error), fixedLength int, prototype any) func(Args, KWArgs) (ElementKlass, error) {
	panic("finish me")
}

// TODO: finish // convert to kwArgs and check wtf we are doing here...
func ArrayOfsE[T any](b func(args Args) (*T, error), fixedLength int, prototype any) func(Args, KWArgs) (ElementKlass, error) {
	panic("finish me")
}

// TODO: finish // convert to kwArgs and check wtf we are doing here...
func ListOfE[T any](b func(arg Arg) (*T, error)) func(Args) (*T, error) {
	panic("finish me")
}
