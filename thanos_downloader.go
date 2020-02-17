// Copyright 2020 Anas Ait Said Oubrahim

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/anas-aso/thanos_downloader/pkg/blocks"
	"github.com/anas-aso/thanos_downloader/pkg/interval"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/thanos-io/thanos/pkg/block"
	"github.com/thanos-io/thanos/pkg/objstore/client"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	logger := log.NewJSONLogger(log.NewSyncWriter(os.Stderr))

	// Application flags
	downloadDirPath := kingpin.Flag("data.path", "target data directory path.").Default("prom_data").String()
	startTimeSeconds := kingpin.Flag("interval.start", "start time of the requested interval in Unix time format (seconds).").Default("0").Int64()
	endTimeSeconds := kingpin.Flag("interval.end", "end time of the requested interval in Unix time format (seconds).").Default(strconv.FormatInt(time.Now().Unix(), 10)).Int64()
	confPath := kingpin.Flag("config.path", "path to the Thanos format config file.").Default("config.yaml").String()
	kingpin.Version("Thanos Downloader : 0.0.1")
	kingpin.Parse()

	// create a TimeInterval from the provided values
	requestedInterval, err := interval.NewTimeInterval(*startTimeSeconds*1000, *endTimeSeconds*1000)
	if err != nil {
		level.Error(logger).Log("msg", err)
		os.Exit(1)
	}

	config, err := ioutil.ReadFile(*confPath)
	if err != nil {
		level.Error(logger).Log("msg", err)
		os.Exit(1)
	}

	// create a bucket client from the provided config
	bkt, err := client.NewBucket(logger, config, nil, "downloader")
	if err != nil {
		level.Error(logger).Log("msg", err)
		os.Exit(1)
	}

	// stores the blocks ULIDs that verify the requested time range
	blks := blocks.Blocks{}

	ctx := context.Background()

	// get list of blocks that satisfy the requested time range
	err = bkt.Iter(ctx, "", func(dir string) error {
		// check that "dir" is a valid ULID, if not skip it
		id, ok := block.IsBlockDir(dir)
		if !ok {
			return nil
		}

		meta, err := block.DownloadMeta(ctx, logger, bkt, id)
		if err != nil {
			return err
		}

		blockInterval, _ := interval.NewTimeInterval(meta.MinTime, meta.MaxTime)
		if requestedInterval.IntersectWith(blockInterval) {
			blks = append(blks, blocks.NewLightMeta(meta))
			level.Info(logger).Log("msg", "Found block", "blockID", id, "minTime", time.Unix(0, meta.MinTime*int64(time.Millisecond)), "maxTime", time.Unix(0, meta.MaxTime*int64(time.Millisecond)))
		}
		return nil
	})

	if err != nil {
		level.Error(logger).Log("msg", err)
		os.Exit(1)
	}

	if len(blks) == 0 {
		level.Warn(logger).Log("msg", "Couldn't find any block that statisfies the requested time range.")
		os.Exit(0)
	}

	// Drop overlapping blocks due to Prometheus HA Setup or Downsampling
	blks.DropOverlappingBlocks()

	level.Info(logger).Log("msg", fmt.Sprintf("Found %v block(s) that statisfy(statisfies) the requested time range.", len(blks)))

	// Download blocks sequentially
	for _, b := range blks {
		start := time.Now()
		level.Info(logger).Log("msg", "Starting download", "blockID", b.ULID)
		block.Download(ctx, logger, bkt, b.ULID, filepath.Join(*downloadDirPath, b.ULID.String()))
		level.Info(logger).Log("msg", "Finished download", "blockID", b.ULID, "duration", time.Since(start))
	}
}
