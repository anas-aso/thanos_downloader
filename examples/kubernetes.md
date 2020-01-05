# Prometheus Helm Chart

#### Requirements:
- [Thanos config file](https://thanos.io/storage.md/#configuration) in a kubernetes secret deployed to the target namespace:
```Yaml
apiVersion: v1
kind: Secret
metadata:
  name: thanos-config
type: Opaque
data:
  bucket.yaml: <Thanos_Config_File_Base64_Encoded>
```
- Kubernetes cluster access.

#### How-to
Deploy the stable Prometheus Helm Chart using [this example values file](prometheus_values.yml) :

```bash
export namespace="myNamespace"
export releaseName="thanos-downloader"

helm upgrade --install --namespace ${namespace} --wait --values prometheus_values.yml --version 9.7.2 ${releaseName} stable/prometheus
```
