package shakeAround

type ShakeAround struct {
	*Client

	Device      *DeviceClient
	Page        *PageClient
	Material    *MaterialClient
	Group       *GroupClient
	Relation    *RelationClient
	Stats       *StatsClient
}
