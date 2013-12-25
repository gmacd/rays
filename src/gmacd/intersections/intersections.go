package intersections

const (
	HIT_FROM_INSIDE = -1
	MISS            = 0
	HIT             = 1
)

type HitDetails struct {
	Result int
	Dist   float64
}

func NewHit(dist float64) HitDetails {
	return HitDetails{HIT, dist}
}

func NewHitFromInside(dist float64) HitDetails {
	return HitDetails{HIT_FROM_INSIDE, dist}
}

func NewMiss() HitDetails {
	return HitDetails{MISS, 0}
}

func (hd HitDetails) IsMiss() bool          { return hd.Result == MISS }
func (hd HitDetails) IsHit() bool           { return hd.Result == HIT }
func (hd HitDetails) IsHitFromInside() bool { return hd.Result == HIT_FROM_INSIDE }
func (hd HitDetails) IsAnyHit() bool        { return hd.Result != MISS }
