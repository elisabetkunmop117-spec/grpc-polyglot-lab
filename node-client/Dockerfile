
FROM node:22-alpine

WORKDIR /app

COPY node-client/package.json node-client/package-lock.json ./

RUN npm ci --omit=dev

COPY node-client/ ./

COPY user.proto /user.proto

ENV GRPC_SERVER=go-server:50051

CMD ["node", "client.js"]
