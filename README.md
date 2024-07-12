# Multi-stage-docker-buids
How to reduce size of docker image using distroless

### Distroless images
These images are minimalistic images which have only particular runtime. They serve as base images while creating docker images.
Using these images gives us highest level of security.
### Example
For running a python application we use python distroless image. Some common shell commands like ls, wget wont work in this case.

![image](https://github.com/user-attachments/assets/58f66e42-76e9-45a5-8a53-33e75d070b15)

In above image , simplecalculator2 is an multi stage docker build and with size **1.98MB **
where as simplecalculator1 is not and it has ubuntu as base image with golang installed , hence its size is 686MB
