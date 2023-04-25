package main

import (
	"os"
	"time"

	"github.com/quanghung97/go-elk-kafka/constants"
	config "github.com/quanghung97/kafka-go"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

var configKafka = &config.Kafka{
	KafkaUrl:          "kafka:29092",
	MinBytes:          5,
	MaxBytes:          10e6, // max 10MB
	MaxWait:           3 * time.Second,
	NumPartitions:     12,
	ReplicationFactor: 1,
}

var encoderConfig = ecszap.NewDefaultEncoderConfig()
var core = ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
var zapWriter = zap.New(core, zap.AddCaller()).With(zap.String("app", "hub-worker")).With(zap.String("environment", "local"))

func handleMsg(msg config.Message, err error) {
	logger := zapWriter.With(zap.String("topic", msg.Topic)).With(zap.Int("partition", msg.Partition)).With(zap.String("key", string(msg.Key))).With(zap.String("value", string(msg.Value)))
	logger.Info(constants.HUB_WORKER_LOGS_INFO + "save log successfully")
	// fmt.Printf("message at topic:%v partition:%v %s = %s\n", msg.Topic, msg.Partition, string(msg.Key), string(msg.Value))
}

func handleErr(msg string, moreMsg ...interface{}) {
	zapWriter.Error(constants.HUB_WORKER_LOGS_ERR+msg, zap.Any("more-err", moreMsg))
	// fmt.Printf(msg, a...)
	// fmt.Println()
}

func main() {
	configKafka.ReaderReceiveMessage("topic-hub-logs", "group-hub-logs", handleMsg, handleErr)
}
