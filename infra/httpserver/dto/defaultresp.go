package dto

type DefaultResp struct {
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"message"`
	} `json:"error"`
}

type DefaultRespOption func(*DefaultResp)

func Code(code int) DefaultRespOption {
	return func(d *DefaultResp) {
		if d == nil {
			return
		}

		d.Error.Code = code
	}
}

func Message(msg string) DefaultRespOption {
	return func(d *DefaultResp) {
		if d == nil {
			return
		}

		d.Error.Msg = msg
	}
}

func (resp *DefaultResp) SetError(opts ...DefaultRespOption) {
	if resp == nil {
		return
	}

	for _, opt := range opts {
		opt(resp)
	}
}

type DefaultRespWithData struct {
	DefaultResp
	Data interface{} `json:"data"`
}

func (resp *DefaultRespWithData) SetData(data interface{}) {
	if resp == nil {
		return
	}

	resp.Data = data
}
