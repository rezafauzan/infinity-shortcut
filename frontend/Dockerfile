FROM node:24.13.0-alpine AS builder

WORKDIR /workspace
COPY package*.json .
RUN npm install
COPY . .

ARG VITE_BASE_URL
ENV VITE_BASE_URL=$VITE_BASE_URL

RUN npm run build

FROM nginx:1.28.2-alpine

COPY --from=builder /workspace/dist /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf
COPY ./default.conf /etc/nginx/conf.d/default.conf