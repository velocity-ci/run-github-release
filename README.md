# run-github-release
Provides a Docker image that can create GitHub Releases and upload files 

## Usage

### General
```
docker run --rm \
--env GITHUB_RELEASE_TOKEN=xxxxx \
--volume $(pwd):/app \
--workdir /app \
velocity-ci run-github-release 0.1.0 dist/run-github-release
```

### In Velocity Task

```
---
# task.yml


```