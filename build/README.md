Docker Build Image for Platform Linux/amd64
```bash
docker buildx build --platform=linux/amd64 -f build/Dockerfile -t coderbillzay/jio-health-linux .
```

Docker Build Image for AMD (M1 Mac):
```bash
docker build -f build/Dockerfile -t coderbillzay/jio-health .
```