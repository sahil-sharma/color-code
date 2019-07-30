FROM golang:latest
COPY ./code/ /opt/code
WORKDIR /opt/code
RUN go get github.com/gin-gonic/gin \
    && go get gopkg.in/go-playground/colors.v1 \
    && go build color.go
EXPOSE 8080
CMD ["bash", "-c", "/opt/code/color"]
