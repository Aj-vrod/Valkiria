FROM golang:1.19

# Create app directory
WORKDIR /app

# Copy reference to dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build
RUN go build -o ./Valkiria

# Run
CMD ["./Valkiria"]

EXPOSE 3000