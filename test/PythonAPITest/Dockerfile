FROM manticoresearch/manticore:3.3.0
RUN apt-get update && apt-get install -y python python-pip curl
WORKDIR /var/PythonAPITest
COPY ./source/ ./source/
RUN pip install -r source/requirements.txt && \
cd source && python setup.py install
COPY ./test/ ./
RUN mkdir /var/data/ && chmod +x test.sh
EXPOSE 9308 9306
CMD ["sh", "-c", "./test.sh"]  