# 概要
以下手順を記載
- ローカル環境構築手順
- ローカル環境での開発手順

# ローカル環境
## クラスタ（Minikube）起動
```zsh
$ rm -rf ~/.minikube
$ minikube stop
$ minikube delete
$ minikube start --vm-driver=hyperkit
$ minikube status
```

## k8s環境構築
### tls鍵生成
自分の好きなドメイン名と有効期限で作成

```zsh
$ mkdir -p docker/go/tls
$
$ cd docker/go/tls
$
$ openssl genrsa 2048 > server.key;
$
$ openssl req -new -key server.key > server.csr
...
Common Name (eg, fully qualified host name) []:*.ucwork.local
...
$ openssl x509 -days 3650 -req -signkey server.key < server.csr > server.crt
```

### NySQL用のパスワード生成
```zsh
$ echo -n 'password' > password
```

### 各種デプロイ
このシェルに一通りコマンド書いた。結構時間かかる

```zsh
$ zsh deploy.sh
```
