[![Release](https://img.shields.io/github/release/anas-aso/thanos_downloader.svg?style=flat)](https://github.com/anas-aso/thanos_downloader/releases/latest)
[![Build Status](https://github.com/anas-aso/thanos_downloader/workflows/.github/workflows/test.yml/badge.svg)](https://github.com/anas-aso/thanos_downloader/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/anas-aso/thanos_downloader)](https://goreportcard.com/report/github.com/anas-aso/thanos_downloader)

# Thanos Downloader
A helper application to use Prometheus TSDB data blocks uploaded by [Thanos Sidecar](https://github.com/thanos-io/thanos/blob/master/docs/components/sidecar.md) component with a vanilla Prometheus setup.

The goal of this project is to allow for a simple use of Thanos Sidecar as a backup service for Prometheus data, and be able to use that data with a vanilla Prometheus setup without the need for the remaining Thanos components.

### Configuration
Thanos Downloader uses the same configuration format used by Thanos Sidecar (e.g https://thanos.io/storage.md/#gcs).
The available configuration flags can be found as below :
```bash
$ thanos_downloader --help
usage: thanos_downloader --interval.start=INTERVAL.START --interval.end=INTERVAL.END [<flags>]

Flags:
  --help                            Show context-sensitive help (also try --help-long and --help-man).
  --data.path="prom_data"           target data directory path.
  --interval.start=INTERVAL.START   start time of the requested interval in Unix time format (seconds).
  --interval.end=INTERVAL.END       end time of the requested interval in Unix time format (seconds).
  --config.path="config.yaml"       path to the Thanos format config file.
  --version                         Show application version.
```

### How does it work
1. get list of blocks in the provided bucket
2. filter all blocks that satisfies the requested time range
3. get the block with the highest resolution and most samples count for each time interval. This is needed since the uploaded blocks might be overlapping due to Prometheus HA setups and/or Thanos downsampling.


### Docker
The application is available as a [docker image](https://hub.docker.com/repository/docker/anasaso/thanos_downloader) :
```
docker run --rm -it anasaso/thanos_downloader:latest --help
```

### Examples
- [How-To Docker](examples/docker.md)
- [How-To Kubernetes](examples/kubernetes.md)
