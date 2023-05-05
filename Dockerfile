FROM gcr.io/distroless/static-debian11
ARG APP
COPY --from=base --chown=1001 ./apps/$APP /run
CMD ["/run"]
