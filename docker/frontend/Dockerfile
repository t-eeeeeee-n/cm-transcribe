FROM node:20.9-bullseye-slim

# Set the working directory in the container
WORKDIR /app

# bun
RUN apt-get update && apt-get install -y curl unzip
RUN curl -fsSL https://bun.sh/install | bash
ENV PATH="/root/.bun/bin:${PATH}"
COPY ./frontend/package.json ./frontend/bun.lockb /app/
RUN bun add @next/swc-linux-x64-gnu @next/swc-linux-x64-musl
RUN #bun install

# Copy the current directory contents into the container at /app
COPY ./frontend/ /app/