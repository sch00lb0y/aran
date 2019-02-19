// Copyright 2019 sch00lb0y.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.
package aran

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/y"
)

func TestDB(t *testing.T) {
	d, err := New(DefaultOptions())
	if err != nil {
		t.Fatalf("db is expected to open but got error %s", err.Error())
	}
	d.Set([]byte("hello"), []byte("schoolboy"))
	d.Close()
	d, err = New(DefaultOptions())
	if err != nil {
		t.Fatalf("db is expected to open but got error %s", err.Error())
	}
	val, exist := d.Get([]byte("hello"))
	if !exist {
		t.Fatalf("unable to retrive data")
	}
	if bytes.Compare(val, []byte("schoolboy")) != 0 {
		t.Fatalf("value is not same expected schoolboy but got %s", string(val))
	}
	d.Close()
}

func TestCloser(t *testing.T) {
	closer := y.NewCloser(1)
	go func() {
	loop:
		for {
			select {
			case <-closer.HasBeenClosed():
				fmt.Println("breaking")
				break loop
			}
		}
		closer.Done()
	}()
	closer.SignalAndWait()
}
