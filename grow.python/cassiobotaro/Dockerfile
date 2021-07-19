FROM python:3-alpine
WORKDIR /app
COPY requirements.txt /app
RUN pip install -r requirements.txt --no-cache
COPY . /app
EXPOSE 8080
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8080"]
