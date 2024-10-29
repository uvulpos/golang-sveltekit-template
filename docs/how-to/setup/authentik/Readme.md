# Setup Authentik

To setup a (local) authentication provider, start your docker development stack `make dev` and connect to [your local authentik server (http://localhost:9000/)](http://localhost:9000/)

## Create Provider

First you need to create a new provider, that we can reference in an application:

<a href="./images/create-provider-1.png">
    <img src="./images/create-provider-1.png" width="500px" alt="New Provider New Provider Type">
</a>

<a href="./images/create-provider-2.png">
    <img src="./images/create-provider-2.png" width="500px" alt="New Provider Base Configuration">
</a>

<a href="./images/create-provider-3.png">
    <img src="./images/create-provider-3.png" width="500px" alt="New Provider Advanced Protocol Settings">
</a>

## Create Application

After your provider is created, create a new application and reference your provider:

<a href="./images/create-application.png">
    <img src="./images/create-application.png" width="500px" alt="New Application">
</a>

## Edit .env File

After you have the credentials, you can insert them into the env file. If you only working locally with `make dev`, you need to replace the Client Key + Secret. The URLs are matching with the endpoint. If you are working with a remote solution, you can inspect them in your provider view. Auto-OpenID-Configuration is not supportet yet.

<a href="./images/env.png">
    <img src="./images/env.png" width="500px" alt="Env File">
</a>

After you edited the configuration, restart the docker stack by interrupting the signal and to restart the application by executing `make dev`
