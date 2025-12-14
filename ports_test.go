package netports

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortRange(t *testing.T) {
	t.Parallel()

	ssh, ok := KnownPorts.GroupByNumber()[22]
	assert.True(t, ok)
	assert.Len(t, ssh, 1)

	assert.Equal(
		t,
		"Secure Shell (SSH), secure logins, file transfers (scp, sftp) and port forwarding",
		ssh[0].Description,
	)
}

func ExamplePorts() {
	fmt.Printf("%d %d",
		len(slices.Collect(KnownPorts.Filter(
			FilterByProto(TCP),
			FilterByCategory(CategoryWellKnown, CategoryRegistered),
		))),
		len(slices.Collect(KnownPorts.Filter(
			FilterByProto(UDP),
			FilterByCategory(CategoryWellKnown, CategoryRegistered),
		))),
	)
	// Output: 2853 2448
}
