// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.26/unstable/xdg-foreign/xdg-foreign-unstable-v2.xml
//
// xdg_foreign_unstable_v2 Protocol Copyright:
//
// Copyright © 2015-2016 Red Hat Inc.
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

package xdg_foreign

import "github.com/rajveermalviya/go-wayland/wayland/client"

// Exporter : interface for exporting surfaces
//
// A global interface used for exporting surfaces that can later be imported
// using xdg_importer.
type Exporter struct {
	client.BaseProxy
}

// NewExporter : interface for exporting surfaces
//
// A global interface used for exporting surfaces that can later be imported
// using xdg_importer.
func NewExporter(ctx *client.Context) *Exporter {
	zxdgExporterV2 := &Exporter{}
	ctx.Register(zxdgExporterV2)
	return zxdgExporterV2
}

// Destroy : destroy the xdg_exporter object
//
// Notify the compositor that the xdg_exporter object will no longer be
// used.
//
func (i *Exporter) Destroy() error {
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

// ExportToplevel : export a toplevel surface
//
// The export_toplevel request exports the passed surface so that it can later be
// imported via xdg_importer. When called, a new xdg_exported object will
// be created and xdg_exported.handle will be sent immediately. See the
// corresponding interface and event for details.
//
// A surface may be exported multiple times, and each exported handle may
// be used to create an xdg_imported multiple times. Only xdg_toplevel
// equivalent surfaces may be exported, otherwise an invalid_surface
// protocol error is sent.
//
//  surface: the surface to export
func (i *Exporter) ExportToplevel(surface *client.Surface) (*Exported, error) {
	id := NewExported(i.Context())
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
	client.PutUint32(r[l:l+4], surface.ID())
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return id, err
}

type ExporterError uint32

// ExporterError : error values
//
// These errors can be emitted in response to invalid xdg_exporter
// requests.
const (
	// ExporterErrorInvalidSurface : surface is not an xdg_toplevel
	ExporterErrorInvalidSurface ExporterError = 0
)

func (e ExporterError) Name() string {
	switch e {
	case ExporterErrorInvalidSurface:
		return "invalid_surface"
	default:
		return ""
	}
}

func (e ExporterError) Value() string {
	switch e {
	case ExporterErrorInvalidSurface:
		return "0"
	default:
		return ""
	}
}

func (e ExporterError) String() string {
	return e.Name() + "=" + e.Value()
}

// Importer : interface for importing surfaces
//
// A global interface used for importing surfaces exported by xdg_exporter.
// With this interface, a client can create a reference to a surface of
// another client.
type Importer struct {
	client.BaseProxy
}

// NewImporter : interface for importing surfaces
//
// A global interface used for importing surfaces exported by xdg_exporter.
// With this interface, a client can create a reference to a surface of
// another client.
func NewImporter(ctx *client.Context) *Importer {
	zxdgImporterV2 := &Importer{}
	ctx.Register(zxdgImporterV2)
	return zxdgImporterV2
}

// Destroy : destroy the xdg_importer object
//
// Notify the compositor that the xdg_importer object will no longer be
// used.
//
func (i *Importer) Destroy() error {
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

// ImportToplevel : import a toplevel surface
//
// The import_toplevel request imports a surface from any client given a handle
// retrieved by exporting said surface using xdg_exporter.export_toplevel.
// When called, a new xdg_imported object will be created. This new object
// represents the imported surface, and the importing client can
// manipulate its relationship using it. See xdg_imported for details.
//
//  handle: the exported surface handle
func (i *Importer) ImportToplevel(handle string) (*Imported, error) {
	id := NewImported(i.Context())
	const opcode = 1
	handleLen := client.PaddedLen(len(handle) + 1)
	rLen := 8 + 4 + (4 + handleLen)
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	client.PutString(r[l:l+(4+handleLen)], handle, handleLen)
	l += (4 + handleLen)
	err := i.Context().WriteMsg(r, nil)
	return id, err
}

// Exported : an exported surface handle
//
// An xdg_exported object represents an exported reference to a surface. The
// exported surface may be referenced as long as the xdg_exported object not
// destroyed. Destroying the xdg_exported invalidates any relationship the
// importer may have established using xdg_imported.
type Exported struct {
	client.BaseProxy
	handleHandlers []ExportedHandleHandlerFunc
}

// NewExported : an exported surface handle
//
// An xdg_exported object represents an exported reference to a surface. The
// exported surface may be referenced as long as the xdg_exported object not
// destroyed. Destroying the xdg_exported invalidates any relationship the
// importer may have established using xdg_imported.
func NewExported(ctx *client.Context) *Exported {
	zxdgExportedV2 := &Exported{}
	ctx.Register(zxdgExportedV2)
	return zxdgExportedV2
}

// Destroy : unexport the exported surface
//
// Revoke the previously exported surface. This invalidates any
// relationship the importer may have set up using the xdg_imported created
// given the handle sent via xdg_exported.handle.
//
func (i *Exported) Destroy() error {
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

// ExportedHandleEvent : the exported surface handle
//
// The handle event contains the unique handle of this exported surface
// reference. It may be shared with any client, which then can use it to
// import the surface by calling xdg_importer.import_toplevel. A handle
// may be used to import the surface multiple times.
type ExportedHandleEvent struct {
	Handle string
}
type ExportedHandleHandlerFunc func(ExportedHandleEvent)

// AddHandleHandler : adds handler for ExportedHandleEvent
func (i *Exported) AddHandleHandler(f ExportedHandleHandlerFunc) {
	if f == nil {
		return
	}

	i.handleHandlers = append(i.handleHandlers, f)
}

func (i *Exported) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.handleHandlers) == 0 {
			return
		}
		var e ExportedHandleEvent
		l := 0
		handleLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Handle = client.String(data[l : l+handleLen])
		l += handleLen
		for _, f := range i.handleHandlers {
			f(e)
		}
	}
}

// Imported : an imported surface handle
//
// An xdg_imported object represents an imported reference to surface exported
// by some client. A client can use this interface to manipulate
// relationships between its own surfaces and the imported surface.
type Imported struct {
	client.BaseProxy
	destroyedHandlers []ImportedDestroyedHandlerFunc
}

// NewImported : an imported surface handle
//
// An xdg_imported object represents an imported reference to surface exported
// by some client. A client can use this interface to manipulate
// relationships between its own surfaces and the imported surface.
func NewImported(ctx *client.Context) *Imported {
	zxdgImportedV2 := &Imported{}
	ctx.Register(zxdgImportedV2)
	return zxdgImportedV2
}

// Destroy : destroy the xdg_imported object
//
// Notify the compositor that it will no longer use the xdg_imported
// object. Any relationship that may have been set up will at this point
// be invalidated.
//
func (i *Imported) Destroy() error {
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

// SetParentOf : set as the parent of some surface
//
// Set the imported surface as the parent of some surface of the client.
// The passed surface must be an xdg_toplevel equivalent, otherwise an
// invalid_surface protocol error is sent. Calling this function sets up
// a surface to surface relation with the same stacking and positioning
// semantics as xdg_toplevel.set_parent.
//
//  surface: the child surface
func (i *Imported) SetParentOf(surface *client.Surface) error {
	const opcode = 1
	const rLen = 8 + 4
	var r [rLen]byte
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], surface.ID())
	l += 4
	err := i.Context().WriteMsg(r[:], nil)
	return err
}

type ImportedError uint32

// ImportedError : error values
//
// These errors can be emitted in response to invalid xdg_imported
// requests.
const (
	// ImportedErrorInvalidSurface : surface is not an xdg_toplevel
	ImportedErrorInvalidSurface ImportedError = 0
)

func (e ImportedError) Name() string {
	switch e {
	case ImportedErrorInvalidSurface:
		return "invalid_surface"
	default:
		return ""
	}
}

func (e ImportedError) Value() string {
	switch e {
	case ImportedErrorInvalidSurface:
		return "0"
	default:
		return ""
	}
}

func (e ImportedError) String() string {
	return e.Name() + "=" + e.Value()
}

// ImportedDestroyedEvent : the imported surface handle has been destroyed
//
// The imported surface handle has been destroyed and any relationship set
// up has been invalidated. This may happen for various reasons, for
// example if the exported surface or the exported surface handle has been
// destroyed, if the handle used for importing was invalid.
type ImportedDestroyedEvent struct{}
type ImportedDestroyedHandlerFunc func(ImportedDestroyedEvent)

// AddDestroyedHandler : adds handler for ImportedDestroyedEvent
func (i *Imported) AddDestroyedHandler(f ImportedDestroyedHandlerFunc) {
	if f == nil {
		return
	}

	i.destroyedHandlers = append(i.destroyedHandlers, f)
}

func (i *Imported) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.destroyedHandlers) == 0 {
			return
		}
		var e ImportedDestroyedEvent
		for _, f := range i.destroyedHandlers {
			f(e)
		}
	}
}
