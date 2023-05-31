package mongogridfs

import "go.mongodb.org/mongo-driver/mongo/gridfs"

type gfs struct {
	bucket *gridfs.Bucket
}

type GridFs interface {
	Upload() error
	Download() error
	Delete() error
	Fetch() error
	Search() error
}

// Delete implements GridFs
func (*gfs) Delete() error {
	panic("unimplemented")
}

// Download implements GridFs
func (*gfs) Download() error {
	panic("unimplemented")
}

// Fetch implements GridFs
func (*gfs) Fetch() error {
	panic("unimplemented")
}

// Search implements GridFs
func (*gfs) Search() error {
	panic("unimplemented")
}

// Upload implements GridFs
func (*gfs) Upload() error {
	panic("unimplemented")
}
