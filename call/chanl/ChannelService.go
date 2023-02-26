package chanl

import "litecluster/data"

type ChannelService struct {
	channel IOChannel
}

func (service *ChannelService) FileChannel(path string) *ChannelService {
	dataBuf := &data.FileDataBuf{Path: path}
	bufChannel := &BufChannel{Adapter: dataBuf}
	service.channel = bufChannel
	return service
}

func (service *ChannelService) Build() IOChannel {
	service.channel.Init()
	return service.channel
}
