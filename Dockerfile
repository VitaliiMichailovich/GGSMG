FROM golang:onbuild
COPY main /
COPY client/ client/
COPY templates/ templates/
CMD ["/main"]