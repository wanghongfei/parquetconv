package parquetconv

import (
	"fmt"
	"github.com/xitongsys/parquet-go/ParquetFile"
	"github.com/xitongsys/parquet-go/ParquetReader"
	"runtime"
)

// 将parquet文件中的指定列按行打印到console
func ExtractColumn(parquetFile string, colnameName string) error {
	parFile, err := ParquetFile.NewLocalFileReader(parquetFile)
	if err != nil {
		return err
	}
	defer parFile.Close()

	reader, err := ParquetReader.NewParquetColumnReader(parFile, int64(runtime.NumCPU()))
	if err != nil {
		return err
	}
	defer reader.ReadStop()

	num := int(reader.GetNumRows())
	for ix := 0; ix < num; ix++ {
		names, _, _ := reader.ReadColumnByPath(colnameName, ix)
		for _, id := range names {
			if id == "" {
				continue
			}

			fmt.Println(id)
		}
	}

	return nil
}
