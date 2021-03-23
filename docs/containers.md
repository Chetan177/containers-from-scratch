## Kubernetes, CRI, kubectl, docker

From [kubernetes documentation](https://kubernetes.io/docs/home/), *"Kubernetes is an open source container orchestration engine for automating deployment, scaling, and management of containerized applications."*

This diagram below shows all the components of a kubernetes cluster.

![Kubernetes Cluster](
images/KubernetesCluster.png?raw=true
"Kubernetes Cluster")

Below is a simplified view of the same but also shows the interface between the `kubelet` and the container runtime on any `node`.

![Simplified k8ss cluster with nodes](
images/Kube-cluster-simplified.png?raw=true
"Simplified k8ss cluster with nodes")0


Container runtimes integrate with `kubelet` by implementing the [CRI](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md).

### Container runtimes

Container runtimes can be logically divided into high level runtime (like [containerd](https://containerd.io/)) & low level runtime (like [runc](https://github.com/opencontainers/runc)) based on their function & features. High level runtimes, like `containerd` implements functions like downloading images, image management, unpacking images, and running containers from them. For runnning containers, they shell out to a low level runtime like `runc`. Low level runtimes, like `runc`, are typically implemented using Linux kernel features such as namespaces & cgroups. They are responsible for "containerizing" the application by creating a cgroup, setting resource limits on it, chrooting the container's process to a root filesystem, unsharing to move the process to it's own namespace and running the application in the created cgroup.

### Docker

Docker is built from packaged versions of vanilla `containerd` and `runc`. A simplified architecture diagram looks like this.

![kubetcl docker interaction](
images/docker.png?raw=true
"kubetcl docker interaction")

A brief description from [high level runtimes](https://www.ianlewis.org/en/container-runtimes-part-3-high-level-runtimes) states,

*"dockerd provides features such as building images, and dockerd uses docker-containerd to provide features such as image management and running containers. For instance, Docker's build step is actually just some logic that interprets a Dockerfile, runs the necessary commands in a container using containerd, and saves the resulting container file system as an image."*

Below is a high level diagram of how `kubectl` integrates with `docker`.

![kubetcl docker interaction](
images/k8s-docker-basic.png?raw=true
"kubetcl docker interaction")
