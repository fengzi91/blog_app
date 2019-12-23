FROM gobuffalo/buffalo:v0.13.5 as builder

RUN mkdir -p $GOPATH/src/github.com/fengzi91/blog_app
WORKDIR $GOPATH/src/github.com/fengzi91/blog_app

# this will cache the npm install step, unless package.json changes
ADD package.json .
ADD yarn.lock .
RUN yarn install --no-progress
ADD . .
ENV GO111MODULES=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/
#RUN go get ./...
RUN go get $(go list ./... | grep -v /vendor/)
RUN buffalo build --static -o /bin/app

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .

# Uncomment to run the binary in "production" mode:
# ENV GO_ENV=production
#ENV GO_ENV=development
# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

# Uncomment to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
CMD exec /bin/app