# Use base golang image from Docker Hub
FROM golang:1.12.7

# Download the dlv (delve) debugger for go (you can comment this out if unused)
RUN go get -u github.com/go-delve/delve/cmd/dlv


RUN mkdir /librdkafka-dir && cd /librdkafka-dir
RUN git clone https://github.com/edenhill/librdkafka.git && \
cd librdkafka && \
./configure --prefix /usr && \
make && \
make install



# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

#copy vendor vendor
# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
#WORKDIR /dist

# Copy binary from build to main folder
#RUN cp /build/main .

# Export necessary port
EXPOSE 9191

# Command to run when starting the container
#CMD ["/build/main --port 8080 --host 0.0.0.0"]
ENTRYPOINT /build/main --port 9191 --host 0.0.0.0


