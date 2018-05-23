package zipbuffer

import (
    "archive/zip"
    "bytes"
)

func NewZipBuffer(dataMap map[string]string) (*bytes.Buffer, error) {
    buf := new(bytes.Buffer)
    writer := zip.NewWriter(buf)
    for k,v := range dataMap {
        file, err := writer.Create(k)
        if err != nil {
            return nil, err
        }
        _, err = file.Write([]byte(v))
        if err != nil {
            return nil, err
        }
    }
    err := writer.Close()
    if err != nil {
        return nil, err
    }

    return buf, nil
}
