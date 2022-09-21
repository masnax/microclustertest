package database

//go:generate -command mapper lxd-generate db mapper -t extended.mapper.go
//go:generate mapper reset
//
//go:generate mapper stmt -d github.com/canonical/microcluster/cluster -e extended_table objects table=extended_table
//go:generate mapper stmt -d github.com/canonical/microcluster/cluster -e extended_table objects-by-Key table=extended_table
//go:generate mapper stmt -d github.com/canonical/microcluster/cluster -e extended_table id table=extended_table
//go:generate mapper stmt -d github.com/canonical/microcluster/cluster -e extended_table create table=extended_table
//go:generate mapper stmt -d github.com/canonical/microcluster/cluster -e extended_table delete-by-Key table=extended_table
//go:generate mapper stmt -d github.com/canonical/microcluster/cluster -e extended_table update table=extended_table
//
//go:generate mapper method -i -d github.com/canonical/microcluster/cluster -e extended_table GetMany
//go:generate mapper method -i -d github.com/canonical/microcluster/cluster -e extended_table GetOne
//go:generate mapper method -i -d github.com/canonical/microcluster/cluster -e extended_table ID
//go:generate mapper method -i -d github.com/canonical/microcluster/cluster -e extended_table Exists
//go:generate mapper method -i -d github.com/canonical/microcluster/cluster -e extended_table Create
//go:generate mapper method -i -d github.com/canonical/microcluster/cluster -e extended_table DeleteOne-by-Key
//go:generate mapper method -i -d github.com/canonical/microcluster/cluster -e extended_table Update

// ExtendedTable is an example of a database table. In this case named `extended_table`. The above comments will
// generate database queries and helpers using lxd-generate.
type ExtendedTable struct {
	ID    int
	Key   string `db:"primary=yes"`
	Value string
}

// ExtendedTableFilter is a required struct for use with lxd-generate. It is used for filtering fields on database
// fetches. In this case we will only support filtering by Key.
type ExtendedTableFilter struct {
	Key *string
}