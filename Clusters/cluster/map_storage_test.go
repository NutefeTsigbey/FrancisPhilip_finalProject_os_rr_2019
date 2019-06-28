
package cluster_test

import (
	"testing"

	"github.com/tsuru/docker-cluster/cluster"
	storageTesting "github.com/tsuru/docker-cluster/storage/testing"
)

func TestMapStorageStorage(t *testing.T) {
	storageTesting.RunTestsForStorage(&cluster.MapStorage{}, t)
}
