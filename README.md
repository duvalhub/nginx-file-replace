# nginx-file-replace
The aim of this project is to build a binary that we can put in an nginx container. The binary will provide the necessary functionalities to nginx to serve Single Page App properly.

There are 2 main functionalities required
 * Able to edit files
      This is to be able to inject configurations into our SPA without having to rebuild the app. The SPA must plan ahead and have a configuration file that it sources and get his various configuration from there. We are talking here configs like url to oauth2 authority, credential id to said authority, debug log, etc.
 * Able to create an nginx configuration file
      This is mainly to be able to proxy request towards SPA's backends. The challenge comes from sending request towards same domain name and thus not having CORS requests. The proxy must then be configurable at runtime so we don't have to make a build per environment.


Those 2 features are especially interesting connected to a Config Server. This app has [this dependency](https://github.com/duvalhub/cloudconfigclient) to connect to a Spring Config Server.
