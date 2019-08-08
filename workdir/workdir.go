package workdir

import (
	"os"
	"path/filepath"
)

const defaultMode = 0755

type cleanup func() error

// WorkDir is a response type that contains the Path to a directory created by this package.
type WorkDir struct {
	// Path is the absolute path to the directory requested.
	Path string
	// Cleanup is a function that will cleanup any directory and files under
	// Path.
	Cleanup cleanup
}

type dirType int

const (
	// DirTypeConfig is a directory used to store configuration. This is commonly
	// used to store application configs like yaml or json files used during
	// the bootstrapping phase of an application startup.
	DirTypeConfig dirType = iota
	// DirTypeCache is a directory used to store any temporary cache that is generated
	// by the application. This directory type can be used to store tokens when logging in,
	// or serialized cache such as large responses that you don't want to have to request again
	// for some amount of time. Anything in here should be considered temporary and can be removed
	// at any time.
	DirTypeCache
	// DirTypeData is a directory to store long term data. This data can be database files or assets
	// that need to later be extracted out into another location. Things in this directory should be
	// considered important and backed up.
	DirTypeData
)

// dirTypeEnvDefault is a type that represents the environment variable name
// and its default if it's not set.
type dirTypeEnvDefault struct {
	// envName is the name of the environment variable we should check for first.
	envName string
	// defaultLoc is the default location for the directory type. `default` is a
	// keyword in the Go language, so we use defaultLoc here to prevent syntax
	// collisions.
	defaultLoc string
}

var dirTypeEnv = map[dirType]dirTypeEnvDefault{
	DirTypeConfig: dirTypeEnvDefault{
		envName:    "XDG_CONFIG_HOME",
		defaultLoc: filepath.Join(os.Getenv("HOME"), ".config"),
	},
	DirTypeCache: dirTypeEnvDefault{
		envName:    "XDG_CACHE_HOME",
		defaultLoc: filepath.Join(os.Getenv("HOME"), ".cache"),
	},
	DirTypeData: dirTypeEnvDefault{
		envName:    "XDG_DATA_HOME",
		defaultLoc: filepath.Join(os.Getenv("HOME"), ".local", "share"),
	},
}

// New returns a new WorkDir or an error. If p is not empty, then it will attempt
// to create the directory contained with in it. If p is empty, then it will use dirType
// and namespace to determine what directory should be created and passed back to the caller.
// A standard cleanup function is made available so the caller can decide if they want to
// remove the directory created after they are done. Options allow additional control over
// the directory attributes.
func New(p string, dirType dirType, namespace []string, opts Options) (*WorkDir, error) {
	if p == "" {
		def := dirTypeEnv[dirType]

		p = filepath.Join(def.defaultLoc, filepath.Join(namespace...))

		if os.Getenv(def.envName) != "" {
			p = filepath.Join(os.Getenv(def.envName), filepath.Join(namespace...))
		}
	}

	mode := os.FileMode(defaultMode)
	if opts.Mode != 0 {
		mode = opts.Mode
	}

	if err := os.MkdirAll(p, mode); err != nil {
		return nil, err
	}

	wd := &WorkDir{
		Path: p,
		Cleanup: func() error {
			return os.RemoveAll(p)
		},
	}

	return wd, nil
}
