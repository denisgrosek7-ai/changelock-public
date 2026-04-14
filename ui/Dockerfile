FROM node:22-alpine AS build
WORKDIR /app
ARG VITE_API_BASE_URL=/api
ARG VITE_API_TOKEN=
ARG VITE_API_TIMEOUT_MS=8000
ENV VITE_API_BASE_URL=$VITE_API_BASE_URL
ENV VITE_API_TOKEN=$VITE_API_TOKEN
ENV VITE_API_TIMEOUT_MS=$VITE_API_TIMEOUT_MS
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:1.29-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html
