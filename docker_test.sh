DOCKER_GO_VERSION="1.13"

for DOCKER_GO_VERSION in 1.1{1,2,3}; do 
  echo "## Docker Golang Version: ${DOCKER_GO_VERSION}"
  docker run \
    -ti \
    --rm \
    -v "$(pwd):/go/github.com/koshatul/openpgp-hash-change" \
    --workdir "/go/github.com/koshatul/openpgp-hash-change" \
    golang:${DOCKER_GO_VERSION} \
    bash -c 'rm test*; go run main.go; md5sum test*.key'
done
