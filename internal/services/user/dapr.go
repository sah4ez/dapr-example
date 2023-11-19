package user

import (
	"context"
	"encoding/json"

	"github.com/dapr/go-sdk/service/common"
	"github.com/rs/zerolog/log"
	"github.com/sah4ez/dapr-example/interfaces/types"
	"github.com/sah4ez/dapr-example/pkg/errors"
)

func (svc *Service) GetNameByIDHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.DaprInvocationErr
		return
	}

	log.Ctx(ctx).Info().
		Str("ContentType", in.ContentType).
		Str("Verb", in.Verb).
		Str("QueryString", in.QueryString).
		Str("data", string(in.Data)).
		Msg("call")

	req := struct {
		ID int `json:"id"`
	}{}
	err = json.Unmarshal(in.Data, &req)
	if err != nil {
		return
	}
	var user types.User
	user, err = svc.GetNameByID(ctx, req.ID)
	if err != nil {
		return
	}

	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		return
	}
	// do something with the invocation here
	out = &common.Content{
		Data:        data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
