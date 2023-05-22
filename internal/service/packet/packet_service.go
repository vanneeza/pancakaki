package packetservice

import (
	entity "pancakaki/internal/domain/entity/packet"
	packetrepository "pancakaki/internal/repository/packet"
)

type PacketService interface {
	InsertPacket(newPacket *entity.Packet) (*entity.Packet, error)
	UpdatePacket(updatePacket *entity.Packet) (*entity.Packet, error)
	DeletePacket(deletePacket *entity.Packet) error
	FindpacketById(id int) (*entity.Packet, error)
	FindPacketByName(name string) (*entity.Packet, error)
	FindAllPacket() ([]entity.Packet, error)
}

type packetService struct {
	packetRepo packetrepository.PacketRepository
}

// DeletePacket implements PacketService
func (s *packetService) DeletePacket(deletePacket *entity.Packet) error {
	return s.packetRepo.DeletePacket(deletePacket)
}

// FindAllPacket implements PacketService
func (s *packetService) FindAllPacket() ([]entity.Packet, error) {
	return s.packetRepo.FindAllPacket()
}

// FindPacketByName implements PacketService
func (s *packetService) FindPacketByName(name string) (*entity.Packet, error) {
	return s.packetRepo.FindPacketByName(name)
}

// FindpacketById implements PacketService
func (s *packetService) FindpacketById(id int) (*entity.Packet, error) {
	return s.packetRepo.FindpacketById(id)
}

// InsertPacket implements PacketService
func (s *packetService) InsertPacket(newPacket *entity.Packet) (*entity.Packet, error) {
	return s.packetRepo.InsertPacket(newPacket)
}

// UpdatePacket implements PacketService
func (s *packetService) UpdatePacket(updatePacket *entity.Packet) (*entity.Packet, error) {
	return s.packetRepo.UpdatePacket(updatePacket)
}

func NewPacketService(packetRepo packetrepository.PacketRepository) PacketService {
	return &packetService{packetRepo: packetRepo}
}
