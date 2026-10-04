FROM node:18-alpine

# For Debian/Ubuntu-based images
RUN apk update && apk upgrade --no-cache

WORKDIR /app

# Copy dependency manifests first
COPY package*.json ./

# Cache layer: This step re-runs ONLY if package.json or package-lock.json changes
RUN npm ci --only=production

# Copy the rest of the application code
# Cache layer: Modifying files in src/ will break cache starting FROM THIS POINT onwards
COPY . .

EXPOSE 3000

CMD ["node", "src/index.js"]