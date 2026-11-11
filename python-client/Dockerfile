
FROM python:3.14-slim

WORKDIR /app

COPY python-client/requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY python-client/ ./

ENV GRPC_SERVER=go-server:50051

CMD ["python", "client.py"]
