package main

import (
	"github.com/nelbrecht/gpghello/payload"
	"go.uber.org/zap"
	"golang.org/x/crypto/openpgp"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	keyring, err := openpgp.ReadArmoredKeyRing(payload.X().KeyRingReader)
	if err != nil {
		logger.Error("Read Armored Key Ring", zap.Error(err))
		return
	}
	logger.Info("Key Ring", zap.Any("keyring", keyring))
	entity, err := openpgp.CheckArmoredDetachedSignature(keyring,
		payload.GetExample().VerificationTarget, payload.X().Signature)
	if err != nil {
		logger.Error("Check Detached Signature", zap.Error(err))
		return
	}

	sugar.Infof("Entity: %+v\n", entity)
}
