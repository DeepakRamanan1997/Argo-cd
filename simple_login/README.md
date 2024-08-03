Deployment.yaml file deploys a namespace, a load balancer service, a "login successful" application from private docker repo and secrets containing credentials for dockerhub

.dockerconfigjson: in deployment.yaml file consists of details of docker credentials which is extracted using this command "echo '{"auths":{"https://index.docker.io/v1/":{"username":"deepakrehfjnan","password":"9994419560"}}}' | base64 -w 0" and then added in deployment file in secrets section
