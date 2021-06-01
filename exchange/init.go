package exchange

import (
	c "../configure"
)

type Header struct {
	Flag       byte   `json: flag`
	DestIPAddr c.IPv4 `json: destipaddr`
}

type Message struct {
	MHeader Header `json: header`
	MData   []byte `json: message`
}
