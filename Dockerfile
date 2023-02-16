FROM golang:1.19.5

WORKDIR /reindeer
COPY . /reindeer/

RUN go build
RUN chmod +x ./reindeer

CMD [ "./reindeer" ]