# Multi-stage-docker-buids
How to reduce size of docker image using distroless

### Distroless images
These images are minimalistic images which have only particular runtime. They serve as base images while creating docker images.
Using these images gives us highest level of security.
### Example
For running a python application we use python distroless image. Some common shell commands like ls, wget wont work in this case.
