package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
)

func main() {
	// Connecting with a custom root-certificate set.
	host := flag.String("host", "beta.soluspay.net", "The port you gonna use")
	port := flag.String("port", "443", "tls port")

	flag.Parse()

	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIFWzCCBEOgAwIBAgISA/aBcSpT2ONrmb0YC2Q21jOAMA0GCSqGSIb3DQEBCwUA
MEoxCzAJBgNVBAYTAlVTMRYwFAYDVQQKEw1MZXQncyBFbmNyeXB0MSMwIQYDVQQD
ExpMZXQncyBFbmNyeXB0IEF1dGhvcml0eSBYMzAeFw0yMDAxMTcxNTQ2MzFaFw0y
MDA0MTYxNTQ2MzFaMBwxGjAYBgNVBAMTEWJldGEuc29sdXNwYXkubmV0MIIBIjAN
BgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzP2+Rt4Ha1qhdbinPrk1A5qvlREx
ZqLmXUctS3M4n6d6EFtzBl+ZPVl7toSNhvtHdwq5sszhBa7d3pMKzh4U1f8TAJfs
VEtiz7hSh8vMyKNESr/MRu1/bBx15jH7MtixPzAtU+oneWNSJXNyhQ0EZ9q8xz5A
b3VWXefQlAbIu6Kg1a8bUjTu50I6VsmX+K361IkszOH72WtbYvDHrpHcpaMojcyz
8NQ2Rz3Xxla8YG52TuBuE13jSmyacHaS3kFPTt1pOhQGlJFLSuxCsc6YzaR9n69n
LhkbssyrmemislVVXgMGrT4qE/pcAkT0y1iLWt/XpEVCpAnPnIzYekp9YQIDAQAB
o4ICZzCCAmMwDgYDVR0PAQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggr
BgEFBQcDAjAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQMR1lppxK2sCi3kxYkFNla
ulo3HDAfBgNVHSMEGDAWgBSoSmpjBH3duubRObemRWXv86jsoTBvBggrBgEFBQcB
AQRjMGEwLgYIKwYBBQUHMAGGImh0dHA6Ly9vY3NwLmludC14My5sZXRzZW5jcnlw
dC5vcmcwLwYIKwYBBQUHMAKGI2h0dHA6Ly9jZXJ0LmludC14My5sZXRzZW5jcnlw
dC5vcmcvMBwGA1UdEQQVMBOCEWJldGEuc29sdXNwYXkubmV0MEwGA1UdIARFMEMw
CAYGZ4EMAQIBMDcGCysGAQQBgt8TAQEBMCgwJgYIKwYBBQUHAgEWGmh0dHA6Ly9j
cHMubGV0c2VuY3J5cHQub3JnMIIBBQYKKwYBBAHWeQIEAgSB9gSB8wDxAHYAXqdz
+d9WwOe1Nkh90EngMnqRmgyEoRIShBh1loFxRVgAAAFvtGYokQAABAMARzBFAiBF
N4BqBVwT67dVaReRQMmZKscxYSjG6pxAfnHbzUM8BwIhAORowQ5x2vHhPCBtKwc+
8DVPNEszguOg2VdHcIkoDjx4AHcAsh4FzIuizYogTodm+Su5iiUgZ2va+nDnsklT
Le+LkF4AAAFvtGYofgAABAMASDBGAiEA2iO48gjVFFPSEPtsamuvuAGBPRZGbp9i
LeVM5nS9rzcCIQDDqevqh244ei6IhCUZrbNQiJdMHFXz+lbPpTKG6NylSjANBgkq
hkiG9w0BAQsFAAOCAQEAM1lF562hWsGS+Ytji9iQflRip5ZLx20CmfBzE7c7T7T+
3rpIyB3JGYRFz5HTofqbOI0i580E9AYKLrdM46sZ1erQOQhAMzgFhyTIewy1vQ5b
5zJX2qtfLyC7FNyqGyeFBKGHtxkJFh9P+DZnKEbSnq5WYdDjx8Vcf2RdxhuGk8QG
BhdUWTFeGQFiM0I6URWAtzZMXXEJG62kdh4106emuAD9ORsG4tQtIiBQxlZAy/wd
qGueJsDbEysXiq1ickIr0r5MUzWl11oVbrTY8Nd/ILBgYrJFC+Fpu9LUBUdAOPEw
ulncbpkQRfCz8AFs5YZE2De2sszQuC/FznFY6K/J2w==
-----END CERTIFICATE-----`

	fmt.Printf("The root cert is: %v", rootPEM)

	secondCert := `
-----BEGIN CERTIFICATE-----
MIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/
MSQwIgYDVQQKExtEaWdpdGFsIFNpZ25hdHVyZSBUcnVzdCBDby4xFzAVBgNVBAMT
DkRTVCBSb290IENBIFgzMB4XDTE2MDMxNzE2NDA0NloXDTIxMDMxNzE2NDA0Nlow
SjELMAkGA1UEBhMCVVMxFjAUBgNVBAoTDUxldCdzIEVuY3J5cHQxIzAhBgNVBAMT
GkxldCdzIEVuY3J5cHQgQXV0aG9yaXR5IFgzMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEAnNMM8FrlLke3cl03g7NoYzDq1zUmGSXhvb418XCSL7e4S0EF
q6meNQhY7LEqxGiHC6PjdeTm86dicbp5gWAf15Gan/PQeGdxyGkOlZHP/uaZ6WA8
SMx+yk13EiSdRxta67nsHjcAHJyse6cF6s5K671B5TaYucv9bTyWaN8jKkKQDIZ0
Z8h/pZq4UmEUEz9l6YKHy9v6Dlb2honzhT+Xhq+w3Brvaw2VFn3EK6BlspkENnWA
a6xK8xuQSXgvopZPKiAlKQTGdMDQMc2PMTiVFrqoM7hD8bEfwzB/onkxEz0tNvjj
/PIzark5McWvxI0NHWQWM6r6hCm21AvA2H3DkwIDAQABo4IBfTCCAXkwEgYDVR0T
AQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAYYwfwYIKwYBBQUHAQEEczBxMDIG
CCsGAQUFBzABhiZodHRwOi8vaXNyZy50cnVzdGlkLm9jc3AuaWRlbnRydXN0LmNv
bTA7BggrBgEFBQcwAoYvaHR0cDovL2FwcHMuaWRlbnRydXN0LmNvbS9yb290cy9k
c3Ryb290Y2F4My5wN2MwHwYDVR0jBBgwFoAUxKexpHsscfrb4UuQdf/EFWCFiRAw
VAYDVR0gBE0wSzAIBgZngQwBAgEwPwYLKwYBBAGC3xMBAQEwMDAuBggrBgEFBQcC
ARYiaHR0cDovL2Nwcy5yb290LXgxLmxldHNlbmNyeXB0Lm9yZzA8BgNVHR8ENTAz
MDGgL6AthitodHRwOi8vY3JsLmlkZW50cnVzdC5jb20vRFNUUk9PVENBWDNDUkwu
Y3JsMB0GA1UdDgQWBBSoSmpjBH3duubRObemRWXv86jsoTANBgkqhkiG9w0BAQsF
AAOCAQEA3TPXEfNjWDjdGBX7CVW+dla5cEilaUcne8IkCJLxWh9KEik3JHRRHGJo
uM2VcGfl96S8TihRzZvoroed6ti6WqEBmtzw3Wodatg+VyOeph4EYpr/1wXKtx8/
wApIvJSwtmVi4MFU5aMqrSDE6ea73Mj2tcMyo5jMd6jmeWUHK8so/joWUoHOUgwu
X4Po1QYz+3dszkDqMp4fklxBwXRsW10KXzPMTZ+sOPAveyxindmjkW8lGy+QsRlG
PfZ+G6Z6h7mjem0Y+iWlkYcV4PIWL1iwBi8saCbGS5jN2p8M+X+Q7UNKEkROb3N6
KOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==
-----END CERTIFICATE-----`

	// First, create the set of root certificates. For this example we only
	// have one. It's also possible to omit this in order to use the
	// default root set of the current operating system.
	roots := x509.NewCertPool()
	// ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	// if !ok {
	// 	panic("failed to parse root certificate")
	// }
	ok := roots.AppendCertsFromPEM([]byte(secondCert))
	if !ok {
		panic("error in parsing the another cert")
	}

	conn, err := tls.Dial("tcp", *host+":"+*port, &tls.Config{
		RootCAs: roots,
	})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}
	conn.Close()
}
