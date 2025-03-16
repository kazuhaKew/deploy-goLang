# Build and start all services
## Option 1: Using Docker Compose (Recommended)

```bash
docker-compose up --build -d
```

# Stop all services
```bash
docker-compose down
```

__________________________________________________________
## Option 2: Building Individual Images
### Build goFiber service
```bash
cd /backend-goFiber
docker build -t backend-gofiber:latest .
```

### Build fileServer service
```bash
cd /backend-fileServer
docker build -t backend-fileserver:latest .
```

### Build gorilla service
```bash
cd /backend-gorilla
docker build -t backend-gorilla:latest .
```
__________________________________________________________

Useful Docker Commands
# View running containers
```bash
docker ps
```

# View all containers (including stopped ones)
docker ps -a

# View built images
docker images

# Stop all services started with docker-compose
docker-compose down

# View logs
docker-compose logs -f                  # All services
docker-compose logs -f gofiber          # Specific service
docker-compose logs -f fileserver       # Specific service
docker-compose logs -f gorilla          # Specific service

