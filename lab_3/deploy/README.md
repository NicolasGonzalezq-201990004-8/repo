# Despliegue en 4 VMs

Los `docker compose` de `deploy/compose-mv*.yml` asumen las IPs privadas del laboratorio
(10.35.168.x). Si tus 4 máquinas tienen IPs distintas, puedes sobrescribirlas con
variables de entorno sin editar los YAML.

## Variables clave (mismos defaults que el repositorio)

| Variable | Uso | Valor por defecto |
| --- | --- | --- |
| `BROKER_ADDR` | Endpoint público del broker para clientes y coordinador | `10.35.168.15:50050` |
| `DATANODE_ADDR` | Dirección del datanode que vive en la VM actual | `10.35.168.15:50051` en MV1, `10.35.168.17:50051` en MV3, `10.35.168.112:50051` en MV4 |
| `DATANODE_ADDRESSES` | Lista completa de datanodes para broker/datanodes | `10.35.168.17:50051,10.35.168.112:50051,10.35.168.15:50051` |
| `CONSENSUS1_ADDR` | ATC-1 (VM2) | `10.35.168.16:50060` |
| `CONSENSUS2_ADDR` | ATC-2 (VM3) | `10.35.168.17:50060` |
| `CONSENSUS3_ADDR` | ATC-3 (VM4) | `10.35.168.112:50060` |
| `CONSENSUS_ADDRESSES` | Lista de consenso que usa el broker (MV1) | `10.35.168.16:50060,10.35.168.17:50060,10.35.168.112:50060` |
| `COORDINATOR_ADDR` | Endpoint del coordinador para clientes MR/RYW | `10.35.168.112:50070` |

## Cómo levantar cada VM con tus IPs
1. En cada host, exporta las variables necesarias con las IPs reales de tu red.
   Ejemplo (ajusta las IPs de los demás nodos según tu topología):
   ```bash
   export BROKER_ADDR=192.168.10.15:50050
   export DATANODE_ADDR=192.168.10.15:50051
   export DATANODE_ADDRESSES=192.168.10.15:50051,192.168.10.16:50051,192.168.10.17:50051
   export CONSENSUS1_ADDR=192.168.10.21:50060
   export CONSENSUS2_ADDR=192.168.10.22:50060
   export CONSENSUS3_ADDR=192.168.10.23:50060
   export CONSENSUS_ADDRESSES="$CONSENSUS1_ADDR,$CONSENSUS2_ADDR,$CONSENSUS3_ADDR"
   export COORDINATOR_ADDR=192.168.10.23:50070
   ```

2. Ejecuta el `docker compose` de la VM actual:
   ```bash
   # En VM1 (broker + datanode3 + clientes)
   docker compose -f deploy/compose-mv1.yml up -d

   # En VM2 (consensus1 + clientes)
   docker compose -f deploy/compose-mv2.yml up -d

   # En VM3 (datanode1 + consensus2 + cliente)
   docker compose -f deploy/compose-mv3.yml up -d

   # En VM4 (coordinator + datanode2 + consensus3)
   docker compose -f deploy/compose-mv4.yml up -d
   ```

3. Verifica conectividad desde cada contenedor a sus peers (ping o `grpcurl`) usando
   las IPs configuradas. Si alguno no responde, corrige la variable en esa VM.

## Notas
- Si pruebas solo en una VM, puedes apuntar a los hostnames internos del compose
  (por ejemplo, `BROKER_ADDR=broker:50050`, `DATANODE_ADDRESSES=datanode3:50051`).
- Las variables con lista (`DATANODE_ADDRESSES`, `CONSENSUS_ADDRESSES`) deben contener
  todos los nodos del clúster para que broker y datanodes conozcan a sus pares.
