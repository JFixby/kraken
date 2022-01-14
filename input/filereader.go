package input

type FileReader struct {
	inputFile string
}

func NewFileReader(inputFile string) *FileReader {
	return &FileReader{inputFile}
}
