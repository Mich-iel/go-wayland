// Generated by go-wayland-scanner
// https://github.com/rajveermalviya/go-wayland/cmd/go-wayland-scanner
// XML file : https://raw.githubusercontent.com/wayland-project/wayland-protocols/1.24/staging/drm-lease/drm-lease-v1.xml
//
// drm_lease_v1 Protocol Copyright:
//
// Copyright © 2018 NXP
// Copyright © 2019 Status Research & Development GmbH.
// Copyright © 2021 Xaver Hugl
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

package drm_lease

import (
	"reflect"

	"github.com/rajveermalviya/go-wayland/wayland/client"
)

// DrmLeaseDevice : lease device
//
// This protocol is used by Wayland compositors which act as Direct
// Renderering Manager (DRM) masters to lease DRM resources to Wayland
// clients.
//
// The compositor will advertise one wp_drm_lease_device_v1 global for each
// DRM node. Some time after a client binds to the wp_drm_lease_device_v1
// global, the compositor will send a drm_fd event followed by zero, one or
// more connector events. After all currently available connectors have been
// sent, the compositor will send a wp_drm_lease_device_v1.done event.
//
// When the list of connectors available for lease changes the compositor
// will send wp_drm_lease_device_v1.connector events for added connectors and
// wp_drm_lease_connector_v1.withdrawn events for removed connectors,
// followed by a wp_drm_lease_device_v1.done event.
//
// The compositor will indicate when a device is gone by removing the global
// via a wl_registry.global_remove event. Upon receiving this event, the
// client should destroy any matching wp_drm_lease_device_v1 object.
//
// To destroy a wp_drm_lease_device_v1 object, the client must first issue
// a release request. Upon receiving this request, the compositor will
// immediately send a released event and destroy the object. The client must
// continue to process and discard drm_fd and connector events until it
// receives the released event. Upon receiving the released event, the
// client can safely cleanup any client-side resources.
//
// Warning! The protocol described in this file is currently in the testing
// phase. Backward compatible changes may be added together with the
// corresponding interface version bump. Backward incompatible changes can
// only be done by creating a new major version of the extension.
type DrmLeaseDevice struct {
	client.BaseProxy
	drmFdHandlers     []DrmLeaseDeviceDrmFdHandlerFunc
	connectorHandlers []DrmLeaseDeviceConnectorHandlerFunc
	doneHandlers      []DrmLeaseDeviceDoneHandlerFunc
	releasedHandlers  []DrmLeaseDeviceReleasedHandlerFunc
}

// NewDrmLeaseDevice : lease device
//
// This protocol is used by Wayland compositors which act as Direct
// Renderering Manager (DRM) masters to lease DRM resources to Wayland
// clients.
//
// The compositor will advertise one wp_drm_lease_device_v1 global for each
// DRM node. Some time after a client binds to the wp_drm_lease_device_v1
// global, the compositor will send a drm_fd event followed by zero, one or
// more connector events. After all currently available connectors have been
// sent, the compositor will send a wp_drm_lease_device_v1.done event.
//
// When the list of connectors available for lease changes the compositor
// will send wp_drm_lease_device_v1.connector events for added connectors and
// wp_drm_lease_connector_v1.withdrawn events for removed connectors,
// followed by a wp_drm_lease_device_v1.done event.
//
// The compositor will indicate when a device is gone by removing the global
// via a wl_registry.global_remove event. Upon receiving this event, the
// client should destroy any matching wp_drm_lease_device_v1 object.
//
// To destroy a wp_drm_lease_device_v1 object, the client must first issue
// a release request. Upon receiving this request, the compositor will
// immediately send a released event and destroy the object. The client must
// continue to process and discard drm_fd and connector events until it
// receives the released event. Upon receiving the released event, the
// client can safely cleanup any client-side resources.
//
// Warning! The protocol described in this file is currently in the testing
// phase. Backward compatible changes may be added together with the
// corresponding interface version bump. Backward incompatible changes can
// only be done by creating a new major version of the extension.
func NewDrmLeaseDevice(ctx *client.Context) *DrmLeaseDevice {
	wpDrmLeaseDeviceV1 := &DrmLeaseDevice{}
	ctx.Register(wpDrmLeaseDeviceV1)
	return wpDrmLeaseDeviceV1
}

// CreateLeaseRequest : create a lease request object
//
// Creates a lease request object.
//
// See the documentation for wp_drm_lease_request_v1 for details.
//
func (i *DrmLeaseDevice) CreateLeaseRequest() (*DrmLeaseRequest, error) {
	id := NewDrmLeaseRequest(i.Context())
	const opcode = 0
	const rLen = 8 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return id, err
}

// Release : release this object
//
// Indicates the client no longer wishes to use this object. In response
// the compositor will immediately send the released event and destroy
// this object. It can however not guarantee that the client won't receive
// connector events before the released event. The client must not send any
// requests after this one, doing so will raise a wl_display error.
// Existing connectors, lease request and leases will not be affected.
//
func (i *DrmLeaseDevice) Release() error {
	const opcode = 1
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

func (i *DrmLeaseDevice) Destroy() error {
	i.Context().Unregister(i)
	return nil
}

// DrmLeaseDeviceDrmFdEvent : open a non-master fd for this DRM node
//
// The compositor will send this event when the wp_drm_lease_device_v1
// global is bound, although there are no guarantees as to how long this
// takes - the compositor might need to wait until regaining DRM master.
// The included fd is a non-master DRM file descriptor opened for this
// device and the compositor must not authenticate it.
// The purpose of this event is to give the client the ability to
// query DRM and discover information which may help them pick the
// appropriate DRM device or select the appropriate connectors therein.
type DrmLeaseDeviceDrmFdEvent struct {
	Fd uintptr
}
type DrmLeaseDeviceDrmFdHandlerFunc func(DrmLeaseDeviceDrmFdEvent)

// AddDrmFdHandler : adds handler for DrmLeaseDeviceDrmFdEvent
func (i *DrmLeaseDevice) AddDrmFdHandler(f DrmLeaseDeviceDrmFdHandlerFunc) {
	if f == nil {
		return
	}

	i.drmFdHandlers = append(i.drmFdHandlers, f)
}

func (i *DrmLeaseDevice) RemoveDrmFdHandler(f DrmLeaseDeviceDrmFdHandlerFunc) {
	for j, e := range i.drmFdHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.drmFdHandlers = append(i.drmFdHandlers[:j], i.drmFdHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseDeviceConnectorEvent : advertise connectors available for leases
//
// The compositor will use this event to advertise connectors available for
// lease by clients. This object may be passed into a lease request to
// indicate the client would like to lease that connector, see
// wp_drm_lease_request_v1.request_connector for details. While the
// compositor will make a best effort to not send disconnected connectors,
// no guarantees can be made.
//
// The compositor must send the drm_fd event before sending connectors.
// After the drm_fd event it will send all available connectors but may
// send additional connectors at any time.
type DrmLeaseDeviceConnectorEvent struct {
	Id *DrmLeaseConnector
}
type DrmLeaseDeviceConnectorHandlerFunc func(DrmLeaseDeviceConnectorEvent)

// AddConnectorHandler : adds handler for DrmLeaseDeviceConnectorEvent
func (i *DrmLeaseDevice) AddConnectorHandler(f DrmLeaseDeviceConnectorHandlerFunc) {
	if f == nil {
		return
	}

	i.connectorHandlers = append(i.connectorHandlers, f)
}

func (i *DrmLeaseDevice) RemoveConnectorHandler(f DrmLeaseDeviceConnectorHandlerFunc) {
	for j, e := range i.connectorHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.connectorHandlers = append(i.connectorHandlers[:j], i.connectorHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseDeviceDoneEvent : signals grouping of connectors
//
// The compositor will send this event to indicate that it has sent all
// currently available connectors after the client binds to the global or
// when it updates the connector list, for example on hotplug, drm master
// change or when a leased connector becomes available again. It will
// similarly send this event to group wp_drm_lease_connector_v1.withdrawn
// events of connectors of this device.
type DrmLeaseDeviceDoneEvent struct{}
type DrmLeaseDeviceDoneHandlerFunc func(DrmLeaseDeviceDoneEvent)

// AddDoneHandler : adds handler for DrmLeaseDeviceDoneEvent
func (i *DrmLeaseDevice) AddDoneHandler(f DrmLeaseDeviceDoneHandlerFunc) {
	if f == nil {
		return
	}

	i.doneHandlers = append(i.doneHandlers, f)
}

func (i *DrmLeaseDevice) RemoveDoneHandler(f DrmLeaseDeviceDoneHandlerFunc) {
	for j, e := range i.doneHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.doneHandlers = append(i.doneHandlers[:j], i.doneHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseDeviceReleasedEvent : the compositor has finished using the device
//
// This event is sent in response to the release request and indicates
// that the compositor is done sending connector events.
// The compositor will destroy this object immediately after sending the
// event and it will become invalid. The client should release any
// resources associated with this device after receiving this event.
type DrmLeaseDeviceReleasedEvent struct{}
type DrmLeaseDeviceReleasedHandlerFunc func(DrmLeaseDeviceReleasedEvent)

// AddReleasedHandler : adds handler for DrmLeaseDeviceReleasedEvent
func (i *DrmLeaseDevice) AddReleasedHandler(f DrmLeaseDeviceReleasedHandlerFunc) {
	if f == nil {
		return
	}

	i.releasedHandlers = append(i.releasedHandlers, f)
}

func (i *DrmLeaseDevice) RemoveReleasedHandler(f DrmLeaseDeviceReleasedHandlerFunc) {
	for j, e := range i.releasedHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.releasedHandlers = append(i.releasedHandlers[:j], i.releasedHandlers[j+1:]...)
			return
		}
	}
}

func (i *DrmLeaseDevice) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.drmFdHandlers) == 0 {
			return
		}
		var e DrmLeaseDeviceDrmFdEvent
		e.Fd = fd
		for _, f := range i.drmFdHandlers {
			f(e)
		}
	case 1:
		if len(i.connectorHandlers) == 0 {
			return
		}
		var e DrmLeaseDeviceConnectorEvent
		l := 0
		e.Id = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*DrmLeaseConnector)
		l += 4
		for _, f := range i.connectorHandlers {
			f(e)
		}
	case 2:
		if len(i.doneHandlers) == 0 {
			return
		}
		var e DrmLeaseDeviceDoneEvent
		for _, f := range i.doneHandlers {
			f(e)
		}
	case 3:
		if len(i.releasedHandlers) == 0 {
			return
		}
		var e DrmLeaseDeviceReleasedEvent
		for _, f := range i.releasedHandlers {
			f(e)
		}
	}
}

// DrmLeaseConnector : a leasable DRM connector
//
// Represents a DRM connector which is available for lease. These objects are
// created via wp_drm_lease_device_v1.connector events, and should be passed
// to lease requests via wp_drm_lease_request_v1.request_connector.
// Immediately after the wp_drm_lease_connector_v1 object is created the
// compositor will send a name, a description, a connector_id and a done
// event. When the description is updated the compositor will send a
// description event followed by a done event.
type DrmLeaseConnector struct {
	client.BaseProxy
	nameHandlers        []DrmLeaseConnectorNameHandlerFunc
	descriptionHandlers []DrmLeaseConnectorDescriptionHandlerFunc
	connectorIdHandlers []DrmLeaseConnectorConnectorIdHandlerFunc
	doneHandlers        []DrmLeaseConnectorDoneHandlerFunc
	withdrawnHandlers   []DrmLeaseConnectorWithdrawnHandlerFunc
}

// NewDrmLeaseConnector : a leasable DRM connector
//
// Represents a DRM connector which is available for lease. These objects are
// created via wp_drm_lease_device_v1.connector events, and should be passed
// to lease requests via wp_drm_lease_request_v1.request_connector.
// Immediately after the wp_drm_lease_connector_v1 object is created the
// compositor will send a name, a description, a connector_id and a done
// event. When the description is updated the compositor will send a
// description event followed by a done event.
func NewDrmLeaseConnector(ctx *client.Context) *DrmLeaseConnector {
	wpDrmLeaseConnectorV1 := &DrmLeaseConnector{}
	ctx.Register(wpDrmLeaseConnectorV1)
	return wpDrmLeaseConnectorV1
}

// Destroy : destroy connector
//
// The client may send this request to indicate that it will not use this
// connector. Clients are encouraged to send this after receiving the
// "withdrawn" event so that the server can release the resources
// associated with this connector offer. Neither existing lease requests
// nor leases will be affected.
//
func (i *DrmLeaseConnector) Destroy() error {
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

// DrmLeaseConnectorNameEvent : name
//
// The compositor sends this event once the connector is created to
// indicate the name of this connector. This will not change for the
// duration of the Wayland session, but is not guaranteed to be consistent
// between sessions.
type DrmLeaseConnectorNameEvent struct {
	Name string
}
type DrmLeaseConnectorNameHandlerFunc func(DrmLeaseConnectorNameEvent)

// AddNameHandler : adds handler for DrmLeaseConnectorNameEvent
func (i *DrmLeaseConnector) AddNameHandler(f DrmLeaseConnectorNameHandlerFunc) {
	if f == nil {
		return
	}

	i.nameHandlers = append(i.nameHandlers, f)
}

func (i *DrmLeaseConnector) RemoveNameHandler(f DrmLeaseConnectorNameHandlerFunc) {
	for j, e := range i.nameHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.nameHandlers = append(i.nameHandlers[:j], i.nameHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseConnectorDescriptionEvent : description
//
// The compositor sends this event once the connector is created to provide
// a human-readable description for this connector, which may be presented
// to the user. The compositor may send this event multiple times over the
// lifetime of this object to reflect changes in the description.
type DrmLeaseConnectorDescriptionEvent struct {
	Description string
}
type DrmLeaseConnectorDescriptionHandlerFunc func(DrmLeaseConnectorDescriptionEvent)

// AddDescriptionHandler : adds handler for DrmLeaseConnectorDescriptionEvent
func (i *DrmLeaseConnector) AddDescriptionHandler(f DrmLeaseConnectorDescriptionHandlerFunc) {
	if f == nil {
		return
	}

	i.descriptionHandlers = append(i.descriptionHandlers, f)
}

func (i *DrmLeaseConnector) RemoveDescriptionHandler(f DrmLeaseConnectorDescriptionHandlerFunc) {
	for j, e := range i.descriptionHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.descriptionHandlers = append(i.descriptionHandlers[:j], i.descriptionHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseConnectorConnectorIdEvent : connector_id
//
// The compositor sends this event once the connector is created to
// indicate the DRM object ID which represents the underlying connector
// that is being offered. Note that the final lease may include additional
// object IDs, such as CRTCs and planes.
type DrmLeaseConnectorConnectorIdEvent struct {
	ConnectorId uint32
}
type DrmLeaseConnectorConnectorIdHandlerFunc func(DrmLeaseConnectorConnectorIdEvent)

// AddConnectorIdHandler : adds handler for DrmLeaseConnectorConnectorIdEvent
func (i *DrmLeaseConnector) AddConnectorIdHandler(f DrmLeaseConnectorConnectorIdHandlerFunc) {
	if f == nil {
		return
	}

	i.connectorIdHandlers = append(i.connectorIdHandlers, f)
}

func (i *DrmLeaseConnector) RemoveConnectorIdHandler(f DrmLeaseConnectorConnectorIdHandlerFunc) {
	for j, e := range i.connectorIdHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.connectorIdHandlers = append(i.connectorIdHandlers[:j], i.connectorIdHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseConnectorDoneEvent : all properties have been sent
//
// This event is sent after all properties of a connector have been sent.
// This allows changes to the properties to be seen as atomic even if they
// happen via multiple events.
type DrmLeaseConnectorDoneEvent struct{}
type DrmLeaseConnectorDoneHandlerFunc func(DrmLeaseConnectorDoneEvent)

// AddDoneHandler : adds handler for DrmLeaseConnectorDoneEvent
func (i *DrmLeaseConnector) AddDoneHandler(f DrmLeaseConnectorDoneHandlerFunc) {
	if f == nil {
		return
	}

	i.doneHandlers = append(i.doneHandlers, f)
}

func (i *DrmLeaseConnector) RemoveDoneHandler(f DrmLeaseConnectorDoneHandlerFunc) {
	for j, e := range i.doneHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.doneHandlers = append(i.doneHandlers[:j], i.doneHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseConnectorWithdrawnEvent : lease offer withdrawn
//
// Sent to indicate that the compositor will no longer honor requests for
// DRM leases which include this connector. The client may still issue a
// lease request including this connector, but the compositor will send
// wp_drm_lease_v1.finished without issuing a lease fd. Compositors are
// encouraged to send this event when they lose access to connector, for
// example when the connector is hot-unplugged, when the connector gets
// leased to a client or when the compositor loses DRM master.
type DrmLeaseConnectorWithdrawnEvent struct{}
type DrmLeaseConnectorWithdrawnHandlerFunc func(DrmLeaseConnectorWithdrawnEvent)

// AddWithdrawnHandler : adds handler for DrmLeaseConnectorWithdrawnEvent
func (i *DrmLeaseConnector) AddWithdrawnHandler(f DrmLeaseConnectorWithdrawnHandlerFunc) {
	if f == nil {
		return
	}

	i.withdrawnHandlers = append(i.withdrawnHandlers, f)
}

func (i *DrmLeaseConnector) RemoveWithdrawnHandler(f DrmLeaseConnectorWithdrawnHandlerFunc) {
	for j, e := range i.withdrawnHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.withdrawnHandlers = append(i.withdrawnHandlers[:j], i.withdrawnHandlers[j+1:]...)
			return
		}
	}
}

func (i *DrmLeaseConnector) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.nameHandlers) == 0 {
			return
		}
		var e DrmLeaseConnectorNameEvent
		l := 0
		nameLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Name = client.String(data[l : l+nameLen])
		l += nameLen
		for _, f := range i.nameHandlers {
			f(e)
		}
	case 1:
		if len(i.descriptionHandlers) == 0 {
			return
		}
		var e DrmLeaseConnectorDescriptionEvent
		l := 0
		descriptionLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Description = client.String(data[l : l+descriptionLen])
		l += descriptionLen
		for _, f := range i.descriptionHandlers {
			f(e)
		}
	case 2:
		if len(i.connectorIdHandlers) == 0 {
			return
		}
		var e DrmLeaseConnectorConnectorIdEvent
		l := 0
		e.ConnectorId = client.Uint32(data[l : l+4])
		l += 4
		for _, f := range i.connectorIdHandlers {
			f(e)
		}
	case 3:
		if len(i.doneHandlers) == 0 {
			return
		}
		var e DrmLeaseConnectorDoneEvent
		for _, f := range i.doneHandlers {
			f(e)
		}
	case 4:
		if len(i.withdrawnHandlers) == 0 {
			return
		}
		var e DrmLeaseConnectorWithdrawnEvent
		for _, f := range i.withdrawnHandlers {
			f(e)
		}
	}
}

// DrmLeaseRequest : DRM lease request
//
// A client that wishes to lease DRM resources will attach the list of
// connectors advertised with wp_drm_lease_device_v1.connector that they
// wish to lease, then use wp_drm_lease_request_v1.submit to submit the
// request.
type DrmLeaseRequest struct {
	client.BaseProxy
}

// NewDrmLeaseRequest : DRM lease request
//
// A client that wishes to lease DRM resources will attach the list of
// connectors advertised with wp_drm_lease_device_v1.connector that they
// wish to lease, then use wp_drm_lease_request_v1.submit to submit the
// request.
func NewDrmLeaseRequest(ctx *client.Context) *DrmLeaseRequest {
	wpDrmLeaseRequestV1 := &DrmLeaseRequest{}
	ctx.Register(wpDrmLeaseRequestV1)
	return wpDrmLeaseRequestV1
}

// RequestConnector : request a connector for this lease
//
// Indicates that the client would like to lease the given connector.
// This is only used as a suggestion, the compositor may choose to
// include any resources in the lease it issues, or change the set of
// leased resources at any time. Compositors are however encouraged to
// include the requested connector and other resources necessary
// to drive the connected output in the lease.
//
// Requesting a connector that was created from a different lease device
// than this lease request raises the wrong_device error. Requesting a
// connector twice will raise the duplicate_connector error.
//
func (i *DrmLeaseRequest) RequestConnector(connector *DrmLeaseConnector) error {
	const opcode = 0
	const rLen = 8 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], connector.ID())
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return err
}

// Submit : submit the lease request
//
// Submits the lease request and creates a new wp_drm_lease_v1 object.
// After calling submit the compositor will immediately destroy this
// object, issuing any more requests will cause a wl_diplay error.
// The compositor doesn't make any guarantees about the events of the
// lease object, clients cannot expect an immediate response.
// Not requesting any connectors before submitting the lease request
// will raise the empty_lease error.
//
func (i *DrmLeaseRequest) Submit() (*DrmLease, error) {
	defer i.Context().Unregister(i)
	id := NewDrmLease(i.Context())
	const opcode = 1
	const rLen = 8 + 4
	r := make([]byte, rLen)
	l := 0
	client.PutUint32(r[l:4], i.ID())
	l += 4
	client.PutUint32(r[l:l+4], uint32(rLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(r[l:l+4], id.ID())
	l += 4
	err := i.Context().WriteMsg(r, nil)
	return id, err
}

type DrmLeaseRequestError uint32

// DrmLeaseRequestError :
const (
	// DrmLeaseRequestErrorWrongDevice : requested a connector from a different lease device
	DrmLeaseRequestErrorWrongDevice DrmLeaseRequestError = 0
	// DrmLeaseRequestErrorDuplicateConnector : requested a connector twice
	DrmLeaseRequestErrorDuplicateConnector DrmLeaseRequestError = 1
	// DrmLeaseRequestErrorEmptyLease : requested a lease without requesting a connector
	DrmLeaseRequestErrorEmptyLease DrmLeaseRequestError = 2
)

func (e DrmLeaseRequestError) Name() string {
	switch e {
	case DrmLeaseRequestErrorWrongDevice:
		return "wrong_device"
	case DrmLeaseRequestErrorDuplicateConnector:
		return "duplicate_connector"
	case DrmLeaseRequestErrorEmptyLease:
		return "empty_lease"
	default:
		return ""
	}
}

func (e DrmLeaseRequestError) Value() string {
	switch e {
	case DrmLeaseRequestErrorWrongDevice:
		return "0"
	case DrmLeaseRequestErrorDuplicateConnector:
		return "1"
	case DrmLeaseRequestErrorEmptyLease:
		return "2"
	default:
		return ""
	}
}

func (e DrmLeaseRequestError) String() string {
	return e.Name() + "=" + e.Value()
}

// DrmLease : a DRM lease
//
// A DRM lease object is used to transfer the DRM file descriptor to the
// client and manage the lifetime of the lease.
//
// Some time after the wp_drm_lease_v1 object is created, the compositor
// will reply with the lease request's result. If the lease request is
// granted, the compositor will send a lease_fd event. If the lease request
// is denied, the compositor will send a finished event without a lease_fd
// event.
type DrmLease struct {
	client.BaseProxy
	leaseFdHandlers  []DrmLeaseLeaseFdHandlerFunc
	finishedHandlers []DrmLeaseFinishedHandlerFunc
}

// NewDrmLease : a DRM lease
//
// A DRM lease object is used to transfer the DRM file descriptor to the
// client and manage the lifetime of the lease.
//
// Some time after the wp_drm_lease_v1 object is created, the compositor
// will reply with the lease request's result. If the lease request is
// granted, the compositor will send a lease_fd event. If the lease request
// is denied, the compositor will send a finished event without a lease_fd
// event.
func NewDrmLease(ctx *client.Context) *DrmLease {
	wpDrmLeaseV1 := &DrmLease{}
	ctx.Register(wpDrmLeaseV1)
	return wpDrmLeaseV1
}

// Destroy : destroys the lease object
//
// The client should send this to indicate that it no longer wishes to use
// this lease. The compositor should use drmModeRevokeLease on the
// appropriate file descriptor, if necessary.
//
func (i *DrmLease) Destroy() error {
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

// DrmLeaseLeaseFdEvent : shares the DRM file descriptor
//
// This event returns a file descriptor suitable for use with DRM-related
// ioctls. The client should use drmModeGetLease to enumerate the DRM
// objects which have been leased to them. The compositor guarantees it
// will not use the leased DRM objects itself until it sends the finished
// event. If the compositor cannot or will not grant a lease for the
// requested connectors, it will not send this event, instead sending the
// finished event.
//
// The compositor will send this event at most once during this objects
// lifetime.
type DrmLeaseLeaseFdEvent struct {
	LeasedFd uintptr
}
type DrmLeaseLeaseFdHandlerFunc func(DrmLeaseLeaseFdEvent)

// AddLeaseFdHandler : adds handler for DrmLeaseLeaseFdEvent
func (i *DrmLease) AddLeaseFdHandler(f DrmLeaseLeaseFdHandlerFunc) {
	if f == nil {
		return
	}

	i.leaseFdHandlers = append(i.leaseFdHandlers, f)
}

func (i *DrmLease) RemoveLeaseFdHandler(f DrmLeaseLeaseFdHandlerFunc) {
	for j, e := range i.leaseFdHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.leaseFdHandlers = append(i.leaseFdHandlers[:j], i.leaseFdHandlers[j+1:]...)
			return
		}
	}
}

// DrmLeaseFinishedEvent : sent when the lease has been revoked
//
// The compositor uses this event to either reject a lease request, or if
// it previously sent a lease_fd, to notify the client that the lease has
// been revoked. If the client requires a new lease, they should destroy
// this object and submit a new lease request. The compositor will send
// no further events for this object after sending the finish event.
// Compositors should revoke the lease when any of the leased resources
// become unavailable, namely when a hot-unplug occurs or when the
// compositor loses DRM master.
type DrmLeaseFinishedEvent struct{}
type DrmLeaseFinishedHandlerFunc func(DrmLeaseFinishedEvent)

// AddFinishedHandler : adds handler for DrmLeaseFinishedEvent
func (i *DrmLease) AddFinishedHandler(f DrmLeaseFinishedHandlerFunc) {
	if f == nil {
		return
	}

	i.finishedHandlers = append(i.finishedHandlers, f)
}

func (i *DrmLease) RemoveFinishedHandler(f DrmLeaseFinishedHandlerFunc) {
	for j, e := range i.finishedHandlers {
		if reflect.ValueOf(e).Pointer() == reflect.ValueOf(f).Pointer() {
			i.finishedHandlers = append(i.finishedHandlers[:j], i.finishedHandlers[j+1:]...)
			return
		}
	}
}

func (i *DrmLease) Dispatch(opcode uint16, fd uintptr, data []byte) {
	switch opcode {
	case 0:
		if len(i.leaseFdHandlers) == 0 {
			return
		}
		var e DrmLeaseLeaseFdEvent
		e.LeasedFd = fd
		for _, f := range i.leaseFdHandlers {
			f(e)
		}
	case 1:
		if len(i.finishedHandlers) == 0 {
			return
		}
		var e DrmLeaseFinishedEvent
		for _, f := range i.finishedHandlers {
			f(e)
		}
	}
}
