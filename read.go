// Copyright 2017 the goauth contributors
// See CONTRIBUTORS file for list of names.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package config provides config parsing related functions
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/burntsushi/toml"
)

// ReadEnv looks up the environment variables provided by the caller
// in the dirvar and filevar arguments and passes their values to
// ReadFromDir for further processing.
// The filenames containted in the filevar environment variable should be
// separated with commas (i.e. ',' char).
func ReadEnv(cnf interface{}, dirvar, filevar string) error {
	dir := os.Getenv(dirvar)
	if dir == "" {
		return fmt.Errorf("config.ReadEnv: dirvar %s has emtpy value", dirvar)
	}
	files := os.Getenv(filevar)
	if files == "" {
		return fmt.Errorf("config.ReadEnv: filevar %s has emtpy value", filevar)
	}
	filelist := strings.Split(files, ",")
	return ReadFromDir(cnf, dir, filelist...)
}

// ReadFromDir joins the given filenames with the given dir name and then
// hands those to ReadConfigFiles for parsing.
func ReadFromDir(cnf interface{}, dir string, filenames ...string) error {
	for i, filename := range filenames {
		filenames[i] = filepath.Join(dir, filename)
	}

	return ReadFiles(cnf, filenames...)
}

// ReadFiles tries to parse each of filenames (one after the other in the order stated) into the cnf
// data structure, aborting on the first encountered error and returning it.
// Since data is parsed into the same struct over and over again, the last config file to specify a value
// for any given (sub) setting takes precedence.
func ReadFiles(cnf interface{}, filenames ...string) error {
	for _, filename := range filenames {
		if _, err := toml.DecodeFile(filename, cnf); err != nil {
			return fmt.Errorf("%s: %s", filename, err)
		}
	}

	if pp, ok := cnf.(PostProcessor); ok {
		if err := pp.PostProcess(); err != nil {
			return err
		}
	}

	if s, ok := cnf.(Sanitizer); ok {
		if err := s.Sanitize(); err != nil {
			return err
		}
	}

	if v, ok := cnf.(Validator); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
