// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.26/unstable/relative-pointer/relative-pointer-unstable-v1.xml
//
// relative_pointer_unstable_v1 Protocol Copyright:
//
// Copyright © 2014      Jonas Ådahl
// Copyright © 2015      Red Hat Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice (including the next
// paragraph) shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package relative_pointer

import "github.com/rajveermalviya/go-wayland/wayland/client"

// RelativePointerManager : get relative pointer objects
//
// A global interface used for getting the relative pointer object for a
// given pointer.
type RelativePointerManager struct {
	client.BaseProxy
}

// NewRelativePointerManager : get relative pointer objects
//
// A global interface used for getting the relative pointer object for a
// given pointer.
func NewRelativePointerManager(ctx *client.Context) *RelativePointerManager {
	zwpRelativePointerManagerV1 := &RelativePointerManager{}
	ctx.Register(zwpRelativePointerManagerV1)
	return zwpRelativePointerManagerV1
}

// Destroy : destroy the relative pointer manager object
//
// Used by the client to notify the server that it will no longer use this
// relative pointer manager object.
//
func (i *RelativePointerManager) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const rLen = 8
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// GetRelativePointer : get a relative pointer object
//
// Create a relative pointer interface given a wl_pointer object. See the
// wp_relative_pointer interface for more details.
//
func (i *RelativePointerManager) GetRelativePointer(pointer *client.Pointer) (*RelativePointer, error) {
	id := NewRelativePointer(i.Context())
	const opcode = 1
	const rLen = 8 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	client.PutUint32(r[l:l+4], pointer.ID())
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return id, err
}

// RelativePointer : relative pointer object
//
// A wp_relative_pointer object is an extension to the wl_pointer interface
// used for emitting relative pointer events. It shares the same focus as
// wl_pointer objects of the same seat and will only emit events when it has
// focus.
type RelativePointer struct {
	client.BaseProxy
	relativeMotionHandlers []RelativePointerRelativeMotionHandlerFunc
}

// NewRelativePointer : relative pointer object
//
// A wp_relative_pointer object is an extension to the wl_pointer interface
// used for emitting relative pointer events. It shares the same focus as
// wl_pointer objects of the same seat and will only emit events when it has
// focus.
func NewRelativePointer(ctx *client.Context) *RelativePointer {
	zwpRelativePointerV1 := &RelativePointer{}
	ctx.Register(zwpRelativePointerV1)
	return zwpRelativePointerV1
}

// Destroy : release the relative pointer object
//
func (i *RelativePointer) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const rLen = 8
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// RelativePointerRelativeMotionEvent : relative pointer motion
//
// Relative x/y pointer motion from the pointer of the seat associated with
// this object.
//
// A relative motion is in the same dimension as regular wl_pointer motion
// events, except they do not represent an absolute position. For example,
// moving a pointer from (x, y) to (x', y') would have the equivalent
// relative motion (x' - x, y' - y). If a pointer motion caused the
// absolute pointer position to be clipped by for example the edge of the
// monitor, the relative motion is unaffected by the clipping and will
// represent the unclipped motion.
//
// This event also contains non-accelerated motion deltas. The
// non-accelerated delta is, when applicable, the regular pointer motion
// delta as it was before having applied motion acceleration and other
// transformations such as normalization.
//
// Note that the non-accelerated delta does not represent 'raw' events as
// they were read from some device. Pointer motion acceleration is device-
// and configuration-specific and non-accelerated deltas and accelerated
// deltas may have the same value on some devices.
//
// Relative motions are not coupled to wl_pointer.motion events, and can be
// sent in combination with such events, but also independently. There may
// also be scenarios where wl_pointer.motion is sent, but there is no
// relative motion. The order of an absolute and relative motion event
// originating from the same physical motion is not guaranteed.
//
// If the client needs button events or focus state, it can receive them
// from a wl_pointer object of the same seat that the wp_relative_pointer
// object is associated with.
type RelativePointerRelativeMotionEvent struct {
	UtimeHi   uint32
	UtimeLo   uint32
	Dx        float64
	Dy        float64
	DxUnaccel float64
	DyUnaccel float64
}
type RelativePointerRelativeMotionHandlerFunc func(RelativePointerRelativeMotionEvent)

// AddRelativeMotionHandler : adds handler for RelativePointerRelativeMotionEvent
func (i *RelativePointer) AddRelativeMotionHandler(f RelativePointerRelativeMotionHandlerFunc) {
	if f == nil {
		return
	}

	i.relativeMotionHandlers = append(i.relativeMotionHandlers, f)
}

func (i *RelativePointer) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.relativeMotionHandlers) == 0 {
			return
		}
		var e RelativePointerRelativeMotionEvent
		l := 0
		e.UtimeHi = client.Uint32(data[l : l+4])
		l += 4
		e.UtimeLo = client.Uint32(data[l : l+4])
		l += 4
		e.Dx = client.Fixed(data[l : l+4])
		l += 4
		e.Dy = client.Fixed(data[l : l+4])
		l += 4
		e.DxUnaccel = client.Fixed(data[l : l+4])
		l += 4
		e.DyUnaccel = client.Fixed(data[l : l+4])
		l += 4
		for _, f := range i.relativeMotionHandlers {
			f(e)
		}
	}
}
