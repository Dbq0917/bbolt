//go:build !windows

package bbolt

import "golang.org/x/sys/unix"

// mlock locks memory of db file
func mlock(m *dbMmap, fileSize int) error {
	sizeToLock := fileSize
	if sizeToLock > m.size {
		// Can't lock more than mmaped slice
		sizeToLock = m.size
	}
	if err := unix.Mlock(m.dataref[:sizeToLock]); err != nil {
		return err
	}
	return nil
}

// munlock unlocks memory of db file
func munlock(m *dbMmap, fileSize int) error {
	if m == nil || m.dataref == nil {
		return nil
	}

	sizeToUnlock := fileSize
	if sizeToUnlock > m.size {
		// Can't unlock more than mmaped slice
		sizeToUnlock = m.size
	}

	if err := unix.Munlock(m.dataref[:sizeToUnlock]); err != nil {
		return err
	}
	return nil
}
