# Create image based on the Skaffolder GO image
FROM skaffolder/golang-base 

# Copy source files
WORKDIR /go/src/skaffolder/unity_project
COPY . .

# Install dependencies
RUN go get
RUN go build -o main . 

CMD ["/go/src/skaffolder/unity_project/main"]
