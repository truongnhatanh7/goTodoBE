package subscriber

import (
	"context"

	"log"

	goservice "github.com/200Lab-Education/go-sdk"
	"github.com/truongnhatanh7/goTodoBE/common"
	"github.com/truongnhatanh7/goTodoBE/common/asyncjob"
	"github.com/truongnhatanh7/goTodoBE/pubsub"
)

type subJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type pbEngine struct {
	serviceCtx goservice.ServiceContext
}

func NewEngine(serviceCtx goservice.ServiceContext) *pbEngine {
	return &pbEngine{serviceCtx: serviceCtx}
}

func (engine *pbEngine) Start() error {
	engine.startSubTopic(common.TopicUserLikedItem, true,
		IncreaseLikeCountAfterUserLikeItem(engine.serviceCtx),
		PushNotificationAfterUserLikeItem(engine.serviceCtx),
	)

	engine.startSubTopic(common.TopicUserUnlikedItem, true,
		DecreaseLikeCountAfterUserUnlikeItem(engine.serviceCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *pbEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, jobs ...subJob) error {
	ps := engine.serviceCtx.MustGet(common.PluginPubSub).(pubsub.PubSub)

	c, _ := ps.Subscribe(context.Background(), topic)

	for _, item := range jobs {
		log.Println("Setup subscriber for:", item.Title)
	}

	getJobHandler := func(job *subJob, message *pubsub.Message) asyncjob.JobHandler {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(jobs))

			for i := range jobs {
				jobHdl := getJobHandler(&jobs[i], msg)
				jobHdlArr[i] = asyncjob.NewJob(jobHdl, asyncjob.WithName(jobs[i].Title))
			}

			group := asyncjob.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
