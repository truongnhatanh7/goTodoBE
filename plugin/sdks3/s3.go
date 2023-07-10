package sdks3

import (
	"flag"

	"github.com/200Lab-Education/go-sdk/logger"
	"github.com/truongnhatanh7/goTodoBE/component/uploadprovider"
)

type S3PluginOpt struct {
	Prefix     string
	BucketName string
	Region     string
	ApiKey     string
	Secret     string
	Domain     string
}

type s3Plugin struct {
	name       string
	logger     logger.Logger
	S3provider *uploadprovider.S3Provider
	*S3PluginOpt
}

func NewS3Service(name, prefix string) *s3Plugin {
	return &s3Plugin{
		S3PluginOpt: &S3PluginOpt{
			Prefix:     prefix,
		},
		name: name,
	}
}

func (s3 *s3Plugin) GetPrefix() string {
	return s3.Prefix
}

func (s3 *s3Plugin) Name() string {
	return s3.name
}

func (s3 *s3Plugin) InitFlags() {
	flag.StringVar(&s3.ApiKey, s3.Prefix+"api-key", "", "s3 programmatic user api key")
	flag.StringVar(&s3.Secret, s3.Prefix+"secret-key", "", "s3 secret key")
	flag.StringVar(&s3.BucketName, s3.Prefix+"bucket-name", "", "s3 bucket name")
	flag.StringVar(&s3.Region, s3.Prefix+"region", "", "s3 region")
	flag.StringVar(&s3.Domain, s3.Prefix+"domain", "", "s3 domain")
}

func (s3 *s3Plugin) Configure() error {
	s3.S3provider = uploadprovider.NewS3Provider(
		s3.BucketName,
		s3.Region,
		s3.ApiKey,
		s3.Secret,
		s3.Domain,
	)
	// s3.logger.Info("s3 provider is connected")

	return nil
}

func (s3 *s3Plugin) Run() error {
	return s3.Configure()
}

func (s3 *s3Plugin) Stop() <-chan bool {
	c := make(chan bool)
	go func() {
		c <- true
	}()
	return c
}

func (s3 *s3Plugin) Get() interface{} {
	return s3
}

func (s3 *s3Plugin) GetBucketProvider() *uploadprovider.S3Provider {
	return s3.S3provider
}
