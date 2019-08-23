# Utilizando-K8S

Exercícios de Kubernetes e GCP

## 1) Servidor Web - Nginx

* Configurar um servidor NGINX usando a imagem base do Nginx Alpine;
* Alterar a pagina default para constar a mensagem: Code.education Rocks;
* Efetuar deploy no Kubernetes Engine;
* Disponibilar 3 réplicas.

*Quando alguém acessar o IP externo do LoadBalancer do serviço criado, ou em caso de utilização do Minikube usando "minikube service nome-do-servico", deve ser exibido no browser: Code.education Rocks.*

[Acesse aqui a resolução e explicações](https://github.com/rogeriopontomoura/Utilizando-K8S#1-servidor-web---nginx-1)

## 2) Configuração do MySQL

* Configuar um servidor de banco de dados MySQL;
* Utilizar secret em conjunto com as variáveis de ambiente para passagem de parâmetros;
* Utilizar disco persistente para gravar as informações do banco;
* Efetuar deploy no Kubernetes Engine;

[Acesse aqui a resolução e explicações](https://github.com/rogeriopontomoura/Utilizando-K8S#2-configura%C3%A7%C3%A3o-mysql)

## 3) Desafio Go!

* Crie um aplicativo Go que disponibiliza um servidor web na porta 8000, quando acessado deve exibir em HTML (em negrito) Code.education Rocks!
* A exibição dessa string deve ser baseada no retorno de uma função chamada "greeting". Essa função receberá a string como parâmetro e a retornará entre as tags <b></b>.
* Criar o teste dessa função.
* Ative o processo de CI no Google Cloud Build para garantir que a cada PR criada faça com que os testes sejam executados.
* Gere a imagem da aplicatição de forma otimizada e publique-a no Docker Hub
* Utilizando o Kubernetes, disponibilize o serviço do tipo Load Balancer que quando acessado pelo browser acesse a aplicação criada em Go.


Este exercício foi colocado em outro repositório.

[Clique aqui para acessar](https://github.com/rogeriopontomoura/ServidorWebGO)


# Resolução e instruções

** Configuração do ambiente GCP

### Criar cluster no GCP:

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