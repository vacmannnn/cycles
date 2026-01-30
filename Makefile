.PHONY: build test run clean

# Сборка исполняемого файла
build:
	go build -o bin/trees .

# Запуск тестов
test:
	go test -v ./...

# Запуск программы
run: build
	./bin/trees

# Очистка артефактов сборки
clean:
	rm -rf bin/