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
)

type PartitionTable struct {
    BootFlag uint8
    StartCHS [3]uint8
    Type uint8
    EndCHS [3]uint8
    StartLBA uint32
    NumSectors uint32
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
  partitions *[NUM_PRIMARY_PARTITIONS]PartitionTable) {
    drive.Seek(POS_PART_TABLE, 0)
    for i := 0; i < NUM_PRIMARY_PARTITIONS; i++ {
        err := binary.Read(drive, binary.LittleEndian, &partitions[i])
        if err != nil {
            fmt.Printf("Failed to read partion table: %d\n", i)
            panic(err)
        }
    }
}


func main() {
    drivePath := flag.String("target", "", "target file")
    flag.Parse()
    fmt.Printf("Target drive: %s\n", *drivePath)

    drive := openDrive(*drivePath)

    var partitions [NUM_PRIMARY_PARTITIONS]PartitionTable
    readPartitionRecords(drive, &partitions)

    for i := 0; i < NUM_PRIMARY_PARTITIONS; i++ {
        fmt.Printf("%d, %v\n", i, partitions[i])
    }

}
