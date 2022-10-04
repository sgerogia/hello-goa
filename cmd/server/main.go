package main

import (
	"context"
	"flag"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/sgerogia/hello-goa/logic"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	math "github.com/sgerogia/hello-goa/gen/math"
	token "github.com/sgerogia/hello-goa/gen/token"
)

type domainInfo struct {
	defaultAddr string
	domain      string
	port        string
	secure      bool
}

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "local", "Server host (valid values: local)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		privateKeyF = flag.String("private-key", "", "The RSA private key file for JWT signing in PEM format\nIt must not be password-protected.")
		publicKeyF  = flag.String("public-key", "", "The RSA public key file for JWT verification in PEM format")
		jwtExpiryF = flag.Int("jwt-expiry", 60, "The lifetime of the generated JWT in minutes\nDefaults to 60 (1h)")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[mathapi] ", log.Ltime)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	var u *url.URL
	startServer := true
	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "local":
		{

			d := domainInfo{
				defaultAddr: "http://localhost:8080",
				port: *httpPortF,
				domain: *domainF,
				secure: *secureF,
			}

			u = canonicalAddress(d, logger)
		}

	default:
		startServer = false
		logger.Fatalf("invalid host argument: %q (valid hosts: local)\n", *hostF)
	}

	logger.Printf("URL: %s", u)

	if startServer {
		// Initialize the services.
		var (
			tokenSvc token.Service
			mathSvc  math.Service
		)

		pk, err := os.ReadFile(*publicKeyF)
		check(err, "invalid public key file", logger)

		k, err := os.ReadFile(*privateKeyF)
		check(err, "invalid private key file", logger)

		pKey, err := jwtgo.ParseRSAPublicKeyFromPEM(pk)
		check(err, "invalid public key", logger)

		key, err := jwtgo.ParseRSAPrivateKeyFromPEM(k)
		check(err, "invalid private key", logger)

		{
			tokenSvc = logic.NewToken(logger, u.String(), key, jwtExpiryF)
			mathSvc = logic.NewMath(logger, u.String(), pKey, jwtExpiryF)
		}

		// Wrap the services in endpoints that can be invoked from other services
		// potentially running in different processes.
		var (
			tokenEndpoints *token.Endpoints
			mathEndpoints  *math.Endpoints
		)
		{
			tokenEndpoints = token.NewEndpoints(tokenSvc)
			mathEndpoints = math.NewEndpoints(mathSvc)
		}
		handleHTTPServer(ctx, u, tokenEndpoints, mathEndpoints, &wg, errc, logger, *dbgF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}

// Populates a canonical URL based on a struct of commande line args
func canonicalAddress(d domainInfo, logger *log.Logger) *url.URL {
	u, err := url.Parse(d.defaultAddr)
	check(err, fmt.Sprintf("invalid URL %#v: %s\n", d.defaultAddr, err), logger)
	if d.secure {
		u.Scheme = "https"
	}
	if d.domain != "" {
		u.Host = d.domain
	}
	if d.port != "" {
		h, _, err := net.SplitHostPort(u.Host)
		check(err, fmt.Sprintf("invalid URL %#v: %s\n", u.Host, err), logger)
		u.Host = net.JoinHostPort(h, d.port)
	} else if u.Port() == "" {
		u.Host = net.JoinHostPort(u.Host, "80")
	}

	return u
}

func check(e error, msg string, logger *log.Logger) {
	if e != nil {
		logger.Fatalf(msg)
		panic(any(e))
	}
}

