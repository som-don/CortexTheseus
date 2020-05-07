// Copyright 2020 The CortexTheseus Authors
// This file is part of the CortexTheseus library.
//
// The CortexTheseus library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The CortexTheseus library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the CortexTheseus library. If not, see <http://www.gnu.org/licenses/>.
package torrentfs

import (
	"github.com/CortexFoundation/CortexTheseus/params"
)

// Config ...
type Config struct {
	// Host is the host interface on which to start the storage server. If this
	// field is empty, no storage will be started.
	Port            int      `toml:",omitempty"`
	DataDir         string   `toml:",omitempty"`
	RpcURI          string   `toml:",omitempty"`
	IpcPath         string   `toml:",omitempty"`
	DisableUTP      bool     `toml:",omitempty"`
	DisableTCP      bool     `toml:",omitempty"`
	DisableDHT      bool     `toml:",omitempty"`
	DefaultTrackers []string `toml:",omitempty"`
	BoostNodes      []string `toml:",omitempty"`
	SyncMode        string   `toml:",omitempty"`
	MaxSeedingNum   int      `toml:",omitempty"`
	MaxActiveNum    int      `toml:",omitempty"`
	FullSeed        bool
	Boost           bool
	Quiet           bool
	UploadRate      int
	DownloadRate    int
	Metrics         bool
}

// DefaultConfig contains default settings for the storage.
var DefaultConfig = Config{
	Port:            40401,
	DefaultTrackers: params.MainnetTrackers,
	BoostNodes:      params.TorrentBoostNodes,
	SyncMode:        "full",
	DisableUTP:      true,
	DisableDHT:      false,
	DisableTCP:      false,
	MaxSeedingNum:   1024,
	MaxActiveNum:    1024,
	FullSeed:        false,
	Boost:           false,
	Quiet:           true,
	UploadRate:      -1,
	DownloadRate:    -1,
	Metrics:         true,
}

const (
	queryTimeInterval              = 1
	expansionFactor        float64 = 1.2
	defaultSeedInterval            = 600
	torrentWaitingTime             = 1800
	downloadWaitingTime            = 2700
	defaultBytesLimitation         = 512 * 1024
	defaultTmpFilePath             = ".tmp"
	version                        = "1"
)
