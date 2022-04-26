FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/beego/bee/v2

CMD ["bee", "run", "--downdoc=true", "--gendoc=true"]