FROM scratch
ADD check-receiver /

EXPOSE 8080
VOLUME /var/check-receiver
CMD ["/check-receiver"]
