# Start by building the application.
FROM golang:1.19 as build

WORKDIR /usr/src/traefik-forward-auth
COPY . .

RUN CGO_ENABLED=0 go build -o ./traefik-forward-auth ./cmd

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=build /usr/src/traefik-forward-auth/traefik-forward-auth /usr/bin/traefik-forward-auth

ENTRYPOINT [ "/usr/bin/traefik-forward-auth" ]
CMD []

LABEL org.opencontainers.image.title traefik-forward-auth
LABEL org.opencontainers.image.description "Forward authentication service for the Traefik reverse proxy"
LABEL org.opencontainers.image.licenses MIT
