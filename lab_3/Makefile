# ================================
#  Makefile para lanzar VM1–VM4
# ================================

DEPLOY_DIR=deploy

docker-vm1:
	@echo ">>> Ejecutando Docker Compose para MV1..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv1.yml build --no-cache

docker-vm2:
	@echo ">>> Ejecutando Docker Compose para MV2..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv2.yml build --no-cache

docker-vm3:
	@echo ">>> Ejecutando Docker Compose para MV3..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv3.yml build --no-cache

docker-vm4:
	@echo ">>> Ejecutando Docker Compose para MV4..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv4.yml build --no-cache

start-vm1:
	@echo ">>> Iniciando MV1 (Sin reconstruir)..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv1.yml up

start-vm2:
	@echo ">>> Iniciando MV2 (Sin reconstruir)..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv2.yml up

start-vm3:
	@echo ">>> Iniciando MV3 (Sin reconstruir)..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv3.yml up

start-vm4:
	@echo ">>> Iniciando MV4 (Sin reconstruir)..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv4.yml up

# Para detener contenedores individuales por VM
down-vm1:
	docker-compose -f $(DEPLOY_DIR)/compose-mv1.yml down

down-vm2:
	docker-compose -f $(DEPLOY_DIR)/compose-mv2.yml down

down-vm3:
	docker-compose -f $(DEPLOY_DIR)/compose-mv3.yml down

down-vm4:
	docker-compose -f $(DEPLOY_DIR)/compose-mv4.yml down

# Para limpiar contenedores huérfanos
clean:
	docker system prune -f
