package utils

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/containerd/continuity/fs"

	mongogridfs "kloudlite.io/pkg/mongo-gridfs"
)

func CreateNodeWorkDir(nodeId string) error {
	dir := path.Join(Workdir, nodeId)
	if _, err := os.Stat(dir); err != nil {
		return os.Mkdir(dir, os.ModePerm)
	}

	if enableClear {
		if err := os.RemoveAll(dir); err != nil {
			return err
		}

		return os.Mkdir(dir, os.ModePerm)
	} else {
		return nil
	}
}

func SetupGetWorkDir() error {
	if _, err := os.Stat(Workdir); err != nil {
		return os.Mkdir(Workdir, os.ModePerm)
	}
	return nil
}

func MakeTfWorkFileReady(ctx context.Context, nodeId, tfPath string, gfs mongogridfs.GridFs, createIfNotExists bool) error {
	filename := fmt.Sprintf("%s.zip", nodeId)
	// check if file exists in db
	gf, err := gfs.FetchFileRef(ctx, filename)
	if err != nil {
		return err
	}

	// not found create new dir
	if gf == nil {
		if !createIfNotExists {
			return fmt.Errorf("no state file found with the nodeId %s to operate", nodeId)
		}

		if err := CreateNodeWorkDir(nodeId); err != nil {
			return err
		}

		// a.tfTemplates
		if err := fs.CopyDir(path.Join(Workdir, nodeId), tfPath); err != nil {
			return err
		}

		return nil
	}

	// found file in db, download and extract to the workdir
	fmt.Println(gf.Name, "found, extract it by downloading")

	source := path.Join(Workdir, filename)
	// Download from db
	if err := gfs.Download(ctx, filename, source); err != nil {
		return err
	}

	if s, err := Unzip(source, path.Join(Workdir)); err != nil {
		return err
	} else {
		for _, v := range s {
			fmt.Print(v, " \n")
		}
	}

	return nil
}

func SaveToDb(ctx context.Context, nodeId string, gfs mongogridfs.GridFs) error {
	/*
		Steps:
		  - compress the workdir into zip
		  - check if file present. if yes, upsert file else upload file
	*/

	dir := path.Join(Workdir, nodeId)
	filename := fmt.Sprintf("%s.zip", nodeId)

	// compress the workdir and upsert to db
	if err := func() error {
		if _, err := os.Stat(dir); err != nil {
			return err
		}

		source := fmt.Sprintf("%s.zip", dir)

		// compress
		if err := ZipSource(dir, source); err != nil {
			return err
		}

		if err := gfs.Upsert(ctx, filename, source); err != nil {
			return err
		}

		return nil
	}(); err != nil {
		fmt.Println(ColorText(fmt.Sprint("Error: ", err), 1))
		return err
	}

	return nil
}

const (
	enableClear bool = false
)
