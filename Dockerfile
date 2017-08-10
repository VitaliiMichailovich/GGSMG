FROM golang:onbuild
COPY main /
COPY client /
COPY templates /
CMD ["/main"]