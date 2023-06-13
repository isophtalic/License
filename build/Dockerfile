FROM golang:1.20 AS build
RUN mkdir /config
WORKDIR /config
COPY go.mod go.sum ./
# COPY ../go.mod ../go.sum ./
RUN go mod download
RUN mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go generate && go build -o backend cmd/main.go

FROM ubuntu:22.04 AS runtime
RUN mkdir /app
RUN mkdir /logs
WORKDIR /app
COPY build/app.env.default /app/conf/app.env
# COPY /app.env.default /app/conf/app.env
COPY build/cmd/* /app/
COPY --from=build /app/backend .
RUN bash -c "chmod +x ./*.sh && ls && ./backend info"