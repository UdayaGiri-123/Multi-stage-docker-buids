FROM ubuntu AS build

RUN apt-get update && apt-get install -y golang-go

ENV GO111MODULE=off

COPY . .

RUN CGO_ENABLED=0 go build -o /app .

### Multi stage build 
## Scratch is a minimalistic distroless image
FROM scratch

##copy the files generated from build stage
COPY --from=build /app /app

ENTRYPOINT ["/app"]