sonarqube-scan:
	@echo "Running SonarScanner with project key..."
	@docker run --rm \
		-v "$$(pwd):/usr/src" \
		--network="sonarqube-network" \
		-e SONAR_HOST_URL="http://sonarqube:9000" \
		-e SONAR_SCANNER_OPTS="-Dsonar.projectKey=project-demo" \
		-e SONAR_TOKEN="sqp_bfaba9e565adfb9f8ee41e546e8c9dee798c9bbc" \
		sonarsource/sonar-scanner-cli

