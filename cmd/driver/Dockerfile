FROM golang

ENV GO111MODULE=on
ENV PORT=6666

RUN apt-get update && apt-get install -y \
  python \
  make \
  g++ \
  build-essential \
  libudev-dev \
  libusb-1.0-0-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=1 go build

EXPOSE $PORT
ENTRYPOINT ["/app/missile-launcher-driver"]