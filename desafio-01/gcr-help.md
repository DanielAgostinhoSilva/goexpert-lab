# documentação de ajuda com o google cloud-run

# Subindo um Dockerfile no Google Cloud Run

## Pré-requisitos
Antes de começar, certifique-se de que você tem o `gcloud` instalado e configurado em seu sistema.

## Passos

1. **Define o projeto do GCP**

   Use o seguinte comando para definir seu projeto GCP (substitua `PROJECT-ID` pelo ID do seu projeto):

    ```shell
    gcloud config set project PROJECT-ID
    ```

2. **Ative o Cloud Run**

   Use o seguinte comando para ativar o Cloud Run:

    ```shell
    gcloud services enable run.googleapis.com
    ```

3. **Autenticação no GCP**

   Autentique-se no GCP com o seguinte comando:

    ```shell
    gcloud auth login
    ```

4. **Construindo a imagem Docker**

   Navegue até o diretório que contém o seu Dockerfile e construa sua imagem Docker com o seguinte comando (substitua `PROJECT-ID` pelo ID do seu projeto GCP e `helloworld` pelo nome que você deseja dar à sua imagem):

    ```shell
    gcloud builds submit --tag gcr.io/PROJECT-ID/helloworld
    ```

5. **Implantando a imagem no Cloud Run**

   Para implantar sua imagem no Cloud Run, use o seguinte comando:

    ```shell
    gcloud run deploy --image gcr.io/PROJECT-ID/helloworld --platform managed
    ```

Agora a sua imagem Docker deve estar rodando no Google Cloud Run! Você pode acessar os detalhes do serviço (incluindo a URL na qual o seu serviço estará disponível) no console do Cloud Run.

6. **Parando o serviço**

   Para "parar" o serviço (isso na verdade exclui o serviço), execute o seguinte comando:

    ```shell
    gcloud run services delete SERVICE_NAME
    ```

   Substitua `SERVICE_NAME` pelo nome do seu serviço.