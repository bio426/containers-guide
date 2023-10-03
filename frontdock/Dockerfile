# BUILD
FROM node:lts-alpine AS builder
WORKDIR /app
COPY . .

# Set env variables
ENV VITE_API_URL="miapisecreta"

# Install the dependencies and build
RUN npm install
RUN npm run build

# SERVE
FROM nginx:stable-alpine

# Remove default nginx files
RUN rm /usr/share/nginx/html/* /etc/nginx/conf.d/default.conf

# Get static and config files
COPY ./nginx.spa.conf /etc/nginx/conf.d/
COPY --from=builder /app/dist/ /usr/share/nginx/html/

EXPOSE 8080

# Run is called by defult nginx image
