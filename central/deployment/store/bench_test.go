package store

import (
	"fmt"
	"testing"

	"bitbucket.org/stack-rox/apollo/central/ranking"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/bolthelper"
	"bitbucket.org/stack-rox/apollo/pkg/fixtures"
	"bitbucket.org/stack-rox/apollo/pkg/uuid"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

const maxGRPCSize = 4194304

func getDeploymentStore(b *testing.B) Store {
	db, err := bolthelper.NewTemp(b.Name() + ".db")
	if err != nil {
		b.Fatal(err)
	}
	return New(db, ranking.NewRanker())
}

func BenchmarkAddDeployment(b *testing.B) {
	store := getDeploymentStore(b)
	deployment := fixtures.GetAlert().GetDeployment()
	for i := 0; i < b.N; i++ {
		store.AddDeployment(deployment)
	}
}

func BenchmarkUpdateDeployment(b *testing.B) {
	store := getDeploymentStore(b)
	deployment := fixtures.GetAlert().GetDeployment()
	for i := 0; i < b.N; i++ {
		store.UpdateDeployment(deployment)
	}
}

func BenchmarkGetDeployment(b *testing.B) {
	store := getDeploymentStore(b)
	deployment := fixtures.GetAlert().GetDeployment()
	store.AddDeployment(deployment)
	for i := 0; i < b.N; i++ {
		store.GetDeployment(deployment.GetId())
	}
}

func BenchmarkListDeployment(b *testing.B) {
	store := getDeploymentStore(b)
	deployment := fixtures.GetAlert().GetDeployment()
	store.AddDeployment(deployment)
	for i := 0; i < b.N; i++ {
		store.ListDeployment(deployment.GetId())
	}
}

// This really isn't a benchmark, but just prints out how many ListDeployments can be returned in an API call
func BenchmarkListDeployments(b *testing.B) {
	listDeployment := &v1.ListDeployment{
		Id:        uuid.NewDummy().String(),
		Name:      "quizzical_cat",
		Cluster:   "Production k8s",
		Namespace: "stackrox",
		UpdatedAt: types.TimestampNow(),
		Priority:  10,
	}

	bytes, _ := proto.Marshal(listDeployment)
	fmt.Printf("Max ListDeployments that can be returned: %d\n", maxGRPCSize/len(bytes))
}
