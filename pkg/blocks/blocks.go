package blocks

import (
	"sort"

	"github.com/oklog/ulid"
	"github.com/thanos-io/thanos/pkg/block/metadata"
)

// LightMeta a simplified version of metadata.Meta that contains only the required attributes
type LightMeta struct {
	ULID       ulid.ULID
	MinTime    int64
	MaxTime    int64
	Resolution int64
	NumSamples uint64
}

// NewLightMeta returns a LightMeta from a Thanos metadata struct
func NewLightMeta(m metadata.Meta) LightMeta {
	return LightMeta{
		ULID:       m.ULID,
		MinTime:    m.MinTime,
		MaxTime:    m.MaxTime,
		Resolution: m.Thanos.Downsample.Resolution,
		NumSamples: m.Stats.NumSamples,
	}
}

// Blocks an array of metadata
type Blocks []LightMeta

func (blk Blocks) Len() int {
	return len(blk)
}

func (blk Blocks) Swap(i, j int) {
	blk[i], blk[j] = blk[j], blk[i]
}

func (blk Blocks) Less(i, j int) bool {
	if blk[i].MinTime < blk[j].MinTime {
		return true
	} else if blk[i].MinTime == blk[j].MinTime {
		if blk[i].Resolution < blk[j].Resolution {
			return true
		} else if blk[i].Resolution == blk[j].Resolution {
			return blk[i].NumSamples > blk[j].NumSamples
		}
	}
	return false
}

func equalTime(a, b LightMeta) bool {
	return a.MinTime == b.MinTime && a.MaxTime == b.MaxTime
}

// DropOverlappingBlocks drops overlapping blocks while keeping the ones with a higher resolution
func (blk *Blocks) DropOverlappingBlocks() {
	// Sort elements to be able to do in-place deduplication
	sort.Sort(blk)
	j := 0
	for i := 1; i < len(*blk); i++ {
		if equalTime((*blk)[j], (*blk)[i]) {
			continue
		}
		j++
		(*blk)[j] = (*blk)[i]
	}
	*blk = (*blk)[:j+1]
}
