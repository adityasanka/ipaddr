FROM mcr.microsoft.com/devcontainers/go:0-1.20

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

# Use vim for cmd line
RUN apt-get update && apt-get install -y --no-install-recommends \
vim

ENV VISUAL=vim
ENV EDITOR="$VISUAL"
ENV GIT_EDITOR=vim