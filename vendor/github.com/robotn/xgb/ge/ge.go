// Package ge is the X client API for the Generic Event Extension extension.
package ge

// This file is automatically generated from ge.xml. Edit at your peril!

import (
	"github.com/robotn/xgb"

	"github.com/robotn/xgb/xproto"
)

// Init must be called before using the Generic Event Extension extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 23, "Generic Event Extension").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named Generic Event Extension could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["Generic Event Extension"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["Generic Event Extension"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["Generic Event Extension"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["Generic Event Extension"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["Generic Event Extension"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Generic Event Extension"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'Generic Event Extension'. ge.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["Generic Event Extension"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'Generic Event Extension'. ge.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint16
	MinorVersion uint16
	// padding: 20 bytes
}

// Reply blocks and returns the reply data for a QueryVersion request.
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// queryVersionReply reads a byte slice into a QueryVersionReply value.
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = xgb.Get16(buf[b:])
	b += 2

	v.MinorVersion = xgb.Get16(buf[b:])
	b += 2

	b += 20 // padding

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["Generic Event Extension"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put16(buf[b:], ClientMajorVersion)
	b += 2

	xgb.Put16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}
