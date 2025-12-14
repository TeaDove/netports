// Package netports
// Well known ports parsed from https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
package netports

import (
	_ "embed"
	"encoding/json"
)

//go:embed ports.json
var portsJson []byte

type Port struct {
	// Start
	// Port number range start and end. Single port is represented as [port, port], i.e. {22, 22}
	// multiple ports are represented as [min, max] included, i.e. {2001, 2009}
	Start uint16 `json:"start"`
	End   uint16 `json:"end"`

	Category    PortCategory `json:"category"`
	Description string       `json:"description"`

	Types map[PortProto]RegistrationStatus `json:"types"`
}

type Ports []Port

// KnownPorts
// List of known ports from https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers
var KnownPorts Ports

func init() { //nolint: gochecknoinits // Too lazy too rewrite using gogenerate
	err := json.Unmarshal(portsJson, &KnownPorts)
	if err != nil {
		panic(err)
	}
}

type PortCategory string

const (
	// CategoryWellKnown (0-1024).
	CategoryWellKnown PortCategory = "WellKnown"
	// CategoryRegistered (1024-49151).
	CategoryRegistered PortCategory = "Registered"
	// CategoryOther (49152-65535).
	CategoryOther PortCategory = "Other"
)

type PortProto string

const (
	TCP  PortProto = "tcp"
	UDP  PortProto = "udp"
	SCTP PortProto = "sctp"
	DCCP PortProto = "dccp"
)

type RegistrationStatus string

const (
	// RegistrationYes
	// Described protocol is assigned by IANA for this port, and is: standardized, specified, or widely used for such.
	RegistrationYes RegistrationStatus = "Yes"
	// RegistrationUnofficial
	// Described protocol is not assigned by IANA for this port, but is: standardized, specified, or widely used for such.
	RegistrationUnofficial RegistrationStatus = "Unofficial"
	// RegistrationAssigned
	// Described protocol is assigned by IANA for this port, but is not: standardized, specified, or widely used for such.
	RegistrationAssigned RegistrationStatus = "Assigned"
	// RegistrationNo
	// Described protocol is not assigned by IANA for this port, standardized, specified, or widely used for such.
	RegistrationNo RegistrationStatus = "No"
	// RegistrationReserved
	// Port is reserved by IANA, generally to prevent collision having its previous use removed.
	// The port number may be available for assignment upon request to IANA.
	RegistrationReserved RegistrationStatus = "Reserved"
)
