docker run --rm -d \
    --name asynqmon \
    -p 8080:8080 \
    hibiken/asynqmon --redis-addr=127.0.0.1:6379 --redis-password=123456