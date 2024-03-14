package format

import (
	"github.com/osmanemek/vdk/av/avutil"
	"github.com/osmanemek/vdk/format/aac"
	"github.com/osmanemek/vdk/format/flv"
	"github.com/osmanemek/vdk/format/mp4"
	"github.com/osmanemek/vdk/format/rtmp"
	"github.com/osmanemek/vdk/format/rtsp"
	"github.com/osmanemek/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
