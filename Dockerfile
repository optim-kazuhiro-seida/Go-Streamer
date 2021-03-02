FROM rust:latest as build
ADD . /usr/src/app
WORKDIR /usr/src/app

RUN rustup target add x86_64-unknown-linux-musl
RUN cargo build --release --target x86_64-unknown-linux-musl


FROM alpine:latest AS runtime
WORKDIR /app
COPY --from=build /usr/src/app/target/x86_64-unknown-linux-musl/release/go_struct_generator /app/go_struct_generator
ENV STREAMER_DIRECTORY /file
ENV STREAMER_RECURSION ""
CMD ["/app/go_struct_generator"]
