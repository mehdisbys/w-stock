package output

import (
	"fmt"
	"os"

	"sainsburys-stock/domain"
)

type WriterCSV struct {
	destination string
	results     chan domain.Result
	file        *os.File
}

func NewWriterCSV(destination string, results chan domain.Result) (*WriterCSV, error) {
	w := &WriterCSV{destination: destination, results: results}
	return w, w.createOutputFile()
}

func (w *WriterCSV) createOutputFile() error {
	var err error
	w.file, err = os.Create(w.destination)
	if err != nil {
		return err
	}
	return nil
}

func (w *WriterCSV) Output() error {

	defer w.file.Close()

	for f := range w.results {

		if f.NotFound {
			_, err := w.file.WriteString(fmt.Sprintf("%s Not Found\n", f.Product))
			if err != nil {
				return err
			}
			continue
		}

		_, err := w.file.WriteString(fmt.Sprintf("%s, %d, %d, %f, %s\n",
			f.Product, f.SKU, int(f.InventoryAvailable), f.RetailPrice, f.URL))
		if err != nil {
			return err
		}
	}
	return nil
}
