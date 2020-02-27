package main

import (
	"bytes"
	"crypto"
	"log"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

func main() {
	file1 := "test1.key"
	file2 := "test2.key"
	file3 := "test3.key"
	file4 := "test4.key"

	keyring, err := readKeyring(file1)
	if err != nil {
		entity, err := openpgp.NewEntity("name", "comment", "email@domain.org", &packet.Config{
			DefaultHash:   crypto.SHA256,
			DefaultCipher: packet.CipherAES256,
		})
		if err != nil {
			log.Fatal(err)
		}
		keyring = append(keyring, entity)
	}

	// log.Printf("Key[file2]: %#v", keyring[0])
	// log.Printf("Key[file2].PrivateKey.CreationTime: %s", keyring[0].PrivateKey.CreationTime.Format(time.RFC3339Nano))
	// log.Printf("Key[file2].PrimaryKey.CreationTime: %s", keyring[0].PrimaryKey.CreationTime.Format(time.RFC3339Nano))
	err = writeKeyring(file2, keyring[0])
	if err != nil {
		log.Fatal(err)
	}

	keyring, err = readKeyring(file2)
	if err != nil {
		log.Fatal(err)
	}

	// log.Printf("Key[file1]: %#v", keyring[0])
	// log.Printf("Key[file1].PrivateKey.CreationTime: %s", keyring[0].PrivateKey.CreationTime.Format(time.RFC3339Nano))
	// log.Printf("Key[file1].PrimaryKey.CreationTime: %s", keyring[0].PrimaryKey.CreationTime.Format(time.RFC3339Nano))
	err = writeKeyring(file3, keyring[0])
	if err != nil {
		log.Fatal(err)
	}

	keyring, err = readKeyring(file3)
	if err != nil {
		log.Fatal(err)
	}

	err = writeKeyring(file4, keyring[0])
	if err != nil {
		log.Fatal(err)
	}

	keyring, err = readKeyring(file4)
	if err != nil {
		log.Fatal(err)
	}

	err = writeKeyring(file1, keyring[0])
	if err != nil {
		log.Fatal(err)
	}

}

func readKeyring(path string) (keyring openpgp.EntityList, err error) {
	r, err := os.Open(path)
	if err != nil {
		return openpgp.EntityList{}, err
	}
	defer r.Close()

	return openpgp.ReadArmoredKeyRing(r)
}

func writeKeyring(path string, entity *openpgp.Entity) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	serializedEntity := bytes.NewBuffer(nil)

	err = entity.SerializePrivate(serializedEntity, nil)
	if err != nil {
		return err
	}

	headers := map[string]string{"Version": "GnuPG v1"}

	w, err := armor.Encode(f, openpgp.PrivateKeyType, headers)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(serializedEntity.Bytes())
	if err != nil {
		return err
	}

	return nil
}
