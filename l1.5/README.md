# L1.5 Timed Channel

Sequentially generates integers, sends them through a channel, and stops after the specified duration.

## Run

```bash
go run . -duration 3s -interval 250ms
```

- `-duration` controls how long the program runs before shutting down (default 5s).
- `-interval` sets the delay between emitted values (default 500ms).

## Output

Values appear in order until the timer fires, after which the program prints `finished` and exits.
