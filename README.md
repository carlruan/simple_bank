T -> Possible, F -> Impossible

| Mysql isolation levels | read uncommitted | read committed | repeatable read | serializable |
|:----------------------:|:----------------:|:--------------:|:---------------:|:------------:|
|       dirty read       |        T         |       F        |        F        |      F       |
|  non-repeatable read   |        T         |       T        |        F        |      F       |
|      phantom read      |        T         |       T        |        F        |      F       |
| serialization anomaly  |        T         |       T        |        T        |      F       |

| Postgres isolation levels | read uncommitted | read committed | repeatable read | serializable |
|:-------------------------:|:----------------:|:--------------:|:---------------:|:------------:|
|        dirty read         |        F         |       F        |        F        |      F       |
|    non-repeatable read    |        T         |       T        |        F        |      F       |
|       phantom read        |        T         |       T        |        F        |      F       |
|   serialization anomaly   |        T         |       T        |        T        |      F       |
