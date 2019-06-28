
package mongodb

import (
	"testing"

	storageTesting "github.com/tsuru/docker-cluster/storage/testing"
)

func TestMongodbStorage(t *testing.T) {
	mongo, err := Mongodb("mongodb://localhost", "test-docker-cluster")
	if err != nil {
		t.Fatal(err)
	}
	stor := mongo.(*mongodbStorage)
	err = stor.session.DB("test-docker-cluster").DropDatabase()
	if err != nil {
		t.Fatal(err)
	}
	storageTesting.RunTestsForStorage(mongo, t)
}
