package main

import (
	"acronis/client"
	"context"
	"fmt"
)

func main() {
	cl := client.NewDataStoreClient("http://localhost:8081")
	ctx := context.Background()

	fmt.Println("=======Setting List Data 1============")
	resp1, err := cl.SetData(ctx, map[string]interface{}{"data": []interface{}{"hi"}})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Response:", resp1)

	fmt.Println("=======Setting List Data 2============")

	resp2, err := cl.SetData(ctx, map[string]interface{}{"data": []interface{}{"hello", 1, 1.2, "ook"}})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Response:", resp2)

	fmt.Println("=======Setting string Data ============")

	resp3, err := cl.SetData(ctx, map[string]interface{}{"data": "I am mayank"})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Response:", resp3)

	fmt.Println("=======Getting All Data============")
	getData, err := cl.GetAllData(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Get All Data Response:", getData)

	fmt.Println("=======Getting Data By Key============")
	getDatabyKey, err := cl.GetDataByKey(ctx, resp1.Key)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Get Data by Key Response:", getDatabyKey)

	fmt.Println("=======Update Data By Key============")
	UptDatabyKey, err := cl.UpdateData(ctx, resp2.Key, map[string]interface{}{"data": []interface{}{"updated the value"}})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Updated Data by Key Response:", UptDatabyKey)

	fmt.Println("=======Delete Data By Key============")
	err = cl.DeleteData(ctx, resp3.Key)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Data Deleted by Key Response")

	fmt.Println("=======Push Data into existing list============")
	UptListDatabyKey, err := cl.PushDataIntoList(ctx, resp1.Key, map[string]interface{}{"uptvalue": "updated the reposen1"})
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Updated List Data by value Response:", UptListDatabyKey)

	fmt.Println("#################################################")
	getDatabyKey, err = cl.GetDataByKey(ctx, resp1.Key)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Get Data by Key Response:", getDatabyKey)

	fmt.Println("#################################################")

	fmt.Println("=======Pop Data into existing list============")
	DelListDatabyKey, err := cl.PopDataFromList(ctx, resp1.Key, map[string]interface{}{"uptvalue": "updated the reposen1"})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Delete List Data by value Response:", DelListDatabyKey)

	fmt.Println("*************************************************")

	getData, err = cl.GetAllData(ctx)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Get All Data Response:", getData)

	fmt.Println("*************************************************")

}
