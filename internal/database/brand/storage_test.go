package brand

import (
	"context"
	"fmt"
	"market/internal/config"
	"market/internal/domain/brand"
	"market/pkg/database/postgresql"
	"testing"
)

func GetConn() *BrandStorage {
	conf := config.GetConfig("config_test.json")
	pg, _ := postgresql.NewConnection(5, conf)
	conn := NewBrand(pg)
	return conn
}

func TestCreate(t *testing.T) {
	conn := GetConn()
	var listtest = []struct {
		mass brand.CreateBrandDTO
	}{
		{mass: brand.CreateBrandDTO{
			"Hello",
			"yuio",
			"lalala",
		},
		},
		{mass: brand.CreateBrandDTO{
			"Its my world",
			"yandex/helloworld",
			"Pribett",
		}},
	}

	list, err := conn.AllRowsDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)

	for _, el := range listtest {
		_, err := conn.CreateRowDB(context.TODO(), el.mass)
		if err != nil {
			t.Error(err)
		}
	}

	list, err = conn.AllRowsDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func TestDelete(t *testing.T) {
	conn := GetConn()
	listtest := []struct {
		br brand.DeleteBrandDTO
	}{
		{
			br: brand.DeleteBrandDTO{15},
		},
		{
			br: brand.DeleteBrandDTO{16},
		},
	}

	for _, test := range listtest {
		if err := conn.DeleteRowDB(context.TODO(), test.br); err != nil {
			fmt.Println(err)
		}
	}

	list, err := conn.AllRowsDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func TestGet(t *testing.T) {
	conn := GetConn()

	listtest := []struct {
		mass brand.GetBrandDTO
	}{
		{
			brand.GetBrandDTO{"5"},
		},
		{
			brand.GetBrandDTO{"3"},
		},
	}

	for _, test := range listtest {
		br, err := conn.GetRowDB(context.TODO(), test.mass)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(br)
	}
}

func TestAll(t *testing.T) {
	conn := GetConn()
	list, err := conn.AllRowsDB(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func TestUpdate(t *testing.T) {
	conn := GetConn()
	listtest := []struct {
		mas brand.UpdateBrandDTO
	}{
		{
			brand.UpdateBrandDTO{
				3,
				"New Title",
				"New Image Path",
				"New Description",
			},
		},
		{
			brand.UpdateBrandDTO{
				8,
				"its",
				"me",
				"friend",
			},
		},
	}

	for _, test := range listtest {
		br, err := conn.UpdateRowDB(context.TODO(), test.mas)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(br)
	}
}
