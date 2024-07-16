#builder stage
FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod .
COPY main.go .
COPY gcp ./gcp
COPY utils ./utils
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

#final stage
FROM google/cloud-sdk:483.0.0-alpine as developer
WORKDIR /app
COPY --from=builder /app/myapp .
COPY XXXXXX.json .
RUN apk --no-cache add ca-certificates tzdata

RUN gcloud auth activate-service-account XXXXXXX@developer.gserviceaccount.com --key-file=XXXXXX.json --project=XXXXXX
RUN gcloud components install kubectl
RUN gcloud components install gke-gcloud-auth-plugin

ENTRYPOINT ["/app/myapp"]

EXPOSE 5432