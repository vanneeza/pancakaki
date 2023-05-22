package packetrepository

import (
	"database/sql"
	"fmt"
	entity "pancakaki/internal/domain/entity/packet"
)

type PacketRepository interface {
	InsertPacket(newPacket *entity.Packet) (*entity.Packet, error)
	UpdatePacket(updatePacket *entity.Packet) (*entity.Packet, error)
	DeletePacket(deletePacket *entity.Packet) error
	FindpacketById(id int) (*entity.Packet, error)
	FindPacketByName(name string) (*entity.Packet, error)
	FindAllPacket() ([]entity.Packet, error)
}

type packetRepository struct {
	db *sql.DB
}

// DeletePacket implements PacketRepository
func (repo *packetRepository) DeletePacket(deletePacket *entity.Packet) error {
	stmt, err := repo.db.Prepare("UPDATE tbl_packet SET is_delete = true WHERE id = $1")
	if err != nil {
		return fmt.Errorf("failed to delete packet : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(deletePacket.Id)
	if err != nil {
		return fmt.Errorf("failed to delete packet : %w", err)
	}

	return nil
}

// FindAllPacket implements PacketRepository
func (repo *packetRepository) FindAllPacket() ([]entity.Packet, error) {
	var packets []entity.Packet
	rows, err := repo.db.Query("SELECT id, name, interval FROM tbl_packet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var packet entity.Packet
		err := rows.Scan(&packet.Id, &packet.Name, &packet.Interval)
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)
	}

	return packets, nil
}

// FindPacketByName implements PacketRepository
func (repo *packetRepository) FindPacketByName(name string) (*entity.Packet, error) {
	var packet entity.Packet
	stmt, err := repo.db.Prepare("SELECT id, name, interval FROM tbl_packet WHERE name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&packet.Id, &packet.Name, &packet.Interval)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("packet with name %s not found", name)
	} else if err != nil {
		return nil, err
	}

	return &packet, nil
}

// FindpacketById implements PacketRepository
func (repo *packetRepository) FindpacketById(id int) (*entity.Packet, error) {
	var packet entity.Packet
	stmt, err := repo.db.Prepare("SELECT id, name, interval FROM tbl_packet WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&packet.Id, &packet.Name, &packet.Interval)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("packet with id %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &packet, nil
}

// InsertPacket implements PacketRepository
func (repo *packetRepository) InsertPacket(newPacket *entity.Packet) (*entity.Packet, error) {
	stmt, err := repo.db.Prepare("INSERT INTO tbl_packet (name, interval) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return nil, fmt.Errorf("failed to insert packet : %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newPacket.Name, newPacket.Interval).Scan(&newPacket.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert packet : %w", err)
	}

	return newPacket, nil
}

// UpdatePacket implements PacketRepository
func (repo *packetRepository) UpdatePacket(updatePacket *entity.Packet) (*entity.Packet, error) {
	stmt, err := repo.db.Prepare("UPDATE tbl_packet SET name = $1, interval = $2 WHERE id = $3")
	if err != nil {
		return nil, fmt.Errorf("failed to update packet : %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatePacket.Name, updatePacket.Interval, updatePacket.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update packet : %w", err)
	}

	return updatePacket, nil
}

func NewPacketRepository(db *sql.DB) PacketRepository {
	return &packetRepository{db: db}
}
