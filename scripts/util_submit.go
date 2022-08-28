package scripts

import (
	"context"
	"fmt"
)

func Submit(tx []byte) {
	sig, err := Client.SendRawTransaction(context.TODO(), tx)
	if err != nil {
		panic(err)
	}

	fmt.Println(sig.String())
}
