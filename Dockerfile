FROM scratch
EXPOSE 8080
ENTRYPOINT ["/go-chaos"]
COPY ./bin/ /