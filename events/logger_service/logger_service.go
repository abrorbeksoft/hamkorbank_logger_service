package logger_service

import (
	"context"
	"encoding/json"
	"fmt"
	"logger_service/pkg/logger"
)

func (t *triggerListener) ListenErrors(ctx context.Context, data []byte) error {
	var resp NotFound
	err := json.Unmarshal(data, &resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}
	t.log.Error("Not found", logger.String("Not found key", resp.NotFound))
	return nil
}

func (t *triggerListener) ListenInfo(ctx context.Context, data []byte) error {
	var resp Phone
	err := json.Unmarshal(data, &resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}

	t.log.Info("Data came from rest_service ", logger.Any("phone", resp))
	return nil
}

func (t *triggerListener) ListenDebug(ctx context.Context, data []byte) error {
	var resp Message
	err := json.Unmarshal(data, &resp)
	if err != nil {
		t.log.Error("error while consuming ", logger.Error(err))
		return err
	}
	str := fmt.Sprintf("Debug on data %s", resp.RecordId)
	t.log.Info(str, logger.Any("phone", resp))
	return nil
}

func (t *triggerListener) ListenAll(ctx context.Context, data []byte) error {
	t.log.Info("Listen all", logger.Any("data", string(data)))
	return nil
}
