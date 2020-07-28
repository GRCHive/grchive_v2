import { v4 as uuidv4 } from 'uuid'

export enum DatabaseType {
    Other = 1,
    PostgreSQL = 2,
    MySQL = 3,
    MariaDB = 4,
    SAPHana = 5,
    IBMDb2 = 6,
    Oracle = 7,
    MsSql = 8,
    MsAccess = 9,
    SQLite = 10,
    H2 = 11
}

export function databaseTypeToString(typ : DatabaseType) : string {
    switch (typ) {
        case DatabaseType.PostgreSQL:
            return 'PostgreSQL'
        case DatabaseType.MySQL:
            return 'MySQL'
        case DatabaseType.MariaDB:
            return 'MariaDB'
        case DatabaseType.SAPHana:
            return 'SAP HANA'
        case DatabaseType.IBMDb2:
            return 'IBM DB2'
        case DatabaseType.Oracle:
            return 'Oracle'
        case DatabaseType.MsSql:
            return 'Microsoft SQL'
        case DatabaseType.MsAccess:
            return 'Microsoft Access'
        case DatabaseType.SQLite:
            return 'SQLite'
        case DatabaseType.H2:
            return 'H2'
        default: 
            return 'Unknown'
    }
}

export interface RawDatabase {
    Id                  : number
    EngagementId        : number
    UniqueId            : string
    Name                : string
    Description         : string
    TypeId              : DatabaseType
    OtherType           : string | null
    Version             : string
}

export function createEmptyDatabase() : RawDatabase {
    return {
        Id: -1,
        EngagementId: -1,
        UniqueId: uuidv4(),
        Name: '',
        Description: '',
        TypeId: DatabaseType.Other,
        OtherType: null,
        Version: ''
    }
}
