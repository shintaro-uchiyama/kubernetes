#!/bin/zsh
# Nginx
kubectl create secret tls tls-secret --key docker/go/tls/server.key --cert docker/go/tls/server.crt &&
kubectl create configmap nginx-conf --from-file=docker/nginx/conf &&
kubectl create -f nginx-deployment.yaml &&

# MySQL
kubectl create secret generic mysql-pass --from-file=./password &&
kubectl create -f mysql-deployment.yaml &&

# golang
kubectl create -f go-deployment.yaml &&

# ingress
## helm
helm init --upgrade &&
runningPodCnt=0
while :
do
runningPodCnt=$(kubectl get pods --namespace kube-system | grep 'tiller-deploy' | grep 'Running' | grep '1/1' | wc -l | xargs)
if [ $runningPodCnt = 1 ]; then
  echo "tiller deployed"
  break
fi
done

## ingress controller
helm install \
  --namespace kube-system \
  --set controller.hostNetwork=true \
  --set controller.kind=DaemonSet \
  --set controller.extraArgs.enable-ssl-passthrough="" \
  stable/nginx-ingress &&

## ingress rule
kubectl create -f ingress/ingress.yaml

## check ingress controller deployed
runningPodCnt=0
while :
do
runningPodCnt=$(kubectl get pods --namespace kube-system | grep 'nginx-ingress-controller' | grep 'Running' | grep '1/1' | wc -l | xargs)
if [ $runningPodCnt = 1 ]; then
  echo "nginx controller deployed"
  break
fi
done