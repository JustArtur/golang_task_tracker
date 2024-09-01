FROM golang:1.23

WORKDIR /task_tracker_app/

COPY . .

EXPOSE 8080:8080
RUN go install github.com/zzwx/fresh@latest

CMD ["fresh"]

