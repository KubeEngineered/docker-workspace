import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpExchange;
import java.net.InetSocketAddress;
import java.io.OutputStream;
import java.io.IOException;

public class Main {
    public static void main(String[] args) throws Exception {
        int port = 8080;
        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);

        // Root Context ("/")
        server.createContext("/", exchange -> {
            String response = "Hello from Java app (v2.0) running inside Docker container!\n";
            sendResponse(exchange, 200, response);
        });

        // Health check endpoint ("/health")
        server.createContext("/health", exchange -> {
            String response = "{\"status\": \"UP\"}\n";
            sendResponse(exchange, 200, response);
        });

        System.out.println("Server listening on port " + port + "...");
        server.start();
    }

    private static void sendResponse(HttpExchange exchange, int statusCode, String response) throws IOException {
        byte[] bytes = response.getBytes();
        exchange.sendResponseHeaders(statusCode, bytes.length);
        try (OutputStream os = exchange.getResponseBody()) {
            os.write(bytes);
        }
    }
}