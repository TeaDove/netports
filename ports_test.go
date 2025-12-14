package netports

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortRange(t *testing.T) {
	t.Parallel()

	ssh, ok := KnownPorts.GroupByProto(TCP)[22]
	assert.True(t, ok)
	assert.Len(t, ssh, 1)

	assert.Equal(
		t,
		"Secure Shell (SSH), secure logins, file transfers (scp, sftp) and port forwarding",
		ssh[0].Description,
	)
}

func ExamplePorts() {
	fmt.Printf("%d %d", len(KnownPorts.GroupByProto(TCP)), len(KnownPorts.GroupByProto(UDP)))
	// Output: 18015 13036
}
