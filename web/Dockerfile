FROM node:16-alpine AS builder
WORKDIR /src
COPY package.json ./
COPY tsconfig.json ./
RUN yarn install
COPY ./src ./src
RUN yarn build

FROM nginx:1.21-alpine
COPY --from=builder /src/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 8080