Эти сервисы являются MVP реализацией тестового задания. Код был написан на быструю руку, не было реализовано валидация, логирование и прочие мелкие детали :)

Для генерации кода от .proto файлов: `protoc --go_out=./названиеДиректории --go_opt=paths=source_relative --go-grpc_out=./названиеДиректории --go-grpc_opt=paths=source_relative proto/data.proto`.

Для запуска проекта: `docker-compose up`.

Postman коллекции представлены в папке `postman`.