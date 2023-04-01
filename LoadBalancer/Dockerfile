
FROM python:3.9.5-alpine

RUN mkdir /LoadBalancer
ADD . /LoadBalancer
WORKDIR /LoadBalancer

#RUN echo "http://dl-4.alpinelinux.org/alpine/v3.14/main" >> /etc/apk/repositories && \
#    echo "http://dl-4.alpinelinux.org/alpine/v3.14/community" >> /etc/apk/repositories
#
#RUN apk update

# RUN apk add --update --no-cache --virtual .tmp gcc libc-dev linux-headers 
# RUN apk add --no-cache jpeg-dev zlib-dev postgresql-dev chromium chromium-chromedriver
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt

COPY . .

ENV FLASK_APP=loadbalancer.py

CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]
