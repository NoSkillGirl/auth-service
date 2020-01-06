FROM ubuntu:16.04
RUN apt-get update
COPY auth-service /
CMD ["/auth-service"]
EXPOSE 8083