# Using Docker CLI (local setup)

#### Requirements:
- [Thanos config file](https://thanos.io/storage.md/#configuration)
- Enough disk space to store the requested data

#### How-to
After adjusting the variables according to your setup, run the following:

```bash
# Target directory to store the downloaded data
export dataDir=${HOME}/tmp/prom_data

# Path to Thanos config file
export configPath=${HOME}/tmp/config.yaml

# Set interval to 1 hour of data that is older than 2 hours
export intervalEnd=$(expr $(date +%s) - 7200)
export intervalStart=$(expr $(date +%s) - 10800)

# You can drop --interval.start and --interval.end to download all existing data
docker run  --rm -it \
            -v ${configPath}:/config.yaml \
            -v ${dataDir}:/data \
            anasaso/thanos_downloader:latest \
                --config.path /config.yaml \
                --data.path /data \
                --interval.start ${intervalStart} \
                --interval.end ${intervalEnd}

# Create empty Prometheus config file to prevent Prometheus from scrapping new data
promConfig="prometheus_$(date +%s).yml"
touch /tmp/${promConfig}

# Run Prometheus server with the downloaded data
docker run  --rm -it \
            -p 9090:9090 \
            -v /tmp/${promConfig}:/prometheus.yml \
            -v ${dataDir}:/data \
            prom/prometheus:v2.15.1 \
                --storage.tsdb.path=/data \
                --config.file=/prometheus.yml
```
