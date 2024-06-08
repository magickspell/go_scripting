package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"main/pb/my"
	"os"
)

func main() {
	file, err := os.OpenFile("../output.bin", os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}

	result, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	params := &my.SellerParams{}
	err = proto.Unmarshal(result, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("print marshal protobuf")
	fmt.Println(
		proto.Marshal(&my.SellerParams{
			Result: []*my.SellerParams_Item{
				{SellerId: 4},
			},
		}),
	)
	fmt.Println()

	fmt.Printf("result:\n%+v\n\n\n", params)
	fmt.Printf("first item:\n%+v\n\n\n", params.GetResult()[0])
	fmt.Println()
	fmt.Println()
	fmt.Printf("first item:\n%+v\n\n\n", params.GetResult()[0].GetParams()["not_exist"]) // найдет nil (нулевой значение)
	fmt.Println()
	fmt.Println()
	fmt.Printf("first item:\n%+v\n\n\n", params.GetResult()[0].GetParams()["not_exist"].GetDouble()) // найдет 0
}
