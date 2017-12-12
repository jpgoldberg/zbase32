// package zbase32 -- A base32 encoding meant for human consumption
// This is based off of http://philzimmermann.com/docs/human-oriented-base-32-encoding.txt
//
// This just defines a base32.NewEncoding using the zbase32 alphabet
// So all of the methods available to base32.Encoding are available
// to a ZEncoding
package zbase32

import (
	base32 "encoding/base32"
)

const encodeZooko = "ybndrfg8ejkmcpqxot1uwisza345h769"

// ZEncoding is the zbase32 encoding. This is very non-standard
var ZEncoding = base32.NewEncoding(encodeZooko).WithPadding(base32.NoPadding)
