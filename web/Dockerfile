FROM node:20-alpine
WORKDIR /app/frontend/

# パッケージをインストール
COPY ./package*.json ./
RUN npm ci

# ソースコードをコピーしてビルド
COPY . .
RUN npm run build

# /app/build/ に できたソースをおいておく
CMD [ "cp", "./build/", "/app/build" ]