# miners-and-postmans

Небольшой учебный проект на Go про конкурентное выполнение задач.

Программа запускает две группы workers:

- **miners** — добывают уголь;
- **postmans** — доставляют письма.

Обе группы работают параллельно в отдельных goroutines, отправляют результаты через каналы и завершаются по сигналу из `context`.

## Что используется

- goroutines
- channels
- `sync.WaitGroup`
- `context`
- `sync/atomic`
- `sync.Mutex`

## Структура проекта

```text
.
├── main.go
├── go.mod
├── miner/
│   └── miner.go
└── postman/
    └── postman.go
