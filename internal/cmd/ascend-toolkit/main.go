package main

import (
	"context"
	"fmt"
	"os"

	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"

	"github.com/innoai-tech/ascend-toolkit/cmd/ascend-toolkit/collector"
)

func main() {
	ctx := context.Background()

	c := &collector.Collector{
		Debug: os.Getenv("ASCEND_TOOLKIT_DEBUG") == "1",
	}

	if err := c.Init(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	defer func() {
		if err := c.Shutdown(ctx); err != nil {
			fmt.Println(err)
		}
	}()

	if err := c.Collect(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	_ = jsonv2.MarshalWrite(os.Stdout, c.Ascend, jsontext.WithIndent(" "))
	_, _ = fmt.Fprintln(os.Stdout)
}
