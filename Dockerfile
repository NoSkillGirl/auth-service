FROM ubuntu:16.04
COPY auth-service /auth-service
RUN chmod +x /auth-service
CMD /auth-service
EXPOSE 8083