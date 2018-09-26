package updater

import (
	"github.com/mcuadros/go-version"
)

const (
	OSWindows = "windows"
)

const (
	Arch386   = "386"
	ArchAmd64 = "amd64"
)

// Package represents a file available for download.
type Package struct {
	URL  string `json:"url"`
	Size string `json:"size"`
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

// DeltaPackage represents a file containing only the changes between two version.
type DeltaPackage struct {
	Package
	Version string `json:"version"`
}

// Version represents a single version available for installation.
type Version struct {
	Version       string          `json:"version"`
	ReleaseDate   string          `json:"release_date"`
	Changelog     string          `json:"changelog"`
	Packages      []*Package      `json:"packages"`
	DeltaPackages []*DeltaPackage `json:"delta_packages"`
}

// Manifest represents a list of versions available for download.
type Manifest []*Version

// Len determines the number of versions in the manifest.
func (m Manifest) Len() int {
	return len(m)
}

// Less determines if version i should be sorted before version j.
func (m Manifest) Less(i, j int) bool {
	return version.Compare(m[i].Version, m[j].Version, "<")
}

// Swap swaps the two versions.
func (m Manifest) Swap(i, j int) {
	m[i].Version, m[j].Version = m[j].Version, m[i].Version
}
