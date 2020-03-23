FROM localhost:5000/org/base:build-push-test

ENTRYPOINT ["echo", "hello-world build-push-from-registry"]