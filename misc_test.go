//
// Copyright © 2012-2019 Guy M. Allard
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package stompngo

import (
	"testing"
)

/*
	Test A zero Byte Message, a corner case.
*/
func TestMiscBytes0(t *testing.T) {
	for _, sp := range Protocols() {
		//tlg.Printf("TestMiscBytes0 protocol:%s\n", sp)
		// Write phase
		//tlg.Printf("TestMiscBytes0 WRITEPHASE\n")
		n, _ = openConn(t)
		ch := login_headers
		ch = headersProtocol(ch, sp)
		conn, e = Connect(n, ch)
		if e != nil {
			t.Fatalf("TestMiscBytes0 CONNECT expected nil, got %v\n", e)
		}
		//
		ms := "" // No data
		d := tdest("/queue/misc.zero.byte.msg." + sp)
		sh := Headers{HK_DESTINATION, d}
		e = conn.Send(sh, ms)
		if e != nil {
			t.Fatalf("TestMiscBytes0 Expected nil error, got [%v]\n", e)
		}
		//
		checkReceived(t, conn, false)
		e = conn.Disconnect(empty_headers)
		checkDisconnectError(t, e)
		_ = closeConn(t, n)

		// Read phase
		//tlg.Printf("TestMiscBytes0 READPHASE\n")
		n, _ = openConn(t)
		ch = login_headers
		ch = headersProtocol(ch, sp)
		conn, _ = Connect(n, ch)
		//
		sbh := sh.Add(HK_ID, d)
		sc, e = conn.Subscribe(sbh)
		if e != nil {
			t.Fatalf("TestMiscBytes0 Expected no subscribe error, got [%v]\n", e)
		}
		if sc == nil {
			t.Fatalf("TestMiscBytes0 Expected subscribe channel, got [nil]\n")
		}

		// Read MessageData
		var md MessageData
		select {
		case md = <-sc:
		case md = <-conn.MessageData:
			t.Fatalf("TestMiscBytes0 read channel error:  expected [nil], got: [%v]\n",
				md.Message.Command)
		}

		if md.Error != nil {
			t.Fatalf("TestMiscBytes0 Expected no message data error, got [%v]\n",
				md.Error)
		}

		// The real tests here
		if len(md.Message.Body) != 0 {
			t.Fatalf("TestMiscBytes0 Expected body length 0, got [%v]\n",
				len(md.Message.Body))
		}
		if string(md.Message.Body) != ms {
			t.Fatalf("TestMiscBytes0 Expected [%v], got [%v]\n",
				ms, string(md.Message.Body))
		}
		//
		//tlg.Printf("TestMiscBytes0 CLEANUP\n")
		checkReceived(t, conn, false)
		e = conn.Disconnect(empty_headers)
		checkDisconnectError(t, e)
		_ = closeConn(t, n)
	}
}

/*
	Test A One Byte Message, a corner case.
*/
func TestMiscBytes1(t *testing.T) {
	for _, sp := range Protocols() {
		//tlg.Printf("TestMiscBytes0 protocol:%s\n", sp)
		// Write phase
		//tlg.Printf("TestMiscBytes0 WRITEPHASE\n")
		n, _ = openConn(t)
		ch := login_headers
		ch = headersProtocol(ch, sp)
		conn, e = Connect(n, ch)
		if e != nil {
			t.Fatalf("TestMiscBytes1 CONNECT expected nil, got %v\n", e)
		}
		//
		ms := "Z" // One Byte
		d := tdest("/queue/misc.zero.byte.msg." + sp)
		sh := Headers{HK_DESTINATION, d}
		e = conn.Send(sh, ms)
		if e != nil {
			t.Fatalf("TestMiscBytes1 Expected nil error, got [%v]\n", e)
		}
		//
		checkReceived(t, conn, false)
		e = conn.Disconnect(empty_headers)
		checkDisconnectError(t, e)
		_ = closeConn(t, n)

		// Read phase
		//tlg.Printf("TestMiscBytes1 READPHASE\n")
		n, _ = openConn(t)
		ch = login_headers
		ch = headersProtocol(ch, sp)
		conn, _ = Connect(n, ch)
		//
		sbh := sh.Add(HK_ID, d)
		sc, e = conn.Subscribe(sbh)
		if e != nil {
			t.Fatalf("TestMiscBytes1 Expected no subscribe error, got [%v]\n", e)
		}
		if sc == nil {
			t.Fatalf("TestMiscBytes1 Expected subscribe channel, got [nil]\n")
		}

		// Read MessageData
		var md MessageData
		select {
		case md = <-sc:
		case md = <-conn.MessageData:
			t.Fatalf("TestMiscBytes1 read channel error:  expected [nil], got: [%v]\n",
				md.Message.Command)
		}

		if md.Error != nil {
			t.Fatalf("TestMiscBytes1 Expected no message data error, got [%v]\n",
				md.Error)
		}

		// The real tests here
		if len(md.Message.Body) != 1 {
			t.Fatalf("TestMiscBytes1 Expected body length 1, got [%v]\n",
				len(md.Message.Body))
		}
		if string(md.Message.Body) != ms {
			t.Fatalf("TestMiscBytes1 Expected [%v], got [%v]\n",
				ms, string(md.Message.Body))
		}
		//
		//tlg.Printf("TestMiscBytes1 CLEANUP\n")
		checkReceived(t, conn, false)
		e = conn.Disconnect(empty_headers)
		checkDisconnectError(t, e)
		_ = closeConn(t, n)
	}
}

/*
	Test nil Headers.
*/
func TestMiscNilHeaders(t *testing.T) {
	for _, _ = range Protocols() {
		n, _ = openConn(t)
		//
		_, e = Connect(n, nil)
		if e == nil {
			t.Fatalf("TestMiscNilHeaders Expected [%v], got [nil]\n",
				EHDRNIL)
		}
		if e != EHDRNIL {
			t.Fatalf("TestMiscNilHeaders Expected [%v], got [%v]\n",
				EHDRNIL, e)
		}
		//
		ch := check11(TEST_HEADERS)
		conn, _ = Connect(n, ch)
		//
		e = nil // reset
		e = conn.Abort(nil)
		if e == nil {
			t.Fatalf("TestMiscNilHeaders Abort Expected [%v], got [nil]\n",
				EHDRNIL)
		}
		//
		e = nil // reset
		e = conn.Ack(nil)
		if e == nil {
			t.Fatalf("TestMiscNilHeaders Ack Expected [%v], got [nil]\n",
				EHDRNIL)
		}
		//
		e = nil // reset
		e = conn.Begin(nil)
		if e == nil {
			t.Fatalf("TestMiscNilHeaders Begin Expected [%v], got [nil]\n",
				EHDRNIL)
		}
		//
		e = nil // reset
		e = conn.Commit(nil)
		if e == nil {
			t.Fatalf("TestMiscNilHeaders Commit Expected [%v], got [nil]\n",
				EHDRNIL)
		}
		//
		e = nil // reset
		e = conn.Disconnect(nil)
		if e == nil {
			t.Fatalf("TestMiscNilHeaders Disconnect Expected [%v], got [nil]\n",
				EHDRNIL)
		}
		//
		if conn.Protocol() > SPL_10 {
			e = nil // reset
			e = conn.Disconnect(nil)
			if e == nil {
				t.Fatalf("TestMiscNilHeaders Nack Expected [%v], got [nil]\n",
					EHDRNIL)
			}
		}
		//
		e = nil // reset
		e = conn.Send(nil, "")
		if e == nil {
			t.Fatalf("TestMiscNilHeaders Send Expected [%v], got [nil]\n",
				EHDRNIL)
		}
		//
	}
}

/*
Test max function.
*/
func TestMiscMax(t *testing.T) {
	for _, _ = range Protocols() {
		var l int64 = 1 // low
		var h int64 = 2 // high
		mr := max(l, h)
		if mr != 2 {
			t.Fatalf("TestMiscMax Expected [%v], got [%v]\n", h, mr)
		}
		mr = max(h, l)
		if mr != 2 {
			t.Fatalf("TestMiscMax Expected [%v], got [%v]\n", h, mr)
		}
	}
}

/*
Test hasValue function.
*/
func TestMiscHasValue(t *testing.T) {
	for _, _ = range Protocols() {
		sa := []string{"a", "b"}
		if !hasValue(sa, "a") {
			t.Fatalf("TestMiscHasValue Expected [true], got [false] for [%v]\n", "a")
		}
		if hasValue(sa, "z") {
			t.Fatalf("TestMiscHasValue Expected [false], got [true] for [%v]\n", "z")
		}
	}
}

/*
Test Uuid function.
*/
func TestMiscUuid(t *testing.T) {
	for _, _ = range Protocols() {
		id := Uuid()
		if id == "" {
			t.Fatalf("TestMiscUuid Expected a UUID, got empty string\n")
		}
		if len(id) != 36 {
			t.Fatalf("TestMiscUuid Expected a 36 character UUID, got length [%v]\n",
				len(id))
		}
	}
}

/*
	Test Bad Headers
*/
func TestMiscBadHeaders(t *testing.T) {
	for _, sp = range Protocols() {
		//
		n, _ = openConn(t)
		neh := Headers{"a", "b", "c"} // not even number header count
		conn, e = Connect(n, neh)

		// Connection should be nil (i.e. no connection)
		if e == nil {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [nil]\n", EHDRLEN)
		}
		if e != EHDRLEN {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [%v]\n", EHDRLEN, e)
		}
		//
		bvh := Headers{HK_HOST, "localhost", HK_ACCEPT_VERSION, "3.14159"}
		conn, e = Connect(n, bvh)
		if e == nil {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [nil]\n", EBADVERCLI)
		}
		if e != EBADVERCLI {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [%v]\n", EBADVERCLI, e)
		}

		//
		ch := login_headers
		ch = headersProtocol(ch, sp)
		//log.Printf("TestMiscBadHeaders Protocol %s, CONNECT Headers: %v\n", sp, ch)
		conn, e = Connect(n, ch)
		if e != nil {
			t.Fatalf("TestMiscBadHeaders CONNECT 2 expected nil, got %v connectresponse: %v\n",
				e, conn.ConnectResponse)
		}

		// Connection should not be nil (i.e. good connection)
		_, e = conn.Subscribe(neh)
		if e == nil {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [nil]\n", EHDRLEN)
		}
		if e != EHDRLEN {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [%v]\n", EHDRLEN, e)
		}
		//
		e = conn.Unsubscribe(neh)
		if e == nil {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [nil]\n", EHDRLEN)
		}
		if e != EHDRLEN {
			t.Fatalf("TestMiscBadHeaders Expected [%v], got [%v]\n", EHDRLEN, e)
		}
		//
		if conn != nil && conn.Connected() {
			e = conn.Disconnect(empty_headers)
			checkDisconnectError(t, e)
		}
		_ = closeConn(t, n)
	}
}
