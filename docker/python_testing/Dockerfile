FROM manticoresearch/manticore
RUN apt-get update && apt-get install -y python python3-pip curl  python3-setuptools && rm -rf /var/lib/apt/lists/*
COPY ./prep.sh /
CMD []  
