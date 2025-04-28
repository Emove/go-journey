package file

import "os"

func OpenWithNotExist(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func Append(name string, data []byte) error {
	file, err := OpenWithNotExist(name)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
