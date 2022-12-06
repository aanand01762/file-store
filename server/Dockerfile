FROM golang:alpine

WORKDIR /app

# Copy source code  
RUN mkdir store-files
RUN mkdir pkg
COPY pkg pkg 
COPY go.mod go.sum main.go ./

# Build the code 
RUN go build

# Expose server port 
EXPOSE 8080

# Run the executable
CMD ["./server"]
