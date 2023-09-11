run:
	go run ./cmd

check:
	curl http://localhost:8088


build:
	go build -o=app ./cmd

# 初回起動時にweb-serverが立ち上がらなければ2回実行する
dkc-run:
	docker-compose up -d