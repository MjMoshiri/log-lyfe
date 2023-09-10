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
