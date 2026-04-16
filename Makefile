.PHONY: bootstrap lint test build run-local audit-up audit-run audit-migrate ui-install ui-dev ui-build observability-up metrics-check package helm-lint helm-template helm-template-prod helm-package kind-up kind-down kind-load-images e2e

bootstrap:
	go mod download

lint:
	gofmt -w ./internal ./services
	go test ./...

test:
	go test ./...

build:
	go build ./services/policy-engine
	go build ./services/attestation-verifier
	go build ./services/deploy-gate
	go build ./services/runtime-agent
	go build ./services/audit-writer

run-local:
	docker compose -f docker-compose.dev.yml up --build

audit-up:
	docker compose -f docker-compose.dev.yml up -d postgres audit-writer

audit-run:
	go run ./services/audit-writer

audit-migrate:
	go run ./services/audit-writer -migrate-only

ui-install:
	cd ui && npm install

ui-dev:
	cd ui && npm run dev:host

ui-build:
	cd ui && npm run build

observability-up:
	docker compose -f docker-compose.dev.yml --profile observability up -d prometheus

metrics-check:
	curl -sS http://127.0.0.1:8094/metrics | rg '^changelock_'

package:
	zip -r changelock-skeleton.zip .

helm-lint:
	helm lint charts/changelock

helm-template:
	helm template changelock charts/changelock

helm-template-prod:
	helm template changelock charts/changelock -f charts/changelock/values-prod-example.yaml

helm-package:
	helm package charts/changelock

kind-load-images:
	./scripts/load_images_kind.sh

kind-up:
	./scripts/bootstrap_local_kind.sh

kind-down:
	kind delete cluster --name $${CLUSTER_NAME:-changelock}

e2e:
	./scripts/run_e2e.sh
