package support

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/go-playground/assert/v2"
	"strings"
	"testing"
)

const (
	testAlgorithmPrivateKeyStr = `-----BEGIN TESTING KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZUJN33V+dSfvd
fL0Mu+39XrZNXFFMQSy1V15FpncHeV47SmV0TzTqZc7hHB0ddqAdDi8Z5k3TKqb7
6sOwYr5TcAfuR6PIPaleyE0/0KrljBum2Isa2Nyq7Dgc3ElBQ6YN4l/a+DpvKaz1
FSKmKrhLNskqokWVSlu4g8OlKlbPXQ9ibII14MZRQrrkTmHYHzfi7GXXM0thAKuR
0HNvyhTHBh4/lrYM3GaMvmWwkwvsMavnOex6+eioZHBOb1/EIZ/LzC6zuHArPpyW
3daGaZ1rtQB1vVzTyERAVVFsXXgBHvfFud3w3ShsJYk8JvMwK2RpJ5/gV0QSARcm
LDRUAlPzAgMBAAECggEBAMc7rDeUaXiWv6bMGbZ3BTXpg1FhdddnWUnYE8HfX/km
OFI7XtBHXcgYFpcjYz4D5787pcsk7ezPidAj58zqenuclmjKnUmT3pfbI5eCA2v4
C9HnbYDrmUPK1ZcADtka4D6ScDccpNYNa1g2TFHzkIrEa6H+q7S3O2fqxY/DRVtN
0JIXalBb8daaqL5QVzSmM2BMVnHy+YITJWIkP2a3pKs9C0W65JGDsnG0wVrHinHF
+cnhFZIbaPEI//DAFMc9NkrWOKVRTEgcCUxCFaHOZVNxDWZD7A2ZfJB2rK6eg//y
gEiFDR2h6mTaDowMB4YF2n2dsIO4/dCG8vPHI20jn4ECgYEA/ZGu6lEMlO0XZnam
AZGtiNgLcCfM/C2ZERZE7QTRPZH1WdK92Al9ndldsswFw4baJrJLCmghjF/iG4zi
hhBvLnOLksnZUfjdumxoHDWXo2QBWbI5QsWIE7AuTiWgWj1I7X4fCXSQf6i+M/y2
6TogQ7d0ANpZFyOkTNMn/tiJvLECgYEA22XqlamG/yfAGWery5KNH2DGlTIyd6xJ
WtJ9j3jU99lZ0bCQ5xhiBbU9ImxCi3zgTsoqLWgA/p00HhNFNoUcTl9ofc0G3zwT
D1y0ZzcnVKxGJdZ6ohW52V0hJStAigtjYAsUgjm7//FH7PiQDBDP1Wa6xSRkDQU/
aSbQxvEE8+MCgYEA3bb8krW7opyM0XL9RHH0oqsFlVO30Oit5lrqebS0oHl3Zsr2
ZGgoBlWBsEzk3UqUhTFwm/DhJLTSJ/TQPRkxnhQ5/mewNhS9C7yua7wQkzVmWN+V
YeUGTvDGDF6qDz12/vJAgSwDDRym8x4NcXD5tTw7mmNRcwIfL22SkysThIECgYAV
BgccoEoXWS/HP2/u6fQr9ZIR6eV8Ij5FPbZacTG3LlS1Cz5XZra95UgebFFUHHtC
EY1JHJY7z8SWvTH8r3Su7eWNaIAoFBGffzqqSVazfm6aYZsOvRY6BfqPHT3p/H1h
Tq6AbBffxrcltgvXnCTORjHPglU0CjSxVs7awW3AEQKBgB5WtaC8VLROM7rkfVIq
+RXqE5vtJfa3e3N7W3RqxKp4zHFAPfr82FK5CX2bppEaxY7SEZVvVInKDc5gKdG/
jWNRBmvvftZhY59PILHO2X5vO4FXh7suEjy6VIh0gsnK36mmRboYIBGsNuDHjXLe
BDa+8mDLkWu5nHEhOxy2JJZl
-----END TESTING KEY-----`
	testAlgorithmExpectSignature = "BKyAfU4iMCuvXMXS0Wzam3V/cnxZ+JaqigPM5OhljS2iOT95OO6Fsuml2JkFANJU9K6q9bLlDhPXuoV" +
		"z+pp4hAm6pHU4ld815U4jsKu1RkyaII+1CYBUYC8TK0XtJ8FwUXXz8vZHh58rrAVN1XwNyvD1vfpxrMT4SL536GLwvpUHlCqIMzoZUguLl" +
		"i/K8V29QiOhuH6IEqLNJn8e9b3nwNcQ7be3CzYGpDAKBfDGPCqCv8Rw5zndhlffk2FEA70G4hvMwe51qMN/RAJbknXG23bSlObuTCN7Ndj" +
		"1aJGH6/L+hdwfLpUtJm4QYVazzW7DFD27EpSQEqA8bX9+8m1rLg=="
)

func getPrivateKey() *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(testingKey(testAlgorithmPrivateKeyStr)))
	if block == nil {
		panic("decode private key err")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		panic(fmt.Errorf("parse private key err:%s", err.Error()))
	}
	testAlgorithmPrivateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		panic(fmt.Errorf("%s is not rsa private key", testingKey(testAlgorithmPrivateKeyStr)))
	}

	return testAlgorithmPrivateKey
}

func testingKey(s string) string { return strings.ReplaceAll(s, "TESTING KEY", "PRIVATE KEY") }

func TestSha256WithRsa(t *testing.T) {
	testAlgorithmPrivateKey := getPrivateKey()
	type args struct {
		source     string
		privateKey *rsa.PrivateKey
	}
	tests := []struct {
		name          string
		args          args
		wantSignature string
		wantErr       bool
	}{
		{
			name: "sign uniformMessage in sha256withrsa success",
			args: args{
				source:     "source",
				privateKey: testAlgorithmPrivateKey,
			},
			wantSignature: testAlgorithmExpectSignature,
			wantErr:       false,
		},
		{
			name: "sign uniformMessage in sha256withrsa err",
			args: args{
				source:     "source",
				privateKey: nil,
			},
			wantErr:       true,
			wantSignature: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSignature, _ := SignSHA256WithRSA(tt.args.source, tt.args.privateKey)
			assert.Equal(t, gotSignature, tt.wantSignature)
			//if gotSignature != tt.wantSignature {
			//	t.Errorf("SignSHA256WithRSA() gotSignature = %v, want %v", gotSignature, tt.wantSignature)
			//}
		})
	}
}
