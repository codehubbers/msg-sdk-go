package identity

// GetLocalKeypair retrieves the local keypair for the Noise protocol.
// Actual implementation is pending.
// TODO: Implement keypair retrieval logic
import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/crypto/curve25519"
)

// Keypair holds a Noise-compatible X25519 keypair
type Keypair struct {
	PrivateKey [32]byte
	PublicKey  [32]byte
}

// GenerateKeypair creates a new X25519 keypair
func GenerateKeypair() (*Keypair, error) {
	var priv [32]byte
	_, err := rand.Read(priv[:])
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)
	}

	pub, err := curve25519.X25519(priv[:], curve25519.Basepoint)
	if err != nil {
		return nil, fmt.Errorf("failed to derive public key: %v", err)
	}

	var pub32 [32]byte
	copy(pub32[:], pub)

	return &Keypair{
		PrivateKey: priv,
		PublicKey:  pub32,
	}, nil
}

// SaveKeypair stores the keypair to a file as JSON
func SaveKeypair(kp *Keypair, filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(kp)
}
func DefaultKeypairPath() string {
	usr, err := user.Current()
	if err != nil {
		
		return filepath.Join(os.TempDir(), "keypair.json")
	}

	dir := filepath.Join(usr.HomeDir, ".msg-sdk")
	os.MkdirAll(dir, 0700) // silently create dir if missing

	return filepath.Join(dir, "keypair.json")
}
// LoadKeypair loads a keypair from file
func LoadKeypair(filepath string) (*Keypair, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var kp Keypair
	err = json.NewDecoder(f).Decode(&kp)
	if err != nil {
		return nil, err
	}
	return &kp, nil
}
