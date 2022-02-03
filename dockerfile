FROM frolvlad/alpine-glibc
WORKDIR /app
COPY assets /app/assets
COPY static /app/static
COPY oauth2  /app/
EXPOSE 8000
ENTRYPOINT ["/app/oauth2"]