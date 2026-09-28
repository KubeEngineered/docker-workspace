# ---- Build stage ----
FROM eclipse-temurin:21-jdk-alpine AS builder
WORKDIR /app
COPY Main.java .
RUN javac Main.java

# ---- Runtime stage ----
FROM eclipse-temurin:21-jre-alpine
WORKDIR /app
RUN addgroup -S javagroup && adduser -S javauser -G javagroup
COPY --from=builder /app/Main.class .
USER javauser
EXPOSE 8080
CMD ["java", "Main"]
