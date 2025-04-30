
# Sonar qube Community

Para rodar o Sonar qube, no linux deve-se realizar a seguinte configuração:
Abrir o arquivo ```/etc/sysctl.conf ``` e adicionar a seguinte informação no final do arquivo:
```sh
vm.max_map_count=262144
```

Logo em seguida executar o comadno:
```sh
/sbin/sysctl -p
```

Para subir o Sonar Qube localmente deve-se executar o comando:
```sh
docker-compose up --build -d
```

Em seguida acesse: http://0.0.0.0:9001/, usuário e senha padrão do Sonarqube:
* Usuário: **admin** 
* Senha: **admin**

Execute o seguinte comando abaixo para começar análise do código:
```sh
make sonarqube-scan
```

O comando acima irá executar o seguinte comando que deverá ser configurado dentro do Makefile:
```sh
docker run --rm \
	-v "$$(pwd):/usr/src" \
	--network="sonarqube-network \
	-e SONAR_HOST_URL="http://sonarqube:9000" \
	-e SONAR_SCANNER_OPTS="-Dsonar.projectKey=PROJECT-KEY" \
	-e SONAR_TOKEN="PROJECT-TOKEN" \
	sonarsource/sonar-scanner-cli
```