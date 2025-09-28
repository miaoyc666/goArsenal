package ecdsa

import (
    "crypto/ecdsa"
    "crypto/rand"
    "encoding/hex"

    "github.com/ethereum/go-ethereum/crypto"
)

/*
Description  : 椭圆曲线数字签名算法工具函数, 包含私钥序列化和反序列化
*/

type KeyManager struct {
    ServerPrivateKey *ecdsa.PrivateKey
    ClientPublicKey  *ecdsa.PublicKey
}

func NewKeyManager() (*KeyManager, error) {
    privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
    if err != nil {
       return nil, err
    }
    return &KeyManager{ServerPrivateKey: privateKey}, nil
}

func LoadKeyManager(privateKeyStr string) (*KeyManager, error) {
    privateKey, err := DeserializePrivateKey(privateKeyStr)
    if err != nil {
       return nil, err
    }
    return &KeyManager{ServerPrivateKey: privateKey}, nil
}

func (km *KeyManager) SetClientPublicKey(hexKey string) error {
    clientPublicKeyBytes, err := hex.DecodeString(hexKey)
    if err != nil {
       return err
    }
    clientPublicKey, err := crypto.UnmarshalPubkey(clientPublicKeyBytes)
    if err != nil {
       return err
    }
    km.ClientPublicKey = clientPublicKey
    return nil
}

func (km *KeyManager) GetServerPublicKeyHex() string {
    serverPublicKeyBytes := crypto.FromECDSAPub(&km.ServerPrivateKey.PublicKey)
    return hex.EncodeToString(serverPublicKeyBytes)
}

func (km *KeyManager) CalcSharedSecret() string {
    sharedX, _ := km.ServerPrivateKey.PublicKey.Curve.ScalarMult(km.ClientPublicKey.X, km.ClientPublicKey.Y, km.ServerPrivateKey.D.Bytes())
    return hex.EncodeToString(sharedX.Bytes())
}

func SerializePrivateKey(privateKey *ecdsa.PrivateKey) (string, error) {
    // 序列化： *ecdsa.PrivateKey -> []byte -> hex
    priBytes := crypto.FromECDSA(privateKey)
    return hex.EncodeToString(priBytes), nil
}

func DeserializePrivateKey(privateKeyStr string) (privateKey *ecdsa.PrivateKey, err error) {
    // 反序列化：hex -> []byte -> *ecdsa.PrivateKey
    priBytes, _ := hex.DecodeString(privateKeyStr)
    privateKey, err = crypto.ToECDSA(priBytes)
    if err != nil {
       return nil, err
    }
    return
}
