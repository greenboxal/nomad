package api

// Resources encapsulates the required resources of
// a given task or task group.
type Resources struct {
	CPU      int
	MemoryMB int
	DiskMB   int
	IOPS     int
	Networks []*NetworkResource
}

// Merge merges this resource with another resource.
func (r *Resources) Merge(other *Resources) {
	if other.CPU != 0 {
		r.CPU = other.CPU
	}
	if other.MemoryMB != 0 {
		r.MemoryMB = other.MemoryMB
	}
	if other.DiskMB != 0 {
		r.DiskMB = other.DiskMB
	}
	if other.IOPS != 0 {
		r.IOPS = other.IOPS
	}
	if len(other.Networks) != 0 {
		r.Networks = other.Networks
	}
}

type Port struct {
	Label string
	Value int
}

// NetworkResource is used to describe required network
// resources of a given task.
type NetworkResource struct {
	Public        bool
	CIDR          string
	ReservedPorts []Port
	DynamicPorts  []Port
	IP            string
	MBits         int
}
