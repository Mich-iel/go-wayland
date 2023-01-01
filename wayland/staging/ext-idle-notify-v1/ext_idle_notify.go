// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.31/staging/ext-idle-notify/ext-idle-notify-v1.xml
//
// ext_idle_notify_v1 Protocol Copyright:
//
// Copyright © 2015 Martin Gräßlin
// Copyright © 2022 Simon Ser
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

package ext_idle_notify

import "github.com/rajveermalviya/go-wayland/wayland/client"

// IdleNotifier : idle notification manager
//
// This interface allows clients to monitor user idle status.
//
// After binding to this global, clients can create ext_idle_notification_v1
// objects to get notified when the user is idle for a given amount of time.
type IdleNotifier struct {
	client.BaseProxy
}

// NewIdleNotifier : idle notification manager
//
// This interface allows clients to monitor user idle status.
//
// After binding to this global, clients can create ext_idle_notification_v1
// objects to get notified when the user is idle for a given amount of time.
func NewIdleNotifier(ctx *client.Context) *IdleNotifier {
	extIdleNotifierV1 := &IdleNotifier{}
	ctx.Register(extIdleNotifierV1)
	return extIdleNotifierV1
}

// Destroy : destroy the manager
//
// Destroy the manager object. All objects created via this interface
// remain valid.
func (i *IdleNotifier) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// GetIdleNotification : create a notification object
//
// Create a new idle notification object.
//
// The notification object has a minimum timeout duration and is tied to a
// seat. The client will be notified if the seat is inactive for at least
// the provided timeout. See ext_idle_notification_v1 for more details.
//
// A zero timeout is valid and means the client wants to be notified as
// soon as possible when the seat is inactive.
//
//	timeout: minimum idle timeout in msec
func (i *IdleNotifier) GetIdleNotification(timeout uint32, seat *client.Seat) (*IdleNotification, error) {
	id := NewIdleNotification(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(timeout))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], seat.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}

// IdleNotification : idle notification
//
// This interface is used by the compositor to send idle notification events
// to clients.
//
// Initially the notification object is not idle. The notification object
// becomes idle when no user activity has happened for at least the timeout
// duration, starting from the creation of the notification object. User
// activity may include input events or a presence sensor, but is
// compositor-specific. If an idle inhibitor is active (e.g. another client
// has created a zwp_idle_inhibitor_v1 on a visible surface), the compositor
// must not make the notification object idle.
//
// When the notification object becomes idle, an idled event is sent. When
// user activity starts again, the notification object stops being idle,
// a resumed event is sent and the timeout is restarted.
type IdleNotification struct {
	client.BaseProxy
	idledHandlers   []IdleNotificationIdledHandlerFunc
	resumedHandlers []IdleNotificationResumedHandlerFunc
}

// NewIdleNotification : idle notification
//
// This interface is used by the compositor to send idle notification events
// to clients.
//
// Initially the notification object is not idle. The notification object
// becomes idle when no user activity has happened for at least the timeout
// duration, starting from the creation of the notification object. User
// activity may include input events or a presence sensor, but is
// compositor-specific. If an idle inhibitor is active (e.g. another client
// has created a zwp_idle_inhibitor_v1 on a visible surface), the compositor
// must not make the notification object idle.
//
// When the notification object becomes idle, an idled event is sent. When
// user activity starts again, the notification object stops being idle,
// a resumed event is sent and the timeout is restarted.
func NewIdleNotification(ctx *client.Context) *IdleNotification {
	extIdleNotificationV1 := &IdleNotification{}
	ctx.Register(extIdleNotificationV1)
	return extIdleNotificationV1
}

// Destroy : destroy the notification object
//
// Destroy the notification object.
func (i *IdleNotification) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// IdleNotificationIdledEvent : notification object is idle
//
// This event is sent when the notification object becomes idle.
//
// It's a compositor protocol error to send this event twice without a
// resumed event in-between.
type IdleNotificationIdledEvent struct{}
type IdleNotificationIdledHandlerFunc func(IdleNotificationIdledEvent)

// AddIdledHandler : adds handler for IdleNotificationIdledEvent
func (i *IdleNotification) AddIdledHandler(f IdleNotificationIdledHandlerFunc) {
	if f == nil {
		return
	}

	i.idledHandlers = append(i.idledHandlers, f)
}

// IdleNotificationResumedEvent : notification object is no longer idle
//
// This event is sent when the notification object stops being idle.
//
// It's a compositor protocol error to send this event twice without an
// idled event in-between. It's a compositor protocol error to send this
// event prior to any idled event.
type IdleNotificationResumedEvent struct{}
type IdleNotificationResumedHandlerFunc func(IdleNotificationResumedEvent)

// AddResumedHandler : adds handler for IdleNotificationResumedEvent
func (i *IdleNotification) AddResumedHandler(f IdleNotificationResumedHandlerFunc) {
	if f == nil {
		return
	}

	i.resumedHandlers = append(i.resumedHandlers, f)
}

func (i *IdleNotification) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.idledHandlers) == 0 {
			return
		}
		var e IdleNotificationIdledEvent
		for _, f := range i.idledHandlers {
			f(e)
		}
	case 1:
		if len(i.resumedHandlers) == 0 {
			return
		}
		var e IdleNotificationResumedEvent
		for _, f := range i.resumedHandlers {
			f(e)
		}
	}
}
