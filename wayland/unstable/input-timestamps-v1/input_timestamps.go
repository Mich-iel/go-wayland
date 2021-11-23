// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.24/unstable/input-timestamps/input-timestamps-unstable-v1.xml
//
// input_timestamps_unstable_v1 Protocol Copyright:
//
// Copyright © 2017 Collabora, Ltd.
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

package input_timestamps

import (
	"reflect"

	"github.com/rajveermalviya/go-wayland/wayland/client"
)

// InputTimestampsManager : context object for high-resolution input timestamps
//
// A global interface used for requesting high-resolution timestamps
// for input events.
type InputTimestampsManager struct {
	client.BaseProxy
}

// NewInputTimestampsManager : context object for high-resolution input timestamps
//
// A global interface used for requesting high-resolution timestamps
// for input events.
func NewInputTimestampsManager(ctx *client.Context) *InputTimestampsManager {
	zwpInputTimestampsManagerV1 := &InputTimestampsManager{}
	ctx.Register(zwpInputTimestampsManagerV1)
	return zwpInputTimestampsManagerV1
}

// Destroy : destroy the input timestamps manager object
//
// Informs the server that the client will no longer be using this
// protocol object. Existing objects created by this object are not
// affected.
//
func (i *InputTimestampsManager) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const rLen = 8
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return err
}

// GetKeyboardTimestamps : subscribe to high-resolution keyboard timestamp events
//
// Creates a new input timestamps object that represents a subscription
// to high-resolution timestamp events for all wl_keyboard events that
// carry a timestamp.
//
// If the associated wl_keyboard object is invalidated, either through
// client action (e.g. release) or server-side changes, the input
// timestamps object becomes inert and the client should destroy it
// by calling zwp_input_timestamps_v1.destroy.
//
//  keyboard: the wl_keyboard object for which to get timestamp events
func (i *InputTimestampsManager) GetKeyboardTimestamps(keyboard *client.Keyboard) (*InputTimestamps, error) {
	id := NewInputTimestamps(i.Context())
	const opcode = 1
	const rLen = 8 + 4 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	client.PutUint32(r[l:l+4], keyboard.ID())
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return id, err
}

// GetPointerTimestamps : subscribe to high-resolution pointer timestamp events
//
// Creates a new input timestamps object that represents a subscription
// to high-resolution timestamp events for all wl_pointer events that
// carry a timestamp.
//
// If the associated wl_pointer object is invalidated, either through
// client action (e.g. release) or server-side changes, the input
// timestamps object becomes inert and the client should destroy it
// by calling zwp_input_timestamps_v1.destroy.
//
//  pointer: the wl_pointer object for which to get timestamp events
func (i *InputTimestampsManager) GetPointerTimestamps(pointer *client.Pointer) (*InputTimestamps, error) {
	id := NewInputTimestamps(i.Context())
	const opcode = 2
	const rLen = 8 + 4 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	client.PutUint32(r[l:l+4], pointer.ID())
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return id, err
}

// GetTouchTimestamps : subscribe to high-resolution touch timestamp events
//
// Creates a new input timestamps object that represents a subscription
// to high-resolution timestamp events for all wl_touch events that
// carry a timestamp.
//
// If the associated wl_touch object becomes invalid, either through
// client action (e.g. release) or server-side changes, the input
// timestamps object becomes inert and the client should destroy it
// by calling zwp_input_timestamps_v1.destroy.
//
//  touch: the wl_touch object for which to get timestamp events
func (i *InputTimestampsManager) GetTouchTimestamps(touch *client.Touch) (*InputTimestamps, error) {
	id := NewInputTimestamps(i.Context())
	const opcode = 3
	const rLen = 8 + 4 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	client.PutUint32(r[l:l+4], touch.ID())
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return id, err
}

// InputTimestamps : context object for input timestamps
//
// Provides high-resolution timestamp events for a set of subscribed input
// events. The set of subscribed input events is determined by the
// zwp_input_timestamps_manager_v1 request used to create this object.
type InputTimestamps struct {
	client.BaseProxy
	timestampHandlers []InputTimestampsTimestampHandlerFunc
}

// NewInputTimestamps : context object for input timestamps
//
// Provides high-resolution timestamp events for a set of subscribed input
// events. The set of subscribed input events is determined by the
// zwp_input_timestamps_manager_v1 request used to create this object.
func NewInputTimestamps(ctx *client.Context) *InputTimestamps {
	zwpInputTimestampsV1 := &InputTimestamps{}
	ctx.Register(zwpInputTimestampsV1)
	return zwpInputTimestampsV1
}

// Destroy : destroy the input timestamps object
//
// Informs the server that the client will no longer be using this
// protocol object. After the server processes the request, no more
// timestamp events will be emitted.
//
func (i *InputTimestamps) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const rLen = 8
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return err
}

// InputTimestampsTimestampEvent : high-resolution timestamp event
//
// The timestamp event is associated with the first subsequent input event
// carrying a timestamp which belongs to the set of input events this
// object is subscribed to.
//
// The timestamp provided by this event is a high-resolution version of
// the timestamp argument of the associated input event. The provided
// timestamp is in the same clock domain and is at least as accurate as
// the associated input event timestamp.
//
// The timestamp is expressed as tv_sec_hi, tv_sec_lo, tv_nsec triples,
// each component being an unsigned 32-bit value. Whole seconds are in
// tv_sec which is a 64-bit value combined from tv_sec_hi and tv_sec_lo,
// and the additional fractional part in tv_nsec as nanoseconds. Hence,
// for valid timestamps tv_nsec must be in [0, 999999999].
type InputTimestampsTimestampEvent struct {
	TvSecHi uint32
	TvSecLo uint32
	TvNsec  uint32
}
type InputTimestampsTimestampHandlerFunc func(InputTimestampsTimestampEvent)

// AddTimestampHandler : adds handler for InputTimestampsTimestampEvent
func (i *InputTimestamps) AddTimestampHandler(f InputTimestampsTimestampHandlerFunc) {
	if f == nil {
		return
	}

	i.timestampHandlers = append(i.timestampHandlers, f)
}

func (i *InputTimestamps) RemoveTimestampHandler(f InputTimestampsTimestampHandlerFunc) {
	for j, e := range i.timestampHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.timestampHandlers = append(i.timestampHandlers[:j], i.timestampHandlers[j+1:]...)
			return
		}
	}
}

func (i *InputTimestamps) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.timestampHandlers) == 0 {
			return
		}
		var e InputTimestampsTimestampEvent
		l := 0
		e.TvSecHi = client.Uint32(data[l : l+4])
		l += 4
		e.TvSecLo = client.Uint32(data[l : l+4])
		l += 4
		e.TvNsec = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.timestampHandlers {
			f(e)
		}
	}
}
