CREATE KEYSPACE IF NOT EXISTS testkeyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1};

USE testkeyspace;

CREATE TABLE IF NOT EXISTS events (
                                      bucket BIGINT,
                                      action TEXT,
                                      actor TEXT,
                                      timestamp TIMESTAMP,
                                      event_id TEXT,
                                      version TEXT,
                                      action_metadata MAP<TEXT, TEXT>,
                                      actor_metadata MAP<TEXT, TEXT>,
                                      PRIMARY KEY (bucket, timestamp, event_id)
);

CREATE KEYSPACE IF NOT EXISTS logger WITH replication = {'class': 'SimpleStrategy', 'replication_factor' : 1}
AND durable_writes = true;

USE logger;

CREATE TABLE IF NOT EXISTS events (
                                      bucket BIGINT,
                                      action TEXT,
                                      actor TEXT,
                                      timestamp TIMESTAMP,
                                      event_id TEXT,
                                      version TEXT,
                                      action_metadata MAP<TEXT, TEXT>,
                                      actor_metadata MAP<TEXT, TEXT>,
                                      PRIMARY KEY (bucket, timestamp, event_id)
) WITH CLUSTERING ORDER BY (timestamp DESC, event_id ASC);