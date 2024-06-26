CREATE TABLE IF NOT EXISTS dialog_table
(
    id         TEXT                     NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,

    CONSTRAINT pk_dialog_table PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS participant_table
(
    id             TEXT                     NOT NULL,
    dialog_id      TEXT                     NOT NULL,

    user_id TEXT                     NOT NULL,

    created_at     TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at     TIMESTAMP WITH TIME ZONE,

    CONSTRAINT pk_participant_table PRIMARY KEY (id),
    CONSTRAINT fk_participant_table_dialog_table FOREIGN KEY (dialog_id) REFERENCES dialog_table (id)
);

CREATE TABLE IF NOT EXISTS message_table
(
    id         TEXT                     NOT NULL,
    dialog_id  TEXT                     NOT NULL,

    sender_id  TEXT                     NOT NULL,
    text       TEXT                     NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,

    CONSTRAINT pk_message_table PRIMARY KEY (id),
    CONSTRAINT fk_message_table_dialog_table FOREIGN KEY (dialog_id) REFERENCES dialog_table (id)
);