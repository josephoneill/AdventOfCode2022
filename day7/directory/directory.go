package directory

import (
	"github.com/AdventOfCode2022/day7/file"
)

type Directory struct {
	Name            string
	FullPath        string
	Directories     []*Directory
	ParentDirectory *Directory
	Files           []file.File
}

func New() Directory {
	d := Directory{Directories: []*Directory{}, Files: []file.File{}}
	return d
}

func (d *Directory) AddDirectory(dir *Directory) {
	d.Directories = append(d.Directories, dir)
}

func (d *Directory) AddFile(f file.File) {
	d.Files = append(d.Files, f)
}

// FindDirectory and DirectoryAlreadyExists are the same thing essentially
func (d *Directory) FindDirectory(name string) *Directory {
	if d == nil {
		return &Directory{}
	}

	for _, val := range d.Directories {
		if val.Name == name {
			return val
		}
	}

	return &Directory{}
}

func (d *Directory) DirectoryAlreadyExists(name string) bool {
	if d == nil {
		return false
	}

	for _, val := range d.Directories {
		if val.Name == name {
			return true
		}
	}

	return false
}

func (d *Directory) TotalDirectoryFileSize() float64 {
	fileSizeTotal := 0.0
	for _, val := range d.Files {
		fileSizeTotal += val.Size
	}

	for _, val := range d.Directories {
		fileSizeTotal += val.TotalDirectoryFileSize()
	}

	return fileSizeTotal
}
