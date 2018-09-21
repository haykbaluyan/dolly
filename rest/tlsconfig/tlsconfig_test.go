package tlsconfig_test

import (
	"crypto/tls"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-phorce/dolly/rest/tlsconfig"
	"github.com/go-phorce/dolly/testify"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_BuildFromFiles(t *testing.T) {
	pemCert, pemKey, err := testify.MakeSelfCertRSAPem(1)
	require.NoError(t, err)
	require.NotNil(t, pemCert)
	require.NotNil(t, pemKey)

	tmpDir := filepath.Join(os.TempDir(), "tests", "tlsconfig")
	err = os.MkdirAll(tmpDir, os.ModePerm)
	require.NoError(t, err)

	pemFile := filepath.Join(tmpDir, "BuildFromFiles.pem")
	keyFile := filepath.Join(tmpDir, "BuildFromFiles-key.pem")

	err = ioutil.WriteFile(pemFile, pemCert, os.ModePerm)
	require.NoError(t, err)
	err = ioutil.WriteFile(keyFile, pemKey, os.ModePerm)
	require.NoError(t, err)

	cfg, err := tlsconfig.BuildFromFiles(pemFile, keyFile, "", true)
	assert.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, tls.RequireAndVerifyClientCert, cfg.ClientAuth)
}