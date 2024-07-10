package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"

	"github.com/spf13/afero"
	"github.com/taubyte/tau/tools/spore-drive/config"
)

func generateSSHKeyPair(bits int) (privateKey []byte, publicKey []byte, err error) {
	// Generate an RSA private key
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// Encode the private key to PEM format
	privPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	// Generate the public key
	pub, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, nil, err
	}

	// Encode the public key to PEM format
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pub,
		},
	)

	return privPEM, pubPEM, nil
}

func main() {
	p, err := config.New(afero.NewBasePathFs(afero.NewOsFs(), "fixtures"), "/config")
	if err != nil {
		panic(err)
	}

	err = p.Cloud().Domain().SetRoot("test.com")
	if err != nil {
		panic(err)
	}

	err = p.Cloud().Domain().SetGenerated("gtest.com")
	if err != nil {
		panic(err)
	}

	err = p.Cloud().Domain().Validation().Generate()
	if err != nil {
		panic(err)
	}

	err = p.Cloud().P2P().Swarm().Generate()
	if err != nil {
		panic(err)
	}

	err = p.Auth().Add("main").SetUsername("tau")
	if err != nil {
		panic(err)
	}

	err = p.Auth().Add("main").SetPassword("testtest")
	if err != nil {
		panic(err)
	}

	err = p.Auth().Add("withkey").SetUsername("tau")
	if err != nil {
		panic(err)
	}

	err = p.Auth().Add("withkey").SetKey("keys/test.pem")
	if err != nil {
		panic(err)
	}

	_, pubKeyData, err := generateSSHKeyPair(256)
	if err != nil {
		panic(err)
	}

	pubKeyFile, err := p.Auth().Add("withkey").Open()
	if err != nil {
		panic(err)
	}
	defer pubKeyFile.Close()

	_, err = io.Copy(pubKeyFile, bytes.NewBuffer(pubKeyData))
	if err != nil {
		panic(err)
	}

	err = p.Shapes().Add("shape1").Services().Set("auth", "seer")
	if err != nil {
		panic(err)
	}

	err = p.Shapes().Add("shape1").Ports().Set("main", 4242)
	if err != nil {
		panic(err)
	}

	err = p.Shapes().Add("shape1").Ports().Set("lite", 4262)
	if err != nil {
		panic(err)
	}

	err = p.Shapes().Add("shape2").Services().Set("gateway", "patrick", "monkey")
	if err != nil {
		panic(err)
	}

	err = p.Shapes().Add("shape2").Ports().Set("main", 6242)
	if err != nil {
		panic(err)
	}

	err = p.Shapes().Add("shape2").Ports().Set("lite", 6262)
	if err != nil {
		panic(err)
	}

	err = p.Shapes().Add("shape2").Plugins().Set("plugin1@v0.1")
	if err != nil {
		panic(err)
	}

	host1 := p.Hosts().Add("host1")
	err = host1.Addresses().Add("1.2.3.4/24")
	if err != nil {
		panic(err)
	}

	err = host1.Addresses().Add("4.3.2.1/24")
	if err != nil {
		panic(err)
	}

	err = host1.SSH().SetFullAddress("1.2.3.4:4242")
	if err != nil {
		panic(err)
	}

	err = host1.SSH().Auth().Add("main")
	if err != nil {
		panic(err)
	}

	err = host1.SetLocation(1.25, 25.1)
	if err != nil {
		panic(err)
	}

	err = host1.Shapes().Add("shape1").SetKey("CAESQIWC2KRhsEexLpN4DsJwki4S56IN5IreCANf89+F+OpTWn7Tf+RwZnUbiZYdxsTFrbBJQ9S+A0oFp8a1SSAN2EE=")
	if err != nil {
		panic(err)
	}

	err = host1.Shapes().Add("shape2").SetKey("CAESQHLGyFbnI2GP7e3Gib9ut7IFDxrkbTbs7LFAJYhe0w0LXEtYrH7HyODglOFY3oXQ+kCfoFcvqvZnAD6K5UavO2c=")
	if err != nil {
		panic(err)
	}

	host2 := p.Hosts().Add("host2")
	err = host2.Addresses().Add("8.2.3.4/24")
	if err != nil {
		panic(err)
	}

	err = host2.Addresses().Add("4.3.2.8/24")
	if err != nil {
		panic(err)
	}

	err = host2.SSH().SetFullAddress("8.2.3.4:4242")
	if err != nil {
		panic(err)
	}

	err = host2.SSH().Auth().Add("main")
	if err != nil {
		panic(err)
	}

	err = host2.SetLocation(1.25, 25.1)
	if err != nil {
		panic(err)
	}

	err = host2.Shapes().Add("shape1").SetKey("CAESQDpF3eQuEbGsjSRkf3uE6E4SV3dvwSSMUcNJkimOUc0hO6gPoZjsq/NO/FwVz8FoZ4LG/5DSF2B/Rl+vJCNLlUI=")
	if err != nil {
		panic(err)
	}

	err = host2.Shapes().Add("shape2").SetKey("CAESQIA03gtBTeL8eYNQKcJ+VqKLgarHfofd5I/CV/zEsxHiqfihV9ZXjl0qtaTPEWExBgqRn+w2YLD6FQy8zBdEabI=")
	if err != nil {
		panic(err)
	}

	p.Sync()
}