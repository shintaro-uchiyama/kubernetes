FROM mysql:8.0.15
LABEL maintainer="shintaro.0112@gmial.com"

# 初期実行ファイル展開
ADD ./init/ /docker-entrypoint-initdb.d

# 設定ファイルを配置
# (MySQLは設定ファイルの権限が777だと読み込まない)
COPY ./conf/charset.cnf /etc/mysql/conf.d/charset.cnf
RUN chmod 644 /etc/mysql/conf.d/*

RUN apt-get update && \
  apt-get install -y locales && \
  rm -rf /var/lib/apt/lists/* && \
  echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
  locale-gen ja_JP.UTF-8
ENV LC_ALL ja_JP.UTF-8