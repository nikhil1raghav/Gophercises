package secrets

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"
)
func File(encodingKey , filePath string) *Vault{
	return &Vault{
		encodingKey: encodingKey,
		filePath: filePath,
		keyValues: make(map[string]string),
	}
}
func Memory(encodingKey string) Vault{
	return Vault{
		encodingKey: encodingKey,
		keyValues: make(map[string]string),
	}
}
type Vault struct {
	encodingKey string
	keyValues map[string]string
	filePath string
	mutex sync.Mutex
}
func (v *Vault) readKeyValues(r io.Reader) error{
	dec:=json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}
func (v *Vault) writeKeyValues(w io.Writer) error{
	enc:=json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}
func (v *Vault) load() error{
	f, err:=os.Open(v.filePath)
	if err!=nil{
		return nil
	}
	defer f.Close()
	r, err:=DecryptReader(v.encodingKey, f)
	if err!=nil{
		return err
	}
	return v.readKeyValues(r)
}
func (v *Vault) save() error{
	f, err:=os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err!=nil{
		return err
	}
	defer f.Close()
	w, err:=EncryptWriter(v.encodingKey, f)
	if err!=nil{
		return err
	}
	return v.writeKeyValues(w)
}
/*
func (v *Vault) loadKeyValues() error{
	f, err:=os.Open(v.filePath)
	if err!=nil{
		return nil
	}
	dec:=json.NewDecoder(f)
	err = dec.Decode(&v.keyValues)
	if err!=nil{
		return err
	}
	return nil
}
func (v *Vault) SaveKeyValues() error{
	var sb strings.Builder
	enc:=json.NewEncoder(&sb)
	err:=enc.Encode(v.keyValues)
	if err!=nil{
		return err
	}
	f, err:=os.OpenFile(v.filePath, os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()
	if err!=nil{
		return err
	}
	fmt.Fprint(f, sb.String())


	return nil
}

 */
func (v *Vault) Get(key string) (string, error){
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err :=v.load()
	if err!=nil{
		return "", err
	}
	ret, ok:=v.keyValues[key]
	if !ok{
		return "", errors.New("No such string")
	}
	return ret,nil
}
func (v *Vault) Set(key, value string)  error{
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err:=v.load()
	if err!=nil{
		return err
	}
	v.keyValues[key]=value
	err=v.save()
	if err!=nil{
		return err
	}
	return nil
}
