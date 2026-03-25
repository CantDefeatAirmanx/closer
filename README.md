# closer

```go
package main

import (
	"context"
	"log"

	"github.com/CantDefeatAirmanx/closer"
)

func main() {
	ctx := context.Background()
	cl, done := closer.NewCloser(ctx)

	defer func() {
		go func() {
			if err := cl.CloseAll(ctx); err != nil {
				log.Printf("shutdown err: %v", err)
			}
		}()
		<-done
	}()

	cl.AddNamed("Dependency", func(ctx context.Context) error {
        // return Dependency.Close()
		return nil
	})
}
```
