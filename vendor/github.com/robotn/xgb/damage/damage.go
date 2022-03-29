// Package damage is the X client API for the DAMAGE extension.
package damage

// This file is automatically generated from damage.xml. Edit at your peril!

import (
	"github.com/robotn/xgb"

	"github.com/robotn/xgb/xfixes"
	"github.com/robotn/xgb/xproto"
)

// Init must be called before using the DAMAGE extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 6, "DAMAGE").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named DAMAGE could be found on on the server.")
	}

	c.ExtLock.Lock()
	c.Extensions["DAMAGE"] = reply.MajorOpcode
	c.ExtLock.Unlock()
	for evNum, fun := range xgb.NewExtEventFuncs["DAMAGE"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["DAMAGE"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	return nil
}

func init() {
	xgb.NewExtEventFuncs["DAMAGE"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["DAMAGE"] = make(map[int]xgb.NewErrorFun)
}

// BadBadDamage is the error number for a BadBadDamage.
const BadBadDamage = 0

type BadDamageError struct {
	Sequence uint16
	NiceName string
}

// BadDamageErrorNew constructs a BadDamageError value that implements xgb.Error from a byte slice.
func BadDamageErrorNew(buf []byte) xgb.Error {
	v := BadDamageError{}
	v.NiceName = "BadDamage"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	return v
}

// SequenceId returns the sequence id attached to the BadBadDamage error.
// This is mostly used internally.
func (err BadDamageError) SequenceId() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadBadDamage error. If no bad value exists, 0 is returned.
func (err BadDamageError) BadId() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadDamage error.

func (err BadDamageError) Error() string {
	fieldVals := make([]string, 0, 0)
	fieldVals = append(fieldVals, "NiceName: "+err.NiceName)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", err.Sequence))
	return "BadBadDamage {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtErrorFuncs["DAMAGE"][0] = BadDamageErrorNew
}

type Damage uint32

func NewDamageId(c *xgb.Conn) (Damage, error) {
	id, err := c.NewId()
	if err != nil {
		return 0, err
	}
	return Damage(id), nil
}

// Notify is the event number for a NotifyEvent.
const Notify = 0

type NotifyEvent struct {
	Sequence  uint16
	Level     byte
	Drawable  xproto.Drawable
	Damage    Damage
	Timestamp xproto.Timestamp
	Area      xproto.Rectangle
	Geometry  xproto.Rectangle
}

// NotifyEventNew constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func NotifyEventNew(buf []byte) xgb.Event {
	v := NotifyEvent{}
	b := 1 // don't read event number

	v.Level = buf[b]
	b += 1

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Drawable = xproto.Drawable(xgb.Get32(buf[b:]))
	b += 4

	v.Damage = Damage(xgb.Get32(buf[b:]))
	b += 4

	v.Timestamp = xproto.Timestamp(xgb.Get32(buf[b:]))
	b += 4

	v.Area = xproto.Rectangle{}
	b += xproto.RectangleRead(buf[b:], &v.Area)

	v.Geometry = xproto.Rectangle{}
	b += xproto.RectangleRead(buf[b:], &v.Geometry)

	return v
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Level
	b += 1

	b += 2 // skip sequence number

	xgb.Put32(buf[b:], uint32(v.Drawable))
	b += 4

	xgb.Put32(buf[b:], uint32(v.Damage))
	b += 4

	xgb.Put32(buf[b:], uint32(v.Timestamp))
	b += 4

	{
		structBytes := v.Area.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.Geometry.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf
}

// SequenceId returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v NotifyEvent) SequenceId() uint16 {
	return v.Sequence
}

// String is a rudimentary string representation of NotifyEvent.
func (v NotifyEvent) String() string {
	fieldVals := make([]string, 0, 6)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("Level: %d", v.Level))
	fieldVals = append(fieldVals, xgb.Sprintf("Drawable: %d", v.Drawable))
	fieldVals = append(fieldVals, xgb.Sprintf("Damage: %d", v.Damage))
	fieldVals = append(fieldVals, xgb.Sprintf("Timestamp: %d", v.Timestamp))
	return "Notify {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtEventFuncs["DAMAGE"][0] = NotifyEventNew
}

const (
	ReportLevelRawRectangles   = 0
	ReportLevelDeltaRectangles = 1
	ReportLevelBoundingBox     = 2
	ReportLevelNonEmpty        = 3
)

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

// AddCookie is a cookie used only for Add requests.
type AddCookie struct {
	*xgb.Cookie
}

// Add sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Add(c *xgb.Conn, Drawable xproto.Drawable, Region xfixes.Region) AddCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Add' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(addRequest(c, Drawable, Region), cookie)
	return AddCookie{cookie}
}

// AddChecked sends a checked request.
// If an error occurs, it can be retrieved using AddCookie.Check()
func AddChecked(c *xgb.Conn, Drawable xproto.Drawable, Region xfixes.Region) AddCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Add' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(addRequest(c, Drawable, Region), cookie)
	return AddCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook AddCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Add
// addRequest writes a Add request to a byte slice.
func addRequest(c *xgb.Conn, Drawable xproto.Drawable, Region xfixes.Region) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["DAMAGE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], uint32(Region))
	b += 4

	return buf
}

// CreateCookie is a cookie used only for Create requests.
type CreateCookie struct {
	*xgb.Cookie
}

// Create sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Create(c *xgb.Conn, Damage Damage, Drawable xproto.Drawable, Level byte) CreateCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Create' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(createRequest(c, Damage, Drawable, Level), cookie)
	return CreateCookie{cookie}
}

// CreateChecked sends a checked request.
// If an error occurs, it can be retrieved using CreateCookie.Check()
func CreateChecked(c *xgb.Conn, Damage Damage, Drawable xproto.Drawable, Level byte) CreateCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Create' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(createRequest(c, Damage, Drawable, Level), cookie)
	return CreateCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook CreateCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Create
// createRequest writes a Create request to a byte slice.
func createRequest(c *xgb.Conn, Damage Damage, Drawable xproto.Drawable, Level byte) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["DAMAGE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Damage))
	b += 4

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	buf[b] = Level
	b += 1

	b += 3 // padding

	return buf
}

// DestroyCookie is a cookie used only for Destroy requests.
type DestroyCookie struct {
	*xgb.Cookie
}

// Destroy sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Destroy(c *xgb.Conn, Damage Damage) DestroyCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Destroy' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroyRequest(c, Damage), cookie)
	return DestroyCookie{cookie}
}

// DestroyChecked sends a checked request.
// If an error occurs, it can be retrieved using DestroyCookie.Check()
func DestroyChecked(c *xgb.Conn, Damage Damage) DestroyCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Destroy' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroyRequest(c, Damage), cookie)
	return DestroyCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook DestroyCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Destroy
// destroyRequest writes a Destroy request to a byte slice.
func destroyRequest(c *xgb.Conn, Damage Damage) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["DAMAGE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Damage))
	b += 4

	return buf
}

// QueryVersionCookie is a cookie used only for QueryVersion requests.
type QueryVersionCookie struct {
	*xgb.Cookie
}

// QueryVersion sends a checked request.
// If an error occurs, it will be returned with the reply by calling QueryVersionCookie.Reply()
func QueryVersion(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, ClientMajorVersion, ClientMinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// QueryVersionUnchecked sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func QueryVersionUnchecked(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) QueryVersionCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'QueryVersion' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
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
	MajorVersion uint32
	MinorVersion uint32
	// padding: 16 bytes
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

	v.MajorVersion = xgb.Get32(buf[b:])
	b += 4

	v.MinorVersion = xgb.Get32(buf[b:])
	b += 4

	b += 16 // padding

	return v
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(c *xgb.Conn, ClientMajorVersion uint32, ClientMinorVersion uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["DAMAGE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], ClientMajorVersion)
	b += 4

	xgb.Put32(buf[b:], ClientMinorVersion)
	b += 4

	return buf
}

// SubtractCookie is a cookie used only for Subtract requests.
type SubtractCookie struct {
	*xgb.Cookie
}

// Subtract sends an unchecked request.
// If an error occurs, it can only be retrieved using xgb.WaitForEvent or xgb.PollForEvent.
func Subtract(c *xgb.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) SubtractCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Subtract' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(false, false)
	c.NewRequest(subtractRequest(c, Damage, Repair, Parts), cookie)
	return SubtractCookie{cookie}
}

// SubtractChecked sends a checked request.
// If an error occurs, it can be retrieved using SubtractCookie.Check()
func SubtractChecked(c *xgb.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) SubtractCookie {
	c.ExtLock.RLock()
	defer c.ExtLock.RUnlock()
	if _, ok := c.Extensions["DAMAGE"]; !ok {
		panic("Cannot issue request 'Subtract' using the uninitialized extension 'DAMAGE'. damage.Init(connObj) must be called first.")
	}
	cookie := c.NewCookie(true, false)
	c.NewRequest(subtractRequest(c, Damage, Repair, Parts), cookie)
	return SubtractCookie{cookie}
}

// Check returns an error if one occurred for checked requests that are not expecting a reply.
// This cannot be called for requests expecting a reply, nor for unchecked requests.
func (cook SubtractCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for Subtract
// subtractRequest writes a Subtract request to a byte slice.
func subtractRequest(c *xgb.Conn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	c.ExtLock.RLock()
	buf[b] = c.Extensions["DAMAGE"]
	c.ExtLock.RUnlock()
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Damage))
	b += 4

	xgb.Put32(buf[b:], uint32(Repair))
	b += 4

	xgb.Put32(buf[b:], uint32(Parts))
	b += 4

	return buf
}
