package cluster

import (
	"fmt"
	"net"
	"time"
)

// MemberID represents the ID of a member node in a cluster
type MemberID string

// Member represents a member node in a cluster
type Member interface {

	// ID returns The ID of the member node
	ID() MemberID

	// IP returns the IP address of the member node
	IP() net.IP

	// Port returns the replication port number of the member node
	Port() uint16
}

// NewMember creates a member with the specified IP and port number
func NewMember(ip net.IP, port uint16) Member {
	return &member{
		MemberIP:   ip,
		MemberPort: port,
	}
}

// member is an implementation of the Member interface
type member struct {
	MemberIP   net.IP    `json:"ip,omitempty"`
	MemberPort uint16    `json:"port,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (m *member) ID() MemberID {
	return MemberID(m.String())
}

func (m *member) IP() net.IP {
	return m.MemberIP
}

func (m *member) Port() uint16 {
	return m.MemberPort
}

func (m *member) String() string {
	return fmt.Sprintf("%s:%d", m.MemberIP.String(), m.MemberPort)
}
