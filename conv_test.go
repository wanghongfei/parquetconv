package parquetconv

import "testing"

func TestExtractColumn(t *testing.T) {
	err := ExtractColumn("/Users/whf/Documents/parquet/pure/part-00000-ca929d3c-cfad-45a4-99df-6d7d589c835c-c000.snappy.parquet", "id")
	if nil != err {
		t.Error(err)
	}
}

