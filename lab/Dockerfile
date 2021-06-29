# Base Container
FROM golang:1.16-stretch as build

ENV CGO_ENABLED=0 \
    GO111MODULE=on
RUN go env -w GOPROXY=direct

WORKDIR /quizdeck
COPY . .
RUN ls
RUN go mod download && \
    go build -o ./build/quizdeck ./lab && \
    chmod +x ./build/quizdeck

# App Container
FROM public.ecr.aws/lambda/provided:al2

COPY --from=build /quizdeck/build/quizdeck /quizdeck

RUN adduser -S rain && chown rain /quizdeck
USER rain

ENTRYPOINT [ "/quizdeck" ]
