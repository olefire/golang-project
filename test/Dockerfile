FROM python

RUN pip install requests
RUN pip install bson

WORKDIR /app

COPY main.py /app

CMD ["python3", "./main.py"]