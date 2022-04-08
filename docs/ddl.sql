create table users (
	user_id varchar(50) primary key not null,
    username varchar(255),
    created_at timestamp,
    updated_at timestamp
)

create table ticket (
	ticket_id varchar(20) primary key not null,
    acara varchar(50),
    harga decimal (38,2),
    created_at timestamp,
    updated_at timestamp
)

create table checkout (
	id serial primary key not null,
	checkout_id varchar(20),
    user_id varchar(50),
    ticket_id varchar(20),
    is_purchased boolean,
    created_at timestamp,
    updated_at timestamp
)

CREATE UNIQUE INDEX user_id on users(user_id);
CREATE UNIQUE INDEX ticket_id on ticket(ticket_id);
CREATE INDEX checkout_id on checkout(checkout_id);