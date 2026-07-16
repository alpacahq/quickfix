// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import "io"

func writeLoop(connection io.Writer, messageOut chan []byte, log Log) {
	var writeFailed bool
	for {
		msg, ok := <-messageOut
		if !ok {
			return
		}

		// After the first write failure the peer socket is dead. Keep draining
		// messageOut (without writing or re-logging) so the session goroutine
		// is not blocked on send, while closing the connection so readLoop
		// fails and triggers a clean disconnect that closes messageOut.
		if writeFailed {
			continue
		}

		if _, err := connection.Write(msg); err != nil {
			log.OnEvent(err.Error())
			writeFailed = true
			if closer, ok := connection.(io.Closer); ok {
				_ = closer.Close()
			}
		}
	}
}

func readLoop(parser *parser, msgIn chan fixIn, log Log) {
	defer close(msgIn)

	for {
		msg, err := parser.ReadMessage()
		if err != nil {
			log.OnEvent(err.Error())
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
