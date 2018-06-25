FROM ubuntu:latest
EXPOSE 80
COPY kubectl /usr/local/bin/
COPY oc /usr/local/bin/
COPY volume2volume /usr/local/bin/
RUN /bin/bash -c 'apt update ; apt -y install curl ; echo "curl -s 'https://raw.githubusercontent.com/mohammed-ali-1/v2v/master/v2v'" >> /etc/bash.bashrc'
CMD sleep 99999999999
