# Stage 1: Base image
FROM python:3.11-alpine AS base

ENV PYTHONDONTWRITEBYTECODE 1

COPY req.txt .

RUN apk --no-cache add libpq-dev gcc python3-dev && rm -f /var/lib/apt/lists/*

RUN pip install --no-cache-dir -r req.txt

# Stage 2: Final image
FROM python:3.11-alpine

COPY --from=base /usr/local/lib/python3.11/site-packages /usr/local/lib/python3.11/site-packages
COPY --from=base /usr/local/bin/ /usr/local/bin/

ENV PYTHONUNBUFFERED 1
ENV PYTHONDONTWRITEBYTECODE 1

COPY . /tn_service
WORKDIR /tn_service

RUN chmod +x ./compile_proto.sh
RUN sh compile_proto.sh


CMD ["python", "main.py"]
