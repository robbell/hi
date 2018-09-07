FROM scratch
ADD main main
ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/main"]