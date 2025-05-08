docker run -d \
  --name jenkins-core-agent \
  -p 2222:22 \
  -v /var/run/docker.sock:/var/run/docker.sock \
  jenkins-core-agent:latest
