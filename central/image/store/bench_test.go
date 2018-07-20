package store

import (
	"fmt"
	"testing"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/bolthelper"
	"bitbucket.org/stack-rox/apollo/pkg/fixtures"
	"bitbucket.org/stack-rox/apollo/pkg/uuid"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

const maxGRPCSize = 4194304

func getImageStore(b *testing.B) Store {
	db, err := bolthelper.NewTemp(b.Name() + ".db")
	if err != nil {
		b.Fatal(err)
	}
	return New(db)
}

func BenchmarkAddImage(b *testing.B) {
	store := getImageStore(b)
	image := fixtures.GetImage()
	for i := 0; i < b.N; i++ {
		store.AddImage(image)
	}
}

func BenchmarkUpdateImage(b *testing.B) {
	store := getImageStore(b)
	image := fixtures.GetImage()
	for i := 0; i < b.N; i++ {
		store.UpdateImage(image)
	}
}

func BenchmarkGetImage(b *testing.B) {
	store := getImageStore(b)
	image := fixtures.GetImage()
	store.AddImage(image)
	for i := 0; i < b.N; i++ {
		store.GetImage(image.GetName().GetSha())
	}
}

func BenchmarkListImage(b *testing.B) {
	store := getImageStore(b)
	image := fixtures.GetImage()
	store.AddImage(image)
	for i := 0; i < b.N; i++ {
		store.ListImage(image.GetName().GetSha())
	}
}

// This really isn't a benchmark, but just prints out how many ListImages can be returned in an API call
func BenchmarkMaxListImage(b *testing.B) {
	listImage := &v1.ListImage{
		Sha:  uuid.NewDummy().String(),
		Name: "quizzical_cat",
		SetComponents: &v1.ListImage_Components{
			Components: 10,
		},
		SetCves: &v1.ListImage_Cves{
			Cves: 10,
		},
		SetFixable: &v1.ListImage_FixableCves{
			FixableCves: 10,
		},
		Created: types.TimestampNow(),
	}

	bytes, _ := proto.Marshal(listImage)
	fmt.Printf("Max ListImages that can be returned: %d\n", maxGRPCSize/len(bytes))
}
