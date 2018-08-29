# run-github-release
Provides a Docker image that can create GitHub Releases and upload files 

## Usage

### General
```
docker run --rm \
--env GITHUB_ACCESS_TOKEN=xxxxx \
--volume $(pwd):/app \
--workdir /app \
civelocity/run.github-release \
velocity-ci run-github-release 0.1.0 dist/run-github-release
```

### In Velocity Task

```
---
# task.yml
steps:
  - description: Upload Release
    type: run
    image: civelocity/run.github-release
    environment:
      GITHUB_ACCESS_TOKEN: ${github_release_token}
    command: velocity-ci parameter.aws-ssm ${GIT_DESCRIBE} dist/aws-ssm
```