package filesystem

import (
	"GoProject_1/repositories/models"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type SupplierFileRepository struct {
}

func (sfr *SupplierFileRepository) GetAll() (suppliers []*models.Supplier, err error) {
	supplierRepo, err := ioutil.ReadDir("./datastore/files/suppliers/")
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range supplierRepo {
		file, err := os.Open("./datastore/files/suppliers/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		defer file.Close()
		supplier := &models.Supplier{}
		var data []byte
		for {
			chunk := make([]byte, 64)
			n, err := file.Read(chunk)
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			data = append(data, chunk[:n]...)
		}
		err = json.Unmarshal(data, &supplier)
		if err != nil {
			return nil, err
		}
		suppliers = append(suppliers, supplier)
	}
	return suppliers, nil
}
