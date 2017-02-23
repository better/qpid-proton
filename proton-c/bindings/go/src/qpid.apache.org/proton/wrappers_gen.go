/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

//
// NOTE: DO NOT EDIT. This file was generated by genwrap.go from the proton header files.
// Update the generator and re-run if you need to modify this code.
//

package proton

import (
	"time"
	"unsafe"
)

// #include <proton/types.h>
// #include <proton/error.h>
// #include <proton/condition.h>
// #include <proton/event.h>
// #include <stdlib.h>
// #include <proton/session.h>
// #include <proton/link.h>
// #include <proton/delivery.h>
// #include <proton/disposition.h>
// #include <proton/condition.h>
// #include <proton/terminus.h>
// #include <proton/connection.h>
// #include <proton/transport.h>
// #include <proton/sasl.h>
import "C"

type EventType int

const (
	EConnectionInit         EventType = C.PN_CONNECTION_INIT
	EConnectionBound        EventType = C.PN_CONNECTION_BOUND
	EConnectionUnbound      EventType = C.PN_CONNECTION_UNBOUND
	EConnectionLocalOpen    EventType = C.PN_CONNECTION_LOCAL_OPEN
	EConnectionRemoteOpen   EventType = C.PN_CONNECTION_REMOTE_OPEN
	EConnectionLocalClose   EventType = C.PN_CONNECTION_LOCAL_CLOSE
	EConnectionRemoteClose  EventType = C.PN_CONNECTION_REMOTE_CLOSE
	EConnectionFinal        EventType = C.PN_CONNECTION_FINAL
	ESessionInit            EventType = C.PN_SESSION_INIT
	ESessionLocalOpen       EventType = C.PN_SESSION_LOCAL_OPEN
	ESessionRemoteOpen      EventType = C.PN_SESSION_REMOTE_OPEN
	ESessionLocalClose      EventType = C.PN_SESSION_LOCAL_CLOSE
	ESessionRemoteClose     EventType = C.PN_SESSION_REMOTE_CLOSE
	ESessionFinal           EventType = C.PN_SESSION_FINAL
	ELinkInit               EventType = C.PN_LINK_INIT
	ELinkLocalOpen          EventType = C.PN_LINK_LOCAL_OPEN
	ELinkRemoteOpen         EventType = C.PN_LINK_REMOTE_OPEN
	ELinkLocalClose         EventType = C.PN_LINK_LOCAL_CLOSE
	ELinkRemoteClose        EventType = C.PN_LINK_REMOTE_CLOSE
	ELinkLocalDetach        EventType = C.PN_LINK_LOCAL_DETACH
	ELinkRemoteDetach       EventType = C.PN_LINK_REMOTE_DETACH
	ELinkFlow               EventType = C.PN_LINK_FLOW
	ELinkFinal              EventType = C.PN_LINK_FINAL
	EDelivery               EventType = C.PN_DELIVERY
	ETransport              EventType = C.PN_TRANSPORT
	ETransportAuthenticated EventType = C.PN_TRANSPORT_AUTHENTICATED
	ETransportError         EventType = C.PN_TRANSPORT_ERROR
	ETransportHeadClosed    EventType = C.PN_TRANSPORT_HEAD_CLOSED
	ETransportTailClosed    EventType = C.PN_TRANSPORT_TAIL_CLOSED
	ETransportClosed        EventType = C.PN_TRANSPORT_CLOSED
	EConnectionWake         EventType = C.PN_CONNECTION_WAKE
	EListenerAccept         EventType = C.PN_LISTENER_ACCEPT
	EListenerClose          EventType = C.PN_LISTENER_CLOSE
	EProactorInterrupt      EventType = C.PN_PROACTOR_INTERRUPT
	EProactorTimeout        EventType = C.PN_PROACTOR_TIMEOUT
	EProactorInactive       EventType = C.PN_PROACTOR_INACTIVE
	EListenerOpen           EventType = C.PN_LISTENER_OPEN
)

func (e EventType) String() string {
	switch e {

	case C.PN_CONNECTION_INIT:
		return "ConnectionInit"
	case C.PN_CONNECTION_BOUND:
		return "ConnectionBound"
	case C.PN_CONNECTION_UNBOUND:
		return "ConnectionUnbound"
	case C.PN_CONNECTION_LOCAL_OPEN:
		return "ConnectionLocalOpen"
	case C.PN_CONNECTION_REMOTE_OPEN:
		return "ConnectionRemoteOpen"
	case C.PN_CONNECTION_LOCAL_CLOSE:
		return "ConnectionLocalClose"
	case C.PN_CONNECTION_REMOTE_CLOSE:
		return "ConnectionRemoteClose"
	case C.PN_CONNECTION_FINAL:
		return "ConnectionFinal"
	case C.PN_SESSION_INIT:
		return "SessionInit"
	case C.PN_SESSION_LOCAL_OPEN:
		return "SessionLocalOpen"
	case C.PN_SESSION_REMOTE_OPEN:
		return "SessionRemoteOpen"
	case C.PN_SESSION_LOCAL_CLOSE:
		return "SessionLocalClose"
	case C.PN_SESSION_REMOTE_CLOSE:
		return "SessionRemoteClose"
	case C.PN_SESSION_FINAL:
		return "SessionFinal"
	case C.PN_LINK_INIT:
		return "LinkInit"
	case C.PN_LINK_LOCAL_OPEN:
		return "LinkLocalOpen"
	case C.PN_LINK_REMOTE_OPEN:
		return "LinkRemoteOpen"
	case C.PN_LINK_LOCAL_CLOSE:
		return "LinkLocalClose"
	case C.PN_LINK_REMOTE_CLOSE:
		return "LinkRemoteClose"
	case C.PN_LINK_LOCAL_DETACH:
		return "LinkLocalDetach"
	case C.PN_LINK_REMOTE_DETACH:
		return "LinkRemoteDetach"
	case C.PN_LINK_FLOW:
		return "LinkFlow"
	case C.PN_LINK_FINAL:
		return "LinkFinal"
	case C.PN_DELIVERY:
		return "Delivery"
	case C.PN_TRANSPORT:
		return "Transport"
	case C.PN_TRANSPORT_AUTHENTICATED:
		return "TransportAuthenticated"
	case C.PN_TRANSPORT_ERROR:
		return "TransportError"
	case C.PN_TRANSPORT_HEAD_CLOSED:
		return "TransportHeadClosed"
	case C.PN_TRANSPORT_TAIL_CLOSED:
		return "TransportTailClosed"
	case C.PN_TRANSPORT_CLOSED:
		return "TransportClosed"
	case C.PN_CONNECTION_WAKE:
		return "ConnectionWake"
	case C.PN_LISTENER_ACCEPT:
		return "ListenerAccept"
	case C.PN_LISTENER_CLOSE:
		return "ListenerClose"
	case C.PN_PROACTOR_INTERRUPT:
		return "ProactorInterrupt"
	case C.PN_PROACTOR_TIMEOUT:
		return "ProactorTimeout"
	case C.PN_PROACTOR_INACTIVE:
		return "ProactorInactive"
	case C.PN_LISTENER_OPEN:
		return "ListenerOpen"
	}
	return "Unknown"
}

// Wrappers for declarations in session.h

type Session struct{ pn *C.pn_session_t }

func (s Session) IsNil() bool          { return s.pn == nil }
func (s Session) CPtr() unsafe.Pointer { return unsafe.Pointer(s.pn) }
func (s Session) Free() {
	C.pn_session_free(s.pn)
}
func (s Session) State() State {
	return State(C.pn_session_state(s.pn))
}
func (s Session) Error() error {
	return PnError(C.pn_session_error(s.pn))
}
func (s Session) Condition() Condition {
	return Condition{C.pn_session_condition(s.pn)}
}
func (s Session) RemoteCondition() Condition {
	return Condition{C.pn_session_remote_condition(s.pn)}
}
func (s Session) Connection() Connection {
	return Connection{C.pn_session_connection(s.pn)}
}
func (s Session) Open() {
	C.pn_session_open(s.pn)
}
func (s Session) Close() {
	C.pn_session_close(s.pn)
}
func (s Session) IncomingCapacity() uint {
	return uint(C.pn_session_get_incoming_capacity(s.pn))
}
func (s Session) SetIncomingCapacity(capacity uint) {
	C.pn_session_set_incoming_capacity(s.pn, C.size_t(capacity))
}
func (s Session) OutgoingWindow() uint {
	return uint(C.pn_session_get_outgoing_window(s.pn))
}
func (s Session) SetOutgoingWindow(window uint) {
	C.pn_session_set_outgoing_window(s.pn, C.size_t(window))
}
func (s Session) OutgoingBytes() uint {
	return uint(C.pn_session_outgoing_bytes(s.pn))
}
func (s Session) IncomingBytes() uint {
	return uint(C.pn_session_incoming_bytes(s.pn))
}
func (s Session) Next(state State) Session {
	return Session{C.pn_session_next(s.pn, C.pn_state_t(state))}
}

// Wrappers for declarations in link.h

type SndSettleMode C.pn_snd_settle_mode_t

const (
	SndUnsettled SndSettleMode = C.PN_SND_UNSETTLED
	SndSettled   SndSettleMode = C.PN_SND_SETTLED
	SndMixed     SndSettleMode = C.PN_SND_MIXED
)

func (e SndSettleMode) String() string {
	switch e {

	case C.PN_SND_UNSETTLED:
		return "SndUnsettled"
	case C.PN_SND_SETTLED:
		return "SndSettled"
	case C.PN_SND_MIXED:
		return "SndMixed"
	}
	return "unknown"
}

type RcvSettleMode C.pn_rcv_settle_mode_t

const (
	RcvFirst  RcvSettleMode = C.PN_RCV_FIRST
	RcvSecond RcvSettleMode = C.PN_RCV_SECOND
)

func (e RcvSettleMode) String() string {
	switch e {

	case C.PN_RCV_FIRST:
		return "RcvFirst"
	case C.PN_RCV_SECOND:
		return "RcvSecond"
	}
	return "unknown"
}

type Link struct{ pn *C.pn_link_t }

func (l Link) IsNil() bool          { return l.pn == nil }
func (l Link) CPtr() unsafe.Pointer { return unsafe.Pointer(l.pn) }
func (l Link) Free() {
	C.pn_link_free(l.pn)
}
func (l Link) Name() string {
	return C.GoString(C.pn_link_name(l.pn))
}
func (l Link) IsSender() bool {
	return bool(C.pn_link_is_sender(l.pn))
}
func (l Link) IsReceiver() bool {
	return bool(C.pn_link_is_receiver(l.pn))
}
func (l Link) State() State {
	return State(C.pn_link_state(l.pn))
}
func (l Link) Error() error {
	return PnError(C.pn_link_error(l.pn))
}
func (l Link) Condition() Condition {
	return Condition{C.pn_link_condition(l.pn)}
}
func (l Link) RemoteCondition() Condition {
	return Condition{C.pn_link_remote_condition(l.pn)}
}
func (l Link) Session() Session {
	return Session{C.pn_link_session(l.pn)}
}
func (l Link) Next(state State) Link {
	return Link{C.pn_link_next(l.pn, C.pn_state_t(state))}
}
func (l Link) Open() {
	C.pn_link_open(l.pn)
}
func (l Link) Close() {
	C.pn_link_close(l.pn)
}
func (l Link) Detach() {
	C.pn_link_detach(l.pn)
}
func (l Link) Source() Terminus {
	return Terminus{C.pn_link_source(l.pn)}
}
func (l Link) Target() Terminus {
	return Terminus{C.pn_link_target(l.pn)}
}
func (l Link) RemoteSource() Terminus {
	return Terminus{C.pn_link_remote_source(l.pn)}
}
func (l Link) RemoteTarget() Terminus {
	return Terminus{C.pn_link_remote_target(l.pn)}
}
func (l Link) Current() Delivery {
	return Delivery{C.pn_link_current(l.pn)}
}
func (l Link) Advance() bool {
	return bool(C.pn_link_advance(l.pn))
}
func (l Link) Credit() int {
	return int(C.pn_link_credit(l.pn))
}
func (l Link) Queued() int {
	return int(C.pn_link_queued(l.pn))
}
func (l Link) RemoteCredit() int {
	return int(C.pn_link_remote_credit(l.pn))
}
func (l Link) Drained() int {
	return int(C.pn_link_drained(l.pn))
}
func (l Link) Available() int {
	return int(C.pn_link_available(l.pn))
}
func (l Link) SndSettleMode() SndSettleMode {
	return SndSettleMode(C.pn_link_snd_settle_mode(l.pn))
}
func (l Link) RcvSettleMode() RcvSettleMode {
	return RcvSettleMode(C.pn_link_rcv_settle_mode(l.pn))
}
func (l Link) SetSndSettleMode(mode SndSettleMode) {
	C.pn_link_set_snd_settle_mode(l.pn, C.pn_snd_settle_mode_t(mode))
}
func (l Link) SetRcvSettleMode(mode RcvSettleMode) {
	C.pn_link_set_rcv_settle_mode(l.pn, C.pn_rcv_settle_mode_t(mode))
}
func (l Link) RemoteSndSettleMode() SndSettleMode {
	return SndSettleMode(C.pn_link_remote_snd_settle_mode(l.pn))
}
func (l Link) RemoteRcvSettleMode() RcvSettleMode {
	return RcvSettleMode(C.pn_link_remote_rcv_settle_mode(l.pn))
}
func (l Link) Unsettled() int {
	return int(C.pn_link_unsettled(l.pn))
}
func (l Link) Offered(credit int) {
	C.pn_link_offered(l.pn, C.int(credit))
}
func (l Link) Flow(credit int) {
	C.pn_link_flow(l.pn, C.int(credit))
}
func (l Link) Drain(credit int) {
	C.pn_link_drain(l.pn, C.int(credit))
}
func (l Link) SetDrain(drain bool) {
	C.pn_link_set_drain(l.pn, C.bool(drain))
}
func (l Link) Draining() bool {
	return bool(C.pn_link_draining(l.pn))
}
func (l Link) MaxMessageSize() uint64 {
	return uint64(C.pn_link_max_message_size(l.pn))
}
func (l Link) SetMaxMessageSize(size uint64) {
	C.pn_link_set_max_message_size(l.pn, C.uint64_t(size))
}
func (l Link) RemoteMaxMessageSize() uint64 {
	return uint64(C.pn_link_remote_max_message_size(l.pn))
}

// Wrappers for declarations in delivery.h

type Delivery struct{ pn *C.pn_delivery_t }

func (d Delivery) IsNil() bool          { return d.pn == nil }
func (d Delivery) CPtr() unsafe.Pointer { return unsafe.Pointer(d.pn) }
func (d Delivery) Tag() DeliveryTag {
	return DeliveryTag{C.pn_delivery_tag(d.pn)}
}
func (d Delivery) Link() Link {
	return Link{C.pn_delivery_link(d.pn)}
}
func (d Delivery) Local() Disposition {
	return Disposition{C.pn_delivery_local(d.pn)}
}
func (d Delivery) LocalState() uint64 {
	return uint64(C.pn_delivery_local_state(d.pn))
}
func (d Delivery) Remote() Disposition {
	return Disposition{C.pn_delivery_remote(d.pn)}
}
func (d Delivery) RemoteState() uint64 {
	return uint64(C.pn_delivery_remote_state(d.pn))
}
func (d Delivery) Settled() bool {
	return bool(C.pn_delivery_settled(d.pn))
}
func (d Delivery) Pending() uint {
	return uint(C.pn_delivery_pending(d.pn))
}
func (d Delivery) Partial() bool {
	return bool(C.pn_delivery_partial(d.pn))
}
func (d Delivery) Writable() bool {
	return bool(C.pn_delivery_writable(d.pn))
}
func (d Delivery) Readable() bool {
	return bool(C.pn_delivery_readable(d.pn))
}
func (d Delivery) Updated() bool {
	return bool(C.pn_delivery_updated(d.pn))
}
func (d Delivery) Update(state uint64) {
	C.pn_delivery_update(d.pn, C.uint64_t(state))
}
func (d Delivery) Clear() {
	C.pn_delivery_clear(d.pn)
}
func (d Delivery) Current() bool {
	return bool(C.pn_delivery_current(d.pn))
}
func (d Delivery) Settle() {
	C.pn_delivery_settle(d.pn)
}
func (d Delivery) Dump() {
	C.pn_delivery_dump(d.pn)
}
func (d Delivery) Buffered() bool {
	return bool(C.pn_delivery_buffered(d.pn))
}

// Wrappers for declarations in disposition.h

type Disposition struct{ pn *C.pn_disposition_t }

func (d Disposition) IsNil() bool          { return d.pn == nil }
func (d Disposition) CPtr() unsafe.Pointer { return unsafe.Pointer(d.pn) }
func (d Disposition) Type() uint64 {
	return uint64(C.pn_disposition_type(d.pn))
}
func (d Disposition) Condition() Condition {
	return Condition{C.pn_disposition_condition(d.pn)}
}
func (d Disposition) Data() Data {
	return Data{C.pn_disposition_data(d.pn)}
}
func (d Disposition) SectionNumber() uint16 {
	return uint16(C.pn_disposition_get_section_number(d.pn))
}
func (d Disposition) SetSectionNumber(section_number uint16) {
	C.pn_disposition_set_section_number(d.pn, C.uint32_t(section_number))
}
func (d Disposition) SectionOffset() uint64 {
	return uint64(C.pn_disposition_get_section_offset(d.pn))
}
func (d Disposition) SetSectionOffset(section_offset uint64) {
	C.pn_disposition_set_section_offset(d.pn, C.uint64_t(section_offset))
}
func (d Disposition) IsFailed() bool {
	return bool(C.pn_disposition_is_failed(d.pn))
}
func (d Disposition) SetFailed(failed bool) {
	C.pn_disposition_set_failed(d.pn, C.bool(failed))
}
func (d Disposition) IsUndeliverable() bool {
	return bool(C.pn_disposition_is_undeliverable(d.pn))
}
func (d Disposition) SetUndeliverable(undeliverable bool) {
	C.pn_disposition_set_undeliverable(d.pn, C.bool(undeliverable))
}
func (d Disposition) Annotations() Data {
	return Data{C.pn_disposition_annotations(d.pn)}
}

// Wrappers for declarations in condition.h

type Condition struct{ pn *C.pn_condition_t }

func (c Condition) IsNil() bool          { return c.pn == nil }
func (c Condition) CPtr() unsafe.Pointer { return unsafe.Pointer(c.pn) }
func (c Condition) IsSet() bool {
	return bool(C.pn_condition_is_set(c.pn))
}
func (c Condition) Clear() {
	C.pn_condition_clear(c.pn)
}
func (c Condition) Name() string {
	return C.GoString(C.pn_condition_get_name(c.pn))
}
func (c Condition) SetName(name string) int {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	return int(C.pn_condition_set_name(c.pn, nameC))
}
func (c Condition) Description() string {
	return C.GoString(C.pn_condition_get_description(c.pn))
}
func (c Condition) SetDescription(description string) int {
	descriptionC := C.CString(description)
	defer C.free(unsafe.Pointer(descriptionC))

	return int(C.pn_condition_set_description(c.pn, descriptionC))
}
func (c Condition) Info() Data {
	return Data{C.pn_condition_info(c.pn)}
}
func (c Condition) IsRedirect() bool {
	return bool(C.pn_condition_is_redirect(c.pn))
}
func (c Condition) RedirectHost() string {
	return C.GoString(C.pn_condition_redirect_host(c.pn))
}
func (c Condition) RedirectPort() int {
	return int(C.pn_condition_redirect_port(c.pn))
}
func (c Condition) Copy(src Condition) int {
	return int(C.pn_condition_copy(c.pn, src.pn))
}

// Wrappers for declarations in terminus.h

type TerminusType C.pn_terminus_type_t

const (
	Unspecified TerminusType = C.PN_UNSPECIFIED
	Source      TerminusType = C.PN_SOURCE
	Target      TerminusType = C.PN_TARGET
	Coordinator TerminusType = C.PN_COORDINATOR
)

func (e TerminusType) String() string {
	switch e {

	case C.PN_UNSPECIFIED:
		return "Unspecified"
	case C.PN_SOURCE:
		return "Source"
	case C.PN_TARGET:
		return "Target"
	case C.PN_COORDINATOR:
		return "Coordinator"
	}
	return "unknown"
}

type Durability C.pn_durability_t

const (
	Nondurable    Durability = C.PN_NONDURABLE
	Configuration Durability = C.PN_CONFIGURATION
	Deliveries    Durability = C.PN_DELIVERIES
)

func (e Durability) String() string {
	switch e {

	case C.PN_NONDURABLE:
		return "Nondurable"
	case C.PN_CONFIGURATION:
		return "Configuration"
	case C.PN_DELIVERIES:
		return "Deliveries"
	}
	return "unknown"
}

type ExpiryPolicy C.pn_expiry_policy_t

const (
	ExpireWithLink       ExpiryPolicy = C.PN_EXPIRE_WITH_LINK
	ExpireWithSession    ExpiryPolicy = C.PN_EXPIRE_WITH_SESSION
	ExpireWithConnection ExpiryPolicy = C.PN_EXPIRE_WITH_CONNECTION
	ExpireNever          ExpiryPolicy = C.PN_EXPIRE_NEVER
)

func (e ExpiryPolicy) String() string {
	switch e {

	case C.PN_EXPIRE_WITH_LINK:
		return "ExpireWithLink"
	case C.PN_EXPIRE_WITH_SESSION:
		return "ExpireWithSession"
	case C.PN_EXPIRE_WITH_CONNECTION:
		return "ExpireWithConnection"
	case C.PN_EXPIRE_NEVER:
		return "ExpireNever"
	}
	return "unknown"
}

type DistributionMode C.pn_distribution_mode_t

const (
	DistModeUnspecified DistributionMode = C.PN_DIST_MODE_UNSPECIFIED
	DistModeCopy        DistributionMode = C.PN_DIST_MODE_COPY
	DistModeMove        DistributionMode = C.PN_DIST_MODE_MOVE
)

func (e DistributionMode) String() string {
	switch e {

	case C.PN_DIST_MODE_UNSPECIFIED:
		return "DistModeUnspecified"
	case C.PN_DIST_MODE_COPY:
		return "DistModeCopy"
	case C.PN_DIST_MODE_MOVE:
		return "DistModeMove"
	}
	return "unknown"
}

type Terminus struct{ pn *C.pn_terminus_t }

func (t Terminus) IsNil() bool          { return t.pn == nil }
func (t Terminus) CPtr() unsafe.Pointer { return unsafe.Pointer(t.pn) }
func (t Terminus) Type() TerminusType {
	return TerminusType(C.pn_terminus_get_type(t.pn))
}
func (t Terminus) SetType(type_ TerminusType) int {
	return int(C.pn_terminus_set_type(t.pn, C.pn_terminus_type_t(type_)))
}
func (t Terminus) Address() string {
	return C.GoString(C.pn_terminus_get_address(t.pn))
}
func (t Terminus) SetAddress(address string) int {
	addressC := C.CString(address)
	defer C.free(unsafe.Pointer(addressC))

	return int(C.pn_terminus_set_address(t.pn, addressC))
}
func (t Terminus) SetDistributionMode(mode DistributionMode) int {
	return int(C.pn_terminus_set_distribution_mode(t.pn, C.pn_distribution_mode_t(mode)))
}
func (t Terminus) Durability() Durability {
	return Durability(C.pn_terminus_get_durability(t.pn))
}
func (t Terminus) SetDurability(durability Durability) int {
	return int(C.pn_terminus_set_durability(t.pn, C.pn_durability_t(durability)))
}
func (t Terminus) ExpiryPolicy() ExpiryPolicy {
	return ExpiryPolicy(C.pn_terminus_get_expiry_policy(t.pn))
}
func (t Terminus) SetExpiryPolicy(policy ExpiryPolicy) int {
	return int(C.pn_terminus_set_expiry_policy(t.pn, C.pn_expiry_policy_t(policy)))
}
func (t Terminus) Timeout() time.Duration {
	return (time.Duration(C.pn_terminus_get_timeout(t.pn)) * time.Second)
}
func (t Terminus) SetTimeout(timeout time.Duration) int {
	return int(C.pn_terminus_set_timeout(t.pn, C.pn_seconds_t(timeout/time.Second)))
}
func (t Terminus) IsDynamic() bool {
	return bool(C.pn_terminus_is_dynamic(t.pn))
}
func (t Terminus) SetDynamic(dynamic bool) int {
	return int(C.pn_terminus_set_dynamic(t.pn, C.bool(dynamic)))
}
func (t Terminus) Properties() Data {
	return Data{C.pn_terminus_properties(t.pn)}
}
func (t Terminus) Capabilities() Data {
	return Data{C.pn_terminus_capabilities(t.pn)}
}
func (t Terminus) Outcomes() Data {
	return Data{C.pn_terminus_outcomes(t.pn)}
}
func (t Terminus) Filter() Data {
	return Data{C.pn_terminus_filter(t.pn)}
}
func (t Terminus) Copy(src Terminus) int {
	return int(C.pn_terminus_copy(t.pn, src.pn))
}

// Wrappers for declarations in connection.h

type Connection struct{ pn *C.pn_connection_t }

func (c Connection) IsNil() bool          { return c.pn == nil }
func (c Connection) CPtr() unsafe.Pointer { return unsafe.Pointer(c.pn) }
func (c Connection) Free() {
	C.pn_connection_free(c.pn)
}
func (c Connection) Release() {
	C.pn_connection_release(c.pn)
}
func (c Connection) Error() error {
	return PnError(C.pn_connection_error(c.pn))
}
func (c Connection) State() State {
	return State(C.pn_connection_state(c.pn))
}
func (c Connection) Open() {
	C.pn_connection_open(c.pn)
}
func (c Connection) Close() {
	C.pn_connection_close(c.pn)
}
func (c Connection) Reset() {
	C.pn_connection_reset(c.pn)
}
func (c Connection) Condition() Condition {
	return Condition{C.pn_connection_condition(c.pn)}
}
func (c Connection) RemoteCondition() Condition {
	return Condition{C.pn_connection_remote_condition(c.pn)}
}
func (c Connection) Container() string {
	return C.GoString(C.pn_connection_get_container(c.pn))
}
func (c Connection) SetContainer(container string) {
	containerC := C.CString(container)
	defer C.free(unsafe.Pointer(containerC))

	C.pn_connection_set_container(c.pn, containerC)
}
func (c Connection) SetUser(user string) {
	userC := C.CString(user)
	defer C.free(unsafe.Pointer(userC))

	C.pn_connection_set_user(c.pn, userC)
}
func (c Connection) User() string {
	return C.GoString(C.pn_connection_get_user(c.pn))
}
func (c Connection) Hostname() string {
	return C.GoString(C.pn_connection_get_hostname(c.pn))
}
func (c Connection) SetHostname(hostname string) {
	hostnameC := C.CString(hostname)
	defer C.free(unsafe.Pointer(hostnameC))

	C.pn_connection_set_hostname(c.pn, hostnameC)
}
func (c Connection) RemoteContainer() string {
	return C.GoString(C.pn_connection_remote_container(c.pn))
}
func (c Connection) RemoteHostname() string {
	return C.GoString(C.pn_connection_remote_hostname(c.pn))
}
func (c Connection) OfferedCapabilities() Data {
	return Data{C.pn_connection_offered_capabilities(c.pn)}
}
func (c Connection) DesiredCapabilities() Data {
	return Data{C.pn_connection_desired_capabilities(c.pn)}
}
func (c Connection) Properties() Data {
	return Data{C.pn_connection_properties(c.pn)}
}
func (c Connection) RemoteOfferedCapabilities() Data {
	return Data{C.pn_connection_remote_offered_capabilities(c.pn)}
}
func (c Connection) RemoteDesiredCapabilities() Data {
	return Data{C.pn_connection_remote_desired_capabilities(c.pn)}
}
func (c Connection) RemoteProperties() Data {
	return Data{C.pn_connection_remote_properties(c.pn)}
}
func (c Connection) Transport() Transport {
	return Transport{C.pn_connection_transport(c.pn)}
}

// Wrappers for declarations in transport.h

type Transport struct{ pn *C.pn_transport_t }

func (t Transport) IsNil() bool          { return t.pn == nil }
func (t Transport) CPtr() unsafe.Pointer { return unsafe.Pointer(t.pn) }
func (t Transport) SetServer() {
	C.pn_transport_set_server(t.pn)
}
func (t Transport) Free() {
	C.pn_transport_free(t.pn)
}
func (t Transport) User() string {
	return C.GoString(C.pn_transport_get_user(t.pn))
}
func (t Transport) RequireAuth(required bool) {
	C.pn_transport_require_auth(t.pn, C.bool(required))
}
func (t Transport) IsAuthenticated() bool {
	return bool(C.pn_transport_is_authenticated(t.pn))
}
func (t Transport) RequireEncryption(required bool) {
	C.pn_transport_require_encryption(t.pn, C.bool(required))
}
func (t Transport) IsEncrypted() bool {
	return bool(C.pn_transport_is_encrypted(t.pn))
}
func (t Transport) Condition() Condition {
	return Condition{C.pn_transport_condition(t.pn)}
}
func (t Transport) Error() error {
	return PnError(C.pn_transport_error(t.pn))
}
func (t Transport) Bind(connection Connection) int {
	return int(C.pn_transport_bind(t.pn, connection.pn))
}
func (t Transport) Unbind() int {
	return int(C.pn_transport_unbind(t.pn))
}
func (t Transport) Log(message string) {
	messageC := C.CString(message)
	defer C.free(unsafe.Pointer(messageC))

	C.pn_transport_log(t.pn, messageC)
}
func (t Transport) ChannelMax() uint32 {
	return uint32(C.pn_transport_get_channel_max(t.pn))
}
func (t Transport) SetChannelMax(channel_max uint32) int {
	return int(C.pn_transport_set_channel_max(t.pn, C.uint16_t(channel_max)))
}
func (t Transport) RemoteChannelMax() uint32 {
	return uint32(C.pn_transport_remote_channel_max(t.pn))
}
func (t Transport) MaxFrame() uint16 {
	return uint16(C.pn_transport_get_max_frame(t.pn))
}
func (t Transport) SetMaxFrame(size uint16) {
	C.pn_transport_set_max_frame(t.pn, C.uint32_t(size))
}
func (t Transport) RemoteMaxFrame() uint16 {
	return uint16(C.pn_transport_get_remote_max_frame(t.pn))
}
func (t Transport) IdleTimeout() time.Duration {
	return (time.Duration(C.pn_transport_get_idle_timeout(t.pn)) * time.Millisecond)
}
func (t Transport) SetIdleTimeout(timeout time.Duration) {
	C.pn_transport_set_idle_timeout(t.pn, C.pn_millis_t(timeout/time.Millisecond))
}
func (t Transport) RemoteIdleTimeout() time.Duration {
	return (time.Duration(C.pn_transport_get_remote_idle_timeout(t.pn)) * time.Millisecond)
}
func (t Transport) Input(bytes string, available uint) int {
	bytesC := C.CString(bytes)
	defer C.free(unsafe.Pointer(bytesC))

	return int(C.pn_transport_input(t.pn, bytesC, C.size_t(available)))
}
func (t Transport) Output(bytes string, size uint) int {
	bytesC := C.CString(bytes)
	defer C.free(unsafe.Pointer(bytesC))

	return int(C.pn_transport_output(t.pn, bytesC, C.size_t(size)))
}
func (t Transport) Capacity() int {
	return int(C.pn_transport_capacity(t.pn))
}
func (t Transport) Process(size uint) int {
	return int(C.pn_transport_process(t.pn, C.size_t(size)))
}
func (t Transport) CloseTail() int {
	return int(C.pn_transport_close_tail(t.pn))
}
func (t Transport) Pending() int {
	return int(C.pn_transport_pending(t.pn))
}
func (t Transport) Peek(dst string, size uint) int {
	dstC := C.CString(dst)
	defer C.free(unsafe.Pointer(dstC))

	return int(C.pn_transport_peek(t.pn, dstC, C.size_t(size)))
}
func (t Transport) Pop(size uint) {
	C.pn_transport_pop(t.pn, C.size_t(size))
}
func (t Transport) CloseHead() int {
	return int(C.pn_transport_close_head(t.pn))
}
func (t Transport) Quiesced() bool {
	return bool(C.pn_transport_quiesced(t.pn))
}
func (t Transport) Closed() bool {
	return bool(C.pn_transport_closed(t.pn))
}
func (t Transport) Tick(now time.Time) time.Time {
	return goTime(C.pn_transport_tick(t.pn, pnTime(now)))
}
func (t Transport) Connection() Connection {
	return Connection{C.pn_transport_connection(t.pn)}
}

// Wrappers for declarations in sasl.h

type SASLOutcome C.pn_sasl_outcome_t

const (
	SASLNone SASLOutcome = C.PN_SASL_NONE
	SASLOk   SASLOutcome = C.PN_SASL_OK
	SASLAuth SASLOutcome = C.PN_SASL_AUTH
	SASLSys  SASLOutcome = C.PN_SASL_SYS
	SASLPerm SASLOutcome = C.PN_SASL_PERM
	SASLTemp SASLOutcome = C.PN_SASL_TEMP
)

func (e SASLOutcome) String() string {
	switch e {

	case C.PN_SASL_NONE:
		return "SASLNone"
	case C.PN_SASL_OK:
		return "SASLOk"
	case C.PN_SASL_AUTH:
		return "SASLAuth"
	case C.PN_SASL_SYS:
		return "SASLSys"
	case C.PN_SASL_PERM:
		return "SASLPerm"
	case C.PN_SASL_TEMP:
		return "SASLTemp"
	}
	return "unknown"
}

type SASL struct{ pn *C.pn_sasl_t }

func (s SASL) IsNil() bool          { return s.pn == nil }
func (s SASL) CPtr() unsafe.Pointer { return unsafe.Pointer(s.pn) }
func (s SASL) Done(outcome SASLOutcome) {
	C.pn_sasl_done(s.pn, C.pn_sasl_outcome_t(outcome))
}
func (s SASL) Outcome() SASLOutcome {
	return SASLOutcome(C.pn_sasl_outcome(s.pn))
}
func (s SASL) User() string {
	return C.GoString(C.pn_sasl_get_user(s.pn))
}
func (s SASL) Mech() string {
	return C.GoString(C.pn_sasl_get_mech(s.pn))
}
func (s SASL) AllowedMechs(mechs string) {
	mechsC := C.CString(mechs)
	defer C.free(unsafe.Pointer(mechsC))

	C.pn_sasl_allowed_mechs(s.pn, mechsC)
}
func (s SASL) SetAllowInsecureMechs(insecure bool) {
	C.pn_sasl_set_allow_insecure_mechs(s.pn, C.bool(insecure))
}
func (s SASL) AllowInsecureMechs() bool {
	return bool(C.pn_sasl_get_allow_insecure_mechs(s.pn))
}
func (s SASL) ConfigName(name string) {
	nameC := C.CString(name)
	defer C.free(unsafe.Pointer(nameC))

	C.pn_sasl_config_name(s.pn, nameC)
}
func (s SASL) ConfigPath(path string) {
	pathC := C.CString(path)
	defer C.free(unsafe.Pointer(pathC))

	C.pn_sasl_config_path(s.pn, pathC)
}
