// Package noise is an opinionated, easy-to-use P2P network stack for decentralized applications, and
// cryptographic protocols written in Go.
//
// noise is made to be minimal, robust, developer-friendly, performant, secure, and cross-platform across
// multitudes of devices by making use of a small amount of well-tested, production-grade dependencies.
package noise

// Handler is called whenever a node receives data from either an inbound/outbound peer connection. Multiple handlers
// may be registered to a node by (*Node).Handle before the node starts listening for new peers.
//
// Returning an error in a handler closes the connection and marks the connection to have closed unexpectedly or due
// to error. Should you intend to wish to skip a handler from processing some given data, return a nil error.
type Handler func(ctx HandlerContext) error

// Protocol is an interface that may be implemented by libraries and projects built on top of Noise to hook callbacks
// onto a series of events that are emitted throughout a nodes lifecycle. They may be registered to a node by
// (*Node).Bind before the node starts listening for new peers.
type Protocol struct {
	// VersionMajor, VersionMinor, and VersionPatch mark the version of this protocol with respect to semantic
	// versioning.
	VersionMajor, VersionMinor, VersionPatch uint

	// Bind is called when the node has successfully started listening for new peers. Important node information
	// such as the nodes binding host, binding port, public address, and ID are not initialized until after
	// (*Node).Listen has successfully been called. Bind gets called the very moment such information has successfully
	// been initialized.
	//
	// Errors returned from implementations of Bind will propagate back up to (*Node).Listen as a returned error.
	Bind func(node *Node) error
}
