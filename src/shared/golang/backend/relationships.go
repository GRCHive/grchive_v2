package backend

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type RelationshipWrapper struct {
	Id       int64       `db:"id"`
	Explicit bool        `db:"explicit"`
	Data     interface{} `db:"data"`
}

func resourceRelationshipToTableName(from ResourceIdentifier, to ResourceIdentifier) (string, error) {
	if from > to {
		return resourceRelationshipToTableName(to, from)
	}

	switch from {
	case RIRisk:
		switch to {
		case RIControl:
			return "risks_controls", nil
		}
	}

	return "", errors.New("Unsupported relationship.")
}

func (b *BackendInterface) ListRelationships(from ResourceIdentifier, fromId int64, to ResourceIdentifier) ([]RelationshipWrapper, error) {
	relationshipTable, err := resourceRelationshipToTableName(from, to)
	if err != nil {
		return nil, err
	}

	resourceTable, err := resourceToTableName(to)
	if err != nil {
		return nil, err
	}

	fromFk, err := resourceToFKDatabaseId(from)
	if err != nil {
		return nil, err
	}

	toFk, err := resourceToFKDatabaseId(to)
	if err != nil {
		return nil, err
	}

	rows, err := b.db.Queryx(fmt.Sprintf(`
		SELECT res.*
		FROM %[2]s AS res
		INNER JOIN %[1]s AS rel
			ON rel.%[4]s = res.id
		WHERE rel.%[3]s = $1
		ORDER BY rel.id DESC
	`, relationshipTable, resourceTable, fromFk, toFk), fromId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ret := []RelationshipWrapper{}
	for rows.Next() {
		wrapper := RelationshipWrapper{
			Explicit: true,
		}
		wrapper.Data, err = CreateResource(to)
		if err != nil {
			return nil, err
		}

		err = rows.StructScan(wrapper.Data)
		if err != nil {
			return nil, err
		}

		ret = append(ret, wrapper)
	}

	return ret, nil
}

func (b *BackendInterface) CreateRelationship(
	tx *sqlx.Tx,
	from ResourceIdentifier, fromId int64,
	to ResourceIdentifier, toId int64,
) error {
	relationshipTable, err := resourceRelationshipToTableName(from, to)
	if err != nil {
		return err
	}

	fromFk, err := resourceToFKDatabaseId(from)
	if err != nil {
		return err
	}

	toFk, err := resourceToFKDatabaseId(to)
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		INSERT INTO %[1]s (%[2]s, %[3]s)
		VALUES ($1, $2)
	`, relationshipTable, fromFk, toFk), fromId, toId)
	return err
}

func (b *BackendInterface) DeleteRelationship(
	tx *sqlx.Tx,
	from ResourceIdentifier, fromId int64,
	to ResourceIdentifier, toId int64,
) error {
	relationshipTable, err := resourceRelationshipToTableName(from, to)
	if err != nil {
		return err
	}

	fromFk, err := resourceToFKDatabaseId(from)
	if err != nil {
		return err
	}

	toFk, err := resourceToFKDatabaseId(to)
	if err != nil {
		return err
	}

	_, err = tx.Exec(fmt.Sprintf(`
		DELETE FROM %[1]s
		WHERE
			%[2]s = $1 AND
			%[3]s = $2
	`, relationshipTable, fromFk, toFk), fromId, toId)
	return err
}
