FROM golang:1.21.0-alpine AS builder

COPY go.mod go.sum /github.com/Maxorella/NIRS/
WORKDIR /github.com/Maxorella/NIRS/

RUN go mod download

COPY . .

#RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/main/main.go

FROM scratch AS runner

WORKDIR /docker-nirs/

COPY --from=builder /github.com/Maxorella/NIRS/.bin .
COPY --from=builder /github.com/Maxorella/NIRS/config config/

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ="Europe/Moscow"
ENV ZONEINFO=/zoneinfo.zip

EXPOSE 80 443

ENTRYPOINT ["./.bin"]