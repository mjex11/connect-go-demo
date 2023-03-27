package main

import (
    "context"
    "log"
    "net/http"

    promotionv1 "promotion/gen/promotion/v1"
    "promotion/gen/promotion/v1/promotionv1connect"

    "github.com/bufbuild/connect-go"
)

func main() {
    client := promotionv1connect.NewPromotionServiceClient(
        http.DefaultClient,
        "http://0.0.0.0:8080",
    )
    log.Println("あああああ")
    res, err := client.GetPointBlance(
        context.Background(),
        connect.NewRequest(&promotionv1.GetPointBlanceRequest{MemberId: 55}),
    )
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("いいいい")
    log.Println(res.Msg.Point)
}
