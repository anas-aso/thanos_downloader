module github.com/anas-aso/thanos_downloader

go 1.15

require (
	github.com/go-kit/kit v0.10.0
	github.com/oklog/ulid v1.3.1
	github.com/thanos-io/thanos v0.15.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

// requried to avoid issues due to "k8s.io/client-go@v12.0.0+incompatible"
replace k8s.io/client-go => k8s.io/client-go v0.19.1
