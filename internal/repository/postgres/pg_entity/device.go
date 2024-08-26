package pgEntity

import (
	"github.com/balobas/auth_service/internal/entity"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const devicesTableName = "user_devices"

type DeviceRow struct {
	Uid         pgtype.UUID
	UserUid     pgtype.UUID
	Name        string
	OS          string
	ConnectedAt pgtype.Timestamp
}

func NewDeviceRow() *DeviceRow {
	return &DeviceRow{}
}

var devicesTableColumns = []string{
	"uid",
	"user_uid",
	"name",
	"os",
	"connected_at",
}

func (d *DeviceRow) FromEntity(device entity.UserDevice) *DeviceRow {
	d.Uid = pgtype.UUID{
		Bytes:  device.Uid,
		Status: pgtype.Present,
	}
	d.UserUid = pgtype.UUID{
		Bytes:  device.UserUid,
		Status: pgtype.Present,
	}
	d.Name = device.Name
	d.OS = device.OS
	if device.ConnectedAt.Unix() == 0 {
		d.ConnectedAt = pgtype.Timestamp{
			Status: pgtype.Null,
		}
	} else {
		d.ConnectedAt = pgtype.Timestamp{
			Time:   device.ConnectedAt,
			Status: pgtype.Present,
		}
	}
	return d
}

func (d *DeviceRow) ToEntity() entity.UserDevice {
	return entity.UserDevice{
		Uid:         d.Uid.Bytes,
		UserUid:     d.UserUid.Bytes,
		Name:        d.Name,
		OS:          d.OS,
		ConnectedAt: d.ConnectedAt.Time,
	}
}

func (d *DeviceRow) IdColumnName() string {
	return "uid"
}

func (d *DeviceRow) Values() []interface{} {
	return []interface{}{
		d.Uid,
		d.UserUid,
		d.Name,
		d.OS,
		d.ConnectedAt,
	}
}

func (d *DeviceRow) Columns() []string {
	return devicesTableColumns
}

func (d *DeviceRow) Table() string {
	return devicesTableName
}

func (d *DeviceRow) GetId() interface{} {
	return d.Uid
}

func (d *DeviceRow) ScanId(row pgx.Row) error {
	return row.Scan(&d.Uid)
}

func (d *DeviceRow) Scan(row pgx.Row) error {
	return row.Scan(&d.Uid, &d.UserUid, &d.Name, &d.OS, &d.ConnectedAt)
}

func (d *DeviceRow) ColumnsForUpdate() []string {
	return nil
}

func (d *DeviceRow) ValuesForUpdate() []interface{} {
	return nil
}
