# ================================
#  Makefile para lanzar VM1–VM4
# ================================

DEPLOY_DIR=deploy

docker-vm1:
	@echo ">>> Ejecutando Docker Compose para MV1..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv1.yml up --build

docker-vm2:
	@echo ">>> Ejecutando Docker Compose para MV2..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv2.yml up --build

docker-vm3:
	@echo ">>> Ejecutando Docker Compose para MV3..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv3.yml up --build

docker-vm4:
	@echo ">>> Ejecutando Docker Compose para MV4..."
	docker-compose -f $(DEPLOY_DIR)/compose-mv4.yml up --build

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
