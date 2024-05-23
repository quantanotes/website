FROM oven/bun:alpine as js-build
WORKDIR /app
COPY static static
COPY www www
COPY bun.lockb *.js *.json ./
RUN bun install --frozen-lockfile
RUN bun vite build
RUN bun vite build --ssr

FROM golang:1.22-alpine as go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY --from=js-build /app/static /app/static
COPY main.go .
COPY internal internal
RUN go build -o app

FROM oven/bun:alpine
WORKDIR /app
COPY .env .env
COPY --from=go-build /app/app /app/app
COPY --from=js-build /app/ssr /app/ssr

EXPOSE 3000
RUN apk add --no-cache supervisor
COPY supervisord.conf /etc/supervisord.conf
CMD ["supervisord", "-c", "/etc/supervisord.conf"]