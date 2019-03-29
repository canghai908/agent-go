package cmd

import (
	"math"
	"strconv"

	"github.com/shirou/gopsutil/mem"
)

//ByteSize float64
type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

var (
	TKB int64 = 1000
	TMB int64 = TKB * 1000
	TGB int64 = TMB * 1000
	TTB int64 = TGB * 1000
	TPB int64 = TTB * 1000
	TEB int64 = TPB * 1000
)

//GetMem mem
func GetMem(mtype string) string {
	v, _ := mem.VirtualMemory()
	switch mtype {
	case "total":
		memTotal := int64(math.Ceil(float64(v.Total) / float64(MB)))
		return strconv.FormatInt(memTotal, 10)
	case "used":
		return strconv.FormatInt(int64(math.Ceil(float64(v.Available)/float64(MB))), 10)
	case "usedper":
		return strconv.FormatFloat(v.UsedPercent, 'E', -1, 64)
	default:
		return ""
	}
	return "nil"

}
