package main

import (
	"crypto"
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/packet"
)

func main() {
	file1 := "test1.key"
	file2 := "test2.key"
	file3 := "test3.key"
	file4 := "test4.key"

	privateKey, err := readprivateKey(file1)
	if err != nil {
		entity, err := openpgp.NewEntity("name", "comment", "email@domain.org", &packet.Config{
			DefaultHash:   crypto.SHA256,
			DefaultCipher: packet.CipherAES256,
		})
		if err != nil {
			log.Fatal(err)
		}
		privateKey = entity.PrivateKey
	}

	// log.Printf("Key[file2]: %#v", privateKey)
	// log.Printf("Key[file2].PrivateKey.CreationTime: %s", privateKey.PrivateKey.CreationTime.Format(time.RFC3339Nano))
	// log.Printf("Key[file2].PrimaryKey.CreationTime: %s", privateKey.PrimaryKey.CreationTime.Format(time.RFC3339Nano))
	err = writeprivateKey(file2, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = readprivateKey(file2)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("Key[file1]: %#v", privateKey)
	// log.Printf("Key[file1].PrivateKey.CreationTime: %s", privateKey.PrivateKey.CreationTime.Format(time.RFC3339Nano))
	// log.Printf("Key[file1].PrimaryKey.CreationTime: %s", privateKey.PrimaryKey.CreationTime.Format(time.RFC3339Nano))
	err = writeprivateKey(file3, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = readprivateKey(file3)
	if err != nil {
		log.Fatal(err)
	}

	err = writeprivateKey(file4, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err = readprivateKey(file4)
	if err != nil {
		log.Fatal(err)
	}

	err = writeprivateKey(file1, privateKey)
	if err != nil {
		log.Fatal(err)
	}

}

func readprivateKey(path string) (privateKey *packet.PrivateKey, err error) {
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	pr := packet.NewReader(r)

	entity, err := openpgp.ReadEntity(pr)
	if err != nil {
		return nil, err
	}

	privateKey = entity.PrivateKey

	return
}

func writeprivateKey(path string, privateKey *packet.PrivateKey) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = privateKey.Serialize(f)
	if err != nil {
		return err
	}

	return nil
}
