# Spider

## Kubernetes の環境選択

```bash
kubectl config get-contexts
```

```bash
kubectl config use-context docker-desktop
```

## Istioセットアップ

```bash
istioctl manifest apply --set profile=demo
```

``` bash
kubectl label namespace default istio-injection=enabled
```

## deploy

```bash
docker build --no-cache -t spider/server server
```

```bash
kubectl apply -f server/manifest.yaml
```

## gateway

```bash
kubectl apply -f gateway.yaml
```
