# ---------- Frontend Build Stage ----------
FROM node:20-alpine AS frontend-builder

WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install

COPY frontend .
RUN npm run build


# ---------- Backend Build Stage ----------
FROM golang:1.25-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server


# ---------- Final Runtime Image ----------
FROM alpine:latest

WORKDIR /app

# copy compiled backend
COPY --from=backend-builder /app/server .

# copy frontend build output
COPY --from=frontend-builder /frontend/dist ./frontend/dist

EXPOSE 8080

CMD ["./server"]