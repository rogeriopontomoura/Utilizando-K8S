# Utilizando-K8S
Exercícios de Kubernetes e GCP

## 1) Servidor Web - Nginx

Utilize a imagem base do Nginx Alpine

Disponibilize 3 réplicas

Quando alguém acessar o IP externo do LoadBalancer do serviço criado, ou em caso de utilização do Minikube usando "minikube service nome-do-servico", deve ser exibido no browser: Code.education Rocks.

## 2) Configuração do MySQL

Faça o processo de configuração de um servidor de banco de dados MySQL
Utilize secret em conjunto com as variáveis de ambiente
Utilize disco persistente para gravar as informações dos dados

## 3) Desafio Go!

Crie um aplicativo Go que disponibilize um servidor web na porta 8000 que quando acessado seja exibido em HTML (em negrito) Code.education Rocks!
A exibição dessa string deve ser baseada no retorno de uma função chamada "greeting". Essa função receberá a string como parâmetro e a retornará entre as tags <b></b>.
Como ótimo desenvolvedor(a), você deverá criar o teste dessa função.
Ative o processo de CI no Google Cloud Build para garantir que a cada PR criada faça com que os testes sejam executados.
Gere a imagem desse aplicativo de forma otimizada e publique-a no Docker Hub
Utilizando o Kubernetes, disponibilize o serviço do tipo Load Balancer que quando acessado pelo browser acesse a aplicação criada em Go.


### Resolução

** Configuração do ambiente GCP

Criar cluster no GCP:

Definir projeto padrão:

> gcloud config set project [ID_PROJETO]

Definir zona padrão:

> gcloud config set compute/zone [ZONA]

Criar cluster do Kubernetes:

> gcloud container clusters create "[NOME-CLUSTER]"

Receber credenciais de autenticação para o cluster:

> gcloud container clusters get-credentials [NOME-CLUSTER]

Verificar quantidade de nodes:

> kubectl get nodes

## 1) Servidor Web - Nginx

### Passo 1 - Clonar Repositório

Clonar este repositório e entrar na pasta nginx

### Passo 2 - Efetuar deploy do configmap:

> kube ctl apply -f configmap.yaml

### Passo 3 - Efetuar deploy da aplicação:

> kube ctl apply -f deployment.yaml

Verificar status do deploy:

> kube ctl get pods

### Passo 4 - Efetuar deploy do service

> kubectl apply -f service.yaml

Verificar status do serviço:

> kubectl get services

Acessar pelo navegador o IP informado no LoadBalancer


## 2) Configuração MySQL

### Passo 1 - Clonar Repositório

Clonar este repositório e entrar na pasta MySQL

### Passo 2 - Efetuar deploy do persistent-volume:

> kube ctl apply -f persistent-volume.yaml

### Passo 3 - Efetuar deploy da secret:

> kube ctl apply -f secret.yaml

### Passo 3 - Efetuar deploy da aplicação:

> kube ctl apply -f deployment.yaml

Verificar status do deploy:

> kube ctl get pods

### Passo 4 - Efetuar deploy do service

> kubectl apply -f service.yaml

Verificar status do serviço:

> kubectl get services

### Passo 5 - Testar serviço do MySQL

Execute um cliente MySQL para se conectar ao servidor:

> kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -ppassword

OBS: o password deve ser o informado no arquivo secret

Esse comando cria um novo Pod no cluster que executa um cliente MySQL e o conecta ao servidor por meio do Serviço. Se ele se conecta, você sabe que seu banco de dados MySQL com estado está funcionando