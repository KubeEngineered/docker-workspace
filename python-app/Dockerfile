# --------- Stage 1: builder ---------
FROM python:3.11-slim

WORKDIR /app

# Copy requirements file into the builder stage first
COPY requirements.txt .

# Install into a local target dir so we can copy just the installed
# packages into the final stage (keeps build tools out of the final image)
RUN pip install --no-cache-dir -r requirements.txt

# Copy source code
COPY app.py .

EXPOSE 5000

ENV APP_VERSION=v1.0.0

CMD ["python", "app.py"]
