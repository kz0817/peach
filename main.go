package main

import (
    "flag"
    "os"
    "fmt"
    "encoding/binary"
)

const (
    POS_PART_TABLE = 0x01be
    NUM_PRIMARY_PARTITIONS = 4
    POS_SIGNATURE = 0x01fe
    SECTOR_SIZE = 512
)

type PartitionTableEntry struct {
    BootFlag uint8
    StartCHS [3]uint8
    Type uint8
    EndCHS [3]uint8
    StartLBA uint32
    NumSectors uint32
}

type Partition struct {
    StartSector uint64
    NumSectors uint64
    StartOffset uint64
    Size uint64
}

func openDrive(drivePath string) *os.File {
    drive, err := os.Open(drivePath)
    if err != nil {
        fmt.Printf("Failed to open: %s\n", drivePath)
        panic(err)
    }
    return drive
}

func readPartitionRecords(
  drive *os.File,
  partitionEntries *[NUM_PRIMARY_PARTITIONS]PartitionTableEntry) {

    drive.Seek(POS_PART_TABLE, 0)

    for i := 0; i < NUM_PRIMARY_PARTITIONS; i++ {
        err := binary.Read(drive, binary.LittleEndian, &partitionEntries[i])
        if err != nil {
            fmt.Printf("Failed to read partion table: %d\n", i)
            panic(err)
        }
    }
}


func (part *Partition) Load(entry *PartitionTableEntry) {
    part.StartSector = uint64(entry.StartLBA)
    part.NumSectors = uint64(entry.NumSectors)
    part.StartOffset = uint64(entry.StartLBA * SECTOR_SIZE)
    part.Size = uint64(entry.NumSectors * SECTOR_SIZE)
}


func main() {
    drivePath := flag.String("target", "", "target file")
    flag.Parse()
    fmt.Printf("Target drive: %s\n", *drivePath)

    drive := openDrive(*drivePath)

    var partitionEntries [NUM_PRIMARY_PARTITIONS]PartitionTableEntry
    readPartitionRecords(drive, &partitionEntries)


    var partitions []Partition
    for i := 0; i < NUM_PRIMARY_PARTITIONS; i++ {
        partitionEntry := &partitionEntries[i]
        if partitionEntry.StartLBA != 0 {
            var part Partition
            part.Load(partitionEntry)
            partitions = append(partitions, part)
        }
    }

    for i, p := range partitions {
        fmt.Printf("[%d] %v\n", i, p)
    }

}
