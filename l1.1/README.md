# Embedded struct / embedded field

## Структуры

- `Human` — базовая структура с полями и методами.
- `Action` — структура, **встраивающая** `Human` (embedded field), добавляется переопределение метода `Who()` и собственная логика.

## Как запустить

```bash
cd l1.1/cmd/app
go run .
```

Ожидаемый вывод:

```
Hi, I'm Sergey (20 y.o.)
Hi, I'm Sergey (21 y.o.)
Action/DevOps
Human
Hi, I'm Sergey (22 y.o.)
Action/DevOps
```

## Tесты

```bash
cd l1.1
go test ./...
```

## Описание

- `internal/domain/human.go` — базовая структура и методы.
- `internal/domain/action.go` — встраивание (`Human` без имени поля), переопределение `Who()`,
  явный вызов метода «родителя» через `a.Human.Method()`.
- `internal/domain/domain_test.go` — тесты
