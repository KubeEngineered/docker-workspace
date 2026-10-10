FROM python:3.11-slim

WORKDIR /app

# Copy dependency manifests first to leverage layer caching
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the application source code
COPY app.py .

EXPOSE 5000

# Default environment variables (can be overridden at runtime via docker run -e)
ENV APP_VERSION=v2.0.0 \
    ENVIRONMENT=lab-testing

CMD ["python", "app.py"]