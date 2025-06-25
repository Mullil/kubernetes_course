FROM python:3.13-slim

WORKDIR /app

COPY log_output.py .

COPY files/logs.txt .

RUN pip install fastapi uvicorn requests

EXPOSE 3000

CMD ["python", "log_output.py"]


