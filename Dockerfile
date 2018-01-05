FROM scratch
COPY main .
EXPOSE 8285
CMD ["./main"]
