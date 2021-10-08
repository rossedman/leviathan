FROM alpine
COPY leviathan /usr/bin/leviathan
ENTRYPOINT ["/usr/bin/leviathan"]