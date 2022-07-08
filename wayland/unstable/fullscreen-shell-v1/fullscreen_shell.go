// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.26/unstable/fullscreen-shell/fullscreen-shell-unstable-v1.xml
//
// fullscreen_shell_unstable_v1 Protocol Copyright:
//
// Copyright © 2016 Yong Bakos
// Copyright © 2015 Jason Ekstrand
// Copyright © 2015 Jonas Ådahl
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

package fullscreen_shell

import "github.com/rajveermalviya/go-wayland/wayland/client"

// FullscreenShell : displays a single surface per output
//
// Displays a single surface per output.
//
// This interface provides a mechanism for a single client to display
// simple full-screen surfaces.  While there technically may be multiple
// clients bound to this interface, only one of those clients should be
// shown at a time.
//
// To present a surface, the client uses either the present_surface or
// present_surface_for_mode requests.  Presenting a surface takes effect
// on the next wl_surface.commit.  See the individual requests for
// details about scaling and mode switches.
//
// The client can have at most one surface per output at any time.
// Requesting a surface to be presented on an output that already has a
// surface replaces the previously presented surface.  Presenting a null
// surface removes its content and effectively disables the output.
// Exactly what happens when an output is "disabled" is
// compositor-specific.  The same surface may be presented on multiple
// outputs simultaneously.
//
// Once a surface is presented on an output, it stays on that output
// until either the client removes it or the compositor destroys the
// output.  This way, the client can update the output's contents by
// simply attaching a new buffer.
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
type FullscreenShell struct {
	client.BaseProxy
	capabilityHandlers []FullscreenShellCapabilityHandlerFunc
}

// NewFullscreenShell : displays a single surface per output
//
// Displays a single surface per output.
//
// This interface provides a mechanism for a single client to display
// simple full-screen surfaces.  While there technically may be multiple
// clients bound to this interface, only one of those clients should be
// shown at a time.
//
// To present a surface, the client uses either the present_surface or
// present_surface_for_mode requests.  Presenting a surface takes effect
// on the next wl_surface.commit.  See the individual requests for
// details about scaling and mode switches.
//
// The client can have at most one surface per output at any time.
// Requesting a surface to be presented on an output that already has a
// surface replaces the previously presented surface.  Presenting a null
// surface removes its content and effectively disables the output.
// Exactly what happens when an output is "disabled" is
// compositor-specific.  The same surface may be presented on multiple
// outputs simultaneously.
//
// Once a surface is presented on an output, it stays on that output
// until either the client removes it or the compositor destroys the
// output.  This way, the client can update the output's contents by
// simply attaching a new buffer.
//
// Warning! The protocol described in this file is experimental and
// backward incompatible changes may be made. Backward compatible changes
// may be added together with the corresponding interface version bump.
// Backward incompatible changes are done by bumping the version number in
// the protocol and interface names and resetting the interface version.
// Once the protocol is to be declared stable, the 'z' prefix and the
// version number in the protocol and interface names are removed and the
// interface version number is reset.
func NewFullscreenShell(ctx *client.Context) *FullscreenShell {
	zwpFullscreenShellV1 := &FullscreenShell{}
	ctx.Register(zwpFullscreenShellV1)
	return zwpFullscreenShellV1
}

// Release : release the wl_fullscreen_shell interface
//
// Release the binding from the wl_fullscreen_shell interface.
//
// This destroys the server-side object and frees this binding.  If
// the client binds to wl_fullscreen_shell multiple times, it may wish
// to free some of those bindings.
//
func (i *FullscreenShell) Release() error {
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

// PresentSurface : present surface for display
//
// Present a surface on the given output.
//
// If the output is null, the compositor will present the surface on
// whatever display (or displays) it thinks best.  In particular, this
// may replace any or all surfaces currently presented so it should
// not be used in combination with placing surfaces on specific
// outputs.
//
// The method parameter is a hint to the compositor for how the surface
// is to be presented.  In particular, it tells the compositor how to
// handle a size mismatch between the presented surface and the
// output.  The compositor is free to ignore this parameter.
//
// The "zoom", "zoom_crop", and "stretch" methods imply a scaling
// operation on the surface.  This will override any kind of output
// scaling, so the buffer_scale property of the surface is effectively
// ignored.
//
// This request gives the surface the role of a fullscreen shell surface.
// If the surface already has another role, it raises a role protocol
// error.
//
func (i *FullscreenShell) PresentSurface(surface *client.Surface, method uint32, output *client.Output) error {
	const opcode = 1
	const rLen = 8 + 4 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	if surface == nil {
		client.PutUint32(r[l:l+4], 0)
		l += 4
	} else {
		client.PutUint32(r[l:l+4], surface.ID())
		l += 4
	}
	client.PutUint32(r[l:l+4], uint32(method))
	l += 4
	if output == nil {
		client.PutUint32(r[l:l+4], 0)
		l += 4
	} else {
		client.PutUint32(r[l:l+4], output.ID())
		l += 4
	}
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

// PresentSurfaceForMode : present surface for display at a particular mode
//
// Presents a surface on the given output for a particular mode.
//
// If the current size of the output differs from that of the surface,
// the compositor will attempt to change the size of the output to
// match the surface.  The result of the mode-switch operation will be
// returned via the provided wl_fullscreen_shell_mode_feedback object.
//
// If the current output mode matches the one requested or if the
// compositor successfully switches the mode to match the surface,
// then the mode_successful event will be sent and the output will
// contain the contents of the given surface.  If the compositor
// cannot match the output size to the surface size, the mode_failed
// will be sent and the output will contain the contents of the
// previously presented surface (if any).  If another surface is
// presented on the given output before either of these has a chance
// to happen, the present_cancelled event will be sent.
//
// Due to race conditions and other issues unknown to the client, no
// mode-switch operation is guaranteed to succeed.  However, if the
// mode is one advertised by wl_output.mode or if the compositor
// advertises the ARBITRARY_MODES capability, then the client should
// expect that the mode-switch operation will usually succeed.
//
// If the size of the presented surface changes, the resulting output
// is undefined.  The compositor may attempt to change the output mode
// to compensate.  However, there is no guarantee that a suitable mode
// will be found and the client has no way to be notified of success
// or failure.
//
// The framerate parameter specifies the desired framerate for the
// output in mHz.  The compositor is free to ignore this parameter.  A
// value of 0 indicates that the client has no preference.
//
// If the value of wl_output.scale differs from wl_surface.buffer_scale,
// then the compositor may choose a mode that matches either the buffer
// size or the surface size.  In either case, the surface will fill the
// output.
//
// This request gives the surface the role of a fullscreen shell surface.
// If the surface already has another role, it raises a role protocol
// error.
//
func (i *FullscreenShell) PresentSurfaceForMode(surface *client.Surface, output *client.Output, framerate int32) (*FullscreenShellModeFeedback, error) {
	feedback := NewFullscreenShellModeFeedback(i.Context())
	const opcode = 2
	const rLen = 8 + 4 + 4 + 4 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], surface.ID())
	l += 4
	client.PutUint32(r[l:l+4], output.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(framerate))
	l += 4
	client.PutUint32(r[l:l+4], feedback.ID())
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return feedback, err
}

type FullscreenShellCapability uint32

// FullscreenShellCapability : capabilities advertised by the compositor
//
// Various capabilities that can be advertised by the compositor.  They
// are advertised one-at-a-time when the wl_fullscreen_shell interface is
// bound.  See the wl_fullscreen_shell.capability event for more details.
//
// ARBITRARY_MODES:
// This is a hint to the client that indicates that the compositor is
// capable of setting practically any mode on its outputs.  If this
// capability is provided, wl_fullscreen_shell.present_surface_for_mode
// will almost never fail and clients should feel free to set whatever
// mode they like.  If the compositor does not advertise this, it may
// still support some modes that are not advertised through wl_global.mode
// but it is less likely.
//
// CURSOR_PLANE:
// This is a hint to the client that indicates that the compositor can
// handle a cursor surface from the client without actually compositing.
// This may be because of a hardware cursor plane or some other mechanism.
// If the compositor does not advertise this capability then setting
// wl_pointer.cursor may degrade performance or be ignored entirely.  If
// CURSOR_PLANE is not advertised, it is recommended that the client draw
// its own cursor and set wl_pointer.cursor(NULL).
const (
	// FullscreenShellCapabilityArbitraryModes : compositor is capable of almost any output mode
	FullscreenShellCapabilityArbitraryModes FullscreenShellCapability = 1
	// FullscreenShellCapabilityCursorPlane : compositor has a separate cursor plane
	FullscreenShellCapabilityCursorPlane FullscreenShellCapability = 2
)

func (e FullscreenShellCapability) Name() string {
	switch e {
	case FullscreenShellCapabilityArbitraryModes:
		return "arbitrary_modes"
	case FullscreenShellCapabilityCursorPlane:
		return "cursor_plane"
	default:
		return ""
	}
}

func (e FullscreenShellCapability) Value() string {
	switch e {
	case FullscreenShellCapabilityArbitraryModes:
		return "1"
	case FullscreenShellCapabilityCursorPlane:
		return "2"
	default:
		return ""
	}
}

func (e FullscreenShellCapability) String() string {
	return e.Name() + "=" + e.Value()
}

type FullscreenShellPresentMethod uint32

// FullscreenShellPresentMethod : different method to set the surface fullscreen
//
// Hints to indicate to the compositor how to deal with a conflict
// between the dimensions of the surface and the dimensions of the
// output. The compositor is free to ignore this parameter.
const (
	// FullscreenShellPresentMethodDefault : no preference, apply default policy
	FullscreenShellPresentMethodDefault FullscreenShellPresentMethod = 0
	// FullscreenShellPresentMethodCenter : center the surface on the output
	FullscreenShellPresentMethodCenter FullscreenShellPresentMethod = 1
	// FullscreenShellPresentMethodZoom : scale the surface, preserving aspect ratio, to the largest size that will fit on the output
	FullscreenShellPresentMethodZoom FullscreenShellPresentMethod = 2
	// FullscreenShellPresentMethodZoomCrop : scale the surface, preserving aspect ratio, to fully fill the output cropping if needed
	FullscreenShellPresentMethodZoomCrop FullscreenShellPresentMethod = 3
	// FullscreenShellPresentMethodStretch : scale the surface to the size of the output ignoring aspect ratio
	FullscreenShellPresentMethodStretch FullscreenShellPresentMethod = 4
)

func (e FullscreenShellPresentMethod) Name() string {
	switch e {
	case FullscreenShellPresentMethodDefault:
		return "default"
	case FullscreenShellPresentMethodCenter:
		return "center"
	case FullscreenShellPresentMethodZoom:
		return "zoom"
	case FullscreenShellPresentMethodZoomCrop:
		return "zoom_crop"
	case FullscreenShellPresentMethodStretch:
		return "stretch"
	default:
		return ""
	}
}

func (e FullscreenShellPresentMethod) Value() string {
	switch e {
	case FullscreenShellPresentMethodDefault:
		return "0"
	case FullscreenShellPresentMethodCenter:
		return "1"
	case FullscreenShellPresentMethodZoom:
		return "2"
	case FullscreenShellPresentMethodZoomCrop:
		return "3"
	case FullscreenShellPresentMethodStretch:
		return "4"
	default:
		return ""
	}
}

func (e FullscreenShellPresentMethod) String() string {
	return e.Name() + "=" + e.Value()
}

type FullscreenShellError uint32

// FullscreenShellError : wl_fullscreen_shell error values
//
// These errors can be emitted in response to wl_fullscreen_shell requests.
const (
	// FullscreenShellErrorInvalidMethod : present_method is not known
	FullscreenShellErrorInvalidMethod FullscreenShellError = 0
	// FullscreenShellErrorRole : given wl_surface has another role
	FullscreenShellErrorRole FullscreenShellError = 1
)

func (e FullscreenShellError) Name() string {
	switch e {
	case FullscreenShellErrorInvalidMethod:
		return "invalid_method"
	case FullscreenShellErrorRole:
		return "role"
	default:
		return ""
	}
}

func (e FullscreenShellError) Value() string {
	switch e {
	case FullscreenShellErrorInvalidMethod:
		return "0"
	case FullscreenShellErrorRole:
		return "1"
	default:
		return ""
	}
}

func (e FullscreenShellError) String() string {
	return e.Name() + "=" + e.Value()
}

// FullscreenShellCapabilityEvent : advertises a capability of the compositor
//
// Advertises a single capability of the compositor.
//
// When the wl_fullscreen_shell interface is bound, this event is emitted
// once for each capability advertised.  Valid capabilities are given by
// the wl_fullscreen_shell.capability enum.  If clients want to take
// advantage of any of these capabilities, they should use a
// wl_display.sync request immediately after binding to ensure that they
// receive all the capability events.
type FullscreenShellCapabilityEvent struct {
	Capability uint32
}
type FullscreenShellCapabilityHandlerFunc func(FullscreenShellCapabilityEvent)

// AddCapabilityHandler : adds handler for FullscreenShellCapabilityEvent
func (i *FullscreenShell) AddCapabilityHandler(f FullscreenShellCapabilityHandlerFunc) {
	if f == nil {
		return
	}

	i.capabilityHandlers = append(i.capabilityHandlers, f)
}

func (i *FullscreenShell) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.capabilityHandlers) == 0 {
			return
		}
		var e FullscreenShellCapabilityEvent
		l := 0
		e.Capability = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.capabilityHandlers {
			f(e)
		}
	}
}

// FullscreenShellModeFeedback :
type FullscreenShellModeFeedback struct {
	client.BaseProxy
	modeSuccessfulHandlers   []FullscreenShellModeFeedbackModeSuccessfulHandlerFunc
	modeFailedHandlers       []FullscreenShellModeFeedbackModeFailedHandlerFunc
	presentCancelledHandlers []FullscreenShellModeFeedbackPresentCancelledHandlerFunc
}

// NewFullscreenShellModeFeedback :
func NewFullscreenShellModeFeedback(ctx *client.Context) *FullscreenShellModeFeedback {
	zwpFullscreenShellModeFeedbackV1 := &FullscreenShellModeFeedback{}
	ctx.Register(zwpFullscreenShellModeFeedbackV1)
	return zwpFullscreenShellModeFeedbackV1
}

func (i *FullscreenShellModeFeedback) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

// FullscreenShellModeFeedbackModeSuccessfulEvent : mode switch succeeded
//
// This event indicates that the attempted mode switch operation was
// successful.  A surface of the size requested in the mode switch
// will fill the output without scaling.
//
// Upon receiving this event, the client should destroy the
// wl_fullscreen_shell_mode_feedback object.
type FullscreenShellModeFeedbackModeSuccessfulEvent struct{}
type FullscreenShellModeFeedbackModeSuccessfulHandlerFunc func(FullscreenShellModeFeedbackModeSuccessfulEvent)

// AddModeSuccessfulHandler : adds handler for FullscreenShellModeFeedbackModeSuccessfulEvent
func (i *FullscreenShellModeFeedback) AddModeSuccessfulHandler(f FullscreenShellModeFeedbackModeSuccessfulHandlerFunc) {
	if f == nil {
		return
	}

	i.modeSuccessfulHandlers = append(i.modeSuccessfulHandlers, f)
}

// FullscreenShellModeFeedbackModeFailedEvent : mode switch failed
//
// This event indicates that the attempted mode switch operation
// failed.  This may be because the requested output mode is not
// possible or it may mean that the compositor does not want to allow it.
//
// Upon receiving this event, the client should destroy the
// wl_fullscreen_shell_mode_feedback object.
type FullscreenShellModeFeedbackModeFailedEvent struct{}
type FullscreenShellModeFeedbackModeFailedHandlerFunc func(FullscreenShellModeFeedbackModeFailedEvent)

// AddModeFailedHandler : adds handler for FullscreenShellModeFeedbackModeFailedEvent
func (i *FullscreenShellModeFeedback) AddModeFailedHandler(f FullscreenShellModeFeedbackModeFailedHandlerFunc) {
	if f == nil {
		return
	}

	i.modeFailedHandlers = append(i.modeFailedHandlers, f)
}

// FullscreenShellModeFeedbackPresentCancelledEvent : mode switch cancelled
//
// This event indicates that the attempted mode switch operation was
// cancelled.  Most likely this is because the client requested a
// second mode switch before the first one completed.
//
// Upon receiving this event, the client should destroy the
// wl_fullscreen_shell_mode_feedback object.
type FullscreenShellModeFeedbackPresentCancelledEvent struct{}
type FullscreenShellModeFeedbackPresentCancelledHandlerFunc func(FullscreenShellModeFeedbackPresentCancelledEvent)

// AddPresentCancelledHandler : adds handler for FullscreenShellModeFeedbackPresentCancelledEvent
func (i *FullscreenShellModeFeedback) AddPresentCancelledHandler(f FullscreenShellModeFeedbackPresentCancelledHandlerFunc) {
	if f == nil {
		return
	}

	i.presentCancelledHandlers = append(i.presentCancelledHandlers, f)
}

func (i *FullscreenShellModeFeedback) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.modeSuccessfulHandlers) == 0 {
			return
		}
		var e FullscreenShellModeFeedbackModeSuccessfulEvent
		for _, f := range i.modeSuccessfulHandlers {
			f(e)
		}
	case 1:
		if len(i.modeFailedHandlers) == 0 {
			return
		}
		var e FullscreenShellModeFeedbackModeFailedEvent
		for _, f := range i.modeFailedHandlers {
			f(e)
		}
	case 2:
		if len(i.presentCancelledHandlers) == 0 {
			return
		}
		var e FullscreenShellModeFeedbackPresentCancelledEvent
		for _, f := range i.presentCancelledHandlers {
			f(e)
		}
	}
}
